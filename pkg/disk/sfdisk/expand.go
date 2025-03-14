package sfdisk

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"

	utilsos "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils/os"

	"k8s.io/klog/v2"
)

func ExpandPartition(ctx context.Context, disk, partition string) error {
	logger := klog.FromContext(ctx)

	dump, err := exec.CommandContext(ctx, "sfdisk", "--dump", disk).Output()
	if err != nil {
		return fmt.Errorf("failed to dump current partition table of %s: %w", disk, utilsos.ErrWithStderr(err))
	}
	dumpStr := string(dump)
	logger.V(4).Info("sfdisk dump before expansion", "dump", dumpStr)

	// Don't cancel this, we don't want to corrupt the partition table
	// --lock according to https://systemd.io/BLOCK_DEVICE_LOCKING/
	// --no-reread --no-tell-kernel is necessary for online expansion. We will use partx to update the kernel.
	cmd := exec.Command("sfdisk", "--lock", "--no-reread", "--no-tell-kernel", disk, "-N", partition)
	cmd.Stdin = bytes.NewReader([]byte(",+")) // enlarge the partition as much as possible
	result, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("failed to expand partition %s on %s: %w\noriginal table looked like:\n%s", partition, disk, utilsos.ErrWithStderr(err), dumpStr)
	}
	logger.V(3).Info("sfdisk success", "output", string(result))

	cmd = exec.Command("partx", "--update", disk, "--nr", partition)
	result, err = cmd.Output()
	if err != nil {
		return fmt.Errorf("failed to update partition %s on %s: %w", partition, disk, utilsos.ErrWithStderr(err))
	}
	logger.V(3).Info("partx success", "output", string(result))
	return nil
}
