package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRepositoryPrefix(t *testing.T) {
	tests := []struct {
		name   string
		region string
		prefix string
		url    string
		want   string
	}{
		{
			name: "all empty",
			want: DefImageRegistry + "/" + DefImageNamespace,
		},
		{
			name:   "only region",
			region: "cn-beijing",
			want:   "registry-cn-beijing-vpc.ack.aliyuncs.com/" + DefImageNamespace,
		},
		{
			name:   "url",
			region: "cn-beijing",
			url:    "registry-cn-wulanchabu.ack.aliyuncs.com",
			want:   "registry-cn-wulanchabu.ack.aliyuncs.com/" + DefImageNamespace,
		},
		{
			name:   "prefix",
			region: "cn-beijing",
			url:    "registry-cn-wulanchabu.ack.aliyuncs.com",
			prefix: "customized-prefix",
			want:   "customized-prefix",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("DEFAULT_REGISTRY", tt.url)
			t.Setenv("IMAGE_REPOSITORY_PREFIX", tt.prefix)
			actual := GetRepositoryPrefix(tt.region)
			assert.Equal(t, tt.want, actual)
		})
	}
}
