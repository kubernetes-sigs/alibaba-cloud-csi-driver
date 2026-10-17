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
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	mount "k8s.io/mount-utils"
)

func TestPublishPreservesInnerReadonlyConfiguration(t *testing.T) {
	for _, tc := range []struct {
		name   string
		mutate func(*csi.NodePublishVolumeRequest, *csi.NodePublishVolumeRequest)
	}{
		{"writable", func(_, _ *csi.NodePublishVolumeRequest) {}},
		{"annotation boolean", func(_, real *csi.NodePublishVolumeRequest) { real.Readonly = true }},
		{"annotation single reader", func(_, real *csi.NodePublishVolumeRequest) {
			real.VolumeCapability.AccessMode.Mode = csi.VolumeCapability_AccessMode_SINGLE_NODE_READER_ONLY
		}},
		{"annotation multi reader", func(_, real *csi.NodePublishVolumeRequest) {
			real.VolumeCapability.AccessMode.Mode = csi.VolumeCapability_AccessMode_MULTI_NODE_READER_ONLY
		}},
		{"annotation ro flags", func(_, real *csi.NodePublishVolumeRequest) {
			real.VolumeCapability.GetMount().MountFlags = []string{"tls", "vers=3", "rw,ro"}
		}},
		{"annotation options", func(_, real *csi.NodePublishVolumeRequest) {
			real.VolumeContext["options"] = "tls,ram,rw,ro"
		}},
		{"mount flags override readonly options", func(_, real *csi.NodePublishVolumeRequest) {
			real.VolumeContext["options"] = "tls,ram,ro"
			real.VolumeCapability.GetMount().MountFlags = []string{"tls", "ram", "rw"}
		}},
		{"explicit readonly does not normalize flags", func(_, real *csi.NodePublishVolumeRequest) {
			real.Readonly = true
			real.VolumeCapability.GetMount().MountFlags = []string{"tls,rw,vers=3", `context="value,rw,value"`, "rw"}
			real.VolumeContext["options"] = "tls,ram,rw,nolock"
		}},
		{"quoted option is not ro", func(in, _ *csi.NodePublishVolumeRequest) {
			in.VolumeCapability.GetMount().MountFlags = []string{`context="value,ro,value"`}
		}},
		{"unrecognized context is not merged", func(in, _ *csi.NodePublishVolumeRequest) {
			in.VolumeContext["options"] = "ro"
			in.VolumeContext["ro"] = "true"
		}},
		{"obsolete controller digest is ignored", func(in, _ *csi.NodePublishVolumeRequest) {
			in.VolumeContext["csi.alibabacloud.com/substrate-binding-digest"] = "obsolete"
		}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			real := writablePublishFixture()
			in := &csi.NodePublishVolumeRequest{VolumeCapability: writablePublishFixture().VolumeCapability, VolumeContext: map[string]string{}}
			tc.mutate(in, real)
			node, downstream, request := readonlyNodeFixture(t, real)
			request.Readonly = in.Readonly
			request.VolumeCapability = in.VolumeCapability
			for key, value := range in.VolumeContext {
				request.VolumeContext[key] = value
			}
			request.VolumeContext["server"] = "must-not-replace-the-bound-server"
			request.VolumeContext["path"] = "/must-not-replace-the-bound-path"
			want := proto.Clone(real).(*csi.NodePublishVolumeRequest)
			want.TargetPath = testTarget
			want.VolumeContext["csi.alibabacloud.com/substrate-mode"] = "true"
			for key, value := range nodeContextFixture() {
				want.VolumeContext[key] = value
			}
			_, err := node.NodePublishVolume(t.Context(), request)
			require.NoError(t, err)
			require.True(t, proto.Equal(want, downstream.published), "forwarded request differs from the annotation")
		})
	}
}

// TestOuterReadonlyNeverReachesTheForwardedRequest pins the passthrough rule: the read-only
// state of a Substrate volume comes from the Actor publish request, and the outer request is
// neither honoured nor rejected.
func TestOuterReadonlyNeverReachesTheForwardedRequest(t *testing.T) {
	for _, tc := range []struct {
		name   string
		mutate func(*csi.NodePublishVolumeRequest)
	}{
		{"boolean", func(in *csi.NodePublishVolumeRequest) { in.Readonly = true }},
		{"single reader", func(in *csi.NodePublishVolumeRequest) {
			in.VolumeCapability.AccessMode.Mode = csi.VolumeCapability_AccessMode_SINGLE_NODE_READER_ONLY
		}},
		{"multi reader", func(in *csi.NodePublishVolumeRequest) {
			in.VolumeCapability.AccessMode.Mode = csi.VolumeCapability_AccessMode_MULTI_NODE_READER_ONLY
		}},
		{"mount flag", func(in *csi.NodePublishVolumeRequest) { in.VolumeCapability.GetMount().MountFlags = []string{"ro"} }},
		{"combined mount flags", func(in *csi.NodePublishVolumeRequest) {
			in.VolumeCapability.GetMount().MountFlags = []string{"tls, ro ,rw"}
		}},
		{"volume context options", func(in *csi.NodePublishVolumeRequest) { in.VolumeContext["options"] = "tls,ram,ro" }},
	} {
		t.Run(tc.name, func(t *testing.T) {
			real := writablePublishFixture()
			node, downstream, in := readonlyNodeFixture(t, real)
			lookup := node.opts.Lookup
			calls := 0
			node.opts.Lookup = func(ctx context.Context, ref ActorReference) (ActorInfo, error) {
				calls++
				return lookup(ctx, ref)
			}
			tc.mutate(in)
			_, err := node.NodePublishVolume(t.Context(), in)
			require.NoError(t, err)
			require.Equal(t, 1, calls)
			want := proto.Clone(real).(*csi.NodePublishVolumeRequest)
			want.TargetPath = testTarget
			want.VolumeContext["csi.alibabacloud.com/substrate-mode"] = "true"
			for key, value := range nodeContextFixture() {
				want.VolumeContext[key] = value
			}
			require.True(t, proto.Equal(want, downstream.published), "the outer request changed the forwarded request")
			require.False(t, downstream.published.Readonly)
		})
	}
}

func TestReadonlyPublishNeedsNoStateDirectory(t *testing.T) {
	node, downstream, in := readonlyNodeFixture(t, publishFixture())
	node.opts.StateDir = ""
	for range 2 {
		node = NewNode(node.opts)
		_, err := node.NodePublishVolume(t.Context(), in)
		require.NoError(t, err)
		require.True(t, downstream.published.Readonly)
	}
}

func TestGoldenPlaceholderIgnoresOuterReadonlyAndStaysWritable(t *testing.T) {
	root, err := filepath.EvalSymlinks(t.TempDir())
	require.NoError(t, err)
	target := filepath.Join(root, testUID, "volumes", "data")
	mounter := mount.NewFakeMounter(nil)
	lookups := 0
	lookup := func(context.Context, ActorReference) (ActorInfo, error) {
		lookups++
		return ActorInfo{UID: testUID, Atespace: "ate-golden", Name: "template-uid", TemplateUID: "template-uid", Golden: true}, nil
	}
	node := NewNode(NodeOptions{ActorRoot: root, StateDir: t.TempDir(), Lookup: lookup, Mounter: mounter})
	attributes := goldenMetadataFixture()
	attributes[PodUIDKey] = testUID
	in := &csi.NodePublishVolumeRequest{VolumeId: testID, TargetPath: target, Readonly: true, VolumeContext: attributes}
	_, err = node.NodePublishVolume(t.Context(), in)
	require.NoError(t, err)
	require.Equal(t, 1, lookups)
	require.Len(t, mounter.MountPoints, 1)
	require.Equal(t, []string{"bind"}, mounter.MountPoints[0].Opts)
	in.Readonly = false
	node = NewNode(node.opts)
	_, err = node.NodePublishVolume(t.Context(), in)
	require.NoError(t, err)
	require.Equal(t, 2, lookups)
	require.Len(t, mounter.MountPoints, 1)
	_, err = node.NodeUnpublishVolume(t.Context(), &csi.NodeUnpublishVolumeRequest{VolumeId: testID, TargetPath: target})
	require.NoError(t, err)
	require.Empty(t, mounter.MountPoints)
}

func TestObsoleteMetadataCannotBlockStatelessPublish(t *testing.T) {
	real := writablePublishFixture()
	node, downstream, in := readonlyNodeFixture(t, real)
	path := filepath.Join(node.opts.StateDir, "legacy.json")
	require.NoError(t, os.WriteFile(path, []byte("corrupt or obsolete metadata"), 0600))
	_, err := node.NodePublishVolume(t.Context(), in)
	require.NoError(t, err)
	require.False(t, downstream.published.Readonly)
	data, err := os.ReadFile(path)
	require.NoError(t, err)
	require.Equal(t, "corrupt or obsolete metadata", string(data))
}

func TestGoldenPlaceholderLifecycleWithoutNAS(t *testing.T) {
	for _, tc := range []struct{ name, retryID, unpublishID string }{
		{"stable ID", testID, testID},
		{"changed publish ID", "publish-alias", "publish-alias"},
		{"changed unpublish ID", testID, "cleanup-alias"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			root, err := filepath.EvalSymlinks(t.TempDir())
			require.NoError(t, err)
			state := t.TempDir()
			target := filepath.Join(root, testUID, "volumes", "data")
			mounter := mount.NewFakeMounter(nil)
			nas := &recordingNAS{err: errors.New("golden must not call NAS")}
			lookup := func(context.Context, ActorReference) (ActorInfo, error) {
				return ActorInfo{UID: testUID, Atespace: "ate-golden", Name: "template-uid", TemplateUID: "template-uid", Golden: true}, nil
			}
			attributes := goldenMetadataFixture()
			node := NewNode(NodeOptions{ActorRoot: root, StateDir: state, Mounter: mounter, Backends: nasBackends(nas), Lookup: lookup})
			attributes[PodUIDKey] = "worker-pod-uid"
			in := &csi.NodePublishVolumeRequest{VolumeId: testID, TargetPath: target, VolumeContext: attributes}
			_, err = node.NodePublishVolume(t.Context(), in)
			require.NoError(t, err)
			require.Nil(t, nas.published)
			require.DirExists(t, target)
			require.Len(t, mounter.MountPoints, 1)
			in.VolumeId = tc.retryID
			_, err = node.NodePublishVolume(t.Context(), in)
			require.NoError(t, err)
			require.Len(t, mounter.MountPoints, 1)
			placeholder, err := node.placeholderSource(target)
			require.NoError(t, err)
			require.NoError(t, os.WriteFile(filepath.Join(placeholder, "build-output"), []byte("temporary"), 0600))
			restarted := NewNode(NodeOptions{ActorRoot: root, StateDir: state, Mounter: mounter, Backends: nasBackends(nas)})
			_, err = restarted.NodeUnpublishVolume(t.Context(), &csi.NodeUnpublishVolumeRequest{VolumeId: tc.unpublishID, TargetPath: target})
			require.NoError(t, err)
			require.Nil(t, nas.unpublished)
			require.Empty(t, mounter.MountPoints)
			// What the warmup wrote into the placeholder is not the bridge's to delete.
			retained, readErr := os.ReadFile(filepath.Join(placeholder, "build-output"))
			require.NoError(t, readErr)
			require.Equal(t, "temporary", string(retained))
			_, err = restarted.NodeUnpublishVolume(t.Context(), &csi.NodeUnpublishVolumeRequest{VolumeId: tc.unpublishID, TargetPath: target})
			require.NoError(t, err)
			require.DirExists(t, placeholder)
		})
	}
}

// A placeholder source that holds nothing is reclaimed; one the Golden warmup wrote
// into is left alone, because deleting that recursively would discard the written
// data without any ownership check.
func TestUnpublishPlaceholderRemovesEmptySourceOnly(t *testing.T) {
	for _, tc := range []struct {
		name      string
		warmupOut bool
	}{
		{"empty source is removed", false},
		{"source with warmup output is kept", true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			root := t.TempDir()
			target := filepath.Join(root, testUID, "volumes", "data")
			mounter := mount.NewFakeMounter(nil)
			node := NewNode(NodeOptions{ActorRoot: root, StateDir: t.TempDir(), Mounter: mounter, Lookup: func(context.Context, ActorReference) (ActorInfo, error) {
				return ActorInfo{UID: testUID, Name: "template-uid", Atespace: "ate-golden", TemplateUID: "template-uid", Golden: true}, nil
			}})
			source, err := node.placeholderSource(target)
			require.NoError(t, err)
			require.NoError(t, os.MkdirAll(source, 0755))
			require.NoError(t, os.MkdirAll(target, 0755))
			if tc.warmupOut {
				require.NoError(t, os.WriteFile(filepath.Join(source, "build-output"), []byte("temporary"), 0600))
			}
			mounter.MountPoints = []mount.MountPoint{{Device: source, Path: target, Type: "none"}}

			_, err = node.NodeUnpublishVolume(t.Context(), &csi.NodeUnpublishVolumeRequest{VolumeId: testID, TargetPath: target})
			require.NoError(t, err)
			require.Empty(t, mounter.MountPoints)
			if tc.warmupOut {
				retained, readErr := os.ReadFile(filepath.Join(source, "build-output"))
				require.NoError(t, readErr)
				require.Equal(t, "temporary", string(retained))
			} else {
				require.NoDirExists(t, source)
			}
		})
	}
}

func TestGoldenPlaceholderInspectionErrors(t *testing.T) {
	for _, tc := range []struct{ name, action string }{
		{"source", "inspect placeholder source"},
		{"target", "inspect Golden target"},
	} {
		for _, operation := range []string{"publish", "unpublish"} {
			t.Run(tc.name+"/"+operation, func(t *testing.T) {
				root := t.TempDir()
				target := filepath.Join(root, testUID, "volumes", "data")
				mounter := mount.NewFakeMounter([]mount.MountPoint{{Device: "/dev/backing", Path: target, Type: "ext4"}})
				node := NewNode(NodeOptions{ActorRoot: root, StateDir: t.TempDir(), Mounter: mounter, Lookup: func(context.Context, ActorReference) (ActorInfo, error) {
					return ActorInfo{UID: testUID, Name: "template-uid", Atespace: "ate-golden", TemplateUID: "template-uid", Golden: true}, nil
				}})
				source, err := node.placeholderSource(target)
				require.NoError(t, err)
				if tc.name == "source" {
					require.NoError(t, os.WriteFile(filepath.Dir(source), []byte("not a directory"), 0600))
				} else {
					require.NoError(t, os.MkdirAll(source, 0755))
				}
				if operation == "publish" {
					attributes := goldenMetadataFixture()
					attributes[PodUIDKey] = "worker-pod-uid"
					_, err = node.NodePublishVolume(t.Context(), &csi.NodePublishVolumeRequest{VolumeId: testID, TargetPath: target, VolumeContext: attributes})
				} else {
					_, err = node.NodeUnpublishVolume(t.Context(), &csi.NodeUnpublishVolumeRequest{VolumeId: testID, TargetPath: target})
				}
				require.Equal(t, codes.Internal, status.Code(err))
				require.ErrorContains(t, err, tc.action)
				require.Len(t, mounter.MountPoints, 1)
			})
		}
	}
}

func TestPlaceholderPathIsStableAndTargetScoped(t *testing.T) {
	options := NodeOptions{StateDir: t.TempDir()}
	one, err := NewNode(options).placeholderSource(testTarget)
	require.NoError(t, err)
	again, err := NewNode(options).placeholderSource(testTarget)
	require.NoError(t, err)
	require.Equal(t, one, again)
	two, err := NewNode(options).placeholderSource("/another/target")
	require.NoError(t, err)
	require.NotEqual(t, one, two)
}

func TestMountedPlaceholderDoesNotRecreateMissingSource(t *testing.T) {
	node := NewNode(NodeOptions{StateDir: t.TempDir(), Mounter: mount.NewFakeMounter(nil)})
	source, err := node.placeholderSource(testTarget)
	require.NoError(t, err)
	node.opts.Mounter = mount.NewFakeMounter([]mount.MountPoint{{Device: source, Path: testTarget, Type: "none"}})
	require.Equal(t, codes.FailedPrecondition, status.Code(node.publishPlaceholder(testTarget)))
	require.NoDirExists(t, source)
}

func TestPlaceholderRejectsSymlinkSource(t *testing.T) {
	node := NewNode(NodeOptions{StateDir: t.TempDir(), Mounter: mount.NewFakeMounter(nil)})
	source, err := node.placeholderSource(testTarget)
	require.NoError(t, err)
	require.NoError(t, os.MkdirAll(filepath.Dir(source), 0700))
	require.NoError(t, os.Symlink(t.TempDir(), source))
	require.Equal(t, codes.FailedPrecondition, status.Code(node.publishPlaceholder(testTarget)))
}

// Only the topmost mount at a target is visible to a later mount and unmount, so a placeholder
// whose source is hidden underneath a foreign mount must neither claim it nor delete from it.
func TestGoldenDoesNotClaimItsSourceHiddenByAnotherMount(t *testing.T) {
	for _, operation := range []string{"publish", "unpublish"} {
		t.Run(operation, func(t *testing.T) {
			root := t.TempDir()
			target := filepath.Join(root, testUID, "volumes", "data")
			mounter := mount.NewFakeMounter(nil)
			node := NewNode(NodeOptions{ActorRoot: root, StateDir: t.TempDir(), Mounter: mounter, Lookup: func(context.Context, ActorReference) (ActorInfo, error) {
				return ActorInfo{UID: testUID, Name: "template-uid", Atespace: "ate-golden", TemplateUID: "template-uid", Golden: true}, nil
			}})
			source, err := node.placeholderSource(target)
			require.NoError(t, err)
			require.NoError(t, os.MkdirAll(source, 0755))
			require.NoError(t, os.MkdirAll(target, 0755))
			require.NoError(t, os.WriteFile(filepath.Join(source, "retain"), []byte("source data"), 0600))
			mounter.MountPoints = []mount.MountPoint{
				{Device: source, Path: target, Type: "none"},
				{Device: "foreign", Path: target, Type: "tmpfs"},
			}
			before := append([]mount.MountPoint(nil), mounter.MountPoints...)
			if operation == "publish" {
				attributes := goldenMetadataFixture()
				attributes[PodUIDKey] = "worker-pod-uid"
				_, err = node.NodePublishVolume(t.Context(), &csi.NodePublishVolumeRequest{VolumeId: testID, TargetPath: target, VolumeContext: attributes})
			} else {
				_, err = node.NodeUnpublishVolume(t.Context(), &csi.NodeUnpublishVolumeRequest{VolumeId: testID, TargetPath: target})
			}
			require.Equal(t, codes.FailedPrecondition, status.Code(err))
			require.Equal(t, before, mounter.MountPoints)
			data, err := os.ReadFile(filepath.Join(source, "retain"))
			require.NoError(t, err)
			require.Equal(t, "source data", string(data))
		})
	}
}
