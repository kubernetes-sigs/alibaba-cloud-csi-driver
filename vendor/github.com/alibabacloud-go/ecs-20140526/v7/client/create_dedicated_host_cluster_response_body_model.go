// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDedicatedHostClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDedicatedHostClusterId(v string) *CreateDedicatedHostClusterResponseBody
	GetDedicatedHostClusterId() *string
	SetRequestId(v string) *CreateDedicatedHostClusterResponseBody
	GetRequestId() *string
}

type CreateDedicatedHostClusterResponseBody struct {
	// The ID of the host group.
	//
	// example:
	//
	// dc-bp12wlf6bw0vz9v2****
	DedicatedHostClusterId *string `json:"DedicatedHostClusterId,omitempty" xml:"DedicatedHostClusterId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E2A664A6-2933-4C64-88AE-5033D003****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDedicatedHostClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDedicatedHostClusterResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDedicatedHostClusterResponseBody) GetDedicatedHostClusterId() *string {
	return s.DedicatedHostClusterId
}

func (s *CreateDedicatedHostClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDedicatedHostClusterResponseBody) SetDedicatedHostClusterId(v string) *CreateDedicatedHostClusterResponseBody {
	s.DedicatedHostClusterId = &v
	return s
}

func (s *CreateDedicatedHostClusterResponseBody) SetRequestId(v string) *CreateDedicatedHostClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDedicatedHostClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
