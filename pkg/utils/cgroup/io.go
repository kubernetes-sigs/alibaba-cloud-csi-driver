package cgroup

import (
	"bytes"
	"fmt"
	"math"
	"strconv"

	utilsio "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils/io"
)

type IOLimits struct {
	BPS struct {
		Read  uint64
		Write uint64
	}
	IOPS struct {
		Read  uint32
		Write uint32
	}
}

func MaxIOLimits() IOLimits {
	l := IOLimits{}
	l.BPS.Read = math.MaxUint64
	l.BPS.Write = math.MaxUint64
	l.IOPS.Read = math.MaxUint32
	l.IOPS.Write = math.MaxUint32
	return l
}

type IO interface {
	SetLimits(dirFd int, major, minor uint32, limits *IOLimits) error
}

type V1 struct{}

func (V1) SetLimits(dirFd int, major, minor uint32, limits *IOLimits) error {
	devPrefix := fmt.Sprintf("%d:%d ", major, minor)
	type f struct {
		path  string
		limit uint64
	}
	pendingFiles := make([]f, 0, 4)
	if limits.BPS.Read < math.MaxUint64 {
		pendingFiles = append(pendingFiles, f{path: "blkio.throttle.read_bps_device", limit: limits.BPS.Read})
	}
	if limits.BPS.Write < math.MaxUint64 {
		pendingFiles = append(pendingFiles, f{path: "blkio.throttle.write_bps_device", limit: limits.BPS.Write})
	}
	if limits.IOPS.Read < math.MaxUint32 {
		pendingFiles = append(pendingFiles, f{path: "blkio.throttle.read_iops_device", limit: uint64(limits.IOPS.Read)})
	}
	if limits.IOPS.Write < math.MaxUint32 {
		pendingFiles = append(pendingFiles, f{path: "blkio.throttle.write_iops_device", limit: uint64(limits.IOPS.Write)})
	}
	for _, f := range pendingFiles {
		content := devPrefix + strconv.FormatUint(f.limit, 10) + "\n"
		if err := utilsio.WriteTrunc(dirFd, f.path, []byte(content)); err != nil {
			return fmt.Errorf("failed to write %s: %w", f.path, err)
		}
	}
	return nil
}

type V2 struct{}

func (V2) SetLimits(dirFd int, major, minor uint32, limits *IOLimits) error {
	content := bytes.Buffer{}
	content.WriteString(fmt.Sprintf("%d:%d", major, minor))
	if limits.BPS.Read < math.MaxUint64 {
		content.WriteString(fmt.Sprintf(" rbps=%d", limits.BPS.Read))
	} else {
		content.WriteString(" rbps=max")
	}
	if limits.BPS.Write < math.MaxUint64 {
		content.WriteString(fmt.Sprintf(" wbps=%d", limits.BPS.Write))
	} else {
		content.WriteString(" wbps=max")
	}
	if limits.IOPS.Read < math.MaxUint32 {
		content.WriteString(fmt.Sprintf(" riops=%d", limits.IOPS.Read))
	} else {
		content.WriteString(" riops=max")
	}
	if limits.IOPS.Write < math.MaxUint32 {
		content.WriteString(fmt.Sprintf(" wiops=%d", limits.IOPS.Write))
	} else {
		content.WriteString(" wiops=max")
	}
	content.WriteString("\n")
	err := utilsio.WriteTrunc(dirFd, "io.max", content.Bytes())
	if err != nil {
		return fmt.Errorf("failed to write io.max: %w", err)
	}
	return nil
}
