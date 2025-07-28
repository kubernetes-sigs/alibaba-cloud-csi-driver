package alinas

import (
	"context"
	"io"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy/server"
	mounter "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/klog/v2"
	"k8s.io/mount-utils"
)

const (
	fstypeCpfsNfs = "cpfs-nfs"
	fstypeAlinas  = "alinas"
)

func init() {
	server.RegisterDriver(&Driver{mounter: mount.New("")})
}

type Driver struct {
	mounter mount.Interface
}

func (h *Driver) Name() string {
	return "alinas"
}

func (h *Driver) Fstypes() []string {
	return []string{fstypeAlinas, fstypeCpfsNfs}
}

func (h *Driver) Mount(ctx context.Context, req *proxy.MountRequest, fuseFd int) error {
	klog.InfoS("Mounting", "fstype", req.Fstype, "source", req.Source, "target", req.Target, "options", req.Options)
	options := append(req.Options, "no_start_watchdog")
	if req.Fstype == fstypeAlinas {
		// options = append(options, "no_atomic_move", "auto_fallback_nfs")
		options = append(options, "no_atomic_move")
		options = addAutoFallbackNFSMountOptions(options)
	}

	return h.mounter.Mount(req.Source, req.Target, req.Fstype, options)
}

func (h *Driver) Init() {
	go runCommandForever("aliyun-alinas-mount-watchdog")
	go runCommandForever("aliyun-cpfs-mount-watchdog")
	setupDefaultConfigs()
}

func (h *Driver) Terminate() {}

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
		for _, option := range mounter.SplitMountOptions(options) {
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
	if isEFC && !isVSC {
		mountOptions = append(mountOptions, "auto_fallback_nfs")
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
