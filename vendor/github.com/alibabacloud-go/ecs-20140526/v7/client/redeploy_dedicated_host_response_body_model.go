// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRedeployDedicatedHostResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RedeployDedicatedHostResponseBody
	GetRequestId() *string
}

type RedeployDedicatedHostResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// FCED4B7A-53D5-4C04-ABE3-26D4F3890D57
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RedeployDedicatedHostResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RedeployDedicatedHostResponseBody) GoString() string {
	return s.String()
}

func (s *RedeployDedicatedHostResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RedeployDedicatedHostResponseBody) SetRequestId(v string) *RedeployDedicatedHostResponseBody {
	s.RequestId = &v
	return s
}

func (s *RedeployDedicatedHostResponseBody) Validate() error {
	return dara.Validate(s)
}
