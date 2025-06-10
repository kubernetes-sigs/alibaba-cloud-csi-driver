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
	"os"
	"path"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
	"time"

	aliyunep "github.com/aliyun/alibaba-cloud-sdk-go/sdk/endpoints"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/containerd/ttrpc"
	volumeSnapshotV1 "github.com/kubernetes-csi/external-snapshotter/client/v8/apis/volumesnapshot/v1"
	snapClientset "github.com/kubernetes-csi/external-snapshotter/client/v8/clientset/versioned"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/metadata"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/common"
	proto "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/disk/proto"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	utilshttp "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils/http"
	utilsio "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils/io"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/version"
	perrors "github.com/pkg/errors"
	"golang.org/x/sys/unix"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/client-go/tools/record"
	"k8s.io/klog/v2"
	k8smount "k8s.io/mount-utils"
)

var (
	// KubernetesAlicloudIdentity is the system identity for ecs client request
	KubernetesAlicloudIdentity = fmt.Sprintf("Kubernetes.Alicloud/CsiProvision.Disk-%s", version.VERSION)
	SupportedFilesystemTypes   = sets.New(EXT4_FSTYPE, EXT3_FSTYPE, XFS_FSTYPE, NTFS_FSTYPE)
)

const (
	DISK_TAG_PREFIX            = "diskTags/"
	instanceTypeInfoAnnotation = "alibabacloud.com/instance-type-info"
)

// DefaultOptions is the struct for access key
type DefaultOptions struct {
	Global struct {
		KubernetesClusterTag string
		AccessKeyID          string `json:"accessKeyID"`
		AccessKeySecret      string `json:"accessKeySecret"`
		Region               string `json:"region"`
	}
}

// RoleAuth define STS Token Response
type RoleAuth struct {
	AccessKeyID     string
	AccessKeySecret string
	Expiration      time.Time
	SecurityToken   string
	LastUpdated     time.Time
	Code            string
}

// InstanceTypeInfo is the annotation value of "alibabacloud.com/instance-type-info"
type InstanceTypeInfo struct {
	DiskQuantity int
}

func newEcsClient(regionID string, ac utils.AccessControl) (ecsClient *ecs.Client) {
	var err error
	switch ac.UseMode {
	case utils.AccessKey:
		ecsClient, err = ecs.NewClientWithAccessKey(regionID, ac.AccessKeyID, ac.AccessKeySecret)
	case utils.Credential:
		ecsClient, err = ecs.NewClientWithOptions(regionID, ac.Config, ac.Credential)
	default:
		ecsClient, err = ecs.NewClientWithStsToken(regionID, ac.AccessKeyID, ac.AccessKeySecret, ac.StsToken)
	}
	scheme := "HTTPS"
	if os.Getenv("ALICLOUD_CLIENT_SCHEME") == "HTTP" {
		scheme = "HTTP"
	}
	ecsClient.SetHTTPSInsecure(false)
	ecsClient.GetConfig().WithScheme(scheme)
	if err != nil {
		return nil
	}

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

	network := utils.GetNetworkType()
	if network != "" {
		klog.Infof("Use ALIBABA_CLOUD_NETWORK_TYPE: %s", network)
		ecsClient.SetEndpointRules(nil, "regional", network)
	}

	header := utilshttp.MustParseHeaderEnv("ECS_HEADERS")
	if len(header) > 0 {
		ecsClient.SetTransport(utilshttp.RoundTripperWithHeader(nil, header))
	}
	return
}

func updateEcsClient(client *ecs.Client) *ecs.Client {
	ac := utils.GetAccessControl()
	if ac.UseMode == utils.ManagedToken {
		client = newEcsClient(GlobalConfigVar.Region, ac)
	}
	if client.Client.GetConfig() != nil {
		client.Client.GetConfig().UserAgent = KubernetesAlicloudIdentity
	}
	return client
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
	if strings.HasPrefix(device, "/dev/") {
		deviceName = strings.TrimPrefix(device, "/dev/")
	} else if strings.HasPrefix(device, "/") {
		deviceName = strings.TrimPrefix(device, "/")
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

// get diskID
func getVolumeConfig(volumeID string) string {
	volumeFile := path.Join(VolumeDir, volumeID+".conf")
	if !utils.IsFileExisting(volumeFile) {
		return ""
	}

	value, err := os.ReadFile(volumeFile)
	if err != nil {
		return ""
	}
	devicePath := strings.TrimSpace(string(value))
	return devicePath
}

// save diskID and volume name
func saveVolumeConfig(volumeID, devicePath string) error {
	if err := utils.CreateDest(VolumeDir); err != nil {
		return err
	}
	if err := utils.CreateDest(VolumeDirRemove); err != nil {
		return err
	}
	if err := removeVolumeConfig(volumeID); err != nil {
		return err
	}
	// cleanup all config files that is pointing to devicePath. Such files may be leaked
	// if previous UnstageVolume is skipped. This is possible if the VolumeDevice is
	// detached by others (e.g. DeleteVolume).
	files, err := os.ReadDir(VolumeDir)
	if err != nil {
		return err
	}
	for _, file := range files {
		if file.Type().IsRegular() && strings.HasSuffix(file.Name(), ".conf") {
			tmpVolID := strings.TrimSuffix(file.Name(), ".conf")
			if getVolumeConfig(tmpVolID) == devicePath {
				if err := removeVolumeConfig(tmpVolID); err != nil {
					return fmt.Errorf("failed to remove volume config for %s: %w", tmpVolID, err)
				}
			}
		}
	}

	volumeFile := path.Join(VolumeDir, volumeID+".conf")
	if err := os.WriteFile(volumeFile, []byte(devicePath), 0644); err != nil {
		return err
	}
	return nil
}

// move config file to remove dir
func removeVolumeConfig(volumeID string) error {
	volumeFile := path.Join(VolumeDir, volumeID+".conf")
	if utils.IsFileExisting(volumeFile) {
		timeStr := time.Now().Format("2006-01-02-15:04:05")
		removeFile := path.Join(VolumeDirRemove, volumeID+"-"+timeStr+".conf")
		if err := os.Rename(volumeFile, removeFile); err != nil {
			return err
		}
	}
	return nil
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
				for _, key := range []string{TopologyZoneKey, common.TopologyKeyZone} {
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
		for _, tag := range strings.Split(v, ",") {
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
func getDiskVolumeOptions(req *csi.CreateVolumeRequest, m metadata.MetadataProvider, recorder record.EventRecorder) (*diskVolumeArgs, error) {
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
		if mnt != nil && mnt.FsType != "" && !SupportedFilesystemTypes.Has(mnt.FsType) {
			return nil, fmt.Errorf("fsType %s is not supported, please use %v", mnt.FsType, SupportedFilesystemTypes.UnsortedList())
		}
	}

	// disk Type
	diskType, err := validateDiskType(volOptions)
	if err != nil {
		return nil, fmt.Errorf("illegal required parameter type: %s", volOptions["type"])
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
	if value, ok := opts["type"]; !ok || (ok && value == DiskHighAvail) {
		diskType = []Category{DiskSSD, DiskEfficiency}
		return
	}
	for _, cusType := range strings.Split(opts["type"], ",") {
		c := Category(cusType)
		if _, ok := AllCategories[c]; ok {
			diskType = append(diskType, c)
		} else {
			return nil, fmt.Errorf("illegal required parameter type: %s", cusType)
		}
	}
	if len(diskType) == 0 {
		return diskType, fmt.Errorf("illegal required parameter type: %s", opts["type"])
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

func getMountedVolumeDevice(mnts []k8smount.MountInfo, targetPath string) string {
	for _, mnt := range mnts {
		if mnt.MountPoint == targetPath {
			return mnt.Root
		}
	}
	return ""
}

func isDeviceMountedAt(mnts []k8smount.MountInfo, device, targetPath string) bool {
	for _, mnt := range mnts {
		if mnt.MountPoint == targetPath && mnt.Source == device {
			return true
		}
	}
	return false
}

const mountInfoPath = "/proc/self/mountinfo"

func CheckDeviceAvailable(devicePath, volumeID, targetPath string) error {
	return checkDeviceAvailable(mountInfoPath, devicePath, volumeID, targetPath)
}

func checkDeviceAvailable(mountinfoPath, devicePath, volumeID, targetPath string) error {
	if devicePath == "" {
		return fmt.Errorf("devicePath is empty, cannot used for Volume")
	}

	mnts, err := k8smount.ParseMountInfo(mountinfoPath)
	if err != nil {
		return err
	}

	// TODO: remove in next version
	// Since devicePath has been change to symlink, we will never run into the following logics.
	// block volume
	if devicePath == "devtmpfs" {
		device := getMountedVolumeDevice(mnts, targetPath)
		newVolumeID, err := GetVolumeIDByDevice(device)
		if err != nil {
			return nil
		}
		if newVolumeID != volumeID {
			return fmt.Errorf("device [%s] associate with volumeID: [%s] rather than volumeID: [%s]", device, newVolumeID, volumeID)
		}

		return nil
	}

	if !utils.IsFileExisting(devicePath) {
		return fmt.Errorf("devicePath(%s) is empty, cannot used for Volume", devicePath)
	}

	// check the device is used for system
	if devicePath == "/dev/vda" || devicePath == "/dev/vda1" {
		klog.Warningf("checkDeviceAvailable: devicePath(%s) may be system device: %s", devicePath, volumeID)
	}

	if isDeviceMountedAt(mnts, devicePath, utils.KubeletRootDir) {
		return fmt.Errorf("devicePath(%s) is used as DataDisk for kubelet, cannot used for Volume", devicePath)
	}
	return nil
}

// GetVolumeDeviceName get device name
func GetVolumeDeviceName(diskID string) (string, error) {
	device, err := DefaultDeviceManager.GetDeviceByVolumeID(diskID)
	if err == nil {
		return device, nil
	}
	device = getVolumeConfig(diskID)
	if device != "" {
		klog.Infof("GetVolumeDeviceName: got disk %s device name %s by config file", diskID, device)
		return device, nil
	}
	// return error from GetDeviceByVolumeID if config file not found
	return device, err
}

// isPathAvailable
func isPathAvailable(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("Open Path (%s) with error: %v ", path, err)
	}
	defer f.Close()
	_, err = f.Readdirnames(1)
	if err != nil && err != io.EOF {
		return fmt.Errorf("Read Path (%s) with error: %v ", path, err)
	}
	return nil
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

func GetAvailableDiskTypes(ctx context.Context, c cloud.ECSInterface, m metadata.MetadataProvider) (types []string, err error) {
	request := ecs.CreateDescribeAvailableResourceRequest()
	request.InstanceType = metadata.MustGet(m, metadata.InstanceType)
	request.DestinationResource = describeResourceType
	request.ZoneId = metadata.MustGet(m, metadata.ZoneID)
	request.ResourceType = "disk"

	response, err := c.DescribeAvailableResource(request)
	if err != nil {
		return nil, fmt.Errorf("failed to DescribeAvailableResource for instance type %s: %v", request.InstanceType, err)
	}

	klog.V(4).Infof("UpdateNode: record ecs openapi req: %+v, resp: %+v", request, response)
	for _, zone := range response.AvailableZones.AvailableZone {
		if zone.ZoneId != request.ZoneId {
			continue
		}
		for _, resource := range zone.AvailableResources.AvailableResource {
			if resource.Type != describeResourceType {
				continue
			}
			for _, supportedResource := range resource.SupportedResources.SupportedResource {
				types = append(types, supportedResource.Value)
			}
		}
	}
	if len(types) == 0 {
		return nil, fmt.Errorf("no supported disk type found. response: %s", response.GetHttpContentString())
	}
	return types, nil
}

func hasDiskTypeLabel(node *v1.Node) bool {
	for k := range node.Labels {
		if strings.HasPrefix(k, nodeDiskTypeLabelPrefix) {
			return true
		}
	}
	return false
}

func patchForNode(node *v1.Node, maxVolumesNum int, diskTypes []string) []byte {
	maxVolumesNumStr := strconv.Itoa(maxVolumesNum)
	needUpdate := node.Annotations[nodeDiskCountAnnotation] != maxVolumesNumStr

	instanceStorageLabels := map[string]string{}
	for _, diskType := range diskTypes {
		labelKey := nodeDiskTypeLabelPrefix + diskType
		instanceStorageLabels[labelKey] = "available"
	}
	for l, v := range instanceStorageLabels {
		if node.Labels[l] != v {
			needUpdate = true
			break
		}
	}

	if !needUpdate {
		return nil
	}
	patch, err := json.Marshal(map[string]interface{}{
		"metadata": map[string]interface{}{
			"labels": instanceStorageLabels,
			"annotations": map[string]string{
				nodeDiskCountAnnotation: maxVolumesNumStr,
			},
		},
	})
	if err != nil {
		klog.Fatalf("failed to marshal patch json")
	}
	return patch
}

func volumeCreate(attempt createAttempt, diskID string, volSizeBytes int64, volumeContext map[string]string, zoneID string, contextSource *csi.VolumeContentSource) *csi.Volume {
	segments := map[string]string{}
	cateDesc := AllCategories[attempt.Category]
	volumeContext[labelAppendPrefix+TopologyRegionKey] = GlobalConfigVar.Region
	if cateDesc.Regional {
		segments[common.TopologyKeyRegion] = GlobalConfigVar.Region
	} else {
		segments[TopologyZoneKey] = zoneID
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
		labelKey := nodeDiskTypeLabelPrefix + string(attempt.Category)
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

	ecsClient := updateEcsClient(GlobalConfigVar.EcsClient)
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

func getAttachedCloudDisks(ecsClient cloud.ECSInterface, m metadata.MetadataProvider) (diskIds []string, err error) {
	req := ecs.CreateDescribeDisksRequest()
	req.InstanceId = metadata.MustGet(m, metadata.InstanceID)
	req.RegionId = metadata.MustGet(m, metadata.RegionID)
	req.MaxResults = requests.NewInteger(100)
	for {
		resp, err := ecsClient.DescribeDisks(req)
		if err != nil {
			return nil, fmt.Errorf("DescribeDisks failed: %w", err)
		}
		req.NextToken = resp.NextToken

		for _, disk := range resp.Disks.Disk {
			// Don't include detaching disks here
			// The device node for them in /dev may not present
			if disk.Status == "Detaching" {
				continue
			}
			// local disks have their own quota in LocalStorageAmount, and are not counted for DiskQuantity
			if strings.HasPrefix(disk.Category, "cloud") || strings.HasPrefix(disk.Category, "elastic_ephemeral_disk") {
				diskIds = append(diskIds, disk.DiskId)
			}
		}

		if len(req.NextToken) == 0 {
			break
		}
	}
	return diskIds, nil
}

func getAvailableDiskCount(node *v1.Node, ecsClient cloud.ECSInterface, m metadata.MetadataProvider) (int, error) {
	if count, err := getAvailableDiskCountFromAnnotation(node); err == nil {
		return count, nil
	}
	return getAvailableDiskCountFromOpenAPI(ecsClient, m)
}

func getAvailableDiskCountFromAnnotation(node *v1.Node) (int, error) {
	if node == nil || node.Annotations[instanceTypeInfoAnnotation] == "" {
		err := fmt.Errorf("empty instance type info annotation")
		klog.Info(err)
		return 0, err
	}
	var typeInfo InstanceTypeInfo
	if err := json.Unmarshal([]byte(node.Annotations[instanceTypeInfoAnnotation]), &typeInfo); err != nil {
		klog.Errorf("error unmarshalling instance type info annotation: %v", err)
		return 0, err
	}
	return typeInfo.DiskQuantity, nil
}

func getAvailableDiskCountFromOpenAPI(ecsClient cloud.ECSInterface, m metadata.MetadataProvider) (int, error) {
	req := ecs.CreateDescribeInstanceTypesRequest()
	req.RegionId = metadata.MustGet(m, metadata.RegionID)
	instanceType := metadata.MustGet(m, metadata.InstanceType)
	req.InstanceTypes = &[]string{instanceType}
	response, err := ecsClient.DescribeInstanceTypes(req)
	if err != nil {
		return 0, fmt.Errorf("DescribeInstanceTypes failed: %w", err)
	}
	for _, i := range response.InstanceTypes.InstanceType {
		if i.InstanceTypeId != instanceType {
			continue
		}
		return i.DiskQuantity, nil
	}
	return 0, fmt.Errorf("unexpected DescribeInstanceTypes response for %s: %s", instanceType, response.GetHttpContentString())
}

func getVolumeCountFromOpenAPI(getNode func() (*v1.Node, error), c cloud.ECSInterface, m metadata.MetadataProvider, dev utilsio.DiskLister) (int, error) {
	// An attached disk is not managed by us if:
	// 1. it is not in node.Status.VolumesInUse or node.Status.VolumesAttached; and
	// 2. it does not have the xattr set.
	// 1 may fail because the info in node.Status is removed before ControllerUnpublishVolume.
	// 2 may fail because the disk may be just attached and not have the xattr set yet.
	// Combine 1 and 2 to get the accurate "not managed" disk list.
	// We should exclude these disks from available count.
	// e.g. static/dynamic PVs are managed, OS disk or manually attached disks are not managed.

	managedDisks := sets.New[string]()

	diskPaths, err := dev.ListDisks()
	if err != nil {
		return 0, fmt.Errorf("failed to list devices: %w", err)
	}
	for _, p := range diskPaths {
		var diskID [32]byte
		sz, err := unix.Getxattr(p, DiskXattrName, diskID[:])
		if err == nil {
			// this disk has xattr, it is managed by us
			managedDisks.Insert(string(diskID[:sz]))
		} else if !utilsio.IsXattrNotFound(err) {
			klog.Warningf("getVolumeCount: failed to get xattr of %s, assuming not managed by us: %s", p, err)
		}
	}

	// To ensure all the managed attachedDisks also present in managedDisks,
	// ECS OpenAPI should goes after ListDisks because the just detached disk should
	// disappear from ListDisks after OpenAPI;
	// ECS OpenAPI should goes before getNode because the just attached disk should
	// appear in node before OpenAPI;
	attachedDisks, err := getAttachedCloudDisks(c, m)
	if err != nil {
		return 0, err
	}
	klog.Infof("getVolumeCount: found %d attached disks", len(attachedDisks))

	node, err := getNode()
	if err != nil {
		return 0, err
	}

	availableCount, err := getAvailableDiskCount(node, c, m)
	if err != nil {
		return 0, err
	}

	prefix := fmt.Sprintf("kubernetes.io/csi/%s^", driverName)
	getDiskId := func(n v1.UniqueVolumeName) string {
		if strings.HasPrefix(string(n), prefix) {
			return string(n[len(prefix):])
		}
		return ""
	}

	for _, volume := range node.Status.VolumesInUse {
		if disk := getDiskId(volume); disk != "" {
			managedDisks.Insert(disk)
		}
	}
	for _, volume := range node.Status.VolumesAttached {
		if disk := getDiskId(volume.Name); disk != "" {
			managedDisks.Insert(disk)
		}
	}

	for _, disk := range attachedDisks {
		if !managedDisks.Has(disk) {
			klog.Infof("getVolumeCount: disk %s is not managed by us", disk)
			availableCount--
		}
	}

	return availableCount, nil
}

// hasMountOption return boolean value indicating whether the slice contains a mount option
func hasMountOption(options []string, opt string) bool {
	for _, o := range options {
		if o == opt {
			return true
		}
	}
	return false
}

// checkRundVolumeExpand
func checkRundVolumeExpand(req *csi.NodeExpandVolumeRequest) (bool, error) {
	klog.Infof("checkRundVolumeExpand: volumePath: %s", req.VolumePath)
	pvName := utils.GetPvNameFormPodMnt(req.VolumePath)
	if pvName == "" {
		klog.Errorf("checkRundVolumeExpand: cannot get pvname from volumePath %s", req.VolumePath)
		return false, perrors.Errorf("cannot get pvname from volumePath %s for volume %s", req.VolumePath, req.VolumeId)
	}
	var grpcVolume string
	socketFile := filepath.Join(RundSocketDir, pvName)
	if utils.IsFileExisting(socketFile) {
		grpcVolume = pvName
	} else {
		socketFile = filepath.Join(RundSocketDir, req.VolumeId)
		if !utils.IsFileExisting(socketFile) {
			klog.Infof("checkRundVolumeExpand: socketfile: %s not exists, trying runc expanding", socketFile)
			return false, nil
		}
		grpcVolume = req.VolumeId
	}
	klog.Infof("checkRundVolumeExpand: rund socket dir: %s exists", socketFile)

	// connect to rund server with timeout
	clientConn, err := net.DialTimeout("unix", socketFile, 1*time.Second)
	if err != nil {
		klog.Errorf("checkRundExpand: volume %s, volumepath %s, connect to rund server with error: %s", req.VolumeId, req.VolumePath, err.Error())
		return true, perrors.Errorf("checkRundExpand: volume %s, volumepath %s, connect to rund server with error: %s", req.VolumeId, req.VolumePath, err.Error())
	}
	defer clientConn.Close()

	// send volume spec to rund to expand volume fs
	volumeSize := strconv.FormatInt(req.GetCapacityRange().GetRequiredBytes(), 10)
	client := proto.NewExtendedStatusClient(ttrpc.NewClient(clientConn))
	resp, err := client.ExpandVolume(context.Background(), &proto.ExpandVolumeRequest{
		Volume: grpcVolume,
	})
	if err != nil {
		klog.Errorf("checkRundExpand: volume %s, volumepath %s, connect to rund server with error response: %s", req.VolumeId, req.VolumePath, err.Error())
		return true, perrors.Errorf("checkRundExpand: volume %s, volumepath %s, connect to rund server with error response: %s", req.VolumeId, req.VolumePath, err.Error())
	}

	klog.Infof("RundVolumeExpand: Expand VolumeFS(%s) to(%s) successful with response: %s", pvName, volumeSize, resp)
	return true, nil
}

func checkOption(opt string) bool {
	switch opt {
	case "enable", "true", "yes":
		return true
	default:
		return false
	}
}

// getPvPvcFromDiskId returns a pv instance with specified disk ID
func getPvPvcFromDiskId(diskId string) (*v1.PersistentVolume, *v1.PersistentVolumeClaim, error) {
	ctx := context.Background()
	pv, err := GlobalConfigVar.ClientSet.CoreV1().PersistentVolumes().Get(ctx, diskId, metav1.GetOptions{})
	if err != nil {
		klog.Errorf("getPvcFromDiskId: failed to get pv from apiserver: %v", err)
		return nil, nil, err
	}
	pvcName, pvcNamespace := pv.Spec.ClaimRef.Name, pv.Spec.ClaimRef.Namespace
	pvc, err := GlobalConfigVar.ClientSet.CoreV1().PersistentVolumeClaims(pvcNamespace).Get(ctx, pvcName, metav1.GetOptions{})
	if err != nil {
		klog.Errorf("getPvcFromDiskId: failed to get pvc from apiserver: %v", err)
		return nil, nil, err
	}
	return pv, pvc, nil
}

// UpdatePvcWithAnnotations update pvc
func updatePvcWithAnnotations(ctx context.Context, pvc *v1.PersistentVolumeClaim, annotations map[string]string, option string) (*v1.PersistentVolumeClaim, error) {
	switch option {
	case "add":
		for key, value := range annotations {
			if pvc.Annotations == nil {
				pvc.Annotations = map[string]string{key: value}
			} else {
				pvc.Annotations[key] = value
			}
		}
	case "delete":
		if pvc.Annotations != nil {
			for key := range annotations {
				delete(pvc.Annotations, key)
			}
		}
	}
	return GlobalConfigVar.ClientSet.CoreV1().PersistentVolumeClaims(pvc.Namespace).Update(ctx, pvc, metav1.UpdateOptions{})
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
