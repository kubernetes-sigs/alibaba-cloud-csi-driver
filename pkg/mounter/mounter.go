package mounter

import mountutils "k8s.io/mount-utils"

type Mounter interface {
	mountutils.Interface
	MountWithSecrets(source, target, fstype string, options []string, secrets map[string]string) error
}
