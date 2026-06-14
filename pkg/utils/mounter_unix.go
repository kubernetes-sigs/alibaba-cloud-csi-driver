//go:build unix

package utils

import (
	"fmt"
	"slices"

	"k8s.io/klog/v2"
	k8smount "k8s.io/mount-utils"
	utilexec "k8s.io/utils/exec"
)

// formatAndMount uses unix utils to format and mount the given disk
func FormatAndMount(diskMounter *k8smount.SafeFormatAndMount, source string, target string, fstype string, mkfsOptions []string, mountOptions []string, omitFsCheck bool) error {
	klog.Infof("formatAndMount: mount options : %+v", mountOptions)
	readOnly := slices.Contains(mountOptions, "ro")

	// check device fs
	mountOptions = append(mountOptions, "defaults")
	if !readOnly || !omitFsCheck {
		// Run fsck on the disk to fix repairable issues, only do this for volumes requested as rw.
		args := []string{"-a", source}

		out, err := diskMounter.Exec.Command("fsck", args...).CombinedOutput()
		if err != nil {
			ee, isExitError := err.(utilexec.ExitError)
			switch {
			case err == utilexec.ErrExecutableNotFound:
				klog.Warningf("'fsck' not found on system; continuing mount without running 'fsck'.")
			case isExitError && ee.ExitStatus() == fsckErrorsCorrected:
				klog.Infof("Device %s has errors which were corrected by fsck.", source)
			case isExitError && ee.ExitStatus() == fsckErrorsUncorrected:
				return fmt.Errorf("'fsck' found errors on device %s but could not correct them: %s", source, string(out))
			case isExitError && ee.ExitStatus() > fsckErrorsUncorrected:
			}
		}
	}

	// Try to mount the disk
	mountErr := diskMounter.Interface.Mount(source, target, fstype, mountOptions)
	if mountErr != nil {
		// Mount failed. This indicates either that the disk is unformatted or
		// it contains an unexpected filesystem.
		existingFormat, err := diskMounter.GetDiskFormat(source)
		if err != nil {
			return err
		}
		if existingFormat == "" {
			return FormatNewDisk(readOnly, source, fstype, target, mkfsOptions, mountOptions, diskMounter)
		}
		// Disk is already formatted and failed to mount
		if len(fstype) == 0 || fstype == existingFormat {
			// This is mount error
			return mountErr
		}
		// Detect if an encrypted disk is empty disk, since atari type partition is detected by blkid.
		if existingFormat == "unknown data, probably partitions" {
			klog.Infof("FormatAndMount: enter special partition logics")
			fsType, ptType, _ := GetDiskFStypePTtype(source)
			if fsType == "" && ptType == "atari" {
				return FormatNewDisk(readOnly, source, fstype, target, mkfsOptions, mountOptions, diskMounter)
			}
		}
		// Block device is formatted with unexpected filesystem, let the user know
		return fmt.Errorf("failed to mount the volume as %q, it already contains %s. Mount error: %v", fstype, existingFormat, mountErr)
	}

	return mountErr
}
