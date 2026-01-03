// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyElasticityAssuranceAutoRenewAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyElasticityAssuranceAutoRenewAttributeResponseBody
	GetRequestId() *string
}

type ModifyElasticityAssuranceAutoRenewAttributeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 2A4EA075-CB5B-41B7-B0EB-70D339F64DE7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyElasticityAssuranceAutoRenewAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyElasticityAssuranceAutoRenewAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyElasticityAssuranceAutoRenewAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyElasticityAssuranceAutoRenewAttributeResponseBody) SetRequestId(v string) *ModifyElasticityAssuranceAutoRenewAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyElasticityAssuranceAutoRenewAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
