package metadata

import (
	"context"
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
	diskQuantity // int32
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
	case diskQuantity:
		return "DiskQuantity"
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

func (p inferMachineKind) GetAny(ctx *mcontext, key MetadataKey) (any, error) {
	v, err := p.next.GetAny(ctx, key)
	if key == machineKind && err == ErrUnknownMetadataKey {
		t, err := p.next.GetAny(ctx, InstanceType)
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

const DISABLE_IMDS_ENV = "ALIBABA_CLOUD_NO_ECS_METADATA"
const KUBE_NODE_NAME_ENV = "KUBE_NODE_NAME"

type strMetadataProvider interface {
	Get(key MetadataKey) (string, error)
}

type MetadataProvider interface {
	strMetadataProvider
	MachineKind() (MachineKind, error)
	DiskQuantity() (int32, error)
	WithSession(ctx context.Context) MetadataProvider
}

type fetcherID uint

const (
	imdsFetcherID fetcherID = iota
	efloFetcherID
	kubernetesNodeMetadataFetcherID
	profileFetcherID
	openAPIFetcherID
	ecsInstanceTypeFetcherID
	stsFetcherID

	numFetchers
)

type Metadata struct {
	providers multi
	ctx       context.Context
	errors    [numFetchers]error
}

type mcontext struct {
	context.Context
	inMemory bool
	logger   klog.Logger
	errors   *[numFetchers]error
}

func (m *Metadata) WithSession(ctx context.Context) MetadataProvider {
	return &Metadata{
		providers: m.providers,
		ctx:       ctx,
	}
}

func get(m *Metadata, key MetadataKey) (any, error) {
	ctx := mcontext{
		Context: m.ctx,
		logger:  klog.FromContext(m.ctx),
		errors:  &m.errors,
		// Try inMemory first.
		// If the data is already in memory, we don't need to make network requests.
		inMemory: true,
	}
	v, err := m.providers.GetAny(&ctx, key)
	if errors.Is(err, ErrUnknownMetadataKey) {
		ctx.inMemory = false // if not found, allow fetch from remote
		v, err = m.providers.GetAny(&ctx, key)
	}
	return v, err
}

func getT[T any](m *Metadata, key MetadataKey) (T, error) {
	v, err := get(m, key)
	if err != nil {
		var zero T
		return zero, err
	}
	return v.(T), nil
}

func (m *Metadata) Get(key MetadataKey) (string, error) { return getT[string](m, key) }
func (m *Metadata) MachineKind() (MachineKind, error)   { return getT[MachineKind](m, machineKind) }
func (m *Metadata) DiskQuantity() (int32, error)        { return getT[int32](m, diskQuantity) }

type fetcher interface {
	FetchFor(ctx *mcontext, key MetadataKey) (middleware, error)
	ID() fetcherID
}

type middleware interface {
	GetAny(ctx *mcontext, key MetadataKey) (any, error)
}

type lazyInit struct {
	next    middleware
	initMu  sync.Mutex
	fetcher fetcher
}

func (p *lazyInit) init(ctx *mcontext, key MetadataKey) error {
	p.initMu.Lock()
	defer p.initMu.Unlock()

	if p.next != nil {
		return nil // already initialized
	}
	if ctx.inMemory { // not allowed to fetch
		return ErrUnknownMetadataKey
	}

	ctxErr := &ctx.errors[p.fetcher.ID()]
	if *ctxErr != nil {
		return *ctxErr // last attempt failed
	}

	next, err := p.fetcher.FetchFor(ctx, key)
	if err == ErrUnknownMetadataKey {
		return err
	} else if err != nil {
		err = fmt.Errorf("%T failed: %w", p.fetcher, err)
		*ctxErr = err
		return err
	} else {
		p.fetcher = nil
		p.next = next
		return nil
	}
}

func (p *lazyInit) GetAny(ctx *mcontext, key MetadataKey) (any, error) {
	err := p.init(ctx, key)
	if err != nil {
		return nil, err
	}
	return p.next.GetAny(ctx, key)
}

type immutable struct {
	next   middleware
	name   string
	mu     sync.Mutex
	values map[MetadataKey]any
}

func (p *immutable) GetAny(ctx *mcontext, key MetadataKey) (any, error) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if v, ok := p.values[key]; ok {
		return v, nil
	}
	v, err := p.next.GetAny(ctx, key)
	if err != nil {
		return "", err
	}
	ctx.logger.V(2).Info("retrieved metadata", "provider", p.name, "key", key, "value", v)
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

func (p strProvider) GetAny(_ *mcontext, key MetadataKey) (any, error) {
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
	return &Metadata{
		providers: providers,
		ctx:       context.Background(),
	}
}

func (m *Metadata) EnableIMDS(httpRT http.RoundTripper) {
	if os.Getenv(DISABLE_IMDS_ENV) != "" {
		klog.Infof("ECS metadata is disabled by environment variable %s", DISABLE_IMDS_ENV)
		return
	}
	m.providers = append(m.providers, inferMachineKind{&lazyInit{
		fetcher: &IMDSFetcer{httpRT: httpRT},
	}}, NewIMDSDynamic(httpRT))
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
	// use the previous providers to get instance id,
	// do not recurse into ourselves
	mPre := m.providers
	m.providers = append(m.providers, inferMachineKind{&lazyInit{
		fetcher: &OpenAPIFetcher{
			ecsClient: ecsClient,
			mPre:      mPre,
		},
	}})
	mPre2 := m.providers // be sure to include OpenAPIFetcher for instance type
	m.providers = append(m.providers, &lazyInit{
		fetcher: &ECSInstanceTypeFetcher{
			ecsClient: ecsClient,
			mPre:      mPre2,
		},
	})
}

func (m *Metadata) EnableSts(stsClient cloud.STSInterface) {
	m.providers = append(m.providers, &lazyInit{
		fetcher: &StsFetcher{
			stsClient: stsClient,
		},
	})
}

func (m *Metadata) EnableEFLO(efloClient cloud.EFLOInterface) {
	// use the previous providers to get instance id,
	// do not recurse into ourselves
	mPre := m.providers
	m.providers = append(m.providers, &lazyInit{
		fetcher: &EfloFetcher{
			efloClient: efloClient,
			mPre:       mPre,
		},
	})
}

type multi []middleware

func (m multi) GetAny(ctx *mcontext, key MetadataKey) (any, error) {
	errors := []error{}
	for _, p := range m {
		v, err := p.GetAny(ctx, key)
		if err == nil {
			return v, nil
		}
		if err == ErrUnknownMetadataKey {
			continue
		}
		// print a warning if we failed to get a value,
		// because the error is hidden if other providers succeed
		ctx.logger.Error(err, "metadata provider failed", "key", key)
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
		MachineKind  MachineKind
		DiskQuantity int32
	}
}

func (p *FakeProvider) Get(key MetadataKey) (string, error) {
	if v, ok := p.Values[key]; ok {
		return v, nil
	}
	return "", ErrUnknownMetadataKey
}

func (p *FakeProvider) MachineKind() (MachineKind, error) {
	return p.V.MachineKind, nil
}

func (p *FakeProvider) DiskQuantity() (int32, error) {
	return p.V.DiskQuantity, nil
}

func (p *FakeProvider) WithSession(ctx context.Context) MetadataProvider {
	return p
}
