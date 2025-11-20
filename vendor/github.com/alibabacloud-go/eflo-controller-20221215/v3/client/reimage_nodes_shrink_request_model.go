// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReimageNodesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ReimageNodesShrinkRequest
	GetClusterId() *string
	SetIgnoreFailedNodeTasks(v bool) *ReimageNodesShrinkRequest
	GetIgnoreFailedNodeTasks() *bool
	SetNodesShrink(v string) *ReimageNodesShrinkRequest
	GetNodesShrink() *string
	SetUserData(v string) *ReimageNodesShrinkRequest
	GetUserData() *string
}

type ReimageNodesShrinkRequest struct {
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
	NodesShrink *string `json:"Nodes,omitempty" xml:"Nodes,omitempty"`
	// The user data.
	//
	// example:
	//
	// #!/bin/sh
	//
	// echo "Hello World. The time is now $(date -R)!" | tee /root/userdata_test.txt
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s ReimageNodesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ReimageNodesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ReimageNodesShrinkRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ReimageNodesShrinkRequest) GetIgnoreFailedNodeTasks() *bool {
	return s.IgnoreFailedNodeTasks
}

func (s *ReimageNodesShrinkRequest) GetNodesShrink() *string {
	return s.NodesShrink
}

func (s *ReimageNodesShrinkRequest) GetUserData() *string {
	return s.UserData
}

func (s *ReimageNodesShrinkRequest) SetClusterId(v string) *ReimageNodesShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *ReimageNodesShrinkRequest) SetIgnoreFailedNodeTasks(v bool) *ReimageNodesShrinkRequest {
	s.IgnoreFailedNodeTasks = &v
	return s
}

func (s *ReimageNodesShrinkRequest) SetNodesShrink(v string) *ReimageNodesShrinkRequest {
	s.NodesShrink = &v
	return s
}

func (s *ReimageNodesShrinkRequest) SetUserData(v string) *ReimageNodesShrinkRequest {
	s.UserData = &v
	return s
}

func (s *ReimageNodesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
