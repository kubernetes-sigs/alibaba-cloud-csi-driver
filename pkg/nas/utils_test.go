//go:build !windows

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

package nas

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseMountFlags(t *testing.T) {
	tests := []struct {
		name        string
		mntOptions  []string
		wantVers    string
		wantOptions []string
	}{
		{
			"vers=3.0 normalized to 3",
			[]string{"mnt=/test", "vers=3.0"},
			"3", []string{"mnt=/test"},
		},
		{
			"vers=4.1 not normalized",
			[]string{"mnt=/test", "vers=4.1"},
			"4.1", []string{"mnt=/test"},
		},
		{
			"no vers",
			[]string{"mnt=/test", "a=b", "c=d"},
			"", []string{"mnt=/test", "a=b", "c=d"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotVers, gotOptions := ParseMountFlags(tt.mntOptions)
			assert.Equal(t, tt.wantVers, gotVers)
			assert.Equal(t, tt.wantOptions, gotOptions)
		})
	}
}

func Test_addTLSMountOptions(t *testing.T) {
	type args struct {
		baseOptions []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"already set tls",
			args{[]string{"vers=3", "tls"}},
			[]string{"vers=3", "tls"},
		},
		{
			"tls not set",
			args{[]string{"vers=3", "nolock"}},
			[]string{"vers=3", "nolock", "tls"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, addTLSMountOptions(tt.args.baseOptions))
		})
	}
}

func TestGetMountRootAndRelPath(t *testing.T) {
	tests := []struct {
		name             string
		mountFsType      string
		opt              *Options
		expectedRootPath string
		expectedRelpath  string
	}{
		{
			name: "nil options",
		},
		{
			name:             "general nas",
			mountFsType:      "nfs",
			opt:              &Options{FSType: "general", Path: "/test", Server: "test-server"},
			expectedRootPath: "test-server:/",
			expectedRelpath:  "test",
		},
		{
			name:             "extreme nas with share-prefixed path",
			mountFsType:      "nfs",
			opt:              &Options{FSType: "extreme", Path: "/share/test", Server: "test-server"},
			expectedRootPath: "test-server:/",
			expectedRelpath:  "test",
		},
		{
			name:             "extreme nas with plain path",
			mountFsType:      "nfs",
			opt:              &Options{FSType: "extreme", Path: "/test", Server: "test-server"},
			expectedRootPath: "test-server:/",
			expectedRelpath:  "test",
		},
		{
			name:             "cpfs with share-prefixed path",
			mountFsType:      "cpfs-nfs",
			opt:              &Options{FSType: "cpfs", Path: "/share/test", Server: "test-server"},
			expectedRootPath: "test-server:/share",
			expectedRelpath:  "test",
		},
		{
			name:             "cpfs with root path",
			mountFsType:      "cpfs-nfs",
			opt:              &Options{FSType: "cpfs", Path: "/", Server: "test-server"},
			expectedRootPath: "",
			expectedRelpath:  "",
		},
		{
			name:             "cpfs with plain path",
			mountFsType:      "cpfs-nfs",
			opt:              &Options{FSType: "cpfs", Path: "/test", Server: "test-server"},
			expectedRootPath: "",
			expectedRelpath:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualRootPath, actualRelpath := getMountRootAndRelPath(tt.mountFsType, tt.opt)
			assert.Equal(t, tt.expectedRootPath, actualRootPath)
			assert.Equal(t, tt.expectedRelpath, actualRelpath)
		})
	}
}

func TestIsSubDir(t *testing.T) {
	tests := []struct {
		name            string
		parent          string
		child           string
		expected        bool
		expectedRelPath string
	}{
		{
			name:            "child is subdir of parent",
			parent:          "/parent",
			child:           "/parent/child",
			expected:        true,
			expectedRelPath: "child",
		},
		{
			name:            "child is subdir of parent with trailing slash",
			parent:          "/parent/",
			child:           "/parent/child",
			expected:        true,
			expectedRelPath: "child",
		},
		{
			name:            "child is not subdir of parent",
			parent:          "/parent",
			child:           "/child",
			expected:        false,
			expectedRelPath: "",
		},
		{
			name:            "parent equals child",
			parent:          "/parent",
			child:           "/parent",
			expected:        false,
			expectedRelPath: "",
		},
		{
			name:            "parent and child has the same prefix",
			parent:          "/parent",
			child:           "/parent-child",
			expected:        false,
			expectedRelPath: "",
		},
		{
			name:            "parent equals child with trailing slash",
			parent:          "/parent/",
			child:           "/parent",
			expected:        false,
			expectedRelPath: "",
		},
		{
			name:            "child with relative path",
			parent:          "/parent",
			child:           "parent",
			expected:        false,
			expectedRelPath: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualRelpath, actual := extractSubDir(tt.parent, tt.child)
			assert.Equal(t, tt.expected, actual)
			assert.Equal(t, tt.expectedRelPath, actualRelpath)
		})
	}
}
