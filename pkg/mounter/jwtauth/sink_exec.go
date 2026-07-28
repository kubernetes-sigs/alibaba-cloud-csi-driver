package jwtauth

import (
	"context"
	"fmt"
	"os/exec"
	"time"

	"k8s.io/klog/v2"
)

const (
	// refreshCommand pushes a rotated STS credential to a live alinas/cpfs
	// mount. It ships with aliyun-alinas-utils.
	refreshCommand = "alinas-tls-cert-refresh"

	// execTimeout bounds a single refresh command invocation so a hung
	// command cannot stall the refresh loop or Stop.
	execTimeout = 30 * time.Second
)

// ExecSink delivers rotated STS credentials to a live mount by executing the
// alinas-tls-cert-refresh command. Nothing is written to disk by this sink,
// so Cleanup is a no-op.
//
// The vendor CLI only accepts the credential via argv, which is briefly
// visible in /proc/<pid>/cmdline while the command runs. This is a constraint
// of the CLI interface; the arguments are therefore never logged here.
type ExecSink struct {
	mountPoint string

	// runCommand runs the refresh command and returns its combined output.
	// It is a field so tests can inject a fake runner.
	runCommand func(ctx context.Context, name string, args ...string) ([]byte, error)
}

var _ CredentialSink = &ExecSink{}

func NewExecSink(mountPoint string) *ExecSink {
	return &ExecSink{
		mountPoint: mountPoint,
		runCommand: func(ctx context.Context, name string, args ...string) ([]byte, error) {
			return exec.CommandContext(ctx, name, args...).CombinedOutput()
		},
	}
}

func (s *ExecSink) Apply(cred *STSToken) error {
	ctx, cancel := context.WithTimeout(context.Background(), execTimeout)
	defer cancel()

	// SECURITY: never log these arguments - they contain the credential.
	args := []string{
		"--mount-point", s.mountPoint,
		"--ak", cred.AccessKeyID,
		"--sk", cred.AccessKeySecret,
		"--token", cred.SecurityToken,
	}
	output, err := s.runCommand(ctx, refreshCommand, args...)
	if err != nil {
		return fmt.Errorf("%s failed for mount point %s: %w, output: %s",
			refreshCommand, s.mountPoint, err, string(output))
	}
	klog.V(4).InfoS("refreshed alinas mount credential", "command", refreshCommand, "mountpoint", s.mountPoint)
	return nil
}

// Cleanup is a no-op: ExecSink never persists anything.
func (s *ExecSink) Cleanup() {}
