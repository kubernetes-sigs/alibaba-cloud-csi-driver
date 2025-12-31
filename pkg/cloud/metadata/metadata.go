package metadata

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/client-go/kubernetes"
	"k8s.io/klog/v2"
)

type MetadataKey int

const (
	RegionID MetadataKey = iota
	ZoneID
	InstanceID
	InstanceType
	AccountID
	ClusterID
	VmocType
	DataPlaneZoneID
	RAMRoleName
	RRSATokenFile

	// non-string metadata, not public, can only access with corresponding methods
	machineKind
)

const LingjunConfigFile = "/host/etc/eflo_config/lingjun_config"

func (k MetadataKey) String() string {
	switch k {
	case RegionID:
		return "RegionID"
	case ZoneID:
		return "ZoneID"
	case InstanceID:
		return "InstanceID"
	case InstanceType:
		return "InstanceType"
	case AccountID:
		return "AccountID"
	case ClusterID:
		return "ClusterID"
	case VmocType:
		return "VmocType"
	case DataPlaneZoneID:
		return "DataPlaneZoneID"
	case RAMRoleName:
		return "RAMRoleName"
	case RRSATokenFile:
		return "RRSATokenFile"
	case machineKind:
		return "MachineKind"
	default:
		return fmt.Sprintf("MetadataKey(%d)", k)
	}
}

type MachineKind int

const (
	MachineKindUnknown MachineKind = iota
	MachineKindECS
	MachineKindLingjun
)

// returns MachineKindECS if instance type starts with "ecs."
// implemented as a middleware to determine the kind as early as possible
// and avoids extra requests to slower fetcher
type inferMachineKind struct {
	next middleware
}

func (p inferMachineKind) GetAny(key MetadataKey) (any, error) {
	v, err := p.next.GetAny(key)
	if key == machineKind && err == ErrUnknownMetadataKey {
		t, err := p.next.GetAny(InstanceType)
		if err != nil {
			return nil, err
		}
		if strings.HasPrefix(t.(string), "ecs.") {
			return MachineKindECS, nil
		}
	}
	return v, err
}

var ErrUnknownMetadataKey = errors.New("unknown metadata key")

const DISABLE_ECS_ENV = "ALIBABA_CLOUD_NO_ECS_METADATA"
const KUBE_NODE_NAME_ENV = "KUBE_NODE_NAME"

type strMetadataProvider interface {
	Get(key MetadataKey) (string, error)
}

type MetadataProvider interface {
	strMetadataProvider
	MachineKind() (MachineKind, error)
}

type Metadata struct {
	providers multi
}

func getT[T any](m *Metadata, key MetadataKey) (T, error) {
	v, err := m.providers.GetAny(key)
	if err != nil {
		var zero T
		return zero, err
	}
	return v.(T), nil
}

func (m *Metadata) Get(key MetadataKey) (string, error) { return getT[string](m, key) }
func (m *Metadata) MachineKind() (MachineKind, error)   { return getT[MachineKind](m, machineKind) }

type fetcher interface {
	FetchFor(key MetadataKey) (middleware, error)
}

type middleware interface {
	GetAny(key MetadataKey) (any, error)
}

type lazyInit struct {
	next    middleware
	err     error
	initMu  sync.Mutex
	fetcher fetcher
}

func (p *lazyInit) GetAny(key MetadataKey) (any, error) {
	p.initMu.Lock()
	if p.next == nil && p.err == nil {
		next, err := p.fetcher.FetchFor(key)
		if err == ErrUnknownMetadataKey {
			p.initMu.Unlock()
			return nil, err
		}
		if err != nil {
			err = fmt.Errorf("%T failed: %w", p.fetcher, err)
			// print a warning if we failed to get a value,
			// because the error is hide if other providers succeed
			klog.Warning(err)
		}
		p.fetcher = nil
		p.next = next
		p.err = err
	}
	p.initMu.Unlock()
	if p.err != nil {
		return nil, p.err
	}
	return p.next.GetAny(key)
}

type immutable struct {
	next   middleware
	name   string
	mu     sync.Mutex
	values map[MetadataKey]any
}

func (p *immutable) GetAny(key MetadataKey) (any, error) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if v, ok := p.values[key]; ok {
		return v, nil
	}
	v, err := p.next.GetAny(key)
	if err != nil {
		return "", err
	}
	klog.V(2).InfoS("retrieved metadata", "provider", p.name, "key", key, "value", v)
	p.values[key] = v
	return v, nil
}

func newImmutable(next middleware, name string) *immutable {
	return &immutable{
		next:   next,
		name:   name,
		values: map[MetadataKey]any{},
	}
}

type strProvider struct {
	p strMetadataProvider
}

func (p strProvider) GetAny(key MetadataKey) (any, error) {
	return p.p.Get(key)
}

func NewMetadata() *Metadata {
	providers := multi{
		newImmutable(strProvider{ENVMetadata{}}, "env"),
	}
	lm, _ := NewLingJunMetadata(LingjunConfigFile)
	if lm != nil { // error is already logged
		providers = append(providers, newImmutable(lm, "lingjun"))
	}
	return &Metadata{providers}
}

func (m *Metadata) EnableEcs(httpRT http.RoundTripper) {
	if os.Getenv(DISABLE_ECS_ENV) != "" {
		klog.Infof("ECS metadata is disabled by environment variable %s", DISABLE_ECS_ENV)
		return
	}
	m.providers = append(m.providers, inferMachineKind{&lazyInit{
		fetcher: &EcsFetcher{httpRT: httpRT},
	}}, strProvider{NewEcsDynamic(httpRT)})
}

func (m *Metadata) EnableKubernetes(client kubernetes.Interface) {
	nodeName := os.Getenv(KUBE_NODE_NAME_ENV)
	if nodeName == "" {
		klog.Warningf("%s environment variable is not set, skipping Kubernetes Node metadata", KUBE_NODE_NAME_ENV)
	} else {
		m.enableKubernetesNode(client, nodeName)
	}
	m.providers = append(m.providers, &lazyInit{
		fetcher: &ProfileFetcher{
			client: client,
		},
	})
}

func (m *Metadata) enableKubernetesNode(client kubernetes.Interface, nodeName string) {
	m.providers = append(m.providers, inferMachineKind{&lazyInit{
		fetcher: &KubernetesNodeMetadataFetcher{
			client:   client.CoreV1().Nodes(),
			nodeName: nodeName,
		},
	}})
}

func (m *Metadata) EnableOpenAPI(ecsClient cloud.ECSv2Interface) {
	mPre := Metadata{
		// use the previous providers to get instance id,
		// do not recurse into ourselves
		providers: m.providers,
	}
	m.providers = append(m.providers, inferMachineKind{&lazyInit{
		fetcher: &OpenAPIFetcher{
			ecsClient: ecsClient,
			mPre:      &mPre,
		},
	}})
}

func (m *Metadata) EnableSts(stsClient cloud.STSInterface) {
	m.providers = append(m.providers, &lazyInit{
		fetcher: &StsFetcher{
			stsClient: stsClient,
		},
	})
}

type multi []middleware

func (m multi) GetAny(key MetadataKey) (any, error) {
	errors := []error{}
	for _, p := range m {
		v, err := p.GetAny(key)
		if err == nil {
			return v, nil
		}
		if err == ErrUnknownMetadataKey {
			continue
		}
		errors = append(errors, err)
	}
	if len(errors) == 0 {
		return nil, ErrUnknownMetadataKey
	}
	return nil, utilerrors.NewAggregate(errors)
}

func MustGet(m strMetadataProvider, key MetadataKey) string {
	value, err := m.Get(key)
	if err != nil {
		err = fmt.Errorf("failed to get metadata %s: %w", key, err)
		panic(err)
	}
	return value
}

// GetFallbackZoneID returns the zone ID for provision resource when it is specified by user
func GetFallbackZoneID(m MetadataProvider) (string, error) {
	zoneID, err := m.Get(DataPlaneZoneID)
	if err == nil {
		return zoneID, nil
	}
	if errors.Is(err, ErrUnknownMetadataKey) ||
		apierrors.IsNotFound(err) || apierrors.IsForbidden(err) {
		// Maybe Kubernetes provider is not enabled, or not in ACK cluster.
		// Use the zoneID of controller as the last resort.
		return m.Get(ZoneID)
	} else {
		return zoneID, err
	}
}

// FakeProvider is a fake metadata provider for testing
type FakeProvider struct {
	Values map[MetadataKey]string
	V      struct {
		MachineKind MachineKind
	}
}

func (p FakeProvider) Get(key MetadataKey) (string, error) {
	if v, ok := p.Values[key]; ok {
		return v, nil
	}
	return "", ErrUnknownMetadataKey
}

func (p FakeProvider) MachineKind() (MachineKind, error) {
	return p.V.MachineKind, nil
}
