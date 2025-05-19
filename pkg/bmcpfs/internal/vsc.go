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
	"time"

	efloclient "github.com/alibabacloud-go/eflo-controller-20221215/v2/client"
	nasclient "github.com/alibabacloud-go/nas-20170626/v3/client"
	"github.com/alibabacloud-go/tea/tea"
	"k8s.io/klog/v2"
	"k8s.io/utils/clock"
)

const (
	VscTypePrimary = "primary"
)

const (
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
		// klog.ErrorS(err, "eflo:CreateVsc failed", "instanceId", instanceId)
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
		// klog.ErrorS(err, "eflo:ListVscs failed", "instanceId", instanceId)
		return nil, fmt.Errorf("eflo:ListVscs failed: %w", err)
	}
	klog.V(2).InfoS("eflo:ListVscs succeeded", "instanceId", instanceId, "response", resp.Body)
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

const (
	CPFSVscStatusAttaching = "Attaching"
	CPFSVscStatusAttached  = "Attached"
	CPFSVscStatusDetaching = "Detaching"
	CPFSVscStatusDetached  = "Detached"
	CPFSVscStatusFailed    = "Failed"
)

const (
	defaultPollInterval = time.Second * 2
)

type CPFSVscAttachInfo = nasclient.DescribeFilesystemsVscAttachInfoResponseBodyVscAttachInfoVscAttachInfo

type CPFSVscAttachInfoWaitCond func(CPFSVscAttachInfo) (done, failed bool)

type CPFSAttachDetacher interface {
	Attach(ctx context.Context, fsId, vscId string) error
	Detach(ctx context.Context, fsId, vscId string) error
}

func NewCPFSAttachDetacher(client *nasclient.Client) CPFSAttachDetacher {
	return &cpfsAttachDetacher{
		client:       client,
		pollInterval: defaultPollInterval,
		clk:          clock.RealClock{},
	}
}

type cpfsAttachDetacher struct {
	client       *nasclient.Client
	pollInterval time.Duration
	clk          clock.WithTicker
}

func (ad *cpfsAttachDetacher) Attach(ctx context.Context, fsId, vscId string) error {
	if err := ad.attach(fsId, vscId); err != nil {
		return err
	}
	return ad.waitFor(ctx, fsId, vscId, func(i CPFSVscAttachInfo) (bool, bool) {
		switch tea.StringValue(i.Status) {
		case CPFSVscStatusAttaching:
			return false, false
		case CPFSVscStatusAttached:
			return true, false
		default:
			return false, true
		}
	})
}

func (ad *cpfsAttachDetacher) Detach(ctx context.Context, fsId, vscId string) error {
	// TODO: ignore error if fsId not found
	if err := ad.detach(fsId, vscId); err != nil {
		return err
	}
	return ad.waitFor(ctx, fsId, vscId, func(i CPFSVscAttachInfo) (bool, bool) {
		switch tea.StringValue(i.Status) {
		case CPFSVscStatusDetaching:
			return false, false
		case CPFSVscStatusDetached:
			return true, false
		default:
			return false, true
		}
	})
}

func (ad *cpfsAttachDetacher) waitFor(ctx context.Context, fsId, vscId string, cond CPFSVscAttachInfoWaitCond) error {
	ticker := ad.clk.NewTicker(ad.pollInterval)
	defer ticker.Stop()
	for {
		attachInfo, err := ad.describe(fsId, vscId)
		if err != nil {
			return err
		}
		if attachInfo == nil {
			return fmt.Errorf("filesystem %s not attached to %s", fsId, vscId)
		}
		done, failed := cond(*attachInfo)
		if failed {
			return fmt.Errorf("unexpected CPFS VSC AttachInfo status: %s", tea.StringValue(attachInfo.Status))
		}
		if done {
			return nil
		}
		select {
		case <-ticker.C():
		case <-ctx.Done():
			return fmt.Errorf("wait attach status: %w", ctx.Err())
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
	klog.InfoS("nas:AttachVscToFilesystemsRequest succeeded", "filesystem", fsId, "vscId", vscId, "response", resp.Body)
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
	klog.InfoS("nas:DetachVscFromFilesystems succeeded", "filesystem", fsId, "vscId", vscId, "response", resp.Body)
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
	if resp.Body.VscAttachInfo == nil || len(resp.Body.VscAttachInfo.VscAttachInfo) == 0 {
		return nil, nil
	}
	return resp.Body.VscAttachInfo.VscAttachInfo[0], nil
}
