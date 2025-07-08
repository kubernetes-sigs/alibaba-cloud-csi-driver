package ossfs2

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

	args := []string{"mount", req.Target}
	// ossfs2.0 forbid to use FUSE args
	// args = append(args, req.MountFlags...)
	args = append(args, []string{"-c", passwdFile}...)
	for _, o := range options {
		args = append(args, fmt.Sprintf("--%s", o))
	}
	args = append(args, "-f")

	cmd := exec.Command("ossfs2", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Start()
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

func (h *Driver) Warmup(targetPath, warmupDir string, workerCount int, totalBytes int64, perFileMaxBytes int64) {
	klog.Infof("Starting FUSE mountpoint warmup for: %s, workerCount: %d, totalBytes: %v, perFileMaxBytes: %v", targetPath, workerCount, totalBytes, perFileMaxBytes)
	startTime := time.Now()

	warmupFullPath := filepath.Join(targetPath, warmupDir)
	// Find the immediate entries (files and subdirectories) - our chunks
	entries, err := os.ReadDir(warmupFullPath)
	if err != nil {
		klog.Errorf("Error reading mountpoint directory %s: %v", warmupFullPath, err)
		return
	}

	if len(entries) == 0 {
		klog.Errorf("No entries found in %s to process.", warmupFullPath)
		return
	}

	// Use a channel to send entry paths to workers
	entryChan := make(chan string, len(entries))
	// Use a channel to receive bytes read from workers
	bytesReadChan := make(chan int64, workerCount)
	var totalBytesRead int64
	var totalBytesMu sync.Mutex
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		for bytes := range bytesReadChan {
			totalBytesMu.Lock()
			totalBytesRead += bytes
			totalBytesMu.Unlock()
		}
		wg.Done()
	}()

	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go worker(i, entryChan, bytesReadChan, perFileMaxBytes, &wg)
	}

	for _, entry := range entries {
		entryPath := filepath.Join(warmupFullPath, entry.Name())
		entryChan <- entryPath
	}
	close(entryChan)

	wg.Wait()
	// all works done
	duration := time.Since(startTime)
	klog.Infof("Finished FUSE mountpoint warmup for: %s, total bytes read: %d MiB, duration: %v", warmupFullPath, totalBytesRead/(1024*1024), duration)
	close(bytesReadChan)
}
