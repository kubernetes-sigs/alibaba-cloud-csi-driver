package interceptors

import (
	"context"
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/jwtauth"
	mounterutils "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
	"golang.org/x/sys/unix"
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

	// delivered records that a credential has been written at least once.
	// Deliver is only ever called by one goroutine: Refresher.Start delivers
	// the initial credential synchronously and only then starts the refresh loop.
	delivered bool

	// notifyNeeded is set by Deliver when files were actually rotated (not
	// initial, not unchanged). Notify checks and clears it.
	notifyNeeded bool
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

func (s *fileCredentialSink) Deliver(cred *jwtauth.STSToken) error {
	if cred == nil {
		return fmt.Errorf("nil credential")
	}

	// Open the parent directory through openat with O_NOFOLLOW at each level,
	// creating intermediate directories as needed. The returned fd is pinned to
	// the real directory inode, so a symlink swap on the path after this call
	// does not redirect subsequent writes.
	parentFd, err := openParentDir(s.dir, customFuseCredentialBaseDir)
	if err != nil {
		return fmt.Errorf("credential path rejected: %w", err)
	}
	defer closeFd(parentFd)

	basename := filepath.Base(s.dir)
	if !s.delivered {
		// Reject a pre-existing symlink at dir itself.
		var stat unix.Stat_t
		err := unix.Fstatat(parentFd, basename, &stat, unix.AT_SYMLINK_NOFOLLOW)
		if err == nil && stat.Mode&unix.S_IFMT == unix.S_IFLNK {
			return fmt.Errorf("credential path %s is a pre-existing symlink; refusing to write through it", s.dir)
		}
	}

	// Construct a path through /proc/self/fd/<parentFd>/<basename> so that
	// rotateTokenFiles operates under the pinned directory fd, immune to
	// symlink swaps on the original path.
	pinnedDir := fmt.Sprintf("/proc/self/fd/%d/%s", parentFd, basename)

	rotated, err := rotateTokenFiles(pinnedDir, credentialSecrets(cred))
	if err != nil {
		return fmt.Errorf("write credential files to %s: %w", s.dir, err)
	}

	initial := !s.delivered
	s.delivered = true
	s.notifyNeeded = !initial && rotated

	if initial {
		klog.V(4).InfoS("delivered initial customfuse credential", "dir", s.dir, "expiration", cred.Expiration)
	} else if rotated {
		klog.V(4).InfoS("rotated customfuse credential files", "dir", s.dir, "expiration", cred.Expiration)
	} else {
		klog.V(4).InfoS("customfuse credential unchanged, skipping hook", "dir", s.dir)
	}
	return nil
}

// openParentDir opens the parent directory of dir by walking each path
// component from allowedRoot with openat(O_DIRECTORY|O_NOFOLLOW), creating
// intermediate directories as needed with mkdirat. Returns an fd pinned to
// the real parent directory inode. The caller must close the fd.
func openParentDir(dir, allowedRoot string) (int, error) {
	if err := ValidateCredentialDir(dir); err != nil {
		return -1, err
	}

	cleaned := filepath.Clean(dir)
	parent := filepath.Dir(cleaned)
	rel, err := filepath.Rel(allowedRoot, parent)
	if err != nil {
		return -1, fmt.Errorf("cannot compute relative path from %s to %s: %w", allowedRoot, parent, err)
	}

	fd, err := unix.Open(allowedRoot, unix.O_DIRECTORY|unix.O_NOFOLLOW|unix.O_CLOEXEC, 0)
	if err != nil {
		return -1, fmt.Errorf("open credential root %s: %w", allowedRoot, err)
	}

	if rel == "." {
		return fd, nil
	}

	for _, component := range strings.Split(rel, string(filepath.Separator)) {
		nextFd, err := unix.Openat(fd, component, unix.O_DIRECTORY|unix.O_NOFOLLOW|unix.O_CLOEXEC, 0)
		if err != nil {
			switch err { //nolint:errorlint // syscall errors are not wrapped
			case unix.ENOENT:
				if mkErr := unix.Mkdirat(fd, component, 0o755); mkErr != nil && mkErr != unix.EEXIST {
					closeFd(fd)
					return -1, fmt.Errorf("mkdirat %s: %w", component, mkErr)
				}
				nextFd, err = unix.Openat(fd, component, unix.O_DIRECTORY|unix.O_NOFOLLOW|unix.O_CLOEXEC, 0)
				if err != nil {
					closeFd(fd)
					return -1, fmt.Errorf("openat %s after mkdir: %w", component, err)
				}
			case unix.ELOOP, unix.ENOTDIR:
				closeFd(fd)
				return -1, fmt.Errorf("symlink at %s in credential path; refusing to write through a symlinked parent", component)
			default:
				closeFd(fd)
				return -1, fmt.Errorf("openat %s: %w", component, err)
			}
		}
		closeFd(fd)
		fd = nextFd
	}
	return fd, nil
}

func closeFd(fd int) {
	_ = unix.Close(fd)
}

func (s *fileCredentialSink) Notify() error {
	if !s.notifyNeeded {
		return nil
	}
	if err := s.runRefreshHook(); err != nil {
		return err
	}
	s.notifyNeeded = false
	return nil
}

// Cleanup removes the credential files and the symlink. It runs after the
// refresher has stopped and the mount is gone, so nothing is left readable
// once there is no client to serve.
func (s *fileCredentialSink) Cleanup() {
	if !s.delivered {
		return
	}
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
