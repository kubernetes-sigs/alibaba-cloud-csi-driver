//go:build !windows

package nas

import (
	"context"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog/v2"
)

// Substrate volume handle format: "substrate-agenticfs:<filesystemID>:<spaceID>".
// The prefix is a persistent contract: parsing must remain for existing
// volumes even if runtime Substrate support is removed.
const substrateAgenticVolumePrefix = "substrate-agenticfs:"

func substrateAgenticVolumeID(filesystemID, spaceID string) string {
	return substrateAgenticVolumePrefix + filesystemID + ":" + spaceID
}

// controllerVolume returns a synthetic PV for a Substrate volume ID, or nil
// when the volume has no backend to reclaim. Substrate volumes have no PV;
// this reconstructs just enough for the NAS controller to recover its
// filesystem and AgenticSpace identifiers.
func (cs *controllerServer) controllerVolume(ctx context.Context, volumeID string) (*corev1.PersistentVolume, error) {
	if strings.HasPrefix(volumeID, substrateAgenticVolumePrefix) {
		parts := strings.SplitN(volumeID[len(substrateAgenticVolumePrefix):], ":", 2)
		if len(parts) != 2 || !validNASResourceID(parts[0]) || !validNASResourceID(parts[1]) {
			return nil, status.Error(codes.InvalidArgument, "invalid Substrate AgenticFS volume handle")
		}
		return &corev1.PersistentVolume{Spec: corev1.PersistentVolumeSpec{
			PersistentVolumeSource: corev1.PersistentVolumeSource{CSI: &corev1.CSIPersistentVolumeSource{
				VolumeHandle: volumeID,
				VolumeAttributes: map[string]string{
					volumeAsKey:         agenticFsVolumeAs,
					filesystemIDKey:     parts[0],
					vcKeyAgenticSpaceId: parts[1],
				},
			}},
		}}, nil
	}
	// Substrate sharepath/subpath volume IDs carry actor-derived identifiers
	// that are not valid PV names; skip backend cleanup.
	if strings.HasPrefix(volumeID, "substrate-") && strings.ContainsAny(volumeID, "/:") {
		klog.InfoS("controllerVolume: Substrate volume ID is not a valid PV name, skipping backend cleanup", "volumeId", volumeID)
		return nil, nil
	}
	pv, err := cs.kubeClient.CoreV1().PersistentVolumes().Get(ctx, volumeID, metav1.GetOptions{})
	if apierrors.IsNotFound(err) && strings.HasPrefix(volumeID, "substrate-") {
		klog.InfoS("controllerVolume: Substrate volume has no PV, skipping backend cleanup", "volumeId", volumeID)
		return nil, nil
	}
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return pv, nil
}

func validNASResourceID(id string) bool {
	if len(id) == 0 || len(id) > 128 {
		return false
	}
	for _, ch := range id {
		switch {
		case ch >= 'a' && ch <= 'z', ch >= 'A' && ch <= 'Z', ch >= '0' && ch <= '9', ch == '-':
		default:
			return false
		}
	}
	return true
}
