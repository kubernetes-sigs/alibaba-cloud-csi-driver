package metadata

import (
	"errors"
	"fmt"
	"net/http"
	"os"
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
	default:
		return fmt.Sprintf("MetadataKey(%d)", k)
	}
}

var ErrUnknownMetadataKey = errors.New("unknown metadata key")

const DISABLE_ECS_ENV = "ALIBABA_CLOUD_NO_ECS_METADATA"
const KUBE_NODE_NAME_ENV = "KUBE_NODE_NAME"

type MetadataProvider interface {
	Get(key MetadataKey) (string, error)
}

type Metadata struct {
	providers []MetadataProvider
}

type MetadataFetcher interface {
	FetchFor(key MetadataKey) (MetadataProvider, error)
}

type lazyInitProvider struct {
	provider MetadataProvider
	err      error
	initMu   sync.Mutex
	fetcher  MetadataFetcher
}

func (p *lazyInitProvider) Get(key MetadataKey) (string, error) {
	p.initMu.Lock()
	if p.provider == nil && p.err == nil {
		provider, err := p.fetcher.FetchFor(key)
		if err == ErrUnknownMetadataKey {
			p.initMu.Unlock()
			return "", err
		}
		if err != nil {
			err = fmt.Errorf("%T failed: %w", p.fetcher, err)
			// print a warning if we failed to get a value,
			// because the error is hide if other providers succeed
			klog.Warning(err)
		}
		p.fetcher = nil
		p.provider = provider
		p.err = err
	}
	p.initMu.Unlock()
	if p.err != nil {
		return "", p.err
	}
	return p.provider.Get(key)
}

type immutableProvider struct {
	provider MetadataProvider
	name     string
	mu       sync.Mutex
	values   map[MetadataKey]string
}

func (p *immutableProvider) Get(key MetadataKey) (string, error) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if v, ok := p.values[key]; ok {
		return v, nil
	}
	v, err := p.provider.Get(key)
	if err != nil {
		return "", err
	}
	klog.V(2).InfoS("retrieved metadata", "provider", p.name, "key", key, "value", v)
	p.values[key] = v
	return v, nil
}

func newImmutableProvider(provider MetadataProvider, name string) *immutableProvider {
	return &immutableProvider{
		provider: provider,
		name:     name,
		values:   map[MetadataKey]string{},
	}
}

func NewMetadata() *Metadata {
	defaultMetadata := &Metadata{
		providers: []MetadataProvider{
			newImmutableProvider(&ENVMetadata{}, "env"),
		},
	}
	lm, err := NewLingJunMetadata(LingjunConfigFile)
	if err != nil {
		return defaultMetadata
	}
	if lm != nil {
		defaultMetadata.providers = append(defaultMetadata.providers, newImmutableProvider(lm, "lingjun"))
	}
	return defaultMetadata
}

func (m *Metadata) EnableEcs(httpRT http.RoundTripper) {
	if os.Getenv(DISABLE_ECS_ENV) != "" {
		klog.Infof("ECS metadata is disabled by environment variable %s", DISABLE_ECS_ENV)
		return
	}
	m.providers = append(m.providers, &lazyInitProvider{
		fetcher: &EcsFetcher{httpRT: httpRT},
	}, NewEcsDynamic(httpRT))
}

func (m *Metadata) EnableKubernetes(client kubernetes.Interface) {
	nodeName := os.Getenv(KUBE_NODE_NAME_ENV)
	if nodeName == "" {
		klog.Warningf("%s environment variable is not set, skipping Kubernetes Node metadata", KUBE_NODE_NAME_ENV)
	} else {
		m.providers = append(m.providers, &lazyInitProvider{
			fetcher: &KubernetesNodeMetadataFetcher{
				client:   client.CoreV1().Nodes(),
				nodeName: nodeName,
			},
		})
	}
	m.providers = append(m.providers, &lazyInitProvider{
		fetcher: &ProfileFetcher{
			client: client,
		},
	})
}

func (m *Metadata) EnableOpenAPI(ecsClient cloud.ECSv2Interface) {
	mPre := Metadata{
		// use the previous providers to get instance id,
		// do not recurse into ourselves
		providers: m.providers,
	}
	m.providers = append(m.providers, &lazyInitProvider{
		fetcher: &OpenAPIFetcher{
			ecsClient: ecsClient,
			mPre:      &mPre,
		},
	})
}

func (m *Metadata) EnableSts(stsClient cloud.STSInterface) {
	m.providers = append(m.providers, &lazyInitProvider{
		fetcher: &StsFetcher{
			stsClient: stsClient,
		},
	})
}

func (m *Metadata) Get(key MetadataKey) (string, error) {
	errors := []error{}
	for _, p := range m.providers {
		v, err := p.Get(key)
		if err == nil {
			return v, nil
		}
		if err == ErrUnknownMetadataKey {
			continue
		}
		errors = append(errors, err)
	}
	if len(errors) == 0 {
		return "", ErrUnknownMetadataKey
	}
	return "", utilerrors.NewAggregate(errors)
}

func MustGet(m MetadataProvider, key MetadataKey) string {
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
}

func (p FakeProvider) Get(key MetadataKey) (string, error) {
	if v, ok := p.Values[key]; ok {
		return v, nil
	}
	return "", ErrUnknownMetadataKey
}
