package interceptors

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"syscall"
	"time"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/jwtauth"
	"k8s.io/klog/v2"
)

const (
	credentialRefreshHookTimeout = 30 * time.Second
)

// stdinCredentialSink delivers credentials to a hook script via stdin.
// Each Apply runs the hook with the credential written as four lines
// (AccessKeyId, AccessKeySecret, SecurityToken, Expiration) to its stdin.
// The hook decides what to do: write files, call a CLI, signal a process.
//
// The initial Apply (before the entrypoint starts) is skipped because the
// entrypoint receives its first credential through its own stdin, set up
// by the driver's ExtendedMount. Only subsequent rotations run the hook.
type stdinCredentialSink struct {
	hook string
	env  []string

	runHook func(ctx context.Context, name string, env []string, stdin []byte) ([]byte, error)

	applied bool // true after the first Apply
}

var _ jwtauth.CredentialSink = &stdinCredentialSink{}

func newStdinCredentialSink(hook string, env []string) *stdinCredentialSink {
	return &stdinCredentialSink{
		hook:    hook,
		env:     env,
		runHook: runHookWithStdin,
	}
}

func (s *stdinCredentialSink) Apply(cred *jwtauth.STSToken) error {
	if cred == nil {
		return fmt.Errorf("nil credential")
	}

	if !s.applied {
		s.applied = true
		klog.V(4).InfoS("skipped initial hook (entrypoint reads stdin)", "expiration", cred.Expiration)
		return nil
	}

	if s.hook == "" {
		klog.V(4).InfoS("no hook configured, credential rotated silently", "expiration", cred.Expiration)
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), credentialRefreshHookTimeout)
	defer cancel()

	stdin := formatCredential(cred)
	env := make([]string, 0, len(s.env))
	env = append(env, s.env...)

	output, err := s.runHook(ctx, s.hook, env, stdin)
	if err != nil {
		return fmt.Errorf("credential refresh hook %s failed: %w, output: %s", s.hook, err, string(output))
	}
	klog.V(4).InfoS("ran customfuse credential refresh hook", "hook", s.hook, "expiration", cred.Expiration)
	return nil
}

func (s *stdinCredentialSink) Cleanup() {}

// formatCredential produces the stdin payload: four lines, one field per line.
// Entrypoint and hook scripts read them with:
//
//	read -r AccessKeyId
//	read -r AccessKeySecret
//	read -r SecurityToken
//	read -r Expiration
func formatCredential(cred *jwtauth.STSToken) []byte {
	return []byte(cred.AccessKeyID + "\n" +
		cred.AccessKeySecret + "\n" +
		cred.SecurityToken + "\n" +
		cred.Expiration + "\n")
}

// runHookWithStdin runs a hook in its own process group with credential data
// on stdin. Context cancellation kills the entire process group.
func runHookWithStdin(ctx context.Context, name string, env []string, stdin []byte) ([]byte, error) {
	cmd := exec.CommandContext(ctx, name)
	cmd.Env = env
	cmd.Stdin = bytes.NewReader(stdin)
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}

	var buf bytes.Buffer
	cmd.Stdout = &buf
	cmd.Stderr = &buf

	if err := cmd.Start(); err != nil {
		return nil, err
	}

	done := make(chan error, 1)
	go func() { done <- cmd.Wait() }()

	select {
	case err := <-done:
		return buf.Bytes(), err
	case <-ctx.Done():
		_ = syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)
		<-done
		return buf.Bytes(), ctx.Err()
	}
}
