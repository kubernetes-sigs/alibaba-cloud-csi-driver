package interceptors

import (
	"context"
	"fmt"
	"os/exec"
	"strings"
	"time"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/jwtauth"
	"k8s.io/klog/v2"
)

const (
	// alinasCertRefreshCommand pushes a rotated STS credential to a live
	// alinas/cpfs mount. It ships with aliyun-alinas-utils.
	alinasCertRefreshCommand = "alinas-tls-cert-refresh"

	// alinasCertRefreshTimeout bounds a single refresh command invocation so a
	// hung command cannot stall the refresh loop or Stop.
	alinasCertRefreshTimeout = 30 * time.Second

	// redactedPlaceholder replaces credential material in command output.
	redactedPlaceholder = "<redacted>"
)

// alinasCertRefreshSink is the NAS/CPFS credential delivery mechanism: it
// pushes each rotated STS credential to a live alinas mount by executing
// alinas-tls-cert-refresh. Nothing is written to disk, so Cleanup is a no-op.
//
// The credential goes on argv, which is briefly visible in /proc/<pid>/cmdline
// while the command runs; the CLI also reads it as JSON on stdin, which would
// avoid that. Either way it must not escape this sink: the arguments are never
// logged, and the CLI echoes arguments it does not recognize back into its own
// output, so that output is redacted before it reaches an error.
type alinasCertRefreshSink struct {
	mountPoint string

	// runCommand runs the refresh command and returns its combined output.
	// It is a field so tests can inject a fake runner.
	runCommand func(ctx context.Context, name string, args ...string) ([]byte, error)
}

var _ jwtauth.CredentialSink = &alinasCertRefreshSink{}

func newAlinasCertRefreshSink(mountPoint string) *alinasCertRefreshSink {
	return &alinasCertRefreshSink{
		mountPoint: mountPoint,
		runCommand: func(ctx context.Context, name string, args ...string) ([]byte, error) {
			return exec.CommandContext(ctx, name, args...).CombinedOutput()
		},
	}
}

func (s *alinasCertRefreshSink) Apply(cred *jwtauth.STSToken) error {
	ctx, cancel := context.WithTimeout(context.Background(), alinasCertRefreshTimeout)
	defer cancel()
	return refreshAlinasCredential(ctx, s.runCommand, s.mountPoint, cred)
}

// RefreshAlinasCredential installs cred on a live alinas/cpfs mount by executing
// the vendor refresh command. Shared with the jwtauth refresh loop, so a rotated
// credential reaches a mount the same way whoever decided to rotate it.
func RefreshAlinasCredential(ctx context.Context, mountPoint string, cred *jwtauth.STSToken) error {
	run := func(ctx context.Context, name string, args ...string) ([]byte, error) {
		return exec.CommandContext(ctx, name, args...).CombinedOutput()
	}
	ctx, cancel := context.WithTimeout(ctx, alinasCertRefreshTimeout)
	defer cancel()
	return refreshAlinasCredential(ctx, run, mountPoint, cred)
}

func refreshAlinasCredential(
	ctx context.Context,
	runCommand func(ctx context.Context, name string, args ...string) ([]byte, error),
	mountPoint string,
	cred *jwtauth.STSToken,
) error {
	// SECURITY: never log these arguments; they contain the credential.
	args := []string{
		"--mount-point", mountPoint,
		"--ak", cred.AccessKeyID,
		"--sk", cred.AccessKeySecret,
		"--token", cred.SecurityToken,
	}
	output, err := runCommand(ctx, alinasCertRefreshCommand, args...)
	if err != nil {
		return fmt.Errorf("%s failed for mount point %s: %w, output: %s",
			alinasCertRefreshCommand, mountPoint, err, redactCredential(string(output), cred))
	}
	klog.V(4).InfoS("refreshed alinas mount credential", "command", alinasCertRefreshCommand, "mountpoint", mountPoint)
	return nil
}

// Cleanup is a no-op: this sink never persists anything.
func (s *alinasCertRefreshSink) Cleanup() {}

// redactCredential replaces every field of cred that appears in s with a
// placeholder, so command output stays diagnosable (exit status, mount point,
// the CLI's own message) without ever carrying credential material into an
// error or a log line.
func redactCredential(s string, cred *jwtauth.STSToken) string {
	if cred == nil {
		return s
	}
	for _, secret := range []string{cred.AccessKeySecret, cred.SecurityToken, cred.AccessKeyID} {
		if secret == "" {
			continue
		}
		s = strings.ReplaceAll(s, secret, redactedPlaceholder)
	}
	return s
}
