package ossfs2

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"sync"
	"syscall"
	"time"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy/server"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
	"github.com/moby/sys/mountinfo"
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

func (h *Driver) runCmd(req *proxy.MountRequest, passwdFile string, fuseFd int) (cmd *exec.Cmd, err error) {
	// 1. use /dev/fd/3 as target for ossfs
	args := []string{"mount", "/dev/fd/3"}
	// ossfs2.0 forbid to use FUSE args
	// args = append(args, req.MountFlags...)
	args = append(args, []string{"-c", passwdFile}...)
	for _, o := range req.Options {
		args = append(args, fmt.Sprintf("--%s", o))
	}
	args = append(args, "-f")

	cmd = exec.Command("ossfs2", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	// 2. pase fd on /dev/fuse by fuseFd
	// Notes: Mountproxy client (csi-plugin) opens a fd on /dev/fuse,
	//        client passes it by unix conn.
	//        Mountproxy server (ossfs pod) receives it by unix conn,
	//        server mounts with ossfs with /dev/fd/3 (on /dev/fuse) as target.
	cmd.ExtraFiles = []*os.File{os.NewFile(uintptr(fuseFd), "/dev/fuse")}

	err = cmd.Start()
	if err != nil {
		err = fmt.Errorf("start ossfs2 failed: %w", err)
		return
	}
	klog.InfoS("Started ossfs2", "pid", cmd.Process.Pid, "args", cmd.Args)
	return
}

func (h *Driver) Mount(ctx context.Context, req *proxy.MountRequest, fuseFd int) error {
	target := req.Target

	// prepare passwd file
	var passwdFile string
	if passwd := req.Secrets[utils.GetPasswdFileName("ossfs2")]; passwd != "" {
		tmpDir, err := os.MkdirTemp("", "ossfs2-")
		if err != nil {
			return err
		}
		passwdFile = filepath.Join(tmpDir, "passwd")
		err = os.WriteFile(passwdFile, []byte(passwd), 0o600)
		if err != nil {
			return err
		}
		klog.V(4).InfoS("created ossfs2 configuration file", "path", passwdFile)
	}

	cmd, err := h.runCmd(req, passwdFile, fuseFd)
	if err != nil {
		return fmt.Errorf("start ossfs2 failed: %w", err)
	}

	info, err := mountinfo.GetMounts(func(i *mountinfo.Info) (skip bool, stop bool) {
		return i.Mountpoint != target, i.Mountpoint == target
	})
	if len(info) == 0 {
		return fmt.Errorf("mountpoint %s is not mounted", target)
	}
	chanId := info[0].Minor

	ossfsExited := make(chan error, 1)
	h.wg.Add(1)
	go func() {
		defer h.wg.Done()
		// release to avoid unexpected forking
		defer syscall.Close(fuseFd)
		defer server.CleanPasswdFile(passwdFile, target)

		for {
			pid := cmd.Process.Pid
			h.pids.Store(pid, cmd)

			err = cmd.Wait()
			if err != nil {
				klog.ErrorS(err, "ossfs2 exited with error", "mountpoint", target, "pid", pid)
			} else {
				klog.InfoS("ossfs2 exited", "mountpoint", target, "pid", pid)
			}
			h.pids.Delete(pid)

			hasClosed := server.SafeSend(ossfsExited, err)
			if !hasClosed {
				return
			}

			waitStatus, ok := cmd.ProcessState.Sys().(syscall.WaitStatus)
			if ok && waitStatus.Signal() == syscall.SIGTERM {
				return
			}

			err = fmt.Errorf("Ossfs2 exited, need recovery")
			klog.ErrorS(err, "mountpoint", target)
			klog.InfoS("Flush fuse connection", "chanId", chanId, "mountpoint", target)
			err = os.WriteFile(filepath.Join(server.FuseConnectionsDir, strconv.Itoa(chanId), "flush"), []byte("1"), 0o600)
			if err != nil {
				klog.ErrorS(err, "Failed to flush fuse connection", "chanId", chanId, "mountpoint", target)
				return
			}

			cmd, err = h.runCmd(req, passwdFile, fuseFd)
			if err != nil {
				return
			}
		}
	}()
	defer func() {
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
			klog.ErrorS(err, "Failed to terminate ossfs2", "pid", cmd.Process.Pid)
		}
		select {
		case <-ossfsExited:
		case <-time.After(time.Second * 2):
			kerr := cmd.Process.Kill()
			if kerr != nil && errors.Is(kerr, os.ErrProcessDone) {
				klog.ErrorS(err, "Failed to kill ossfs2", "pid", cmd.Process.Pid)
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
