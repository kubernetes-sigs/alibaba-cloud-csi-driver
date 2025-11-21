package mounter

import (
	"context"
	"fmt"

	"k8s.io/klog/v2"
	mountutils "k8s.io/mount-utils"
)

type Mounter interface {
	mountutils.Interface
	ExtendedMount(ctx context.Context, op *MountOperation) error
}

type MountOperation struct {
	Source      string
	Target      string
	FsType      string
	Options     []string
	Args        []string
	Secrets     map[string]string
	MetricsPath string
	VolumeID    string

	MountResult any
}

type MountInterceptor interface {
	BeforeMount(op *MountOperation, err error) (*MountOperation, error)
	AfterMount(op *MountOperation, err error) error
}

type MountWorkflow struct {
	Mounter
	interceptors []MountInterceptor
}

var _ Mounter = &MountWorkflow{}

func (w *MountWorkflow) ExtendedMount(ctx context.Context, op *MountOperation) error {
	var err error
	for i, interceptor := range w.interceptors {
		if op, err = interceptor.BeforeMount(op, err); err != nil && i != len(w.interceptors)-1 {
			// Log error but continue to the next interceptor, since some interceptors may
			// want to handle the error, e.g. the OssMonitorInterceptor.
			// Only log for the intermediate interceptors, otherwise the final error will be printed twice.
			klog.ErrorS(err, "error executing BeforeMount interceptor")
		}
	}
	if err != nil {
		return fmt.Errorf("error executing BeforeMount interceptor: %w", err)
	}

	err = w.Mounter.ExtendedMount(ctx, op)
	for _, interceptor := range w.interceptors {
		if afterErr := interceptor.AfterMount(op, err); afterErr != nil {
			// Log error but continue to the next interceptor.
			// Do not override the original error from mount operation.
			klog.ErrorS(afterErr, "error executing AfterMount interceptor")
		}
	}
	return err
}

func NewForMounter(m Mounter, interceptors ...MountInterceptor) Mounter {
	return &MountWorkflow{
		Mounter:      m,
		interceptors: interceptors,
	}
}
