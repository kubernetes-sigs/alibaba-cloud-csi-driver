package cgroup

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	utilsio "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils/io"
	"golang.org/x/sys/unix"
	"k8s.io/klog/v2"
)

type CgroupDriver int

const (
	Systemd CgroupDriver = iota
	CgroupFS
)

var (
	probePath = [...]string{
		Systemd:  "kubepods.slice",
		CgroupFS: "kubepods",
	}
	podTemplate = [...][]string{
		Systemd: {
			"%s/kubepods-besteffort.slice/kubepods-besteffort-pod%s.slice",
			"%s/kubepods-burstable.slice/kubepods-burstable-pod%s.slice",
			"%s/kubepods-pod%s.slice",
		},
		CgroupFS: {
			"%s/besteffort/pod%s",
			"%s/burstable/pod%s",
			"%s/pod%s",
		},
	}
)

type PodCGroupSlice struct {
	KubepodsSlicePath string
	CgroupDriver      CgroupDriver
}

func (cg PodCGroupSlice) FindPodDir(podUID string) (int, error) {
	if cg.KubepodsSlicePath == "" {
		return -1, errors.New("no cgroup detected")
	}

	//  systemd: /sys/fs/cgroup/blkio/kubepods.slice/kubepods-besteffort.slice/kubepods-besteffort-podaadcc749_6776_4933_990d_d50f260f5d46.slice/blkio.throttle.write_bps_device
	// cgroupfs: /sys/fs/cgroup/blkio/kubepods/burstable/podaa704696-c39c-4f18-bace-53f58a2db1cc/blkio.throttle.write_bps_device
	if cg.CgroupDriver == Systemd {
		podUID = strings.ReplaceAll(podUID, "-", "_")
	}
	for _, p := range podTemplate[cg.CgroupDriver] {
		p = fmt.Sprintf(p, cg.KubepodsSlicePath, podUID)
		fd, err := unix.Open(p, utilsio.O_PATH, 0)
		if err == nil {
			return fd, nil
		}
		if !errors.Is(err, fs.ErrNotExist) {
			return -1, fmt.Errorf("failed to open CGroup path %q: %w", p, err)
		}
	}
	return -1, errors.New("pod CGroup path not found")
}

// DetectCgroupVersion detects the cgroup version of the cgroup mounted at cgroupMountPath.
// It only cares about kubepods.slice
//
// Takes the path to the cgroupfs mount point as an argument. Usually /sys/fs/cgroup.
// Returns the path to kubepods.slice, the IO implementation for the detected version.
func DetectCgroupVersion(cgroupfsMountPath string) (PodCGroupSlice, IO, error) {
	for driver, p := range probePath {
		v2Path := filepath.Join(cgroupfsMountPath, p)
		_, err := os.Stat(v2Path + "/io.max")
		if err == nil {
			klog.Infof("Detected cgroup v2 at %s", v2Path)
			return PodCGroupSlice{v2Path, CgroupDriver(driver)}, V2{}, nil
		}
		if !errors.Is(err, os.ErrNotExist) {
			return PodCGroupSlice{}, nil, err
		}

		v1Path := filepath.Join(cgroupfsMountPath, "blkio", p)
		_, err = os.Stat(v1Path + "/blkio.throttle.read_bps_device")
		if err == nil {
			klog.Infof("Detected cgroup v1 at %s", v1Path)
			return PodCGroupSlice{v1Path, CgroupDriver(driver)}, V1{}, nil
		}
		if !errors.Is(err, os.ErrNotExist) {
			return PodCGroupSlice{}, nil, err
		}
	}

	klog.Warningf("No cgroup detected at %s. IO limit will not work", cgroupfsMountPath)
	return PodCGroupSlice{}, nil, nil
}
