//go:build !windows

package disk

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type fakePCIDevice struct {
	addr     string
	vendor   uint16
	class    uint32
	totalVFs int // > 0 for SR-IOV PF
}

func fakeSysfs(t *testing.T, devices ...fakePCIDevice) string {
	dir := t.TempDir()
	for _, d := range devices {
		p := filepath.Join(dir, d.addr)
		require.NoError(t, os.Mkdir(p, 0o755))
		write := func(name, content string) {
			require.NoError(t, os.WriteFile(filepath.Join(p, name), []byte(content), 0o644))
		}
		write("vendor", fmt.Sprintf("0x%04x\n", d.vendor))
		write("class", fmt.Sprintf("0x%06x\n", d.class))
		if d.totalVFs > 0 {
			write("sriov_totalvfs", fmt.Sprintf("%d\n", d.totalVFs))
		}
	}
	return dir
}

func TestHasSRIOVStorageController(t *testing.T) {
	gpu := fakePCIDevice{addr: "0000:00:02.0", vendor: 0x10de, class: 0x030200, totalVFs: 16}
	cases := []struct {
		name   string
		device fakePCIDevice
		vf     bool
	}{
		{"SCSI with SR-IOV", fakePCIDevice{addr: "0000:4b:00.0", vendor: pciVendorAlibaba, class: 0x010000, totalVFs: 256}, true},
		{"other storage with SR-IOV", fakePCIDevice{addr: "0000:4b:00.0", vendor: pciVendorAlibaba, class: 0x018000, totalVFs: 256}, true},
		{"without SR-IOV", fakePCIDevice{addr: "0000:4b:00.0", vendor: pciVendorAlibaba, class: 0x010000}, false},
		{"not function 0", fakePCIDevice{addr: "0000:4b:00.1", vendor: pciVendorAlibaba, class: 0x010000, totalVFs: 256}, false},
		{"NVMe", fakePCIDevice{addr: "0000:4b:00.0", vendor: pciVendorAlibaba, class: 0x010802, totalVFs: 256}, false},
		{"other vendor", fakePCIDevice{addr: "0000:4b:00.0", vendor: 0x8086, class: 0x010000, totalVFs: 256}, false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			vf, err := hasSRIOVStorageController(fakeSysfs(t, gpu, c.device))
			require.NoError(t, err)
			assert.Equal(t, c.vf, vf)
		})
	}

	_, err := hasSRIOVStorageController(filepath.Join(t.TempDir(), "not-exist"))
	assert.Error(t, err)
}
