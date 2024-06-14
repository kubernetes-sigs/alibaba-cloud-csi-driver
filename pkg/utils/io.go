package utils

import (
	"errors"
	"fmt"
	"math"
	"math/bits"
	"strconv"
	"strings"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils/cgroup"
	utilsio "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils/io"
	log "github.com/sirupsen/logrus"
	"golang.org/x/sys/unix"
)

// CGROUPFS_MOUNT_PATH is the path to the host cgroupfs mounted in the container
const CGROUPFS_MOUNT_PATH = "/host/sys/fs/cgroup"

type PodCGroup struct {
	slice cgroup.PodCGroupSlice
	io    cgroup.IO
	dev   utilsio.DevTmpFS
}

func NewPodCGroup() (*PodCGroup, error) {
	slice, io, err := cgroup.DetectCgroupVersion(CGROUPFS_MOUNT_PATH)
	if err != nil {
		return nil, err
	}

	return &PodCGroup{
		slice: slice,
		io:    io,
		dev:   utilsio.RealDevTmpFS{},
	}, nil
}

func (cg *PodCGroup) ApplyConfig(devicePath string, req *csi.NodePublishVolumeRequest) error {
	return cg.setVolumeIOLimit(devicePath, req)
}

// parseIOLimits config io limit for device
// readIOPS: 1000
// writeIOPS: 10000
// readBPS: 100K
// writeBPS: 1M
func parseIOLimits(ctx map[string]string) (*cgroup.IOLimits, error) {
	readIOPS := ctx["readIOPS"]
	writeIOPS := ctx["writeIOPS"]
	readBPS := ctx["readBPS"]
	writeBPS := ctx["writeBPS"]

	// if no quota set, return;
	if readIOPS == "" && writeIOPS == "" && readBPS == "" && writeBPS == "" {
		return nil, nil
	}

	// io limit parse
	limit := cgroup.MaxIOLimits()
	var err error
	limit.BPS.Read, err = getBpsLimit(readBPS)
	if err != nil {
		return nil, fmt.Errorf("failed to parse readBPS: %w", err)
	}

	limit.BPS.Write, err = getBpsLimit(writeBPS)
	if err != nil {
		return nil, fmt.Errorf("failed to parse writeBPS: %w", err)
	}

	limit.IOPS.Read, err = getIOPSLimit(readIOPS)
	if err != nil {
		return nil, fmt.Errorf("failed to parse readIOPS: %w", err)
	}

	limit.IOPS.Write, err = getIOPSLimit(writeIOPS)
	if err != nil {
		return nil, fmt.Errorf("failed to parse writeIOPS: %w", err)
	}
	return &limit, nil
}

func (cg *PodCGroup) setVolumeIOLimit(devicePath string, req *csi.NodePublishVolumeRequest) (err error) {
	limits, err := parseIOLimits(req.VolumeContext)
	if err != nil {
		return err
	}
	if limits == nil {
		return nil
	}

	// Get Device major/minor number
	maj, min, err := cg.dev.DevFor(devicePath)
	if err != nil {
		return fmt.Errorf("Volume Cannot get major/minor device number for %s: %w", devicePath, err)
	}

	// Get pod uid
	podUID := req.VolumeContext["csi.storage.k8s.io/pod.uid"]
	if podUID == "" {
		log.Errorf("Volume(%s) Cannot get poduid and cannot set volume limit", req.VolumeId)
		return errors.New("Cannot get poduid and cannot set volume limit: " + req.VolumeId)
	}

	fd, err := cg.slice.FindPodDir(podUID)
	if err != nil {
		return err
	}
	defer func() {
		if err := unix.Close(fd); err != nil {
			log.Errorf("Volume(%s) close fd(%d) failed: %v", req.VolumeId, fd, err)
		}
	}()

	if err := cg.io.SetLimits(fd, maj, min, limits); err != nil {
		return err
	}
	log.Infof("Successfully Set Volume(%s) IO Limit: %+v", req.VolumeId, limits)
	return nil
}

func getBpsLimit(limit string) (uint64, error) {
	if limit == "" {
		return math.MaxUint64, nil
	}

	limit = strings.ToLower(limit)
	var convertNumber uint64 = 1
	switch limit[len(limit)-1] {
	case 'k', 'K':
		convertNumber = 1024
	case 'm', 'M':
		convertNumber = 1024 * 1024
	case 'g', 'G':
		convertNumber = 1024 * 1024 * 1024
	}

	intStr := limit
	if convertNumber != 1 {
		intStr = limit[:len(limit)-1]
	}

	intValue, err := strconv.ParseUint(intStr, 10, 64)
	if err != nil {
		return 0, err
	}
	hi, intValue := bits.Mul64(intValue, convertNumber)
	if hi > 0 {
		return 0, fmt.Errorf("%s %w", limit, strconv.ErrRange)
	}
	return intValue, nil
}

func getIOPSLimit(limit string) (uint32, error) {
	if limit == "" {
		return math.MaxUint32, nil
	}

	intValue, err := strconv.ParseUint(limit, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint32(intValue), nil
}
