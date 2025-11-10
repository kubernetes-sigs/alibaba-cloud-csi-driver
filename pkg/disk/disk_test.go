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

package disk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseLingjunNodeDiskTypes(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:  "empty input",
			input: "",
			expected: []string{
				string(DiskESSDAuto),
				string(DiskESSD),
			},
		},
		{
			name:  "whitespace only input",
			input: "   ",
			expected: []string{
				string(DiskESSDAuto),
				string(DiskESSD),
			},
		},
		{
			name:  "valid single disk type",
			input: "cloud_ssd",
			expected: []string{
				"cloud_ssd",
			},
		},
		{
			name:  "valid multiple disk types",
			input: "cloud_ssd,cloud_essd",
			expected: []string{
				"cloud_ssd",
				"cloud_essd",
			},
		},
		{
			name:  "valid multiple disk types with spaces",
			input: "cloud_ssd, cloud_essd , cloud_efficiency",
			expected: []string{
				"cloud_ssd",
				"cloud_essd",
				"cloud_efficiency",
			},
		},
		{
			name:  "duplicate disk types",
			input: "cloud_ssd,cloud_essd,cloud_ssd",
			expected: []string{
				"cloud_ssd",
				"cloud_essd",
			},
		},
		{
			name:  "mix of valid and invalid disk types",
			input: "cloud_ssd,invalid_type,cloud_essd",
			expected: []string{
				"cloud_ssd",
				"cloud_essd",
			},
		},
		{
			name:  "all invalid disk types",
			input: "invalid_type1,invalid_type2",
			expected: []string{
				string(DiskESSDAuto),
				string(DiskESSD),
			},
		},
		{
			name:  "valid disk types with empty entries",
			input: "cloud_ssd,,cloud_essd,",
			expected: []string{
				"cloud_ssd",
				"cloud_essd",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parseLingjunNodeDiskTypes(tt.input)
			assert.ElementsMatch(t, tt.expected, result, "Mismatch between expected and actual disk types")
		})
	}
}
