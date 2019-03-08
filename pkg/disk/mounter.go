package disk

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang/glog"
	"os"
	"os/exec"
	"strings"
)

type findmntResponse struct {
	FileSystems []fileSystem `json:"filesystems"`
}

type fileSystem struct {
	Target      string `json:"target"`
	Propagation string `json:"propagation"`
	FsType      string `json:"fstype"`
	Options     string `json:"options"`
}

// Mounter is responsible for formatting and mounting volumes
type Mounter interface {
	// If the folder doesn't exist, it will call 'mkdir -p'
	EnsureFolder(target string) error
	// If the block doesn't exist, create it
	EnsureBlock(target string) error
	// Format formats the source with the given filesystem type
	Format(source, fsType string) error

	// Mount mounts source to target with the given fstype and options.
	Mount(source, target, fsType string, options ...string) error

	// Mount mounts source to target for block file.
	MountBlock(source, target string, options ...string) error
	// Unmount unmounts the given target
	Unmount(target string) error

	// IsFormatted checks whether the source device is formatted or not. It
	// returns true if the source device is already formatted.
	IsFormatted(source string) (bool, error)

	// IsMounted checks whether the target path is a correct mount (i.e:
	// propagated). It returns true if it's mounted. An error is returned in
	// case of system errors or if it's mounted incorrectly.
	IsMounted(target string) (bool, error)
}

// TODO(arslan): this is Linux only for now. Refactor this into a package with
// architecture specific code in the future, such as mounter_darwin.go,
// mounter_linux.go, etc..
type mounter struct {
}

// NewMounter returns a new mounter instance
func NewMounter() Mounter {
	return &mounter{}
}
func (m *mounter) EnsureFolder(target string) error {
	mdkirCmd := "mkdir"
	_, err := exec.LookPath(mdkirCmd)
	if err != nil {
		if err == exec.ErrNotFound {
			return fmt.Errorf("%q executable not found in $PATH", mdkirCmd)
		}
		return err
	}

	mkdirArgs := []string{"-p", target}
	glog.V(4).Infof("mkdir for folder, the command is %s %v", mdkirCmd, mkdirArgs)
	_, err = exec.Command(mdkirCmd, mkdirArgs...).CombinedOutput()
	if err != nil {
		return fmt.Errorf("mkdir for folder error: %v", err)
	}
	return nil
}

func (m *mounter) EnsureBlock(target string) error {
	fi, err := os.Lstat(target)
	if err != nil && !os.IsNotExist(err) {
		return err
	}
	if err == nil && fi.IsDir() {
		os.Remove(target)
	}
	targetPathFile, err := os.OpenFile(target, os.O_CREATE|os.O_RDWR, 0750)
	if err != nil {
		glog.V(4).Infof("Failed to create block:%s with error: %v", target, err)
		return fmt.Errorf("create block error: %v", err)
	}
	if err := targetPathFile.Close(); err != nil {
		glog.V(4).Infof("Failed to close targetPath:%s with error: %v", target, err)
		return fmt.Errorf("close block error: %v", err)
	}
	return nil
}

func (m *mounter) Format(source, fsType string) error {
	mkfsCmd := fmt.Sprintf("mkfs.%s", fsType)

	_, err := exec.LookPath(mkfsCmd)
	if err != nil {
		if err == exec.ErrNotFound {
			return fmt.Errorf("%q executable not found in $PATH", mkfsCmd)
		}
		return err
	}

	mkfsArgs := []string{}
	if fsType == "" {
		return errors.New("fs type is not specified for formatting the volume")
	}
	if source == "" {
		return errors.New("source is not specified for formatting the volume")
	}
	mkfsArgs = append(mkfsArgs, source)
	if fsType == "ext4" || fsType == "ext3" {
		mkfsArgs = []string{"-F", source}
	}

	glog.V(4).Infof("Format %s with fsType %s, the command is %s %v", source, fsType, mkfsCmd, mkfsArgs)
	out, err := exec.Command(mkfsCmd, mkfsArgs...).CombinedOutput()
	if err != nil {
		return fmt.Errorf("formatting disk failed: %v cmd: '%s %s' output: %q",
			err, mkfsCmd, strings.Join(mkfsArgs, " "), string(out))
	}

	return nil
}

func (m *mounter) MountBlock(source, target string, opts ...string) error {
	mountCmd := "mount"
	mountArgs := []string{}

	if source == "" {
		return errors.New("source is not specified for mounting the volume")
	}
	if target == "" {
		return errors.New("target is not specified for mounting the volume")
	}

	if len(opts) > 0 {
		mountArgs = append(mountArgs, "-o", strings.Join(opts, ","))
	}
	mountArgs = append(mountArgs, source)
	mountArgs = append(mountArgs, target)
	// create target, os.Mkdirall is noop if it exists
	_, err := os.Create(target)
	if err != nil {
		return err
	}

	glog.V(4).Infof("Mount %s to %s, the command is %s %v", source, target, mountCmd, mountArgs)
	out, err := exec.Command(mountCmd, mountArgs...).CombinedOutput()
	if err != nil {
		return fmt.Errorf("mounting failed: %v cmd: '%s %s' output: %q",
			err, mountCmd, strings.Join(mountArgs, " "), string(out))
	}
	return nil
}

func (m *mounter) Mount(source, target, fsType string, opts ...string) error {
	mountCmd := "mount"
	mountArgs := []string{}

	if fsType == "" {
		return errors.New("fs type is not specified for mounting the volume")
	}

	if source == "" {
		return errors.New("source is not specified for mounting the volume")
	}

	if target == "" {
		return errors.New("target is not specified for mounting the volume")
	}

	mountArgs = append(mountArgs, "-t", fsType)

	if len(opts) > 0 {
		mountArgs = append(mountArgs, "-o", strings.Join(opts, ","))
	}

	mountArgs = append(mountArgs, source)
	mountArgs = append(mountArgs, target)

	// create target, os.Mkdirall is noop if it exists
	err := os.MkdirAll(target, 0750)
	if err != nil {
		return err
	}

	glog.V(4).Infof("Mount %s to %s with fsType %s, the command is %s %v", source, target, fsType, mountCmd, mountArgs)

	out, err := exec.Command(mountCmd, mountArgs...).CombinedOutput()
	if err != nil {
		return fmt.Errorf("mounting failed: %v cmd: '%s %s' output: %q",
			err, mountCmd, strings.Join(mountArgs, " "), string(out))
	}

	return nil
}

func (m *mounter) Unmount(target string) error {
	umountCmd := "umount"
	if target == "" {
		return errors.New("target is not specified for unmounting the volume")
	}

	umountArgs := []string{target}

	glog.V(4).Infof("Unmount %s, the command is %s %v", target, umountCmd, umountArgs)

	out, err := exec.Command(umountCmd, umountArgs...).CombinedOutput()
	if err != nil {
		return fmt.Errorf("unmounting failed: %v cmd: '%s %s' output: %q",
			err, umountCmd, target, string(out))
	}

	return nil
}

func (m *mounter) IsFormatted(source string) (bool, error) {
	if source == "" {
		return false, errors.New("source is not specified")
	}

	fileCmd := "file"
	_, err := exec.LookPath(fileCmd)
	if err != nil {
		if err == exec.ErrNotFound {
			return false, fmt.Errorf("%q executable not found in $PATH", fileCmd)
		}
		return false, err
	}

	args := []string{"-sL", source}

	glog.V(4).Infof("IsFormatted: %s, the command is %s %v", source, fileCmd, args)
	out, err := exec.Command(fileCmd, args...).CombinedOutput()
	if err != nil {
		return false, fmt.Errorf("checking formatting failed: %v cmd: %q output: %q",
			err, fileCmd, string(out))
	}

	output := strings.TrimPrefix(string(out), fmt.Sprintf("%s:", source))
	if strings.TrimSpace(output) == "data" {
		return false, nil
	}

	return true, nil
}

func (m *mounter) IsMounted(target string) (bool, error) {
	if target == "" {
		return false, errors.New("target is not specified for checking the mount")
	}

	findmntCmd := "findmnt"
	_, err := exec.LookPath(findmntCmd)
	if err != nil {
		if err == exec.ErrNotFound {
			return false, fmt.Errorf("%q executable not found in $PATH", findmntCmd)
		}
		return false, err
	}

	findmntArgs := []string{"-o", "TARGET,PROPAGATION,FSTYPE,OPTIONS", "-M", target, "-J"}

	glog.V(4).Infof("IsMounted: %s, the command is %s %v", target, findmntCmd, findmntArgs)
	out, err := exec.Command(findmntCmd, findmntArgs...).CombinedOutput()
	if err != nil {
		// findmnt exits with non zero exit status if it couldn't find anything
		if strings.TrimSpace(string(out)) == "" {
			return false, nil
		}

		return false, fmt.Errorf("checking mounted failed: %v cmd: %q output: %q",
			err, findmntCmd, string(out))
	}

	// no response means there is no mount
	if string(out) == "" {
		return false, nil
	}

	var resp *findmntResponse
	err = json.Unmarshal(out, &resp)
	if err != nil {
		return false, fmt.Errorf("couldn't unmarshal data: %q: %s", string(out), err)
	}

	targetFound := false
	for _, fs := range resp.FileSystems {
		// check if the mount is propagated correctly. It should be set to shared.
		// if fs.Propagation != "shared" {
		// 	return true, fmt.Errorf("mount propagation for target %q is not enabled", target)
		// }

		// the mountpoint should match as well
		if fs.Target == target {
			targetFound = true
		}
	}

	return targetFound, nil
}