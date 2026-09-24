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
)

const substrateAgenticVolumePrefix = "substrate-agenticfs:"

func substrateAgenticVolumeID(filesystemID, spaceID string) string {
	return substrateAgenticVolumePrefix + "v1:" + filesystemID + ":" + spaceID
}

func (cs *controllerServer) controllerVolume(ctx context.Context, volumeID string) (*corev1.PersistentVolume, error) {
	if strings.HasPrefix(volumeID, substrateAgenticVolumePrefix) {
		parts := strings.Split(volumeID, ":")
		if len(parts) != 4 || parts[1] != "v1" || !validNASResourceID(parts[2]) || !validNASResourceID(parts[3]) {
			return nil, status.Error(codes.InvalidArgument, "invalid or unsupported Substrate AgenticFS volume handle")
		}
		filesystemID, agenticSpaceID := parts[2], parts[3]
		// Adapt to the existing controller interface in memory; no Kubernetes PV is created.
		return &corev1.PersistentVolume{Spec: corev1.PersistentVolumeSpec{
			PersistentVolumeSource: corev1.PersistentVolumeSource{CSI: &corev1.CSIPersistentVolumeSource{
				VolumeHandle: volumeID,
				VolumeAttributes: map[string]string{
					volumeAsKey:         agenticFsVolumeAs,
					filesystemIDKey:     filesystemID,
					vcKeyAgenticSpaceId: agenticSpaceID,
				},
			}},
		}}, nil
	}
	pv, err := cs.kubeClient.CoreV1().PersistentVolumes().Get(ctx, volumeID, metav1.GetOptions{})
	if apierrors.IsNotFound(err) && strings.HasPrefix(volumeID, "substrate-") {
		return nil, status.Error(codes.FailedPrecondition,
			"legacy Substrate volume has no PV or backend identifiers: cannot confirm deletion; recover its filesystem/space metadata and follow the ownership-checked manual cleanup procedure")
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
