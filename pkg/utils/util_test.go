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
	"testing"

	"github.com/stretchr/testify/assert"
	"k8s.io/klog/v2"
	k8smount "k8s.io/mount-utils"
)

func TestGetServiceType(t *testing.T) {
	defer func() { klog.OsExit = os.Exit }()
	klog.OsExit = func(c int) { panic(c) }
	tests := []struct {
		name                 string
		runAsController      bool
		runControllerService bool
		runNodeService       bool
		serviceTypeEnv       string
		want                 ServiceType
		fatal                bool
	}{
		{
			name:                 "default",
			runControllerService: true,
			runNodeService:       true,
			want:                 Controller | Node,
		},
		{
			name:            "Run as controller",
			runAsController: true,
			want:            Controller,
		},
		{
			name:           "env provisioner",
			serviceTypeEnv: ProvisionerService,
			want:           Controller,
		},
		{
			name:           "env plugin",
			serviceTypeEnv: PluginService,
			want:           Node,
		},
		{
			name:                 "Run controller",
			runControllerService: true,
			runNodeService:       false,
			want:                 Controller,
		},
		{
			name:                 "Run node",
			runControllerService: false,
			runNodeService:       true,
			want:                 Node,
		},
		{
			name:                 "nothing",
			runControllerService: false,
			runNodeService:       false,
			want:                 0,
		},
		{
			name:           "invalid env",
			serviceTypeEnv: "invalid",
			fatal:          true,
		},
		{
			name:            "conflict env and run-as-controller",
			runAsController: true,
			serviceTypeEnv:  PluginService,
			fatal:           true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("SERVICE_TYPE", tt.serviceTypeEnv)
			if tt.fatal {
				assert.Panics(t, func() {
					GetServiceType(tt.runAsController, tt.runControllerService, tt.runNodeService)
				})
				return
			}
			got := GetServiceType(tt.runAsController, tt.runControllerService, tt.runNodeService)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestIsDirTmpfs(t *testing.T) {
	mounter := k8smount.NewFakeMounter([]k8smount.MountPoint{
		{
			Path: "/",
		},
		{
			Path: "/some/tmpfs",
			Type: "tmpfs",
		},
		{
			Path: "/some/other",
			Type: "ext4",
		},
	})
	isTmpfs, err := IsDirTmpfs(mounter, "/some/tmpfs")
	assert.Nil(t, err)
	assert.True(t, isTmpfs)
}

func TestIsDirTmpfsFalse(t *testing.T) {
	mounter := k8smount.NewFakeMounter([]k8smount.MountPoint{
		{
			Path: "/",
		},
		{
			Path: "/some/other",
			Type: "ext4",
		},
	})
	isTmpfs, err := IsDirTmpfs(mounter, "/some/tmpfs")
	assert.Nil(t, err)
	assert.False(t, isTmpfs)
}
