/*
Copyright 2026 The Kubernetes Authors.

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

package main

import (
	"testing"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/substrate"
	"github.com/stretchr/testify/require"
)

func TestExpandDriverNames(t *testing.T) {
	for _, tc := range []struct {
		name string
		in   []string
		want []string
	}{
		{"short names", []string{"disk", "nas", "oss", substrate.DriverShortName}, []string{TypePluginDISK, TypePluginNAS, TypePluginOSS, substrate.DriverName}},
		{"full names are kept", []string{substrate.DriverName, TypePluginNAS}, []string{substrate.DriverName, TypePluginNAS}},
		{"agent is not a driver name", []string{ExtenderAgent}, []string{ExtenderAgent}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			expandDriverNames(tc.in)
			require.Equal(t, tc.want, tc.in)
		})
	}
}

func TestDriverMetricType(t *testing.T) {
	for _, tc := range []struct {
		name string
		in   string
		want string
	}{
		{"disk", TypePluginDISK, "disk"},
		{"nas", TypePluginNAS, "nas"},
		{"bmcpfs", TypePluginBMCPFS, "bmcpfs"},
		{"substrate", substrate.DriverName, substrate.DriverShortName},
	} {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.want, driverMetricType(tc.in))
		})
	}
}
