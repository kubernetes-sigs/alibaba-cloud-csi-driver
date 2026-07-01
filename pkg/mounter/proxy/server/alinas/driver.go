package alinas

import (
	"context"
	"io"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/interceptors"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy/server"
	mounterutils "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/klog/v2"
	"k8s.io/mount-utils"
)

const (
	fstypeCpfsNfs = "cpfs-nfs"
	fstypeAlinas  = "alinas"
)

func init() {
	server.RegisterDriver(NewDriver())
}

func NewDriver() *Driver {
	driver := &Driver{}
	driver.Mounter = mounter.NewForMounter(
		&extendedMounter{driver: driver, Interface: mount.New("")},
		interceptors.AlinasSecretInterceptor,
	)
	return driver
}

type Driver struct {
	mounter.Mounter
	targets       sync.Map
	ResetFlagPath string
}

func (h *Driver) Name() string {
	return "alinas"
}

func (h *Driver) Fstypes() []string {
	return []string{fstypeAlinas, fstypeCpfsNfs}
}

func (h *Driver) Mount(ctx context.Context, req *proxy.MountRequest) error {
	return h.ExtendedMount(ctx, &mounter.MountOperation{
		Source:      req.Source,
		Target:      req.Target,
		FsType:      req.Fstype,
		Options:     req.Options,
		Secrets:     req.Secrets,
		MetricsPath: req.MetricsPath,
		VolumeID:    req.VolumeID,
	})
}

func (h *Driver) Init() {
	go runCommandForever("aliyun-alinas-mount-watchdog")
	go runCommandForever("aliyun-cpfs-mount-watchdog")
	setupDefaultConfigs()
}

// defaultResetFlagPath is the path to the reset flag file written by envd.
// When this file exists at termination time, the driver will unmount all
// tracked NAS mount points before exiting.
var defaultResetFlagPath = "/etc/aliyun/alinas/reset"

func (h *Driver) resetFlagPath() string {
	if h.ResetFlagPath != "" {
		return h.ResetFlagPath
	}
	return defaultResetFlagPath
}

func (h *Driver) Terminate() {
	if !server.CleanupNASMountsOnExit() {
		return
	}
	if !h.shouldCleanup() {
		klog.InfoS("Reset flag not found, skipping NAS mount cleanup", "path", h.resetFlagPath())
		return
	}
	h.targets.Range(func(key, _ any) bool {
		target := key.(string)
		klog.InfoS("Unmounting NAS mount point on exit", "target", target)
		if err := h.Unmount(target); err != nil {
			klog.ErrorS(err, "Failed to unmount NAS mount point", "target", target)
		}
		return true
	})
}

func (h *Driver) shouldCleanup() bool {
	_, err := os.Stat(h.resetFlagPath())
	return err == nil
}

// ApplyOptionDefaults applies driver-specific option defaults.
// alinas does not apply any environment-detected defaults.
func (h *Driver) ApplyOptionDefaults(options []string) []string {
	return options
}

func runCommandForever(command string, args ...string) {
	wait.Forever(func() {
		klog.InfoS("Starting", "command", command)
		cmd := exec.Command(command, args...)
		cmd.SysProcAttr = &syscall.SysProcAttr{
			Setpgid: true,
		}
		err := cmd.Run()
		if err != nil {
			klog.ErrorS(err, "Exited", "command", command)
		}
	}, time.Second)
}

// addAutoFallbackNFSMountOptions adds auto_fallback_nfs mount option when using efc
func addAutoFallbackNFSMountOptions(mountOptions []string) []string {
	isEFC := false
	isVSC := false
	for _, options := range mountOptions {
		for _, option := range mounterutils.SplitMountOptions(options) {
			if option == "" {
				continue
			}
			key, value, _ := strings.Cut(option, "=")
			switch key {
			case "efc":
				isEFC = true
			case "net":
				isVSC = value == "vsc"
			}
		}
	}
	if isEFC {
		mountOptions = append(mountOptions, "no_add_cgroup")
		if !isVSC {
			mountOptions = append(mountOptions, "auto_fallback_nfs")
		}
	}
	return mountOptions
}

const (
	defaultConfigDir = "/etc/aliyun-defaults"
	configDir        = "/etc/aliyun"
)

func setupDefaultConfigs() {
	for _, name := range []string{"cpfs", "alinas"} {
		srcDir := filepath.Join(defaultConfigDir, name)
		dstDir := filepath.Join(configDir, name)
		if err := filepath.WalkDir(srcDir, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}

			if d.IsDir() {
				return nil
			}

			relPath, err := filepath.Rel(srcDir, path)
			if err != nil {
				return err
			}

			dstPath := filepath.Join(dstDir, relPath)

			if _, err := os.Stat(dstPath); err == nil {
				// File already exists, skip
				return nil
			} else if !os.IsNotExist(err) {
				return err
			}

			klog.InfoS("Copying default config file", "path", dstPath)
			return copyFile(path, dstPath)
		}); err != nil {
			panic(err)
		}
	}
}

func copyFile(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}

	return dstFile.Sync()
}

type extendedMounter struct {
	driver *Driver
	mount.Interface
}

var _ mounter.Mounter = &extendedMounter{}

func (m *extendedMounter) ExtendedMount(ctx context.Context, op *mounter.MountOperation) error {
	klog.FromContext(ctx).Info("Mounting", "fstype", op.FsType, "source", op.Source, "target", op.Target, "options", op.Options)
	op.Options = append(op.Options, "no_start_watchdog")
	if op.FsType == fstypeAlinas {
		// options = append(options, "no_atomic_move", "auto_fallback_nfs")
		op.Options = append(op.Options, "no_atomic_move")
		op.Options = addAutoFallbackNFSMountOptions(op.Options)
	}
	err := m.Mount(op.Source, op.Target, op.FsType, op.Options)
	if err == nil {
		m.driver.targets.Store(op.Target, struct{}{})
	}
	return err
}

func (m *extendedMounter) Unmount(target string) error {
	err := m.Interface.Unmount(target)
	if err == nil {
		m.driver.targets.Delete(target)
	}
	return err
}
