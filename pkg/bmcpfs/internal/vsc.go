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

	efloclient "github.com/alibabacloud-go/eflo-controller-20221215/v3/client"
	nasclient "github.com/alibabacloud-go/nas-20170626/v4/client"
	"github.com/alibabacloud-go/tea/tea"
	"golang.org/x/time/rate"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/klog/v2"
	"k8s.io/utils/clock"
)

const (
	// VscTypePrimary is the type for primary VSC
	VscTypePrimary = "primary"

	// VscStatusCreating is the status when VSC is being created
	VscStatusCreating = "Creating"
	// VscStatusNormal is the status when VSC is normal
	VscStatusNormal = "Normal"
	// VscStatusDeleting is the status when VSC is being deleted
	VscStatusDeleting = "Deleting"
)

// Vsc represents a VSC (Virtual Switch Cluster)
type Vsc struct {
	NodeID string
	VscID  string
	Type   string
	Status string
}

// VscManager is the interface for managing VSCs
type VscManager interface {
	CreatePrimaryVscFor(instanceID string) (string, error)
	GetPrimaryVscOf(instanceID string) (*Vsc, error)
	GetVsc(vscID string) (*Vsc, error)
}

// NewVscManager creates a new VSC manager
func NewVscManager(client *efloclient.Client) VscManager {
	return &LingjunVscManager{client: client}
}

// LingjunVscManager is the implementation of VscManager for Lingjun
type LingjunVscManager struct {
	client *efloclient.Client
}

// CreatePrimaryVscFor creates a primary VSC for the given instance
func (m *LingjunVscManager) CreatePrimaryVscFor(instanceID string) (string, error) {
	req := &efloclient.CreateVscRequest{
		NodeId:  &instanceID,
		VscType: tea.String(VscTypePrimary),
	}
	resp, err := m.client.CreateVsc(req)
	if err != nil {
		return "", fmt.Errorf("eflo:CreateVsc failed: %w", err)
	}
	klog.InfoS("eflo:CreateVsc succeeded", "instanceId", instanceID, "response", resp.Body)
	if tea.StringValue(resp.Body.VscId) == "" {
		return "", errors.New("unexpected response of CreateVsc")
	}
	return tea.StringValue(resp.Body.VscId), nil
}

// GetPrimaryVscOf gets the primary VSC of the given instance
func (m *LingjunVscManager) GetPrimaryVscOf(instanceID string) (*Vsc, error) {
	req := &efloclient.ListVscsRequest{
		NodeIds:    []*string{&instanceID},
		MaxResults: tea.Int32(100),
	}
	resp, err := m.client.ListVscs(req)
	if err != nil {
		return nil, fmt.Errorf("eflo:ListVscs failed: %w", err)
	}
	klog.V(4).InfoS("eflo:ListVscs succeeded", "instanceId", instanceID, "response", resp.Body)
	for _, vsc := range resp.Body.Vscs {
		if tea.StringValue(vsc.VscType) == VscTypePrimary {
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

// GetVsc gets the VSC with the given ID
func (m *LingjunVscManager) GetVsc(vscID string) (*Vsc, error) {
	req := &efloclient.DescribeVscRequest{
		VscId: &vscID,
	}
	resp, err := m.client.DescribeVsc(req)
	if err != nil {
		return nil, fmt.Errorf("eflo:DescribeVsc failed: %w", err)
	}
	klog.InfoS("eflo:DescribeVsc succeeded", "vscId", vscID, "response", resp.Body)
	return &Vsc{
		NodeID: tea.StringValue(resp.Body.NodeId),
		VscID:  tea.StringValue(resp.Body.VscId),
		Type:   tea.StringValue(resp.Body.VscType),
		Status: tea.StringValue(resp.Body.Status),
	}, nil
}

type vscWithErr struct {
	*Vsc
	err error
}

// PrimaryVscManagerWithCache is a VSC manager with caching capabilities
type PrimaryVscManagerWithCache struct {
	VscManager
	retryTimes int

	cond *sync.Cond
	// Instance ID to VSC
	cache map[string]vscWithErr
	// To create primary vsc for node
	queue workqueue.TypedRateLimitingInterface[string]
}

const (
	defaultVscManagerRetryTimes  = 3
	defaultVscManagerWorkerCount = 3
)

// NewPrimaryVscManagerWithCache creates a new PrimaryVscManagerWithCache
func NewPrimaryVscManagerWithCache(efloClient *efloclient.Client) *PrimaryVscManagerWithCache {
	m := &PrimaryVscManagerWithCache{
		VscManager: NewVscManager(efloClient),
		retryTimes: defaultVscManagerRetryTimes,
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
	instanceID, quit := m.queue.Get()
	if quit {
		return false
	}
	defer m.queue.Done(instanceID)

	newVsc, err := m.getOrCreatePrimaryFor(instanceID)

	m.cond.L.Lock()
	m.cache[instanceID] = vscWithErr{newVsc, err}
	m.cond.L.Unlock()

	if err == nil {
		m.queue.Forget(instanceID)
		m.cond.Broadcast()
	} else {
		sdkErr := &tea.SDKError{}
		if errors.As(err, &sdkErr) || m.queue.NumRequeues(instanceID) > m.retryTimes {
			klog.ErrorS(err, "Failed to ensure VSC", "instance", instanceID)
			m.queue.Forget(instanceID)
			m.cond.Broadcast()
		} else {
			klog.InfoS("Retrying to ensure VSC", "instance", instanceID, "error", err)
			m.queue.AddRateLimited(instanceID)
		}
	}
	return true
}

func (m *PrimaryVscManagerWithCache) getOrCreatePrimaryFor(instanceID string) (*Vsc, error) {
	var err error
	// try to get existing vsc
	vsc, err := m.VscManager.GetPrimaryVscOf(instanceID)
	if err != nil {
		return nil, err
	}
	// primary vsc of the instance not found, create it
	var vscID string
	if vsc == nil {
		vscID, err = m.CreatePrimaryVscFor(instanceID)
		if err != nil {
			return nil, err
		}
		vsc, err = m.GetVsc(vscID)
		if err != nil {
			return nil, err
		}
	}
	if vsc == nil {
		return nil, fmt.Errorf("vsc %s not found after creation", vscID)
	}
	// check vsc status
	if vsc.Status != VscStatusNormal {
		return vsc, fmt.Errorf("unexpected vsc status: %s", vsc.Status)
	}
	return vsc, nil
}

// EnsurePrimaryVsc ensures a primary VSC exists for the given instance
func (m *PrimaryVscManagerWithCache) EnsurePrimaryVsc(ctx context.Context, instanceID string, refresh bool) (string, error) {
	m.cond.L.Lock()
	defer m.cond.L.Unlock()

	if !refresh {
		vsc, exists := m.cache[instanceID]
		if exists && vsc.err == nil {
			return vsc.VscID, nil
		}
	}

	delete(m.cache, instanceID)
	m.queue.Add(instanceID)
	for {
		vsc, exists := m.cache[instanceID]
		if exists {
			var vscID string
			if vsc.Vsc != nil {
				vscID = vsc.VscID
			}
			return vscID, vsc.err
		}
		select {
		case <-ctx.Done():
			return "", ctx.Err()
		default:
			m.cond.Wait()
		}
	}
}

// GetPrimaryVscOf gets the primary VSC of the given instance with cache support
func (m *PrimaryVscManagerWithCache) GetPrimaryVscOf(instanceID string) (*Vsc, error) {
	m.cond.L.Lock()
	cachedVsc, exists := m.cache[instanceID]
	if exists && cachedVsc.Vsc != nil {
		m.cond.L.Unlock()
		return cachedVsc.Vsc, nil
	}
	m.cond.L.Unlock()

	vsc, err := m.VscManager.GetPrimaryVscOf(instanceID)
	if err != nil {
		return nil, err
	}

	// update the cache
	if vsc != nil {
		m.cond.L.Lock()
		clonedVsc := new(Vsc)
		*clonedVsc = *vsc
		m.cache[instanceID] = vscWithErr{clonedVsc, nil}
		m.cond.L.Unlock()
	}

	return vsc, nil
}

const (
	// CPFSVscStatusAttaching is the status when attaching
	CPFSVscStatusAttaching = "Attaching"
	// CPFSVscStatusAttached is the status when attached
	CPFSVscStatusAttached = "Attached"
	// CPFSVscStatusDetaching is the status when detaching
	CPFSVscStatusDetaching = "Detaching"
	// CPFSVscStatusDetached is the status when detached
	CPFSVscStatusDetached = "Detached"
	// CPFSVscStatusFailed is the status when failed
	CPFSVscStatusFailed = "Failed"

	// VscAttachNotSupported is the error code when VSC attach is not supported
	VscAttachNotSupported = "AttachVscTarget.VscAttachNotSupported"
)

const (
	defaultPollInterval  = time.Second * 2
	defaultADWaitTimeout = time.Second * 10
)

func newAttachNotSupportedError(err error, volumeID, nodeID string) *AttachNotSupportedError {
	return &AttachNotSupportedError{
		message:  err.Error(),
		volumeID: volumeID,
		vscID:    nodeID,
	}
}

// NewAttachNotSupportedError creates a new AttachNotSupportedError
func NewAttachNotSupportedError(err error, volumeID, nodeID string) *AttachNotSupportedError {
	return newAttachNotSupportedError(err, volumeID, nodeID)
}

// AttachNotSupportedError represents an error when VSC attach is not supported
type AttachNotSupportedError struct {
	message  string
	volumeID string
	vscID    string
}

func (e *AttachNotSupportedError) Error() string {
	return "volumeID: " + e.volumeID + "vscId: " + e.vscID + e.message
}

// IsAttachNotSupportedError checks if the error is an AttachNotSupportedError
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

// CPFSVscAttachInfo represents CPFS VSC attach information
type CPFSVscAttachInfo = nasclient.DescribeFilesystemsVscAttachInfoResponseBodyVscAttachInfoVscAttachInfo

// CPFSVscAttachInfoCond is a function type for checking CPFS VSC attach conditions
type CPFSVscAttachInfoCond func(*CPFSVscAttachInfo) (done bool, err error)

// CPFSAttachDetacher is the interface for attaching/detaching CPFS
type CPFSAttachDetacher interface {
	Attach(ctx context.Context, fsID, vscID string) error
	Detach(ctx context.Context, fsID, vscID string) error
}

// NewCPFSAttachDetacher creates a new CPFS attach detacher
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

func (ad *cpfsAttachDetacher) Attach(ctx context.Context, fsID, vscID string) error {
	attachInfo, err := ad.describe(fsID, vscID)
	if err != nil {
		return err
	}
	if attachInfo != nil {
		klog.InfoS("Already attached", "filesystem", fsID, "vscId", vscID)
		switch tea.StringValue(attachInfo.Status) {
		case CPFSVscStatusAttaching:
		case CPFSVscStatusAttached:
			return nil
		default:
			return fmt.Errorf("unexpected attachinfo status: %v", tea.StringValue(attachInfo.Status))
		}
	} else {
		if err := ad.attach(fsID, vscID); err != nil {
			if strings.Contains(err.Error(), VscAttachNotSupported) {
				return newAttachNotSupportedError(err, fsID, vscID)
			}
			return err
		}
	}
	return ad.waitFor(ctx, fsID, vscID, func(i *CPFSVscAttachInfo) (bool, error) {
		if i == nil {
			return false, fmt.Errorf("filesystem %s not attached to %s", fsID, vscID)
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

func (ad *cpfsAttachDetacher) Detach(ctx context.Context, fsID, vscID string) error {
	if err := ad.detach(fsID, vscID); err != nil {
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
	return ad.waitFor(ctx, fsID, vscID, func(i *CPFSVscAttachInfo) (bool, error) {
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

func (ad *cpfsAttachDetacher) waitFor(ctx context.Context, fsID, vscID string, cond CPFSVscAttachInfoCond, cause string) error {
	deadline := ad.clk.NewTimer(ad.waitTimeout)
	ticker := ad.clk.NewTicker(ad.pollInterval)
	defer ticker.Stop()
	for {
		attachInfo, err := ad.describe(fsID, vscID)
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

func (ad *cpfsAttachDetacher) attach(fsID, vscID string) error {
	req := &nasclient.AttachVscToFilesystemsRequest{
		ResourceIds: []*nasclient.AttachVscToFilesystemsRequestResourceIds{
			{
				FileSystemId: &fsID,
				VscId:        &vscID,
			},
		},
	}
	resp, err := ad.client.AttachVscToFilesystems(req)
	if err != nil {
		return fmt.Errorf("nas:AttachVscToFilesystems failed: %w", err)
	}
	klog.InfoS("nas:AttachVscToFilesystemsRequest succeeded", "filesystem", fsID, "vscId", vscID, "requestid", resp.Body.RequestId)
	return nil
}

func (ad *cpfsAttachDetacher) detach(fsID, vscID string) error {
	req := &nasclient.DetachVscFromFilesystemsRequest{
		ResourceIds: []*nasclient.DetachVscFromFilesystemsRequestResourceIds{
			{
				FileSystemId: &fsID,
				VscId:        &vscID,
			},
		},
	}
	resp, err := ad.client.DetachVscFromFilesystems(req)
	if err != nil {
		return fmt.Errorf("nas:DetachVscFromFilesystems failed: %w", err)
	}
	klog.InfoS("nas:DetachVscFromFilesystems succeeded", "filesystem", fsID, "vscId", vscID, "requestid", resp.Body.RequestId)
	return nil
}

func (ad *cpfsAttachDetacher) describe(fsID, vscID string) (*CPFSVscAttachInfo, error) {
	req := &nasclient.DescribeFilesystemsVscAttachInfoRequest{
		ResourceIds: []*nasclient.DescribeFilesystemsVscAttachInfoRequestResourceIds{
			{
				FileSystemId: &fsID,
				VscId:        &vscID,
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
