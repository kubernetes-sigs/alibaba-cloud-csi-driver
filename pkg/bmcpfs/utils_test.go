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
	"github.com/stretchr/testify/require"
)

func TestParseVolumeID_Success(t *testing.T) {
	fsID, fileSetID, err := parseVolumeID("fs-123+fset-456")

	require.NoError(t, err)
	assert.Equal(t, "fs-123", fsID)
	assert.Equal(t, "fset-456", fileSetID)
}

func TestParseVolumeID_InvalidFormat_NoSeparator(t *testing.T) {
	fsID, fileSetID, err := parseVolumeID("fs-123fset-456")

	require.Error(t, err)
	assert.Empty(t, fsID)
	assert.Empty(t, fileSetID)
	assert.Contains(t, err.Error(), "invalid volume ID format")
}

func TestParseVolumeID_InvalidFormat_MultipleSeparators(t *testing.T) {
	fsID, fileSetID, err := parseVolumeID("fs-123+fset-456+extra")

	require.Error(t, err)
	assert.Empty(t, fsID)
	assert.Empty(t, fileSetID)
	assert.Contains(t, err.Error(), "invalid volume ID format")
}

func TestParseVolumeID_EmptyFilesystemID(t *testing.T) {
	fsID, fileSetID, err := parseVolumeID("+fset-456")

	require.Error(t, err)
	assert.Empty(t, fsID)
	assert.Empty(t, fileSetID)
	assert.Contains(t, err.Error(), "filesystem ID or fileset ID is empty")
}

func TestParseVolumeID_EmptyFilesetID(t *testing.T) {
	fsID, fileSetID, err := parseVolumeID("fs-123+")

	require.Error(t, err)
	assert.Empty(t, fsID)
	assert.Empty(t, fileSetID)
	assert.Contains(t, err.Error(), "filesystem ID or fileset ID is empty")
}

func TestParseVolumeID_EmptyVolumeID(t *testing.T) {
	fsID, fileSetID, err := parseVolumeID("")

	require.Error(t, err)
	assert.Empty(t, fsID)
	assert.Empty(t, fileSetID)
	assert.Contains(t, err.Error(), "invalid volume ID format")
}

func TestParseVolumeID_WithSpecialCharacters(t *testing.T) {
	fsID, fileSetID, err := parseVolumeID("fs-abc-123+fset-def-456")

	require.NoError(t, err)
	assert.Equal(t, "fs-abc-123", fsID)
	assert.Equal(t, "fset-def-456", fileSetID)
}
