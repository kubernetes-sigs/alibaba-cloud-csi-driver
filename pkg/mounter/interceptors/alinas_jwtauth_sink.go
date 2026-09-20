package interceptors

import (
	"bytes"
	"context"
	"encoding/json"
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
// The credential goes in on stdin, never on argv: argv is world-readable through
// /proc/<pid>/cmdline for as long as the command runs, which is enough for any
// process on the node (and anything sampling ps) to lift an AK/SK/token. It must
// not escape through the command's output either: the CLI echoes arguments it
// does not recognize, so its output is redacted before reaching an error.
type alinasCertRefreshSink struct {
	mountPoint string

	// runCommand runs the refresh command with stdin and returns its combined
	// output. It is a field so tests can inject a fake runner.
	runCommand runner
}

// runner executes the refresh command, feeding it stdin, and returns its combined
// output.
type runner func(ctx context.Context, stdin []byte, name string, args ...string) ([]byte, error)

func execCommand(ctx context.Context, stdin []byte, name string, args ...string) ([]byte, error) {
	cmd := exec.CommandContext(ctx, name, args...)
	cmd.Stdin = bytes.NewReader(stdin)
	return cmd.CombinedOutput()
}

var _ jwtauth.CredentialSink = &alinasCertRefreshSink{}

func newAlinasCertRefreshSink(mountPoint string) *alinasCertRefreshSink {
	return &alinasCertRefreshSink{
		mountPoint: mountPoint,
		runCommand: execCommand,
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
	ctx, cancel := context.WithTimeout(ctx, alinasCertRefreshTimeout)
	defer cancel()
	return refreshAlinasCredential(ctx, execCommand, mountPoint, cred)
}

func refreshAlinasCredential(
	ctx context.Context,
	runCommand runner,
	mountPoint string,
	cred *jwtauth.STSToken,
) error {
	// SECURITY: this JSON is the credential; never log it. The CLI prefers it here
	// precisely so it stays out of argv.
	stdin, err := json.Marshal(map[string]string{
		"ak":    cred.AccessKeyID,
		"sk":    cred.AccessKeySecret,
		"token": cred.SecurityToken,
	})
	if err != nil {
		return fmt.Errorf("marshal credential for %s: %w", alinasCertRefreshCommand, err)
	}
	output, err := runCommand(ctx, stdin, alinasCertRefreshCommand, "--mount-point", mountPoint)
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
