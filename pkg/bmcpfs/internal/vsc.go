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
	VscTypePrimary = "primary"

	VscStatusCreating = "Creating"
	VscStatusNormal   = "Normal"
	VscStatusDeleting = "Deleting"
)

type Vsc struct {
	NodeID string
	VscID  string
	Type   string
	Status string
}

type VscManager interface {
	CreatePrimaryVscFor(instanceId string) (string, error)
	GetPrimaryVscOf(instanceId string) (*Vsc, error)
	GetVsc(vscId string) (*Vsc, error)
}

func NewVscManager(client *efloclient.Client) VscManager {
	return &LingjunVscManager{client: client}
}

type LingjunVscManager struct {
	client *efloclient.Client
}

func (m *LingjunVscManager) CreatePrimaryVscFor(instanceId string) (string, error) {
	req := &efloclient.CreateVscRequest{
		NodeId:  &instanceId,
		VscType: tea.String(VscTypePrimary),
	}
	resp, err := m.client.CreateVsc(req)
	if err != nil {
		return "", fmt.Errorf("eflo:CreateVsc failed: %w", err)
	}
	klog.InfoS("eflo:CreateVsc succeeded", "instanceId", instanceId, "response", resp.Body)
	if tea.StringValue(resp.Body.VscId) == "" {
		return "", errors.New("unexpected response of CreateVsc")
	}
	return tea.StringValue(resp.Body.VscId), nil
}

func (m *LingjunVscManager) GetPrimaryVscOf(instanceId string) (*Vsc, error) {
	req := &efloclient.ListVscsRequest{
		NodeIds:    []*string{&instanceId},
		MaxResults: tea.Int32(100),
	}
	resp, err := m.client.ListVscs(req)
	if err != nil {
		return nil, fmt.Errorf("eflo:ListVscs failed: %w", err)
	}
	klog.V(4).InfoS("eflo:ListVscs succeeded", "instanceId", instanceId, "response", resp.Body)
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

func (m *LingjunVscManager) GetVsc(vscId string) (*Vsc, error) {
	req := &efloclient.DescribeVscRequest{
		VscId: &vscId,
	}
	resp, err := m.client.DescribeVsc(req)
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

type vscWithErr struct {
	*Vsc
	err error
}

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
	instanceId, quit := m.queue.Get()
	if quit {
		return false
	}
	defer m.queue.Done(instanceId)

	newVsc, err := m.getOrCreatePrimaryFor(instanceId)

	m.cond.L.Lock()
	m.cache[instanceId] = vscWithErr{newVsc, err}
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
		vsc, err = m.GetVsc(vscId)
		if err != nil {
			return nil, err
		}
	}
	if vsc == nil {
		return nil, fmt.Errorf("vsc %s not found after creation", vscId)
	}
	// check vsc status
	if vsc.Status != VscStatusNormal {
		return vsc, fmt.Errorf("unexpected vsc status: %s", vsc.Status)
	}
	return vsc, nil
}

func (m *PrimaryVscManagerWithCache) EnsurePrimaryVsc(ctx context.Context, instanceId string, refresh bool) (string, error) {
	m.cond.L.Lock()
	defer m.cond.L.Unlock()

	if !refresh {
		vsc, exists := m.cache[instanceId]
		if exists && vsc.err == nil {
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
	if exists && cachedVsc.Vsc != nil {
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
		m.cache[instanceId] = vscWithErr{clonedVsc, nil}
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
