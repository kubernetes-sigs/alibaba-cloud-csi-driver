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
	"context"
	"fmt"
	"strings"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/bmcpfs/internal"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/common"
	"k8s.io/klog/v2"
)

const (
	// keys for static volume support
	annFileSetID = "csi.alibabacloud.com/fileset-id"

	// keys for volume context management
	labelAppendPrefix = "csi.alibabacloud.com/"
	annAppendPrefix   = "csi.alibabacloud.com/"

	// TopologyRegionKey is the key for region topology
	TopologyRegionKey = "region"
	// TopologyZoneKey is the key for zone topology
	TopologyZoneKey = "zone"

	// volume context keys
	labelVolumeType  = "volume-type"
	annVolumeTopoKey = "volume-topology"
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
// VolumeId format: "filesystemId+filesetId"
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
		if mnt != nil && mnt.FsType != "" && mnt.FsType != "cpfs" {
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
	volumeContext[labelAppendPrefix+labelVolumeType] = "bmcpfs"
	volumeContext[annFileSetID] = fileSetID

	klog.Infof("createVolumeResponse: volumeContext: %+v", volumeContext)

	return &csi.Volume{
		CapacityBytes: volSizeBytes,
		VolumeId:      bmcpfsID,
		VolumeContext: volumeContext,
	}
}

// staticFileSetCreate handles static volume creation for existing filesets
func staticFileSetCreate(req *csi.CreateVolumeRequest, fileSetManager internal.CPFSFileSetManager) (*csi.Volume, error) {
	params := req.GetParameters()
	fileSetID := params[annFileSetID]
	if fileSetID == "" {
		return nil, nil
	}

	bmcpfsID := params["bmcpfsId"]
	if bmcpfsID == "" {
		return nil, fmt.Errorf("bmcpfsId parameter is required for static volume")
	}

	// Verify the fileset exists by describing it
	fileSet, err := fileSetManager.DescribeFileset(context.Background(), bmcpfsID, fileSetID)
	if err != nil {
		return nil, fmt.Errorf("failed to describe filesets: %v", err)
	}

	if fileSet == nil {
		return nil, fmt.Errorf("fileset %s not found in filesystem %s", fileSetID, bmcpfsID)
	}

	// Validate size if provided
	volSizeBytes := int64(req.GetCapacityRange().GetRequiredBytes())
	if volSizeBytes > 0 && *fileSet.Quota.SizeLimit != volSizeBytes {
		klog.Warningf("requested size %d does not match existing fileset size %d, using existing size", volSizeBytes, *fileSet.Quota.SizeLimit)
	}

	volumeContext := req.GetParameters()
	volumeContext = updateVolumeContext(volumeContext)
	return createVolumeResponse(fileSetID, bmcpfsID, volSizeBytes, volumeContext), nil
}
