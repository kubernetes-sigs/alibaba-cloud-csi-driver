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
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils/agentidentity"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	mount "k8s.io/mount-utils"
)

type resolvedMount struct {
	Actor   ActorInfo
	Driver  string
	Request *csi.NodePublishVolumeRequest
}

// resolveMount asks the Actor API what the volume actually is and returns the publish
// request that describes it. The identity read off the mount target is the question; the
// Actor API is the only authority on the answer.
func resolveMount(ctx context.Context, lookup ActorLookup, id volumeIdentity, attributes map[string]string) (resolvedMount, error) {
	// 1. Find the Actor that owns the actor directory in the target.
	actor, err := lookupActor(ctx, lookup, id, attributes)
	if err != nil {
		return resolvedMount{}, err
	}
	resolved := resolvedMount{Actor: actor}
	// 2. A Golden Actor is building its image, so it has no backend request to select yet.
	if actor.Golden {
		if actor.Atespace != "ate-golden" || actor.TemplateUID == "" || actor.Name != actor.TemplateUID {
			return resolvedMount{}, status.Error(codes.FailedPrecondition, "golden actor association is invalid")
		}
		return resolved, nil
	}
	// 3. An Actor in the golden atespace that is not golden itself is not verified.
	if actor.Atespace == "ate-golden" {
		return resolvedMount{}, status.Error(codes.FailedPrecondition, "unverified golden actor")
	}
	// 4. Pick this template volume's request out of the Actor annotation. The entry names the
	// driver it was generated for; routing to a backend happens later, on that name.
	request, driver, err := selectPublishEntry(actor.Annotation, id.VolumeName)
	if err != nil {
		return resolvedMount{}, err
	}
	// 5. Refuse anything the bridge may not forward, whatever the backend: the invariants every
	// driver shares. Driver-specific rules are applied by the routed backend's own validator.
	if err := validateCommonRequest(request, actor); err != nil {
		return resolvedMount{}, err
	}
	resolved.Driver = driver
	resolved.Request = request
	return resolved, nil
}

func lookupActor(ctx context.Context, lookup ActorLookup, id volumeIdentity, attributes map[string]string) (ActorInfo, error) {
	if lookup == nil {
		return ActorInfo{}, status.Error(codes.FailedPrecondition, "actor lookup must be configured")
	}
	ref := ActorReference{UID: id.ActorUID, Name: attributes[agentidentity.ActorNameKey], Atespace: attributes[agentidentity.ActorNamespaceKey]}
	if ref.Name == "" || ref.Atespace == "" {
		return ActorInfo{}, status.Error(codes.InvalidArgument, "actor name and namespace are required")
	}
	actor, err := lookup(ctx, ref)
	if err != nil {
		return ActorInfo{}, err
	}
	if actor.UID != id.ActorUID {
		return ActorInfo{}, status.Error(codes.PermissionDenied, "actor lookup returned a different identity")
	}
	return actor, nil
}

// selectPublishEntry decodes the Actor annotation and returns the entry of one template volume,
// together with the driver that entry was generated for. The annotation carries the request the
// shared generator produced for the ACS Sandbox form, so the bridge reads it instead of
// reconstructing it. The bridge does not decide here whether it can serve that driver: it returns
// the name and lets the caller route to a registered backend, so supporting another driver is a
// registration rather than a loosened check.
//
// By design the bridge sets no size limit of its own on that annotation: the object is already
// bounded by the API that stores it, and a limit here would reject a configuration the Actor
// directory itself accepted.
func selectPublishEntry(annotation, volumeName string) (*csi.NodePublishVolumeRequest, string, error) {
	if annotation == "" {
		return nil, "", status.Errorf(codes.FailedPrecondition, "actor has no %s annotation", PublishRequestsAnnotation)
	}
	var entries []publishEntry
	if err := json.Unmarshal([]byte(annotation), &entries); err != nil {
		return nil, "", status.Error(codes.FailedPrecondition, "invalid actor publish requests JSON")
	}
	var matched *publishEntry
	for i := range entries {
		if entries[i].VolumeName != volumeName {
			continue
		}
		if matched != nil {
			return nil, "", status.Errorf(codes.FailedPrecondition, "actor has more than one publish request for volume %q", volumeName)
		}
		matched = &entries[i]
	}
	if matched == nil {
		return nil, "", status.Errorf(codes.FailedPrecondition, "actor has no publish request for volume %q", volumeName)
	}
	request := new(csi.NodePublishVolumeRequest)
	if err := protojson.Unmarshal(matched.Request, request); err != nil {
		return nil, "", status.Error(codes.FailedPrecondition, "invalid CSI publish request")
	}
	return request, matched.Driver, nil
}

// validateCommonRequest enforces the invariants every backend shares, so a driver-specific
// validator only has to reason about its own mount. A field is refused here only when forwarding
// it would either succeed as the wrong mount or make the backend key its work on an empty
// identifier; anything a backend can answer itself is left to it. By design a stored field the
// bridge replaces (the target) or a backend ignores (the staging path) is forwarded as it is
// instead of being checked.
func validateCommonRequest(request *csi.NodePublishVolumeRequest, actor ActorInfo) error {
	// Backends lock and log on the volume ID of the request they receive.
	if request.VolumeId == "" {
		return status.Error(codes.FailedPrecondition, "publish request requires a volume ID")
	}
	// Backends read the mount flags and the access mode off these; a block capability or an
	// unknown mode would be mounted as a writable filesystem rather than rejected.
	if request.GetVolumeCapability().GetMount() == nil {
		return status.Error(codes.FailedPrecondition, "publish request requires a filesystem capability")
	}
	if request.GetVolumeCapability().GetAccessMode().GetMode() == csi.VolumeCapability_AccessMode_UNKNOWN {
		return status.Error(codes.FailedPrecondition, "publish request requires an access mode")
	}
	// A backend may take credentials straight from the request, and the annotation is Actor
	// metadata with a much wider readership than a Secret. Reject credentials before forwarding;
	// the producer must separately prevent persisting them in the Actor annotation in the first
	// place.
	if len(request.Secrets) != 0 {
		return status.Error(codes.FailedPrecondition, "publish request must not carry credentials")
	}
	// Both NAS and OSS reconcile a stored sandboxId against the Actor identity they are handed in
	// substrate mode. A mismatch means the stored request describes a different Actor's sandbox,
	// so refuse it here rather than let a backend mount under the wrong identity. Stored PodInfo
	// needs no such comparison, because backendRequest replaces it with the current placement.
	for key, value := range request.VolumeContext {
		if strings.EqualFold(key, "sandboxId") && value != "" && value != actor.UID {
			return status.Error(codes.PermissionDenied, "publish request carries a different actor identity")
		}
	}
	return validateActorMetadata(request.VolumeContext, actor)
}

// validateNASRequest is the NAS backend's driver-specific gate: an AgenticFS filesystem mounted
// through Agent Identity over alinas, with no client-selection or CNFS routing that would silently
// mount a different volume.
func validateNASRequest(request *csi.NodePublishVolumeRequest, _ ActorInfo) error {
	// Without Agent Identity the mount proxy is never told to use the Actor's own bearer token
	// (pkg/nas/utils.go:379), so the target would be mounted under some other identity; without
	// alinas it would not go through the AccessPoint the request names.
	if !strings.EqualFold(request.VolumeContext["authType"], "agent-identity") || request.VolumeContext["mountProtocol"] != "alinas" {
		return status.Error(codes.FailedPrecondition, "only AgenticFS with Agent Identity is supported")
	}
	for key, value := range request.VolumeContext {
		switch strings.ToLower(key) {
		case "useclient", "containernetworkfilesystem":
			// Both rewrite what NAS mounts and where: the client type and filesystem from the
			// volume context, the server endpoint from the CNFS object (pkg/nas/nodeserver.go:217,
			// :251, :313-322). They do not fail, they mount a different volume.
			if value != "" {
				return status.Error(codes.FailedPrecondition, "NAS client selection and CNFS routing are not supported through the Substrate bridge")
			}
		}
	}
	return nil
}

func validateActorMetadata(values map[string]string, actor ActorInfo) error {
	if !agentidentity.HasActorIdentity(values) {
		return nil
	}
	if values[agentidentity.ActorUIDKey] != actor.UID || values[agentidentity.ActorNameKey] != actor.Name || values[agentidentity.ActorNamespaceKey] != actor.Atespace {
		return status.Error(codes.PermissionDenied, "actor metadata does not match the volume owner")
	}
	return nil
}

func setActorMetadata(values map[string]string, actor ActorInfo) {
	values[agentidentity.ActorUIDKey] = actor.UID
	values[agentidentity.ActorNameKey] = actor.Name
	values[agentidentity.ActorNamespaceKey] = actor.Atespace
}

// mountAt returns the mount currently visible at target. When several mounts share a path the
// last one in the table is the topmost, and that is the one the Actor would see.
func (n *Node) mountAt(target string) (*mount.MountPoint, error) {
	if n.opts.Mounter == nil {
		return nil, status.Error(codes.FailedPrecondition, "mount inspector is required")
	}
	entries, err := n.opts.Mounter.List()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "read mount table: %v", err)
	}
	var live *mount.MountPoint
	for i := range entries {
		if entries[i].Path == target {
			live = &entries[i]
		}
	}
	return live, nil
}

func nasMount(entry *mount.MountPoint) bool {
	return entry.Type == "nfs" || entry.Type == "nfs4" || entry.Type == "alinas"
}

func (n *Node) placeholderSource(target string) (string, error) {
	if !filepath.IsAbs(n.opts.StateDir) || filepath.Clean(n.opts.StateDir) == "/" {
		return "", status.Error(codes.FailedPrecondition, "an absolute dedicated placeholder directory is required")
	}
	key := fmt.Sprintf("%x", sha256.Sum256([]byte(target)))
	return filepath.Join(n.opts.StateDir, "placeholders", key), nil
}

func (n *Node) placeholderMounted(source, target string) (bool, error) {
	info, err := os.Lstat(source)
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	if err != nil {
		return false, status.Errorf(codes.Internal, "inspect placeholder source: %v", err)
	}
	if !info.IsDir() || info.Mode()&os.ModeSymlink != 0 {
		return false, status.Error(codes.FailedPrecondition, "placeholder source is not an owned directory")
	}
	live, err := n.mountAt(target)
	if err != nil || live == nil {
		return false, err
	}
	if live.Device == source {
		return true, nil
	}
	// Bind sources may appear as their backing device in the table. Compare the
	// visible target inode instead of any hidden lower mount or historical reference.
	targetInfo, err := os.Stat(target)
	if err != nil {
		return false, status.Errorf(codes.Internal, "inspect Golden target: %v", err)
	}
	return os.SameFile(info, targetInfo), nil
}

func (n *Node) publishPlaceholder(target string) error {
	live, err := n.mountAt(target)
	if err != nil {
		return err
	}
	source, err := n.placeholderSource(target)
	if err != nil {
		return err
	}
	if live != nil {
		owned, err := n.placeholderMounted(source, target)
		if err != nil {
			return err
		}
		if !owned {
			return status.Error(codes.FailedPrecondition, "Golden target does not reference its expected source")
		}
		return nil
	}
	if info, err := os.Lstat(source); err == nil && (!info.IsDir() || info.Mode()&os.ModeSymlink != 0) {
		return status.Error(codes.FailedPrecondition, "placeholder source is not an owned directory")
	} else if err != nil && !errors.Is(err, os.ErrNotExist) {
		return status.Errorf(codes.Internal, "inspect placeholder source: %v", err)
	}
	if err := os.MkdirAll(source, 0755); err != nil {
		return status.Errorf(codes.Internal, "create placeholder source: %v", err)
	}
	if err := os.MkdirAll(target, 0755); err != nil {
		return status.Errorf(codes.Internal, "create Golden target directory: %v", err)
	}
	if err := n.opts.Mounter.Mount(source, target, "", []string{"bind"}); err != nil {
		return status.Errorf(codes.Internal, "bind mount Golden placeholder: %v", err)
	}
	return nil
}

func (n *Node) unpublishPlaceholder(source, target string) error {
	if err := n.opts.Mounter.Unmount(target); err != nil {
		return status.Errorf(codes.Internal, "unmount Golden placeholder: %v", err)
	}
	// The mount removed here may have been covering a mount the bridge does not own. Only an
	// empty target may be turned back into a plain directory.
	if remaining, err := n.mountAt(target); err != nil {
		return err
	} else if remaining != nil {
		return status.Error(codes.FailedPrecondition, "target still has a mount after removing the Golden placeholder")
	}
	if err := os.Remove(target); err != nil && !errors.Is(err, os.ErrNotExist) {
		return status.Errorf(codes.Internal, "remove Golden target directory: %v", err)
	}
	entries, err := n.opts.Mounter.List()
	if err != nil {
		return status.Errorf(codes.Internal, "read mount table: %v", err)
	}
	for _, entry := range entries {
		if entry.Device == source || entry.Path == source || strings.HasPrefix(entry.Path, source+string(os.PathSeparator)) {
			return status.Error(codes.FailedPrecondition, "refusing to remove mounted placeholder storage")
		}
	}
	refs, err := n.opts.Mounter.GetMountRefs(source)
	if err != nil {
		return status.Errorf(codes.Internal, "list placeholder mount references: %v", err)
	}
	if len(refs) > 0 {
		return status.Error(codes.FailedPrecondition, "placeholder source still has mount references")
	}
	if err := os.RemoveAll(source); err != nil {
		return status.Errorf(codes.Internal, "remove placeholder storage: %v", err)
	}
	return nil
}
