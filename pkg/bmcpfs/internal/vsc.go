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
	"errors"
	"fmt"

	efloclient "github.com/alibabacloud-go/eflo-controller-20221215/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	"k8s.io/klog/v2"
)

const (
	VscTypePrimary = "primary"
)

const (
	VscStatusCreating  = "Creating"
	VscStatusAvailable = "Available"
	VscStatusDeleting  = "Deleting"
	VscStatusInvalid   = "Invalid"
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

type CPFSAttachDetacher interface {
	Attach()
	Detach()
	Describe()
}

type cpfsAttachDetacher struct {
}
