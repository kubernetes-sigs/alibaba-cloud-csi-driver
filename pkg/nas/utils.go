package nas

import (
	"fmt"
	"os"
	"strings"

	"github.com/AliyunContainerService/csi-plugin/pkg/utils"
)

func DoMount(nfsServer, nfsPath, nfsVers, mountOptions, mountPoint, volumeId string) error {
	if !utils.IsFileExisting(mountPoint) {
		CreateDest(mountPoint)
	}
	mntCmd := fmt.Sprintf("mount -t nfs %s:%s %s", nfsServer, nfsPath, mountPoint)
	if mountOptions != "" {
		mntCmd = fmt.Sprintf("mount -t nfs -o %s %s:%s %s", mountOptions, nfsServer, nfsPath, mountPoint)
	}
	_, err := utils.Run(mntCmd)
	if err != nil && nfsPath != "/" {
		if strings.Contains(err.Error(), "reason given by server: No such file or directory") || strings.Contains(err.Error(), "access denied by server while mounting") {
			if err := createNasSubDir(nfsServer, nfsPath, nfsVers, mountOptions, volumeId); err != nil {
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

func CheckNfsPathMounted(mountpoint, server, path string) bool {
	mntCmd := fmt.Sprintf("mount | grep %s | grep %s | grep %s | grep -v grep | wc -l", mountpoint, server, path)
	if out, err := utils.Run(mntCmd); err == nil && strings.TrimSpace(out) != "0" {
		return true
	}
	return false
}

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
	return nfsServer, nfsPath
}
