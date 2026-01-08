// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFilesetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetFilesetResponseBodyData) *GetFilesetResponseBody
	GetData() *GetFilesetResponseBodyData
	SetRequestId(v string) *GetFilesetResponseBody
	GetRequestId() *string
}

type GetFilesetResponseBody struct {
	Data *GetFilesetResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// BC7C825C-5F65-4B56-BEF6-98C56C7C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetFilesetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetFilesetResponseBody) GoString() string {
	return s.String()
}

func (s *GetFilesetResponseBody) GetData() *GetFilesetResponseBodyData {
	return s.Data
}

func (s *GetFilesetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetFilesetResponseBody) SetData(v *GetFilesetResponseBodyData) *GetFilesetResponseBody {
	s.Data = v
	return s
}

func (s *GetFilesetResponseBody) SetRequestId(v string) *GetFilesetResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFilesetResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetFilesetResponseBodyData struct {
	// example:
	//
	// 2025-11-21 12:49:25
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// false
	DeletionProtection *bool `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	// example:
	//
	// Description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 1024
	FileCountUsage *int64 `json:"FileCountUsage,omitempty" xml:"FileCountUsage,omitempty"`
	// example:
	//
	// cpfs-0244729a8ef8****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// example:
	//
	// pathtoroot/fset/
	FileSystemPath *string `json:"FileSystemPath,omitempty" xml:"FileSystemPath,omitempty"`
	// example:
	//
	// fset-03250e8fe78d****
	FsetId *string                          `json:"FsetId,omitempty" xml:"FsetId,omitempty"`
	Quota  *GetFilesetResponseBodyDataQuota `json:"Quota,omitempty" xml:"Quota,omitempty" type:"Struct"`
	// example:
	//
	// 1024
	SpaceUsage *int64 `json:"SpaceUsage,omitempty" xml:"SpaceUsage,omitempty"`
	// example:
	//
	// CREATED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 2025-11-22 12:49:25
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetFilesetResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetFilesetResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetFilesetResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetFilesetResponseBodyData) GetDeletionProtection() *bool {
	return s.DeletionProtection
}

func (s *GetFilesetResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetFilesetResponseBodyData) GetFileCountUsage() *int64 {
	return s.FileCountUsage
}

func (s *GetFilesetResponseBodyData) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *GetFilesetResponseBodyData) GetFileSystemPath() *string {
	return s.FileSystemPath
}

func (s *GetFilesetResponseBodyData) GetFsetId() *string {
	return s.FsetId
}

func (s *GetFilesetResponseBodyData) GetQuota() *GetFilesetResponseBodyDataQuota {
	return s.Quota
}

func (s *GetFilesetResponseBodyData) GetSpaceUsage() *int64 {
	return s.SpaceUsage
}

func (s *GetFilesetResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetFilesetResponseBodyData) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetFilesetResponseBodyData) SetCreateTime(v string) *GetFilesetResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetFilesetResponseBodyData) SetDeletionProtection(v bool) *GetFilesetResponseBodyData {
	s.DeletionProtection = &v
	return s
}

func (s *GetFilesetResponseBodyData) SetDescription(v string) *GetFilesetResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetFilesetResponseBodyData) SetFileCountUsage(v int64) *GetFilesetResponseBodyData {
	s.FileCountUsage = &v
	return s
}

func (s *GetFilesetResponseBodyData) SetFileSystemId(v string) *GetFilesetResponseBodyData {
	s.FileSystemId = &v
	return s
}

func (s *GetFilesetResponseBodyData) SetFileSystemPath(v string) *GetFilesetResponseBodyData {
	s.FileSystemPath = &v
	return s
}

func (s *GetFilesetResponseBodyData) SetFsetId(v string) *GetFilesetResponseBodyData {
	s.FsetId = &v
	return s
}

func (s *GetFilesetResponseBodyData) SetQuota(v *GetFilesetResponseBodyDataQuota) *GetFilesetResponseBodyData {
	s.Quota = v
	return s
}

func (s *GetFilesetResponseBodyData) SetSpaceUsage(v int64) *GetFilesetResponseBodyData {
	s.SpaceUsage = &v
	return s
}

func (s *GetFilesetResponseBodyData) SetStatus(v string) *GetFilesetResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetFilesetResponseBodyData) SetUpdateTime(v string) *GetFilesetResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *GetFilesetResponseBodyData) Validate() error {
	if s.Quota != nil {
		if err := s.Quota.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetFilesetResponseBodyDataQuota struct {
	// example:
	//
	// 10001
	FileCountLimit *int64 `json:"FileCountLimit,omitempty" xml:"FileCountLimit,omitempty"`
	// example:
	//
	// 10,737,418,240
	SizeLimit *int64 `json:"SizeLimit,omitempty" xml:"SizeLimit,omitempty"`
}

func (s GetFilesetResponseBodyDataQuota) String() string {
	return dara.Prettify(s)
}

func (s GetFilesetResponseBodyDataQuota) GoString() string {
	return s.String()
}

func (s *GetFilesetResponseBodyDataQuota) GetFileCountLimit() *int64 {
	return s.FileCountLimit
}

func (s *GetFilesetResponseBodyDataQuota) GetSizeLimit() *int64 {
	return s.SizeLimit
}

func (s *GetFilesetResponseBodyDataQuota) SetFileCountLimit(v int64) *GetFilesetResponseBodyDataQuota {
	s.FileCountLimit = &v
	return s
}

func (s *GetFilesetResponseBodyDataQuota) SetSizeLimit(v int64) *GetFilesetResponseBodyDataQuota {
	s.SizeLimit = &v
	return s
}

func (s *GetFilesetResponseBodyDataQuota) Validate() error {
	return dara.Validate(s)
}
