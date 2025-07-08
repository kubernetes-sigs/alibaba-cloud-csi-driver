package server

import (
	"context"
	"fmt"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
)

type Driver interface {
	Name() string
	Fstypes() []string
	Init()
	Terminate()
	Mount(ctx context.Context, req *proxy.MountRequest) error
	Warmup(string, string, int, int64, int64)
}

var (
	fstypeToDriver = map[string]Driver{}
	nameToDriver   = map[string]Driver{}
)

func RegisterDriver(driver Driver) {
	nameToDriver[driver.Name()] = driver
}

func handleMountRequest(ctx context.Context, req *proxy.MountRequest) error {
	h := fstypeToDriver[req.Fstype]
	if h == nil {
		return fmt.Errorf("fstype %q not supported", req.Fstype)
	}
	err := h.Mount(ctx, req)
	if err != nil {
		return err
	}
	for _, warmupDir := range req.WarmupDirs {
		h.Warmup(req.Target, warmupDir, req.WarmupWorkers, req.WarmupTotalGBs*utils.GiB, req.WarmupPerFileMaxGBs*utils.GiB)
	}
	return nil
}
