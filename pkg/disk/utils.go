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

package disk

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"iter"
	"maps"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
	"sync"
	"time"

	ecs20140526 "github.com/alibabacloud-go/ecs-20140526/v7/client"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	v1_credentials "github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials"
	aliyunep "github.com/aliyun/alibaba-cloud-sdk-go/sdk/endpoints"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/container-storage-interface/spec/lib/go/csi"
	volumeSnapshotV1 "github.com/kubernetes-csi/external-snapshotter/client/v8/apis/volumesnapshot/v1"
	snapClientset "github.com/kubernetes-csi/external-snapshotter/client/v8/clientset/versioned"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/metadata"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/common"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	utilshttp "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils/http"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/version"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/client-go/tools/record"
	"k8s.io/klog/v2"
	"k8s.io/utils/ptr"
)

var (
	// KubernetesAlicloudIdentity is the system identity for ecs client request
	KubernetesAlicloudIdentity = fmt.Sprintf("Kubernetes.Alicloud/CsiProvision.Disk-%s", version.VERSION)
	SupportedFilesystemTypes   = sets.New(EXT4_FSTYPE, EXT3_FSTYPE, XFS_FSTYPE, NTFS_FSTYPE)
)

const (
	instanceTypeInfoAnnotation = "alibabacloud.com/instance-type-info"
	// PVC annotation key of KMS key ID, override the storage class parameter kmsKeyId
	KMSKeyID = "alibabacloud.com/kms-key-id"
	// CreateDiskARN ARN parameter of the CreateDisk interface
	CreateDiskARN = "alibabacloud.com/createdisk-arn"
)

// RoleAuth define STS Token Response
type RoleAuth struct {
	AccessKeyID     string
	AccessKeySecret string
	Expiration      time.Time
	SecurityToken   string
	LastUpdated     time.Time
	Code            string
}

var ecsOpenAPIDialer = net.Dialer{
	Timeout:   30 * time.Second,
	KeepAlive: 30 * time.Second,
}

var ecsOpenAPITransport = http.Transport{
	Proxy:                 http.ProxyFromEnvironment,
	DialContext:           ecsOpenAPIDialer.DialContext,
	ForceAttemptHTTP2:     true,
	MaxIdleConns:          100,
	MaxIdleConnsPerHost:   100, // Set this equal to MaxIdleConns as we should only talk to one endpoint with this Transport instance.
	MaxConnsPerHost:       500, // Protect our backend. Should be large enough to handle any workload.
	IdleConnTimeout:       90 * time.Second,
	TLSHandshakeTimeout:   10 * time.Second,
	ExpectContinueTimeout: 1 * time.Second,
}

var ecsEndpointOnce sync.Once

func NewEcsClient(regionID string, cred v1_credentials.CredentialsProvider) (ecsClient *ecs.Client) {
	scheme := "HTTPS"
	if e := os.Getenv("ALICLOUD_CLIENT_SCHEME"); e != "" {
		scheme = e
	}
	config := sdk.NewConfig().WithScheme(scheme).WithUserAgent(KubernetesAlicloudIdentity)

	ecsClient, err := ecs.NewClientWithOptions(regionID, config, cred)
	if err != nil {
		return nil
	}

	ecsClient.SetHTTPSInsecure(false)

	ecsEndpointOnce.Do(func() {
		ep := os.Getenv("ECS_ENDPOINT")
		if ep != "" {
			klog.Infof("Use ECS_ENDPOINT: %s", ep)
		} else if os.Getenv("INTERNAL_MODE") == "true" {
			var err error
			ep, err = cloud.ECSQueryLocalEndpoint(regionID, ecsClient)
			if err != nil {
				klog.Fatalf("Internal mode, but resolve ECS endpoint failed: %v", err)
			}
			klog.Infof("Resolved ECS localAPI endpoint: %s", ep)
		}
		if ep != "" {
			err := aliyunep.AddEndpointMapping(regionID, "Ecs", ep)
			if err != nil {
				klog.Fatalf("Set ECS endpoint failed: %v", err)
			}
		}
	})

	network := utils.GetNetworkType()
	if network != "" {
		klog.Infof("Use ALIBABA_CLOUD_NETWORK_TYPE: %s", network)
		ecsClient.SetEndpointRules(nil, "regional", network)
	}

	var rt http.RoundTripper = &ecsOpenAPITransport
	header := utilshttp.MustParseHeaderEnv("ECS_HEADERS")
	if len(header) > 0 {
		rt = utilshttp.RoundTripperWithHeader(rt, header)
	}
	ecsClient.SetTransport(rt)
	return
}

// IsFileExisting check file exist in volume driver
func IsFileExisting(filename string) bool {
	_, err := os.Stat(filename)
	if err == nil {
		return true
	}
	// Notice: this err may be is not dictionary error, it will returns true
	if os.IsNotExist(err) {
		return false
	}
	return true
}

// IsDirEmpty check whether the given directory is empty
func IsDirEmpty(name string) (bool, error) {
	f, err := os.Open(name)
	if err != nil {
		return false, err
	}
	defer f.Close()

	// read in ONLY one file
	_, err = f.Readdir(1)
	// and if the file is EOF... well, the dir is empty.
	if err == io.EOF {
		return true, nil
	}
	return false, err
}

// GetDeviceByBdf get device name by bdf
func GetDeviceByBdf(bdf string, enLog bool) (device string, err error) {
	virtioPciPath := fmt.Sprintf("/sys/bus/pci/drivers/virtio-pci/%s", bdf)
	dirs, err := os.ReadDir(virtioPciPath)
	if err != nil {
		return "", err
	}
	var virtioNumbers []string
	for _, dir := range dirs {
		if dir.IsDir() && strings.HasPrefix(dir.Name(), "virtio") {
			virtioNumbers = append(virtioNumbers, dir.Name())
		}
	}
	if enLog {
		klog.Infof("Device bdf: %s, virtio numbers: %v", bdf, virtioNumbers)
	}
	if len(virtioNumbers) == 0 {
		return "", fmt.Errorf("virtio device not found, bdf: %s", bdf)
	} else if len(virtioNumbers) > 1 {
		return "", fmt.Errorf("virtio device found multiple: %v, bdf: %s", virtioNumbers, bdf)
	}

	devices, err := filepath.Glob("/sys/block/*/device")
	if err != nil {
		return "", fmt.Errorf("Glob：%v", err)
	}
	for _, device := range devices {
		targetPath, _ := os.Readlink(device)
		if filepath.Base(targetPath) == virtioNumbers[0] {
			devicePath := fmt.Sprintf("/dev/%s", filepath.Base(filepath.Dir(device)))
			if enLog {
				klog.Infof("Device bdf: %s, device: %s", bdf, devicePath)
			}
			return devicePath, nil
		}
	}
	return "", fmt.Errorf("virtio device not found, bdf: %s", bdf)
}

func checkRootAndSubDeviceFS(rootDevicePath, subDevicePath string) error {
	if !utils.IsFileExisting(rootDevicePath) || !utils.IsFileExisting(subDevicePath) {
		return fmt.Errorf("input device path does not exist: %s, %s ", rootDevicePath, subDevicePath)
	}
	fstype, pttype, err := utils.GetDiskFStypePTtype(rootDevicePath)
	if err != nil {
		return fmt.Errorf("GetDiskFormat of root device %s failed: %w", rootDevicePath, err)
	}
	if fstype != "" {
		return fmt.Errorf("root device %s has filesystem: %s, and ptType: %s, disk is not supported ", rootDevicePath, fstype, pttype)
	}

	fstype, _, err = utils.GetDiskFStypePTtype(subDevicePath)
	if err != nil {
		return fmt.Errorf("GetDiskFormat of sub device %s failed: %w", subDevicePath, err)
	}
	if fstype == "" {
		return fmt.Errorf("root device %s has partition, and you should format %s by hands ", rootDevicePath, subDevicePath)
	}
	return nil
}

func prepareMountInfos(req *csi.NodePublishVolumeRequest) ([]string, string) {
	mnt := req.VolumeCapability.GetMount()

	options := []string{"bind"}
	fsType := "ext4"
	if mnt != nil {
		options = append(options, mnt.MountFlags...)
		if mnt.FsType != "" {
			fsType = mnt.FsType
		}
	}

	klog.Infof("prepareMountInfos: VolumeCapability: %+v, req.ReadOnly: %+v", mnt, req.Readonly)
	if req.Readonly {
		options = append(options, "ro")
	}
	return options, fsType
}

// GetVolumeIDByDevice get volumeID by specific device name according to device meta-info
func GetVolumeIDByDevice(device string) (volumeID string, err error) {
	// get volume by serial number feature
	deviceName := device
	if after, ok := strings.CutPrefix(device, "/dev/"); ok {
		deviceName = after
	} else if after, ok := strings.CutPrefix(device, "/"); ok {
		deviceName = after
	}

	virtioSerialFile := filepath.Join("/sys/block/", deviceName, "/serial")
	volumeIDContent := utils.GetFileContent(virtioSerialFile)
	if volumeIDContent != "" {
		return "d-" + volumeIDContent, nil
	}
	nvmeSerialFile := filepath.Join("/sys/block/", deviceName, "device", "/serial")
	volumeIDContent = utils.GetFileContent(nvmeSerialFile)
	if volumeIDContent != "" {
		return "d-" + volumeIDContent, nil
	}

	return "", nil
}

// iterZone iterates zones from given topology requirement.
func iterZone(requirement *csi.TopologyRequirement) iter.Seq[string] {
	return func(yield func(string) bool) {
		if requirement == nil {
			return
		}
		for _, topo := range [][]*csi.Topology{requirement.Preferred, requirement.Requisite} {
			for _, topology := range topo {
				segs := topology.GetSegments()
				for _, key := range []string{TopologyZoneKey, v1.LabelTopologyZone} {
					zone, exists := segs[key]
					if exists && !yield(zone) {
						return
					}
				}
			}
		}
	}
}

func parseTags(params map[string]string) (map[string]string, error) {
	// Note that we cannot assume the iteration order of the map, we must ensure consistent output.
	seenTags := map[string]string{}
	// process old diskTags format first, so that new custom tags can override them consistently
	if v := params["diskTags"]; v != "" {
		for tag := range strings.SplitSeq(v, ",") {
			k, v, found := strings.Cut(tag, ":")
			if !found {
				return nil, fmt.Errorf("invalid diskTags %q, no \":\" found", tag)
			}
			seenTags[k] = v
		}
	}
	// new custom tags
	for k, v := range params {
		if strings.HasPrefix(k, DISK_TAG_PREFIX) {
			seenTags[k[len(DISK_TAG_PREFIX):]] = v
		}
	}
	// k8s PV info as disk tags, override any custom tags
	if v := params[common.PVCNameKey]; v != "" {
		seenTags[common.PVCNameTag] = v
	}
	if v := params[common.PVNameKey]; v != "" {
		seenTags[common.PVNameTag] = v
	}
	if v := params[common.PVCNamespaceKey]; v != "" {
		seenTags[common.PVCNamespaceTag] = v
	}
	return seenTags, nil
}

func pvcRef(params map[string]string) *v1.ObjectReference {
	return &v1.ObjectReference{
		APIVersion: "v1",
		Kind:       "PersistentVolumeClaim",
		Namespace:  params[common.PVCNamespaceKey],
		Name:       params[common.PVCNameKey],
	}
}

var ErrUnknownZone = errors.New("unknown zone")

func getZone(req *csi.CreateVolumeRequest, recorder record.EventRecorder) (string, error) {
	volOptions := req.GetParameters()
	zoneID, ok := volOptions[ZoneID]
	if !ok {
		zoneID, ok = volOptions[strings.ToLower(ZoneID)]
	}
	var paramZone sets.Set[string]
	if ok {
		paramZone = sets.New(strings.Split(zoneID, ",")...)
	}

	// We use the first zone in the intersection of AccessibilityRequirements and Parameter.
	// New user is encouraged to use standard AccessibilityRequirements.
	// For this to work on Kubernetes, --strict-topology needs to be set on external-provisioner
	// to ensure only the selected node is passed in AccessibilityRequirements.

	// Use set instead of slice because we have instanceID in topology,
	// so the same zone may appears multiple times.
	missedZones := sets.New[string]()
	for zone := range iterZone(req.AccessibilityRequirements) {
		if paramZone == nil {
			return zone, nil
		} else if paramZone.Has(zone) {
			return zone, nil
		}
		missedZones.Insert(zone)
	}

	if req.AccessibilityRequirements.GetRequisite() != nil {
		// CSI Spec:
		// If the list of requisite topologies is specified and the SP is
		// unable to to make the provisioned volume available from any of the
		// requisite topologies it MUST fail the CreateVolume call.
		//
		// However, for backward capability, we just record event.
		msg := "no zone info found in accessibility requirements"
		if len(missedZones) > 0 {
			msg = fmt.Sprintf("conflicting zone, parameters specified: %s, accessibility requires: %v", zoneID, sets.List(missedZones))
		}
		recorder.Event(pvcRef(req.Parameters), v1.EventTypeWarning, "ConflictingZone", msg)
	}

	// Best effort to pick what we have
	for zone := range paramZone {
		// Use a random zone from parameters
		return zone, nil
	}
	return "", ErrUnknownZone
}

// getDiskVolumeOptions
func getDiskVolumeOptions(
	req *csi.CreateVolumeRequest,
	m metadata.MetadataProvider,
	recorder record.EventRecorder,
	snapshotId string,
) (*diskVolumeArgs, error) {
	var ok bool
	diskVolArgs := &diskVolumeArgs{
		DiskTags: map[string]string{},
	}
	volOptions := req.GetParameters()

	diskVolArgs.RegionID, ok = volOptions["regionId"]
	if !ok {
		diskVolArgs.RegionID = GlobalConfigVar.Region
	}

	diskVolArgs.NodeSelected = volOptions[NodeScheduleTag]

	for _, cap := range req.GetVolumeCapabilities() {
		mnt := cap.GetMount()
		if mnt != nil && mnt.FsType != "" {
			// Note: skip filesystem type validation when CreateDisk from a snapshot.
			if snapshotId != "" {
				continue
			}
			if SupportedFilesystemTypes.Has(mnt.FsType) {
				continue
			}
			return nil, fmt.Errorf("fsType %s is not supported, please use %v", mnt.FsType, SupportedFilesystemTypes.UnsortedList())
		}
	}

	// disk Type
	diskType, err := validateDiskType(volOptions)
	if err != nil {
		return nil, fmt.Errorf("illegal required parameter type: %s", volOptions[DISK_TYPE])
	}

	if slices.ContainsFunc(diskType, func(t Category) bool { return !AllCategories[t].Regional }) {
		diskVolArgs.ZoneID, err = getZone(req, recorder)
		if err != nil {
			if err == ErrUnknownZone {
				klog.V(1).InfoS("No zone info. Fallback to metadata", "name", req.Name)
				diskVolArgs.ZoneID, err = metadata.GetFallbackZoneID(m)
				if err != nil {
					return nil, fmt.Errorf("no zone info, and failed to get zone id from metadata: %w", err)
				}
			} else {
				return nil, err
			}
		}
		if diskVolArgs.ZoneID == "" {
			return nil, fmt.Errorf("empty zone ID is invalid")
		}
	}
	diskVolArgs.Type = diskType
	pls, err := validateDiskPerformanceLevel(volOptions)
	if err != nil {
		return nil, err
	}
	diskVolArgs.PerformanceLevel = pls

	// readonly, default false
	value, ok := volOptions["readOnly"]
	if !ok {
		diskVolArgs.ReadOnly = false
	} else {
		value = strings.ToLower(value)
		if value == "yes" || value == "true" || value == "1" {
			diskVolArgs.ReadOnly = true
		} else {
			diskVolArgs.ReadOnly = false
		}
	}

	// encrypted or not
	value, ok = volOptions["encrypted"]
	if !ok {
		diskVolArgs.Encrypted = false
	} else {
		value = strings.ToLower(value)
		if value == "yes" || value == "true" || value == "1" {
			diskVolArgs.Encrypted = true
		} else {
			diskVolArgs.Encrypted = false
		}
	}

	// MultiAttach
	{
		multiAttachRequired, err := validateCapabilities(req.VolumeCapabilities)
		if err != nil {
			return nil, err
		}
		diskVolArgs.MultiAttach = false
		if v, ok := volOptions["multiAttach"]; ok {
			switch strings.ToLower(v) {
			case "true", "enabled":
				diskVolArgs.MultiAttach = true
			}
		}
		if multiAttachRequired && !diskVolArgs.MultiAttach {
			return nil, errors.New("multiAttach is required for this access mode." +
				"Please note the limits in https://www.alibabacloud.com/help/en/ecs/user-guide/enable-multi-attach before enabling multiAttach")
		}
	}

	// DiskTags
	diskVolArgs.DiskTags, err = parseTags(volOptions)
	if err != nil {
		return nil, err
	}

	// kmsKeyId
	diskVolArgs.KMSKeyID, ok = volOptions[KMSKeyID]
	if !ok {
		diskVolArgs.KMSKeyID = volOptions["kmsKeyId"]
	}

	if arnStr, ok := volOptions[CreateDiskARN]; ok {
		if err := json.Unmarshal([]byte(arnStr), &diskVolArgs.ARN); err != nil {
			return nil, fmt.Errorf("failed to unmarshal arn, string: %s, err: %v", arnStr, err)
		}
	}

	// resourceGroupId
	diskVolArgs.ResourceGroupID, ok = volOptions["resourceGroupId"]
	if !ok {
		diskVolArgs.ResourceGroupID = ""
	}

	diskVolArgs.StorageClusterID, ok = volOptions["storageClusterId"]
	if !ok {
		diskVolArgs.StorageClusterID = ""
	}

	if diskVolArgs.StorageClusterID != "" {
		if len(diskVolArgs.PerformanceLevel) == 0 {
			return nil, fmt.Errorf("performanceLevel is necessary when storageClusterID: '%s' specified", diskVolArgs.StorageClusterID)
		}
	}

	// volumeExpandAutoSnapshot parameter is obsolete and has no effect now

	if value, ok = volOptions[VOLUME_DELETE_AUTO_SNAPSHOT_OP_RETENT_DAYS_KEY]; ok {
		iValue, err := strconv.Atoi(value)
		if err != nil {
			return nil, fmt.Errorf("getDiskVolumeOptions: parameters volumeDeleteSnapshotRetentionDays[%s] is illegal", value)
		}
		if iValue <= SNAPSHOT_MAX_RETENTION_DAYS && iValue >= SNAPSHOT_MIN_RETENTION_DAYS {
			diskVolArgs.DelAutoSnap = value
		}
	}

	value, ok = volOptions[PROVISIONED_IOPS_KEY]
	if ok {
		iValue, err := strconv.ParseInt(value, 10, 64)
		if err != nil || iValue < 0 {
			return nil, fmt.Errorf("getDiskVolumeOptions: parameters provisionedIops[%s] is illegal", value)
		}
		diskVolArgs.ProvisionedIops = iValue
	}

	diskVolArgs.BurstingEnabled = false
	value, ok = volOptions[BURSTING_ENABLED_KEY]
	if ok {
		value = strings.ToLower(value)
		if value == "yes" || value == "true" || value == "1" {
			diskVolArgs.BurstingEnabled = true
		}
	}

	if req.GetCapacityRange() == nil {
		return nil, fmt.Errorf("capacity range is required")
	}
	volSizeBytes := int64(req.GetCapacityRange().GetRequiredBytes())
	requestGB := (volSizeBytes + 1024*1024*1024 - 1) / (1024 * 1024 * 1024)
	if requestGB < MinimumDiskSizeInGB {
		switch strings.ToLower(volOptions["volumeSizeAutoAvailable"]) {
		case "yes", "true", "1":
			klog.Infof("CreateVolume: volume size was less than allowed limit. Setting request Size to %vGB. volumeSizeAutoAvailable is set.", MinimumDiskSizeInGB)
			requestGB = MinimumDiskSizeInGB
		}
	}
	diskVolArgs.RequestGB = requestGB

	return diskVolArgs, nil
}

func validateDiskType(opts map[string]string) (diskType []Category, err error) {
	if value, ok := opts[DISK_TYPE]; !ok || (ok && value == DiskHighAvail) {
		diskType = []Category{DiskSSD, DiskEfficiency}
		return
	}
	for cusType := range strings.SplitSeq(opts[DISK_TYPE], ",") {
		c := Category(cusType)
		if _, ok := AllCategories[c]; ok {
			diskType = append(diskType, c)
		} else {
			return nil, fmt.Errorf("illegal required parameter type: %s", cusType)
		}
	}
	if len(diskType) == 0 {
		return diskType, fmt.Errorf("illegal required parameter type: %s", opts[DISK_TYPE])
	}
	return
}

func validateDiskPerformanceLevel(opts map[string]string) ([]PerformanceLevel, error) {
	opt := opts[ESSD_PERFORMANCE_LEVEL]
	if opt == "" {
		return nil, nil
	}
	klog.Infof("validateDiskPerformanceLevel: pl: %v", opt)
	allPLs := AllCategories[DiskESSD].PerformanceLevel
	plsStr := strings.Split(opt, ",")
	pls := make([]PerformanceLevel, 0, len(plsStr))
	for _, plStr := range plsStr {
		pl := PerformanceLevel(plStr)
		if _, ok := allPLs[pl]; !ok {
			return nil, fmt.Errorf("illegal performance level type: %s", plStr)
		}
		pls = append(pls, pl)
	}
	return pls, nil
}

// return if MultiAttach is required
func validateCapabilities(capabilities []*csi.VolumeCapability) (bool, error) {
	multiAttachRequired := false
	for _, cap := range capabilities {
		if cap.GetAccessMode() == nil {
			continue
		}
		switch cap.AccessMode.Mode {
		case csi.VolumeCapability_AccessMode_SINGLE_NODE_WRITER,
			csi.VolumeCapability_AccessMode_SINGLE_NODE_READER_ONLY:
			// single node mode is always supported
			continue
		case csi.VolumeCapability_AccessMode_MULTI_NODE_READER_ONLY:
			// multi node read only mode is always supported if MultiAttach is enabled
			multiAttachRequired = true
			continue
		case csi.VolumeCapability_AccessMode_MULTI_NODE_MULTI_WRITER,
			csi.VolumeCapability_AccessMode_MULTI_NODE_SINGLE_WRITER:
			// only supported on block volume
			multiAttachRequired = true
			if _, ok := cap.AccessType.(*csi.VolumeCapability_Block); !ok {
				return multiAttachRequired, errors.New("multi-node writing is only supported for block volume. " +
					"For Kubernetes users, if unsure, use ReadWriteOnce access mode in PersistentVolumeClaim for disk volume")
			}
		default:
			return multiAttachRequired, fmt.Errorf("volume capability %v is not supported", cap.AccessMode.Mode)
		}
	}
	return multiAttachRequired, nil
}

func parseMutableParameters(mutableParameters map[string]string) (ModifyParameters, error) {
	var params ModifyParameters
	for k, v := range mutableParameters {
		switch k {
		case DISK_TYPE:
			params.Category = Category(v)
		case ESSD_PERFORMANCE_LEVEL:
			params.PerformanceLevel = PerformanceLevel(v)
		case PROVISIONED_IOPS_KEY:
			iops, err := strconv.Atoi(v)
			if err != nil {
				return params, fmt.Errorf("invalid %s: %w", PROVISIONED_IOPS_KEY, err)
			}
			params.ProvisionedIops = &iops
		case BURSTING_ENABLED_KEY:
			en, err := strconv.ParseBool(v)
			if err != nil {
				return params, fmt.Errorf("invalid %s: %w", BURSTING_ENABLED_KEY, err)
			}
			params.BurstingEnabled = &en
		default:
			switch {
			case strings.HasPrefix(k, DISK_TAG_PREFIX):
				tagKey := k[len(DISK_TAG_PREFIX):]
				params.Tags = append(params.Tags, ecs.TagResourcesTag{
					Key:   tagKey,
					Value: v,
				})
			case strings.HasPrefix(k, REMOVE_DISK_TAG_PREFIX):
				tagKey := k[len(REMOVE_DISK_TAG_PREFIX):]
				params.RemoveTags = append(params.RemoveTags, tagKey)
			default:
				return params, fmt.Errorf("unknown parameter %s", k)
			}
		}
	}
	return params, nil
}

func importMutableParameters(params *diskVolumeArgs, mutable *ModifyParameters) {
	if len(mutable.Category) > 0 {
		params.Type = []Category{mutable.Category}
	}
	if len(mutable.PerformanceLevel) > 0 {
		params.PerformanceLevel = []PerformanceLevel{mutable.PerformanceLevel}
	}
	if mutable.ProvisionedIops != nil {
		params.ProvisionedIops = int64(*mutable.ProvisionedIops)
	}
	if mutable.BurstingEnabled != nil {
		params.BurstingEnabled = *mutable.BurstingEnabled
	}
	for _, tag := range mutable.RemoveTags {
		delete(params.DiskTags, tag)
	}
	for _, tag := range mutable.Tags {
		params.DiskTags[tag.Key] = tag.Value
	}
}

func getBlockDeviceCapacity(devicePath string) int64 {

	file, err := os.Open(devicePath)
	if err != nil {
		klog.Errorf("getBlockDeviceCapacity:: failed to open devicePath: %v", devicePath)
		return 0
	}
	pos, err := file.Seek(0, io.SeekEnd)
	if err != nil {
		klog.Errorf("getBlockDeviceCapacity:: failed to read devicePath: %v", devicePath)
		return 0
	}
	return pos
}

func appendDiskTypes(resp *ecs20140526.DescribeAvailableResourceResponse, zoneID string, types []string) []string {
	if resp == nil || resp.Body == nil || resp.Body.AvailableZones == nil {
		return types
	}
	for _, zone := range resp.Body.AvailableZones.AvailableZone {
		if zone == nil || zone.AvailableResources == nil {
			continue
		}
		if ptr.Deref(zone.ZoneId, "") != zoneID {
			continue
		}
		for _, resource := range zone.AvailableResources.AvailableResource {
			if resource == nil || resource.SupportedResources == nil {
				continue
			}
			if ptr.Deref(resource.Type, "") != describeResourceType {
				continue
			}
			for _, supported := range resource.SupportedResources.SupportedResource {
				if supported == nil || supported.Value == nil {
					continue
				}
				types = append(types, *supported.Value)
			}
		}
	}
	return types
}

// GetAvailableDiskTypes queries ECS for the disk categories supported by the
// given instance type, combining zonal and regional results. zoneID may be
// empty to skip the zonal query.
func GetAvailableDiskTypes(ctx context.Context, c cloud.ECSv2Interface, instanceType, zoneID, regionID string) ([]string, error) {
	logger := klog.FromContext(ctx)

	var types []string
	req := &ecs20140526.DescribeAvailableResourceRequest{
		InstanceType:        &instanceType,
		DestinationResource: new(describeResourceType),
		ZoneId:              &zoneID,
		ResourceType:        new("disk"),
		RegionId:            &regionID,
	}
	resp, err := c.DescribeAvailableResource(req)
	if err != nil {
		return nil, fmt.Errorf("failed to DescribeAvailableResource(%s) for instance type %s: %w", zoneID, instanceType, err)
	}
	logger.V(4).Info("DescribeAvailableResource", "zoneID", zoneID, "response", resp)
	types = appendDiskTypes(resp, zoneID, types)

	reqRegional := &ecs20140526.DescribeAvailableResourceRequest{
		InstanceType:        &instanceType,
		DestinationResource: new(describeResourceType),
		ResourceType:        new("disk"),
		Scope:               new("region"),
		RegionId:            &regionID,
	}
	respRegional, err := c.DescribeAvailableResource(reqRegional)
	if err != nil {
		return nil, fmt.Errorf("failed to DescribeAvailableResource(region) for instance type %s: %w", instanceType, err)
	}
	logger.V(4).Info("DescribeAvailableResource(region)", "response", respRegional)
	types = appendDiskTypes(respRegional, "", types)

	if len(types) == 0 {
		return nil, fmt.Errorf("no supported disk type found for instance type %s in zone %q: requestId: %s, %s",
			instanceType, zoneID, ptr.Deref(resp.Body.RequestId, ""), ptr.Deref(respRegional.Body.RequestId, ""))
	}
	return types, nil
}

func hasDiskTypeLabel(node *v1.Node) bool {
	for k := range node.Labels {
		if strings.HasPrefix(k, NodeDiskTypeLabelPrefix) {
			return true
		}
	}
	return false
}

// BuildNodePatch returns a strategic-merge patch that sets the given annotation
// key and disk-type labels, or nil if the node already has the expected values.
func BuildNodePatch(node *v1.Node, maxVolumes int, diskTypes []string, annotationKey string) []byte {
	maxStr := strconv.Itoa(maxVolumes)
	need := node.Annotations[annotationKey] != maxStr

	wantLabels := make(map[string]string, len(diskTypes))
	for _, t := range diskTypes {
		wantLabels[NodeDiskTypeLabelPrefix+t] = "available"
	}
	if !need {
		for k, v := range wantLabels {
			if node.Labels[k] != v {
				need = true
				break
			}
		}
	}
	if !need {
		return nil
	}
	patch := map[string]any{
		"metadata": map[string]any{
			"labels": wantLabels,
			"annotations": map[string]string{
				annotationKey: maxStr,
			},
		},
	}
	b, err := json.Marshal(patch)
	if err != nil {
		klog.Fatalf("failed to marshal patch json: %v", err)
	}
	return b
}

func patchForNode(node *v1.Node, maxVolumesNum int, diskTypes []string) []byte {
	return BuildNodePatch(node, maxVolumesNum, diskTypes, NodeDiskCountAnnotation)
}

func volumeCreate(attempt createAttempt, diskID string, volSizeBytes int64, volumeContext map[string]string, zoneID string, contextSource *csi.VolumeContentSource) *csi.Volume {
	segments := map[string]string{}
	cateDesc := AllCategories[attempt.Category]
	volumeContext[labelAppendPrefix+TopologyRegionKey] = GlobalConfigVar.Region
	if cateDesc.Regional {
		segments[RegionalDiskTopologyKey] = GlobalConfigVar.Region
	} else {
		segments[ZonalDiskTopologyKey] = zoneID
		volumeContext[labelAppendPrefix+TopologyZoneKey] = zoneID
	}
	if attempt.Instance != "" {
		segments[common.ECSInstanceIDTopologyKey] = attempt.Instance
	}

	accessibleTopology := []*csi.Topology{{Segments: segments}}
	if attempt.Category != "" {
		// Add PV Label
		if attempt.Category == DiskESSD && attempt.PerformanceLevel == "" {
			attempt.PerformanceLevel = "PL1"
		}
		// TODO delete performanceLevel key
		// delete(volumeContext, "performanceLevel")
		volumeContext[labelAppendPrefix+labelVolumeType] = attempt.String()
		// TODO delete type key
		// delete(volumeContext, "type")

		// Add PV NodeAffinity
		labelKey := NodeDiskTypeLabelPrefix + string(attempt.Category)
		expressions := []v1.NodeSelectorRequirement{{
			Key:      labelKey,
			Operator: v1.NodeSelectorOpIn,
			Values:   []string{"available"},
		}}
		terms := []v1.NodeSelectorTerm{{
			MatchExpressions: expressions,
		}}
		diskTypeTopo := &v1.NodeSelector{
			NodeSelectorTerms: terms,
		}
		diskTypeTopoBytes, _ := json.Marshal(diskTypeTopo)
		volumeContext[annAppendPrefix+annVolumeTopoKey] = string(diskTypeTopoBytes)
	}

	klog.Infof("volumeCreate: volumeContext: %+v", volumeContext)
	tmpVol := &csi.Volume{
		CapacityBytes:      volSizeBytes,
		VolumeId:           diskID,
		VolumeContext:      volumeContext,
		AccessibleTopology: accessibleTopology,
		ContentSource:      contextSource,
	}

	return tmpVol
}

func parseSnapshotID(req *csi.CreateVolumeRequest) (string, error) {
	volumeSource := req.GetVolumeContentSource()
	if volumeSource != nil {
		sourceSnapshot := volumeSource.GetSnapshot()
		if sourceSnapshot == nil {
			return "", fmt.Errorf("CreateVolume: unsupported volumeContentSource type: %T", volumeSource.Type)
		}
		return sourceSnapshot.GetSnapshotId(), nil
	}
	// set snapshotID if pvc labels/annotation set it.
	return req.Parameters[DiskSnapshotID], nil
}

func volumeContentSource(snapshotID string) *csi.VolumeContentSource {
	if snapshotID == "" {
		return nil
	}
	return &csi.VolumeContentSource{
		Type: &csi.VolumeContentSource_Snapshot{
			Snapshot: &csi.VolumeContentSource_SnapshotSource{
				SnapshotId: snapshotID,
			},
		},
	}
}

// staticVolumeCreate 检查输入参数，如果包含了云盘ID，则直接使用云盘进行返回；
// 根据云盘ID请求云盘的具体属性，并作为pv参数返回；
func staticVolumeCreate(req *csi.CreateVolumeRequest) (*csi.Volume, error) {
	paras := req.GetParameters()
	diskID := paras[annDiskID]
	if diskID == "" {
		return nil, nil
	}

	ecsClient := GlobalConfigVar.EcsClient
	disk, err := findDiskByID(diskID, ecsClient)
	if err != nil {
		return nil, err
	}
	if disk == nil {
		return nil, fmt.Errorf("disk %s cannot be found from ecs api", diskID)
	}

	volumeContext := req.GetParameters()
	volumeContext = updateVolumeContext(volumeContext)
	volumeContext["type"] = disk.Category
	volSizeBytes := int64(req.GetCapacityRange().GetRequiredBytes())
	diskSizeBytes := utils.Gi2Bytes(int64(disk.Size))
	if volSizeBytes != diskSizeBytes {
		return nil, fmt.Errorf("disk %s is not expected capacity: expected(%d), disk(%d)", diskID, volSizeBytes, diskSizeBytes)
	}

	// Set VolumeContentSource
	snapshotID, err := parseSnapshotID(req)
	if err != nil {
		return nil, err
	}

	tags, err := parseTags(req.Parameters)
	if err != nil {
		return nil, err
	}
	tagDiskUserTags(diskID, tags)

	attempt := createAttempt{
		Category(disk.Category), PerformanceLevel(disk.PerformanceLevel),
		"", // no instanceID for virtual-kubelet.
	}
	return volumeCreate(attempt, diskID, volSizeBytes, volumeContext, disk.ZoneId, volumeContentSource(snapshotID)), nil
}

// updateVolumeContext remove unnecessary volume context
func updateVolumeContext(volumeContext map[string]string) map[string]string {

	cloned := maps.Clone(volumeContext)
	for _, key := range []string{
		LastApplyKey,
		common.PVNameKey,
		common.PVCNameKey,
		common.PVCNamespaceKey,
		StorageProvisionerKey, "csi.alibabacloud.com/reclaimPolicy",
		"csi.alibabacloud.com/storageclassName",
		"allowVolumeExpansion", "volume.kubernetes.io/selected-node"} {

		delete(cloned, key)
	}

	return cloned
}

// GetAttachedCloudDisks queries ECS (v2 API) for cloud disks attached to the
// given instance, excluding detaching disks and local-only categories.
func GetAttachedCloudDisks(ecsClient cloud.ECSv2Interface, instanceID, regionID string) ([]string, error) {
	req := &ecs20140526.DescribeDisksRequest{
		InstanceId: &instanceID,
		RegionId:   &regionID,
		MaxResults: new(int32(100)),
	}
	var ids []string
	for {
		resp, err := ecsClient.DescribeDisks(req)
		if err != nil {
			return nil, fmt.Errorf("DescribeDisks: %w", err)
		}
		if resp == nil || resp.Body == nil {
			break
		}
		if resp.Body.Disks != nil {
			for _, disk := range resp.Body.Disks.Disk {
				if disk == nil || disk.DiskId == nil {
					continue
				}
				if ptr.Deref(disk.Status, "") == "Detaching" {
					continue
				}
				category := ptr.Deref(disk.Category, "")
				if strings.HasPrefix(category, "cloud") || strings.HasPrefix(category, "elastic_ephemeral_disk") {
					ids = append(ids, *disk.DiskId)
				}
			}
		}
		next := ptr.Deref(resp.Body.NextToken, "")
		if next == "" {
			break
		}
		req.NextToken = resp.Body.NextToken
	}
	return ids, nil
}

// CollectManagedVolumes extracts CSI-managed disk IDs from node status.
func CollectManagedVolumes(node *v1.Node) sets.Set[string] {
	prefix := fmt.Sprintf("kubernetes.io/csi/%s^", DriverName)
	managed := sets.New[string]()
	extract := func(name string) {
		if id, found := strings.CutPrefix(name, prefix); found {
			managed.Insert(id)
		}
	}
	for _, v := range node.Status.VolumesInUse {
		extract(string(v))
	}
	for _, v := range node.Status.VolumesAttached {
		extract(string(v.Name))
	}
	return managed
}

func getVolumeCountFromLabelerAnnotation(node *v1.Node) (int, error) {
	v, ok := node.Annotations[MaxDiskAnnotation]
	if !ok {
		return 0, fmt.Errorf("max-disk annotation not found")
	}
	n, err := strconv.Atoi(v)
	if err != nil {
		return 0, fmt.Errorf("invalid max-disk annotation: %w", err)
	}
	return n, nil
}

func checkOption(opt string) bool {
	switch opt {
	case "enable", "true", "yes":
		return true
	default:
		return false
	}
}

func makeVolumeSnapshot(snapName string, snapContentName string) *volumeSnapshotV1.VolumeSnapshot {
	vs := &volumeSnapshotV1.VolumeSnapshot{
		ObjectMeta: metav1.ObjectMeta{
			Name: snapName,
		},
		Spec: volumeSnapshotV1.VolumeSnapshotSpec{
			Source: volumeSnapshotV1.VolumeSnapshotSource{
				VolumeSnapshotContentName: &snapContentName,
			},
		},
	}
	return vs
}

func makeVolumeSnapshotContent(snapName, snapContentName, snapshotID string) *volumeSnapshotV1.VolumeSnapshotContent {
	vs := &volumeSnapshotV1.VolumeSnapshotContent{
		ObjectMeta: metav1.ObjectMeta{
			Name: snapContentName,
		},
		Spec: volumeSnapshotV1.VolumeSnapshotContentSpec{
			VolumeSnapshotRef: v1.ObjectReference{
				APIVersion: "snapshot.storage.k8s.io/v1",
				Kind:       "VolumeSnapshot",
				Name:       snapName,
				Namespace:  "default",
			},
			DeletionPolicy: volumeSnapshotV1.VolumeSnapshotContentDelete,
			Source: volumeSnapshotV1.VolumeSnapshotContentSource{
				SnapshotHandle: &snapshotID,
			},
			Driver: "diskplugin.csi.alibabacloud.com",
		},
	}
	return vs
}

func createStaticSnap(volumeID, snapshotID string, snapClient snapClientset.Interface) error {

	volumeSnapshotName := fmt.Sprintf("%s-delprotect", volumeID)
	volumeSnapshotContentName := fmt.Sprintf("%s-delprotect-content", volumeID)

	volumeSnapshot := makeVolumeSnapshot(volumeSnapshotName, volumeSnapshotContentName)
	volumeSnapshotContent := makeVolumeSnapshotContent(volumeSnapshotName, volumeSnapshotContentName, snapshotID)

	_, err := snapClient.SnapshotV1().VolumeSnapshots("default").Create(context.Background(), volumeSnapshot, metav1.CreateOptions{})
	if err != nil {
		return err
	}
	_, err = snapClient.SnapshotV1().VolumeSnapshotContents().Create(context.Background(), volumeSnapshotContent, metav1.CreateOptions{})
	if err != nil {
		return err
	}
	return nil
}

// Logging disk block device size
type DiskSize struct {
	Bytes int64
}

func (d DiskSize) String() string {
	// Alibaba cloud disks are at least in the order of GiB
	if d.Bytes%GBSIZE == 0 {
		return fmt.Sprintf("%d GiB", d.Bytes/GBSIZE)
	}
	return fmt.Sprintf("%.3f GiB (0x%X)", float64(d.Bytes)/GBSIZE, d.Bytes)
}
