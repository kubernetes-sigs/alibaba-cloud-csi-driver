package metadata

import (
	"context"
	"net/http"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/metadata/imds"
)

type IMDSDynamic struct {
	client imds.Client
}

func NewIMDSDynamic(httpRT http.RoundTripper) *IMDSDynamic {
	return &IMDSDynamic{
		client: *imds.NewClient(httpRT),
	}
}

func (m *IMDSDynamic) fetch(path string) (string, error) {
	data, err := m.client.Fetch(context.TODO(), path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (m *IMDSDynamic) Get(key MetadataKey) (string, error) {
	switch key {
	case RAMRoleName:
		return m.fetch("meta-data/ram/security-credentials/")
	default:
		return "", ErrUnknownMetadataKey
	}
}
