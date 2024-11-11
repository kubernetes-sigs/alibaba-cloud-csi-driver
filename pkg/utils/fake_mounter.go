package utils

type fakeMounter struct {
}

func NewFakeMounter() Mounter {
	return &fakeMounter{}
}

func (m *fakeMounter) EnsureBlock(target string) error {
	return nil
}

func (m *fakeMounter) EnsureFolder(target string) error {
	return nil
}

func (m *fakeMounter) Format(source, fsType string) error {
	return nil
}

func (m *fakeMounter) Mount(source, target, fsType string, options ...string) error {
	return nil
}

func (m *fakeMounter) MountBlock(source, target string, options ...string) error {
	return nil
}

func (m *fakeMounter) Unmount(target string) error {
	return nil
}

func (m *fakeMounter) HasMountRefs(mountPath string, mountRefs []string) bool {
	return false
}
