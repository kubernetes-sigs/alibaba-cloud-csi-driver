package mounter

import (
	"fmt"
	"strings"

	mountutils "k8s.io/mount-utils"
)

type Mounter interface {
	mountutils.Interface
	Name() string
	MountWithSecrets(source, target, fstype string, options []string, secrets map[string]string) error
	RotateToken(target, fstype string, secrets map[string]string) error
}

func ErrNotImplemented(driver, mounterType, method string) error {
	return fmt.Errorf("%s(%s): %s not implemented", mounterType, driver, method)
}

func IsNotImplementedErr(err error) bool {
	if err == nil {
		return false
	}
	return strings.Contains(err.Error(), "not implemented")
}
