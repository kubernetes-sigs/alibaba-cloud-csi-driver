package metadata

import (
	"encoding/json"
	"os"
	"strings"
	"sync"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/sets"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

// instanceTypeInfo is the annotation value of "alibabacloud.com/instance-type-info"
type instanceTypeInfo struct {
	DiskQuantity int32
}

type KubernetesNodeMetadata struct {
	node         *v1.Node
	instanceType func() (*instanceTypeInfo, error)
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
		"alibabacloud.com/instance-id", // lingjun node label
		"alibabacloud.com/ecs-instance-id",
		"sigma.ali/ecs-instance-id",
	}
	VmocLabels = []string{
		"rm.alibaba-inc.com/vbm",
	}
)

const instanceTypeInfoAnnotation = "alibabacloud.com/instance-type-info"

var MetadataLabels = map[MetadataKey][]string{
	RegionID:     RegionIDLabels,
	ZoneID:       ZoneIDLabels,
	InstanceType: InstanceTypeLabels,
	InstanceID:   InstanceIdLabels,
	VmocType:     VmocLabels,
}

func init() {
	envInstanceIdKey := os.Getenv("NODE_LABEL_ECS_ID_KEY")
	if envInstanceIdKey != "" {
		InstanceIdLabels = append([]string{envInstanceIdKey}, InstanceIdLabels...)
		MetadataLabels[InstanceID] = InstanceIdLabels
	}
}

// LingjunWorkerLabel is the label key used to identify Lingjun nodes.
const LingjunWorkerLabel = "alibabacloud.com/lingjun-worker"

// NodeMetadataLabelKeys returns the set of label keys needed for metadata extraction.
func NodeMetadataLabelKeys() sets.Set[string] {
	keys := sets.New[string]()
	for _, candidates := range MetadataLabels {
		keys.Insert(candidates...)
	}
	keys.Insert(LingjunWorkerLabel)
	return keys
}

// NodeMetadataAnnotationKeys returns the set of annotation keys needed for metadata extraction.
func NodeMetadataAnnotationKeys() sets.Set[string] {
	return sets.New(instanceTypeInfoAnnotation)
}

func NewKubernetesNodeMetadata(node *v1.Node) *KubernetesNodeMetadata {
	instanceType := func() (*instanceTypeInfo, error) {
		return nil, ErrUnknownMetadataKey
	}
	if value := node.Annotations[instanceTypeInfoAnnotation]; value != "" {
		instanceType = sync.OnceValues(func() (*instanceTypeInfo, error) {
			var info instanceTypeInfo
			err := json.Unmarshal([]byte(value), &info)
			return &info, err
		})
	}
	return &KubernetesNodeMetadata{node: node, instanceType: instanceType}
}

// Get returns the string metadata value for the given key.
func (m *KubernetesNodeMetadata) Get(key MetadataKey) (string, error) {
	labels := MetadataLabels[key]
	for _, label := range labels {
		if value := m.node.Labels[label]; value != "" {
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

// DiskQuantity returns the DiskQuantity decoded from the
// "alibabacloud.com/instance-type-info" annotation, or [ErrUnknownMetadataKey]
// if the annotation is absent.
func (m *KubernetesNodeMetadata) DiskQuantity() (int32, error) {
	if i, err := m.instanceType(); err == nil {
		return i.DiskQuantity, nil
	} else {
		return 0, err
	}
}

func (m *KubernetesNodeMetadata) MachineKind() (MachineKind, error) {
	if m.node.Labels[LingjunWorkerLabel] == "true" {
		return MachineKindLingjun, nil
	}
	return MachineKindUnknown, ErrUnknownMetadataKey
}

func (m *KubernetesNodeMetadata) GetAny(_ *mcontext, key MetadataKey) (any, error) {
	switch key {
	case diskQuantity:
		return m.DiskQuantity()
	case machineKind:
		return m.MachineKind()
	}
	return m.Get(key)
}

type KubernetesNodeMetadataFetcher struct {
	client   corev1.NodeInterface
	nodeName string
}

func (f *KubernetesNodeMetadataFetcher) ID() fetcherID { return kubernetesNodeMetadataFetcherID }

func (f *KubernetesNodeMetadataFetcher) FetchFor(ctx *mcontext, key MetadataKey) (middleware, error) {
	_, ok := MetadataLabels[key]
	if !ok {
		switch key {
		case diskQuantity, machineKind: // supported
		default:
			return nil, ErrUnknownMetadataKey
		}
	}
	node, err := f.client.Get(ctx, f.nodeName, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return newImmutable(NewKubernetesNodeMetadata(node), "Kubernetes"), nil
}
