package metadata

import (
	"context"
	"os"
	"strings"

	v1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type KubernetesMetadata struct {
	node    *v1.Node
	profile *v1.ConfigMap
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
		"sigma.ali/ecs-instance-id",
	}
)

var MetadataLabels = map[MetadataKey][]string{
	RegionID:     RegionIDLabels,
	ZoneID:       ZoneIDLabels,
	InstanceType: InstanceTypeLabels,
	InstanceID:   InstanceIdLabels,
}

var MetadataProfileDataKeys = map[MetadataKey]string{
	ClusterID: "clusterid",
}

func init() {
	envInstanceIdKey := os.Getenv("NODE_LABEL_ECS_ID_KEY")
	if envInstanceIdKey != "" {
		InstanceIdLabels = append([]string{envInstanceIdKey}, InstanceIdLabels...)
		MetadataLabels[InstanceID] = InstanceIdLabels
	}
}

func NewKubernetesMetadata(nodeName string, client kubernetes.Interface) (*KubernetesMetadata, error) {
	node, err := client.CoreV1().Nodes().Get(context.Background(), nodeName, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	profile, err := client.CoreV1().ConfigMaps("kube-system").Get(context.Background(), "ack-cluster-profile", metav1.GetOptions{})
	if err != nil && !apierrors.IsNotFound(err) {
		return nil, err
	}
	return &KubernetesMetadata{node: node, profile: profile}, nil
}

func (m *KubernetesMetadata) Get(key MetadataKey) (string, error) {

	if key, ok := MetadataProfileDataKeys[key]; ok && m.profile != nil {
		return m.profile.Data[key], nil
	}

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
	client   kubernetes.Interface
	nodeName string
}

func (f *KubernetesMetadataFetcher) FetchFor(key MetadataKey) (MetadataProvider, error) {
	_, ok1 := MetadataLabels[key]
	_, ok2 := MetadataProfileDataKeys[key]
	if !ok1 && !ok2 {
		return nil, ErrUnknownMetadataKey
	}
	p, err := NewKubernetesMetadata(f.nodeName, f.client)
	if err != nil {
		return nil, err
	}
	return newImmutableProvider(p, "Kubernetes"), nil
}
