//go:build linux

/*
Copyright 2026 The Kubernetes Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package substrate

import (
	"context"
	"net"
	"os"
	"path/filepath"
	"sync/atomic"
	"testing"
	"time"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	mount "k8s.io/mount-utils"
)

func TestGoldenPlaceholderRealMountThroughGRPC(t *testing.T) {
	if os.Getenv("BRIDGE_REAL_MOUNT_TEST") != "1" {
		t.Skip("requires an explicitly enabled isolated privileged Linux test")
	}
	require.Zero(t, os.Geteuid())
	root, state := t.TempDir(), t.TempDir()
	target := filepath.Join(root, testUID, "volumes", "data")
	var unavailable atomic.Bool
	lookup := func(context.Context, ActorReference) (ActorInfo, error) {
		if unavailable.Load() {
			return ActorInfo{}, status.Error(codes.Unavailable, "Actor API unavailable")
		}
		return ActorInfo{UID: testUID, Atespace: "ate-golden", Name: "template-uid", TemplateUID: "template-uid", Golden: true}, nil
	}
	node := NewNode(NodeOptions{ActorRoot: root, StateDir: state, Lookup: lookup, Mounter: mount.NewWithoutSystemd("")})
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	require.NoError(t, err)
	server := grpc.NewServer()
	csi.RegisterControllerServer(server, &Controller{})
	csi.RegisterNodeServer(server, node)
	done := make(chan error, 1)
	go func() { done <- server.Serve(listener) }()
	t.Cleanup(func() { server.Stop(); require.NoError(t, <-done) })
	conn, err := grpc.NewClient(listener.Addr().String(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err)
	t.Cleanup(func() { require.NoError(t, conn.Close()) })
	created, err := csi.NewControllerClient(conn).CreateVolume(t.Context(), &csi.CreateVolumeRequest{Name: testID, Parameters: goldenMetadataFixture(), VolumeCapabilities: []*csi.VolumeCapability{publishFixture().VolumeCapability}})
	require.NoError(t, err)
	vc := created.Volume.VolumeContext
	vc[PodUIDKey] = "worker-pod-uid"
	client := csi.NewNodeClient(conn)
	in := &csi.NodePublishVolumeRequest{VolumeId: testID, TargetPath: target, VolumeContext: vc}
	_, err = client.NodePublishVolume(t.Context(), in)
	require.NoError(t, err)
	t.Cleanup(func() {
		_, cleanupErr := client.NodeUnpublishVolume(context.Background(), &csi.NodeUnpublishVolumeRequest{VolumeId: testID, TargetPath: target})
		require.NoError(t, cleanupErr)
	})
	in.VolumeId = "publish-alias"
	_, err = client.NodePublishVolume(t.Context(), in)
	require.NoError(t, err)
	require.NoError(t, os.WriteFile(filepath.Join(target, "probe"), []byte("golden-placeholder"), 0600))
	source, err := node.placeholderSource(target)
	require.NoError(t, err)
	read, err := os.ReadFile(filepath.Join(source, "probe"))
	require.NoError(t, err)
	require.Equal(t, []byte("golden-placeholder"), read)
	unavailable.Store(true)
	_, err = client.NodeUnpublishVolume(t.Context(), &csi.NodeUnpublishVolumeRequest{VolumeId: "cleanup-alias", TargetPath: target})
	require.NoError(t, err)
	require.NoDirExists(t, target)
	require.NoDirExists(t, source)
}

func TestGoldenRealMountIgnoresOuterReadonlyAndStaysWritable(t *testing.T) {
	if os.Getenv("BRIDGE_REAL_MOUNT_TEST") != "1" {
		t.Skip("set BRIDGE_REAL_MOUNT_TEST=1 inside an isolated privileged Linux container")
	}
	root := t.TempDir()
	target := filepath.Join(root, "actors", testUID, "volumes", "data")
	lookup := func(context.Context, ActorReference) (ActorInfo, error) {
		return ActorInfo{UID: testUID, Atespace: "ate-golden", Name: "template-uid", TemplateUID: "template-uid", Golden: true}, nil
	}
	attributes := goldenMetadataFixture()
	attributes[PodUIDKey] = "worker-pod-uid"
	node := NewNode(NodeOptions{ActorRoot: filepath.Join(root, "actors"), StateDir: filepath.Join(root, "state"), Lookup: lookup, Mounter: mount.New("")})
	_, err := node.NodePublishVolume(t.Context(), &csi.NodePublishVolumeRequest{VolumeId: testID, TargetPath: target, Readonly: true, VolumeContext: attributes})
	require.NoError(t, err)
	t.Cleanup(func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, cleanupErr := node.NodeUnpublishVolume(cleanupCtx, &csi.NodeUnpublishVolumeRequest{VolumeId: testID, TargetPath: target})
		require.NoError(t, cleanupErr)
	})
	require.NoError(t, os.WriteFile(filepath.Join(target, "probe"), []byte("writable placeholder"), 0600))
}

func TestGoldenRealStackedMountOwnership(t *testing.T) {
	if os.Getenv("BRIDGE_REAL_MOUNT_TEST") != "1" {
		t.Skip("requires an explicitly enabled isolated privileged Linux test")
	}
	for _, scenario := range []string{"covered publish", "covered unpublish", "lower mount remains"} {
		t.Run(scenario, func(t *testing.T) {
			root, state := t.TempDir(), t.TempDir()
			target := filepath.Join(root, testUID, "volumes", "data")
			mounter := mount.NewWithoutSystemd("")
			node := NewNode(NodeOptions{ActorRoot: root, StateDir: state, Mounter: mounter})
			source, err := node.placeholderSource(target)
			require.NoError(t, err)
			require.NoError(t, os.MkdirAll(source, 0755))
			require.NoError(t, os.MkdirAll(target, 0755))
			require.NoError(t, os.WriteFile(filepath.Join(source, "retain"), []byte("source data"), 0600))
			t.Cleanup(func() {
				for range 2 {
					live, err := node.mountAt(target)
					require.NoError(t, err)
					if live == nil {
						return
					}
					require.NoError(t, mounter.Unmount(target))
				}
			})
			if scenario == "lower mount remains" {
				require.NoError(t, mounter.Mount("tmpfs", target, "tmpfs", nil))
				require.NoError(t, os.WriteFile(filepath.Join(target, "foreign"), []byte("foreign data"), 0600))
				require.NoError(t, mounter.Mount(source, target, "", []string{"bind"}))
			} else {
				require.NoError(t, mounter.Mount(source, target, "", []string{"bind"}))
				require.NoError(t, mounter.Mount("tmpfs", target, "tmpfs", nil))
				require.NoError(t, os.WriteFile(filepath.Join(target, "foreign"), []byte("foreign data"), 0600))
			}

			if scenario == "covered publish" {
				err = node.publishPlaceholder(target)
			} else {
				_, err = node.NodeUnpublishVolume(t.Context(), &csi.NodeUnpublishVolumeRequest{VolumeId: testID, TargetPath: target})
			}
			require.Equal(t, codes.FailedPrecondition, status.Code(err))
			live, err := node.mountAt(target)
			require.NoError(t, err)
			require.NotNil(t, live)
			require.Equal(t, "tmpfs", live.Type)
			foreign, err := os.ReadFile(filepath.Join(target, "foreign"))
			require.NoError(t, err)
			require.Equal(t, "foreign data", string(foreign))
			retained, err := os.ReadFile(filepath.Join(source, "retain"))
			require.NoError(t, err)
			require.Equal(t, "source data", string(retained))
		})
	}
}
