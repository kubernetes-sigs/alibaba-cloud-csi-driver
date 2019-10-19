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

package cpfs

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	log "github.com/sirupsen/logrus"
)

func createCpfsSubDir(cpfsOptions, cpfsServer, cpfsFileSystem, cpfsSubpath string, volumeId string) error {
	// step 1: create mount path
	cpfsTmpPath := filepath.Join(CPFS_TEMP_MNTPath, volumeId)
	if err := utils.CreateDest(cpfsTmpPath); err != nil {
		log.Infof("Create Cpfs temp Directory err: " + err.Error())
		return err
	}
	if utils.IsMounted(cpfsTmpPath) {
		utils.Umount(cpfsTmpPath)
	}

	// step 2: do mount
	mntCmd := fmt.Sprintf("mount -t lustre %s:/%s %s", cpfsServer, cpfsFileSystem, cpfsTmpPath)
	if cpfsOptions != "" {
		mntCmd = fmt.Sprintf("mount -t lustre -o %s %s:/%s %s", cpfsOptions, cpfsServer, cpfsFileSystem, cpfsTmpPath)
	}
	_, err := utils.Run(mntCmd)
	if err != nil {
		log.Errorf("Cpfs, Mount to temp directory fail: %s", err.Error())
		return err
	}
	subPath := path.Join(cpfsTmpPath, cpfsSubpath)
	if err := utils.CreateDest(subPath); err != nil {
		log.Infof("Cpfs, Create Sub Directory err: " + err.Error())
		return err
	}

	// step 3: umount after create
	utils.Umount(cpfsTmpPath)
	log.Infof("Create Sub Directory successful: %s", cpfsSubpath)
	return nil
}

func GetCpfsDetails(cpfsServersString string) (string, string, string) {
	cpfsServer, cpfsFileSystem, cpfsPath := "", "", ""
	cpfsServerList := strings.Split(cpfsServersString, ",")
	serverNum := len(cpfsServerList)

	if _, ok := storageClassServerPos[cpfsServersString]; !ok {
		storageClassServerPos[cpfsServersString] = 0
	}
	zoneIndex := storageClassServerPos[cpfsServersString] % serverNum
	selectedServer := cpfsServerList[zoneIndex]
	storageClassServerPos[cpfsServersString]++

	serverParts := strings.Split(selectedServer, ":/")
	if len(serverParts) == 2 {
		cpfsServer = serverParts[0]
		fsPath := serverParts[1]
		fsPathList := strings.Split(fsPath, "/")
		if len(fsPathList) == 1 {
			cpfsFileSystem = fsPathList[0]
			cpfsPath = "/"
		} else if len(fsPathList) == 2 {
			cpfsFileSystem = fsPathList[0]
			cpfsPath = fsPathList[1]
		} else {
			cpfsServer = ""
		}
	} else {
		cpfsServer = ""
		cpfsFileSystem = ""
		cpfsPath = ""
	}
	return cpfsServer, cpfsFileSystem, cpfsPath
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

//DoMount execute the mount command for cpfs dir
func DoMount(cpfsServer, cpfsFileSystem, cpfsPath, mountOptions, mountPoint, volumeId string) error {
	if !utils.IsFileExisting(mountPoint) {
		CreateDest(mountPoint)
	}
	mntCmd := fmt.Sprintf("mount -t lustre %s:/%s%s %s", cpfsServer, cpfsFileSystem, cpfsPath, mountPoint)
	if mountOptions != "" {
		mntCmd = fmt.Sprintf("mount -t lustre -o %s %s:/%s%s %s", mountOptions, cpfsServer, cpfsFileSystem, cpfsPath, mountPoint)
	}
	_, err := utils.Run(mntCmd)
	if err != nil && cpfsPath != "/" {
		if strings.Contains(err.Error(), "No such file or directory") {
			if err := createCpfsSubDir(mountOptions, cpfsServer, cpfsFileSystem, cpfsPath, volumeId); err != nil {
				return err
			}
			if _, err := utils.Run(mntCmd); err != nil {
				return err
			}
		} else {
			return err
		}
	} else if err != nil {
		return err
	}

	return nil
}
