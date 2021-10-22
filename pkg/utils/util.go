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
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/go-ping/ping"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"io/ioutil"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	v1core "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/record"
	k8svol "k8s.io/kubernetes/pkg/volume"
	"k8s.io/kubernetes/pkg/volume/util/fs"
	"net/http"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"reflect"
	"strings"
	"syscall"
	"time"
)

//DefaultOptions used for global ak
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
	// MetadataMaxRetrycount ...
	MetadataMaxRetrycount = 4
)

// KubernetesAlicloudIdentity set a identity label
var KubernetesAlicloudIdentity = fmt.Sprintf("Kubernetes.Alicloud/CsiPlugin")

// RoleAuth define STS Token Response
type RoleAuth struct {
	AccessKeyID     string
	AccessKeySecret string
	Expiration      time.Time
	SecurityToken   string
	LastUpdated     time.Time
	Code            string
}

//CreateEvent is create events
func CreateEvent(recorder record.EventRecorder, objectRef *v1.ObjectReference, eventType string, reason string, err string) {
	recorder.Event(objectRef, eventType, reason, err)
}

//NewEventRecorder is create snapshots event recorder
func NewEventRecorder() record.EventRecorder {
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Fatalf("NewControllerServer: Failed to create config: %v", err)
	}
	clientset, err := kubernetes.NewForConfig(config)
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

// Run run shell command
func Run(cmd string) (string, error) {
	out, err := exec.Command("sh", "-c", cmd).CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("Failed to run cmd: " + cmd + ", with out: " + string(out) + ", with error: " + err.Error())
	}
	return string(out), nil
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

//IsLikelyNotMountPoint return status of mount point,this function fix IsMounted return 0 bug
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
	cmd := fmt.Sprintf("mount | grep %s | grep -v grep | wc -l", mountPath)
	out, err := Run(cmd)
	if err != nil {
		log.Infof("IsMounted: exec error: %s, %s", cmd, err.Error())
		return false
	}
	if strings.TrimSpace(out) == "0" {
		return false
	}
	return true
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

//GetMetaData get metadata from ecs meta-server
func GetMetaData(resource string) (string, error) {
	resp, err := http.Get(MetadataURL + resource)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

// RetryGetMetaData ...
func RetryGetMetaData(resource string) string {
	var nodeID string
	for i := 0; i < MetadataMaxRetrycount; i++ {
		nodeID, _ = GetMetaData(resource)
		if strings.Contains(nodeID, "Error 500 Internal Server Error") {
			if i == MetadataMaxRetrycount-1 {
				log.Fatalf("NewDriver:: Access metadata server failed: %v", nodeID)
			}
			continue
		}
		return nodeID
	}
	return nodeID
}

// GetRegionIDAndInstanceID get regionID and instanceID object
func GetRegionIDAndInstanceID(nodeName string) (string, string, error) {
	strs := strings.SplitN(nodeName, ".", 2)
	if len(strs) < 2 {
		return "", "", fmt.Errorf("failed to get regionID and instanceId from nodeName")
	}
	return strs[0], strs[1], nil
}

// WriteJSONFile write a json object
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
func NewEcsClient(accessKeyID, accessKeySecret, accessToken string) (ecsClient *ecs.Client) {
	var err error
	if accessToken == "" {
		ecsClient, err = ecs.NewClientWithAccessKey(DefaultRegion, accessKeyID, accessKeySecret)
		if err != nil {
			return nil
		}
	} else {
		ecsClient, err = ecs.NewClientWithStsToken(DefaultRegion, accessKeyID, accessKeySecret, accessToken)
		if err != nil {
			return nil
		}
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
	available, capacity, usage, inodes, inodesFree, inodesUsed, err := fs.FsInfo(path)
	if err != nil {
		return nil, err
	}

	metrics := &k8svol.Metrics{Time: metav1.Now()}
	metrics.Available = resource.NewQuantity(available, resource.BinarySI)
	metrics.Capacity = resource.NewQuantity(capacity, resource.BinarySI)
	metrics.Used = resource.NewQuantity(usage, resource.BinarySI)
	metrics.Inodes = resource.NewQuantity(inodes, resource.BinarySI)
	metrics.InodesFree = resource.NewQuantity(inodesFree, resource.BinarySI)
	metrics.InodesUsed = resource.NewQuantity(inodesUsed, resource.BinarySI)

	metricAvailable, ok := (*(metrics.Available)).AsInt64()
	if !ok {
		log.Errorf("failed to fetch available bytes for target: %s", path)
		return nil, status.Error(codes.Unknown, "failed to fetch available bytes")
	}
	metricCapacity, ok := (*(metrics.Capacity)).AsInt64()
	if !ok {
		log.Errorf("failed to fetch capacity bytes for target: %s", path)
		return nil, status.Error(codes.Unknown, "failed to fetch capacity bytes")
	}
	metricUsed, ok := (*(metrics.Used)).AsInt64()
	if !ok {
		log.Errorf("failed to fetch used bytes for target %s", path)
		return nil, status.Error(codes.Unknown, "failed to fetch used bytes")
	}
	metricInodes, ok := (*(metrics.Inodes)).AsInt64()
	if !ok {
		log.Errorf("failed to fetch available inodes for target %s", path)
		return nil, status.Error(codes.Unknown, "failed to fetch available inodes")
	}
	metricInodesFree, ok := (*(metrics.InodesFree)).AsInt64()
	if !ok {
		log.Errorf("failed to fetch free inodes for target: %s", path)
		return nil, status.Error(codes.Unknown, "failed to fetch free inodes")
	}
	metricInodesUsed, ok := (*(metrics.InodesUsed)).AsInt64()
	if !ok {
		log.Errorf("failed to fetch used inodes for target: %s", path)
		return nil, status.Error(codes.Unknown, "failed to fetch used inodes")
	}

	return &csi.NodeGetVolumeStatsResponse{
		Usage: []*csi.VolumeUsage{
			{
				Available: metricAvailable,
				Total:     metricCapacity,
				Used:      metricUsed,
				Unit:      csi.VolumeUsage_BYTES,
			},
			{
				Available: metricInodesFree,
				Total:     metricInodes,
				Used:      metricInodesUsed,
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

// WriteJosnFile save json data to file
func WriteJosnFile(obj interface{}, file string) error {
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
