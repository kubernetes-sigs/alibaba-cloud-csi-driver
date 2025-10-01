// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRecycleBinAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileSystemId(v string) *UpdateRecycleBinAttributeRequest
	GetFileSystemId() *string
	SetReservedDays(v int64) *UpdateRecycleBinAttributeRequest
	GetReservedDays() *int64
}

type UpdateRecycleBinAttributeRequest struct {
	// The ID of the file system.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1ca404****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The retention period of the files in the recycle bin. Unit: days.
	//
	// Valid values: 1 to 180.
	//
	// Default value: 3.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	ReservedDays *int64 `json:"ReservedDays,omitempty" xml:"ReservedDays,omitempty"`
}

func (s UpdateRecycleBinAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRecycleBinAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateRecycleBinAttributeRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *UpdateRecycleBinAttributeRequest) GetReservedDays() *int64 {
	return s.ReservedDays
}

func (s *UpdateRecycleBinAttributeRequest) SetFileSystemId(v string) *UpdateRecycleBinAttributeRequest {
	s.FileSystemId = &v
	return s
}

func (s *UpdateRecycleBinAttributeRequest) SetReservedDays(v int64) *UpdateRecycleBinAttributeRequest {
	s.ReservedDays = &v
	return s
}

func (s *UpdateRecycleBinAttributeRequest) Validate() error {
	return dara.Validate(s)
}
