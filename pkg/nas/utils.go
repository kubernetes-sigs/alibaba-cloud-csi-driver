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
	"fmt"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cnfs/v1beta1"
	"io/ioutil"
	"k8s.io/client-go/dynamic"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"

	"errors"
	aliyunep "github.com/aliyun/alibaba-cloud-sdk-go/sdk/endpoints"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	aliNas "github.com/aliyun/alibaba-cloud-sdk-go/services/nas"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	log "github.com/sirupsen/logrus"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// MetadataURL is metadata url
	MetadataURL = "http://100.100.100.200/latest/meta-data/"
	// RegionTag is region id
	RegionTag = "region-id"
	// NsenterCmd is nsenter mount command
	NsenterCmd = "/nsenter --mount=/proc/1/ns/mnt"
	// LoopLockFile lock file for nas loopsetup
	LoopLockFile = "loopsetup.nas.csi.alibabacloud.com.lck"
	// LoopImgFile image file for nas loopsetup
	LoopImgFile = "loopsetup.nas.csi.alibabacloud.com.img"
	// Resize2fsFailedFilename ...
	Resize2fsFailedFilename = "resize2fs_failed.txt"
	// Resize2fsFailedFixCmd ...
	Resize2fsFailedFixCmd = "%s fsck -a %s"
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

//DoNfsMount execute the mount command for nas dir
func DoNfsMount(nfsServer, nfsPath, nfsVers, mountOptions, mountPoint, volumeID, podUID, useNasClient string) error {
	if !utils.IsFileExisting(mountPoint) {
		CreateDest(mountPoint)
	}

	if CheckNfsPathMounted(mountPoint, nfsServer, nfsPath) {
		log.Infof("DoNfsMount: nfs server already mounted: %s, %s", nfsServer, nfsPath)
		return nil
	}

	var mntCmd string
	var err error
	nfsServer = "959684a3e6-ege66.cn-zhangjiakou.nas.aliyuncs.com"
	if len(useNasClient) != 0 && useNasClient == "true" {
		mntCmd = fmt.Sprintf("systemd-run --scope -- mount -t alinas -o unas -o client_owner=%s %s:%s %s", podUID, nfsServer, nfsPath, mountPoint)
		if out, err := utils.ConnectorRun(mntCmd); err != nil {
			if err != nil {
				log.Errorf("Mount by nas rich client, error: %s", err.Error())
				return errors.New("Create nas user space volume fail: " + err.Error() + ", out: " + out)
			}
		}
	} else {
		mntCmd = fmt.Sprintf("mount -t nfs -o vers=%s %s:%s %s", nfsVers, nfsServer, nfsPath, mountPoint)
		if mountOptions != "" {
			mntCmd = fmt.Sprintf("mount -t nfs -o vers=%s,%s %s:%s %s", nfsVers, mountOptions, nfsServer, nfsPath, mountPoint)
		}
		_, err = utils.Run(mntCmd)
	}
	if err != nil && nfsPath != "/" {
		if strings.Contains(err.Error(), "reason given by server: No such file or directory") || strings.Contains(err.Error(), "access denied by server while mounting") {
			if err := createNasSubDir(nfsServer, nfsPath, nfsVers, mountOptions, volumeID); err != nil {
				log.Errorf("DoNfsMount: Create SubPath error: %s", err.Error())
				return err
			}
			if _, err := utils.Run(mntCmd); err != nil {
				log.Errorf("DoNfsMount, Mount Nfs sub directory fail: %s", err.Error())
				return err
			}
		} else {
			return err
		}
	} else if err != nil {
		return err
	}
	log.Infof("DoNfsMount: mount nfs successful with command: %s", mntCmd)
	return nil
}

//CheckNfsPathMounted check whether the given nfs path was mounted
func CheckNfsPathMounted(mountpoint, server, path string) bool {
	// mntCmd := fmt.Sprintf("findmnt %s | grep %s | grep %s | grep -v grep | wc -l", mountpoint, server, path)
	mntCmd := fmt.Sprintf("cat /proc/mounts | grep %s | grep %s | grep -v grep | wc -l", mountpoint, path)
	// mntCmd := fmt.Sprintf("grep -E -- '%s.*%s' /proc/mounts", mountpoint, path)
	if out, err := utils.Run(mntCmd); err == nil && strings.TrimSpace(out) != "0" {
		return true
	}
	return false
}

//CreateDest create the target
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

// GetMetaData get host regionid, zoneid
func GetMetaData(resource string) string {
	resp, err := http.Get(MetadataURL + resource)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	return string(body)
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
	if len(regionID) == 0 {
		regionID = GetMetaData(RegionTag)
	}
	var err error
	switch ac.UseMode {
	case utils.AccessKey:
		nasClient, err = aliNas.NewClientWithAccessKey(regionID, ac.AccessKeyID, ac.AccessKeySecret)
	case utils.Credential:
		nasClient, err = aliNas.NewClientWithOptions(regionID, ac.Config, ac.Credential)
	default:
		nasClient, err = aliNas.NewClientWithStsToken(regionID, ac.AccessKeyID, ac.AccessKeySecret, ac.StsToken)
	}

	if err != nil {
		return nil
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

func waitTimeout(wg *sync.WaitGroup, timeout int) bool {
	c := make(chan struct{})
	go func() {
		defer close(c)
		wg.Wait()
	}()
	select {
	case <-c:
		return false
	case <-time.After(time.Duration(timeout) * time.Second):
		return true
	}

}

func createNasSubDir(nfsServer, nfsPath, nfsVers, nfsOptions string, volumeID string) error {
	// step 1: create mount path
	nasTmpPath := filepath.Join(NasTempMntPath, volumeID)
	if err := utils.CreateDest(nasTmpPath); err != nil {
		log.Infof("Create Nas temp Directory err: " + err.Error())
		return err
	}
	if utils.IsMounted(nasTmpPath) {
		utils.Umount(nasTmpPath)
	}

	// step 2: do mount, and create subpath by extreme
	if nfsOptions != "" {
		nfsVers = nfsVers + "," + nfsOptions
	}
	usePath := nfsPath
	rootPath := "/"
	//server is extreme nas
	if strings.Contains(nfsServer, "extreme.nas.aliyuncs.com") {
		//1.No need to deal with the case where nfsPath only beginning with /share or /share/
		//2.No need to deal with the case where nfspath does not contain /share or /share/ at the beginning
		//3.Need to deal with the case where nfsPath is /share/subpath
		if strings.HasPrefix(nfsPath, "/share/") && len(nfsPath) > len("/share/") {
			rootPath = "/share/"
			usePath = nfsPath[6:]
		}
	}

	mntCmdRootPath := fmt.Sprintf("mount -t nfs -o vers=%s %s:%s %s", nfsVers, nfsServer, rootPath, nasTmpPath)
	_, err := utils.Run(mntCmdRootPath)
	if err != nil {
		log.Errorf("Nas, mount directory rootPath fail, rootPath:%s, err:%s", rootPath, err.Error())
		return err
	}

	subPath := path.Join(nasTmpPath, usePath)
	if err := utils.CreateDest(subPath); err != nil {
		log.Infof("Nas, Create Sub Directory fail, subPath:%s, err: " + err.Error())
		return err
	}

	// step 3: umount after create
	utils.Umount(nasTmpPath)
	log.Infof("Create Sub Directory successful, nfsPath:%s, subPath:%s", nfsPath, subPath)
	return nil
}

func setNasVolumeCapacity(nfsServer, nfsPath string, volSizeBytes int64) error {
	if nfsPath == "" || nfsPath == "/" {
		return fmt.Errorf("Volume %s:%s not support set quota to root path ", nfsServer, nfsPath)
	}
	pvSizeGB := volSizeBytes / (1024 * 1024 * 1024)
	nasClient := updateNasClient(GlobalConfigVar.NasClient, GetMetaData(RegionTag))
	fsList := strings.Split(nfsServer, "-")
	if len(fsList) < 1 {
		return fmt.Errorf("volume error nas server(%s) ", nfsServer)
	}
	quotaRequest := aliNas.CreateSetDirQuotaRequest()
	quotaRequest.FileSystemId = fsList[0]
	quotaRequest.Path = nfsPath
	quotaRequest.UserType = "AllUsers"
	quotaRequest.QuotaType = "Enforcement"
	pvSizeGBStr := strconv.FormatInt(pvSizeGB, 10)
	quotaRequest.SizeLimit = requests.Integer(pvSizeGBStr)
	quotaRequest.RegionId = GetMetaData(RegionTag)
	_, err := nasClient.SetDirQuota(quotaRequest)
	if err != nil {
		if strings.Contains(err.Error(), "The specified FileSystem does not exist.") {
			return fmt.Errorf("extreme did not support quota, please change %s to General Purpose NAS", nfsServer)
		}
		return fmt.Errorf("volume set nas quota with error: %s", err.Error())
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
	updateNasConfig := false
	sunRPCFile := "/etc/modprobe.d/sunrpc.conf"
	if !utils.IsFileExisting(sunRPCFile) {
		updateNasConfig = true
	} else {
		chkCmd := fmt.Sprintf("cat %s | grep tcp_slot_table_entries | grep 128 | grep -v grep | wc -l", sunRPCFile)
		out, err := utils.Run(chkCmd)
		if err != nil {
			log.Warnf("Update Nas system config check error: %s", err.Error())
			return
		}
		if strings.TrimSpace(out) == "0" {
			updateNasConfig = true
		}
	}

	if updateNasConfig {
		upCmd := fmt.Sprintf("echo \"options sunrpc tcp_slot_table_entries=128\" >> %s && echo \"options sunrpc tcp_max_slot_table_entries=128\" >> %s && sysctl -w sunrpc.tcp_slot_table_entries=128", sunRPCFile, sunRPCFile)
		_, err := utils.Run(upCmd)
		if err != nil {
			log.Warnf("Update Nas system config error: %s", err.Error())
			return
		}
		log.Warnf("Successful update Nas system config")
	}
}

// ParseMountFlags parse mountOptions
func ParseMountFlags(mntOptions []string) (string, string) {
	if len(mntOptions) > 0 {
		mntOptionsStr := strings.Join(mntOptions, ",")
		// mntOptions should re-split, as some like ["a,b,c", "d"]
		mntOptionsList := strings.Split(mntOptionsStr, ",")
		tmpOptionsList := []string{}

		if strings.Contains(mntOptionsStr, "vers=3.0") {
			for _, tmpOptions := range mntOptionsList {
				if tmpOptions != "vers=3.0" {
					tmpOptionsList = append(tmpOptionsList, tmpOptions)
				}
			}
			return "3", strings.Join(tmpOptionsList, ",")
		} else if strings.Contains(mntOptionsStr, "vers=3") {
			for _, tmpOptions := range mntOptionsList {
				if tmpOptions != "vers=3" {
					tmpOptionsList = append(tmpOptionsList, tmpOptions)
				}
			}
			return "3", strings.Join(tmpOptionsList, ",")
		} else if strings.Contains(mntOptionsStr, "vers=4.0") {
			for _, tmpOptions := range mntOptionsList {
				if tmpOptions != "vers=4.0" {
					tmpOptionsList = append(tmpOptionsList, tmpOptions)
				}
			}
			return "4.0", strings.Join(tmpOptionsList, ",")
		} else if strings.Contains(mntOptionsStr, "vers=4.1") {
			for _, tmpOptions := range mntOptionsList {
				if tmpOptions != "vers=4.1" {
					tmpOptionsList = append(tmpOptionsList, tmpOptions)
				}
			}
			return "4.1", strings.Join(tmpOptionsList, ",")
		} else {
			return "", strings.Join(mntOptions, ",")
		}
	}
	return "", ""
}

func createLosetupPv(fullPath string, volSizeBytes int64) error {
	blockNum := volSizeBytes / (4 * 1024)
	fileName := filepath.Join(fullPath, LoopImgFile)
	if utils.IsFileExisting(fileName) {
		log.Infof("createLosetupPv: image file is exist, just skip: %s", fileName)
		return nil
	}
	imgCmd := fmt.Sprintf("dd if=/dev/zero of=%s bs=4k seek=%d count=0", fileName, blockNum)
	_, err := utils.Run(imgCmd)
	if err != nil {
		return err
	}

	formatCmd := fmt.Sprintf("mkfs.ext4 -F -m0 %s", fileName)
	_, err = utils.Run(formatCmd)
	if err != nil {
		return err
	}
	return nil
}

// /var/lib/kubelet/pods/5e03c7f7-2946-4ee1-ad77-2efbc4fdb16c/volumes/kubernetes.io~csi/nas-f5308354-725a-4fd3-b613-0f5b384bd00e/mount
func mountLosetupPv(mountPoint string, opt *Options, volumeID string) error {
	// if target path mounted already, return
	if utils.IsMounted(mountPoint) {
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
	err := DoNfsMount(opt.Server, opt.Path, opt.Vers, opt.Options, nfsPath, volumeID, podID, "false")
	if err != nil {
		return fmt.Errorf("mountLosetupPv: mount losetup volume failed: %s", err.Error())
	}

	lockFile := filepath.Join(nfsPath, LoopLockFile)
	if opt.LoopLock == "true" && isLosetupUsed(lockFile, opt, volumeID) {
		return fmt.Errorf("mountLosetupPv: nfs losetup file is used by others %s", lockFile)
	}
	imgFile := filepath.Join(nfsPath, LoopImgFile)
	failedFile := filepath.Join(nfsPath, Resize2fsFailedFilename)
	if utils.IsFileExisting(failedFile) {
		// path/to/whatever does not exist
		cmd := fmt.Sprintf(Resize2fsFailedFixCmd, NsenterCmd, imgFile)
		_, err = utils.Run(cmd)
		if err != nil {
			return fmt.Errorf("mountLosetupPv: mount nfs losetup error %s", err.Error())
		}
		err = os.Remove(failedFile)
		if err != nil {
			log.Errorf("mountLosetupPv: failed to remove failed file: %v", err)
		}
	}
	mountCmd := fmt.Sprintf("%s mount -o loop %s %s", NsenterCmd, imgFile, mountPoint)
	_, err = utils.Run(mountCmd)
	if err != nil {
		return fmt.Errorf("mountLosetupPv: mount nfs losetup error %s", err.Error())
	}
	lockContent := GlobalConfigVar.NodeID + ":" + GlobalConfigVar.NodeIP
	if err := ioutil.WriteFile(lockFile, ([]byte)(lockContent), 0644); err != nil {
		return err
	}
	return nil
}

func isLosetupUsed(lockFile string, opt *Options, volumeID string) bool {
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
		if !isLosetupMount(volumeID) {
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

func checkLosetupUnmount(mountPoint string) error {
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
		if err := utils.Umount(nfsPath); err != nil {
			return fmt.Errorf("checkLosetupUnmount: umount nfs path error %v", err)
		}
		log.Infof("Losetup Unmount successful %s", mountPoint)
	} else {
		log.Infof("Losetup Unmount, image file not exist, skipping %s", mountPoint)
	}
	return nil
}

func isLosetupMount(volumeID string) bool {
	keyWord := volumeID + "/" + LoopImgFile
	cmd := fmt.Sprintf("mount | grep %s |grep -v grep |wc -l", keyWord)
	out, err := utils.Run(cmd)
	if err != nil {
		log.Infof("isLosetupMount: exec error: %s, %s", cmd, err.Error())
		return false
	}
	if strings.TrimSpace(out) == "0" {
		return false
	}
	return true
}

func getPvObj(volumeID string) (*v1.PersistentVolume, error) {
	return GlobalConfigVar.KubeClient.CoreV1().PersistentVolumes().Get(context.Background(), volumeID, metav1.GetOptions{})
}

func isValidCnfsParameter(server string, cnfsName string) error {
	if len(server) == 0 && len(cnfsName) == 0 {
		msg := fmt.Sprintf("Server and ContainerNetworkFileSystem need to be configured at least one.")
		log.Errorf(msg)
		return errors.New(msg)
	}

	if len(server) != 0 && len(cnfsName) != 0 {
		msg := fmt.Sprintf("Server and ContainerNetworkFileSystem can only be configured to use one.")
		log.Errorf(msg)
		return errors.New(msg)
	}
	return nil
}

//GetFsIDByServer func is get fsID from serverName
func GetFsIDByServer(server string) string {
	if len(server) == 0 {
		return ""
	}
	serverArray := strings.Split(server, "-")
	return serverArray[0]
}
