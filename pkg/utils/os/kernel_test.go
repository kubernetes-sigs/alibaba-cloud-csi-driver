package os

import (
	"testing"
)

func TestParseKernelVersion(t *testing.T) {
	tests := []struct {
		name       string
		release    string
		wantMajor  int
		wantMinor  int
		wantPatch  int
		wantSub    int
		wantOSDist string
		wantArch   string
		wantErr    bool
	}{
		{
			name:       "Alibaba Cloud Linux 2 x86_64",
			release:    "5.10.134-18.al8.x86_64",
			wantMajor:  5,
			wantMinor:  10,
			wantPatch:  134,
			wantSub:    18,
			wantOSDist: "al8",
			wantArch:   "x86_64",
		},
		{
			name:       "Alibaba Cloud Linux 3 newer sublevel",
			release:    "5.10.134-19.al8.x86_64",
			wantMajor:  5,
			wantMinor:  10,
			wantPatch:  134,
			wantSub:    19,
			wantOSDist: "al8",
			wantArch:   "x86_64",
		},
		{
			name:       "Alibaba Cloud Linux aarch64",
			release:    "5.10.134-18.al8.aarch64",
			wantMajor:  5,
			wantMinor:  10,
			wantPatch:  134,
			wantSub:    18,
			wantOSDist: "al8",
			wantArch:   "aarch64",
		},
		{
			name:       "CentOS 8",
			release:    "5.10.134-18.el8.x86_64",
			wantMajor:  5,
			wantMinor:  10,
			wantPatch:  134,
			wantSub:    18,
			wantOSDist: "el8",
			wantArch:   "x86_64",
		},
		{
			name:       "higher kernel version",
			release:    "5.10.135-20.al8.x86_64",
			wantMajor:  5,
			wantMinor:  10,
			wantPatch:  135,
			wantSub:    20,
			wantOSDist: "al8",
			wantArch:   "x86_64",
		},
		{
			name:       "Alibaba Cloud Linux 4 x86_64",
			release:    "6.6.102-5.alnx4.x86_64",
			wantMajor:  6,
			wantMinor:  6,
			wantPatch:  102,
			wantSub:    5,
			wantOSDist: "alnx4",
			wantArch:   "x86_64",
		},
		{
			name:       "kernel 6.x",
			release:    "6.6.0-1.al8.x86_64",
			wantMajor:  6,
			wantMinor:  6,
			wantPatch:  0,
			wantSub:    1,
			wantOSDist: "al8",
			wantArch:   "x86_64",
		},
		{
			name:       "multi-segment sublevel",
			release:    "5.10.134-19.3.1.al8.x86_64",
			wantMajor:  5,
			wantMinor:  10,
			wantPatch:  134,
			wantSub:    19,
			wantOSDist: "al8",
			wantArch:   "x86_64",
		},
		{
			name:       "extra version and sublevel segments",
			release:    "5.10.134.1.2-19.3.1.al8.x86_64",
			wantMajor:  5,
			wantMinor:  10,
			wantPatch:  134,
			wantSub:    19,
			wantOSDist: "al8",
			wantArch:   "x86_64",
		},
		{
			name:      "no suffix",
			release:   "5.10.134",
			wantMajor: 5,
			wantMinor: 10,
			wantPatch: 134,
		},
		{
			name:    "empty string",
			release: "",
			wantErr: true,
		},
		{
			name:    "missing patch",
			release: "5.10",
			wantErr: true,
		},
		{
			name:    "invalid major",
			release: "abc.10.134-18.al8.x86_64",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			kv, err := ParseKernelVersion(tt.release)
			if tt.wantErr {
				if err == nil {
					t.Errorf("ParseKernelVersion(%q) expected error, got nil", tt.release)
				}
				return
			}
			if err != nil {
				t.Fatalf("ParseKernelVersion(%q) unexpected error: %v", tt.release, err)
			}
			if kv.Major != tt.wantMajor || kv.Minor != tt.wantMinor || kv.Patch != tt.wantPatch {
				t.Errorf("version = %d.%d.%d, want %d.%d.%d", kv.Major, kv.Minor, kv.Patch, tt.wantMajor, tt.wantMinor, tt.wantPatch)
			}
			if kv.Sublevel != tt.wantSub {
				t.Errorf("sublevel = %d, want %d", kv.Sublevel, tt.wantSub)
			}
			if kv.OSDist != tt.wantOSDist {
				t.Errorf("osdist = %q, want %q", kv.OSDist, tt.wantOSDist)
			}
			if kv.Arch != tt.wantArch {
				t.Errorf("arch = %q, want %q", kv.Arch, tt.wantArch)
			}
			if kv.String() != tt.release {
				t.Errorf("String() = %q, want %q", kv.String(), tt.release)
			}
		})
	}
}

func TestKernelVersionLess(t *testing.T) {
	tests := []struct {
		name   string
		a      string
		b      string
		aLessB bool
	}{
		{"same version", "5.10.134-18.al8.x86_64", "5.10.134-18.al8.x86_64", false},
		{"higher patch", "5.10.134-18.al8.x86_64", "5.10.135-18.al8.x86_64", true},
		{"higher sublevel", "5.10.134-17.al8.x86_64", "5.10.134-18.al8.x86_64", true},
		{"higher minor", "5.9.134-18.al8.x86_64", "5.10.134-18.al8.x86_64", true},
		{"higher major", "4.10.134-18.al8.x86_64", "5.10.134-18.al8.x86_64", true},
		{"equal different dist", "5.10.134-18.al8.x86_64", "5.10.134-18.el8.x86_64", false},
		{"kernel 6 > kernel 5", "5.10.134-18.al8.x86_64", "6.0.0-1.al8.x86_64", true},
		{"patch beats sublevel", "5.10.134-99.al8.x86_64", "5.10.135-1.al8.x86_64", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a, _ := ParseKernelVersion(tt.a)
			b, _ := ParseKernelVersion(tt.b)
			if got := a.Less(b); got != tt.aLessB {
				t.Errorf("Less(%q, %q) = %v, want %v", tt.a, tt.b, got, tt.aLessB)
			}
		})
	}
}

func TestCheckKernelForRecovery_ValidationError(t *testing.T) {
	tests := []struct {
		name    string
		release string
		machine string
		wantErr bool
	}{
		{"kernel too old - sublevel", "5.10.134-17.al8.x86_64", "x86_64", true},
		{"kernel too old - patch", "5.10.133-18.al8.x86_64", "x86_64", true},
		{"kernel too old - minor", "5.9.134-18.al8.x86_64", "x86_64", true},
		{"not alinux - el8", "5.10.134-18.el8.x86_64", "x86_64", true},
		{"not alinux - al7", "5.10.134-18.al7.x86_64", "x86_64", true},
		{"wrong arch - s390x", "5.10.134-18.al8.x86_64", "s390x", true},
		{"wrong arch - aarch64", "5.10.134-18.al8.aarch64", "aarch64", true},
		{"wrong arch - i386", "5.10.134-18.al8.i386", "i386", true},
		{"wrong arch - armv7l", "5.10.134-18.al8.armv7l", "armv7l", true},
		{"wrong arch - ppc64le", "5.10.134-18.al8.ppc64le", "ppc64le", true},
		{"wrong arch - empty", "5.10.134-18.al8", "", true},
		{"valid al8 x86_64", "5.10.134-18.al8.x86_64", "x86_64", false},
		{"valid higher version", "5.10.135-20.al8.x86_64", "x86_64", false},
		{"valid kernel 6", "6.6.0-1.al8.x86_64", "x86_64", false},
		{"valid alnx4 x86_64", "6.6.102-5.alnx4.x86_64", "x86_64", true}, // TODO: change to false when ossfs2 supports alinux4+
		{"valid alnx5 x86_64", "6.6.102-5.alnx5.x86_64", "x86_64", true}, // TODO: change to false when ossfs2 supports alinux4+
		{"valid multi-segment sublevel", "5.10.134-19.3.1.al8.x86_64", "x86_64", false},
		{"not alinux - al9", "5.10.134-18.al9.x86_64", "x86_64", true},
		{"not alinux - alnx3", "5.10.134-18.alnx3.x86_64", "x86_64", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := checkKernelForRecoveryWithInputs(tt.release, tt.machine)
			gotErr := err != nil
			if gotErr != tt.wantErr {
				t.Errorf("checkKernelForRecoveryWithInputs(%q, %q) error=%v, wantErr=%v", tt.release, tt.machine, err, tt.wantErr)
			}
		})
	}
}

func TestIsSupportedOSForRecovery(t *testing.T) {
	tests := []struct {
		osDist string
		want   bool
	}{
		{"al8", true},
		{"alnx4", false}, // TODO: change to true when ossfs2 supports alinux4+
		{"alnx5", false},
		{"alnx10", false},
		{"al7", false},
		{"al9", false},
		{"al10", false},
		{"al1", false},
		{"alnx3", false},
		{"alnx", false},
		{"el8", false},
		{"al", false},
		{"alx", false},
		{"", false},
	}

	for _, tt := range tests {
		t.Run(tt.osDist, func(t *testing.T) {
			if got := isSupportedOSForRecovery(tt.osDist); got != tt.want {
				t.Errorf("isSupportedOSForRecovery(%q) = %v, want %v", tt.osDist, got, tt.want)
			}
		})
	}
}
