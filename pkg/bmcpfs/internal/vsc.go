/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package internal

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"

	ecsClient "github.com/alibabacloud-go/ecs-20140526/v7/client"
	efloClient "github.com/alibabacloud-go/eflo-controller-20221215/v3/client"
	nasclient "github.com/alibabacloud-go/nas-20170626/v4/client"
	"github.com/alibabacloud-go/tea/tea"
	"golang.org/x/time/rate"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/klog/v2"
	"k8s.io/utils/clock"
)

// vscDialect captures backend-specific values for VSC type/status enums
// that are conceptually identical between Lingjun (eflo) and ECS backends
// but use different string literals.
type vscDialect struct {
	PrimaryType  string // VSC type for "primary"
	StatusNormal string // status meaning the VSC is healthy/usable
}

var (
	efloVscDialect = vscDialect{
		PrimaryType:  "primary",
		StatusNormal: "Normal",
	}
	ecsVscDialect = vscDialect{
		PrimaryType:  "Primary",
		StatusNormal: "In_use",
	}
)

// isECSInstance reports whether the instanceId belongs to an ECS instance.
// ECS instance IDs are prefixed with "i-"; Lingjun node IDs are not.
func isECSInstance(instanceId string) bool {
	return strings.HasPrefix(instanceId, "i-")
}

type Vsc struct {
	NodeID string
	VscID  string
	Type   string
	Status string
}

// VscBackend is the interface each cloud backend implements independently.
type VscBackend interface {
	CreatePrimaryVsc(instanceId string) (string, error)
	GetPrimaryVscOf(instanceId string) (*Vsc, error)
	GetVsc(vscId string) (*Vsc, error)
}

type VscManager interface {
	CreatePrimaryVscFor(instanceId string) (string, error)
	GetPrimaryVscOf(instanceId string) (*Vsc, error)
	// GetVsc retrieves a single VSC by ID. The instanceId is used only for
	// routing to the correct backend; individual backends do not need it.
	GetVsc(vscId, instanceId string) (*Vsc, error)
}

func NewVscManager(eflo *efloClient.Client, ecs *ecsClient.Client) VscManager {
	return &dispatchingVscManager{
		eflo: &efloVscBackend{client: eflo},
		ecs:  &ecsVscBackend{client: ecs},
	}
}

// dispatchingVscManager routes calls to the correct backend based on instanceId.
type dispatchingVscManager struct {
	eflo VscBackend
	ecs  VscBackend
}

func (m *dispatchingVscManager) backendFor(instanceId string) VscBackend {
	if isECSInstance(instanceId) {
		return m.ecs
	}
	return m.eflo
}

func (m *dispatchingVscManager) CreatePrimaryVscFor(instanceId string) (string, error) {
	return m.backendFor(instanceId).CreatePrimaryVsc(instanceId)
}

func (m *dispatchingVscManager) GetPrimaryVscOf(instanceId string) (*Vsc, error) {
	return m.backendFor(instanceId).GetPrimaryVscOf(instanceId)
}

func (m *dispatchingVscManager) GetVsc(vscId, instanceId string) (*Vsc, error) {
	return m.backendFor(instanceId).GetVsc(vscId)
}

// efloVscBackend implements VscBackend for Lingjun (eflo) nodes.
type efloVscBackend struct {
	client *efloClient.Client
}

func (b *efloVscBackend) CreatePrimaryVsc(instanceId string) (string, error) {
	req := &efloClient.CreateVscRequest{
		NodeId:  &instanceId,
		VscType: tea.String(efloVscDialect.PrimaryType),
	}
	resp, err := b.client.CreateVsc(req)
	if err != nil {
		return "", fmt.Errorf("eflo:CreateVsc failed: %w", err)
	}
	klog.InfoS("eflo:CreateVsc succeeded", "instanceId", instanceId, "response", resp.Body)
	if tea.StringValue(resp.Body.VscId) == "" {
		return "", errors.New("unexpected response of eflo:CreateVsc")
	}
	return tea.StringValue(resp.Body.VscId), nil
}

func (b *efloVscBackend) GetPrimaryVscOf(instanceId string) (*Vsc, error) {
	req := &efloClient.ListVscsRequest{
		NodeIds:    []*string{&instanceId},
		MaxResults: tea.Int32(100),
	}
	resp, err := b.client.ListVscs(req)
	if err != nil {
		return nil, fmt.Errorf("eflo:ListVscs failed: %w", err)
	}
	klog.V(4).InfoS("eflo:ListVscs succeeded", "instanceId", instanceId, "response", resp.Body)
	for _, vsc := range resp.Body.Vscs {
		if tea.StringValue(vsc.VscType) == efloVscDialect.PrimaryType {
			return &Vsc{
				NodeID: tea.StringValue(vsc.NodeId),
				VscID:  tea.StringValue(vsc.VscId),
				Type:   tea.StringValue(vsc.VscType),
				Status: tea.StringValue(vsc.Status),
			}, nil
		}
	}
	return nil, nil
}

func (b *efloVscBackend) GetVsc(vscId string) (*Vsc, error) {
	req := &efloClient.DescribeVscRequest{
		VscId: &vscId,
	}
	resp, err := b.client.DescribeVsc(req)
	if err != nil {
		return nil, fmt.Errorf("eflo:DescribeVsc failed: %w", err)
	}
	klog.InfoS("eflo:DescribeVsc succeeded", "vscId", vscId, "response", resp.Body)
	return &Vsc{
		NodeID: tea.StringValue(resp.Body.NodeId),
		VscID:  tea.StringValue(resp.Body.VscId),
		Type:   tea.StringValue(resp.Body.VscType),
		Status: tea.StringValue(resp.Body.Status),
	}, nil
}

// ecsVscBackend implements VscBackend for ECS nodes with VSC enabled.
type ecsVscBackend struct {
	client *ecsClient.Client
}

func (b *ecsVscBackend) CreatePrimaryVsc(instanceId string) (string, error) {
	req := &ecsClient.CreateVscRequest{
		InstanceId: &instanceId,
		VscType:    tea.String(ecsVscDialect.PrimaryType),
	}
	resp, err := b.client.CreateVsc(req)
	if err != nil {
		return "", fmt.Errorf("ecs:CreateVsc failed: %w", err)
	}
	klog.InfoS("ecs:CreateVsc succeeded", "instanceId", instanceId, "response", resp.Body)
	if tea.StringValue(resp.Body.VscId) == "" {
		return "", errors.New("unexpected response of ecs:CreateVsc")
	}
	return tea.StringValue(resp.Body.VscId), nil
}

func (b *ecsVscBackend) GetPrimaryVscOf(instanceId string) (*Vsc, error) {
	req := &ecsClient.DescribeVscsRequest{
		InstanceId: &instanceId,
		MaxResults: tea.Int32(100),
	}
	resp, err := b.client.DescribeVscs(req)
	if err != nil {
		return nil, fmt.Errorf("ecs:DescribeVscs failed: %w", err)
	}
	klog.V(4).InfoS("ecs:DescribeVscs succeeded", "instanceId", instanceId, "response", resp.Body)
	for _, vsc := range resp.Body.Vscs {
		if tea.StringValue(vsc.VscType) == ecsVscDialect.PrimaryType {
			return &Vsc{
				NodeID: tea.StringValue(vsc.InstanceId),
				VscID:  tea.StringValue(vsc.VscId),
				Type:   tea.StringValue(vsc.VscType),
				Status: tea.StringValue(vsc.Status),
			}, nil
		}
	}
	return nil, nil
}

func (b *ecsVscBackend) GetVsc(vscId string) (*Vsc, error) {
	req := &ecsClient.DescribeVscsRequest{
		VscIds: []*string{&vscId},
	}
	resp, err := b.client.DescribeVscs(req)
	if err != nil {
		return nil, fmt.Errorf("ecs:DescribeVscs failed: %w", err)
	}
	klog.V(4).InfoS("ecs:DescribeVscs succeeded", "vscId", vscId, "response", resp.Body)
	if len(resp.Body.Vscs) == 0 {
		return nil, nil
	}
	return &Vsc{
		NodeID: tea.StringValue(resp.Body.Vscs[0].InstanceId),
		VscID:  tea.StringValue(resp.Body.Vscs[0].VscId),
		Type:   tea.StringValue(resp.Body.Vscs[0].VscType),
		Status: tea.StringValue(resp.Body.Vscs[0].Status),
	}, nil
}

type vscWithErr struct {
	*Vsc
	err      error
	cachedAt time.Time
}

func (v *vscWithErr) isExpired(ttl time.Duration) bool {
	return time.Since(v.cachedAt) > ttl
}

type PrimaryVscManagerWithCache struct {
	VscManager
	retryTimes int
	cacheTTL   time.Duration

	cond *sync.Cond
	// Instance ID to VSC
	cache map[string]vscWithErr
	// To create primary vsc for node
	queue workqueue.TypedRateLimitingInterface[string]
}

const (
	defaultVscManagerRetryTimes  = 3
	defaultVscManagerWorkerCount = 3
	defaultVscCacheTTL           = 3 * time.Minute
)

func NewPrimaryVscManagerWithCache(efloClient *efloClient.Client, ecsClient *ecsClient.Client) *PrimaryVscManagerWithCache {
	m := &PrimaryVscManagerWithCache{
		VscManager: NewVscManager(efloClient, ecsClient),
		retryTimes: defaultVscManagerRetryTimes,
		cacheTTL:   defaultVscCacheTTL,
		cond:       sync.NewCond(&sync.Mutex{}),
		cache:      make(map[string]vscWithErr),
		queue: workqueue.NewTypedRateLimitingQueue(
			workqueue.NewTypedMaxOfRateLimiter(
				workqueue.NewTypedItemExponentialFailureRateLimiter[string](500*time.Millisecond, 1000*time.Second),
				&workqueue.TypedBucketRateLimiter[string]{Limiter: rate.NewLimiter(rate.Limit(10), 100)},
			),
		),
	}

	for range defaultVscManagerWorkerCount {
		go func() {
			for m.handleNext() {
			}
		}()
	}
	return m
}

func (m *PrimaryVscManagerWithCache) handleNext() bool {
	instanceId, quit := m.queue.Get()
	if quit {
		return false
	}
	defer m.queue.Done(instanceId)

	newVsc, err := m.getOrCreatePrimaryFor(instanceId)

	m.cond.L.Lock()
	m.cache[instanceId] = vscWithErr{Vsc: newVsc, err: err, cachedAt: time.Now()}
	m.cond.L.Unlock()

	if err == nil {
		m.queue.Forget(instanceId)
		m.cond.Broadcast()
	} else {
		sdkErr := &tea.SDKError{}
		if errors.As(err, &sdkErr) || m.queue.NumRequeues(instanceId) > m.retryTimes {
			klog.ErrorS(err, "Failed to ensure VSC", "instance", instanceId)
			m.queue.Forget(instanceId)
			m.cond.Broadcast()
		} else {
			klog.InfoS("Retrying to ensure VSC", "instance", instanceId, "error", err)
			m.queue.AddRateLimited(instanceId)
		}
	}
	return true
}

func (m *PrimaryVscManagerWithCache) getOrCreatePrimaryFor(instanceId string) (*Vsc, error) {
	var err error
	// try to get existing vsc
	vsc, err := m.VscManager.GetPrimaryVscOf(instanceId)
	if err != nil {
		return nil, err
	}
	// primary vsc of the instance not found, create it
	var vscId string
	if vsc == nil {
		vscId, err = m.CreatePrimaryVscFor(instanceId)
		if err != nil {
			return nil, err
		}
		vsc, err = m.GetVsc(vscId, instanceId)
		if err != nil {
			return nil, err
		}
	}
	if vsc == nil {
		return nil, fmt.Errorf("vsc %s not found after creation", vscId)
	}
	// check vsc status
	expected := efloVscDialect.StatusNormal
	if isECSInstance(instanceId) {
		expected = ecsVscDialect.StatusNormal
	}
	if vsc.Status != expected {
		return vsc, fmt.Errorf("unexpected vsc status: %s", vsc.Status)
	}
	return vsc, nil
}

func (m *PrimaryVscManagerWithCache) EnsurePrimaryVsc(ctx context.Context, instanceId string, refresh bool) (string, error) {
	m.cond.L.Lock()
	defer m.cond.L.Unlock()

	if !refresh {
		vsc, exists := m.cache[instanceId]
		if exists && vsc.err == nil && !vsc.isExpired(m.cacheTTL) {
			return vsc.VscID, nil
		}
	}

	delete(m.cache, instanceId)
	m.queue.Add(instanceId)
	for {
		vsc, exists := m.cache[instanceId]
		if exists {
			var vscId string
			if vsc.Vsc != nil {
				vscId = vsc.VscID
			}
			return vscId, vsc.err
		}
		select {
		case <-ctx.Done():
			return "", ctx.Err()
		default:
			m.cond.Wait()
		}
	}
}

func (m *PrimaryVscManagerWithCache) GetPrimaryVscOf(instanceId string) (*Vsc, error) {
	m.cond.L.Lock()
	cachedVsc, exists := m.cache[instanceId]
	if exists && cachedVsc.Vsc != nil && !cachedVsc.isExpired(m.cacheTTL) {
		m.cond.L.Unlock()
		return cachedVsc.Vsc, nil
	}
	m.cond.L.Unlock()

	vsc, err := m.VscManager.GetPrimaryVscOf(instanceId)
	if err != nil {
		return nil, err
	}

	// update the cache
	if vsc != nil {
		m.cond.L.Lock()
		clonedVsc := new(Vsc)
		*clonedVsc = *vsc
		m.cache[instanceId] = vscWithErr{Vsc: clonedVsc, err: nil, cachedAt: time.Now()}
		m.cond.L.Unlock()
	}

	return vsc, nil
}

const (
	CPFSVscStatusAttaching = "Attaching"
	CPFSVscStatusAttached  = "Attached"
	CPFSVscStatusDetaching = "Detaching"
	CPFSVscStatusDetached  = "Detached"
	CPFSVscStatusFailed    = "Failed"

	VscAttachNotSupported = "AttachVscTarget.VscAttachNotSupported"
)

const (
	defaultPollInterval  = time.Second * 2
	defaultADWaitTimeout = time.Second * 10
)

func newAttachNotSupportedError(err error, volumeId, nodeId string) *AttachNotSupportedError {
	return &AttachNotSupportedError{
		message:  err.Error(),
		volumeId: volumeId,
		vscId:    nodeId,
	}
}

// NewAttachNotSupportedError creates a new AttachNotSupportedError
func NewAttachNotSupportedError(err error, volumeId, nodeId string) *AttachNotSupportedError {
	return newAttachNotSupportedError(err, volumeId, nodeId)
}

type AttachNotSupportedError struct {
	message  string
	volumeId string
	vscId    string
}

func (e *AttachNotSupportedError) Error() string {
	return "volumeID: " + e.volumeId + "vscId: " + e.vscId + e.message
}

func IsAttachNotSupportedError(err error) bool {
	if err == nil {
		return false
	}
	var attachErr *AttachNotSupportedError
	if errors.As(err, &attachErr) {
		return true
	}
	sdkErr := &tea.SDKError{}
	return errors.As(err, &sdkErr) && tea.StringValue(sdkErr.Code) == VscAttachNotSupported
}

type CPFSVscAttachInfo = nasclient.DescribeFilesystemsVscAttachInfoResponseBodyVscAttachInfoVscAttachInfo

type CPFSVscAttachInfoCond func(*CPFSVscAttachInfo) (done bool, err error)

type CPFSAttachDetacher interface {
	Attach(ctx context.Context, fsId, vscId string) error
	Detach(ctx context.Context, fsId, vscId string) error
}

func NewCPFSAttachDetacher(client *nasclient.Client) CPFSAttachDetacher {
	return &cpfsAttachDetacher{
		client:       client,
		pollInterval: defaultPollInterval,
		clk:          clock.RealClock{},
		waitTimeout:  defaultADWaitTimeout,
	}
}

type cpfsAttachDetacher struct {
	client       *nasclient.Client
	pollInterval time.Duration
	clk          clock.WithTicker
	waitTimeout  time.Duration
}

func (ad *cpfsAttachDetacher) Attach(ctx context.Context, fsId, vscId string) error {
	attachInfo, err := ad.describe(fsId, vscId)
	if err != nil {
		return err
	}
	if attachInfo != nil {
		klog.InfoS("Already attached", "filesystem", fsId, "vscId", vscId)
		switch tea.StringValue(attachInfo.Status) {
		case CPFSVscStatusAttaching:
		case CPFSVscStatusAttached:
			return nil
		default:
			return fmt.Errorf("unexpected attachinfo status: %v", tea.StringValue(attachInfo.Status))
		}
	} else {
		if err := ad.attach(fsId, vscId); err != nil {
			if strings.Contains(err.Error(), VscAttachNotSupported) {
				return newAttachNotSupportedError(err, fsId, vscId)
			}
			return err
		}
	}
	return ad.waitFor(ctx, fsId, vscId, func(i *CPFSVscAttachInfo) (bool, error) {
		if i == nil {
			return false, fmt.Errorf("filesystem %s not attached to %s", fsId, vscId)
		}
		switch tea.StringValue(i.Status) {
		case CPFSVscStatusAttaching:
			return false, nil
		case CPFSVscStatusAttached:
			return true, nil
		default:
			return false, fmt.Errorf("unexpected attachinfo status: %v", tea.StringValue(i.Status))
		}
	}, "wait for cpfs to be attached")
}

func (ad *cpfsAttachDetacher) Detach(ctx context.Context, fsId, vscId string) error {
	if err := ad.detach(fsId, vscId); err != nil {
		sdkErr := new(tea.SDKError)
		if errors.As(err, &sdkErr) {
			errCode := tea.StringValue(sdkErr.Code)
			// attached by legacy inner api, ignore it
			if errCode == "Resource.Check.Fail" || errCode == "InvalidFileSystem.NotFound" {
				klog.InfoS("Ignore detaching error", "error", err)
				return nil
			}
		}
		return err
	}
	return ad.waitFor(ctx, fsId, vscId, func(i *CPFSVscAttachInfo) (bool, error) {
		if i == nil {
			return true, nil
		}
		switch tea.StringValue(i.Status) {
		case CPFSVscStatusDetaching:
			return false, nil
		case CPFSVscStatusDetached:
			return true, nil
		default:
			return false, fmt.Errorf("unexpected attachinfo status: %v", tea.StringValue(i.Status))
		}
	}, "wait for cpfs to be detached")
}

func (ad *cpfsAttachDetacher) waitFor(ctx context.Context, fsId, vscId string, cond CPFSVscAttachInfoCond, cause string) error {
	deadline := ad.clk.NewTimer(ad.waitTimeout)
	ticker := ad.clk.NewTicker(ad.pollInterval)
	defer ticker.Stop()
	for {
		attachInfo, err := ad.describe(fsId, vscId)
		if err != nil {
			return err
		}
		done, err := cond(attachInfo)
		if err != nil {
			return err
		}
		if done {
			return nil
		}
		select {
		case <-ticker.C():
		case <-deadline.C():
			return fmt.Errorf("%s: timeout", cause)
		case <-ctx.Done():
			return fmt.Errorf("%s: %w", cause, ctx.Err())
		}
	}
}

func (ad *cpfsAttachDetacher) attach(fsId, vscId string) error {
	req := &nasclient.AttachVscToFilesystemsRequest{
		ResourceIds: []*nasclient.AttachVscToFilesystemsRequestResourceIds{
			{
				FileSystemId: &fsId,
				VscId:        &vscId,
			},
		},
	}
	resp, err := ad.client.AttachVscToFilesystems(req)
	if err != nil {
		return fmt.Errorf("nas:AttachVscToFilesystems failed: %w", err)
	}
	klog.InfoS("nas:AttachVscToFilesystemsRequest succeeded", "filesystem", fsId, "vscId", vscId, "requestid", resp.Body.RequestId)
	return nil
}

func (ad *cpfsAttachDetacher) detach(fsId, vscId string) error {
	req := &nasclient.DetachVscFromFilesystemsRequest{
		ResourceIds: []*nasclient.DetachVscFromFilesystemsRequestResourceIds{
			{
				FileSystemId: &fsId,
				VscId:        &vscId,
			},
		},
	}
	resp, err := ad.client.DetachVscFromFilesystems(req)
	if err != nil {
		return fmt.Errorf("nas:DetachVscFromFilesystems failed: %w", err)
	}
	klog.InfoS("nas:DetachVscFromFilesystems succeeded", "filesystem", fsId, "vscId", vscId, "requestid", resp.Body.RequestId)
	return nil
}

func (ad *cpfsAttachDetacher) describe(fsId, vscId string) (*CPFSVscAttachInfo, error) {
	req := &nasclient.DescribeFilesystemsVscAttachInfoRequest{
		ResourceIds: []*nasclient.DescribeFilesystemsVscAttachInfoRequestResourceIds{
			{
				FileSystemId: &fsId,
				VscId:        &vscId,
			},
		},
	}
	resp, err := ad.client.DescribeFilesystemsVscAttachInfo(req)
	if err != nil {
		return nil, fmt.Errorf("nas:DescribeFilesystemsVscAttachInfo failed: %w", err)
	}
	klog.V(4).InfoS("nas:DescribeFilesystemsVscAttachInfo succeeded", "response", resp.Body)
	if resp.Body.VscAttachInfo == nil || len(resp.Body.VscAttachInfo.VscAttachInfo) == 0 {
		return nil, nil
	}
	return resp.Body.VscAttachInfo.VscAttachInfo[0], nil
}
