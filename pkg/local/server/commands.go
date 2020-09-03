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
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/local/lib"
	"golang.org/x/net/context"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
)

const (
	// NsenterCmd is the nsenter command
	NsenterCmd = "/nsenter --mount=/proc/1/ns/mnt --ipc=/proc/1/ns/ipc --net=/proc/1/ns/net --uts=/proc/1/ns/uts "
)

// ListLV lists lvm volumes
func ListLV(listspec string) ([]*lib.LV, error) {
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
	vgs := make([]*lib.VG, len(outLines))
	for i, line := range outLines {
		line = strings.TrimSpace(line)
		vg, err := lib.ParseVG(line)
		if err != nil {
			return nil, err
		}
		vgs[i] = vg
	}
	return vgs, nil
}

// CreateVG create volume group
func CreateVG(ctx context.Context, name string, physicalVolume string, tags []string) (string, error) {
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
	regions, err := lib.GetRegions()
	if err != nil {
		return nil, err
	}

	namespaces := []*lib.NameSpace{}
	for _, region := range regions.Regions {
		for _, ns := range region.Namespaces {
			namespaces = append(namespaces, ns.ToProto())
		}
	}
	return namespaces, nil
}
