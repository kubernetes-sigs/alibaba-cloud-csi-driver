// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDeploymentSetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeploymentSetId(v string) *CreateDeploymentSetResponseBody
	GetDeploymentSetId() *string
	SetRequestId(v string) *CreateDeploymentSetResponseBody
	GetRequestId() *string
}

type CreateDeploymentSetResponseBody struct {
	// The ID of the deployment set.
	//
	// example:
	//
	// ds-bp1frxuzdg87zh4pzq****
	DeploymentSetId *string `json:"DeploymentSetId,omitempty" xml:"DeploymentSetId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDeploymentSetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDeploymentSetResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDeploymentSetResponseBody) GetDeploymentSetId() *string {
	return s.DeploymentSetId
}

func (s *CreateDeploymentSetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDeploymentSetResponseBody) SetDeploymentSetId(v string) *CreateDeploymentSetResponseBody {
	s.DeploymentSetId = &v
	return s
}

func (s *CreateDeploymentSetResponseBody) SetRequestId(v string) *CreateDeploymentSetResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDeploymentSetResponseBody) Validate() error {
	return dara.Validate(s)
}
