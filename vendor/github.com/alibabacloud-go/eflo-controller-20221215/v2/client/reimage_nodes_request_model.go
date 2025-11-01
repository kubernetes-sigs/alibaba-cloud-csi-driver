// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReimageNodesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ReimageNodesRequest
	GetClusterId() *string
	SetIgnoreFailedNodeTasks(v bool) *ReimageNodesRequest
	GetIgnoreFailedNodeTasks() *bool
	SetNodes(v []*ReimageNodesRequestNodes) *ReimageNodesRequest
	GetNodes() []*ReimageNodesRequestNodes
	SetUserData(v string) *ReimageNodesRequest
	GetUserData() *string
}

type ReimageNodesRequest struct {
	// The cluster ID.
	//
	// example:
	//
	// i15dfa12e8f27c44f4a006c2c8bb
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Specifies whether to allow skipping failed nodes. Default value: False.
	//
	// example:
	//
	// False
	IgnoreFailedNodeTasks *bool `json:"IgnoreFailedNodeTasks,omitempty" xml:"IgnoreFailedNodeTasks,omitempty"`
	// The nodes.
	Nodes []*ReimageNodesRequestNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	// The user data.
	//
	// example:
	//
	// #!/bin/sh
	//
	// echo "Hello World. The time is now $(date -R)!" | tee /root/userdata_test.txt
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s ReimageNodesRequest) String() string {
	return dara.Prettify(s)
}

func (s ReimageNodesRequest) GoString() string {
	return s.String()
}

func (s *ReimageNodesRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ReimageNodesRequest) GetIgnoreFailedNodeTasks() *bool {
	return s.IgnoreFailedNodeTasks
}

func (s *ReimageNodesRequest) GetNodes() []*ReimageNodesRequestNodes {
	return s.Nodes
}

func (s *ReimageNodesRequest) GetUserData() *string {
	return s.UserData
}

func (s *ReimageNodesRequest) SetClusterId(v string) *ReimageNodesRequest {
	s.ClusterId = &v
	return s
}

func (s *ReimageNodesRequest) SetIgnoreFailedNodeTasks(v bool) *ReimageNodesRequest {
	s.IgnoreFailedNodeTasks = &v
	return s
}

func (s *ReimageNodesRequest) SetNodes(v []*ReimageNodesRequestNodes) *ReimageNodesRequest {
	s.Nodes = v
	return s
}

func (s *ReimageNodesRequest) SetUserData(v string) *ReimageNodesRequest {
	s.UserData = &v
	return s
}

func (s *ReimageNodesRequest) Validate() error {
	if s.Nodes != nil {
		for _, item := range s.Nodes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ReimageNodesRequestNodes struct {
	// The hostname.
	//
	// example:
	//
	// 457db5ca-241d-11ed-9fd7-acde48001122
	Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	// The system image ID.
	//
	// example:
	//
	// m-8vbf8rpv2nn14y7oybjy
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The logon password.
	//
	// example:
	//
	// ***
	LoginPassword *string `json:"LoginPassword,omitempty" xml:"LoginPassword,omitempty"`
	// The node ID.
	//
	// example:
	//
	// e01-cn-zvp2tgykr0b
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
}

func (s ReimageNodesRequestNodes) String() string {
	return dara.Prettify(s)
}

func (s ReimageNodesRequestNodes) GoString() string {
	return s.String()
}

func (s *ReimageNodesRequestNodes) GetHostname() *string {
	return s.Hostname
}

func (s *ReimageNodesRequestNodes) GetImageId() *string {
	return s.ImageId
}

func (s *ReimageNodesRequestNodes) GetLoginPassword() *string {
	return s.LoginPassword
}

func (s *ReimageNodesRequestNodes) GetNodeId() *string {
	return s.NodeId
}

func (s *ReimageNodesRequestNodes) SetHostname(v string) *ReimageNodesRequestNodes {
	s.Hostname = &v
	return s
}

func (s *ReimageNodesRequestNodes) SetImageId(v string) *ReimageNodesRequestNodes {
	s.ImageId = &v
	return s
}

func (s *ReimageNodesRequestNodes) SetLoginPassword(v string) *ReimageNodesRequestNodes {
	s.LoginPassword = &v
	return s
}

func (s *ReimageNodesRequestNodes) SetNodeId(v string) *ReimageNodesRequestNodes {
	s.NodeId = &v
	return s
}

func (s *ReimageNodesRequestNodes) Validate() error {
	return dara.Validate(s)
}
