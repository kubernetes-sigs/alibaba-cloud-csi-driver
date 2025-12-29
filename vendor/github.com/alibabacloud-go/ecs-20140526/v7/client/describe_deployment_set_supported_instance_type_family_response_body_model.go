// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDeploymentSetSupportedInstanceTypeFamilyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceTypeFamilies(v string) *DescribeDeploymentSetSupportedInstanceTypeFamilyResponseBody
	GetInstanceTypeFamilies() *string
	SetRequestId(v string) *DescribeDeploymentSetSupportedInstanceTypeFamilyResponseBody
	GetRequestId() *string
}

type DescribeDeploymentSetSupportedInstanceTypeFamilyResponseBody struct {
	// The instance families that support the deployment strategy.
	//
	// example:
	//
	// ecs.i2g,ecs.i1,ecs.i2ne,ecs.i2gne
	InstanceTypeFamilies *string `json:"InstanceTypeFamilies,omitempty" xml:"InstanceTypeFamilies,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B7DB-A3DC7DE1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDeploymentSetSupportedInstanceTypeFamilyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeploymentSetSupportedInstanceTypeFamilyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDeploymentSetSupportedInstanceTypeFamilyResponseBody) GetInstanceTypeFamilies() *string {
	return s.InstanceTypeFamilies
}

func (s *DescribeDeploymentSetSupportedInstanceTypeFamilyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDeploymentSetSupportedInstanceTypeFamilyResponseBody) SetInstanceTypeFamilies(v string) *DescribeDeploymentSetSupportedInstanceTypeFamilyResponseBody {
	s.InstanceTypeFamilies = &v
	return s
}

func (s *DescribeDeploymentSetSupportedInstanceTypeFamilyResponseBody) SetRequestId(v string) *DescribeDeploymentSetSupportedInstanceTypeFamilyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDeploymentSetSupportedInstanceTypeFamilyResponseBody) Validate() error {
	return dara.Validate(s)
}
