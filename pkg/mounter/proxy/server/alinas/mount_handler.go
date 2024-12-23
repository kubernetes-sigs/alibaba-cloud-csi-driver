package alinas

import (
	"context"
	"os/exec"
	"time"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy/server"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/klog/v2"
	"k8s.io/mount-utils"
)

const (
	fstypeCpfsNfs = "cpfs-nfs"
	fstypeAlinas  = "alinas"
)

func init() {
	server.RegisterMountHandler(&MountHandler{
		mounter: mount.New(""),
	}, fstypeCpfsNfs, fstypeAlinas)
}

type MountHandler struct {
	mounter mount.Interface
}

func (h *MountHandler) Mount(ctx context.Context, req *proxy.MountRequest) error {
	klog.InfoS("Mounting", "fstype", req.Fstype, "source", req.Source, "target", req.Target, "options", req.Options)
	options := append(req.Options, "no_start_watchdog")
	if req.Fstype == fstypeAlinas {
		options = append(options, "no_atomic_move", "auto_fallback_nfs")
	}
	return h.mounter.Mount(req.Source, req.Target, req.Fstype, options)
}

func (h *MountHandler) Init() {
	go runCommandForever("aliyun-alinas-mount-watchdog")
	go runCommandForever("aliyun-cpfs-mount-watchdog")
}

func (h *MountHandler) Terminate() {}

func runCommandForever(command string, args ...string) {
	wait.Forever(func() {
		klog.InfoS("Starting", "command", command)
		cmd := exec.Command(command, args...)
		err := cmd.Run()
		if err != nil {
			klog.ErrorS(err, "Exited", "command", command)
		}
	}, time.Second)
}
