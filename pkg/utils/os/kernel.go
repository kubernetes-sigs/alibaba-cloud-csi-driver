package os

import (
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/sys/unix"
)

// ErrPrefixRecoveryKernel is the common prefix for recovery kernel validation errors.
const ErrPrefixRecoveryKernel = "recovery requires kernel >= 5.10.134-18 on Alibaba Cloud Linux 3, x86_64"

// KernelVersion represents a parsed kernel version string from uname.
// Example: "5.10.134-18.al8.x86_64" → {Major:5, Minor:10, Patch:134, Sublevel:18, OSDist:"al8", Arch:"x86_64"}
type KernelVersion struct {
	Major    int
	Minor    int
	Patch    int
	Sublevel int    // the numeric part after the first hyphen (e.g. 18 in "5.10.134-18")
	OSDist   string // the OS distribution tag (e.g. "al8", "el8")
	Arch     string // the architecture (e.g. "x86_64", "aarch64")
	raw      string
}

// UnameRelease calls unix.Uname and returns the release string (equivalent to "uname -r").
// This avoids spawning a subprocess via os/exec.
func UnameRelease() (string, error) {
	var uname unix.Utsname
	if err := unix.Uname(&uname); err != nil {
		return "", fmt.Errorf("uname: %w", err)
	}
	release := utsnameToString(uname.Release[:])
	return release, nil
}

// UnameMachine returns the machine hardware name (equivalent to "uname -m").
func UnameMachine() (string, error) {
	var uname unix.Utsname
	if err := unix.Uname(&uname); err != nil {
		return "", fmt.Errorf("uname: %w", err)
	}
	return utsnameToString(uname.Machine[:]), nil
}

// ParseKernelVersion parses a kernel release string like "5.10.134-18.al8.x86_64".
// It tolerates arbitrary numeric sub-segments in both the version and suffix parts
// (e.g. "5.10.134.1.2-19.3.1.al8.x86_64").
func ParseKernelVersion(release string) (*KernelVersion, error) {
	kv := &KernelVersion{raw: release}
	if release == "" {
		return nil, fmt.Errorf("empty kernel release string")
	}

	versionPart, suffix, hasSuffix := strings.Cut(release, "-")

	// Version part: take first 3 numeric segments as major.minor.patch.
	parts := strings.Split(versionPart, ".")
	if len(parts) < 3 {
		return nil, fmt.Errorf("kernel version %q: expected major.minor.patch", release)
	}
	var err error
	kv.Major, err = strconv.Atoi(parts[0])
	if err != nil {
		return nil, fmt.Errorf("kernel version %q: invalid major: %w", release, err)
	}
	kv.Minor, err = strconv.Atoi(parts[1])
	if err != nil {
		return nil, fmt.Errorf("kernel version %q: invalid minor: %w", release, err)
	}
	kv.Patch, err = strconv.Atoi(parts[2])
	if err != nil {
		return nil, fmt.Errorf("kernel version %q: invalid patch: %w", release, err)
	}

	if !hasSuffix {
		return kv, nil
	}

	// Suffix: e.g. "19.3.1.al8.x86_64" or "18.al8.x86_64"
	// Sublevel is the first numeric segment. OSDist is the first non-numeric
	// segment. Everything after OSDist is Arch.
	suffixParts := strings.Split(suffix, ".")
	if v, err := strconv.Atoi(suffixParts[0]); err == nil {
		kv.Sublevel = v
	}
	for i := 1; i < len(suffixParts); i++ {
		if _, err := strconv.Atoi(suffixParts[i]); err != nil {
			kv.OSDist = suffixParts[i]
			if i+1 < len(suffixParts) {
				kv.Arch = strings.Join(suffixParts[i+1:], ".")
			}
			break
		}
	}
	return kv, nil
}

// Less returns true if kv is strictly less than other.
// Compares by (Major, Minor, Patch, Sublevel) in order.
func (kv *KernelVersion) Less(other *KernelVersion) bool {
	if kv.Major != other.Major {
		return kv.Major < other.Major
	}
	if kv.Minor != other.Minor {
		return kv.Minor < other.Minor
	}
	if kv.Patch != other.Patch {
		return kv.Patch < other.Patch
	}
	return kv.Sublevel < other.Sublevel
}

// String returns the original release string.
func (kv *KernelVersion) String() string {
	return kv.raw
}

// CheckKernelForRecovery validates that the kernel meets the minimum requirements
// for FUSE recovery support:
//   - Kernel version >= 5.10.134-18
//   - OS distribution is Alibaba Cloud Linux 3 (al8)
//   - Architecture is x86_64
//
// Returns an error describing which requirement is not met.
func CheckKernelForRecovery() error {
	release, err := UnameRelease()
	if err != nil {
		return fmt.Errorf("%s: cannot detect kernel version: %w", ErrPrefixRecoveryKernel, err)
	}
	machine, err := UnameMachine()
	if err != nil {
		return fmt.Errorf("%s: cannot detect architecture: %w", ErrPrefixRecoveryKernel, err)
	}
	return checkKernelForRecoveryWithInputs(release, machine)
}

// checkKernelForRecoveryWithInputs validates kernel requirements using the provided
// release and machine strings. It is the testable core of CheckKernelForRecovery.
func checkKernelForRecoveryWithInputs(release, machine string) error {
	kv, err := ParseKernelVersion(release)
	if err != nil {
		return fmt.Errorf("%s: %w", ErrPrefixRecoveryKernel, err)
	}

	minVersion := &KernelVersion{Major: 5, Minor: 10, Patch: 134, Sublevel: 18}
	if kv.Less(minVersion) {
		return fmt.Errorf("%s, got %s", ErrPrefixRecoveryKernel, release)
	}

	if !isSupportedOSForRecovery(kv.OSDist) {
		return fmt.Errorf("%s: unsupported OS distribution %q (kernel %s)", ErrPrefixRecoveryKernel, kv.OSDist, release)
	}

	// TODO: support aarch64 architecture for recovery
	if machine != "x86_64" {
		return fmt.Errorf("recovery requires x86_64 architecture, got %q", machine)
	}

	return nil
}

// isSupportedOSForRecovery checks whether the OS distribution supports FUSE recovery.
// Currently only Alibaba Cloud Linux 3 (osDist "al8") is supported.
func isSupportedOSForRecovery(osDist string) bool {
	if osDist == "al8" {
		return true
	}
	// TODO: enable when ossfs2 supports Alibaba Cloud Linux 4+ (alnx4, alnx5, ...)
	// if verStr, ok := strings.CutPrefix(osDist, "alnx"); ok {
	// 	ver, err := strconv.Atoi(verStr)
	// 	return err == nil && ver >= 4
	// }
	return false
}

// utsnameToString converts a C-style char array from Utsname to a Go string.
func utsnameToString(buf []byte) string {
	var b strings.Builder
	for _, c := range buf {
		if c == 0 {
			break
		}
		b.WriteByte(c)
	}
	return b.String()
}
