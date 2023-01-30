/*

Copyright 2017 Google Inc.

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

package server

import (
	"crypto/sha256"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"unicode"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/local/lib"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/local/manager"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/local/types"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

const (
	// NsenterCmd is the nsenter command
	NsenterCmd = "/nsenter --mount=/proc/1/ns/mnt --ipc=/proc/1/ns/ipc --net=/proc/1/ns/net --uts=/proc/1/ns/uts "
	// ProjQuotaPrefix is the template of quota fullpath
	ProjQuotaPrefix = "/mnt/quotapath.%s/%s"
	// ProjQuotaNamespacePrefix ...
	ProjQuotaNamespacePrefix = "/mnt/quotapath.%s"
)

func CreateLoopDevice(templateFile, pvName, quotaSize string, lp manager.LoopDevice) (string, error) {

	lpPath := filepath.Join(filepath.Dir(templateFile), fmt.Sprintf("%s.img", pvName))
	_, err := os.Stat(lpPath)
	if err == nil {
		log.Errorf("CreateLoopDevice: target sparsefile: %s exists, previous loopdevice not cleaned", lpPath)
	}
	if !os.IsNotExist(err) {
		log.Errorf("CreateLoopDevice: check loopdevice sparsefile path err: %v", err)
		return "", err
	}
	log.Infof("CreateLoopDevice: lpPath %s is not exists", lpPath)
	if value, err := strconv.Atoi(quotaSize); err != nil {
		return "", fmt.Errorf("CreateLoopDevice: invalid quotasize:%s, err:%v", quotaSize, err)
	} else {
		_, size := lp.GetTemplateInfo()
		if value == size*1024*1024*1024 {
			err := lp.CopySparseFile(templateFile, lpPath)
			if err != nil {
				log.Errorf("CreateLoopDevice: copy sparsefile failed err:%v", err)
				return "", err
			}
		} else {
			err := lp.CreateSparseFile(lpPath, quotaSize)
			if err != nil {
				log.Errorf("CreateLoopDevice: create sparsefile failed err:%v", err)
				return "", err
			}
			err = lp.FormatFile(lpPath)
			if err != nil {
				log.Errorf("CreateLoopDevice: failed to create sparsefile. err: %v", err)
				return "", err
			}
		}
	}

	return lp.CreateLoopDevice(lpPath)
}

func DeleteLoopDevice(pvName string) error {
	ld := manager.NewLoopDevice(types.GlobalConfigVar.LocalSparseFileDir, types.GlobalConfigVar.LocalSparseFileTempSize)
	lpPath := filepath.Join(types.GlobalConfigVar.LocalSparseFileDir, fmt.Sprintf("%s.img", pvName))
	out, err := ld.DeleteLoopDevice(lpPath)
	log.Infof("DeleteLoopDevice: delete loopdevice with out:%s, err:%v", out, err)
	return err
}

// ListLV lists lvm volumes
func ListLV(listspec string) ([]*lib.LV, error) {
	if !utils.CheckParameterValidate([]string{listspec}) {
		return nil, fmt.Errorf("inputs illegal: %s", listspec)
	}

	lvs := []*lib.LV{}
	cmdList := []string{NsenterCmd, "lvs", "--units=b", "--separator=\"<:SEP:>\"", "--nosuffix", "--noheadings",
		"-o", "lv_name,lv_size,lv_uuid,lv_attr,copy_percent,lv_kernel_major,lv_kernel_minor,lv_tags", "--nameprefixes", "-a", listspec}
	cmd := strings.Join(cmdList, " ")
	out, err := utils.Run(cmd)
	if err != nil {
		if strings.Contains(err.Error(), "Failed to find logical volume") {
			return lvs, nil
		}
		return nil, err
	}
	outStr := strings.TrimSpace(string(out))
	outLines := strings.Split(outStr, "\n")
	for _, line := range outLines {
		line = strings.TrimSpace(line)
		if !strings.Contains(line, "LVM2_LV_NAME") {
			continue
		}
		lv, err := lib.ParseLV(line)
		if err != nil {
			return nil, errors.New("Parse LVM: " + line + ", with error: " + err.Error())
		}
		lvs = append(lvs, lv)
	}
	return lvs, nil
}

// CreateLV creates a new volume
func CreateLV(ctx context.Context, vg string, name string, size uint64, mirrors uint32, tags []string, striping bool) (string, error) {
	if !utils.CheckParameterValidate([]string{vg, name}) {
		return "", fmt.Errorf("inputs illegal: %s, %s", vg, name)
	}
	if !utils.CheckParameterValidate(tags) {
		return "", fmt.Errorf("inputs illegal: %v", tags)
	}

	if size == 0 {
		return "", errors.New("size must be greater than 0")
	}
	args := []string{NsenterCmd, "lvcreate", "-n", name, "-L", fmt.Sprintf("%db", size), "-W", "y", "-y"}
	if mirrors > 0 {
		args = append(args, "-m", fmt.Sprintf("%d", mirrors), "--nosync")
	}
	for _, tag := range tags {
		args = append(args, "--add-tag", tag)
	}
	if striping {
		pvCount := getPVNumber(vg)
		if pvCount != 0 {
			args = append(args, "-i", strconv.Itoa(pvCount))
		}
	}

	args = append(args, vg)
	cmd := strings.Join(args, " ")
	out, err := utils.Run(cmd)
	return string(out), err
}

func getPVNumber(vgName string) int {
	var pvCount = 0
	vgList, err := ListVG()
	if err != nil {
		return 0
	}
	for _, vg := range vgList {
		if vg.Name == vgName {
			pvCount = int(vg.PvCount)
		}
	}
	return pvCount
}

// ProtectedTagName is a tag that prevents RemoveLV & RemoveVG from removing a volume
const ProtectedTagName = "protected"

// RemoveLV removes a volume
func RemoveLV(ctx context.Context, vg string, name string) (string, error) {
	if !utils.CheckParameterValidate([]string{vg, name}) {
		return "", fmt.Errorf("inputs illegal: %s, %s", vg, name)
	}

	lvs, err := ListLV(fmt.Sprintf("%s/%s", vg, name))
	if err != nil {
		return "", fmt.Errorf("failed to list LVs: %v", err)
	}
	if len(lvs) == 0 {
		return "lvm " + vg + "/" + name + " is not exist, skip remove", nil
	}
	if len(lvs) != 1 {
		return "", fmt.Errorf("expected 1 LV, got %d", len(lvs))
	}
	for _, tag := range lvs[0].Tags {
		if tag == ProtectedTagName {
			return "", errors.New("volume is protected")
		}
	}

	args := []string{NsenterCmd, "lvremove", "-v", "-f", fmt.Sprintf("%s/%s", vg, name)}
	cmd := strings.Join(args, " ")
	out, err := utils.Run(cmd)

	return string(out), err
}

// CloneLV clones a volume via dd
func CloneLV(ctx context.Context, src, dest string) (string, error) {
	// FIXME(farcaller): bloody insecure. And broken.
	if !utils.CheckParameterValidate([]string{src, dest}) {
		return "", fmt.Errorf("inputs illegal: %s, %s", src, dest)
	}

	args := []string{NsenterCmd, "dd", fmt.Sprintf("if=%s", src), fmt.Sprintf("of=%s", dest), "bs=4M"}
	cmd := strings.Join(args, " ")
	out, err := utils.Run(cmd)

	return string(out), err
}

// ListVG get vg info
func ListVG() ([]*lib.VG, error) {
	args := []string{NsenterCmd, "vgs", "--units=b", "--separator=\"<:SEP:>\"", "--nosuffix", "--noheadings",
		"-o", "vg_name,vg_size,vg_free,vg_uuid,vg_tags,pv_count", "--nameprefixes", "-a"}
	cmd := strings.Join(args, " ")
	out, err := utils.Run(cmd)
	if err != nil {
		return nil, err
	}
	outStr := strings.TrimSpace(string(out))
	outLines := strings.Split(outStr, "\n")

	vgs := []*lib.VG{}
	for _, line := range outLines {
		line = strings.TrimSpace(line)
		if !strings.Contains(line, "LVM2_VG_NAME") {
			continue
		}
		vg, err := lib.ParseVG(line)
		if err != nil {
			return nil, err
		}
		vgs = append(vgs, vg)
	}
	return vgs, nil
}

// CreateVG create volume group
func CreateVG(ctx context.Context, name string, physicalVolume string, tags []string) (string, error) {
	if !utils.CheckParameterValidate([]string{name, physicalVolume}) {
		return "", fmt.Errorf("inputs illegal: %s, %s", name, physicalVolume)
	}
	if !utils.CheckParameterValidate(tags) {
		return "", fmt.Errorf("inputs illegal: %s", tags)
	}

	args := []string{NsenterCmd, "vgcreate", name, physicalVolume, "-v"}
	for _, tag := range tags {
		args = append(args, "--add-tag", tag)
	}
	cmd := strings.Join(args, " ")
	out, err := utils.Run(cmd)

	return string(out), err
}

// RemoveVG remove volume group
func RemoveVG(ctx context.Context, name string) (string, error) {
	if !utils.CheckParameterValidate([]string{name}) {
		return "", fmt.Errorf("inputs illegal: %s", name)
	}

	vgs, err := ListVG()
	if err != nil {
		return "", fmt.Errorf("failed to list VGs: %v", err)
	}
	var vg *lib.VG
	for _, v := range vgs {
		if v.Name == name {
			vg = v
			break
		}
	}
	if vg == nil {
		return "", fmt.Errorf("could not find vg to delete")
	}
	for _, tag := range vg.Tags {
		if tag == ProtectedTagName {
			return "", errors.New("volume is protected")
		}
	}

	args := []string{NsenterCmd, "vgremove", "-v", "-f", name}
	cmd := strings.Join(args, " ")
	out, err := utils.Run(cmd)

	return string(out), err
}

// CleanPath deletes all the contents under the given directory
func CleanPath(ctx context.Context, path string) error {
	if !utils.CheckParameterValidate([]string{path}) {
		return fmt.Errorf("inputs illegal: %s", path)
	}

	dir, err := os.Open(path)
	if err != nil {
		return err
	}
	defer dir.Close()

	files, err := dir.Readdirnames(-1)
	if err != nil {
		return err
	}
	errList := []error{}
	for _, file := range files {
		err = os.RemoveAll(filepath.Join(path, file))
		if err != nil {
			errList = append(errList, err)
		}
	}
	if len(errList) == 0 {
		return nil
	}
	return errList[0]
}

// AddTagLV add tag
func AddTagLV(ctx context.Context, vg string, name string, tags []string) (string, error) {
	if !utils.CheckParameterValidate([]string{vg, name}) {
		return "", fmt.Errorf("inputs illegal: %s, %s", vg, name)
	}
	if !utils.CheckParameterValidate(tags) {
		return "", fmt.Errorf("inputs illegal: %s", tags)
	}

	lvs, err := ListLV(fmt.Sprintf("%s/%s", vg, name))
	if err != nil {
		return "", fmt.Errorf("failed to list LVs: %v", err)
	}
	if len(lvs) != 1 {
		return "", fmt.Errorf("expected 1 LV, got %d", len(lvs))
	}

	args := make([]string, 0)
	args = append(args, NsenterCmd)
	args = append(args, "lvchange")
	for _, tag := range tags {
		args = append(args, "--addtag", tag)
	}

	args = append(args, fmt.Sprintf("%s/%s", vg, name))
	cmd := strings.Join(args, " ")
	out, err := utils.Run(cmd)

	return string(out), err
}

// RemoveTagLV remove tag
func RemoveTagLV(ctx context.Context, vg string, name string, tags []string) (string, error) {
	if !utils.CheckParameterValidate([]string{vg, name}) {
		return "", fmt.Errorf("inputs illegal: %s, %s", vg, name)
	}
	if !utils.CheckParameterValidate(tags) {
		return "", fmt.Errorf("inputs illegal: %s", tags)
	}

	lvs, err := ListLV(fmt.Sprintf("%s/%s", vg, name))
	if err != nil {
		return "", fmt.Errorf("failed to list LVs: %v", err)
	}
	if len(lvs) != 1 {
		return "", fmt.Errorf("expected 1 LV, got %d", len(lvs))
	}

	args := make([]string, 0)
	args = append(args, NsenterCmd)
	args = append(args, "lvchange")
	for _, tag := range tags {
		args = append(args, "--deltag", tag)
	}

	args = append(args, fmt.Sprintf("%s/%s", vg, name))
	cmd := strings.Join(args, " ")
	out, err := utils.Run(cmd)
	return string(out), err
}

// CreateNameSpace creates a new namespace
// ndctl create-namespace -r region0 --size=6G -n webpmem1
func CreateNameSpace(ctx context.Context, region string, name string, size uint64) (string, error) {
	if !utils.CheckParameterValidate([]string{region, name}) {
		return "", fmt.Errorf("inputs illegal: %s, %s", region, name)
	}

	if size == 0 {
		return "", errors.New("size must be greater than 0")
	}
	args := []string{NsenterCmd, "ndctl", "create-namespace", "-r", region, "-s", fmt.Sprintf("%d", size), "-n", name}
	cmd := strings.Join(args, " ")
	out, err := utils.Run(cmd)
	return string(out), err
}

// GetNameSpace get pmem namespace
func GetNameSpace(namespaceName string) (*lib.NameSpace, error) {
	namespaces, err := ListNameSpace()
	if err != nil {
		return nil, err
	}
	if len(namespaces) == 0 {
		return nil, nil
	}

	for _, tmp := range namespaces {
		if tmp.Dev == namespaceName {
			return tmp, nil
		}
	}
	return nil, nil
}

// RemoveNameSpace removes a namespace
func RemoveNameSpace(ctx context.Context, namespaceName string) (string, error) {
	if !utils.CheckParameterValidate([]string{namespaceName}) {
		return "", fmt.Errorf("inputs illegal: %s", namespaceName)
	}

	namespace, err := GetNameSpace(namespaceName)
	if err != nil {
		return "", fmt.Errorf("failed to found namespace with error: %v", err)
	}
	if namespace == nil {
		return "namespace  " + namespaceName + " not found, skip", nil
	}

	args := []string{NsenterCmd, "ndctl", "disable-namespace", fmt.Sprintf("%s", namespace.Dev)}
	cmd := strings.Join(args, " ")
	_, err = utils.Run(cmd)
	if err != nil {
		return "", fmt.Errorf("failed to disable namespace with error: %v", err)
	}

	args = []string{NsenterCmd, "ndctl", "destroy-namespace", fmt.Sprintf("%s", namespace.Dev)}
	cmd = strings.Join(args, " ")
	_, err = utils.Run(cmd)
	if err != nil {
		return "", fmt.Errorf("failed to destroy namespace with error: %v", err)
	}
	return "", nil
}

// ListNameSpace list pmem namespace
func ListNameSpace() ([]*lib.NameSpace, error) {
	regions, err := manager.GetRegions()
	log.Infof("show me the regions: %+v", regions)
	if err != nil {
		return nil, err
	}

	namespaces := []*lib.NameSpace{}
	for _, region := range regions.Regions {
		for _, ns := range region.Namespaces {
			newns := ns.ToProto()
			namespaces = append(namespaces, newns)
		}
	}
	return namespaces, nil
}

// ConvertString2int convert pvName to int data
func ConvertString2int(origin string) string {
	h := sha256.New()
	h.Write([]byte(origin))
	hashResult := fmt.Sprintf("%x", h.Sum(nil))
	for {
		if hashResult[0] == '0' {
			hashResult = hashResult[1:]
			continue
		}
		break
	}
	return str2ASCII(hashResult[:9])[:9]
}

func str2ASCII(origin string) string {
	var result string
	for _, c := range origin {
		if !unicode.IsDigit(c) {
			result += strconv.Itoa(int(rune(c)))
		} else {
			result += string(c)
		}
	}
	return result
}

// SetProjectID2PVSubpath ...
func SetProjectID2PVSubpath(subPath, fullPath, rootPath, filesystem string, run utils.CommandRunFunc) (string, error) {
	if !utils.CheckParameterValidate([]string{subPath, fullPath, rootPath, filesystem}) {
		return "", fmt.Errorf("inputs illegal: %s, %s, %s, %s", subPath, fullPath, rootPath, filesystem)
	}

	projectID := ConvertString2int(subPath)
	var args []string
	switch filesystem {
	case "ext4", "ext3":
		args = []string{NsenterCmd, "chattr", "+P -p", fmt.Sprintf("%s %s", projectID, fullPath)}
	case "xfs":
		args = []string{NsenterCmd, "xfs_quota", "-x -c", fmt.Sprintf("'project -s -p %s %s' %s", fullPath, projectID, rootPath)}
	default:
		return "", fmt.Errorf("SetProjectID2PVSubpath:: unsupport filesystem: %s type", filesystem)
	}
	cmd := strings.Join(args, " ")
	_, err := run(cmd)
	if err != nil {
		return "", fmt.Errorf("failed to set projectID to subpath with error: %v", err)
	}
	return projectID, nil
}

func getTotalLimitKBFromCSV(in string) (totalLimit int64, err error) {
	r := csv.NewReader(strings.NewReader(in))
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return 0, err
		}
		trimedStr := strings.TrimSpace(record[0])
		if strings.HasPrefix(trimedStr, "#") && trimedStr != "#0" {
			if err != nil {
				return 0, err
			}
			limitKByte, err := strconv.ParseInt(record[5], 10, 64)
			if err != nil {
				return 0, err
			}
			totalLimit += limitKByte
		}
	}
	return
}

// GetNamespaceAssignedQuota ...
func GetNamespaceAssignedQuota(namespace string) (int, error) {
	if !utils.CheckParameterValidate([]string{namespace}) {
		return 0, fmt.Errorf("inputs illegal: %s", namespace)
	}

	args := []string{NsenterCmd, "repquota", "-P -O csv", fmt.Sprintf(ProjQuotaNamespacePrefix, namespace)}
	cmd := strings.Join(args, " ")
	out, err := utils.Run(cmd)
	if err != nil {
		return 0, fmt.Errorf("failed to request namespace quota with error: %v", err)
	}
	totalLimit, err := getTotalLimitKBFromCSV(out)
	if err != nil {
		return 0, err
	}

	return int(totalLimit), nil
}

// SelectNamespace ...
func SelectNamespace(ctx context.Context, quotaSize string) (string, error) {
	namespaces, err := ListNameSpace()
	if err != nil {
		return "", err
	}
	if len(namespaces) != 2 {
		return "", fmt.Errorf("namespaces count is wrong in current ecs %v", len(namespaces))
	}
	namespace1Quota, err := GetNamespaceAssignedQuota(namespaces[0].Dev)
	namespace2Quota, err := GetNamespaceAssignedQuota(namespaces[1].Dev)
	if err != nil {
		return "", err
	}
	if namespace1Quota > namespace2Quota {
		return namespaces[1].Dev, nil
	}
	return namespaces[0].Dev, nil
}

// CreateProjQuotaSubpath ...
func CreateProjQuotaSubpath(ctx context.Context, subPath, quotaSize, rootPath string) (string, string, string, error) {
	if !utils.CheckParameterValidate([]string{subPath, quotaSize, rootPath}) {
		return "", "", "", fmt.Errorf("inputs illegal: %s, %s, %s", subPath, quotaSize, rootPath)
	}

	var fullPath string
	var err error
	if len(rootPath) == 0 {
		selectedNamespace, err := SelectNamespace(ctx, quotaSize)
		if err != nil {
			return "", "", "", err
		}
		fullPath = fmt.Sprintf(ProjQuotaPrefix, selectedNamespace, subPath)
	} else {
		fullPath = filepath.Join(rootPath, subPath)
	}
	fsType := getFileSystemType(rootPath)
	if fsType == "" {
		return "", "", "", fmt.Errorf("CreateProjQuotaSubpath:: failed to get rootPath: %s filesystem", rootPath)
	}
	err = manager.EnsureFolder(fullPath)
	if err != nil {
		return "", "", "", err
	}
	projectID, err := SetProjectID2PVSubpath(subPath, fullPath, rootPath, fsType, utils.Run)
	if err != nil {
		return "", "", "", err
	}
	return fullPath, "", projectID, nil
}

func findBlockLimitByProjectID(out, projectID string) (string, string, error) {
	r := csv.NewReader(strings.NewReader(out))
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", "", err
		}
		trimedStr := strings.TrimSpace(record[0])
		if strings.HasPrefix(trimedStr, "#") && trimedStr != "#0" {
			if strings.Contains(trimedStr, projectID) {
				return record[4], record[5], nil
			}
		}
	}
	return "", "", fmt.Errorf("findBlockLimitByProjectID: cannot find projectID: %s in output", projectID)
}

func getFileSystemType(rootPath string) string {
	args := []string{NsenterCmd, "df", "-Th", rootPath, " | awk 'NR != 1 {print $2}'"}
	cmd := strings.Join(args, " ")
	out, _ := utils.Run(cmd)
	return strings.TrimSpace(out)
}

func checkSubpathProjQuotaEqual(projQuotaNamespacePath, projectID, blockHardLimitExpected, blockSoftLimitExpected string) (bool, error) {

	args := []string{NsenterCmd, "repquota", "-P -O csv", projQuotaNamespacePath}
	cmd := strings.Join(args, " ")
	out, err := utils.Run(cmd)
	if err != nil {
		return false, fmt.Errorf("failed to request namespace quota with error: %v", err)
	}
	blockSoftLimitActual, blockHardLimitActual, err := findBlockLimitByProjectID(out, projectID)
	if err != nil {
		return false, err
	}
	if blockHardLimitExpected == blockHardLimitActual && blockSoftLimitActual == blockHardLimitExpected {
		return true, nil
	}
	return false, fmt.Errorf("checkSubpathProjQuotaEqual:actualData:%s is not equals with setedData:%s, blockSoftLimit:%s, blockSoftlimit:%s", blockHardLimitExpected, blockHardLimitActual, blockSoftLimitExpected, blockSoftLimitActual)
}

// SetSubpathProjQuota ...
func SetSubpathProjQuota(ctx context.Context, projQuotaSubpath, blockHardlimit, blockSoftlimit string) (string, error) {
	if !utils.CheckParameterValidate([]string{projQuotaSubpath, blockHardlimit, blockSoftlimit}) {
		return "", fmt.Errorf("inputs illegal: %s, %s, %s", projQuotaSubpath, blockHardlimit, blockSoftlimit)
	}

	projectID := ConvertString2int(filepath.Base(projQuotaSubpath))
	rootPath := filepath.Dir(projQuotaSubpath)
	fsType := getFileSystemType(rootPath)
	switch fsType {
	case "ext3", "ext4":
		args := []string{NsenterCmd, "setquota", "-P", fmt.Sprintf("%s %s %s 0 0 %s", projectID, blockHardlimit, blockHardlimit, filepath.Dir(projQuotaSubpath))}
		cmd := strings.Join(args, " ")
		_, err := utils.Run(cmd)
		if err != nil {
			return "", fmt.Errorf("failed to set quota to subpath with error: %v", err)
		}
		ok, err := checkSubpathProjQuotaEqual(filepath.Dir(projQuotaSubpath), projectID, blockHardlimit, blockHardlimit)
		if ok {
			return "", nil
		}
		return "", err
	case "xfs":
		args := []string{NsenterCmd, "xfs_quota", "-x -c", fmt.Sprintf("'limit -p bhard=%sk %s' %s", blockHardlimit, projectID, rootPath)}
		cmd := strings.Join(args, " ")
		_, err := utils.Run(cmd)
		if err != nil {
			return "", fmt.Errorf("failed to set quota to subpath with error: %v", err)
		}
		return "", nil
	default:
		return "", fmt.Errorf("failed to analyse filesystem: %v type", fsType)
	}
}

// RemoveProjQuotaSubpath ...
func RemoveProjQuotaSubpath(ctx context.Context, quotaSubpath string) (string, error) {
	if !utils.CheckParameterValidate([]string{quotaSubpath}) {
		return "", fmt.Errorf("inputs illegal: %s", quotaSubpath)
	}
	if !utils.IsHostFileExist(quotaSubpath) {
		return "path already removed", nil
	}
	if err := utils.CheckQuotaPathValidate(KubeClient, quotaSubpath); err != nil {
		return "", err
	}

	args := []string{NsenterCmd, "rm", "-rf", quotaSubpath}
	cmd := strings.Join(args, " ")
	out, err := utils.Run(cmd)
	if err != nil {
		return "", fmt.Errorf("failed to remove proj quota subpath with error: %v", err)
	}
	return out, nil
}
