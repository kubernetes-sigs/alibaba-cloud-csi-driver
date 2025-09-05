package ossfs2

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"sync"
	"syscall"
	"time"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy/server"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/klog/v2"
	"k8s.io/mount-utils"
)

func init() {
	server.RegisterDriver(NewDriver())
}

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
	return "ossfs2"
}

func (h *Driver) Fstypes() []string {
	return []string{"ossfs2"}
}

func (h *Driver) Mount(ctx context.Context, req *proxy.MountRequest) error {
	options := req.Options

	// prepare passwd file
	passwdFile, err := utils.SaveOssSecretsToFile(req.Secrets, req.Fstype)
	if err != nil {
		return err
	}

	args := []string{"mount", req.Target}
	// ossfs2.0 forbid to use FUSE args
	// args = append(args, req.MountFlags...)
	if passwdFile != "" {
		args = append(args, []string{"-c", passwdFile}...)
	}
	for _, o := range options {
		args = append(args, fmt.Sprintf("--%s", o))
	}
	args = append(args, "-f")

	cmd := exec.Command("ossfs2", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Start()
	if err != nil {
		return fmt.Errorf("start ossfs2 failed: %w", err)
	}

	target := req.Target
	pid := cmd.Process.Pid
	klog.InfoS("Started ossfs2", "pid", pid, "args", args)

	ossfsExited := make(chan error, 1)
	h.wg.Add(1)
	h.pids.Store(pid, cmd)
	go func() {
		defer h.wg.Done()
		defer h.pids.Delete(pid)

		err := cmd.Wait()
		if err != nil {
			klog.ErrorS(err, "ossfs2 exited with error", "mountpoint", target, "pid", pid)
		} else {
			klog.InfoS("ossfs2 exited", "mountpoint", target, "pid", pid)
		}
		ossfsExited <- err
		if err := os.Remove(passwdFile); err != nil {
			klog.ErrorS(err, "Remove configuration file", "mountpoint", target, "path", passwdFile)
		}
		close(ossfsExited)
	}()

	err = wait.PollUntilContextCancel(ctx, time.Second, true, func(ctx context.Context) (done bool, err error) {
		select {
		case err := <-ossfsExited:
			// TODO: collect ossfs outputs to return in error message
			if err != nil {
				return false, fmt.Errorf("ossfs2 exited: %w", err)
			}
			return false, fmt.Errorf("ossfs2 exited")
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
			klog.ErrorS(err, "Failed to terminate ossfs2", "pid", pid)
		}
		select {
		case <-ossfsExited:
		case <-time.After(time.Second * 2):
			kerr := cmd.Process.Kill()
			if kerr != nil && errors.Is(kerr, os.ErrProcessDone) {
				klog.ErrorS(err, "Failed to kill ossfs2", "pid", pid)
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
			klog.ErrorS(err, "Failed to terminate ossfs2", "pid", key)
		}
		klog.V(4).InfoS("Sended sigterm", "pid", key)
		return true
	})
	// wait all ossfs processes to exit
	h.wg.Wait()
	klog.InfoS("All ossfs2 processes exited")
}
