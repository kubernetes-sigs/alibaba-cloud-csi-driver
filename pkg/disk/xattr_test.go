package disk

import (
	"os"
	"testing"
	"testing/synctest"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/sys/unix"
	"k8s.io/klog/v2/ktesting"
)

func testDiskXattr(t *testing.T) {
	orgiDiskXattrName := DiskXattrName
	DiskXattrName = "user.testing-csi-managed-disk"
	t.Cleanup(func() {
		DiskXattrName = orgiDiskXattrName
	})

}

func TestDevMap(t *testing.T) {
	testDiskXattr(t)
	logger, _ := ktesting.NewTestContext(t)

	d := MockDisks{
		base: t.TempDir(),
	}
	d.AddDisk(t, "/vda", []byte("d-disk1"))
	d.AddDisk(t, "/vdb", nil) // d-disk2
	d.AddDisk(t, "/vdc", []byte("d-disk3"))
	d.AddDisk(t, "/vdd", []byte("d-disk4"))
	d.AddDisk(t, "/vde", []byte("d-disk5"))

	m, err := NewDevMap(&d)
	assert.NoError(t, err)

	require.NoError(t, os.Remove(d.base+"/vdc")) // vdc detached
	// detached then attached another disk at vdd
	require.NoError(t, unix.Setxattr(d.base+"/vdd", DiskXattrName, []byte("d-whatever"), 0))
	// detached then attached another unmanaged disk at vde
	require.NoError(t, unix.Removexattr(d.base+"/vde", DiskXattrName))

	p, err := m.Get(logger, "d-disk1")
	assert.NoError(t, err)
	assert.Equal(t, d.base+"/vda", p)

	for _, d := range []string{"d-disk2", "d-disk3", "d-disk4", "d-disk5"} {
		p, err := m.Get(logger, d)
		assert.NoError(t, err)
		assert.Equal(t, "", p)
	}

	require.NoError(t, os.WriteFile(d.base+"/vdf", nil, 0644))
	assert.NoError(t, m.Add("d-disk6", d.base+"/vdf"))
	p, err = m.Get(logger, "d-disk6")
	assert.NoError(t, err)
	assert.Equal(t, d.base+"/vdf", p)

	m.Delete("d-disk6")
	p, err = m.Get(logger, "d-disk6")
	assert.NoError(t, err)
	assert.Equal(t, "", p)

}

func TestDevMapAddError(t *testing.T) {
	testDiskXattr(t)
	m := devMap{}
	assert.Error(t, m.Add("d-disk1", "/not-exist"))
	assert.Error(t, m.Add(longDiskID, "/dev/vda"))
}

func TestDevMapRace(t *testing.T) {
	testDiskXattr(t)

	getXattrTestHook = func() {
		time.Sleep(2 * time.Millisecond)
	}
	defer func() {
		getXattrTestHook = func() {}
	}()

	logger, _ := ktesting.NewTestContext(t)
	synctest.Test(t, func(t *testing.T) {
		m := devMap{}
		base := t.TempDir()
		require.NoError(t, os.WriteFile(base+"/vda", nil, 0644))
		require.NoError(t, m.Add("d-disk1", base+"/vda"))

		require.NoError(t, os.Remove(base+"/vda"))
		require.NoError(t, os.WriteFile(base+"/vdb", nil, 0644))

		// disk1 detached
		// t0: read disk1 xattr, file not found
		// t1: disk1 re-attached as vdb
		// t2: disk1 remove from m, but value mismatch
		// t3: get disk2, vbd returned
		go func() {
			p, err := m.Get(logger, "d-disk1") // t0
			require.NoError(t, err)
			require.Equal(t, "", p) // t2
		}()
		time.Sleep(1 * time.Millisecond)
		require.NoError(t, m.Add("d-disk1", base+"/vdb")) // t1
		time.Sleep(2 * time.Millisecond)
		p, err := m.Get(logger, "d-disk1") // t3
		require.NoError(t, err)
		require.Equal(t, base+"/vdb", p)
	})
}
