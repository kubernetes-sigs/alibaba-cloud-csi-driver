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
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/alibabacloud-go/tea/tea"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/losetup"
	mounter "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/nas/cloud"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/nas/interfaces"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"k8s.io/klog/v2"
	mountutils "k8s.io/mount-utils"
)

const (
	// RegionTag is region id
	RegionTag = "region-id"
	// LoopLockFile lock file for nas loopsetup
	LoopLockFile = "loopsetup.nas.csi.alibabacloud.com.lck"
	// LoopImgFile image file for nas loopsetup
	LoopImgFile = "loopsetup.nas.csi.alibabacloud.com.img"
	// Resize2fsFailedFilename ...
	Resize2fsFailedFilename = "resize2fs_failed.txt"
	// Resize2fsFailedFixCmd ...
	Resize2fsFailedFixCmd = "%s fsck -a %s"
	// GiB bytes
	GiB = 1 << 30
	// see https://help.aliyun.com/zh/nas/modify-the-maximum-number-of-concurrent-nfs-requests
	TcpSlotTableEntries      = "/proc/sys/sunrpc/tcp_slot_table_entries"
	TcpSlotTableEntriesValue = "128\n"

	filesystemIDKey   = "fileSystemId"
	filesystemTypeKey = "fileSystemType"
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

func doMount(mounter mountutils.Interface, opt *Options, targetPath, volumeId, podUid string, agentMode bool) error {
	var (
		mountFstype     string
		source          string
		combinedOptions []string
		isPathNotFound  func(error) bool
	)
	if opt.Accesspoint != "" {
		source = fmt.Sprintf("%s:%s", opt.Accesspoint, opt.Path)
	} else {
		source = fmt.Sprintf("%s:%s", opt.Server, opt.Path)
	}
	if opt.Options != "" {
		combinedOptions = append(combinedOptions, opt.Options)
	}

	switch opt.ClientType {
	case EFCClient:
		combinedOptions = append(combinedOptions, "efc", fmt.Sprintf("bindtag=%s", volumeId), fmt.Sprintf("client_owner=%s", podUid))
		switch opt.FSType {
		case "standard":
		case "cpfs":
			combinedOptions = append(combinedOptions, "protocol=nfs3")
		default:
			return errors.New("EFC Client don't support this storage type:" + opt.FSType)
		}
		mountFstype = "alinas"
		// err = mounter.Mount(source, mountPoint, "alinas", combinedOptions)
		isPathNotFound = isEFCPathNotFoundError
	case NativeClient:
		switch opt.FSType {
		case "cpfs":
		default:
			return errors.New("Native Client don't support this storage type:" + opt.FSType)
		}
		mountFstype = "cpfs"
	default:
		//NFS Mount(Capacdity/Performance Extreme Nas、Cpfs2.0, AliNas)
		versStr := fmt.Sprintf("vers=%s", opt.Vers)
		if !strings.Contains(opt.Options, versStr) {
			combinedOptions = append(combinedOptions, versStr)
		}
		if opt.Accesspoint != "" {
			mountFstype = "alinas"
			// must enable tls when using accesspoint
			combinedOptions = addTLSMountOptions(combinedOptions)
		}
		isPathNotFound = isNFSPathNotFoundError
		// Enable compatibility for BMCPFS VPC mount points using the NAS Driver, to support the same usage in ECI.
		if isCPFS(opt.FSType, opt.Server) && opt.MountProtocol == "efc" {
			mountFstype = "alinas"
			combinedOptions = append(combinedOptions, "efc,protocol=efc,net=tcp,fstype=cpfs")
			if agentMode {
				isPathNotFound = isEFCPathNotFoundError
			} else {
				klog.V(4).InfoS("When using EFC, automatic directory creation is only supported in agent mode.")
			}
		}
		if mountFstype == "" {
			mountFstype = opt.MountProtocol
		}
	}
	err := mounter.Mount(source, targetPath, mountFstype, combinedOptions)
	if err == nil {
		return nil
	}
	if isPathNotFound == nil || !isPathNotFound(err) {
		return err
	}

	rootPath := "/"
	if opt.FSType == "cpfs" || mountFstype == MountProtocolCPFSNFS || isExtrameNAS(opt.FSType, opt.Server) {
		rootPath = "/share"
	}
	relPath, relErr := filepath.Rel(rootPath, opt.Path)
	if relErr != nil || relPath == "." {
		return err
	}
	rootSource := fmt.Sprintf("%s:%s", opt.Server, rootPath)
	klog.Infof("trying to create subpath %s in %s", opt.Path, opt.Server)
	tmpPath := filepath.Join(utils.KubeletRootDir, "csi-plugins", driverName, "node")
	if err := os.MkdirAll(tmpPath, 0o700); err != nil {
		return err
	}
	tmpPath, err = os.MkdirTemp(tmpPath, "subpath-creation_"+volumeId+"_")
	if err != nil {
		return err
	}
	defer os.Remove(tmpPath)
	if err := mounter.Mount(rootSource, tmpPath, mountFstype, combinedOptions); err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Join(tmpPath, relPath), os.ModePerm); err != nil {
		return err
	}
	if err := cleanupMountpoint(mounter, tmpPath); err != nil {
		klog.Errorf("failed to cleanup tmp mountpoint %s: %v", tmpPath, err)
	}
	return mounter.Mount(source, targetPath, mountFstype, combinedOptions)
}

func isEFCPathNotFoundError(err error) bool {
	return strings.Contains(err.Error(), "Failed to bind mount")
}

func isNFSPathNotFoundError(err error) bool {
	errors := []string{
		"reason given by server: No such file or directory",
		"access denied by server while mounting",
	}
	for _, error := range errors {
		if strings.Contains(err.Error(), error) {
			return true
		}
	}
	return false
}

// check system config,
// if tcp_slot_table_entries not set to 128, just config.
func checkSystemNasConfig() error {
	data, err := os.ReadFile(TcpSlotTableEntries)
	if err == nil && string(data) == TcpSlotTableEntriesValue {
		return nil
	}
	return os.WriteFile(TcpSlotTableEntries, []byte(TcpSlotTableEntriesValue), os.ModePerm)
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

func addTLSMountOptions(baseOptions []string) []string {
	for _, options := range baseOptions {
		for _, option := range mounter.SplitMountOptions(options) {
			if option == "" {
				continue
			}
			key, _, _ := strings.Cut(option, "=")
			if key == "tls" {
				return baseOptions
			}
		}
	}
	return append(baseOptions, "tls")
}

func createLosetupPv(fullPath string, volSizeBytes int64) error {
	fileName := filepath.Join(fullPath, LoopImgFile)
	if utils.IsFileExisting(fileName) {
		klog.Infof("createLosetupPv: image file is exist, just skip: %s", fileName)
		return nil
	}
	if err := losetup.TruncateFile(fileName, volSizeBytes); err != nil {
		return err
	}
	return exec.Command("mkfs.ext4", "-F", "-m0", fileName).Run()
}

func unmountLosetupPv(mounter mountutils.Interface, mountPoint string) error {
	pathList := strings.Split(mountPoint, "/")
	if len(pathList) != 10 {
		klog.Infof("MountPoint not format as losetup type: %s", mountPoint)
		return nil
	}
	podID := pathList[5]
	pvName := pathList[8]
	nfsPath := filepath.Join(NasMntPoint, podID, pvName)
	imgFile := filepath.Join(nfsPath, LoopImgFile)
	lockFile := filepath.Join(nfsPath, LoopLockFile)
	if utils.IsFileExisting(imgFile) {
		klog.Infof("cleanup the losetup mountpoint due to the existence of image file %s", imgFile)
		if utils.IsFileExisting(lockFile) {
			if err := os.Remove(lockFile); err != nil {
				return fmt.Errorf("checkLosetupUnmount: remove lock file error %v", err)
			}
		}
		if err := cleanupMountpoint(mounter, nfsPath); err != nil {
			return fmt.Errorf("checkLosetupUnmount: umount nfs path error %v", err)
		}
		klog.Infof("Losetup Unmount successful %s", mountPoint)
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

// rund-csi 2.0
func saveVolumeData(opt *Options, mountPath string) error {
	// save volume data to json file
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
	return utils.AppendJSONData(fileName, volumeData)
}

func cleanupMountpoint(mounter mountutils.Interface, mountPath string) (err error) {
	forceUnmounter, ok := mounter.(mountutils.MounterForceUnmounter)
	if ok {
		err = mountutils.CleanupMountWithForce(mountPath, forceUnmounter, false, time.Second*30)
	} else {
		err = mountutils.CleanupMountPoint(mountPath, mounter, false)
	}
	return
}

func getNASIDFromMapOrServer(params map[string]string, server string) string {
	if id := params[filesystemIDKey]; id != "" {
		return id
	}
	id, _, _ := strings.Cut(server, "-")
	return id
}

func getCPFSIDFromMapOrServer(params map[string]string, server string) string {
	if id := params[filesystemIDKey]; id != "" {
		return id
	}
	elems := strings.Split(server, "-")
	switch len(elems) {
	case 0:
		return ""
	case 1:
		return elems[0]
	default:
		return elems[0] + "-" + elems[1]
	}
}

func getFilesystemTypeFromAPIOrServer(filesystemID, server string, client interfaces.NasClientV2Interface) string {
	fs, err := client.DescribeFileSystems(filesystemID)
	if err != nil {
		klog.ErrorS(err, "DescribeFileSystems failed")
	}
	if fs != nil && fs.Body != nil && fs.Body.FileSystems != nil && len(fs.Body.FileSystems.FileSystem) > 0 {
		fs := fs.Body.FileSystems.FileSystem[0]
		if fs != nil && fs.FileSystemType != nil {
			return tea.StringValue(fs.FileSystemType)
		}
	}
	return cloud.GetFilesystemTypeByMountTargetDomain(server)
}

func isExtrameNAS(filesystemType, server string) bool {
	if filesystemType != "" {
		return filesystemType == "extreme"
	}
	return strings.Contains(server, "extreme.nas.aliyuncs.com")
}

func isCPFS(filesystemType, server string) bool {
	if filesystemType != "" {
		return filesystemType == "cpfs"
	}
	return strings.HasSuffix(server, "cpfs.aliyuncs.com")
}
