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

package utils

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	aliyunep "github.com/aliyun/alibaba-cloud-sdk-go/sdk/endpoints"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/go-ping/ping"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/options"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	v1core "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/record"
	"k8s.io/klog/v2"
	k8smount "k8s.io/mount-utils"
	utilexec "k8s.io/utils/exec"
)

const (
	// MetadataURL is metadata url
	MetadataURL = "http://100.100.100.200/latest/meta-data/"
	// ECS worker role resource url in metadata server
	WorkerRoleResource = "ram/security-credentials/"
	// RegionIDTag is region id
	RegionIDTag = "region-id"
	// InstanceIDTag is instance id
	InstanceIDTag = "instance-id"
	// DefaultRegion is default region
	DefaultRegion = "cn-hangzhou"
	// CsiPluginRunTimeFlagFile tag
	CsiPluginRunTimeFlagFile = "../alibabacloudcsiplugin.json"
	// RuncRunTimeTag tag
	RuncRunTimeTag = "runc"
	// RunvRunTimeTag tag
	RunvRunTimeTag = "runv"
	// RunvRunTimeTag tag
	RundRunTimeTag = "rund"
	// PVMRunTimeTag tag
	PVMRunTimeTag = "sandbox-pvm"
	// ServiceType tag
	ServiceTypeEnv = "SERVICE_TYPE"
	// PluginService represents the csi-plugin type.
	PluginService = "plugin"
	// ProvisionerService represents the csi-provisioner type.
	ProvisionerService = "provisioner"
	// MetadataMaxRetryCount ...
	MetadataMaxRetryCount = 4
	// VolDataFileName file
	VolDataFileName = "vol_data.json"
	// fsckErrorsCorrected tag
	fsckErrorsCorrected = 1
	// fsckErrorsUncorrected tag
	fsckErrorsUncorrected = 4
	// socketPath is path of connector sock
	socketPath = "/host/run/csi-tool/connector/connector.sock"
	// hostPrefix is the prefix of host path
	hostPrefix = "/host"

	// GiB ...
	GiB = 1024 * 1024 * 1024

	PodNameKey      = "csi.storage.k8s.io/pod.name"
	PodNamespaceKey = "csi.storage.k8s.io/pod.namespace"
)

type ServiceType int

type podInfoKey struct{}

var podInfo = podInfoKey{}

const (
	Controller ServiceType = 1 << iota
	Node
)

func GetServiceType(runAsController, runControllerService, runNodeService bool) ServiceType {
	serviceType := ServiceType(0)
	if runAsController {
		klog.Warning("-run-as-controller is deprecated, use -run-node-service=false instead")
		serviceType = Controller
	}
	if st := os.Getenv(ServiceTypeEnv); st != "" {
		klog.Warningf("%s env support is deprecated, use -run-controller-service and -run-node-service instead", ServiceTypeEnv)
		switch st {
		case PluginService:
			if runAsController {
				klog.Fatalf("%s env is set to %s, but -run-as-controller is also set", ServiceTypeEnv, st)
			}
			serviceType = Node
		case ProvisionerService:
			serviceType = Controller
		default:
			klog.Fatalf("invalid %s env value: %s", ServiceTypeEnv, st)
		}
	}
	if serviceType == 0 {
		// nothing deprecated was set, use new flags
		if runControllerService {
			serviceType |= Controller
		}
		if runNodeService {
			serviceType |= Node
		}
	}
	if serviceType == 0 {
		klog.Warning("no service type activated, this configuration may not be useful")
	}
	if serviceType&Controller != 0 {
		klog.Infof("activate CSI controller service")
	}
	if serviceType&Node != 0 {
		klog.Infof("activate CSI node service")
	}
	return serviceType
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

// CreateEvent is create events
func CreateEvent(recorder record.EventRecorder, objectRef *v1.ObjectReference, eventType string, reason string, err string) {
	recorder.Event(objectRef, eventType, reason, err)
}

// NewEventRecorder is create snapshots event recorder
func NewEventRecorder() record.EventRecorder {
	broadcaster := record.NewBroadcaster()
	broadcaster.StartLogging(klog.Infof)
	cfg, err := options.GetRestConfig()
	if err != nil {
		klog.ErrorS(err, "Error building kubeconfig for events")
	} else {
		clientset, err := kubernetes.NewForConfig(cfg)
		if err != nil {
			klog.Fatalf("NewControllerServer: Failed to create client: %v", err)
		}
		sink := &v1core.EventSinkImpl{
			Interface: clientset.CoreV1().Events(""),
		}
		broadcaster.StartRecordingToSink(sink)
	}
	source := v1.EventSource{Component: "csi-controller-server"}
	return broadcaster.NewRecorder(scheme.Scheme, source)
}

var NsenterArgs = []string{"--target=1", "--mount", "--ipc", "--net", "--uts", "--"}

func CommandOnNode(args ...string) *exec.Cmd {
	allArgs := append(NsenterArgs, args...)
	return exec.Command("/usr/bin/nsenter", allArgs...)
}

// CreateDest create de destination dir
func CreateDest(dest string) error {
	fi, err := os.Lstat(dest)

	if os.IsNotExist(err) {
		if err := os.MkdirAll(dest, 0777); err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	if fi != nil && !fi.IsDir() {
		return fmt.Errorf("%v already exist but it's not a directory", dest)
	}
	return nil
}

// IsFileExisting check file exist in volume driver or not
func IsFileExisting(filename string) bool {
	_, err := os.Stat(filename)
	if err == nil {
		return true
	}
	return !os.IsNotExist(err)
}

func Gi2Bytes(gb int64) int64 {
	return gb * GiB
}

// BytesToGiB converts Bytes to GiB
func Bytes2GiB(volumeSizeBytes int64) int64 {
	return volumeSizeBytes / GiB
}

// RoundUpBytes rounds up the volume size in bytes upto multiplications of GiB
// in the unit of Bytes
func RoundUpBytes(volumeSizeBytes int64) int64 {
	return roundUpSize(volumeSizeBytes, GiB) * GiB
}

// TODO: check division by zero and int overflow
func roundUpSize(volumeSizeBytes int64, allocationUnitBytes int64) int64 {
	return (volumeSizeBytes + allocationUnitBytes - 1) / allocationUnitBytes
}

// ReadJSONFile return a json object
func ReadJSONFile(file string) (map[string]string, error) {
	jsonObj := map[string]string{}
	raw, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(raw, &jsonObj)
	if err != nil {
		return nil, err
	}
	return jsonObj, nil
}

// NewEcsClient create a ecsClient object
func NewEcsClient(ac AccessControl) (ecsClient *ecs.Client) {
	if ep := os.Getenv("ECS_ENDPOINT"); ep != "" {
		_ = aliyunep.AddEndpointMapping(DefaultRegion, "Ecs", ep)
	}
	var err error
	switch ac.UseMode {
	case AccessKey:
		ecsClient, err = ecs.NewClientWithAccessKey(DefaultRegion, ac.AccessKeyID, ac.AccessKeySecret)
	case Credential:
		ecsClient, err = ecs.NewClientWithOptions(DefaultRegion, ac.Config, ac.Credential)
	default:
		ecsClient, err = ecs.NewClientWithStsToken(DefaultRegion, ac.AccessKeyID, ac.AccessKeySecret, ac.StsToken)

	}

	if err != nil {
		return nil
	}
	return
}

func getOpenAPIConfig(regionID string) *openapi.Config {
	config := &openapi.Config{RegionId: &regionID}
	if e := os.Getenv("ALICLOUD_CLIENT_SCHEME"); e != "" {
		config.Protocol = &e
	}
	if e := os.Getenv("ALIBABA_CLOUD_NETWORK_TYPE"); e != "" {
		config.Network = &e
	}
	return config
}

func GetStsConfig(regionID string) *openapi.Config {
	config := getOpenAPIConfig(regionID)
	if e := os.Getenv("STS_ENDPOINT"); e != "" {
		config.Endpoint = &e
	}
	return config
}

func GetEcsConfig(regionID string) *openapi.Config {
	config := getOpenAPIConfig(regionID)
	if e := os.Getenv("ECS_ENDPOINT"); e != "" {
		config.Endpoint = &e
	}
	return config
}

// IsDir check file is directory
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// GetFileContent get file content
func GetFileContent(fileName string) string {
	volumeFile := path.Join(fileName)
	if !IsFileExisting(volumeFile) {
		return ""
	}
	value, err := os.ReadFile(volumeFile)
	if err != nil {
		return ""
	}
	devicePath := strings.TrimSpace(string(value))
	return devicePath
}

// GetAccessModes returns a slice containing all of the access modes defined
// in the passed in VolumeCapabilities.
func GetAccessModes(caps []*csi.VolumeCapability) *[]string {
	modes := []string{}
	for _, c := range caps {
		modes = append(modes, c.AccessMode.GetMode().String())
	}
	return &modes
}

// WriteJSONFile save json data to file
func WriteJSONFile(obj interface{}, file string) error {
	rankingsJSON, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	if err := os.WriteFile(file, rankingsJSON, 0644); err != nil {
		return err
	}
	return nil
}

// GetPodRunTime Get Pod runtimeclass config
// Default as runc.
func GetPodRunTime(ctx context.Context, req *csi.NodePublishVolumeRequest, clientSet kubernetes.Interface) string {
	if req.VolumeContext["csi.alibabacloud.com/rootfs-volume"] == "true" {
		// For rootfs volume, always mount directly. Even if it is used by runD, mount it like runC.
		return RuncRunTimeTag
	}
	// if pod name namespace is empty, use default
	runtimeVal := RuncRunTimeTag
	pod, err := GetPodFromContextOrK8s(ctx, clientSet, req)
	if err != nil {
		klog.Errorf("GetPodRunTime: failed to get pod: %v", err)
		return runtimeVal
	}
	namespace, name := pod.Namespace, pod.Name
	// check pod.Spec.RuntimeClassName == "runv"
	if pod.Spec.RuntimeClassName == nil || *pod.Spec.RuntimeClassName == "" {
		klog.Infof("GetPodRunTime: Get default runtime(nil), %s, %s", name, namespace)
		return runtimeVal
	} else {
		runtimeClassName := strings.TrimSpace(*pod.Spec.RuntimeClassName)
		klog.Infof("GetPodRunTime: Get PodInfo Successful: %s, %s, with runtime: %s", name, namespace, runtimeClassName)
		if runtimeClassName == RunvRunTimeTag || runtimeClassName == RundRunTimeTag {
			runtimeVal = runtimeClassName
		}
	}
	// check Annotation[io.kubernetes.cri.untrusted-workload] = true
	if value, ok := pod.Annotations["io.kubernetes.cri.untrusted-workload"]; ok && strings.TrimSpace(value) == "true" {
		klog.Infof("GetPodRunTime: namespace: %s, name: %s pod has untrusted workload tag, use runv runtime logic", namespace, name)
		runtimeVal = RunvRunTimeTag
	}
	return runtimeVal
}

func WithPodInfo(ctx context.Context, client kubernetes.Interface, req *csi.NodePublishVolumeRequest) (context.Context, *v1.Pod) {
	pod, err := getPodFromK8s(ctx, client, req)
	if err != nil {
		klog.Errorf("WithPodInfo: failed to get pod: %v", err)
		return ctx, nil
	}
	return context.WithValue(ctx, podInfo, pod), pod
}

func GetPodFromContextOrK8s(ctx context.Context, client kubernetes.Interface, req *csi.NodePublishVolumeRequest) (*v1.Pod, error) {
	pod, ok := ctx.Value(podInfo).(*v1.Pod)
	if ok {
		return pod, nil
	}
	return getPodFromK8s(ctx, client, req)
}

func getPodFromK8s(ctx context.Context, client kubernetes.Interface, req *csi.NodePublishVolumeRequest) (*v1.Pod, error) {
	name, namespace := req.VolumeContext[PodNameKey], req.VolumeContext[PodNamespaceKey]
	if name == "" || namespace == "" {
		return nil, fmt.Errorf("empty pod name or namespace: '%s', '%s'", name, namespace)
	}
	if client == nil {
		return nil, fmt.Errorf("kube client is nil")
	}
	pod, err := client.CoreV1().Pods(namespace).Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return pod, nil
}

// IsMountPointRunv check the mountpoint is runv style.
// Remember to check this is not a regular mountpoint before calling this function.
func IsMountPointRunv(mountPoint string) bool {
	mountFileName := filepath.Join(mountPoint, CsiPluginRunTimeFlagFile)
	if IsFileExisting(mountFileName) {
		mountInfo := GetFileContent(mountFileName)
		mountInfo = strings.ToLower(mountInfo)
		maps := map[string]string{}
		if err := json.Unmarshal([]byte(mountInfo), &maps); err != nil {
			return false
		}
		if value, ok := maps["mountfile"]; ok && value == mountFileName {
			if valuert, okrt := maps["runtime"]; okrt && valuert == "runv" {
				return true
			}
		}
	}
	return false
}

// Ping check network like shell ping command
func Ping(ipAddress string) (*ping.Statistics, error) {
	pinger, err := ping.NewPinger(ipAddress)
	if err != nil {
		return nil, err
	}
	pinger.SetPrivileged(true)
	pinger.Count = 1
	pinger.Timeout = time.Second * 2
	err = pinger.Run()
	if err != nil {
		return nil, err
	}
	stats := pinger.Statistics()
	return stats, nil
}

// IsDirTmpfs check path is tmpfs mounted or not
func IsDirTmpfs(mounter k8smount.Interface, path string) (bool, error) {
	mnts, err := mounter.List()
	if err != nil {
		return false, err
	}
	for _, mnt := range mnts {
		// in case of multiple mounts stack on the same path, we only check the first one
		if mnt.Path == path {
			return mnt.Type == "tmpfs", nil
		}
	}
	return false, nil
}

// WriteAndSyncFile behaves just like ioutil.WriteFile in the standard library,
// but calls Sync before closing the file. WriteAndSyncFile guarantees the data
// is synced if there is no error returned.
func WriteAndSyncFile(filename string, data []byte, perm os.FileMode) error {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, perm)
	if err != nil {
		return err
	}
	n, err := f.Write(data)
	if err == nil && n < len(data) {
		err = io.ErrShortWrite
	}
	if err == nil {
		err = Fsync(f)
	}
	if err1 := f.Close(); err == nil {
		err = err1
	}
	return err
}

// Fsync is a wrapper around file.Sync(). Special handling is needed on darwin platform.
func Fsync(f *os.File) error {
	return f.Sync()
}

// GetPvNameFormPodMnt get pv name
func GetPvNameFormPodMnt(mntPath string) string {
	if mntPath == "" {
		return ""
	}
	if strings.HasSuffix(mntPath, "/mount") {
		tmpPath := mntPath[0 : len(mntPath)-6]
		pvName := filepath.Base(tmpPath)
		return pvName
	}
	return ""
}

// ConnectorRun Run shell command with host connector
// host connector is daemon running in host.
func ConnectorRun(cmd ...string) (string, error) {
	c, err := net.Dial("unix", socketPath)
	if err != nil {
		klog.Errorf("Oss connector Dial error: %s", err.Error())
		return err.Error(), err
	}
	defer c.Close()

	_, err = c.Write([]byte(strings.Join(cmd, "\x00")))
	if err != nil {
		klog.Errorf("Oss connector write error: %s", err.Error())
		return err.Error(), err
	}

	buf := make([]byte, 2048)
	n, _ := c.Read(buf[:])
	response := string(buf[0:n])
	if strings.HasPrefix(response, "Success") {
		respstr := response[8:]
		return respstr, nil
	}
	return response, errors.New("Exec command error:" + response)
}

// AppendJSONData append map data to json file.
func AppendJSONData(dataFilePath string, appData map[string]string) error {
	curData, err := LoadJSONData(dataFilePath)
	if err != nil {
		return err
	}
	for key, value := range appData {
		if strings.HasPrefix(key, "csi.alibabacloud.com/") {
			curData[key] = value
		}
	}
	rankingsJSON, _ := json.Marshal(curData)
	if err := os.WriteFile(dataFilePath, rankingsJSON, 0644); err != nil {
		return err
	}

	klog.Infof("AppendJSONData: Json data file saved successfully [%s], content: %v", dataFilePath, curData)
	return nil
}

// LoadJSONData loads json info from specified json file
func LoadJSONData(dataFileName string) (map[string]string, error) {
	file, err := os.Open(dataFileName)
	if err != nil {
		return nil, fmt.Errorf("failed to open json data file [%s]: %v", dataFileName, err)
	}
	defer file.Close()
	data := map[string]string{}
	if err := json.NewDecoder(file).Decode(&data); err != nil {
		return nil, fmt.Errorf("failed to parse json data file [%s]: %v", dataFileName, err)
	}
	return data, nil
}

// IsPathAvailable
func IsPathAvailable(path string) error {
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

// formatAndMount uses unix utils to format and mount the given disk
func FormatAndMount(diskMounter *k8smount.SafeFormatAndMount, source string, target string, fstype string, mkfsOptions []string, mountOptions []string, omitFsCheck bool) error {
	klog.Infof("formatAndMount: mount options : %+v", mountOptions)
	readOnly := false
	for _, option := range mountOptions {
		if option == "ro" {
			readOnly = true
			break
		}
	}

	// check device fs
	mountOptions = append(mountOptions, "defaults")
	if !readOnly || !omitFsCheck {
		// Run fsck on the disk to fix repairable issues, only do this for volumes requested as rw.
		args := []string{"-a", source}

		out, err := diskMounter.Exec.Command("fsck", args...).CombinedOutput()
		if err != nil {
			ee, isExitError := err.(utilexec.ExitError)
			switch {
			case err == utilexec.ErrExecutableNotFound:
				klog.Warningf("'fsck' not found on system; continuing mount without running 'fsck'.")
			case isExitError && ee.ExitStatus() == fsckErrorsCorrected:
				klog.Infof("Device %s has errors which were corrected by fsck.", source)
			case isExitError && ee.ExitStatus() == fsckErrorsUncorrected:
				return fmt.Errorf("'fsck' found errors on device %s but could not correct them: %s", source, string(out))
			case isExitError && ee.ExitStatus() > fsckErrorsUncorrected:
			}
		}
	}

	// Try to mount the disk
	mountErr := diskMounter.Interface.Mount(source, target, fstype, mountOptions)
	if mountErr != nil {
		// Mount failed. This indicates either that the disk is unformatted or
		// it contains an unexpected filesystem.
		existingFormat, err := diskMounter.GetDiskFormat(source)
		if err != nil {
			return err
		}
		if existingFormat == "" {
			return FormatNewDisk(readOnly, source, fstype, target, mkfsOptions, mountOptions, diskMounter)
		}
		// Disk is already formatted and failed to mount
		if len(fstype) == 0 || fstype == existingFormat {
			// This is mount error
			return mountErr
		}
		// Detect if an encrypted disk is empty disk, since atari type partition is detected by blkid.
		if existingFormat == "unknown data, probably partitions" {
			klog.Infof("FormatAndMount: enter special partition logics")
			fsType, ptType, _ := GetDiskFStypePTtype(source)
			if fsType == "" && ptType == "atari" {
				return FormatNewDisk(readOnly, source, fstype, target, mkfsOptions, mountOptions, diskMounter)
			}
		}
		// Block device is formatted with unexpected filesystem, let the user know
		return fmt.Errorf("failed to mount the volume as %q, it already contains %s. Mount error: %v", fstype, existingFormat, mountErr)
	}

	return mountErr
}

func FormatNewDisk(readOnly bool, source, fstype, target string, mkfsOptions, mountOptions []string, diskMounter *k8smount.SafeFormatAndMount) error {
	if readOnly {
		// Don't attempt to format if mounting as readonly, return an error to reflect this.
		return errors.New("failed to mount unformatted volume as read only")
	}

	// Use 'ext4' as the default
	if len(fstype) == 0 {
		fstype = "ext4"
	}

	args := mkfsDefaultArgs(fstype, source)

	// add mkfs options
	if len(mkfsOptions) != 0 {
		args = []string{}
		args = append(args, mkfsOptions...)
		args = append(args, source)
	}

	klog.Infof("Disk %q appears to be unformatted, attempting to format as type: %q with options: %v", source, fstype, args)
	startT := time.Now()

	pvName := filepath.Base(source)
	_, err := diskMounter.Exec.Command("mkfs."+fstype, args...).CombinedOutput()
	if err == nil {
		// the disk has been formatted successfully try to mount it again.
		klog.Infof("Disk format succeeded, pvName: %s elapsedTime: %+v ms", pvName, time.Since(startT).Milliseconds())
		return diskMounter.Interface.Mount(source, target, fstype, mountOptions)
	}
	klog.Errorf("format of disk %q failed: type:(%q) target:(%q) options:(%q) error:(%v)", source, fstype, target, args, err)
	return err
}

func mkfsDefaultArgs(fstype, source string) (args []string) {
	// default args
	if fstype == "ext4" || fstype == "ext3" {
		args = []string{
			"-F",  // Force flag
			"-m0", // Zero blocks reserved for super-user
			source,
		}
	} else if fstype == "xfs" {
		args = []string{
			"-f",
			source,
		}
	}
	return
}

// GetDiskFStypePTtype uses 'blkid' to see if the given disk is unformatted
func GetDiskFStypePTtype(disk string) (fstype string, pttype string, err error) {
	args := []string{"-p", "-s", "TYPE", "-s", "PTTYPE", "-o", "export", disk}

	dataOut, err := utilexec.New().Command("blkid", args...).CombinedOutput()
	output := string(dataOut)

	if err != nil {
		if exit, ok := err.(utilexec.ExitError); ok {
			if exit.ExitStatus() == 2 {
				// Disk device is unformatted.
				// For `blkid`, if the specified token (TYPE/PTTYPE, etc) was
				// not found, or no (specified) devices could be identified, an
				// exit code of 2 is returned.
				return "", "", nil
			}
		}
		klog.Errorf("Could not determine if disk %q is formatted (%v)", disk, err)
		return "", "", err
	}

	lines := strings.Split(output, "\n")
	for _, l := range lines {
		if len(l) <= 0 {
			// Ignore empty line.
			continue
		}
		cs := strings.Split(l, "=")
		if len(cs) != 2 {
			return "", "", fmt.Errorf("blkid returns invalid output: %s", output)
		}
		// TYPE is filesystem type, and PTTYPE is partition table type, according
		// to https://www.kernel.org/pub/linux/utils/util-linux/v2.21/libblkid-docs/.
		if cs[0] == "TYPE" {
			fstype = cs[1]
		} else if cs[0] == "PTTYPE" {
			pttype = cs[1]
		}
	}

	if len(pttype) > 0 {
		// Returns a special non-empty string as filesystem type, then kubelet
		// will not format it.
		return fstype, pttype, nil
	}

	return fstype, "", nil
}

func HasSpecificTagKey(tagKey string, disk *ecs.Disk) (bool, string) {
	exists := false
	for _, tag := range disk.Tags.Tag {
		if tag.TagKey == tagKey {
			exists = true
			return exists, tag.TagValue
		}
	}
	return exists, ""
}

func IsPrivateCloud() bool {
	privateTag := os.Getenv("PRIVATE_CLOUD_TAG")
	privateBool, err := strconv.ParseBool(privateTag)
	if err != nil {
		return false
	}
	return privateBool
}

func parseUdevadmInfoSerial(out string) string {
	for _, line := range strings.Split(string(out), "\n") {
		const prefix = "ID_SERIAL_SHORT="
		if strings.HasPrefix(line, prefix) {
			return line[len(prefix):]
		}
	}
	return ""
}

// Get NVME device name by diskID;
// /dev/nvme0n1 0: means device index, 1: means namespace for nvme device;
// udevadm info --query=all --name=/dev/nvme0n1 | grep ID_SERIAL_SHORT | awk -F= '{print $2}'
// bp1bcfmvsobfauvxb3ow
func GetNvmeDeviceByVolumeID(volumeID string) (device string, err error) {
	serialNumber := strings.TrimPrefix(volumeID, "d-")
	devNumFiles, err := filepath.Glob("/sys/block/*/dev")
	if err != nil {
		klog.Infof("GetNvmeDeviceByVolumeID: List device number failed: %v", err)
		return "", err
	}
	for _, deviceFile := range devNumFiles {
		deviceName := filepath.Base(filepath.Dir(deviceFile))
		if strings.HasPrefix(deviceName, "nvme") && !strings.Contains(deviceName, "p") {
			out, err := CommandOnNode("udevadm", "info", "--query=property", "--name=/dev/"+deviceName).CombinedOutput()
			if err != nil {
				klog.Warningf("GetNvmeDeviceByVolumeID: udevadm failed: %s (%v)", string(out), err)
				continue
			}
			snumber := parseUdevadmInfoSerial(string(out))
			if snumber == "" {
				klog.Warningf("GetNvmeDeviceByVolumeID: udevadm did not know serial number for %s", deviceName)
			} else if serialNumber == snumber {
				device = filepath.Join("/dev/", deviceName)
				klog.Infof("GetNvmeDeviceByVolumeID: Get nvme device %s with volumeID %s", device, volumeID)
				return device, nil
			}
		}
	}
	return "", nil
}
