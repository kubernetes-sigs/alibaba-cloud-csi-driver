// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProtocolMountTargetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetExportId(v string) *CreateProtocolMountTargetResponseBody
	GetExportId() *string
	SetRequestId(v string) *CreateProtocolMountTargetResponseBody
	GetRequestId() *string
}

type CreateProtocolMountTargetResponseBody struct {
	// The ID of the export directory for the protocol service.
	//
	// example:
	//
	// exp-123****
	ExportId *string `json:"ExportId,omitempty" xml:"ExportId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateProtocolMountTargetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateProtocolMountTargetResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProtocolMountTargetResponseBody) GetExportId() *string {
	return s.ExportId
}

func (s *CreateProtocolMountTargetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateProtocolMountTargetResponseBody) SetExportId(v string) *CreateProtocolMountTargetResponseBody {
	s.ExportId = &v
	return s
}

func (s *CreateProtocolMountTargetResponseBody) SetRequestId(v string) *CreateProtocolMountTargetResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateProtocolMountTargetResponseBody) Validate() error {
	return dara.Validate(s)
}
