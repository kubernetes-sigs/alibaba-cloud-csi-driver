package utils

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"

	"github.com/container-storage-interface/spec/lib/go/csi"
	log "github.com/sirupsen/logrus"
)

// SetVolumeIOLimit config io limit for device
// readIOPS: 1000
// writeIOPS: 10000
// readBPS: 100K
// writeBPS: 1M
func SetVolumeIOLimit(devicePath string, req *csi.NodePublishVolumeRequest) error {
	readIOPS := req.VolumeContext["readIOPS"]
	writeIOPS := req.VolumeContext["writeIOPS"]
	readBPS := req.VolumeContext["readBPS"]
	writeBPS := req.VolumeContext["writeBPS"]

	// if no quota set, return;
	if readIOPS == "" && writeIOPS == "" && readBPS == "" && writeBPS == "" {
		return nil
	}

	// io limit parse
	readBPSInt, err := getBpsLimt(readBPS)
	if err != nil {
		log.Errorf("Volume(%s) Input Read BPS Limit format error: %s", req.VolumeId, err.Error())
		return err
	}
	writeBPSInt, err := getBpsLimt(writeBPS)
	if err != nil {
		log.Errorf("Volume(%s) Input Write BPS Limit format error: %s", req.VolumeId, err.Error())
		return err
	}
	readIOPSInt := 0
	if readIOPS != "" {
		readIOPSInt, err = strconv.Atoi(readIOPS)
		if err != nil {
			log.Errorf("Volume(%s) Input Read IOPS Limit format error: %s", req.VolumeId, err.Error())
			return err
		}
	}
	writeIOPSInt := 0
	if writeIOPS != "" {
		writeIOPSInt, err = strconv.Atoi(writeIOPS)
		if err != nil {
			log.Errorf("Volume(%s) Input Write IOPS Limit format error: %s", req.VolumeId, err.Error())
			return err
		}
	}

	// Get Device major/minor number
	majMinNum, err := getMajMinDevice(devicePath)
	if err != nil {
		return fmt.Errorf("Volume Cannot get major/minor device number for %s: %w", devicePath, err)
	}

	// Get pod uid
	podUID := req.VolumeContext["csi.storage.k8s.io/pod.uid"]
	if podUID == "" {
		log.Errorf("Volume(%s) Cannot get poduid and cannot set volume limit", req.VolumeId)
		return errors.New("Cannot get poduid and cannot set volume limit: " + req.VolumeId)
	}
	// /sys/fs/cgroup/blkio/kubepods.slice/kubepods-besteffort.slice/kubepods-besteffort-podaadcc749_6776_4933_990d_d50f260f5d46.slice/blkio.throttle.write_bps_device
	podUID = strings.ReplaceAll(podUID, "-", "_")
	podBlkIOPath := ""
	for _, p := range []string{
		fmt.Sprintf("/sys/fs/cgroup/blkio/kubepods.slice/kubepods-besteffort.slice/kubepods-besteffort-pod%s.slice", podUID),
		fmt.Sprintf("/sys/fs/cgroup/blkio/kubepods.slice/kubepods-burstable.slice/kubepods-burstable-pod%s.slice", podUID),
		fmt.Sprintf("/sys/fs/cgroup/blkio/kubepods.slice/kubepods-pod%s.slice", podUID),
	} {
		if IsHostFileExist(p) {
			podBlkIOPath = p
			break
		}
	}
	if podBlkIOPath == "" {
		log.Errorf("Volume(%s), pod blkio/cgroup path not found", req.VolumeId)
		return errors.New("pod blkio/cgroup path not found")
	}

	// io limit set to blkio limit files
	if readIOPSInt != 0 {
		err := writeIoLimit(majMinNum, podBlkIOPath, "blkio.throttle.read_iops_device", readIOPSInt)
		if err != nil {
			return err
		}
	}
	if writeIOPSInt != 0 {
		err := writeIoLimit(majMinNum, podBlkIOPath, "blkio.throttle.write_iops_device", writeIOPSInt)
		if err != nil {
			return err
		}
	}
	if readBPSInt != 0 {
		err := writeIoLimit(majMinNum, podBlkIOPath, "blkio.throttle.read_bps_device", readBPSInt)
		if err != nil {
			return err
		}
	}
	if writeBPSInt != 0 {
		err := writeIoLimit(majMinNum, podBlkIOPath, "blkio.throttle.write_bps_device", writeBPSInt)
		if err != nil {
			return err
		}
	}
	log.Infof("Seccessful Set Volume(%s) IO Limit: readIOPS(%d), writeIOPS(%d), readBPS(%d), writeBPS(%d)", req.VolumeId, readIOPSInt, writeIOPSInt, readBPSInt, writeBPSInt)
	return nil
}

func writeIoLimit(majMinNum, podBlkIOPath, ioFile string, ioLimit int) error {
	targetPath := filepath.Join(podBlkIOPath, ioFile)
	content := majMinNum + " " + strconv.Itoa(ioLimit)
	return WriteTrunc(targetPath, content)
}

func getBpsLimt(bpsLimt string) (int, error) {
	if bpsLimt == "" {
		return 0, nil
	}

	bpsLimt = strings.ToLower(bpsLimt)
	convertNumber := 1
	intBpsStr := bpsLimt
	if strings.HasSuffix(bpsLimt, "k") && len(bpsLimt) > 1 {
		convertNumber = 1024
		intBpsStr = strings.TrimSuffix(bpsLimt, "k")
	} else if strings.HasSuffix(bpsLimt, "m") && len(bpsLimt) > 1 {
		convertNumber = 1024 * 1024
		intBpsStr = strings.TrimSuffix(bpsLimt, "m")
	} else if strings.HasSuffix(bpsLimt, "g") && len(bpsLimt) > 1 {
		convertNumber = 1024 * 1024 * 1024
		intBpsStr = strings.TrimSuffix(bpsLimt, "g")
	}

	bpsIntValue, err := strconv.Atoi(intBpsStr)
	if err != nil {
		return 0, err
	}
	return bpsIntValue * convertNumber, nil
}

func getMajMinDevice(devicePath string) (string, error) {
	fileInfo, err := os.Stat(devicePath)
	if err != nil {
		return "", err
	}
	stat, ok := fileInfo.Sys().(*syscall.Stat_t)
	if !ok {
		return "", errors.New("unsupported platform")
	}
	maj := stat.Rdev / 256
	min := stat.Rdev % 256
	return fmt.Sprintf("%d:%d", maj, min), nil
}
