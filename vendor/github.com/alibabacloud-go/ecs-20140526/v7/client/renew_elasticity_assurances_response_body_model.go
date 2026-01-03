// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewElasticityAssurancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *RenewElasticityAssurancesResponseBody
	GetOrderId() *string
	SetPrivatePoolOptionsIdSet(v *RenewElasticityAssurancesResponseBodyPrivatePoolOptionsIdSet) *RenewElasticityAssurancesResponseBody
	GetPrivatePoolOptionsIdSet() *RenewElasticityAssurancesResponseBodyPrivatePoolOptionsIdSet
	SetRequestId(v string) *RenewElasticityAssurancesResponseBody
	GetRequestId() *string
}

type RenewElasticityAssurancesResponseBody struct {
	// The ID of the renewal order.
	//
	// example:
	//
	// 182372800****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The IDs of the elasticity assurances.
	PrivatePoolOptionsIdSet *RenewElasticityAssurancesResponseBodyPrivatePoolOptionsIdSet `json:"PrivatePoolOptionsIdSet,omitempty" xml:"PrivatePoolOptionsIdSet,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RenewElasticityAssurancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RenewElasticityAssurancesResponseBody) GoString() string {
	return s.String()
}

func (s *RenewElasticityAssurancesResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *RenewElasticityAssurancesResponseBody) GetPrivatePoolOptionsIdSet() *RenewElasticityAssurancesResponseBodyPrivatePoolOptionsIdSet {
	return s.PrivatePoolOptionsIdSet
}

func (s *RenewElasticityAssurancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RenewElasticityAssurancesResponseBody) SetOrderId(v string) *RenewElasticityAssurancesResponseBody {
	s.OrderId = &v
	return s
}

func (s *RenewElasticityAssurancesResponseBody) SetPrivatePoolOptionsIdSet(v *RenewElasticityAssurancesResponseBodyPrivatePoolOptionsIdSet) *RenewElasticityAssurancesResponseBody {
	s.PrivatePoolOptionsIdSet = v
	return s
}

func (s *RenewElasticityAssurancesResponseBody) SetRequestId(v string) *RenewElasticityAssurancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *RenewElasticityAssurancesResponseBody) Validate() error {
	if s.PrivatePoolOptionsIdSet != nil {
		if err := s.PrivatePoolOptionsIdSet.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RenewElasticityAssurancesResponseBodyPrivatePoolOptionsIdSet struct {
	PrivatePoolOptionsId []*string `json:"PrivatePoolOptionsId,omitempty" xml:"PrivatePoolOptionsId,omitempty" type:"Repeated"`
}

func (s RenewElasticityAssurancesResponseBodyPrivatePoolOptionsIdSet) String() string {
	return dara.Prettify(s)
}

func (s RenewElasticityAssurancesResponseBodyPrivatePoolOptionsIdSet) GoString() string {
	return s.String()
}

func (s *RenewElasticityAssurancesResponseBodyPrivatePoolOptionsIdSet) GetPrivatePoolOptionsId() []*string {
	return s.PrivatePoolOptionsId
}

func (s *RenewElasticityAssurancesResponseBodyPrivatePoolOptionsIdSet) SetPrivatePoolOptionsId(v []*string) *RenewElasticityAssurancesResponseBodyPrivatePoolOptionsIdSet {
	s.PrivatePoolOptionsId = v
	return s
}

func (s *RenewElasticityAssurancesResponseBodyPrivatePoolOptionsIdSet) Validate() error {
	return dara.Validate(s)
}
