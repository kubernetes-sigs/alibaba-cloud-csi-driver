package os

import (
	"os"
	"path/filepath"
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

		wantErr bool
	}{
		{
			name:       "Alibaba Cloud Linux 2 x86_64",
			release:    "5.10.134-18.al8.x86_64",
			wantMajor:  5,
			wantMinor:  10,
			wantPatch:  134,
			wantSub:    18,
			wantOSDist: "al8",
		},
		{
			name:       "Alibaba Cloud Linux 3 newer sublevel",
			release:    "5.10.134-19.al8.x86_64",
			wantMajor:  5,
			wantMinor:  10,
			wantPatch:  134,
			wantSub:    19,
			wantOSDist: "al8",
		},
		{
			name:       "Alibaba Cloud Linux aarch64",
			release:    "5.10.134-18.al8.aarch64",
			wantMajor:  5,
			wantMinor:  10,
			wantPatch:  134,
			wantSub:    18,
			wantOSDist: "al8",
		},
		{
			name:       "CentOS 8",
			release:    "5.10.134-18.el8.x86_64",
			wantMajor:  5,
			wantMinor:  10,
			wantPatch:  134,
			wantSub:    18,
			wantOSDist: "el8",
		},
		{
			name:       "higher kernel version",
			release:    "5.10.135-20.al8.x86_64",
			wantMajor:  5,
			wantMinor:  10,
			wantPatch:  135,
			wantSub:    20,
			wantOSDist: "al8",
		},
		{
			name:       "Alibaba Cloud Linux 4 x86_64",
			release:    "6.6.102-5.alnx4.x86_64",
			wantMajor:  6,
			wantMinor:  6,
			wantPatch:  102,
			wantSub:    5,
			wantOSDist: "alnx4",
		},
		{
			name:       "kernel 6.x",
			release:    "6.6.0-1.al8.x86_64",
			wantMajor:  6,
			wantMinor:  6,
			wantPatch:  0,
			wantSub:    1,
			wantOSDist: "al8",
		},
		{
			name:       "multi-segment sublevel",
			release:    "5.10.134-19.3.1.al8.x86_64",
			wantMajor:  5,
			wantMinor:  10,
			wantPatch:  134,
			wantSub:    19,
			wantOSDist: "al8",
		},
		{
			name:       "extra version and sublevel segments",
			release:    "5.10.134.1.2-19.3.1.al8.x86_64",
			wantMajor:  5,
			wantMinor:  10,
			wantPatch:  134,
			wantSub:    19,
			wantOSDist: "al8",
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

func TestCheckKallsymsForSymbol(t *testing.T) {
	tests := []struct {
		name    string
		content string
		symbol  string
		wantErr bool
	}{
		{
			name:    "symbol present in fuse module",
			content: "ffffffffc0b20ed0 t fuse_dev_ioctl_recover\t[fuse]\nffffffffc0b25ed0 t fuse_flush_pq\t[fuse]\n",
			symbol:  "fuse_flush_pq",
			wantErr: false,
		},
		{
			name:    "symbol present without module",
			content: "ffffffff81000000 T startup_64\nffffffff82000000 t fuse_flush_pq\n",
			symbol:  "fuse_flush_pq",
			wantErr: false,
		},
		{
			name:    "symbol not present",
			content: "ffffffffc0b20ed0 t fuse_dev_ioctl_recover\t[fuse]\nffffffff81000000 T startup_64\n",
			symbol:  "fuse_flush_pq",
			wantErr: true,
		},
		{
			name:    "empty file",
			content: "",
			symbol:  "fuse_flush_pq",
			wantErr: true,
		},
		{
			name:    "partial match is not a match",
			content: "ffffffff82000000 t fuse_flush_pq_extended\n",
			symbol:  "fuse_flush_pq",
			wantErr: true,
		},
		{
			name:    "file does not exist",
			content: "",
			symbol:  "fuse_flush_pq",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var path string
			if tt.name == "file does not exist" {
				path = filepath.Join(t.TempDir(), "nonexistent")
			} else {
				path = filepath.Join(t.TempDir(), "kallsyms")
				if err := os.WriteFile(path, []byte(tt.content), 0o644); err != nil {
					t.Fatal(err)
				}
			}
			err := checkKallsymsForSymbol(path, tt.symbol)
			if (err != nil) != tt.wantErr {
				t.Errorf("checkKallsymsForSymbol() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
