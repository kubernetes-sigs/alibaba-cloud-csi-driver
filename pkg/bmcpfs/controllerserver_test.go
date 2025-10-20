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

package bmcpfs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestParseVolumeHandle tests the parseVolumeHandle function with various inputs
func TestParseVolumeHandle(t *testing.T) {
	tests := []struct {
		name           string
		volumeHandle   string
		expectedFsId   string
		expectedFsetId string
	}{
		{
			name:           "single part",
			volumeHandle:   "fs-12345",
			expectedFsId:   "fs-12345",
			expectedFsetId: "",
		},
		{
			name:           "two parts",
			volumeHandle:   "fs-12345+fset-67890",
			expectedFsId:   "fs-12345",
			expectedFsetId: "fset-67890",
		},
		{
			name:           "multiple delimiters",
			volumeHandle:   "fs-12345+fset-67890+extra",
			expectedFsId:   "fs-12345",
			expectedFsetId: "fset-67890+extra",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fsId, fsetId := parseVolumeHandle(tt.volumeHandle)
			assert.Equal(t, tt.expectedFsId, fsId)
			assert.Equal(t, tt.expectedFsetId, fsetId)
		})
	}
}
