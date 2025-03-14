package sfdisk

import (
	"bytes"
	"context"
	"errors"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"testing"

	utilsos "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils/os"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func runCommand(t *testing.T, name string, arg ...string) string {
	cmd := exec.Command(name, arg...)
	output, err := cmd.Output()
	require.NoError(t, utilsos.ErrWithStderr(err))
	return string(output)
}

func TestExpandPartition(t *testing.T) {
	path, err := exec.LookPath("sfdisk")
	if errors.Is(err, exec.ErrNotFound) {
		t.Skip("sfdisk not found")
	}
	require.NoError(t, err)
	t.Logf("sfdisk found at: %s", path)

	testImage := t.TempDir() + "/test.img"
	f, err := os.Create(testImage)
	require.NoError(t, err)
	require.NoError(t, f.Close())
	require.NoError(t, os.Truncate(testImage, 1<<23)) // 8MB

	cmd := exec.Command("sfdisk", testImage)
	cmd.Stdin = bytes.NewReader([]byte("label: gpt\n,\n")) // create a single partition
	result, err := cmd.Output()
	require.NoError(t, utilsos.ErrWithStderr(err))
	t.Logf("create partition success: %s", string(result))

	result, err = exec.Command("losetup", "--find", testImage, "--show", "--partscan").Output()
	if err != nil {
		t.Skip("losetup failed", err)
	}
	disk := strings.TrimSpace(string(result))
	t.Logf("loop device: %s", disk)
	defer func() {
		runCommand(t, "losetup", "-d", disk)
	}()

	part := disk + "p1"
	output := runCommand(t, "mkfs.ext4", "-F", part)
	t.Logf("mkfs.ext4 output: %s", output)
	mntPath := t.TempDir()
	runCommand(t, "mount", part, mntPath)
	defer func() {
		runCommand(t, "umount", mntPath)
	}()

	require.NoError(t, os.Truncate(testImage, 1<<24)) // expand to 16MB
	runCommand(t, "losetup", "--set-capacity", disk)
	require.NoError(t, ExpandPartition(context.Background(), disk, "1"))

	dump := runCommand(t, "sfdisk", "--dump", disk)
	assert.Contains(t, strings.ReplaceAll(dump, " ", ""), "p1:start=2048,size=30687")

	output = runCommand(t, "blockdev", "--getsize64", part)
	realSize, err := strconv.Atoi(strings.TrimSpace(output))
	assert.NoError(t, err)
	assert.GreaterOrEqual(t, realSize, 12<<20)
}
