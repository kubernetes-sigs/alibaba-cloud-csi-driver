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

func TestGetSkipGlobalMount(t *testing.T) {
	tests := []struct {
		name               string
		skipGlobalMount    string
		ossSkipGlobalMount string
		defaultVal         bool
		want               bool
	}{
		{
			name:       "no env set, default false",
			defaultVal: false,
			want:       false,
		},
		{
			name:       "no env set, default true",
			defaultVal: true,
			want:       true,
		},
		{
			name:            "SKIP_GLOBAL_MOUNT takes priority over OSS_SKIP_GLOBAL_MOUNT",
			skipGlobalMount: "true",
			ossSkipGlobalMount: "false",
			defaultVal:         false,
			want:               true,
		},
		{
			name:               "OSS_SKIP_GLOBAL_MOUNT used when SKIP_GLOBAL_MOUNT is empty",
			ossSkipGlobalMount: "true",
			defaultVal:         false,
			want:               true,
		},
		{
			name:            "SKIP_GLOBAL_MOUNT=false overrides default true",
			skipGlobalMount: "false",
			defaultVal:      true,
			want:            false,
		},
		{
			name:               "OSS_SKIP_GLOBAL_MOUNT=0 parsed as false",
			ossSkipGlobalMount: "0",
			defaultVal:         true,
			want:               false,
		},
		{
			name:            "invalid SKIP_GLOBAL_MOUNT falls back to default",
			skipGlobalMount: "invalid",
			defaultVal:      true,
			want:            true,
		},
		{
			name:               "invalid OSS_SKIP_GLOBAL_MOUNT falls back to default",
			ossSkipGlobalMount: "yes",
			defaultVal:         false,
			want:               false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.skipGlobalMount != "" {
				t.Setenv("SKIP_GLOBAL_MOUNT", tt.skipGlobalMount)
			}
			if tt.ossSkipGlobalMount != "" {
				t.Setenv("OSS_SKIP_GLOBAL_MOUNT", tt.ossSkipGlobalMount)
			}
			got := GetSkipGlobalMount(tt.defaultVal)
			assert.Equal(t, tt.want, got)
		})
	}
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
