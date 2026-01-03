// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyElasticityAssuranceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyElasticityAssuranceResponseBody
	GetRequestId() *string
}

type ModifyElasticityAssuranceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 8455DD10-84F8-43C9-8365-5F448EB169B6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyElasticityAssuranceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyElasticityAssuranceResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyElasticityAssuranceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyElasticityAssuranceResponseBody) SetRequestId(v string) *ModifyElasticityAssuranceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyElasticityAssuranceResponseBody) Validate() error {
	return dara.Validate(s)
}
