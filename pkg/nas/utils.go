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

package nas

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"time"

	aliyunep "github.com/aliyun/alibaba-cloud-sdk-go/sdk/endpoints"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	aliNas "github.com/aliyun/alibaba-cloud-sdk-go/services/nas"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cnfs/v1beta1"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/losetup"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	log "github.com/sirupsen/logrus"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/dynamic"
	mountutils "k8s.io/mount-utils"
)

const (
	// MetadataURL is metadata url
	MetadataURL = "http://100.100.100.200/latest/meta-data/"
	// RegionTag is region id
	RegionTag = "region-id"
	// NsenterCmd is nsenter mount command
	NsenterCmd = "nsenter --mount=/proc/1/ns/mnt"
	//Nsenter is nsenter binary command
	Nsenter = "nsenter"
	// LoopLockFile lock file for nas loopsetup
	LoopLockFile = "loopsetup.nas.csi.alibabacloud.com.lck"
	// LoopImgFile image file for nas loopsetup
	LoopImgFile = "loopsetup.nas.csi.alibabacloud.com.img"
	// Resize2fsFailedFilename ...
	Resize2fsFailedFilename = "resize2fs_failed.txt"
	// Resize2fsFailedFixCmd ...
	Resize2fsFailedFixCmd = "%s fsck -a %s"

	GiB = 1 << 30
)

var (
	// VERSION should be updated by hand at each release
	VERSION = "v1.14.8"
	// GITCOMMIT will be overwritten automatically by the build system
	GITCOMMIT = "HEAD"
	// KubernetesAlicloudIdentity is the system identity for ecs client request
	KubernetesAlicloudIdentity = fmt.Sprintf("Kubernetes.Alicloud/CsiProvision.Nas-%s", VERSION)
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

// DoMount execute the mount command for nas dir
func DoMount(mounter mountutils.Interface, fsType, clientType, nfsProtocol, nfsServer, nfsPath, nfsVers, mountOptions, mountPoint, volumeID, podUID string) error {
	if !utils.IsFileExisting(mountPoint) {
		_ = CreateDest(mountPoint)
	}

	notMounted, err := mounter.IsLikelyNotMountPoint(mountPoint)
	if err != nil {
		return err
	}
	if !notMounted {
		log.Infof("%s already mounted", mountPoint)
		return nil
	}

	source := fmt.Sprintf("%s:%s", nfsServer, nfsPath)
	var combinedOptions []string
	if mountOptions != "" {
		combinedOptions = append(combinedOptions, mountOptions)
	}

	switch clientType {
	case EFCClient:
		combinedOptions = append(combinedOptions, "efc", fmt.Sprintf("bindtag=%s", volumeID), fmt.Sprintf("client_owner=%s", podUID))
		switch fsType {
		case "standard":
		case "cpfs":
			combinedOptions = append(combinedOptions, "protocol=nfs3")
		default:
			return errors.New("EFC Client don't support this storage type:" + fsType)
		}
		err = mounter.Mount(source, mountPoint, "alinas", combinedOptions)
	case NativeClient:
		switch fsType {
		case "cpfs":
		default:
			return errors.New("Native Client don't support this storage type:" + fsType)
		}
		err = mounter.Mount(source, mountPoint, "cpfs", combinedOptions)
	default:
		//NFS Mount(Capacdity/Performance Extreme Nas、Cpfs2.0, AliNas)
		versStr := fmt.Sprintf("vers=%s", nfsVers)
		if !strings.Contains(mountOptions, versStr) {
			combinedOptions = append(combinedOptions, versStr)
		}
		//nfsProtocol: cpfs-nfs/nfs/alinas
		err = mounter.Mount(source, mountPoint, nfsProtocol, combinedOptions)
		//try mount nfsPath is successfully.
		if err == nil {
			break
		}
		//mount root-path is failed, return error
		if nfsPath == "/" {
			break
		}
		//Other errors
		if !strings.Contains(err.Error(), "reason given by server: No such file or directory") && !strings.Contains(err.Error(), "access denied by server while mounting") {
			break
		}
		//mount sub-path is failed, if subpath is not exist or cpfs don't output subpath
		if err := createSubDir(mounter, nfsProtocol, nfsServer, nfsPath, combinedOptions, volumeID); err != nil {
			log.Errorf("DoMount: Create subpath is failed, err: %s", err.Error())
			return err
		}
		err = mounter.Mount(source, mountPoint, nfsProtocol, combinedOptions)
	}
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{"source": source, "target": mountPoint, "client": clientType}).Info("mounted successfully")
	return nil
}

// CreateDest create the target
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

// GetNfsDetails get nfs server's details
func GetNfsDetails(nfsServersString string) (string, string) {
	nfsServer, nfsPath := "", ""
	nfsServerList := strings.Split(nfsServersString, ",")
	serverNum := len(nfsServerList)

	if _, ok := storageClassServerPos[nfsServersString]; !ok {
		storageClassServerPos[nfsServersString] = 0
	}
	zoneIndex := storageClassServerPos[nfsServersString] % serverNum
	selectedServer := nfsServerList[zoneIndex]
	storageClassServerPos[nfsServersString]++

	serverParts := strings.Split(selectedServer, ":")
	if len(serverParts) == 1 {
		nfsServer = serverParts[0]
		nfsPath = "/"
	} else if len(serverParts) == 2 {
		nfsServer = serverParts[0]
		nfsPath = serverParts[1]
		if nfsPath == "" {
			nfsPath = "/"
		}
	} else {
		nfsServer = ""
		nfsPath = ""
	}

	// remove / if path end with /;
	if nfsPath != "/" && strings.HasSuffix(nfsPath, "/") {
		nfsPath = nfsPath[0 : len(nfsPath)-1]
	}

	return nfsServer, nfsPath
}

func updateNasClient(client *aliNas.Client, regionID string) *aliNas.Client {
	ac := utils.GetAccessControl()
	if ac.UseMode == utils.EcsRAMRole || ac.UseMode == utils.ManagedToken {
		client = newNasClient(ac, regionID)
	}
	if client.Client.GetConfig() != nil {
		client.Client.GetConfig().UserAgent = KubernetesAlicloudIdentity
	}
	return client
}

func newNasClient(ac utils.AccessControl, regionID string) (nasClient *aliNas.Client) {
	nasClient, err := aliNas.NewClientWithOptions(regionID, ac.Config, ac.Credential)
	if err != nil {
		return nil
	}

	if os.Getenv("ALICLOUD_CLIENT_SCHEME") != "HTTP" {
		nasClient.SetHTTPSInsecure(false)
		nasClient.GetConfig().WithScheme("HTTPS")
	}

	// Set Nas Endpoint
	SetNasEndPoint(regionID)
	return
}

// SetNasEndPoint Set Endpoint for Nas
func SetNasEndPoint(regionID string) {
	// use unitized region endpoint for blew regions.
	// total 16 regions
	unitizedRegions := []string{"cn-hangzhou", "cn-zhangjiakou", "cn-huhehaote", "cn-shenzhen", "ap-southeast-1", "ap-southeast-2",
		"ap-southeast-3", "ap-southeast-5", "eu-central-1", "us-east-1", "ap-northeast-1", "ap-south-1",
		"us-west-1", "eu-west-1", "cn-chengdu", "cn-north-2-gov-1", "cn-beijing", "cn-shanghai", "cn-hongkong",
		"cn-shenzhen-finance-1", "cn-shanghai-finance-1", "cn-hangzhou-finance", "cn-qingdao"}
	for _, tmpRegion := range unitizedRegions {
		if regionID == tmpRegion {
			aliyunep.AddEndpointMapping(regionID, "Nas", "nas-vpc."+regionID+".aliyuncs.com")
			break
		}
	}

	// use environment endpoint setting first;
	if ep := os.Getenv("NAS_ENDPOINT"); ep != "" {
		aliyunep.AddEndpointMapping(regionID, "Nas", ep)
	}
}

func createSubDir(mounter mountutils.Interface, nfsProtocol, nfsServer, nfsPath string, nfsOptions []string, volumeID string) error {
	rootPath := "/"
	usePath := nfsPath
	//The root directory of cpfs-nfs and extreme NAS is /share
	if nfsProtocol == MountProtocolCPFSNFS || strings.Contains(nfsServer, "extreme.nas.aliyuncs.com") {
		//1.No need to deal with the case where nfsPath only beginning with /share or /share/
		//2.No need to deal with the case where nfspath does not contain /share or /share/ at the beginning
		//3.Need to deal with the case where nfsPath is /share/subpath
		if strings.HasPrefix(nfsPath, "/share/") && len(nfsPath) > len("/share/") {
			rootPath = "/share/"
			usePath = nfsPath[len("/share/"):]
		}
	}

	nasTmpPath := filepath.Join(NasTempMntPath, volumeID)
	if err := utils.CreateDest(nasTmpPath); err != nil {
		log.Infof("Create nas tempPath is failed, tmpPath:%s, err: %s", nasTmpPath, err.Error())
		return err
	}
	notMounted, err := mounter.IsLikelyNotMountPoint(nasTmpPath)
	if err != nil {
		return fmt.Errorf("check mountpoint %s: %w", nasTmpPath, err)
	}
	if !notMounted {
		if err := mounter.Unmount(nasTmpPath); err != nil {
			return err
		}
	}
	err = mounter.Mount(fmt.Sprintf("%s:%s", nfsServer, rootPath), nasTmpPath, nfsProtocol, nfsOptions)
	if err != nil {
		log.Errorf("Mount rootPath is failed, rootPath:%s, err:%s", rootPath, err.Error())
		return err
	}
	subPath := path.Join(nasTmpPath, usePath)
	if err := utils.CreateDest(subPath); err != nil {
		log.Infof("Create subPath is failed, subPath:%s, err: %s", subPath, err.Error())
		return err
	}
	log.Infof("Create subPath is successfully, nfsPath:%s, subPath:%s", nfsPath, path.Join(nasTmpPath, usePath))
	return mounter.Unmount(nasTmpPath)
}

func setNasVolumeCapacity(nfsServer, nfsPath string, volSizeBytes int64) error {
	if nfsPath == "" || nfsPath == "/" {
		return fmt.Errorf("Volume %s:%s not support set quota to root path ", nfsServer, nfsPath)
	}
	nasClient := updateNasClient(GlobalConfigVar.NasClient, GlobalConfigVar.Region)
	fsList := strings.Split(nfsServer, "-")
	if len(fsList) < 1 {
		return fmt.Errorf("volume error nas server(%s) ", nfsServer)
	}
	quotaRequest := aliNas.CreateSetDirQuotaRequest()
	quotaRequest.FileSystemId = fsList[0]
	quotaRequest.Path = nfsPath
	quotaRequest.UserType = "AllUsers"
	quotaRequest.QuotaType = "Enforcement"
	quotaRequest.SizeLimit = requests.NewInteger64((volSizeBytes + GiB - 1) / GiB)
	quotaRequest.RegionId, _ = utils.GetMetaData(RegionTag)
	log.Infof("SetDirQuotaRequest: %+v", quotaRequest)
	_, err := nasClient.SetDirQuota(quotaRequest)
	if err != nil {
		return fmt.Errorf("SetDirQuota: %w", err)
	}
	return nil
}

func setNasVolumeCapacityWithID(pvObj *v1.PersistentVolume, crdClient dynamic.Interface, volSizeBytes int64) error {
	if pvObj.Spec.CSI == nil {
		return fmt.Errorf("Volume %s is not CSI type %v ", pvObj.Name, pvObj)
	}

	// Check Pv volume parameters
	if value, ok := pvObj.Spec.CSI.VolumeAttributes["volumeCapacity"]; ok && value == "false" {
		return fmt.Errorf("Volume %s not contain volumeCapacity parameters, not support expand, PV: %v ", pvObj.Name, pvObj)
	}
	nfsServer, nfsPath := "", ""
	if value, ok := pvObj.Spec.CSI.VolumeAttributes["server"]; ok {
		nfsServer = value
	} else {
		if value, ok := pvObj.Spec.CSI.VolumeAttributes["containerNetworkFileSystem"]; ok {
			cnfs, err := v1beta1.GetCnfsObject(crdClient, value)
			if err != nil {
				return err
			}
			nfsServer = cnfs.Status.FsAttributes.Server
		}
	}
	if value, ok := pvObj.Spec.CSI.VolumeAttributes["path"]; ok {
		nfsPath = value
	}
	return setNasVolumeCapacity(nfsServer, nfsPath, volSizeBytes)
}

// check system config,
// if tcp_slot_table_entries not set to 128, just config.
func checkSystemNasConfig() {
	sunRPCFile := "/etc/modprobe.d/sunrpc.conf"
	if utils.IsFileExisting(sunRPCFile) {
		data, err := os.ReadFile(sunRPCFile)
		if err != nil {
			log.Warnf("Update Nas system config check error: %s", err.Error())
			return
		}
		for _, line := range strings.Split(string(data), "\n") {
			if strings.Contains(line, "tcp_slot_table_entries") && strings.Contains(line, "128") {
				return
			}
		}
	}

	sunRpcConfig := "\"options sunrpc tcp_slot_table_entries=128\"\n\"options sunrpc tcp_max_slot_table_entries=128\""
	// startRpcConfig := "sysctl -w sunrpc.tcp_slot_table_entries=128"
	err := utils.WriteAndSyncFile(sunRPCFile, []byte(sunRpcConfig), 0755)
	if err != nil {
		log.Warnf("Write nas rpcconfig %s is failed, err: %s", sunRPCFile, err.Error())
		return
	}
	err = exec.Command("sysctl", "-w", "sunrpc.tcp_slot_table_entries=128").Run()
	if err != nil {
		log.Warnf("Start nas rpcconfig is failed, err: %s", err.Error())
		return
	}
	log.Warnf("Update nas system config is successfully.")
}

// ParseMountFlags parse mountOptions
func ParseMountFlags(mntOptions []string) (string, string) {
	var vers string
	var otherOptions []string
	for _, options := range mntOptions {
		for _, option := range mounter.SplitMountOptions(options) {
			if option == "" {
				continue
			}
			key, value, found := strings.Cut(option, "=")
			if found && key == "vers" {
				vers = value
			} else {
				otherOptions = append(otherOptions, option)
			}
		}
	}
	if vers == "3.0" {
		vers = "3"
	}
	return vers, strings.Join(otherOptions, ",")
}

func createLosetupPv(fullPath string, volSizeBytes int64) error {
	fileName := filepath.Join(fullPath, LoopImgFile)
	if utils.IsFileExisting(fileName) {
		log.Infof("createLosetupPv: image file is exist, just skip: %s", fileName)
		return nil
	}
	if err := losetup.TruncateFile(fileName, volSizeBytes); err != nil {
		return err
	}
	return exec.Command("mkfs.ext4", "-F", "-m0", fileName).Run()
}

// /var/lib/kubelet/pods/5e03c7f7-2946-4ee1-ad77-2efbc4fdb16c/volumes/kubernetes.io~csi/nas-f5308354-725a-4fd3-b613-0f5b384bd00e/mount
func mountLosetupPv(mounter mountutils.Interface, mountPoint string, opt *Options, volumeID string) error {
	// if target path mounted already, return
	notMounted, err := mounter.IsLikelyNotMountPoint(mountPoint)
	if err != nil {
		return fmt.Errorf("check mount point %s: %w", mountPoint, err)
	}
	if !notMounted {
		log.Infof("mountLosetupPv: TargetPath(%s) is mounted as tmpfs, not need mount again", mountPoint)
		return nil
	}

	pathList := strings.Split(mountPoint, "/")
	if len(pathList) != 10 {
		return fmt.Errorf("mountLosetupPv: mountPoint format error, %s", mountPoint)
	}

	podID := pathList[5]
	pvName := pathList[8]

	// /mnt/nasplugin.alibabacloud.com/6c690876-74aa-46f6-a301-da7f4353665d/pv-losetup/
	nfsPath := filepath.Join(NasMntPoint, podID, pvName)
	if err := utils.CreateDest(nfsPath); err != nil {
		return fmt.Errorf("mountLosetupPv: create nfs mountPath error %s ", err.Error())
	}
	//mount nas to use losetup dev
	err = DoMount(mounter, "nas", "nfsclient", opt.MountProtocol, opt.Server, opt.Path, opt.Vers, opt.Options, nfsPath, volumeID, podID)
	if err != nil {
		return fmt.Errorf("mountLosetupPv: mount losetup volume failed: %s", err.Error())
	}

	lockFile := filepath.Join(nfsPath, LoopLockFile)
	if opt.LoopLock == "true" && isLosetupUsed(mounter, lockFile, opt, volumeID) {
		return fmt.Errorf("mountLosetupPv: nfs losetup file is used by others %s", lockFile)
	}
	imgFile := filepath.Join(nfsPath, LoopImgFile)
	failedFile := filepath.Join(nfsPath, Resize2fsFailedFilename)
	if utils.IsFileExisting(failedFile) {
		// path/to/whatever does not exist
		cmd := exec.Command("fsck", "-a", imgFile)
		// cmd := fmt.Sprintf(Resize2fsFailedFixCmd, NsenterCmd, imgFile)
		err := cmd.Run()
		if err != nil {
			return fmt.Errorf("mountLosetupPv: mount nfs losetup error %s", err.Error())
		}
		err = os.Remove(failedFile)
		if err != nil {
			log.Errorf("mountLosetupPv: failed to remove failed file: %v", err)
		}
	}
	err = mounter.Mount(imgFile, mountPoint, "", []string{"loop"})
	if err != nil {
		return fmt.Errorf("mountLosetupPv: mount nfs losetup error %s", err.Error())
	}
	lockContent := GlobalConfigVar.NodeID + ":" + GlobalConfigVar.NodeIP
	if err := os.WriteFile(lockFile, ([]byte)(lockContent), 0644); err != nil {
		return err
	}
	return nil
}

func isLosetupUsed(mounter mountutils.Interface, lockFile string, opt *Options, volumeID string) bool {
	if !utils.IsFileExisting(lockFile) {
		return false
	}
	fileCotent := utils.GetFileContent(lockFile)
	contentParts := strings.Split(fileCotent, ":")
	if len(contentParts) != 2 || contentParts[0] == "" || contentParts[1] == "" {
		return true
	}

	oldNodeID := contentParts[0]
	oldNodeIP := contentParts[1]
	if GlobalConfigVar.NodeID == oldNodeID {
		mounted, err := isLosetupMount(volumeID)
		if err != nil {
			log.Errorf("can not determine whether %s losetup image used: %v", volumeID, err)
			return true
		}
		if !mounted {
			log.Warnf("Lockfile(%s) exist, but Losetup image not mounted %s.", lockFile, opt.Path)
			return false
		}
		log.Warnf("Lockfile(%s) exist, but Losetup image mounted %s.", lockFile, opt.Path)
		return true
	}

	stat, err := utils.Ping(oldNodeIP)
	if err != nil {
		log.Warnf("Ping node %s, but get error: %s, consider as volume used", oldNodeIP, err.Error())
		return true
	}
	if stat.PacketLoss == 100 {
		log.Warnf("Cannot connect to node %s, consider the node as shutdown(%s).", oldNodeIP, lockFile)
		return false
	}
	return true
}

func unmountLosetupPv(mounter mountutils.Interface, mountPoint string) error {
	pathList := strings.Split(mountPoint, "/")
	if len(pathList) != 10 {
		log.Infof("MountPoint not format as losetup type: %s", mountPoint)
		return nil
	}
	podID := pathList[5]
	pvName := pathList[8]
	nfsPath := filepath.Join(NasMntPoint, podID, pvName)
	imgFile := filepath.Join(nfsPath, LoopImgFile)
	lockFile := filepath.Join(nfsPath, LoopLockFile)
	if utils.IsFileExisting(imgFile) {
		if utils.IsFileExisting(lockFile) {
			if err := os.Remove(lockFile); err != nil {
				return fmt.Errorf("checkLosetupUnmount: remove lock file error %v", err)
			}
		}
		if err := mounter.Unmount(nfsPath); err != nil {
			return fmt.Errorf("checkLosetupUnmount: umount nfs path error %v", err)
		}
		log.Infof("Losetup Unmount successful %s", mountPoint)
	} else {
		log.Infof("Losetup Unmount, image file not exist, skipping %s", mountPoint)
	}
	return nil
}

func isLosetupMount(volumeID string) (bool, error) {
	devices, err := losetup.List()
	if err != nil {
		return false, err
	}
	for _, device := range devices {
		if strings.Contains(device.BackFile, volumeID+"/"+LoopImgFile) {
			return true, nil
		}
	}
	return false, nil
}

func getPvObj(volumeID string) (*v1.PersistentVolume, error) {
	return GlobalConfigVar.KubeClient.CoreV1().PersistentVolumes().Get(context.Background(), volumeID, metav1.GetOptions{})
}

func isValidCnfsParameter(server string, cnfsName string) error {
	if len(server) == 0 && len(cnfsName) == 0 {
		msg := "Server and Cnfs need to be configured at least one."
		log.Errorf(msg)
		return errors.New(msg)
	}

	if len(server) != 0 && len(cnfsName) != 0 {
		msg := "Server and Cnfs can only be configured to use one."
		log.Errorf(msg)
		return errors.New(msg)
	}
	return nil
}

// GetFsIDByNasServer func is get fsID from serverName
func GetFsIDByNasServer(server string) string {
	if len(server) == 0 {
		return ""
	}
	serverArray := strings.Split(server, "-")
	return serverArray[0]
}

// GetFsIDByCpfsServer func is get fsID from serverName
func GetFsIDByCpfsServer(server string) string {
	if len(server) == 0 {
		return ""
	}
	serverArray := strings.Split(server, "-")
	return serverArray[0] + "-" + serverArray[1]
}

func saveVolumeData(opt *Options, mountPath string) {
	// save volume data to json file
	if utils.IsKataInstall() {
		volumeData := map[string]string{}
		volumeData["csi.alibabacloud.com/version"] = opt.Vers
		if len(opt.Options) != 0 {
			volumeData["csi.alibabacloud.com/options"] = opt.Options
		}
		if len(opt.Server) != 0 {
			volumeData["csi.alibabacloud.com/server"] = opt.Server
		}
		if len(opt.Path) != 0 {
			volumeData["csi.alibabacloud.com/path"] = opt.Path
		}
		fileName := filepath.Join(filepath.Dir(mountPath), utils.VolDataFileName)
		if strings.HasSuffix(mountPath, "/") {
			fileName = filepath.Join(filepath.Dir(filepath.Dir(mountPath)), utils.VolDataFileName)
		}
		if err := utils.AppendJSONData(fileName, volumeData); err != nil {
			log.Warnf("NodePublishVolume: append nas volume spec to %s with error: %s", fileName, err.Error())
		}
	}
}
