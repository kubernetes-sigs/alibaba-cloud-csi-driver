/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package utils

import (
	"os"
	"path/filepath"
	"syscall"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetAccessControl(t *testing.T) {
	testAccessKey := "testkey"
	testAccessKeySecret := "testvalue"
	t.Setenv("ACCESS_KEY_ID", testAccessKey)
	t.Setenv("ACCESS_KEY_SECRET", testAccessKeySecret)
	ac := GetAccessControl()
	assert.Equal(t, testAccessKey, ac.AccessKeyID)
	assert.Equal(t, testAccessKeySecret, ac.AccessKeySecret)
	assert.Empty(t, ac.StsToken)
	t.Setenv("ACCESS_KEY_ID", "")
	t.Setenv("ACCESS_KEY_SECRET", "")
	ac = GetAccessControl()
	assert.Empty(t, ac.AccessKeyID)
	assert.Empty(t, ac.AccessKeySecret)
	assert.Empty(t, ac.StsToken)
}

func TestValidatePath(t *testing.T) {
	// Resolve the temp dir so expected paths match what filepath.EvalSymlinks
	// returns (on macOS /var is a symlink to /private/var).
	base, err := filepath.EvalSymlinks(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	// An existing nested directory to act as a valid, resolvable parent.
	existingDir := filepath.Join(base, "pods", "uid", "volumes")
	if err := os.MkdirAll(existingDir, 0o755); err != nil {
		t.Fatal(err)
	}
	// A directory whose name contains "proc" to verify it is not treated as /proc.
	procLikeDir := filepath.Join(base, "proc-data")
	if err := os.MkdirAll(procLikeDir, 0o755); err != nil {
		t.Fatal(err)
	}
	// A regular file to verify that resolution errors other than non-existence
	// are propagated instead of triggering the parent recursion.
	regularFile := filepath.Join(t.TempDir(), "regular-file")
	if err := os.WriteFile(regularFile, []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name    string
		path    string
		wantOK  bool
		wantErr string
	}{
		{
			name:   "existing directory",
			path:   existingDir,
			wantOK: true,
		},
		{
			name:   "non-existent leaf under existing parent",
			path:   filepath.Join(existingDir, "mount"),
			wantOK: true,
		},
		{
			name:   "base directory itself",
			path:   base,
			wantOK: true,
		},
		// ".." (or any non-canonical form) is rejected up front: the kernel
		// resolves ".." after following symlinks, so tolerating it would let a
		// symlinked component redirect the mount target away from the
		// lexically-cleaned path that is validated.
		{
			// Note: string concatenation, not filepath.Join, which would
			// clean the path and hide the traversal.
			name:    "traversal within existing tree is rejected",
			path:    existingDir + "/../..",
			wantOK:  false,
			wantErr: "contains '..'",
		},
		{
			name:    "traversal crossing a potential symlink component is rejected",
			path:    "/run/csi-mount/root/../usr/local/bin",
			wantOK:  false,
			wantErr: "contains '..'",
		},
		{
			name:   "trailing slash is allowed",
			path:   existingDir + "/",
			wantOK: true,
		},
		{
			name:   "proc as non-prefix component is fine",
			path:   filepath.Join(procLikeDir, "file"),
			wantOK: true,
		},

		// Hidden-directory-style components (".agent", "..agent") are ordinary
		// path components: filepath.Clean only collapses exact "." and ".."
		// components, so the canonical requirement must not filter legitimate
		// cluster paths like /.agent/xx.
		{
			name:   "hidden root component is allowed",
			path:   "/.agent/xx",
			wantOK: true,
		},
		{
			name:   "hidden component under existing tree is allowed",
			path:   filepath.Join(existingDir, ".agent", "xx"),
			wantOK: true,
		},
		{
			name:   "dot-dot-prefixed component is not traversal",
			path:   filepath.Join(existingDir, "..agent", "xx"),
			wantOK: true,
		},
		{
			// A "." component is harmless (kernel resolves it identically),
			// so it is allowed. Note: string concatenation, not filepath.Join.
			name:   "literal dot component is allowed",
			path:   existingDir + "/./mount",
			wantOK: true,
		},

		// Relative path is rejected before any further processing.
		{
			name:    "relative path rejected",
			path:    "var/lib/kubelet/pods/uid",
			wantOK:  false,
			wantErr: "must be an absolute path",
		},

		// Sensitive system path prefix (/proc), caught on the literal path.
		{
			name:    "proc path rejected",
			path:    "/proc/self/fd/0",
			wantOK:  false,
			wantErr: "under sensitive path /proc",
		},
		{
			name:    "proc root rejected",
			path:    "/proc",
			wantOK:  false,
			wantErr: "under sensitive path /proc",
		},

		// Multiple non-existent trailing components are resolved by walking up
		// to the deepest existing ancestor, so the path is still validated.
		{
			name:   "multiple non-existent components resolve via ancestor",
			path:   filepath.Join(base, "missing-parent", "leaf"),
			wantOK: true,
		},

		// A path through a regular file fails resolution with a non-ENOENT
		// error, which is propagated instead of recursing into the parent.
		{
			name:    "path through regular file returns error",
			path:    filepath.Join(regularFile, "leaf"),
			wantOK:  false,
			wantErr: "failed to resolve symlinks",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ok, err := ValidatePath(tt.path)
			assert.Equal(t, tt.wantOK, ok)
			if tt.wantErr != "" {
				assert.ErrorContains(t, err, tt.wantErr)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestValidatePath_PATHDirs(t *testing.T) {
	base, err := filepath.EvalSymlinks(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	binDir := filepath.Join(base, "bin")
	otherDir := filepath.Join(base, "data")
	prefixDir := filepath.Join(base, "binary") // shares a prefix with binDir
	for _, d := range []string{binDir, otherDir, prefixDir} {
		if err := os.MkdirAll(d, 0o755); err != nil {
			t.Fatal(err)
		}
	}

	t.Setenv("PATH", binDir)

	// Path under a PATH directory is rejected (leaf need not exist).
	ok, err := ValidatePath(filepath.Join(binDir, "ls"))
	assert.False(t, ok)
	assert.ErrorContains(t, err, "under PATH directory")

	// PATH directory itself is rejected.
	ok, err = ValidatePath(binDir)
	assert.False(t, ok)
	assert.ErrorContains(t, err, "under PATH directory")

	// Path not under any PATH directory is allowed.
	ok, err = ValidatePath(filepath.Join(otherDir, "file"))
	assert.True(t, ok)
	assert.NoError(t, err)

	// Path that shares a prefix but is not actually under a PATH dir.
	ok, err = ValidatePath(filepath.Join(prefixDir, "file"))
	assert.True(t, ok)
	assert.NoError(t, err)
}

func TestValidatePath_Symlinks(t *testing.T) {
	base, err := filepath.EvalSymlinks(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	sensitiveDir := filepath.Join(base, "sensitive")
	safeDir := filepath.Join(base, "safe")
	for _, d := range []string{sensitiveDir, safeDir} {
		if err := os.MkdirAll(d, 0o755); err != nil {
			t.Fatal(err)
		}
	}

	t.Setenv("PATH", sensitiveDir)

	// A symlinked parent that resolves into a PATH dir must be rejected, even
	// though the literal path is not under the PATH dir. The leaf does not exist.
	parentLink := filepath.Join(base, "parent-link")
	if err := os.Symlink(sensitiveDir, parentLink); err != nil {
		t.Fatal(err)
	}
	ok, err := ValidatePath(filepath.Join(parentLink, "leaf"))
	assert.False(t, ok)
	assert.ErrorContains(t, err, "under PATH directory")

	// A multi-level missing path through a symlinked ancestor must also be
	// rejected: the recursion resolves the ancestor and re-appends the full
	// non-existent suffix before the PATH check runs.
	ok, err = ValidatePath(filepath.Join(parentLink, "missing1", "missing2"))
	assert.False(t, ok)
	assert.ErrorContains(t, err, "under PATH directory")

	// A leaf that is itself a symlink into a PATH dir must be rejected.
	leafLink := filepath.Join(base, "leaf-link")
	if err := os.Symlink(sensitiveDir, leafLink); err != nil {
		t.Fatal(err)
	}
	ok, err = ValidatePath(leafLink)
	assert.False(t, ok)
	assert.ErrorContains(t, err, "under PATH directory")

	// A symlinked parent that resolves into a safe dir is allowed.
	safeLink := filepath.Join(base, "safe-link")
	if err := os.Symlink(safeDir, safeLink); err != nil {
		t.Fatal(err)
	}
	ok, err = ValidatePath(filepath.Join(safeLink, "leaf"))
	assert.True(t, ok)
	assert.NoError(t, err)
}

func TestValidatePath_SymlinkIntoProc(t *testing.T) {
	if _, err := os.Stat("/proc"); err != nil {
		t.Skip("/proc not available on this platform")
	}
	base := t.TempDir()

	// A symlink whose target resolves into /proc must be rejected, even though
	// the literal path is not under /proc.
	procLink := filepath.Join(base, "proc-link")
	if err := os.Symlink("/proc", procLink); err != nil {
		t.Fatal(err)
	}
	ok, err := ValidatePath(filepath.Join(procLink, "self"))
	assert.False(t, ok)
	assert.ErrorContains(t, err, "under sensitive path /proc")
}

func TestResolveSymlinks(t *testing.T) {
	base, err := filepath.EvalSymlinks(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	existing := filepath.Join(base, "existing")
	if err := os.MkdirAll(existing, 0o755); err != nil {
		t.Fatal(err)
	}

	// An existing path without any symlink resolves to itself.
	got, err := resolveSymlinks(existing)
	assert.NoError(t, err)
	assert.Equal(t, existing, got)

	// A symlinked ancestor is resolved and the full non-existent suffix is
	// re-appended: link -> existing, so link/a/b resolves to existing/a/b.
	link := filepath.Join(base, "link")
	if err := os.Symlink(existing, link); err != nil {
		t.Fatal(err)
	}
	got, err = resolveSymlinks(filepath.Join(link, "a", "b"))
	assert.NoError(t, err)
	assert.Equal(t, filepath.Join(existing, "a", "b"), got)

	// A leaf that is itself a symlink resolves to its target.
	got, err = resolveSymlinks(link)
	assert.NoError(t, err)
	assert.Equal(t, existing, got)

	// A path through a regular file fails with a non-ENOENT error, which must
	// be propagated instead of triggering the parent recursion.
	regularFile := filepath.Join(t.TempDir(), "file")
	if err := os.WriteFile(regularFile, []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	got, err = resolveSymlinks(filepath.Join(regularFile, "leaf"))
	assert.Error(t, err)
	assert.Empty(t, got)
}

// A mount point whose filesystem is dead must not fail validation: both
// NodePublishVolume and NodeUnpublishVolume resolve the target path first, so the
// driver could neither re-establish nor clean up such a mount.
func TestResolveSymlinks_DeadMountPoint(t *testing.T) {
	// Resolve the temp dir up front: on macOS /var is a symlink to /private/var.
	base, err := filepath.EvalSymlinks(t.TempDir())
	require.NoError(t, err)
	dead := filepath.Join(base, "deadmount")
	require.NoError(t, os.MkdirAll(dead, 0o755))

	orig := evalSymlinks
	defer func() { evalSymlinks = orig }()
	failWith := func(target string, errno syscall.Errno) {
		evalSymlinks = func(path string) (string, error) {
			if path == target {
				return "", &os.PathError{Op: "lstat", Path: path, Err: errno}
			}
			return orig(path)
		}
	}

	// Dead leaf: resolve the healthy parent and re-append the base name.
	failWith(dead, syscall.ENOTCONN)
	got, err := resolveSymlinks(dead)
	assert.NoError(t, err)
	assert.Equal(t, dead, got)

	// A mount point is always the leaf, so an ENOTCONN ancestor is not tolerated:
	// its contents are unreadable and the suffix could hide a symlink.
	failWith(dead, syscall.ENOTCONN)
	got, err = resolveSymlinks(filepath.Join(dead, "leaf"))
	assert.Error(t, err)
	assert.Empty(t, got)

	// Errors IsCorruptedMnt would accept but that are not the crashed-daemon case
	// stay hard failures.
	for _, errno := range []syscall.Errno{syscall.ESTALE, syscall.EACCES, syscall.EIO, syscall.EPERM} {
		failWith(dead, errno)
		got, err = resolveSymlinks(dead)
		assert.Error(t, err, "errno %v must not be resolved around", errno)
		assert.Empty(t, got)
	}
}
