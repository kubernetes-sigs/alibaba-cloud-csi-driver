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
	"bytes"
	// md5 is mandated by the EFC STSFile format as a content checksum
	// (md5 of the concatenated credential triple), not used for security.
	"crypto/md5"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Keys accepted in the nodePublishSecretRef secret. The auth mode is derived
// from the key presence (no authType attribute): AK/SK selects AK mode, AK/SK
// plus securityToken selects STS mode.
const (
	secretKeyAccessKeyID     = "accessKeyId"
	secretKeyAccessKeySecret = "accessKeySecret"
	secretKeySecurityToken   = "securityToken"
)

// Credential files consumed by the EFC client. The directory is shared with
// the alinas mount-proxy container via the /run/cnfs hostPath, so the paths
// passed as mount options resolve identically on both sides. EFC re-reads the
// STS file every 10 minutes to refresh signatures; the AK file is read only
// at mount time.
const (
	defaultCredentialsRoot = "/run/cnfs/efc-credentials"
	akFileName             = "ak.json"
	stsFileName            = "sts.json"

	credentialsDirMode  = 0o700
	credentialsFileMode = 0o600
)

// authMode is the RAM auth mode of a volume, derived from the secret key set.
type authMode int

const (
	authModeNone authMode = iota
	authModeAK
	authModeSTS
)

func (m authMode) String() string {
	switch m {
	case authModeAK:
		return "ak"
	case authModeSTS:
		return "ststoken"
	default:
		return "none"
	}
}

// detectAuthMode classifies the nodePublishSecretRef contents by key
// presence: no secrets means no auth, a security token selects STS mode,
// otherwise AK mode. Error messages reference key names only, never values.
func detectAuthMode(secrets map[string]string) (authMode, error) {
	if len(secrets) == 0 {
		return authModeNone, nil
	}
	if secrets[secretKeyAccessKeyID] == "" || secrets[secretKeyAccessKeySecret] == "" {
		return authModeNone, fmt.Errorf("secret keys %q and %q are required", secretKeyAccessKeyID, secretKeyAccessKeySecret)
	}
	if secrets[secretKeySecurityToken] != "" {
		return authModeSTS, nil
	}
	return authModeAK, nil
}

// credentialsDirForVolume returns the per-volume credentials directory,
// rejecting volume IDs that could escape the root. Volume IDs are opaque
// user-controlled strings, so they must not contain path separators or
// traversal sequences.
func credentialsDirForVolume(root, volumeID string) (string, error) {
	if volumeID == "" || strings.Contains(volumeID, "/") || strings.Contains(volumeID, "..") {
		return "", fmt.Errorf("volume ID %q is not a valid credentials directory name", volumeID)
	}
	return filepath.Join(root, volumeID), nil
}

type akCredentials struct {
	AccessKeyID     string `json:"accessKeyId"`
	AccessKeySecret string `json:"accessKeySecret"`
}

type stsCredentials struct {
	AccessKeyID     string `json:"accessKeyId"`
	AccessKeySecret string `json:"accessKeySecret"`
	StsToken        string `json:"stsToken"`
	MD5             string `json:"md5"`
}

// writeAKFile persists the static AK credential for the EFC g_unas_AKFile
// mount option. AK does not support hot update, so callers only invoke this
// on the initial publish.
func writeAKFile(dir string, secrets map[string]string) (string, error) {
	data, err := json.Marshal(akCredentials{
		AccessKeyID:     secrets[secretKeyAccessKeyID],
		AccessKeySecret: secrets[secretKeyAccessKeySecret],
	})
	if err != nil {
		return "", fmt.Errorf("marshal AK credentials: %w", err)
	}
	if _, err := writeCredentialsFile(dir, akFileName, data); err != nil {
		return "", err
	}
	return filepath.Join(dir, akFileName), nil
}

// writeSTSFile persists the STS credential triple for the EFC g_unas_STSFile
// mount option, computing the md5 checksum the EFC client requires
// (md5 of accessKeyId + accessKeySecret + stsToken concatenated). It reports
// whether the file content actually changed, so republish-driven rotation can
// skip no-op rewrites.
func writeSTSFile(dir string, secrets map[string]string) (path string, changed bool, err error) {
	ak, sk, token := secrets[secretKeyAccessKeyID], secrets[secretKeyAccessKeySecret], secrets[secretKeySecurityToken]
	sum := md5.Sum([]byte(ak + sk + token))
	data, err := json.Marshal(stsCredentials{
		AccessKeyID:     ak,
		AccessKeySecret: sk,
		StsToken:        token,
		MD5:             fmt.Sprintf("%x", sum),
	})
	if err != nil {
		return "", false, fmt.Errorf("marshal STS credentials: %w", err)
	}
	changed, err = writeCredentialsFile(dir, stsFileName, data)
	if err != nil {
		return "", false, err
	}
	return filepath.Join(dir, stsFileName), changed, nil
}

// writeCredentialsFile atomically replaces dir/name with data via a temp file
// and rename, so the EFC client never observes a partially written file. The
// write is skipped when the content is unchanged. Returns whether the file
// was (re)written.
func writeCredentialsFile(dir, name string, data []byte) (bool, error) {
	if err := os.MkdirAll(dir, credentialsDirMode); err != nil {
		return false, fmt.Errorf("create credentials dir: %w", err)
	}
	path := filepath.Join(dir, name)
	if old, err := os.ReadFile(path); err == nil && bytes.Equal(old, data) {
		return false, nil
	}
	tmp, err := os.CreateTemp(dir, name+".tmp-*")
	if err != nil {
		return false, fmt.Errorf("create temp credentials file: %w", err)
	}
	tmpName := tmp.Name()
	// Best-effort cleanup on any failure path; a no-op once renamed.
	defer func() { _ = os.Remove(tmpName) }()
	if err := tmp.Chmod(credentialsFileMode); err != nil {
		_ = tmp.Close()
		return false, fmt.Errorf("chmod temp credentials file: %w", err)
	}
	if _, err := tmp.Write(data); err != nil {
		_ = tmp.Close()
		return false, fmt.Errorf("write temp credentials file: %w", err)
	}
	if err := tmp.Sync(); err != nil {
		_ = tmp.Close()
		return false, fmt.Errorf("sync temp credentials file: %w", err)
	}
	if err := tmp.Close(); err != nil {
		return false, fmt.Errorf("close temp credentials file: %w", err)
	}
	if err := os.Rename(tmpName, path); err != nil {
		return false, fmt.Errorf("rename temp credentials file: %w", err)
	}
	return true, nil
}
