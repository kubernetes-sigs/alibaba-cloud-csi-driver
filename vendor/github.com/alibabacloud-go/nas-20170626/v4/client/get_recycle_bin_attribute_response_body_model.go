// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRecycleBinAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRecycleBinAttribute(v *GetRecycleBinAttributeResponseBodyRecycleBinAttribute) *GetRecycleBinAttributeResponseBody
	GetRecycleBinAttribute() *GetRecycleBinAttributeResponseBodyRecycleBinAttribute
	SetRequestId(v string) *GetRecycleBinAttributeResponseBody
	GetRequestId() *string
}

type GetRecycleBinAttributeResponseBody struct {
	// The description of the recycle bin.
	RecycleBinAttribute *GetRecycleBinAttributeResponseBodyRecycleBinAttribute `json:"RecycleBinAttribute,omitempty" xml:"RecycleBinAttribute,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 9E15E394-38A6-457A-A62A-D9797C9A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetRecycleBinAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRecycleBinAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *GetRecycleBinAttributeResponseBody) GetRecycleBinAttribute() *GetRecycleBinAttributeResponseBodyRecycleBinAttribute {
	return s.RecycleBinAttribute
}

func (s *GetRecycleBinAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRecycleBinAttributeResponseBody) SetRecycleBinAttribute(v *GetRecycleBinAttributeResponseBodyRecycleBinAttribute) *GetRecycleBinAttributeResponseBody {
	s.RecycleBinAttribute = v
	return s
}

func (s *GetRecycleBinAttributeResponseBody) SetRequestId(v string) *GetRecycleBinAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRecycleBinAttributeResponseBody) Validate() error {
	if s.RecycleBinAttribute != nil {
		if err := s.RecycleBinAttribute.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetRecycleBinAttributeResponseBodyRecycleBinAttribute struct {
	// The size of the archived data that is dumped to the recycle bin. Unit: bytes.
	//
	// example:
	//
	// 1611661312
	ArchiveSize *int64 `json:"ArchiveSize,omitempty" xml:"ArchiveSize,omitempty"`
	// The time at which the recycle bin was enabled.
	//
	// example:
	//
	// 2021-05-30T10:08:08Z
	EnableTime *string `json:"EnableTime,omitempty" xml:"EnableTime,omitempty"`
	// The retention period of the files in the recycle bin. Unit: days.
	//
	// If the recycle bin is disabled, 0 is returned for this parameter.
	//
	// example:
	//
	// 0
	ReservedDays *int64 `json:"ReservedDays,omitempty" xml:"ReservedDays,omitempty"`
	// The size of the Infrequent Access (IA) data that is dumped to the recycle bin. Unit: bytes.
	//
	// example:
	//
	// 100
	SecondarySize *int64 `json:"SecondarySize,omitempty" xml:"SecondarySize,omitempty"`
	// The size of the files that are dumped to the recycle bin. Unit: bytes.
	//
	// example:
	//
	// 100
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The status of the recycle bin.
	//
	// Valid values:
	//
	// 	- Enable: The recycle bin is enabled.
	//
	// 	- Disable: The recycle bin is disabled.
	//
	// example:
	//
	// Disable
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetRecycleBinAttributeResponseBodyRecycleBinAttribute) String() string {
	return dara.Prettify(s)
}

func (s GetRecycleBinAttributeResponseBodyRecycleBinAttribute) GoString() string {
	return s.String()
}

func (s *GetRecycleBinAttributeResponseBodyRecycleBinAttribute) GetArchiveSize() *int64 {
	return s.ArchiveSize
}

func (s *GetRecycleBinAttributeResponseBodyRecycleBinAttribute) GetEnableTime() *string {
	return s.EnableTime
}

func (s *GetRecycleBinAttributeResponseBodyRecycleBinAttribute) GetReservedDays() *int64 {
	return s.ReservedDays
}

func (s *GetRecycleBinAttributeResponseBodyRecycleBinAttribute) GetSecondarySize() *int64 {
	return s.SecondarySize
}

func (s *GetRecycleBinAttributeResponseBodyRecycleBinAttribute) GetSize() *int64 {
	return s.Size
}

func (s *GetRecycleBinAttributeResponseBodyRecycleBinAttribute) GetStatus() *string {
	return s.Status
}

func (s *GetRecycleBinAttributeResponseBodyRecycleBinAttribute) SetArchiveSize(v int64) *GetRecycleBinAttributeResponseBodyRecycleBinAttribute {
	s.ArchiveSize = &v
	return s
}

func (s *GetRecycleBinAttributeResponseBodyRecycleBinAttribute) SetEnableTime(v string) *GetRecycleBinAttributeResponseBodyRecycleBinAttribute {
	s.EnableTime = &v
	return s
}

func (s *GetRecycleBinAttributeResponseBodyRecycleBinAttribute) SetReservedDays(v int64) *GetRecycleBinAttributeResponseBodyRecycleBinAttribute {
	s.ReservedDays = &v
	return s
}

func (s *GetRecycleBinAttributeResponseBodyRecycleBinAttribute) SetSecondarySize(v int64) *GetRecycleBinAttributeResponseBodyRecycleBinAttribute {
	s.SecondarySize = &v
	return s
}

func (s *GetRecycleBinAttributeResponseBodyRecycleBinAttribute) SetSize(v int64) *GetRecycleBinAttributeResponseBodyRecycleBinAttribute {
	s.Size = &v
	return s
}

func (s *GetRecycleBinAttributeResponseBodyRecycleBinAttribute) SetStatus(v string) *GetRecycleBinAttributeResponseBodyRecycleBinAttribute {
	s.Status = &v
	return s
}

func (s *GetRecycleBinAttributeResponseBodyRecycleBinAttribute) Validate() error {
	return dara.Validate(s)
}
