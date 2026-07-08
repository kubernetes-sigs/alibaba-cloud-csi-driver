package alinas

import (
	"context"
	"fmt"
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
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/jwtauth"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy/server"
	mounterutils "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
	"golang.org/x/sys/unix"
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
		interceptors.AlinasJWTAuthInterceptor,
	)
	return driver
}

type Driver struct {
	mounter.Mounter
	targets       sync.Map
	ResetFlagPath string
	// ConfigDir overrides the base config directory (/etc/aliyun). Used in tests.
	ConfigDir string
}

func (h *Driver) Name() string {
	return "alinas"
}

func (h *Driver) Fstypes() []string {
	return []string{fstypeAlinas, fstypeCpfsNfs}
}

var _ server.Unmounter = (*Driver)(nil)

// Unmount implements server.Unmounter. It unmounts target only if this driver
// mounted it (tracked in h.targets at mount time), returning owned=false when it
// has no record of target so the dispatcher can fall back appropriately.
//
// Routing by ownership (not by kernel fstype) is required: alinas AccessPoint
// mounts are performed with fstype "alinas" but appear as "nfs" in the kernel
// mount table, so any fstype-based decision would fail to route their unmounts
// through the broker and the umount.nfs RPC to tcp 12049 would be dropped by the
// csi_mount_proxy nftables rule (cgroup != 0), blocking ~3s.
func (h *Driver) Unmount(target string) (owned bool, err error) {
	if _, ok := h.targets.Load(target); !ok {
		return false, nil
	}
	// h.Mounter is the extendedMounter, whose Unmount also deletes the target
	// from h.targets and stops the jwtauth refresher on success.
	if err := h.Mounter.Unmount(target); err != nil {
		return true, err
	}
	return true, nil
}

func (h *Driver) Mount(ctx context.Context, req *proxy.MountRequest, fuseFd int) error {
	if fuseFd > 0 {
		// alinas does not support fd-passing; close the received fd to prevent leak
		// and fall back to normal mount.
		_ = unix.Close(fuseFd)
		klog.FromContext(ctx).Error(nil, "alinas does not support fd-passing, falling back to normal mount", "target", req.Target)
	}
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
	setupDefaultConfigs()
	if server.InitNASRSAPEM {
		if err := h.initRSAPrivateKey(); err != nil {
			klog.Fatalf("Failed to init NAS RSA private key: %v", err)
		}
	}
	go runCommandForever("aliyun-alinas-mount-watchdog")
	go runCommandForever("aliyun-cpfs-mount-watchdog")
}

// alinasKeygenCommand is the helper script shipped by aliyun-alinas-utils that
// generates the NAS RSA private key under <configDir>/alinas.
const alinasKeygenCommand = "alinas-keygen"

// rsaPrivateKeyFile is the NAS RSA private key filename generated under
// <configDir>/alinas by alinas-keygen.
const rsaPrivateKeyFile = "privateKey.pem"

func (h *Driver) configDir() string {
	if h.ConfigDir != "" {
		return h.ConfigDir
	}
	return configDir
}

// rsaPrivateKeyPath is the full path of the NAS RSA private key.
func (h *Driver) rsaPrivateKeyPath() string {
	return filepath.Join(h.configDir(), "alinas", rsaPrivateKeyFile)
}

// initRSAPrivateKey generates the NAS RSA private key by invoking the
// alinas-keygen helper script shipped by aliyun-alinas-utils at pod startup.
//
// alinas-keygen effectively runs:
//
//	mkdir -p /etc/aliyun/alinas &&
//	  openssl genpkey -algorithm RSA -out /etc/aliyun/alinas/privateKey.pem -pkeyopt rsa_keygen_bits:3072 &&
//	  chmod 400 /etc/aliyun/alinas/privateKey.pem
//
// The key it emits is an unencrypted PKCS#8 PEM block
// ("-----BEGIN PRIVATE KEY-----").
//
// This function is idempotent regardless of the script's behavior: it stats the
// key first and skips generation when it already exists, so an existing key
// (persisted via the /etc/aliyun hostPath) is never rotated on restart.
func (h *Driver) initRSAPrivateKey() error {
	keyPath := h.rsaPrivateKeyPath()
	if _, err := os.Stat(keyPath); err == nil {
		klog.InfoS("NAS RSA private key already exists, skipping generation", "path", keyPath)
		return nil
	} else if !os.IsNotExist(err) {
		return fmt.Errorf("stat %s: %w", keyPath, err)
	}

	klog.InfoS("Generating NAS RSA private key via alinas-keygen", "command", alinasKeygenCommand, "path", keyPath)

	cmd := exec.Command(alinasKeygenCommand)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("run %s: %w: %s", alinasKeygenCommand, err, strings.TrimSpace(string(out)))
	}
	klog.InfoS("alinas-keygen finished", "output", strings.TrimSpace(string(out)))
	return nil
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
	// Stop all jwtauth credential refreshers regardless of mount cleanup
	// policy, so no refresh goroutine outlives the driver.
	jwtauth.StopAll()

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
		if _, err := h.Unmount(target); err != nil {
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
		// The alinas config is always overwritten with the default shipped in
		// the image so that config updates from a new image take effect on
		// restart. cpfs keeps the copy-if-absent behavior.
		overwrite := name == "alinas"
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

			if !overwrite {
				if _, err := os.Stat(dstPath); err == nil {
					// File already exists, skip
					return nil
				} else if !os.IsNotExist(err) {
					return err
				}
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
	op.Options = append(op.Options, "no_start_watchdog")
	if op.FsType == fstypeAlinas {
		// options = append(options, "no_atomic_move", "auto_fallback_nfs")
		op.Options = append(op.Options, "no_atomic_move")
		op.Options = addAutoFallbackNFSMountOptions(op.Options)
	}
	// SensitiveOptions (e.g. jwtauth STS credentials) are passed separately so
	// mount-utils masks them in logs and error messages.
	err := m.MountSensitive(op.Source, op.Target, op.FsType, op.Options, op.SensitiveOptions)
	if err == nil {
		m.driver.targets.Store(op.Target, struct{}{})
	}
	return err
}

func (m *extendedMounter) Unmount(target string) error {
	err := m.Interface.Unmount(target)
	if err == nil {
		m.driver.targets.Delete(target)
		// The mount is gone; stop any jwtauth credential refresher serving it.
		jwtauth.DefaultManager.StopByTarget(target)
	}
	return err
}
