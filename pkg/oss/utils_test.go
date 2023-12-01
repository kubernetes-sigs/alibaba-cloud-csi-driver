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

package oss

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRAMRoleOption(t *testing.T) {
	result := GetRAMRoleOption()
	assert.NotEqual(t, "", result)
}

func Test_parseOtherOpts(t *testing.T) {
	type args struct {
		otherOpts string
	}
	tests := []struct {
		name             string
		args             args
		wantMountOptions []string
		wantErr          bool
	}{
		{
			"normal",
			args{"-o max_stat_cache_size=0 -o allow_other"},
			[]string{"max_stat_cache_size=0", "allow_other"},
			false,
		},
		{
			"empty",
			args{""},
			nil,
			false,
		},
		{
			"spaces",
			args{"	 "},
			nil,
			false,
		},
		{
			"merged -o",
			args{"-omax_stat_cache_size=0 -o allow_other"},
			[]string{"max_stat_cache_size=0", "allow_other"},
			false,
		},
		{
			"missing -o",
			args{"-omax_stat_cache_size=0 allow_other"},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMountOptions, err := parseOtherOpts(tt.args.otherOpts)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseOtherOpts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotMountOptions, tt.wantMountOptions) {
				t.Errorf("parseOtherOpts() = %v, want %v", gotMountOptions, tt.wantMountOptions)
			}
		})
	}
}
