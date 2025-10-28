package mounter

import (
	mountutils "k8s.io/mount-utils"
)

type ExtendedMountParams struct {
	Secrets     map[string]string
	MetricsPath string
}

type Mounter interface {
	mountutils.Interface
	ExtendedMount(source, target, fstype string, options []string, parms ExtendedMountParams) error
}
