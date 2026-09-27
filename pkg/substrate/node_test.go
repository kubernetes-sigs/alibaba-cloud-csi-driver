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
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	mount "k8s.io/mount-utils"
)

const testUID = "1c3d0394-02df-4aa9-b9ec-27604d5b7601"

const testID = "substrate-" + testUID + "-data"

const testTarget = "/var/lib/ateom-gvisor/actors/" + testUID + "/volumes/data"

type recordingNAS struct {
	csi.UnimplementedNodeServer
	published   *csi.NodePublishVolumeRequest
	unpublished *csi.NodeUnpublishVolumeRequest
	err         error
	mounts      func() mount.Interface
}

func (n *recordingNAS) NodePublishVolume(_ context.Context, req *csi.NodePublishVolumeRequest) (*csi.NodePublishVolumeResponse, error) {
	n.published = req
	if n.err == nil && n.mounts != nil {
		if err := fakeNASMount(n.mounts(), req); err != nil {
			return nil, err
		}
	}
	return &csi.NodePublishVolumeResponse{}, n.err
}

func (n *recordingNAS) NodeUnpublishVolume(_ context.Context, req *csi.NodeUnpublishVolumeRequest) (*csi.NodeUnpublishVolumeResponse, error) {
	n.unpublished = req
	if n.err == nil && n.mounts != nil {
		if err := n.mounts().Unmount(req.TargetPath); err != nil {
			return nil, err
		}
	}
	return &csi.NodeUnpublishVolumeResponse{}, n.err
}

func fakeNASMount(m mount.Interface, req *csi.NodePublishVolumeRequest) error {
	entries, err := m.List()
	if err != nil {
		return err
	}
	for _, entry := range entries {
		if entry.Path == req.TargetPath {
			return nil
		}
	}
	mode := "rw"
	if req.GetReadonly() {
		mode = "ro"
	}
	return m.Mount("server:/", req.TargetPath, "nfs", []string{mode})
}

// nasBackends registers a single NAS backend, the only driver the bridge ships
// with, so tests exercise the same routing production does.
func nasBackends(node csi.NodeServer) map[string]Backend {
	return map[string]Backend{NASDriverName: NASBackend(node)}
}

func publishFixture() *csi.NodePublishVolumeRequest {
	return &csi.NodePublishVolumeRequest{
		VolumeId: "customer-pv-a1b2c3", TargetPath: "/data", Readonly: true,
		VolumeCapability: &csi.VolumeCapability{
			AccessType: &csi.VolumeCapability_Mount{Mount: &csi.VolumeCapability_MountVolume{FsType: "nfs", MountFlags: []string{"tls", "vers=3"}}},
			AccessMode: &csi.VolumeCapability_AccessMode{Mode: csi.VolumeCapability_AccessMode_SINGLE_NODE_READER_ONLY},
		},
		VolumeContext:  map[string]string{"authType": "agent-identity", "mountProtocol": "alinas", "server": "ap-example.fs.cn-beijing.agenticfs.aliyuncs.com", "path": "/tenant-a", "sandboxCredProviderName": "read-only"},
		PublishContext: map[string]string{"preserved": "value"},
	}
}

func annotationFixture(t *testing.T, request *csi.NodePublishVolumeRequest) string {
	t.Helper()
	raw, err := protojson.Marshal(request)
	require.NoError(t, err)
	encoded, err := json.Marshal([]publishEntry{{VolumeName: "data", Driver: NASDriverName, Request: raw}})
	require.NoError(t, err)
	return string(encoded)
}

func writablePublishFixture() *csi.NodePublishVolumeRequest {
	req := publishFixture()
	req.Readonly = false
	req.VolumeCapability.AccessMode.Mode = csi.VolumeCapability_AccessMode_SINGLE_NODE_WRITER
	return req
}

func readonlyNodeFixture(t *testing.T, real *csi.NodePublishVolumeRequest) (*Node, *recordingNAS, *csi.NodePublishVolumeRequest) {
	t.Helper()
	annotation := annotationFixture(t, real)
	downstream := &recordingNAS{}
	node := NewNode(NodeOptions{
		ActorRoot: DefaultActorRoot, StateDir: t.TempDir(), Backends: nasBackends(downstream),
		Mounter: mount.NewFakeMounter(nil),
		Lookup: func(context.Context, ActorReference) (ActorInfo, error) {
			return ActorInfo{UID: testUID, Atespace: "storage-test", Name: "actor", Annotation: annotation}, nil
		},
	})
	downstream.mounts = func() mount.Interface { return node.opts.Mounter }
	return node, downstream, &csi.NodePublishVolumeRequest{
		VolumeId: testID, TargetPath: testTarget,
		VolumeCapability: writablePublishFixture().VolumeCapability,
		VolumeContext:    nodeContextFixture(),
	}
}

func actorMetadataFixture() map[string]string {
	return map[string]string{
		"csi.alibabacloud.com/actor.uid":       testUID,
		"csi.alibabacloud.com/actor.name":      "actor",
		"csi.alibabacloud.com/actor.namespace": "storage-test",
	}
}

func goldenMetadataFixture() map[string]string {
	attributes := actorMetadataFixture()
	attributes["csi.alibabacloud.com/actor.name"] = "template-uid"
	attributes["csi.alibabacloud.com/actor.namespace"] = "ate-golden"
	return attributes
}

func nodeContextFixture() map[string]string {
	attributes := actorMetadataFixture()
	attributes[PodUIDKey] = "worker-pod-uid"
	return attributes
}

func identifiedLookup(t *testing.T) ActorLookup {
	t.Helper()
	annotation := annotationFixture(t, writablePublishFixture())
	return func(context.Context, ActorReference) (ActorInfo, error) {
		return ActorInfo{UID: testUID, Name: "actor", Atespace: "storage-test", Annotation: annotation}, nil
	}
}

func TestPublishPreservesRequestAndUsesActorIdentity(t *testing.T) {
	request := publishFixture()
	annotation := annotationFixture(t, request)
	downstream := &recordingNAS{}
	node := NewNode(NodeOptions{NodeID: "node", ActorRoot: "/var/lib/ateom-gvisor/actors", Backends: nasBackends(downstream), StateDir: t.TempDir(), Mounter: mount.NewFakeMounter(nil),
		Lookup: func(_ context.Context, ref ActorReference) (ActorInfo, error) {
			require.Equal(t, testUID, ref.UID)
			return ActorInfo{UID: ref.UID, Atespace: "storage-test", Name: "actor", Annotation: annotation}, nil
		},
	})
	_, err := node.NodePublishVolume(t.Context(), &csi.NodePublishVolumeRequest{VolumeId: testID, TargetPath: testTarget, VolumeContext: nodeContextFixture()})
	require.NoError(t, err)
	want := proto.Clone(request).(*csi.NodePublishVolumeRequest)
	want.TargetPath = testTarget
	want.VolumeContext["csi.alibabacloud.com/substrate-mode"] = "true"
	for key, value := range nodeContextFixture() {
		want.VolumeContext[key] = value
	}
	require.True(t, proto.Equal(want, downstream.published))
	require.Equal(t, "/data", request.TargetPath)
	require.NotContains(t, request.VolumeContext, "csi.storage.k8s.io/pod.uid")
}

func TestPublishRejectsInvalidBoundary(t *testing.T) {
	for _, tc := range []struct {
		name   string
		mutate func(*csi.NodePublishVolumeRequest, *csi.NodePublishVolumeRequest)
		code   codes.Code
		reason string
	}{
		{"missing actor UID", func(in, _ *csi.NodePublishVolumeRequest) { delete(in.VolumeContext, "csi.alibabacloud.com/actor.uid") }, codes.InvalidArgument, "actor UID must match the target actor directory"},
		{"foreign target", func(in, _ *csi.NodePublishVolumeRequest) { in.TargetPath = "/etc/data" }, codes.InvalidArgument, "target must be inside the actor root"},
		{"actor directory itself", func(in, _ *csi.NodePublishVolumeRequest) { in.TargetPath = "/var/lib/ateom-gvisor/actors/" + testUID }, codes.InvalidArgument, "target must be an actor volume directory"},
		{"non-clean target", func(in, _ *csi.NodePublishVolumeRequest) {
			in.TargetPath = "/var/lib/ateom-gvisor/actors/" + testUID + "/volumes/data/.."
		}, codes.InvalidArgument, "target must be a clean absolute path"},
		{"secrets", func(_, real *csi.NodePublishVolumeRequest) {
			real.Secrets = map[string]string{"key": "must-not-persist"}
		}, codes.FailedPrecondition, "publish request must not carry credentials"},
		{"foreign identity", func(_, real *csi.NodePublishVolumeRequest) { real.VolumeContext["sandboxId"] = "another-actor" }, codes.PermissionDenied, "publish request carries a different actor identity"},
		{"non agent identity", func(_, real *csi.NodePublishVolumeRequest) { real.VolumeContext["authType"] = "access-key" }, codes.FailedPrecondition, "only AgenticFS with Agent Identity is supported"},
		{"non AgenticFS", func(_, real *csi.NodePublishVolumeRequest) { real.VolumeContext["mountProtocol"] = "nfs" }, codes.FailedPrecondition, "only AgenticFS with Agent Identity is supported"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			real := publishFixture()
			in := &csi.NodePublishVolumeRequest{VolumeId: testID, TargetPath: testTarget, VolumeContext: nodeContextFixture()}
			tc.mutate(in, real)
			annotation := annotationFixture(t, real)
			downstream := &recordingNAS{}
			node := NewNode(NodeOptions{ActorRoot: "/var/lib/ateom-gvisor/actors", StateDir: t.TempDir(), Mounter: mount.NewFakeMounter(nil), Backends: nasBackends(downstream), Lookup: func(context.Context, ActorReference) (ActorInfo, error) {
				return ActorInfo{UID: testUID, Name: "actor", Atespace: "storage-test", Annotation: annotation}, nil
			}})
			_, err := node.NodePublishVolume(t.Context(), in)
			require.Equal(t, tc.code, status.Code(err))
			require.Equal(t, tc.reason, status.Convert(err).Message())
			require.Nil(t, downstream.published)
		})
	}
}

func TestUnpublishResolvesIdentityWithoutVolumeContext(t *testing.T) {
	downstream := &recordingNAS{}
	annotation := annotationFixture(t, publishFixture())
	state := t.TempDir()
	node := NewNode(NodeOptions{ActorRoot: "/var/lib/ateom-gvisor/actors", Backends: nasBackends(downstream), StateDir: state, Mounter: mount.NewFakeMounter(nil), Lookup: func(_ context.Context, ref ActorReference) (ActorInfo, error) {
		require.Equal(t, testUID, ref.UID)
		return ActorInfo{UID: ref.UID, Name: "actor", Atespace: "storage-test", Annotation: annotation}, nil
	}})
	_, err := node.NodePublishVolume(t.Context(), &csi.NodePublishVolumeRequest{VolumeId: testID, TargetPath: testTarget, VolumeContext: nodeContextFixture()})
	require.NoError(t, err)
	restarted := NewNode(NodeOptions{ActorRoot: DefaultActorRoot, StateDir: state, Mounter: mount.NewFakeMounter([]mount.MountPoint{{Path: testTarget, Device: "server:/", Type: "nfs"}}), Backends: nasBackends(downstream)})
	_, err = restarted.NodeUnpublishVolume(t.Context(), &csi.NodeUnpublishVolumeRequest{VolumeId: testID, TargetPath: testTarget})
	require.NoError(t, err)
	require.Equal(t, testID, downstream.unpublished.VolumeId)
	require.Equal(t, testTarget, downstream.unpublished.TargetPath)
}

func TestUnpublishDoesNotNeedActorLookup(t *testing.T) {
	downstream := &recordingNAS{}
	state := t.TempDir()
	node := NewNode(NodeOptions{ActorRoot: "/var/lib/ateom-gvisor/actors", Backends: nasBackends(downstream), StateDir: state, Mounter: mount.NewFakeMounter([]mount.MountPoint{{Path: testTarget, Device: "server:/", Type: "nfs"}}), Lookup: func(context.Context, ActorReference) (ActorInfo, error) {
		t.Fatal("unpublish must not call the Actor API")
		return ActorInfo{}, status.Error(codes.Unavailable, "control plane unavailable")
	}})
	_, err := node.NodeUnpublishVolume(t.Context(), &csi.NodeUnpublishVolumeRequest{VolumeId: testID, TargetPath: testTarget})
	require.NoError(t, err)
	require.Equal(t, testID, downstream.unpublished.VolumeId)
}

func TestNodeGetInfoRequiresNodeIdentity(t *testing.T) {
	_, err := NewNode(NodeOptions{}).NodeGetInfo(t.Context(), &csi.NodeGetInfoRequest{})
	require.Equal(t, codes.FailedPrecondition, status.Code(err))
	info, err := NewNode(NodeOptions{NodeID: "worker-node"}).NodeGetInfo(t.Context(), &csi.NodeGetInfoRequest{})
	require.NoError(t, err)
	require.Equal(t, "worker-node", info.NodeId)
}

// Stored PodInfo is not the current placement: the worker changes on resume, and a caller that
// names no Pod at all must not retain a previous worker's labels. Actor ownership is unaffected
// either way, and the bridge caches nothing between the two publishes.
func TestPublishDoesNotRetainStoredWorkerLabels(t *testing.T) {
	real := writablePublishFixture()
	real.VolumeContext[PodUIDKey] = "previous-worker"
	real.VolumeContext["csi.storage.k8s.io/pod.name"] = "stored-pod-name"
	real.VolumeContext["csi.storage.k8s.io/pod.namespace"] = "stored-pod-space"
	node, downstream, request := readonlyNodeFixture(t, real)
	for _, tc := range []struct{ worker, podName, podNamespace string }{
		{"worker-one", "worker-one-name", "worker-one-namespace"},
		{"worker-two", "", ""},
	} {
		t.Run(tc.worker, func(t *testing.T) {
			request.VolumeContext[PodUIDKey] = tc.worker
			delete(request.VolumeContext, "csi.storage.k8s.io/pod.name")
			delete(request.VolumeContext, "csi.storage.k8s.io/pod.namespace")
			if tc.podName != "" {
				request.VolumeContext["csi.storage.k8s.io/pod.name"] = tc.podName
				request.VolumeContext["csi.storage.k8s.io/pod.namespace"] = tc.podNamespace
			}
			_, err := node.NodePublishVolume(t.Context(), request)
			require.NoError(t, err)
			for key, value := range map[string]string{
				PodUIDKey:                          tc.worker,
				"csi.storage.k8s.io/pod.name":      tc.podName,
				"csi.storage.k8s.io/pod.namespace": tc.podNamespace,
			} {
				if value == "" {
					require.NotContains(t, downstream.published.VolumeContext, key)
				} else {
					require.Equal(t, value, downstream.published.VolumeContext[key])
				}
			}
			for key, value := range actorMetadataFixture() {
				require.Equal(t, value, downstream.published.VolumeContext[key])
			}
			require.Equal(t, testTarget, downstream.published.TargetPath)
		})
	}
	_, err := node.NodeUnpublishVolume(t.Context(), &csi.NodeUnpublishVolumeRequest{VolumeId: testID, TargetPath: testTarget})
	require.NoError(t, err)
	require.Equal(t, testID, downstream.unpublished.VolumeId)
}

func TestPublishRejectsIncompleteOrMismatchedActorMetadata(t *testing.T) {
	for _, key := range []string{"csi.alibabacloud.com/actor.uid", "csi.alibabacloud.com/actor.name", "csi.alibabacloud.com/actor.namespace"} {
		t.Run(key, func(t *testing.T) {
			for _, value := range []string{"", "different"} {
				node, downstream, request := readonlyNodeFixture(t, writablePublishFixture())
				node.opts.Lookup = identifiedLookup(t)
				for name, actual := range actorMetadataFixture() {
					request.VolumeContext[name] = actual
				}
				request.VolumeContext[key] = value
				_, err := node.NodePublishVolume(t.Context(), request)
				require.Error(t, err)
				require.Nil(t, downstream.published)
			}
		})
	}
}

func TestPublishRejectsUseclientAndCNFS(t *testing.T) {
	for _, key := range []string{"useclient", "UseClient", "containernetworkfilesystem", "ContainerNetworkFileSystem"} {
		t.Run(key, func(t *testing.T) {
			real := writablePublishFixture()
			real.VolumeContext[key] = "efc"
			annotation := annotationFixture(t, real)
			node, _, in := readonlyNodeFixture(t, real)
			node.opts.Lookup = func(context.Context, ActorReference) (ActorInfo, error) {
				return ActorInfo{UID: testUID, Name: "actor", Atespace: "storage-test", Annotation: annotation}, nil
			}
			_, err := node.NodePublishVolume(t.Context(), in)
			require.Equal(t, codes.FailedPrecondition, status.Code(err))
			require.ErrorContains(t, err, "not supported through the Substrate bridge")
		})
	}
}

func TestPublishRejectsTargetOutsideActorRoot(t *testing.T) {
	node, _, in := readonlyNodeFixture(t, writablePublishFixture())
	in.TargetPath = "/var/lib/ateom-gvisor/actors/11111111-2222-3333-4444-555555555555/volumes/data"
	_, err := node.NodePublishVolume(t.Context(), in)
	require.Equal(t, codes.InvalidArgument, status.Code(err))
	require.ErrorContains(t, err, "actor UID must match the target actor directory")
}

func TestPublishTreatsVolumeIDAsOpaqueAndUsesTargetLayout(t *testing.T) {
	// Substrate's logical volume naming is not a contract: the VolumeId is only
	// forwarded as an opaque label, while the Actor and the publish request key
	// come from the target's position under the actor root.
	real := writablePublishFixture()
	node, downstream, in := readonlyNodeFixture(t, real)
	node.opts.Lookup = identifiedLookup(t)
	in.VolumeId = "whatever-ateapi-chose"
	in.TargetPath = filepath.Join("/var/lib/ateom-gvisor/actors", testUID, "mnt", "data")
	_, err := node.NodePublishVolume(t.Context(), in)
	require.NoError(t, err)
	require.Equal(t, in.TargetPath, downstream.published.TargetPath)
	require.Equal(t, testUID, downstream.published.VolumeContext["csi.alibabacloud.com/actor.uid"])
}

func TestPublishRequiresWorkerPodUID(t *testing.T) {
	real := writablePublishFixture()
	annotation := annotationFixture(t, real)
	node, _, in := readonlyNodeFixture(t, real)
	node.opts.Lookup = func(context.Context, ActorReference) (ActorInfo, error) {
		return ActorInfo{UID: testUID, Name: "actor", Atespace: "storage-test", Annotation: annotation}, nil
	}
	delete(in.VolumeContext, PodUIDKey)
	_, err := node.NodePublishVolume(t.Context(), in)
	require.Equal(t, codes.InvalidArgument, status.Code(err))
	require.ErrorContains(t, err, "worker Pod UID is required")
}

func TestPublishRejectsWhenNoBackendConfigured(t *testing.T) {
	real := writablePublishFixture()
	annotation := annotationFixture(t, real)
	node, _, in := readonlyNodeFixture(t, real)
	node.opts.Lookup = func(context.Context, ActorReference) (ActorInfo, error) {
		return ActorInfo{UID: testUID, Name: "actor", Atespace: "storage-test", Annotation: annotation}, nil
	}
	node.opts.Backends = nil
	_, err := node.NodePublishVolume(t.Context(), in)
	require.Equal(t, codes.FailedPrecondition, status.Code(err))
	require.ErrorContains(t, err, "has no backend in the Substrate bridge")
}

func TestPublishRejectsUnsupportedFilesystemAtTarget(t *testing.T) {
	real := writablePublishFixture()
	annotation := annotationFixture(t, real)
	node, _, in := readonlyNodeFixture(t, real)
	node.opts.Lookup = func(context.Context, ActorReference) (ActorInfo, error) {
		return ActorInfo{UID: testUID, Name: "actor", Atespace: "storage-test", Annotation: annotation}, nil
	}
	node.opts.Mounter = mount.NewFakeMounter([]mount.MountPoint{{Path: testTarget, Device: "/dev/sda1", Type: "ext4"}})
	_, err := node.NodePublishVolume(t.Context(), in)
	require.Equal(t, codes.FailedPrecondition, status.Code(err))
	require.ErrorContains(t, err, "unsupported filesystem")
}

func TestUnpublishRejectsUnsupportedLiveMount(t *testing.T) {
	node := NewNode(NodeOptions{
		ActorRoot: DefaultActorRoot,
		StateDir:  t.TempDir(),
		Mounter:   mount.NewFakeMounter([]mount.MountPoint{{Path: testTarget, Device: "/dev/sda1", Type: "ext4"}}),
		Backends:  nasBackends(&recordingNAS{}),
	})
	_, err := node.NodeUnpublishVolume(t.Context(), &csi.NodeUnpublishVolumeRequest{VolumeId: testID, TargetPath: testTarget})
	require.Equal(t, codes.FailedPrecondition, status.Code(err))
	require.ErrorContains(t, err, "refusing to unmount an unsupported live mount")
}

func TestUnpublishSucceedsWhenNothingMounted(t *testing.T) {
	node := NewNode(NodeOptions{
		ActorRoot: DefaultActorRoot,
		StateDir:  t.TempDir(),
		Mounter:   mount.NewFakeMounter(nil),
		Backends:  nasBackends(&recordingNAS{}),
	})
	_, err := node.NodeUnpublishVolume(t.Context(), &csi.NodeUnpublishVolumeRequest{VolumeId: testID, TargetPath: testTarget})
	require.NoError(t, err)
}

func TestUnpublishRetriesLiveMountWhenNASFails(t *testing.T) {
	state := t.TempDir()
	nas := &recordingNAS{err: status.Error(codes.Unavailable, "daemon unavailable")}
	node := NewNode(NodeOptions{ActorRoot: DefaultActorRoot, StateDir: state, Mounter: mount.NewFakeMounter([]mount.MountPoint{{Path: testTarget, Device: "server:/", Type: "nfs"}}), Backends: nasBackends(nas), Lookup: func(context.Context, ActorReference) (ActorInfo, error) {
		t.Fatal("unpublish must not query actor")
		return ActorInfo{}, nil
	}})
	_, err := node.NodeUnpublishVolume(t.Context(), &csi.NodeUnpublishVolumeRequest{VolumeId: testID, TargetPath: testTarget})
	require.Equal(t, codes.Unavailable, status.Code(err))
	nas.err = nil
	node = NewNode(node.opts)
	_, err = node.NodeUnpublishVolume(t.Context(), &csi.NodeUnpublishVolumeRequest{VolumeId: testID, TargetPath: testTarget})
	require.NoError(t, err)
}

func TestForeignFilesystemIsNotRemoved(t *testing.T) {
	root := t.TempDir()
	target := filepath.Join(root, testUID, "volumes", "data")
	mounter := mount.NewFakeMounter([]mount.MountPoint{{Device: "unknown", Path: target, Type: "tmpfs"}})
	node := NewNode(NodeOptions{ActorRoot: root, StateDir: t.TempDir(), Mounter: mounter, Backends: nasBackends(&recordingNAS{})})
	_, err := node.NodeUnpublishVolume(t.Context(), &csi.NodeUnpublishVolumeRequest{VolumeId: testID, TargetPath: target})
	require.Equal(t, codes.FailedPrecondition, status.Code(err))
	require.Len(t, mounter.MountPoints, 1)
}

func TestStatelessPublishDoesNotWriteBindingFiles(t *testing.T) {
	node, _, in := readonlyNodeFixture(t, writablePublishFixture())
	_, err := node.NodePublishVolume(t.Context(), in)
	require.NoError(t, err)
	entries, err := os.ReadDir(node.opts.StateDir)
	require.NoError(t, err)
	require.Empty(t, entries)
}

func TestStatelessUnpublishUsesLiveNASAndLogicalLockKey(t *testing.T) {
	nas := &recordingNAS{}
	node := NewNode(NodeOptions{ActorRoot: DefaultActorRoot, Backends: nasBackends(nas), Mounter: mount.NewFakeMounter([]mount.MountPoint{{Path: testTarget, Device: "server:/", Type: "nfs"}}), Lookup: func(context.Context, ActorReference) (ActorInfo, error) {
		t.Fatal("unpublish must not query actor")
		return ActorInfo{}, nil
	}})
	_, err := node.NodeUnpublishVolume(t.Context(), &csi.NodeUnpublishVolumeRequest{VolumeId: testID, TargetPath: testTarget})
	require.NoError(t, err)
	require.Equal(t, testID, nas.unpublished.VolumeId)
	require.Equal(t, testTarget, nas.unpublished.TargetPath)
	stateFile := filepath.Join(t.TempDir(), "not-a-directory")
	require.NoError(t, os.WriteFile(stateFile, []byte("unused for NAS"), 0600))
	node.opts.StateDir = stateFile
	_, err = node.NodeUnpublishVolume(t.Context(), &csi.NodeUnpublishVolumeRequest{VolumeId: testID, TargetPath: testTarget})
	require.NoError(t, err)
}

func TestStatelessUnpublishWithoutMountNeedsNoBackendOrState(t *testing.T) {
	node := NewNode(NodeOptions{ActorRoot: DefaultActorRoot, Mounter: mount.NewFakeMounter(nil)})
	for range 2 {
		_, err := node.NodeUnpublishVolume(t.Context(), &csi.NodeUnpublishVolumeRequest{VolumeId: testID, TargetPath: testTarget})
		require.NoError(t, err)
	}
}

func TestIdentityReadsVolumeIdentityFromTheTarget(t *testing.T) {
	node := NewNode(NodeOptions{ActorRoot: DefaultActorRoot})
	for _, tc := range []struct {
		name    string
		target  string
		want    volumeIdentity
		wantErr codes.Code
	}{
		{"actor volume directory", testTarget, volumeIdentity{ActorUID: testUID, VolumeName: "data"}, codes.OK},
		{"deeper layout", testTarget + "/mnt/inner", volumeIdentity{ActorUID: testUID, VolumeName: "inner"}, codes.OK},
		{"actor directory itself", "/var/lib/ateom-gvisor/actors/" + testUID, volumeIdentity{}, codes.InvalidArgument},
		{"outside the actor root", "/var/lib/kubelet/pods/x/volumes/kubernetes.io~csi/y/mount", volumeIdentity{}, codes.InvalidArgument},
		{"relative target", "var/lib/ateom-gvisor/actors/" + testUID + "/volumes/data", volumeIdentity{}, codes.InvalidArgument},
		{"non-clean target", "/var/lib/ateom-gvisor/actors/" + testUID + "/volumes/../data", volumeIdentity{}, codes.InvalidArgument},
	} {
		t.Run(tc.name, func(t *testing.T) {
			id, err := node.identity(tc.target)
			require.Equal(t, tc.wantErr, status.Code(err))
			require.Equal(t, tc.want, id)
		})
	}
	_, err := NewNode(NodeOptions{ActorRoot: "relative/actor/root"}).identity(testTarget)
	require.Equal(t, codes.FailedPrecondition, status.Code(err))
}

// The stored target is replaced by the host directory and NAS does not look at the staging path
// of a publish at all, so neither is a reason to refuse the mount.
func TestPublishForwardsFieldsTheBridgeNeverReads(t *testing.T) {
	for _, storedTarget := range []struct {
		name, value string
	}{
		{"empty", ""},
		{"relative", "relative"},
		{"non-clean", "/guest/../unused"},
		{"clean absolute", "/guest/never/used"},
	} {
		t.Run(storedTarget.name, func(t *testing.T) {
			real := writablePublishFixture()
			real.TargetPath = storedTarget.value
			real.StagingTargetPath = "stored/staging/is/unused/by/nas"
			node, downstream, in := readonlyNodeFixture(t, real)
			in.StagingTargetPath = "/outer/staging/must-not-replace-inner"
			_, err := node.NodePublishVolume(t.Context(), in)
			require.NoError(t, err)
			require.Equal(t, testTarget, downstream.published.TargetPath)
			require.Equal(t, real.StagingTargetPath, downstream.published.StagingTargetPath)
		})
	}
}

// The bridge stores no annotation of its own, so the payload a caller chose to write is not
// something the bridge puts a second, smaller ceiling on.
func TestPublishImposesNoAnnotationSizeLimit(t *testing.T) {
	real := writablePublishFixture()
	real.VolumeContext["opaque-configuration"] = strings.Repeat("x", 300*1024)
	node, downstream, in := readonlyNodeFixture(t, real)
	_, err := node.NodePublishVolume(t.Context(), in)
	require.NoError(t, err)
	require.Equal(t, real.VolumeContext["opaque-configuration"], downstream.published.VolumeContext["opaque-configuration"])
}

func TestCheckCallerMetadataMatchesTheTargetActor(t *testing.T) {
	id := volumeIdentity{ActorUID: testUID, VolumeName: "data"}
	for _, tc := range []struct {
		name    string
		mutate  func(map[string]string)
		wantErr codes.Code
	}{
		{"actor and distinct worker", func(map[string]string) {}, codes.OK},
		{"legacy pod UID only", func(context map[string]string) {
			delete(context, "csi.alibabacloud.com/actor.uid")
			delete(context, "csi.alibabacloud.com/actor.name")
			delete(context, "csi.alibabacloud.com/actor.namespace")
		}, codes.InvalidArgument},
		{"foreign actor", func(context map[string]string) {
			context["csi.alibabacloud.com/actor.uid"] = "00000000-0000-0000-0000-000000000000"
		}, codes.InvalidArgument},
		{"different worker", func(context map[string]string) {
			context[PodUIDKey] = "00000000-0000-0000-0000-000000000000"
		}, codes.OK},
		{"missing worker", func(context map[string]string) { delete(context, PodUIDKey) }, codes.InvalidArgument},
	} {
		t.Run(tc.name, func(t *testing.T) {
			req := &csi.NodePublishVolumeRequest{VolumeContext: nodeContextFixture()}
			tc.mutate(req.VolumeContext)
			require.Equal(t, tc.wantErr, status.Code(checkCallerMetadata(req, id)))
		})
	}
}

func TestBackendRequestChangesOnlyTargetModeAndIdentity(t *testing.T) {
	actor := ActorInfo{UID: testUID, Name: "actor", Atespace: "storage-test"}
	in := &csi.NodePublishVolumeRequest{
		VolumeId:         "whatever-the-caller-called-it",
		TargetPath:       testTarget,
		Readonly:         true,
		VolumeCapability: publishFixture().VolumeCapability,
		VolumeContext:    map[string]string{"server": "must-not-replace-the-bound-server", PodUIDKey: "worker-pod-uid", "csi.storage.k8s.io/pod.namespace": "worker-space"},
	}
	for _, tc := range []struct {
		name        string
		storedPodID string
	}{
		{"stored request without pod identity", ""},
		{"stored request that already names the actor", testUID},
	} {
		t.Run(tc.name, func(t *testing.T) {
			real := writablePublishFixture()
			real.VolumeContext["csi.storage.k8s.io/pod.name"] = "stored-pod-name"
			if tc.storedPodID != "" {
				real.VolumeContext[PodUIDKey] = tc.storedPodID
			}
			want := proto.Clone(real).(*csi.NodePublishVolumeRequest)
			got := backendRequest(in, resolvedMount{Actor: actor, Request: real})
			require.Equal(t, "customer-pv-a1b2c3", got.VolumeId)
			require.False(t, got.Readonly)
			require.Equal(t, csi.VolumeCapability_AccessMode_SINGLE_NODE_WRITER, got.VolumeCapability.GetAccessMode().GetMode())
			require.Equal(t, "ap-example.fs.cn-beijing.agenticfs.aliyuncs.com", got.VolumeContext["server"])
			require.NotContains(t, got.VolumeContext, "csi.storage.k8s.io/pod.name")
			require.Equal(t, "worker-pod-uid", got.VolumeContext[PodUIDKey])
			want.TargetPath = testTarget
			want.VolumeContext["csi.alibabacloud.com/substrate-mode"] = "true"
			want.VolumeContext["csi.alibabacloud.com/actor.uid"] = testUID
			want.VolumeContext["csi.alibabacloud.com/actor.name"] = "actor"
			want.VolumeContext["csi.alibabacloud.com/actor.namespace"] = "storage-test"
			want.VolumeContext[PodUIDKey] = "worker-pod-uid"
			want.VolumeContext["csi.storage.k8s.io/pod.namespace"] = "worker-space"
			delete(want.VolumeContext, "csi.storage.k8s.io/pod.name")
			require.True(t, proto.Equal(want, got), "the forwarded request is not the stored request")
		})
	}
}

func TestCheckBackendTargetRejectsWhatItCannotMountOver(t *testing.T) {
	for _, tc := range []struct {
		name    string
		backend Backend
		node    *Node
		wantErr string
	}{
		{"clear", NASBackend(&recordingNAS{}), NewNode(NodeOptions{Mounter: mount.NewFakeMounter(nil)}), ""},
		{"no backend node server", Backend{Driver: NASDriverName, OwnsMount: nasMount}, NewNode(NodeOptions{Mounter: mount.NewFakeMounter(nil)}), "backend forwarding is not configured"},
		{"foreign filesystem", NASBackend(&recordingNAS{}), NewNode(NodeOptions{Mounter: mount.NewFakeMounter([]mount.MountPoint{{Path: testTarget, Type: "ext4"}})}), "unsupported filesystem"},
		{"NAS mount of any access mode", NASBackend(&recordingNAS{}), NewNode(NodeOptions{Mounter: mount.NewFakeMounter([]mount.MountPoint{{Path: testTarget, Type: "nfs", Opts: []string{"rw"}}})}), ""},
		{"stacked mount hides the NAS mount below", NASBackend(&recordingNAS{}), NewNode(NodeOptions{Mounter: mount.NewFakeMounter([]mount.MountPoint{{Path: testTarget, Type: "nfs"}, {Path: testTarget, Type: "tmpfs"}})}), "unsupported filesystem"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.node.checkBackendTarget(tc.backend, testTarget)
			if tc.wantErr == "" {
				require.NoError(t, err)
				return
			}
			require.ErrorContains(t, err, tc.wantErr)
		})
	}
}
