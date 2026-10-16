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
	"path/filepath"
	"strings"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils/agentidentity"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	mount "k8s.io/mount-utils"
)

const (
	DriverName                = "substrate.csi.alibabacloud.com"
	DriverShortName           = "substrate"
	NASDriverName             = "nasplugin.csi.alibabacloud.com"
	PublishRequestsAnnotation = "ate.dev/csi-volume-publish-requests"
	PodUIDKey                 = agentidentity.PodUIDKey
	DefaultActorRoot          = "/var/lib/ate/actors"
)

type publishEntry struct {
	VolumeName string          `json:"volumeName"`
	Driver     string          `json:"driver"`
	Request    json.RawMessage `json:"request"`
}

// ActorLookup returns the resolved Actor's metadata. Golden requires a verified
// ate-golden Actor whose name equals TemplateUID. Node rechecks these invariants
// at the boundary so alternative lookup implementations follow the same contract.
type ActorLookup func(context.Context, ActorReference) (ActorInfo, error)

type NodeOptions struct {
	NodeID    string
	ActorRoot string
	Lookup    ActorLookup
	// Backends maps a CSI driver name to the backend that forwards mounts for
	// it. The bridge routes a stored publish request to the backend named by
	// that request's driver, so supporting another driver is a registration, not
	// a code change here.
	Backends map[string]Backend
	StateDir string
	Mounter  mount.Interface
}

// Node serializes operations per mount target, not per VolumeId: the VolumeId is
// opaque to the bridge, and two logical volumes reaching one target must still
// serialize. The lock is process-local, like the rest of the bridge's state.
type Node struct {
	csi.UnimplementedNodeServer
	opts  NodeOptions
	locks *utils.VolumeLocks
}

func NewNode(opts NodeOptions) *Node { return &Node{opts: opts, locks: utils.NewVolumeLocks()} }

func (n *Node) NodePublishVolume(ctx context.Context, req *csi.NodePublishVolumeRequest) (*csi.NodePublishVolumeResponse, error) {
	// 1. The mount target, not the VolumeId, says which Actor volume this is.
	id, err := n.identity(req.GetTargetPath())
	if err != nil {
		return nil, err
	}
	// 2. What the caller claims about itself has to line up with that directory.
	if err := checkCallerMetadata(req, id); err != nil {
		return nil, err
	}
	// 3. One operation per target at a time; the lock is process-local.
	if !n.locks.TryAcquire(req.TargetPath) {
		return nil, status.Error(codes.Aborted, "target operation already in progress")
	}
	defer n.locks.Release(req.TargetPath)
	// 4. The Actor API decides what this volume really is, and which request describes it.
	resolved, err := resolveMount(ctx, n.opts.Lookup, id, req.VolumeContext)
	if err != nil {
		return nil, err
	}
	// 5. Actor metadata the caller brought along must describe the same Actor the API named.
	if err := validateActorMetadata(req.VolumeContext, resolved.Actor); err != nil {
		return nil, err
	}
	// 6. A Golden Actor is building its image, so it gets a local placeholder rather than
	// backend storage that would be shared with the running Actor. By design the placeholder is
	// writable whatever the request asked for: it receives writes during template warmup.
	if resolved.Actor.Golden {
		if err := n.publishPlaceholder(req.TargetPath); err != nil {
			return nil, err
		}
		return &csi.NodePublishVolumeResponse{}, nil
	}
	// 7. Route to the backend the stored request's driver names, and let that
	// backend refuse a request it must not be handed. Common invariants were
	// already checked while resolving; this is the driver-specific gate.
	backend, ok := n.backend(resolved.Driver)
	if !ok {
		return nil, status.Errorf(codes.FailedPrecondition, "publish request for driver %q has no backend in the Substrate bridge", resolved.Driver)
	}
	if backend.Validate != nil {
		if err := backend.Validate(resolved.Request, resolved.Actor); err != nil {
			return nil, err
		}
	}
	// 8. The target must be able to host the mount the Actor asks for.
	if err := n.checkBackendTarget(backend, req.TargetPath); err != nil {
		return nil, err
	}
	// 9. The backend does the mount, with the Actor's own request.
	return backend.Node.NodePublishVolume(ctx, backendRequest(req, resolved))
}

// checkCallerMetadata keeps storage ownership separate from placement: Actor metadata
// names the target owner, while atelet supplies the current worker Pod UID. The worker
// can change on resume and must not be compared with the Actor UID.
func checkCallerMetadata(req *csi.NodePublishVolumeRequest, id volumeIdentity) error {
	if req.GetVolumeContext()[agentidentity.ActorUIDKey] != id.ActorUID {
		return status.Error(codes.InvalidArgument, "actor UID must match the target actor directory")
	}
	if req.GetVolumeContext()[PodUIDKey] == "" {
		return status.Error(codes.InvalidArgument, "worker Pod UID is required")
	}
	return nil
}

// checkBackendTarget rejects a target that cannot carry this Actor's mount: no backend node
// server to forward to, or a live mount on a filesystem the routed backend does not own. By design
// nothing else is compared. A live mount the backend recognises is left to it, which answers an
// already-mounted target with success, so a retry stays retryable and matching a live mount
// against a new request stays the caller's unpublish-first obligation instead of a refusal the
// bridge invents.
func (n *Node) checkBackendTarget(backend Backend, target string) error {
	if backend.Node == nil {
		return status.Error(codes.FailedPrecondition, "backend forwarding is not configured")
	}
	live, err := n.mountAt(target)
	if err != nil {
		return err
	}
	if live != nil && !backend.OwnsMount(live) {
		return status.Error(codes.FailedPrecondition, "target is mounted with an unsupported filesystem")
	}
	return nil
}

// backendRequest turns the Actor's stored publish request into the one handed to the routed
// backend: the target becomes the host directory the caller gave us, and the substrate-mode marker
// tells the backend to skip kubelet semantics. Everything else — volume ID, capability, read-only
// state, mount options, other backend volume context entries — pass through untouched, which keeps
// a Substrate mount equivalent to the ACS Sandbox mount of the same volume.
//
// Identity metadata is supplied separately: validated Actor keys own storage and credential
// exchange. PodInfo describes the current worker assignment, never a placement cached in the
// stored request.
//
// By design the caller's own read-only request is not carried over, which is a deliberate
// deviation from the CSI wording that the SP MUST honour readonly on publish: the Actor's request
// decides access, not a Pod object that does not exist here, and a bridge that "corrected" the
// request would make the Substrate and ACS Sandbox forms differ where neither can see it.
func backendRequest(req *csi.NodePublishVolumeRequest, resolved resolvedMount) *csi.NodePublishVolumeRequest {
	real := resolved.Request
	real.TargetPath = req.TargetPath
	real.VolumeContext[utils.SubstrateModeKey] = "true"
	setActorMetadata(real.VolumeContext, resolved.Actor)
	for _, key := range []string{agentidentity.PodUIDKey, agentidentity.PodNameKey, agentidentity.PodNamespaceKey} {
		if value := req.VolumeContext[key]; value != "" {
			real.VolumeContext[key] = value
		} else {
			delete(real.VolumeContext, key)
		}
	}
	return real
}

func (n *Node) NodeUnpublishVolume(ctx context.Context, req *csi.NodeUnpublishVolumeRequest) (*csi.NodeUnpublishVolumeResponse, error) {
	// 1. The target still has to be one of the bridge's actor volume directories. No Actor
	// is looked up: unpublish decides from the mount table alone, which is what keeps it
	// working after the Actor is gone.
	if _, err := n.identity(req.GetTargetPath()); err != nil {
		return nil, err
	}
	// 2. One operation per target at a time.
	if !n.locks.TryAcquire(req.TargetPath) {
		return nil, status.Error(codes.Aborted, "target operation already in progress")
	}
	defer n.locks.Release(req.TargetPath)
	// 3. An absent mount is already unpublished.
	live, err := n.mountAt(req.TargetPath)
	if err != nil {
		return nil, err
	}
	if live == nil {
		return &csi.NodeUnpublishVolumeResponse{}, nil
	}
	// 4. A live mount no backend owns can only be the bridge's own Golden placeholder.
	backend, ok := n.backendForMount(live)
	if !ok {
		return n.unpublishBackendPlaceholder(req)
	}
	// 5. The backend owns its mounts. The volume ID here is the caller's, not the one publish
	// forwarded: the backend uses it as a lock and log key only and the bridge keeps no per-ID
	// state, which is by design so unpublish still works once the Actor and its annotation are gone.
	if _, err := backend.Node.NodeUnpublishVolume(ctx, &csi.NodeUnpublishVolumeRequest{VolumeId: req.VolumeId, TargetPath: req.TargetPath}); err != nil {
		return nil, err
	}
	return &csi.NodeUnpublishVolumeResponse{}, nil
}

// unpublishBackendPlaceholder removes a Golden placeholder and the directory behind it.
// Anything else mounted there is refused: the bridge never unmounts storage it did not
// create.
func (n *Node) unpublishBackendPlaceholder(req *csi.NodeUnpublishVolumeRequest) (*csi.NodeUnpublishVolumeResponse, error) {
	const unsupported = "refusing to unmount an unsupported live mount"
	if !filepath.IsAbs(n.opts.StateDir) || filepath.Clean(n.opts.StateDir) == "/" {
		return nil, status.Error(codes.FailedPrecondition, unsupported)
	}
	source, err := n.placeholderSource(req.TargetPath)
	if err != nil {
		return nil, err
	}
	owned, err := n.placeholderMounted(source, req.TargetPath)
	if err != nil {
		return nil, err
	}
	if !owned {
		return nil, status.Error(codes.FailedPrecondition, unsupported)
	}
	if err := n.unpublishPlaceholder(source, req.TargetPath); err != nil {
		return nil, err
	}
	return &csi.NodeUnpublishVolumeResponse{}, nil
}

// volumeIdentity is what a mount target says about the volume it holds.
type volumeIdentity struct {
	ActorUID   string
	VolumeName string
}

// identity reads the Actor UID and the publish-request key out of the mount
// target instead of the VolumeId: Substrate has not stabilized its logical
// volume naming, so only its host directory layout is treated as a contract.
func (n *Node) identity(target string) (volumeIdentity, error) {
	root := filepath.Clean(n.opts.ActorRoot)
	if !filepath.IsAbs(n.opts.ActorRoot) || root == "/" {
		return volumeIdentity{}, status.Error(codes.FailedPrecondition, "an absolute actor root is required")
	}
	if !filepath.IsAbs(target) || filepath.Clean(target) != target {
		return volumeIdentity{}, status.Error(codes.InvalidArgument, "target must be a clean absolute path")
	}
	rel, ok := strings.CutPrefix(target, root+"/")
	if !ok {
		return volumeIdentity{}, status.Error(codes.InvalidArgument, "target must be inside the actor root")
	}
	segments := strings.Split(rel, "/")
	if len(segments) < 2 {
		return volumeIdentity{}, status.Error(codes.InvalidArgument, "target must be an actor volume directory")
	}
	return volumeIdentity{ActorUID: segments[0], VolumeName: segments[len(segments)-1]}, nil
}

// NodeStage and NodeUnstage answer successfully and do nothing: the bridge mounts at publish and
// has no stage of its own. By design the answer is a success rather than Unimplemented, because
// atelet only skips staging when the driver reports Unimplemented (internal/volume/csi/plugin.go)
// and would otherwise fail the mount. Atelet consequently supplies an outer staging path;
// that outer field is not merged into the stored request. A stored inner staging path is
// independently permitted because the backends do not consume it during NodePublishVolume.
func (*Node) NodeStageVolume(context.Context, *csi.NodeStageVolumeRequest) (*csi.NodeStageVolumeResponse, error) {
	return &csi.NodeStageVolumeResponse{}, nil
}

func (*Node) NodeUnstageVolume(context.Context, *csi.NodeUnstageVolumeRequest) (*csi.NodeUnstageVolumeResponse, error) {
	return &csi.NodeUnstageVolumeResponse{}, nil
}

func (n *Node) NodeGetInfo(context.Context, *csi.NodeGetInfoRequest) (*csi.NodeGetInfoResponse, error) {
	if n.opts.NodeID == "" {
		return nil, status.Error(codes.FailedPrecondition, "bridge node identity is not configured")
	}
	return &csi.NodeGetInfoResponse{NodeId: n.opts.NodeID}, nil
}

func (*Node) NodeGetCapabilities(context.Context, *csi.NodeGetCapabilitiesRequest) (*csi.NodeGetCapabilitiesResponse, error) {
	return &csi.NodeGetCapabilitiesResponse{Capabilities: []*csi.NodeServiceCapability{{Type: &csi.NodeServiceCapability_Rpc{Rpc: &csi.NodeServiceCapability_RPC{Type: csi.NodeServiceCapability_RPC_STAGE_UNSTAGE_VOLUME}}}}}, nil
}
