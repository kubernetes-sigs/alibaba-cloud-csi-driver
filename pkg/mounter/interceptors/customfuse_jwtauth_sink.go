package interceptors

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/jwtauth"
	mounterutils "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
	"k8s.io/klog/v2"
)

const (
	// credentialRefreshHookTimeout bounds a single hook invocation so a hung
	// hook cannot stall the refresh loop or Stop.
	credentialRefreshHookTimeout = 30 * time.Second

	// credentialDirEnvKey is how the refresh hook learns where the credential
	// lives. It matches the mount option name, which the driver turns into an
	// environment variable of the same name, so a hook and the entrypoint that
	// mounted read the identical path from the identical variable.
	credentialDirEnvKey = OptCredentialDir
)

// fileCredentialSink delivers a credential by writing one file per field into a
// directory, and optionally running a hook afterwards.
//
// The driver cannot know what shape its client wants a credential in, and no
// shape is neutral: every concrete format is some vendor's dialect, down to
// disagreeing on the field names for the same three values. So the credential is
// written as the smallest useful primitive, one value per file, and the
// entrypoint composes whatever its client actually reads — a JSON document with
// either spelling of the field names, an ini profile, a properties file, command
// line arguments. Composing from separate values needs no parsing; going the
// other way would force every entrypoint to take apart a format it did not want.
//
// Two things the entrypoint cannot do for itself, and this therefore must:
//
// Deliver before the client starts. The entrypoint reads the credential as it
// comes up, so it has to already be there.
//
// Signal that a rotation happened. A client that re-reads the files needs
// nothing. A client whose credential lives elsewhere — a metadata store, an
// already-parsed in-memory config — cannot observe a file changing at all and
// has to be told, which is what the hook is for. The hook is an arbitrary
// executable precisely so that it can express whatever telling this client
// means: invoking its CLI, signalling it, calling its admin endpoint, rewriting
// its own config. A client that can do neither cannot be served by any delivery
// mechanism; it needs a credential outliving the mount.
type fileCredentialSink struct {
	// dir is the symlink path the client and the hook are pointed at. Its
	// parent must exist; rotateTokenFiles creates the data directories next to
	// it and swaps the symlink atomically.
	dir string

	// hook is an absolute path to an executable run after each successful
	// write, or empty to run nothing. No credential value ever reaches it, as an
	// argument or as an environment value: it is told where to look and reads the
	// files itself, so nothing lands in /proc/<pid>/cmdline or in the environment
	// of an unrelated process.
	hook string

	// env is the environment the hook runs in, captured at mount time: the
	// daemon's own plus this volume's attributes, one key=value each. A hook runs
	// long after the mount call returned, so this is its only way to learn which
	// volume rotated and where that volume came from — a client that has to be
	// told about a rotation has to be told about a specific one. It carries no
	// credential value, and no Secret either; see customFuseHookEnv.
	env []string

	// runHook runs the hook and returns its combined output. It is a field so
	// tests can inject a fake runner.
	runHook func(ctx context.Context, name string, env []string) ([]byte, error)

	// delivered records that a credential has been written at least once, which
	// is what separates the initial delivery from a rotation. The hook tells a
	// *running* client that its credential changed, and on the initial delivery
	// there is no client yet — it is written so the entrypoint can read it as it
	// starts. Running the hook then would at best do nothing, and at worst fail
	// against a client that is not up, failing a mount that was fine.
	//
	// Apply is only ever called by one goroutine: Refresher.Start applies the
	// initial credential synchronously and only then starts the refresh loop.
	delivered bool
}

var _ jwtauth.CredentialSink = &fileCredentialSink{}

func newFileCredentialSink(dir, hook string, env []string) *fileCredentialSink {
	return &fileCredentialSink{
		dir:  dir,
		hook: hook,
		env:  env,
		runHook: func(ctx context.Context, name string, env []string) ([]byte, error) {
			cmd := exec.CommandContext(ctx, name)
			cmd.Env = env
			return cmd.CombinedOutput()
		},
	}
}

func (s *fileCredentialSink) Apply(cred *jwtauth.STSToken) error {
	if cred == nil {
		return fmt.Errorf("nil credential")
	}

	// rotateTokenFiles writes the data directory next to the symlink, so the
	// parent has to exist before the first rotation. The swap is a symlink
	// rename, atomic per path lookup rather than per credential: a client
	// reading the fields as separate lookups can in principle straddle it. A
	// client is expected to verify a credential before relying on it, which it
	// needs regardless for expiry and provider errors.
	if err := os.MkdirAll(filepath.Dir(s.dir), 0o755); err != nil {
		return fmt.Errorf("create credential parent directory: %w", err)
	}

	rotated, err := rotateTokenFiles(s.dir, credentialSecrets(cred))
	if err != nil {
		return fmt.Errorf("write credential files to %s: %w", s.dir, err)
	}

	initial := !s.delivered
	s.delivered = true

	if initial {
		klog.V(4).InfoS("delivered initial customfuse credential", "dir", s.dir, "expiration", cred.Expiration)
		return nil
	}
	if !rotated {
		// Same credential as the one already on disk; the client has nothing to
		// pick up, so do not disturb it with a hook.
		klog.V(4).InfoS("customfuse credential unchanged, skipping hook", "dir", s.dir)
		return nil
	}
	klog.V(4).InfoS("rotated customfuse credential files", "dir", s.dir, "expiration", cred.Expiration)

	return s.runRefreshHook()
}

// Cleanup removes the credential files and the symlink. It runs after the
// refresher has stopped and the mount is gone, so nothing is left readable
// once there is no client to serve.
func (s *fileCredentialSink) Cleanup() {
	cleanupTokenFiles(s.dir)
}

// runRefreshHook runs the configured hook, if any. A hook failure is returned so
// the refresh loop reports it: the files are already updated, but a client that
// depends on the hook has not picked them up, which is a half-applied rotation
// and must not pass silently.
func (s *fileCredentialSink) runRefreshHook() error {
	if s.hook == "" {
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), credentialRefreshHookTimeout)
	defer cancel()

	// The hook gets the environment recorded at mount time plus the credential
	// directory. It is told where to look, never what the credential is. The copy
	// matters: s.env is reused on every rotation, so appending into its spare
	// capacity would write through to it.
	env := make([]string, 0, len(s.env)+1)
	env = append(env, s.env...)
	env = append(env, credentialDirEnvKey+"="+s.dir)
	output, err := s.runHook(ctx, s.hook, env)
	if err != nil {
		return fmt.Errorf("credential refresh hook %s failed: %w, output: %s", s.hook, err, string(output))
	}
	klog.V(4).InfoS("ran customfuse credential refresh hook", "hook", s.hook, "dir", s.dir)
	return nil
}

// credentialSecrets keys the credential by the driver-wide names, which is what
// rotateTokenFiles expects and what the other drivers' rotation already writes.
func credentialSecrets(cred *jwtauth.STSToken) map[string]string {
	return map[string]string{
		mounterutils.KeyAccessKeyId:     cred.AccessKeyID,
		mounterutils.KeyAccessKeySecret: cred.AccessKeySecret,
		mounterutils.KeySecurityToken:   cred.SecurityToken,
		mounterutils.KeyExpiration:      cred.Expiration,
	}
}
