// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceMetadataOptionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyInstanceMetadataOptionsResponseBody
	GetRequestId() *string
}

type ModifyInstanceMetadataOptionsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyInstanceMetadataOptionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceMetadataOptionsResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceMetadataOptionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyInstanceMetadataOptionsResponseBody) SetRequestId(v string) *ModifyInstanceMetadataOptionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyInstanceMetadataOptionsResponseBody) Validate() error {
	return dara.Validate(s)
}
