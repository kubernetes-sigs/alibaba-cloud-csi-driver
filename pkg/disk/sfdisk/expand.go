package sfdisk

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"

	utilsos "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils/os"

	"golang.org/x/sys/unix"
	"k8s.io/klog/v2"
)

func ExpandPartition(ctx context.Context, disk, partition string) error {
	logger := klog.FromContext(ctx)
	fd, err := unix.Open(disk, unix.O_RDONLY, 0)
	if err != nil {
		return err
	}
	defer func() {
		if err := unix.Close(fd); err != nil {
			logger.Error(err, "failed to close", "fd", fd)
		}
	}()

	err = unix.Flock(fd, unix.LOCK_EX) // as suggested in the man sfdisk(8)
	if err != nil {
		return fmt.Errorf("failed to lock %s exclusively: %v", disk, err)
	}
	defer func() {
		if err := unix.Flock(fd, unix.LOCK_UN); err != nil {
			logger.Error(err, "failed to unlock", "fd", fd)
		}
	}()

	dump, err := exec.CommandContext(ctx, "sfdisk", "--dump", disk).Output()
	if err != nil {
		return fmt.Errorf("failed to dump current partition table of %s: %v", disk, utilsos.ErrWithStderr(err))
	}
	dumpStr := string(dump)
	logger.V(4).Info("sfdisk dump before expansion", "dump", dumpStr)

	// Don't cancel this, we don't want to corrupt the partition table
	cmd := exec.Command("sfdisk", disk, "-N", partition)
	cmd.Stdin = bytes.NewReader([]byte(",+")) // enlarge the partition as much as possible
	result, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("failed to expand partition %s on %s: %v\noriginal table looked like:\n%s", partition, disk, utilsos.ErrWithStderr(err), dumpStr)
	}
	logger.V(3).Info("sfdisk success", "output", string(result))
	return nil
}
