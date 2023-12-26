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
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/go-ping/ping"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/options"
	log "github.com/sirupsen/logrus"
	"golang.org/x/sys/unix"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	v1core "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/record"
	k8smount "k8s.io/mount-utils"
	utilexec "k8s.io/utils/exec"
	"k8s.io/utils/mount"
)

// DefaultOptions used for global ak
type DefaultOptions struct {
	Global struct {
		KubernetesClusterTag string
		AccessKeyID          string `json:"accessKeyID"`
		AccessKeySecret      string `json:"accessKeySecret"`
		Region               string `json:"region"`
	}
}

const (
	// UserAKID is user AK ID
	UserAKID = "/etc/.volumeak/akId"
	// UserAKSecret is user AK Secret
	UserAKSecret = "/etc/.volumeak/akSecret"
	// MetadataURL is metadata url
	MetadataURL = "http://100.100.100.200/latest/meta-data/"
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
	// ServiceType tag
	ServiceType = "SERVICE_TYPE"
	// PluginService represents the csi-plugin type.
	PluginService = "plugin"
	// ProvisionerService represents the csi-provisioner type.
	ProvisionerService = "provisioner"
	// InstallSnapshotCRD tag
	InstallSnapshotCRD = "INSTALL_SNAPSHOT_CRD"
	// MetadataMaxRetryCount ...
	MetadataMaxRetryCount = 4
	// VolDataFileName file
	VolDataFileName = "vol_data.json"
	// fsckErrorsCorrected tag
	fsckErrorsCorrected = 1
	// fsckErrorsUncorrected tag
	fsckErrorsUncorrected = 4

	// NsenterCmd is the nsenter command
	NsenterCmd = "nsenter --mount=/proc/1/ns/mnt --ipc=/proc/1/ns/ipc --net=/proc/1/ns/net --uts=/proc/1/ns/uts"

	// socketPath is path of connector sock
	socketPath = "/host/etc/csi-tool/connector.sock"

	// GiB ...
	GiB = 1024 * 1024 * 1024
)

// KubernetesAlicloudIdentity set a identity label
var KubernetesAlicloudIdentity = "Kubernetes.Alicloud/CsiPlugin"

var (
	nodeAddrMap = sync.Map{}
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

// CreateEvent is create events
func CreateEvent(recorder record.EventRecorder, objectRef *v1.ObjectReference, eventType string, reason string, err string) {
	recorder.Event(objectRef, eventType, reason, err)
}

// NewEventRecorder is create snapshots event recorder
func NewEventRecorder() record.EventRecorder {
	cfg, err := clientcmd.BuildConfigFromFlags(options.MasterURL, options.Kubeconfig)
	if err != nil {
		log.Fatalf("Error building kubeconfig: %s", err.Error())
	}
	clientset, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		log.Fatalf("NewControllerServer: Failed to create client: %v", err)
	}
	broadcaster := record.NewBroadcaster()
	broadcaster.StartLogging(log.Infof)
	source := v1.EventSource{Component: "csi-controller-server"}
	if broadcaster != nil {
		sink := &v1core.EventSinkImpl{
			Interface: v1core.New(clientset.CoreV1().RESTClient()).Events(""),
		}
		broadcaster.StartRecordingToSink(sink)
	}
	return broadcaster.NewRecorder(scheme.Scheme, source)
}

// Succeed return a Succeed Result
func Succeed(a ...interface{}) Result {
	return Result{
		Status:  "Success",
		Message: fmt.Sprint(a...),
	}
}

// NotSupport return a NotSupport Result
func NotSupport(a ...interface{}) Result {
	return Result{
		Status:  "Not supported",
		Message: fmt.Sprint(a...),
	}
}

// Fail return a Fail Result
func Fail(a ...interface{}) Result {
	return Result{
		Status:  "Failure",
		Message: fmt.Sprint(a...),
	}
}

// Result struct definition
type Result struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
	Device  string `json:"device,omitempty"`
}

// CommandRunFunc define the run function in utils for ut
type CommandRunFunc func(cmd string) (string, error)

// Run command
func ValidateRun(cmd string) (string, error) {
	arr := strings.Split(cmd, " ")
	withArgs := false
	if len(arr) >= 2 {
		withArgs = true
	}

	name := arr[0]
	var args []string
	err := CheckCmd(cmd, name)
	if err != nil {
		return "", err
	}
	if withArgs {
		args = arr[1:]
		err = CheckCmdArgs(cmd, args...)
		if err != nil {
			return "", err
		}
	}
	safeMount := mount.SafeFormatAndMount{
		Interface: mount.New(""),
		Exec:      utilexec.New(),
	}
	var command utilexec.Cmd
	if withArgs {
		command = safeMount.Exec.Command(name, args...)
	} else {
		command = safeMount.Exec.Command(name)
	}

	stdout, err := command.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("Failed to exec command:%s, name:%s, args:%+v, stdout:%s, stderr:%s", cmd, name, args, string(stdout), err.Error())
	}
	log.Infof("Exec command %s is successfully, name:%s, args:%+v", cmd, name, args)
	return string(stdout), nil
}

var NsenterArgs = []string{"--target=1", "--mount", "--ipc", "--net", "--uts", "--"}

func CommandOnNode(args ...string) *exec.Cmd {
	allArgs := append(NsenterArgs, args...)
	return exec.Command("/nsenter", allArgs...)
}

// run shell command
func Run(cmd string) (string, error) {
	out, err := exec.Command("sh", "-c", cmd).CombinedOutput()
	strOut := string(out)
	if err != nil {
		err = fmt.Errorf("Failed to run cmd: %s, with out: %s: %w", cmd, strOut, err)
	}
	return strOut, err
}

func RunWithFilter(cmd string, filter ...string) ([]string, error) {
	ans := make([]string, 0)
	stdout, err := Run(cmd)
	if err != nil {
		return nil, err
	}
	stdoutArr := strings.Split(string(stdout), "\n")
	for _, stdout := range stdoutArr {
		find := true
		for _, f := range filter {
			if !strings.Contains(stdout, f) {
				find = false
			}
		}
		if find {
			ans = append(ans, stdout)
		}
	}
	return ans, nil
}

// RunTimeout tag
func RunTimeout(cmd string, timeout int) error {
	ctx := context.Background()
	if timeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
		defer cancel()
	}

	cmdCont := exec.CommandContext(ctx, "sh", "-c", cmd)
	if err := cmdCont.Run(); err != nil {
		return err
	}
	return nil
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

// IsLikelyNotMountPoint return status of mount point,this function fix IsMounted return 0 bug
// IsLikelyNotMountPoint determines if a directory is not a mountpoint.
// It is fast but not necessarily ALWAYS correct. If the path is in fact
// a bind mount from one part of a mount to another it will not be detected.
// It also can not distinguish between mountpoints and symbolic links.
// mkdir /tmp/a /tmp/b; mount --bind /tmp/a /tmp/b; IsLikelyNotMountPoint("/tmp/b")
// will return true. When in fact /tmp/b is a mount point. If this situation
// is of interest to you, don't use this function...
func IsLikelyNotMountPoint(file string) (bool, error) {
	stat, err := os.Stat(file)
	if err != nil {
		return true, err
	}
	rootStat, err := os.Stat(filepath.Dir(strings.TrimSuffix(file, "/")))
	if err != nil {
		return true, err
	}
	// If the directory has a different device as parent, then it is a mountpoint.
	if stat.Sys().(*syscall.Stat_t).Dev != rootStat.Sys().(*syscall.Stat_t).Dev {
		return false, nil
	}

	return true, nil
}

// IsMounted return status of mount operation
func IsMounted(mountPath string) bool {
	stdout, err := RunWithFilter("mount", mountPath)
	if err != nil {
		log.Infof("IsMounted: Exec command mount is failed, err: %s, %s", stdout, err.Error())
		return false
	}
	if len(stdout) == 0 {
		return false
	}
	return true
}

// IsMountedInHost return status of host mounted or not
func IsMountedInHost(mountPath string) bool {
	cmd := fmt.Sprintf("%s mount", NsenterCmd)
	stdout, err := RunWithFilter(cmd, mountPath)
	if err != nil {
		log.Infof("IsMounted: Exec command %s is failed, err: %s", cmd, err.Error())
		return false
	}
	if len(stdout) == 0 {
		return false
	}
	return true
}

// UmountInHost do an unmount operation
func UmountInHost(mountPath string) error {
	cmd := fmt.Sprintf("%s umount %s", NsenterCmd, mountPath)
	_, err := Run(cmd)
	if err != nil {
		return err
	}
	return nil
}

// Umount do an unmount operation
func Umount(mountPath string) error {
	cmd := fmt.Sprintf("umount %s", mountPath)
	_, err := Run(cmd)
	if err != nil {
		return err
	}
	return nil
}

// IsFileExisting check file exist in volume driver or not
func IsFileExisting(filename string) bool {
	_, err := os.Stat(filename)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}

// GetRegionAndInstanceID get region and instanceID object
func GetRegionAndInstanceID() (string, string, error) {
	regionID, err := GetMetaData(RegionIDTag)
	if err != nil {
		return "", "", err
	}
	instanceID, err := GetMetaData(InstanceIDTag)
	if err != nil {
		return "", "", err
	}
	return regionID, instanceID, nil
}

// GetRegionID Get RegionID from Environment Variables or Metadata
func GetRegionID() (string, error) {
	var err error
	regionID := os.Getenv("REGION_ID")
	if regionID == "" {
		regionID, err = GetMetaData(RegionIDTag)
	}
	return regionID, err
}

// GetMetaData get metadata from ecs meta-server
func GetMetaData(resource string) (string, error) {
	resp, err := http.Get(MetadataURL + resource)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != 200 {
		msg := fmt.Sprintf("GetMetaData Response StatusCode %d, Response: %++v", resp.StatusCode, resp)
		return "", errors.New(msg)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if strings.Contains(string(body), "Error 500 Internal Server Error") {
		msg := fmt.Sprintf("GetMetaData Response StatusCode %d, Response: %++v", resp.StatusCode, resp)
		return "", errors.New(msg)
	}
	return string(body), nil
}

// RetryGetMetaData ...
func RetryGetMetaData(resource string) string {
	var response string
	for i := 0; i < MetadataMaxRetryCount; i++ {
		response, _ = GetMetaData(resource)
		if response != "" {
			break
		}
		time.Sleep(1 * time.Second)
	}
	if response == "" {
		log.Fatalf("RetryGetMetadata: failed to get metadata %s%s after %d retries", MetadataURL, resource, MetadataMaxRetryCount)
	}
	log.Infof("RetryGetMetaData: successful get metadata %v: %v", resource, response)
	return response
}

// GetRegionIDAndInstanceID get regionID and instanceID object
func GetRegionIDAndInstanceID(nodeName string) (string, string, error) {
	strs := strings.SplitN(nodeName, ".", 2)
	if len(strs) < 2 {
		return "", "", fmt.Errorf("failed to get regionID and instanceId from nodeName")
	}
	return strs[0], strs[1], nil
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

// RoundUpGiB rounds up the volume size in bytes upto multiplications of GiB
// in the unit of GiB
func RoundUpGiB(volumeSizeBytes int64) int64 {
	return roundUpSize(volumeSizeBytes, GiB)
}

// BytesToGiB converts Bytes to GiB
func BytesToGiB(volumeSizeBytes int64) int64 {
	return volumeSizeBytes / GiB
}

// TODO: check division by zero and int overflow
func roundUpSize(volumeSizeBytes int64, allocationUnitBytes int64) int64 {
	return (volumeSizeBytes + allocationUnitBytes - 1) / allocationUnitBytes
}

func KBlock2Bytes(kblocks int64) int64 {
	return kblocks * 1024
}

// ReadJSONFile return a json object
func ReadJSONFile(file string) (map[string]string, error) {
	jsonObj := map[string]string{}
	raw, err := ioutil.ReadFile(file)
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

// IsDir check file is directory
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// GetMetrics get path metric
func GetMetrics(path string) (*csi.NodeGetVolumeStatsResponse, error) {
	if path == "" {
		return nil, fmt.Errorf("getMetrics No path given")
	}
	statfs := &unix.Statfs_t{}
	err := unix.Statfs(path, statfs)
	if err != nil {
		return nil, err
	}

	// Available is blocks available * fragment size
	available := int64(statfs.Bavail) * int64(statfs.Bsize)

	// Capacity is total block count * fragment size
	capacity := int64(statfs.Blocks) * int64(statfs.Bsize)

	// Usage is block being used * fragment size (aka block size).
	usage := (int64(statfs.Blocks) - int64(statfs.Bfree)) * int64(statfs.Bsize)

	inodes := int64(statfs.Files)
	inodesFree := int64(statfs.Ffree)
	inodesUsed := inodes - inodesFree

	return &csi.NodeGetVolumeStatsResponse{
		Usage: []*csi.VolumeUsage{
			{
				Available: available,
				Total:     capacity,
				Used:      usage,
				Unit:      csi.VolumeUsage_BYTES,
			},
			{
				Available: inodesFree,
				Total:     inodes,
				Used:      inodesUsed,
				Unit:      csi.VolumeUsage_INODES,
			},
		},
	}, nil
}

// GetFileContent get file content
func GetFileContent(fileName string) string {
	volumeFile := path.Join(fileName)
	if !IsFileExisting(volumeFile) {
		return ""
	}
	value, err := ioutil.ReadFile(volumeFile)
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
	maps := make(map[string]interface{})
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).String() != "" {
			maps[t.Field(i).Name] = v.Field(i).String()
		}
	}
	rankingsJSON, _ := json.Marshal(maps)
	if err := ioutil.WriteFile(file, rankingsJSON, 0644); err != nil {
		return err
	}
	return nil
}

// GetPodRunTime Get Pod runtimeclass config
// Default as runc.
func GetPodRunTime(req *csi.NodePublishVolumeRequest, clientSet *kubernetes.Clientset) (string, error) {
	// if pod name namespace is empty, use default
	podName, nameSpace := "", ""
	if value, ok := req.VolumeContext["csi.storage.k8s.io/pod.name"]; ok {
		podName = value
	}
	if value, ok := req.VolumeContext["csi.storage.k8s.io/pod.namespace"]; ok {
		nameSpace = value
	}
	if podName == "" || nameSpace == "" {
		log.Warnf("GetPodRunTime: Rreceive Request with Empty name or namespace: %s, %s", podName, nameSpace)
		return "", fmt.Errorf("GetPodRunTime: Rreceive Request with Empty name or namespace")
	}

	podInfo, err := clientSet.CoreV1().Pods(nameSpace).Get(context.Background(), podName, metav1.GetOptions{})
	if err != nil {
		log.Errorf("GetPodRunTime: Get PodInfo(%s, %s) with error: %s", podName, nameSpace, err.Error())
		return "", fmt.Errorf("GetPodRunTime: Get PodInfo(%s, %s) with error: %s", podName, nameSpace, err.Error())
	}
	runTimeValue := RuncRunTimeTag

	// check pod.Spec.RuntimeClassName == "runv"
	if podInfo.Spec.RuntimeClassName == nil {
		log.Infof("GetPodRunTime: Get without runtime(nil), %s, %s", podName, nameSpace)
	} else if *podInfo.Spec.RuntimeClassName == "" {
		log.Infof("GetPodRunTime: Get with empty runtime: %s, %s", podName, nameSpace)
	} else {
		log.Infof("GetPodRunTime: Get PodInfo Successful: %s, %s, with runtime: %s", podName, nameSpace, *podInfo.Spec.RuntimeClassName)
		if strings.TrimSpace(*podInfo.Spec.RuntimeClassName) == RunvRunTimeTag {
			runTimeValue = RunvRunTimeTag
		}
	}

	// Deprecated pouch为了支持k8s 1.12以前没有RuntimeClass的情况做的特殊逻辑，为了代码健壮性，这里做下支持
	if podInfo.Annotations["io.kubernetes.runtime"] == "kata-runtime" {
		log.Infof("RunTime: Send with runtime: %s, %s, %s", podName, nameSpace, "runv")
		runTimeValue = RunvRunTimeTag
	}

	// check Annotation[io.kubernetes.cri.untrusted-workload] = true
	if value, ok := podInfo.Annotations["io.kubernetes.cri.untrusted-workload"]; ok && strings.TrimSpace(value) == "true" {
		runTimeValue = RunvRunTimeTag
	}
	return runTimeValue, nil
}

// IsMountPointRunv check the mountpoint is runv style
func IsMountPointRunv(mountPoint string) bool {
	if IsMounted(mountPoint) {
		return false
	}
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
	pinger.Run()
	stats := pinger.Statistics()
	return stats, nil
}

// IsDirTmpfs check path is tmpfs mounted or not
func IsDirTmpfs(path string) bool {
	cmd := fmt.Sprintf("findmnt %s -o FSTYPE -n", path)
	fsType, err := Run(cmd)
	if err == nil && strings.TrimSpace(fsType) == "tmpfs" {
		return true
	}
	return false
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

func ClearNodeIPCache(nodeID string) {
	nodeAddrMap.Delete(nodeID)
}

// GetNodeAddr get node address
func GetNodeAddr(client kubernetes.Interface, node string, port string) (string, error) {
	ip, err := GetNodeIP(client, node)
	if err != nil {
		return "", err
	}
	return net.JoinHostPort(ip.String(), port), nil
}

// GetNodeIP get node address
func GetNodeIP(client kubernetes.Interface, nodeID string) (net.IP, error) {
	if value, ok := nodeAddrMap.Load(nodeID); ok {
		return value.(net.IP), nil
	}
	node, err := client.CoreV1().Nodes().Get(context.Background(), nodeID, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	addresses := node.Status.Addresses
	addressMap := make(map[v1.NodeAddressType][]v1.NodeAddress)
	for i := range addresses {
		addressMap[addresses[i].Type] = append(addressMap[addresses[i].Type], addresses[i])
	}
	for _, t := range []v1.NodeAddressType{v1.NodeInternalIP, v1.NodeExternalIP} {
		if addresses, ok := addressMap[t]; ok {
			ip := net.ParseIP(addresses[0].Address)
			nodeAddrMap.Store(nodeID, ip)
			return ip, nil
		}
	}
	return nil, fmt.Errorf("Node IP unknown; known addresses: %v", addresses)
}

// CheckParameterValidate is check parameter validating in csi-plugin
func CheckParameterValidate(inputs []string) bool {
	for _, input := range inputs {
		if matched, err := regexp.MatchString("^[A-Za-z0-9=._@:~/-]*$", input); err != nil || !matched {
			return false
		}
	}
	return true
}

// CheckQuotaPathValidate is check quota path validating in csi-plugin
func CheckQuotaPathValidate(kubeClient *kubernetes.Clientset, path string) error {
	pvName := filepath.Base(path)
	_, err := kubeClient.CoreV1().PersistentVolumes().Get(context.Background(), pvName, metav1.GetOptions{})
	if err != nil {
		log.Errorf("utils.CheckQuotaPathValidate %s cannot find volume, error: %s", path, err.Error())
		return err
	}
	return nil
}

// IsHostFileExist is check host file is existing in lvm
func IsHostFileExist(path string) bool {
	args := []string{NsenterCmd, "stat", path}
	cmd := strings.Join(args, " ")
	out, err := Run(cmd)
	if err != nil && strings.Contains(out, "No such file or directory") {
		return false
	}

	return true
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

func DoMountInHost(mntCmd string) error {
	out, err := ConnectorRun(mntCmd)
	if err != nil {
		msg := fmt.Sprintf("Mount is failed in host, mntCmd:%s, err: %s, out: %s", mntCmd, err.Error(), out)
		log.Errorf(msg)
		return errors.New(msg)
	}
	return nil
}

// ConnectorRun Run shell command with host connector
// host connector is daemon running in host.
func ConnectorRun(cmd string) (string, error) {
	c, err := net.Dial("unix", socketPath)
	if err != nil {
		log.Errorf("Oss connector Dial error: %s", err.Error())
		return err.Error(), err
	}
	defer c.Close()

	_, err = c.Write([]byte(cmd))
	if err != nil {
		log.Errorf("Oss connector write error: %s", err.Error())
		return err.Error(), err
	}

	buf := make([]byte, 2048)
	n, err := c.Read(buf[:])
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
	if err := ioutil.WriteFile(dataFilePath, rankingsJSON, 0644); err != nil {
		return err
	}

	log.Infof("AppendJSONData: Json data file saved successfully [%s], content: %v", dataFilePath, curData)
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

// IsKataInstall check kata daemon installed
func IsKataInstall() bool {
	if IsFileExisting("/host/etc/kata-containers") || IsFileExisting("/host/etc/kata-containers2") {
		return true
	}
	return false
}

// IsPathAvailiable
func IsPathAvailiable(path string) error {
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

func WriteMetricsInfo(metricsPathPrefix string, req *csi.NodePublishVolumeRequest, metricsTop string, clientName string, storageBackendName string, fsName string) {
	podUIDPath := metricsPathPrefix + req.VolumeContext["csi.storage.k8s.io/pod.uid"] + "/"
	mountPointPath := podUIDPath + req.GetVolumeId() + "/"
	podInfoName := "pod_info"
	mountPointName := "mount_point_info"
	if !IsFileExisting(mountPointPath) {
		_ = os.MkdirAll(mountPointPath, os.FileMode(0755))
	}
	if !IsFileExisting(podUIDPath + podInfoName) {
		info := req.VolumeContext["csi.storage.k8s.io/pod.namespace"] + " " +
			req.VolumeContext["csi.storage.k8s.io/pod.name"] + " " +
			req.VolumeContext["csi.storage.k8s.io/pod.uid"] + " " +
			metricsTop
		_ = WriteAndSyncFile(podUIDPath+podInfoName, []byte(info), os.FileMode(0644))
	}

	if !IsFileExisting(mountPointPath + mountPointName) {
		info := clientName + " " +
			storageBackendName + " " +
			fsName + " " +
			req.GetVolumeId() + " " +
			req.TargetPath
		_ = WriteAndSyncFile(mountPointPath+mountPointName, []byte(info), os.FileMode(0644))
	}
}

// formatAndMount uses unix utils to format and mount the given disk
func FormatAndMount(diskMounter *k8smount.SafeFormatAndMount, source string, target string, fstype string, mkfsOptions []string, mountOptions []string, omitFsCheck bool) error {
	log.Infof("formatAndMount: mount options : %+v", mountOptions)
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
				log.Warningf("'fsck' not found on system; continuing mount without running 'fsck'.")
			case isExitError && ee.ExitStatus() == fsckErrorsCorrected:
				log.Infof("Device %s has errors which were corrected by fsck.", source)
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
			log.Infof("FormatAndMount: enter special partition logics")
			fsType, ptType, _ := GetDiskPtypePTtype(source)
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

	log.Infof("Disk %q appears to be unformatted, attempting to format as type: %q with options: %v", source, fstype, args)
	startT := time.Now()

	pvName := filepath.Base(source)
	_, err := diskMounter.Exec.Command("mkfs."+fstype, args...).CombinedOutput()
	if err == nil {
		// the disk has been formatted successfully try to mount it again.
		log.Infof("Disk format successed, pvName: %s elapsedTime: %+v ms", pvName, time.Now().Sub(startT).Milliseconds())
		return diskMounter.Interface.Mount(source, target, fstype, mountOptions)
	}
	log.Errorf("format of disk %q failed: type:(%q) target:(%q) options:(%q) error:(%v)", source, fstype, target, args, err)
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

// GetDiskPtypePTtype uses 'blkid' to see if the given disk is unformatted
func GetDiskPtypePTtype(disk string) (fstype string, pttype string, err error) {
	args := []string{"-p", "-s", "TYPE", "-s", "PTTYPE", "-o", "export", disk}

	diskMounter := &k8smount.SafeFormatAndMount{Interface: k8smount.New(""), Exec: utilexec.New()}
	dataOut, err := diskMounter.Exec.Command("blkid", args...).CombinedOutput()
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
		log.Errorf("Could not determine if disk %q is formatted (%v)", disk, err)
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

// Get NVME device name by diskID;
// /dev/nvme0n1 0: means device index, 1: means namespace for nvme device;
// udevadm info --query=all --name=/dev/nvme0n1 | grep ID_SERIAL_SHORT | awk -F= '{print $2}'
// bp1bcfmvsobfauvxb3ow
func GetNvmeDeviceByVolumeID(volumeID string) (device string, err error) {
	serialNumber := strings.TrimPrefix(volumeID, "d-")
	devNumFiles, err := filepath.Glob("/sys/block/*/dev")
	if err != nil {
		log.Infof("GetNvmeDeviceByVolumeID: List device number failed: %v", err)
		return "", err
	}
	for _, deviceFile := range devNumFiles {
		deviceName := filepath.Base(filepath.Dir(deviceFile))
		if strings.HasPrefix(deviceName, "nvme") && !strings.Contains(deviceName, "p") {
			cmd := fmt.Sprintf("%s udevadm info --query=all --name=/dev/%s | grep ID_SERIAL_SHORT | awk -F= '{print $2}'", NsenterCmd, deviceName)
			snumber, err := Run(cmd)
			if err != nil {
				log.Warnf("GetNvmeDeviceByVolumeID: Get device with command %s and got error: %s", cmd, err.Error())
				continue
			}
			snumber = strings.TrimSpace(snumber)
			if serialNumber == strings.TrimSpace(snumber) {
				device = filepath.Join("/dev/", deviceName)
				log.Infof("GetNvmeDeviceByVolumeID: Get nvme device %s with volumeID %s", device, volumeID)
				return device, nil
			}
		}
	}
	return "", nil
}
