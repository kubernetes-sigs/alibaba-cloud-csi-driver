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
	"time"

	ecsClient "github.com/alibabacloud-go/ecs-20140526/v7/client"
	efloClient "github.com/alibabacloud-go/eflo-controller-20221215/v3/client"
	nasclient "github.com/alibabacloud-go/nas-20170626/v4/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/throttle"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils/ttlcache"
	"k8s.io/klog/v2"
	"k8s.io/utils/clock"
)

// newVscThrottler builds a reactive backoff throttler for a single VSC OpenAPI
// action. VSC clients use the new tea SDK, so we classify with V2Classifier.
// Delays mirror pkg/disk's defaults.
func newVscThrottler() *throttle.Throttler {
	return throttle.NewThrottler(clock.RealClock{}, 1*time.Second, 10*time.Second, throttle.V2Classifier)
}

// Backend-specific wire values for the VSC "primary" type and the healthy/usable
// status. The concepts are identical across Lingjun (eflo) and ECS; only the string
// literals differ. Each backend uses its own pair directly.
const (
	efloPrimaryVscType  = "primary" // VSC type for "primary"
	efloVscStatusNormal = "Normal"  // status meaning the VSC is healthy/usable

	ecsPrimaryVscType  = "Primary"
	ecsVscStatusNormal = "In_use"
)

type Vsc struct {
	NodeID string
	VscID  string
	Type   string
	Status string
}

// Backend is the interface each cloud backend (Lingjun/eflo or ECS) implements.
// It holds all per-backend shared state (client, throttlers) and owns its own
// status/type dialect via Ready. A Backend is constructed once and shared across
// every instance it serves.
type Backend interface {
	CreatePrimaryVsc(ctx context.Context, instanceId string) (string, error)
	GetPrimaryVsc(ctx context.Context, instanceId string) (*Vsc, error)
	GetVsc(ctx context.Context, vscId string) (*Vsc, error)
	// Ready reports whether a VSC has reached this backend's usable status.
	Ready(*Vsc) bool
}

// Instance binds an instance/node ID to the Backend that serves it. It is a cheap
// value: Backend carries the shared state, so callers build an Instance per request
// for free. The caller picks the Backend (the CSI layer does so from the node kind
// it parsed); this package does not inspect the ID to choose a Backend.
type Instance struct {
	ID      string
	Backend Backend
}

func (i Instance) createPrimary(ctx context.Context) (string, error) {
	return i.Backend.CreatePrimaryVsc(ctx, i.ID)
}

func (i Instance) getPrimary(ctx context.Context) (*Vsc, error) {
	return i.Backend.GetPrimaryVsc(ctx, i.ID)
}

// efloVscBackend implements Backend for Lingjun (eflo) nodes.
type efloVscBackend struct {
	client            *efloClient.Client
	createThrottler   *throttle.Throttler
	listThrottler     *throttle.Throttler
	describeThrottler *throttle.Throttler
}

// NewEfloBackend builds the Backend for Lingjun (eflo) nodes. Construct it once and
// share it: the throttlers are per-action and must be shared across all instances.
func NewEfloBackend(client *efloClient.Client) Backend {
	return &efloVscBackend{
		client:            client,
		createThrottler:   newVscThrottler(),
		listThrottler:     newVscThrottler(),
		describeThrottler: newVscThrottler(),
	}
}

func (b *efloVscBackend) Ready(v *Vsc) bool {
	return v.Status == efloVscStatusNormal
}

func (b *efloVscBackend) CreatePrimaryVsc(ctx context.Context, instanceId string) (string, error) {
	req := &efloClient.CreateVscRequest{
		NodeId:  &instanceId,
		VscType: new(efloPrimaryVscType),
	}
	resp, err := throttle.Throttled(b.createThrottler, b.client.CreateVsc)(ctx, req)
	if err != nil {
		return "", fmt.Errorf("eflo:CreateVsc failed: %w", err)
	}
	klog.InfoS("eflo:CreateVsc succeeded", "instanceId", instanceId, "response", resp.Body)
	if tea.StringValue(resp.Body.VscId) == "" {
		return "", errors.New("unexpected response of eflo:CreateVsc")
	}
	return tea.StringValue(resp.Body.VscId), nil
}

func (b *efloVscBackend) GetPrimaryVsc(ctx context.Context, instanceId string) (*Vsc, error) {
	req := &efloClient.ListVscsRequest{
		NodeIds:    []*string{&instanceId},
		MaxResults: tea.Int32(100),
	}
	resp, err := throttle.Throttled(b.listThrottler, b.client.ListVscs)(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("eflo:ListVscs failed: %w", err)
	}
	klog.V(4).InfoS("eflo:ListVscs succeeded", "instanceId", instanceId, "response", resp.Body)
	for _, vsc := range resp.Body.Vscs {
		if tea.StringValue(vsc.VscType) == efloPrimaryVscType {
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

func (b *efloVscBackend) GetVsc(ctx context.Context, vscId string) (*Vsc, error) {
	req := &efloClient.DescribeVscRequest{
		VscId: &vscId,
	}
	resp, err := throttle.Throttled(b.describeThrottler, b.client.DescribeVsc)(ctx, req)
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

// ecsVscBackend implements Backend for ECS nodes with VSC enabled.
type ecsVscBackend struct {
	client            *ecsClient.Client
	createThrottler   *throttle.Throttler
	describeThrottler *throttle.Throttler
}

// NewEcsBackend builds the Backend for ECS nodes with VSC enabled. Construct it once
// and share it: the throttlers are per-action and must be shared across all instances.
func NewEcsBackend(client *ecsClient.Client) Backend {
	return &ecsVscBackend{
		client:            client,
		createThrottler:   newVscThrottler(),
		describeThrottler: newVscThrottler(),
	}
}

func (b *ecsVscBackend) Ready(v *Vsc) bool {
	return v.Status == ecsVscStatusNormal
}

func (b *ecsVscBackend) CreatePrimaryVsc(ctx context.Context, instanceId string) (string, error) {
	req := &ecsClient.CreateVscRequest{
		InstanceId: &instanceId,
		VscType:    new(ecsPrimaryVscType),
	}
	resp, err := throttle.Throttled(b.createThrottler, b.client.CreateVsc)(ctx, req)
	if err != nil {
		return "", fmt.Errorf("ecs:CreateVsc failed: %w", err)
	}
	klog.InfoS("ecs:CreateVsc succeeded", "instanceId", instanceId, "response", resp.Body)
	if tea.StringValue(resp.Body.VscId) == "" {
		return "", errors.New("unexpected response of ecs:CreateVsc")
	}
	return tea.StringValue(resp.Body.VscId), nil
}

func (b *ecsVscBackend) GetPrimaryVsc(ctx context.Context, instanceId string) (*Vsc, error) {
	req := &ecsClient.DescribeVscsRequest{
		InstanceId: &instanceId,
		MaxResults: tea.Int32(100),
	}
	resp, err := throttle.Throttled(b.describeThrottler, b.client.DescribeVscs)(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("ecs:DescribeVscs failed: %w", err)
	}
	klog.V(4).InfoS("ecs:DescribeVscs succeeded", "instanceId", instanceId, "response", resp.Body)
	for _, vsc := range resp.Body.Vscs {
		if tea.StringValue(vsc.VscType) == ecsPrimaryVscType {
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

func (b *ecsVscBackend) GetVsc(ctx context.Context, vscId string) (*Vsc, error) {
	req := &ecsClient.DescribeVscsRequest{
		VscIds: []*string{&vscId},
	}
	resp, err := throttle.Throttled(b.describeThrottler, b.client.DescribeVscs)(ctx, req)
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

// errVscNotFound is a sentinel returned by the readCache producer when the
// backend reports no primary VSC for an instance. TTLCache does not cache
// errors, so returning it keeps "not found" out of the positive cache (the
// absence might be transient) while still deduplicating concurrent lookups.
var errVscNotFound = errors.New("primary vsc not found")

// PrimaryVscManagerWithCache adds single-flight TTL caching on top of the VSC
// Backends. It is backend-agnostic: callers pass a resolved Instance (an ID paired
// with its Backend) and this layer only caches, single-flights and polls — it does
// no backend selection itself. Concurrency is reactive: requests are sent freely and
// each backend action backs off only when the cloud actually throttles it (see
// newVscThrottler), rather than being capped by a fixed worker pool.
type PrimaryVscManagerWithCache struct {
	// vscCache is the single VSC cache, shared by GetPrimaryVscOf (read-through)
	// and EnsurePrimaryVsc (write-through on a confirmed VSC), so a publish warms
	// the cache that a later unpublish reads. Keyed by instance ID; ECS ("i-…") and
	// Lingjun IDs never collide, so one cache safely spans both backends.
	vscCache *ttlcache.TTLCache[string, *Vsc]
	// createVsc single-flights EnsurePrimaryVsc per instance so concurrent
	// publishes for a new node create only one VSC. Its TTL is 0: it never caches
	// (vscCache is the cache), it only serializes in-flight creates.
	createVsc *ttlcache.TTLCache[string, *Vsc]

	clk          clock.Clock
	pollInterval time.Duration
	// pollAttempts is how many times getOrCreatePrimaryFor polls GetVsc for the
	// ready status. Must be >= 1, so a freshly created VSC (not fetched before
	// the loop) is always fetched at least once.
	pollAttempts int
}

const (
	defaultVscCacheTTL     = 3 * time.Minute
	defaultVscPollInterval = 2 * time.Second
	// defaultVscPollAttempts bounds in-process polling of a freshly created VSC
	// until it reaches the expected status.
	defaultVscPollAttempts = 5
)

func NewPrimaryVscManagerWithCache() *PrimaryVscManagerWithCache {
	return &PrimaryVscManagerWithCache{
		vscCache:     ttlcache.NewTTLCache[string, *Vsc](defaultVscCacheTTL),
		createVsc:    ttlcache.NewTTLCache[string, *Vsc](0),
		clk:          clock.RealClock{},
		pollInterval: defaultVscPollInterval,
		pollAttempts: defaultVscPollAttempts,
	}
}

func (m *PrimaryVscManagerWithCache) EnsurePrimaryVsc(ctx context.Context, inst Instance) (string, error) {
	vsc, err := m.createVsc.Get(ctx, inst.ID, func() (*Vsc, error) {
		// Decouple the shared computation from any single caller's cancellation;
		// concurrent waiters still cancel their own wait via the ctx passed to Get.
		return m.getOrCreatePrimaryFor(context.WithoutCancel(ctx), inst)
	})
	if err != nil {
		return "", err
	}
	return vsc.VscID, nil
}

func (m *PrimaryVscManagerWithCache) GetPrimaryVscOf(ctx context.Context, inst Instance) (*Vsc, error) {
	vsc, err := m.vscCache.Get(ctx, inst.ID, func() (*Vsc, error) {
		vsc, err := inst.getPrimary(ctx)
		if err != nil {
			return nil, err
		}
		if vsc == nil {
			return nil, errVscNotFound
		}
		return vsc, nil
	})
	if errors.Is(err, errVscNotFound) {
		return nil, nil
	}
	return vsc, err
}

func (m *PrimaryVscManagerWithCache) getOrCreatePrimaryFor(ctx context.Context, inst Instance) (*Vsc, error) {
	// try to get existing vsc (through the read cache)
	vsc, err := m.GetPrimaryVscOf(ctx, inst)
	if err != nil {
		return nil, err
	}

	var vscID string
	if vsc != nil {
		if inst.Backend.Ready(vsc) {
			return vsc, nil // already usable; GetPrimaryVscOf cached it
		}
		vscID = vsc.VscID
	} else {
		// primary vsc of the instance not found, create it
		vscID, err = inst.createPrimary(ctx)
		if err != nil {
			return nil, err
		}
	}

	// A freshly created (or not-yet-ready) VSC may take a moment to become
	// usable; poll GetVsc a bounded number of times until it becomes ready.
	for range m.pollAttempts {
		select {
		case <-ctx.Done():
			return vsc, ctx.Err()
		case <-m.clk.After(m.pollInterval):
		}
		if vsc, err = inst.Backend.GetVsc(ctx, vscID); err != nil {
			return nil, err
		}
		if vsc == nil {
			return nil, fmt.Errorf("vsc %s not found after creation", vscID)
		}
		if inst.Backend.Ready(vsc) {
			m.vscCache.Store(inst.ID, vsc)
			return vsc, nil
		}
	}
	// pollAttempts >= 1, so the loop ran at least once and vsc is non-nil here.
	return vsc, fmt.Errorf("unexpected vsc status: %s", vsc.Status)
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
	if _, ok := errors.AsType[*AttachNotSupportedError](err); ok {
		return true
	}
	sdkErr, ok := errors.AsType[*tea.SDKError](err)
	return ok && tea.StringValue(sdkErr.Code) == VscAttachNotSupported
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
		if sdkErr, ok := errors.AsType[*tea.SDKError](err); ok {
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
