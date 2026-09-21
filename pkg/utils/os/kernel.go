package os

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/sys/unix"
)

// fuseFlushSymbol is the kernel symbol exposed by the alinux FUSE recovery patch.
// Its presence in /proc/kallsyms indicates the kernel supports flushing in-flight
// FUSE requests via /sys/fs/fuse/connections/<id>/flush, which is required for
// ossfs2 crash recovery. See also: bmcpfs checks fuse_dev_ioctl_recover from the
// same patch set.
const fuseFlushSymbol = "fuse_flush_pq"

// procKallsymsPath is the default path to the kernel symbol table.
// Overridden in tests.
var procKallsymsPath = "/proc/kallsyms"

// KernelVersion represents a parsed kernel version string from uname.
// Example: "5.10.134-18.al8.x86_64" → {Major:5, Minor:10, Patch:134, Sublevel:18, OSDist:"al8"}
// Architecture is not parsed here; use UnameMachine() (uname -m) instead.
type KernelVersion struct {
	Major    int
	Minor    int
	Patch    int
	Sublevel int    // the numeric part after the first hyphen (e.g. 18 in "5.10.134-18")
	OSDist   string // the OS distribution tag (e.g. "al8", "el8")
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

// CheckKernelForRecovery checks whether the running kernel supports FUSE
// connection flush, required for ossfs2 crash recovery.
//
// Detection is capability-based: we scan /proc/kallsyms for the fuse_flush_pq
// symbol rather than inferring support from kernel version/OS/arch. This
// matches the approach used by bmcpfs (which checks fuse_dev_ioctl_recover).
func CheckKernelForRecovery() error {
	return checkKallsymsForSymbol(procKallsymsPath, fuseFlushSymbol)
}

// checkKallsymsForSymbol scans a kallsyms-format file for a symbol name.
// kallsyms format: "<addr> <type> <name>\n" or "<addr> <type> <name>\t[<module>]\n"
func checkKallsymsForSymbol(path, symbol string) error {
	f, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("cannot check kernel symbols: %w", err)
	}
	defer func() { _ = f.Close() }()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) >= 3 && fields[2] == symbol {
			return nil
		}
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading %s: %w", path, err)
	}
	return fmt.Errorf("kernel symbol %q not found in %s; FUSE recovery flush not supported by this kernel", symbol, path)
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
