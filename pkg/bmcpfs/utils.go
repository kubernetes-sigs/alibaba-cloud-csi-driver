/*
Copyright 2019 The Kubernetes Authors.

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

package bmcpfs

import (
	"fmt"
	"strings"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/common"
	"k8s.io/klog/v2"
)

const (
	// volume context keys
	VolumeType = "volumeType"
	// keys for static volume support
	FileSetID = "filesetId"
)

// updateVolumeContext removes unnecessary volume context parameters
func updateVolumeContext(volumeContext map[string]string) map[string]string {
	cloned := make(map[string]string)
	for k, v := range volumeContext {
		cloned[k] = v
	}

	// Remove unnecessary keys similar to disk implementation
	keysToRemove := []string{
		"csi.alibabacloud.com/lastApply",
		common.PVNameKey,
		common.PVCNameKey,
		common.PVCNamespaceKey,
		"csi.alibabacloud.com/storage-provisioner",
		"csi.alibabacloud.com/reclaimPolicy",
		"csi.alibabacloud.com/storageclassName",
		"allowVolumeExpansion",
		"volume.kubernetes.io/selected-node",
	}

	for _, key := range keysToRemove {
		delete(cloned, key)
	}

	return cloned
}

// parseVolumeID parses the volume ID to extract filesystem ID and fileset ID
// VolumeId format: "filesystemId"
func parseVolumeID(volumeID string) (filesystemID, filesetID string, err error) {
	parts := strings.Split(volumeID, "+")
	if len(parts) != 2 {
		return "", "", fmt.Errorf("invalid volume ID format: %s, expected format: filesystemId+filesetId", volumeID)
	}

	filesystemID = parts[0]
	filesetID = parts[1]

	if filesystemID == "" || filesetID == "" {
		return "", "", fmt.Errorf("invalid volume ID: filesystem ID or fileset ID is empty")
	}

	return filesystemID, filesetID, nil
}

// validateFileSetParameters validates the parameters for fileset creation
func validateFileSetParameters(req *csi.CreateVolumeRequest) error {
	params := req.GetParameters()

	// Check required parameters
	bmcpfsID := params["bmcpfsId"]
	if bmcpfsID == "" {
		return fmt.Errorf("bmcpfsId parameter is required")
	}

	// Validate filesystem type if provided
	for _, cap := range req.GetVolumeCapabilities() {
		mnt := cap.GetMount()
		if mnt != nil && mnt.FsType != "" && mnt.FsType != "bmcpfs" {
			return fmt.Errorf("fsType %s is not supported, only 'cpfs' is supported", mnt.FsType)
		}
	}

	return nil
}

// createVolumeResponse creates a proper CSI volume response
func createVolumeResponse(fileSetID, bmcpfsID string, volSizeBytes int64, volumeContext map[string]string) *csi.Volume {
	// Add volume type to context
	if volumeContext == nil {
		volumeContext = make(map[string]string)
	}
	volumeContext[VolumeType] = "bmcpfs"
	volumeContext[FileSetID] = fileSetID

	klog.Infof("createVolumeResponse: volumeContext: %+v", volumeContext)

	return &csi.Volume{
		CapacityBytes: volSizeBytes,
		VolumeId:      bmcpfsID,
		VolumeContext: volumeContext,
	}
}
