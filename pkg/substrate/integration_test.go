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

type readonlyObservedNAS struct {
	csi.UnimplementedNodeServer
	published chan bool
}

func (n *readonlyObservedNAS) NodePublishVolume(_ context.Context, req *csi.NodePublishVolumeRequest) (*csi.NodePublishVolumeResponse, error) {
	n.published <- req.Readonly
	return &csi.NodePublishVolumeResponse{}, nil
}

func (*readonlyObservedNAS) NodeUnpublishVolume(context.Context, *csi.NodeUnpublishVolumeRequest) (*csi.NodeUnpublishVolumeResponse, error) {
	return &csi.NodeUnpublishVolumeResponse{}, nil
}

func TestReadonlyPublishLifecycleThroughGRPC(t *testing.T) {
	ctx, cancel := context.WithTimeout(t.Context(), 10*time.Second)
	t.Cleanup(cancel)
	node, _, in := readonlyNodeFixture(t, writablePublishFixture())
	readonlyAnnotation := annotationFixture(t, publishFixture())
	writableAnnotation := annotationFixture(t, writablePublishFixture())
	var innerReadonly atomic.Bool
	innerReadonly.Store(true)
	var lookups atomic.Int32
	node.opts.Lookup = func(context.Context, ActorReference) (ActorInfo, error) {
		lookups.Add(1)
		annotation := writableAnnotation
		if innerReadonly.Load() {
			annotation = readonlyAnnotation
		}
		return ActorInfo{UID: testUID, Name: "actor", Atespace: "storage-test", Annotation: annotation}, nil
	}
	nas := &readonlyObservedNAS{published: make(chan bool, 1)}
	node.opts.Backends = nasBackends(nas)
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
	created, err := csi.NewControllerClient(conn).CreateVolume(ctx, &csi.CreateVolumeRequest{
		Name: testID, Parameters: actorMetadataFixture(), VolumeCapabilities: []*csi.VolumeCapability{in.VolumeCapability},
	})
	require.NoError(t, err)
	in.VolumeContext = created.Volume.VolumeContext
	require.Zero(t, lookups.Load())
	in.VolumeContext[PodUIDKey] = "worker-pod-uid"
	client := csi.NewNodeClient(conn)
	// The Actor request is readonly, so the volume publishes readonly with or without the
	// outer request asking for it.
	in.Readonly = true
	_, err = client.NodePublishVolume(ctx, in)
	require.NoError(t, err)
	require.Equal(t, int32(1), lookups.Load())
	require.Len(t, nas.published, 1)
	require.True(t, <-nas.published)
	in.Readonly = false
	_, err = client.NodePublishVolume(ctx, in)
	require.NoError(t, err)
	require.Equal(t, int32(2), lookups.Load())
	require.True(t, <-nas.published)
	in.VolumeCapability.GetMount().MountFlags = []string{"ro"}
	_, err = client.NodePublishVolume(ctx, in)
	require.NoError(t, err)
	require.True(t, <-nas.published)
	in.VolumeCapability.GetMount().MountFlags = nil
	_, err = client.NodePublishVolume(ctx, in)
	require.NoError(t, err)
	require.True(t, <-nas.published)
	// A writable Actor request publishes writable: the outer request is not consulted, in
	// either direction.
	innerReadonly.Store(false)
	_, err = client.NodePublishVolume(ctx, in)
	require.NoError(t, err)
	require.False(t, <-nas.published)
	in.Readonly = true
	_, err = client.NodePublishVolume(ctx, in)
	require.NoError(t, err)
	require.False(t, <-nas.published)
	in.Readonly = false
	_, err = client.NodeUnpublishVolume(ctx, &csi.NodeUnpublishVolumeRequest{VolumeId: testID, TargetPath: testTarget})
	require.NoError(t, err)
	_, err = client.NodePublishVolume(ctx, in)
	require.NoError(t, err)
	require.Len(t, nas.published, 1)
	require.False(t, <-nas.published)
}

type waitingNAS struct {
	csi.UnimplementedNodeServer
	entered, release chan struct{}
}

func (n *waitingNAS) NodePublishVolume(context.Context, *csi.NodePublishVolumeRequest) (*csi.NodePublishVolumeResponse, error) {
	close(n.entered)
	<-n.release
	return &csi.NodePublishVolumeResponse{}, nil
}

func TestStatelessTargetOperationsStaySerialized(t *testing.T) {
	node, _, in := readonlyNodeFixture(t, writablePublishFixture())
	nas := &waitingNAS{entered: make(chan struct{}), release: make(chan struct{})}
	node.opts.Backends = nasBackends(nas)
	done := make(chan error, 1)
	go func() { _, err := node.NodePublishVolume(t.Context(), in); done <- err }()
	<-nas.entered
	_, err := node.NodeUnpublishVolume(t.Context(), &csi.NodeUnpublishVolumeRequest{VolumeId: testID, TargetPath: testTarget})
	require.Equal(t, codes.Aborted, status.Code(err))
	close(nas.release)
	require.NoError(t, <-done)
}

func TestStatelessPublishAcceptsCorrectionAfterFailureAndRestart(t *testing.T) {
	real := writablePublishFixture()
	annotation := annotationFixture(t, real)
	node, downstream, in := readonlyNodeFixture(t, real)
	node.opts.Lookup = func(context.Context, ActorReference) (ActorInfo, error) {
		return ActorInfo{UID: testUID, Name: "actor", Atespace: "storage-test", Annotation: annotation}, nil
	}
	downstream.err = status.Error(codes.Unavailable, "temporary failure")
	_, err := node.NodePublishVolume(t.Context(), in)
	require.Equal(t, codes.Unavailable, status.Code(err))
	real.VolumeContext["path"] = "/corrected"
	annotation = annotationFixture(t, real)
	downstream.err = nil
	node = NewNode(node.opts)
	_, err = node.NodePublishVolume(t.Context(), in)
	require.NoError(t, err)
	require.Equal(t, "/corrected", downstream.published.VolumeContext["path"])
}

func TestFirstPublishUsesAnnotationChangedAfterCreate(t *testing.T) {
	annotation := annotationFixture(t, publishFixture())
	lookup := func(context.Context, ActorReference) (ActorInfo, error) {
		return ActorInfo{UID: testUID, Atespace: "storage-test", Name: "actor", Annotation: annotation}, nil
	}
	created, err := (&Controller{}).CreateVolume(t.Context(), &csi.CreateVolumeRequest{Name: testID, Parameters: actorMetadataFixture(), VolumeCapabilities: []*csi.VolumeCapability{publishFixture().VolumeCapability}})
	require.NoError(t, err)
	changed := publishFixture()
	changed.VolumeContext["path"] = "/another-tenant"
	annotation = annotationFixture(t, changed)
	nas := &recordingNAS{}
	node := NewNode(NodeOptions{ActorRoot: DefaultActorRoot, StateDir: t.TempDir(), Mounter: mount.NewFakeMounter(nil), Lookup: lookup, Backends: nasBackends(nas)})
	context := created.Volume.VolumeContext
	context[PodUIDKey] = "worker-pod-uid"
	_, err = node.NodePublishVolume(t.Context(), &csi.NodePublishVolumeRequest{VolumeId: testID, TargetPath: testTarget, VolumeContext: context})
	require.NoError(t, err)
	require.Equal(t, "/another-tenant", nas.published.VolumeContext["path"])
}

func TestCreateSucceedsBeforePublishConfigurationIsAvailable(t *testing.T) {
	for _, tc := range []struct {
		name       string
		annotation string
		lookupErr  error
		wantCode   codes.Code
	}{
		{"missing annotation", "", nil, codes.FailedPrecondition},
		{"invalid annotation", "not JSON", nil, codes.FailedPrecondition},
		{"API unavailable", "", status.Error(codes.Unavailable, "API unavailable"), codes.Unavailable},
	} {
		t.Run(tc.name, func(t *testing.T) {
			created, err := (&Controller{}).CreateVolume(t.Context(), &csi.CreateVolumeRequest{Name: testID, Parameters: actorMetadataFixture(), VolumeCapabilities: []*csi.VolumeCapability{publishFixture().VolumeCapability}})
			require.NoError(t, err)
			calls := 0
			nas := &recordingNAS{}
			node := NewNode(NodeOptions{ActorRoot: DefaultActorRoot, StateDir: t.TempDir(), Mounter: mount.NewFakeMounter(nil), Backends: nasBackends(nas), Lookup: func(context.Context, ActorReference) (ActorInfo, error) {
				calls++
				return ActorInfo{UID: testUID, Name: "actor", Atespace: "storage-test", Annotation: tc.annotation}, tc.lookupErr
			}})
			created.Volume.VolumeContext[PodUIDKey] = "worker-pod-uid"
			_, err = node.NodePublishVolume(t.Context(), &csi.NodePublishVolumeRequest{VolumeId: testID, TargetPath: testTarget, VolumeContext: created.Volume.VolumeContext})
			require.Equal(t, tc.wantCode, status.Code(err))
			require.Equal(t, 1, calls)
			require.Nil(t, nas.published)
		})
	}
}

func TestPublishUsesCurrentConfigurationWithoutBindingConflicts(t *testing.T) {
	for _, failedFirst := range []bool{false, true} {
		t.Run(map[bool]string{false: "mounted", true: "failed publish"}[failedFirst], func(t *testing.T) {
			real := writablePublishFixture()
			annotation := annotationFixture(t, real)
			node, nas, request := readonlyNodeFixture(t, real)
			node.opts.Lookup = func(context.Context, ActorReference) (ActorInfo, error) {
				return ActorInfo{UID: testUID, Name: "actor", Atespace: "storage-test", Annotation: annotation}, nil
			}
			if failedFirst {
				nas.err = status.Error(codes.Unavailable, "NAS unavailable")
			}
			_, err := node.NodePublishVolume(t.Context(), request)
			if failedFirst {
				require.Equal(t, codes.Unavailable, status.Code(err))
			} else {
				require.NoError(t, err)
				_, err = node.NodePublishVolume(t.Context(), request)
				require.NoError(t, err)
			}
			node = NewNode(node.opts)
			nas.err, nas.published = nil, nil
			real.VolumeContext["path"] = "/changed-after-binding"
			annotation = annotationFixture(t, real)
			_, err = node.NodePublishVolume(t.Context(), request)
			require.NoError(t, err)
			require.Equal(t, "/changed-after-binding", nas.published.VolumeContext["path"])
			_, err = node.NodeUnpublishVolume(t.Context(), &csi.NodeUnpublishVolumeRequest{VolumeId: testID, TargetPath: testTarget})
			require.NoError(t, err)
			_, err = node.NodePublishVolume(t.Context(), request)
			require.NoError(t, err)
			require.Equal(t, "/changed-after-binding", nas.published.VolumeContext["path"])
		})
	}
}
