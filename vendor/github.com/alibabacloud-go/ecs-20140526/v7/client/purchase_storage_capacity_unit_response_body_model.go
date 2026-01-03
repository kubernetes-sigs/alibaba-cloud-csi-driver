// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPurchaseStorageCapacityUnitResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *PurchaseStorageCapacityUnitResponseBody
	GetOrderId() *string
	SetRequestId(v string) *PurchaseStorageCapacityUnitResponseBody
	GetRequestId() *string
	SetStorageCapacityUnitIds(v *PurchaseStorageCapacityUnitResponseBodyStorageCapacityUnitIds) *PurchaseStorageCapacityUnitResponseBody
	GetStorageCapacityUnitIds() *PurchaseStorageCapacityUnitResponseBodyStorageCapacityUnitIds
}

type PurchaseStorageCapacityUnitResponseBody struct {
	// The order ID.
	//
	// example:
	//
	// 204135153880****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The IDs of the SCUs.
	StorageCapacityUnitIds *PurchaseStorageCapacityUnitResponseBodyStorageCapacityUnitIds `json:"StorageCapacityUnitIds,omitempty" xml:"StorageCapacityUnitIds,omitempty" type:"Struct"`
}

func (s PurchaseStorageCapacityUnitResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PurchaseStorageCapacityUnitResponseBody) GoString() string {
	return s.String()
}

func (s *PurchaseStorageCapacityUnitResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *PurchaseStorageCapacityUnitResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PurchaseStorageCapacityUnitResponseBody) GetStorageCapacityUnitIds() *PurchaseStorageCapacityUnitResponseBodyStorageCapacityUnitIds {
	return s.StorageCapacityUnitIds
}

func (s *PurchaseStorageCapacityUnitResponseBody) SetOrderId(v string) *PurchaseStorageCapacityUnitResponseBody {
	s.OrderId = &v
	return s
}

func (s *PurchaseStorageCapacityUnitResponseBody) SetRequestId(v string) *PurchaseStorageCapacityUnitResponseBody {
	s.RequestId = &v
	return s
}

func (s *PurchaseStorageCapacityUnitResponseBody) SetStorageCapacityUnitIds(v *PurchaseStorageCapacityUnitResponseBodyStorageCapacityUnitIds) *PurchaseStorageCapacityUnitResponseBody {
	s.StorageCapacityUnitIds = v
	return s
}

func (s *PurchaseStorageCapacityUnitResponseBody) Validate() error {
	if s.StorageCapacityUnitIds != nil {
		if err := s.StorageCapacityUnitIds.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PurchaseStorageCapacityUnitResponseBodyStorageCapacityUnitIds struct {
	StorageCapacityUnitId []*string `json:"StorageCapacityUnitId,omitempty" xml:"StorageCapacityUnitId,omitempty" type:"Repeated"`
}

func (s PurchaseStorageCapacityUnitResponseBodyStorageCapacityUnitIds) String() string {
	return dara.Prettify(s)
}

func (s PurchaseStorageCapacityUnitResponseBodyStorageCapacityUnitIds) GoString() string {
	return s.String()
}

func (s *PurchaseStorageCapacityUnitResponseBodyStorageCapacityUnitIds) GetStorageCapacityUnitId() []*string {
	return s.StorageCapacityUnitId
}

func (s *PurchaseStorageCapacityUnitResponseBodyStorageCapacityUnitIds) SetStorageCapacityUnitId(v []*string) *PurchaseStorageCapacityUnitResponseBodyStorageCapacityUnitIds {
	s.StorageCapacityUnitId = v
	return s
}

func (s *PurchaseStorageCapacityUnitResponseBodyStorageCapacityUnitIds) Validate() error {
	return dara.Validate(s)
}
