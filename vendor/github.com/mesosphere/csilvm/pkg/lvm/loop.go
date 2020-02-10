package lvm

import (
	"io/ioutil"
	"os"

	"github.com/mesosphere/csilvm/pkg/cleanup"
	losetup "gopkg.in/freddierice/go-losetup.v1"
)

// This file abstracts operations regarding LVM on loopback devices.
// See http://www.anthonyldechiaro.com/blog/2010/12/19/lvm-loopback-how-to/

// LoopDevice represents a loop device such as `/dev/loop0` backed by a file.
type LoopDevice struct {
	lodev           losetup.Device
	backingFilePath string
}

// CreateLoopDevice returns a file-backed loop device.  The caller is
// responsible for calling `Close()` on the `*LoopDevice` when done
// with it.
//
// CreateLoopDevice may panic if an error occurs during error recovery.
func CreateLoopDevice(size uint64) (device *LoopDevice, err error) {
	var cleanup cleanup.Steps
	defer func() {
		if err != nil {
			cleanup.Unwind()
		}
	}()

	// Create a tempfile to use as the target of our loop device.
	file, err := ioutil.TempFile("", "test-dev")
	if err != nil {
		return nil, err
	}
	// If anything goes wrong, remove the tempfile.
	cleanup.Add(func() error { return os.Remove(file.Name()) })
	// Close the file as we're not going to manipulate its
	// contents manually.
	if err = file.Close(); err != nil {
		return nil, err
	}

	if err = os.Truncate(file.Name(), int64(size)); err != nil {
		return nil, err
	}

	// Attach a loop device
	const (
		offset = 0
		ro     = false
	)
	lodev, err := losetup.Attach(file.Name(), offset, ro)
	if err != nil {
		return nil, err
	}
	cleanup.Add(func() error { return lodev.Detach() })
	// https://www.howtogeek.com/howto/40702/how-to-manage-and-use-lvm-logical-volume-management-in-ubuntu/
	return &LoopDevice{lodev, file.Name()}, nil
}

func (d *LoopDevice) Path() string {
	return d.lodev.Path()
}

func (d *LoopDevice) String() string {
	return d.lodev.Path()
}

// Close detaches the loop device and removes the backing file.
func (d *LoopDevice) Close() error {
	if err := d.lodev.Detach(); err != nil {
		return err
	}
	return os.Remove(d.backingFilePath)
}
