package ossfs

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
	"syscall"
	"time"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy/server"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/klog/v2"
	"k8s.io/mount-utils"
)

func init() {
	server.RegisterDriver(NewDriver())
}

const (
	OssfsPasswdFile = "passwd-ossfs"
)

type Driver struct {
	pids *sync.Map
	wg   sync.WaitGroup
	raw  mount.Interface
}

func NewDriver() *Driver {
	return &Driver{
		pids: new(sync.Map),
		raw:  mount.NewWithoutSystemd(""),
	}
}

func (h *Driver) Name() string {
	return "ossfs"
}

func (h *Driver) Fstypes() []string {
	return []string{"ossfs"}
}

func (h *Driver) Mount(ctx context.Context, req *proxy.MountRequest) error {
	options := req.Options

	// prepare passwd file
	passwdFile, tokenDir, credOpts, err := prepareCredentialFiles(req.Secrets)
	if err != nil {
		return err
	}
	options = append(options, credOpts...)
	if passwdFile != "" {
		klog.V(4).InfoS("created ossfs passwd file", "path", passwdFile)
	}
	if tokenDir != "" {
		klog.V(4).InfoS("created ossfs token directory", "dir", tokenDir)
	}

	args := mount.MakeMountArgs(req.Source, req.Target, "", options)
	args = append(args, req.MountFlags...)
	args = append(args, "-f")

	cmd := exec.Command("ossfs", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Start()
	if err != nil {
		return fmt.Errorf("start ossfs failed: %w", err)
	}

	target := req.Target
	pid := cmd.Process.Pid
	klog.InfoS("Started ossfs", "pid", pid, "args", args)

	ossfsExited := make(chan error, 1)
	h.wg.Add(1)
	h.pids.Store(pid, cmd)
	go func() {
		defer h.wg.Done()
		defer h.pids.Delete(pid)

		err := cmd.Wait()
		if err != nil {
			klog.ErrorS(err, "ossfs exited with error", "mountpoint", target, "pid", pid)
		} else {
			klog.InfoS("ossfs exited", "mountpoint", target, "pid", pid)
		}
		ossfsExited <- err
		if passwdFile != "" {
			if err := os.Remove(passwdFile); err != nil {
				klog.ErrorS(err, "Remove passwd file", "mountpoint", target, "path", passwdFile)
			}
		}
		if tokenDir != "" {
			if err := os.RemoveAll(tokenDir); err != nil {
				klog.ErrorS(err, "Remove token directory", "mountpoint", target, "dir", tokenDir)
			}
		}

		close(ossfsExited)
	}()

	err = wait.PollUntilContextCancel(ctx, time.Second, true, func(ctx context.Context) (done bool, err error) {
		select {
		case err := <-ossfsExited:
			// TODO: collect ossfs outputs to return in error message
			if err != nil {
				return false, fmt.Errorf("ossfs exited: %w", err)
			}
			return false, fmt.Errorf("ossfs exited")
		default:
			notMnt, err := h.raw.IsLikelyNotMountPoint(target)
			if err != nil {
				klog.ErrorS(err, "check mountpoint", "mountpoint", target)
				return false, nil
			}
			if !notMnt {
				klog.InfoS("Successfully mounted", "mountpoint", target)
				return true, nil
			}
			return false, nil
		}
	})

	if err == nil {
		return nil
	}

	if wait.Interrupted(err) {
		// terminate ossfs process when timeout
		terr := cmd.Process.Signal(syscall.SIGTERM)
		if terr != nil {
			klog.ErrorS(err, "Failed to terminate ossfs", "pid", pid)
		}
		select {
		case <-ossfsExited:
		case <-time.After(time.Second * 2):
			kerr := cmd.Process.Kill()
			if kerr != nil && errors.Is(kerr, os.ErrProcessDone) {
				klog.ErrorS(err, "Failed to kill ossfs", "pid", pid)
			}
		}
	}
	return err
}

func (h *Driver) Init() {}

func (h *Driver) Terminate() {
	// terminate all running ossfs
	h.pids.Range(func(key, value any) bool {
		err := value.(*exec.Cmd).Process.Signal(syscall.SIGTERM)
		if err != nil {
			klog.ErrorS(err, "Failed to terminate ossfs", "pid", key)
		}
		klog.V(4).InfoS("Sended sigterm", "pid", key)
		return true
	})
	// wait all ossfs processes to exit
	h.wg.Wait()
	klog.InfoS("All ossfs processes exited")
}

func writeFile(dir, fileName, contents string, perm os.FileMode) error {
	file := filepath.Join(dir, fileName)
	return os.WriteFile(file, []byte(contents), perm)
}

// prepareCredentialFiles returns:
//  1. file:    path of ossfs credential file for fixed AKSK
//  2. dir:     dorectory of ossfs credential files for token
//  3. options: extra options
//  4. error
func prepareCredentialFiles(secrets map[string]string) (file, dir string, options []string, err error) {
	var tmpDir string
	tmpDir, err = os.MkdirTemp("", "ossfs-")
	if err != nil {
		return
	}

	// fixed AKSK
	if passwd := secrets[OssfsPasswdFile]; passwd != "" {
		err = writeFile(tmpDir, "passwd", passwd, 0o600)
		if err != nil {
			return
		}
		file = filepath.Join(tmpDir, "passwd")
		options = append(options, "passwd_file="+file)
		return
	}

	// token
	tokenKey := []string{mounter.KeyAccessKeyId, mounter.KeyAccessKeySecret, mounter.KeySecurityToken, mounter.KeyExpiration}
	tokenDir := filepath.Join(tmpDir, "token")
	var token bool
	for _, key := range tokenKey {
		val := secrets[filepath.Join(OssfsPasswdFile, key)]
		if val == "" {
			continue
		}
		err = os.MkdirAll(tokenDir, 0o644)
		if err != nil {
			klog.Errorf("mkdirall tokendir failed %v", err)
			return
		}
		err = writeFile(tokenDir, key, val, 0o600)
		if err != nil {
			klog.Errorf("writeFile %s failed %v", key, err)
			return
		}
		token = true
	}
	if token {
		dir = tokenDir
		options = append(options, "passwd_file="+tokenDir)
	}
	return
}
