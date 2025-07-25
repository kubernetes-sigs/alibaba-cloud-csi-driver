package metadata

import (
	"context"
	"net/http"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/metadata/imds"
)

type ECSDynamic struct {
	client imds.Client
}

func NewEcsDynamic(httpRT http.RoundTripper) *ECSDynamic {
	return &ECSDynamic{
		client: *imds.NewClient(httpRT),
	}
}

func (m *ECSDynamic) fetch(path string) (string, error) {
	data, err := m.client.Fetch(context.TODO(), path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (m *ECSDynamic) Get(key MetadataKey) (string, error) {
	switch key {
	case RAMRoleName:
		return m.fetch("meta-data/ram/security-credentials/")
	default:
		return "", ErrUnknownMetadataKey
	}
}
