//go:build !linux && !windows

package datacache

import (
	"errors"

	"k8s.io/klog/v2"
)

// The data cache is backed by Linux device-mapper and loop devices. On other
// platforms (the driver only builds for linux in production; this exists so the
// package compiles and its portable logic can be unit tested) the primitives
// are stubs: OpenDmControl reports device-mapper as unavailable, matching a
// Linux node without the dm subsystem.

const (
	dmNameLen         = 128
	dmStatusTableFlag = 1 << 4
)

var errDataCacheUnsupported = errors.New("data cache is only supported on Linux")

// DmControl is never constructed off Linux; OpenDmControl always reports
// device-mapper unavailable. The type exists only so the portable *DmControl
// fields and signatures compile.
type DmControl struct{}

func (c *DmControl) device(logger klog.Logger, volumeID string) dmDevice { panic("unsupported") }

func (c *DmControl) close() error { panic("unsupported") }

func OpenDmControl() (*DmControl, error) { return nil, nil }

func allocCacheFile(logger klog.Logger, path string, size int64) (string, int, error) {
	return "", 0, errDataCacheUnsupported
}
