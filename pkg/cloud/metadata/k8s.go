package metadata

import (
	"context"
	"strings"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

type KubernetesNodeMetadata struct {
	node v1.Node
}

var (
	RegionIDLabels = []string{
		"topology.kubernetes.io/region",
		"failure-domain.beta.kubernetes.io/region",
		"alibabacloud.com/ecs-region-id",
		"sigma.ali/ecs-region-id",
	}
	ZoneIDLabels = []string{
		"topology.kubernetes.io/zone",
		"failure-domain.beta.kubernetes.io/zone",
		"alibabacloud.com/ecs-zone-id",
		"sigma.ali/ecs-zone-id",
	}
	InstanceTypeLabels = []string{
		"node.kubernetes.io/instance-type",
		"beta.kubernetes.io/instance-type",
		"sigma.ali/machine-model",
	}
	InstanceIdLabels = []string{
		"alibabacloud.com/ecs-instance-id",
	}
)

var MetadataLabels = map[MetadataKey][]string{
	RegionID:     RegionIDLabels,
	ZoneID:       ZoneIDLabels,
	InstanceType: InstanceTypeLabels,
	InstanceID:   InstanceIdLabels,
}

func NewKubernetesNodeMetadata(nodeName string, nodeClient corev1.NodeInterface) (*KubernetesNodeMetadata, error) {
	node, err := nodeClient.Get(context.Background(), nodeName, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return &KubernetesNodeMetadata{node: *node}, nil
}

func (m *KubernetesNodeMetadata) Get(key MetadataKey) (string, error) {
	labels := MetadataLabels[key]
	for _, label := range labels {
		if value, ok := m.node.Labels[label]; ok {
			return value, nil
		}
	}

	providerIdSegments := strings.Split(m.node.Spec.ProviderID, ".")
	if len(providerIdSegments) == 2 {
		switch key {
		case RegionID:
			return strings.TrimPrefix(providerIdSegments[0], "alicloud://"), nil
		case InstanceID:
			return providerIdSegments[1], nil
		}
	}

	return "", ErrUnknownMetadataKey
}

type KubernetesMetadataFetcher struct {
	client   corev1.NodeInterface
	nodeName string
}

func (f *KubernetesMetadataFetcher) FetchFor(key MetadataKey) (MetadataProvider, error) {
	if _, ok := MetadataLabels[key]; !ok {
		return nil, ErrUnknownMetadataKey
	}
	return NewKubernetesNodeMetadata(f.nodeName, f.client)
}
