package sfdisk

import (
	"bytes"
	"context"
	"errors"
	"os"
	"os/exec"
	"strings"
	"testing"

	utilsos "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils/os"
	"github.com/stretchr/testify/assert"
)

func TestExpandPartition(t *testing.T) {
	path, err := exec.LookPath("sfdisk")
	if errors.Is(err, exec.ErrNotFound) {
		t.Skip("sfdisk not found")
	}
	assert.NoError(t, err)
	t.Logf("sfdisk found at: %s", path)

	testImage := t.TempDir() + "/test.img"
	_, err = os.Create(testImage)
	assert.NoError(t, err)
	assert.NoError(t, os.Truncate(testImage, 1<<23)) // 8MB

	cmd := exec.Command("sfdisk", testImage)
	cmd.Stdin = bytes.NewReader([]byte("label: gpt\n,\n")) // create a single partition
	result, err := cmd.Output()
	assert.NoError(t, utilsos.ErrWithStderr(err))
	t.Logf("create partition success: %s", string(result))

	assert.NoError(t, os.Truncate(testImage, 1<<24)) // expand to 16MB
	assert.NoError(t, ExpandPartition(context.Background(), testImage, "1"))

	dump, err := exec.Command("sfdisk", "--dump", testImage).Output()
	assert.NoError(t, utilsos.ErrWithStderr(err))
	assert.Contains(t, strings.ReplaceAll(string(dump), " ", ""), "test.img1:start=2048,size=30687")
}
