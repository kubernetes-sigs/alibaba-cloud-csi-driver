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
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDetectAuthMode(t *testing.T) {
	tests := []struct {
		name    string
		secrets map[string]string
		want    authMode
		wantErr string
	}{
		{name: "nil secrets", secrets: nil, want: authModeNone},
		{name: "empty secrets", secrets: map[string]string{}, want: authModeNone},
		{
			name:    "ak pair",
			secrets: map[string]string{"accessKeyId": "ak", "accessKeySecret": "sk"},
			want:    authModeAK,
		},
		{
			name:    "sts triple",
			secrets: map[string]string{"accessKeyId": "ak", "accessKeySecret": "sk", "securityToken": "tok"},
			want:    authModeSTS,
		},
		{
			name:    "sts triple with expiration",
			secrets: map[string]string{"accessKeyId": "ak", "accessKeySecret": "sk", "securityToken": "tok", "expiration": "2026-01-01T00:00:00Z"},
			want:    authModeSTS,
		},
		{
			name:    "missing access key secret",
			secrets: map[string]string{"accessKeyId": "ak"},
			wantErr: "required",
		},
		{
			name:    "empty access key id value",
			secrets: map[string]string{"accessKeyId": "", "accessKeySecret": "sk"},
			wantErr: "required",
		},
		{
			name:    "unknown key ignored",
			secrets: map[string]string{"accessKeyId": "ak", "accessKeySecret": "sk", "security_token": "tok"},
			want:    authModeAK,
		},
		{
			name:    "expiration without token still AK",
			secrets: map[string]string{"accessKeyId": "ak", "accessKeySecret": "sk", "expiration": "x"},
			want:    authModeAK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := detectAuthMode(tt.secrets)
			if tt.wantErr != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.wantErr)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestCredentialsDirForVolume(t *testing.T) {
	tests := []struct {
		name     string
		volumeID string
		want     string
		wantErr  bool
	}{
		{name: "plain", volumeID: "cpfs-123", want: "/root/cpfs-123"},
		{name: "with suffix", volumeID: "cpfs-123+ap-abc", want: "/root/cpfs-123+ap-abc"},
		{name: "empty rejected", volumeID: "", wantErr: true},
		{name: "slash rejected", volumeID: "a/b", wantErr: true},
		{name: "traversal rejected", volumeID: "..", wantErr: true},
		{name: "embedded traversal rejected", volumeID: "a..b", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := credentialsDirForVolume("/root", tt.volumeID)
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestWriteSTSFile(t *testing.T) {
	dir := t.TempDir()
	secrets := map[string]string{"accessKeyId": "ak", "accessKeySecret": "sk", "securityToken": "tok"}

	path, changed, err := writeSTSFile(dir, secrets)
	require.NoError(t, err)
	assert.True(t, changed)
	assert.Equal(t, filepath.Join(dir, stsFileName), path)

	raw, err := os.ReadFile(path)
	require.NoError(t, err)
	var cred stsCredentials
	require.NoError(t, json.Unmarshal(raw, &cred))
	assert.Equal(t, "ak", cred.AccessKeyID)
	assert.Equal(t, "sk", cred.AccessKeySecret)
	assert.Equal(t, "tok", cred.StsToken)
	// md5("aksktok"), i.e. md5 of the concatenated credential triple.
	assert.Equal(t, "cb42cf60c84c55aff7bbb0c35443c165", cred.MD5)

	info, err := os.Stat(path)
	require.NoError(t, err)
	assert.Equal(t, os.FileMode(0o600), info.Mode().Perm())

	// Identical content is a no-op.
	_, changed, err = writeSTSFile(dir, secrets)
	require.NoError(t, err)
	assert.False(t, changed)

	// Rotated token rewrites with a new checksum.
	secrets["securityToken"] = "tok2"
	_, changed, err = writeSTSFile(dir, secrets)
	require.NoError(t, err)
	assert.True(t, changed)
	raw, err = os.ReadFile(path)
	require.NoError(t, err)
	require.NoError(t, json.Unmarshal(raw, &cred))
	assert.Equal(t, "tok2", cred.StsToken)
}

func TestWriteAKFile(t *testing.T) {
	// A non-existent dir so MkdirAll actually creates it and its 0700 mode can
	// be asserted (a pre-existing t.TempDir would not be re-chmodded).
	dir := filepath.Join(t.TempDir(), "vol")
	secrets := map[string]string{"accessKeyId": "ak", "accessKeySecret": "sk"}

	path, err := writeAKFile(dir, secrets)
	require.NoError(t, err)
	assert.Equal(t, filepath.Join(dir, akFileName), path)

	raw, err := os.ReadFile(path)
	require.NoError(t, err)
	var cred akCredentials
	require.NoError(t, json.Unmarshal(raw, &cred))
	assert.Equal(t, "ak", cred.AccessKeyID)
	assert.Equal(t, "sk", cred.AccessKeySecret)

	info, err := os.Stat(path)
	require.NoError(t, err)
	assert.Equal(t, os.FileMode(0o600), info.Mode().Perm())

	dirInfo, err := os.Stat(dir)
	require.NoError(t, err)
	assert.Equal(t, os.FileMode(0o700), dirInfo.Mode().Perm())
}

func TestWriteCredentialsFile_CreatesDir(t *testing.T) {
	dir := filepath.Join(t.TempDir(), "nested", "vol")
	changed, err := writeCredentialsFile(dir, "f.json", []byte("{}"))
	require.NoError(t, err)
	assert.True(t, changed)
	assert.FileExists(t, filepath.Join(dir, "f.json"))
}
