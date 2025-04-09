package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetRoleSessionName(t *testing.T) {
	tests := []struct {
		volumeId string
		target   string
		wantName string
	}{
		{"vol1", "/mnt/target1", "ossfs.vol1." + computeMountPathHash("/mnt/target1")},
		{"hereisalonglongpvnamethatisalreadylongerthan64ibeleive", "/mnt/target2", "ossfs.hereisalonglongpvnamethatisalreadylongerthan64ibeleive.c85"},
	}

	for _, test := range tests {
		sessionName := GetRoleSessionName(test.volumeId, test.target, "ossfs")
		assert.Equal(t, test.wantName, sessionName)
		assert.True(t, len(sessionName) <= MaxRoleSessionNameLimit, "sessionName length should not exceed %d, got: %d", MaxRoleSessionNameLimit, len(sessionName))
	}
}

func Test_GetMountProxySocketPath(t *testing.T) {
	tests := []struct {
		volumeId string
		expected string
	}{
		{"volume1", "/run/fuse.ossfs/022a36dfadf09ba4bf2549819660fea3ded8a9fc2ac564db0ca90af906b2a29a/mounter.sock"},
		{"volume2", "/run/fuse.ossfs/c1d2b5f1fd180bacb639735c65d0359b40115b1be2d4b98f0eb8dd0269f42534/mounter.sock"},
	}

	for _, test := range tests {
		actual := GetMountProxySocketPath(test.volumeId)
		assert.Equal(t, test.expected, actual)
	}
}

func Test_GetAttachPath(t *testing.T) {
	tests := []struct {
		volumeId string
		expected string
	}{
		{"volume1", "/run/fuse.ossfs/022a36dfadf09ba4bf2549819660fea3ded8a9fc2ac564db0ca90af906b2a29a/globalmount"},
		{"volume2", "/run/fuse.ossfs/c1d2b5f1fd180bacb639735c65d0359b40115b1be2d4b98f0eb8dd0269f42534/globalmount"},
	}

	for _, test := range tests {
		actual := GetAttachPath(test.volumeId)
		assert.Equal(t, test.expected, actual)
	}
}
