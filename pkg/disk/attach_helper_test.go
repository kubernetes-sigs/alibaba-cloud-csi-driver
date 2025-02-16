/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package disk

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/util/sets"
)

func TestGetNoSerialDevicesFromSysfs(t *testing.T) {
	sysfs := t.TempDir()
	assert.NoError(t, os.MkdirAll(sysfs+"/block", 0755))

	assert.NoError(t, os.MkdirAll(sysfs+"/devices/pci0000:00/0000:00:07.0/nvme/nvme0/nvme0n1", 0755))
	assert.NoError(t, os.Symlink("../../nvme0", sysfs+"/devices/pci0000:00/0000:00:07.0/nvme/nvme0/nvme0n1/device"))
	assert.NoError(t, os.Symlink("../devices/pci0000:00/0000:00:07.0/nvme/nvme0/nvme0n1", sysfs+"/block/nvme0n1"))

	assert.NoError(t, os.MkdirAll(sysfs+"/devices/pci0000:00/0000:00:08.0/nvme/nvme1/nvme1n1", 0755))
	assert.NoError(t, os.Symlink("../../nvme1", sysfs+"/devices/pci0000:00/0000:00:08.0/nvme/nvme1/nvme1n1/device"))
	assert.NoError(t, os.Symlink("../devices/pci0000:00/0000:00:08.0/nvme/nvme1/nvme1n1", sysfs+"/block/nvme1n1"))

	assert.NoError(t, os.MkdirAll(sysfs+"/devices/pci0000:00/0000:00:0a.0/virtio7/block/vdb", 0755))
	assert.NoError(t, os.Symlink("../devices/pci0000:00/0000:00:0a.0/virtio7/block/vdb", sysfs+"/block/vdb"))

	assert.NoError(t, os.MkdirAll(sysfs+"/devices/pci0000:00/0000:00:0b.0/virtio8/block/vdc", 0755))
	assert.NoError(t, os.Symlink("../devices/pci0000:00/0000:00:0b.0/virtio8/block/vdc", sysfs+"/block/vdc"))

	assert.NoError(t, os.MkdirAll(sysfs+"/devices/pci0000:00/0000:00:0c.0/virtio9/block/vdd", 0755))
	assert.NoError(t, os.Symlink("../devices/pci0000:00/0000:00:0c.0/virtio9/block/vdd", sysfs+"/block/vdd"))

	assert.NoError(t, os.WriteFile(sysfs+"/block/vdb/serial", []byte("serialofvdb"), 0644))
	assert.NoError(t, os.WriteFile(sysfs+"/block/vdc/serial", []byte(""), 0644))
	// vdd no serial file
	assert.NoError(t, os.WriteFile(sysfs+"/block/nvme0n1/device/serial", []byte("serialofnvme0\n"), 0644))
	assert.NoError(t, os.WriteFile(sysfs+"/block/nvme1n1/device/serial", []byte("\n"), 0644))

	devices, err := getNoSerialDevicesFromSysfs(sysfs + "/block")
	assert.NoError(t, err)
	assert.Equal(t, sets.New("/dev/vdc", "/dev/vdd", "/dev/nvme1n1"), devices)
}
