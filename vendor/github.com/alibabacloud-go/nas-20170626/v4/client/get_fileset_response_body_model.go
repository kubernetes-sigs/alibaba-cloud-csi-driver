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
	// The response parameters.
	Data *GetFilesetResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
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
	// The time when the fileset was created.
	//
	// Return format: `yyyy-MM-dd HH:mm:ss`
	//
	// example:
	//
	// 2025-11-21 12:49:25
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Specifies whether the fileset is protected from being released through the console or the [DeleteFileset](https://help.aliyun.com/document_detail/2402263.html) operation.
	//
	// 	- true: Enables release protection. The fileset cannot be released.
	//
	// 	- false (default): Disables release protection. The fileset can be released.
	//
	// >  This parameter can protect filesets only against manual releases, but not against automatic releases.
	//
	// example:
	//
	// false
	DeletionProtection *bool `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	// The description of the fileset.
	//
	// example:
	//
	// Description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The usage of the file quantity.
	//
	// >  Only CPFS for LINGJUN V2.7.0 and later support this parameter.
	//
	// example:
	//
	// 1024
	FileCountUsage *int64 `json:"FileCountUsage,omitempty" xml:"FileCountUsage,omitempty"`
	// The ID of the file system.
	//
	// 	- The IDs of CPFS file systems must start with `cpfs-`. Example: cpfs-125487\\*\\*\\*\\*.
	//
	// 	- The IDs of CPFS for Lingjun file systems must start with `bmcpfs-`. Example: bmcpfs-0015\\*\\*\\*\\*.
	//
	// example:
	//
	// cpfs-0244729a8ef8****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The directory of the fileset in the CPFS file system.
	//
	// example:
	//
	// pathtoroot/fset/
	FileSystemPath *string `json:"FileSystemPath,omitempty" xml:"FileSystemPath,omitempty"`
	// The fileset ID.
	//
	// >  This parameter is required for CPFS file systems.
	//
	// example:
	//
	// fset-03250e8fe78d****
	FsetId *string `json:"FsetId,omitempty" xml:"FsetId,omitempty"`
	// The quota information.
	//
	// >  Only CPFS for Lingjun V2.7.0 and later support this parameter.
	Quota *GetFilesetResponseBodyDataQuota `json:"Quota,omitempty" xml:"Quota,omitempty" type:"Struct"`
	// The capacity usage. Unit: bytes.
	//
	// >  Only CPFS for Lingjun V2.7.0 and later support this parameter.
	//
	// example:
	//
	// 1024
	SpaceUsage *int64 `json:"SpaceUsage,omitempty" xml:"SpaceUsage,omitempty"`
	// The fileset status. Valid values:
	//
	// 	- CREATING: The fileset is being created.
	//
	// 	- CREATED: The fileset has been created and is running properly.
	//
	// 	- RELEASING: The fileset is being released.
	//
	// 	- RELEASED: The fileset has been deleted.
	//
	// example:
	//
	// CREATED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the fileset was last updated.
	//
	// Return format: `yyyy-MM-dd HH:mm:ss`
	//
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
	// The file quantity quota. Valid values:
	//
	// 	- Minimum value: 10,000.
	//
	// 	- Maximum value: 10,000,000,000.
	//
	// example:
	//
	// 10001
	FileCountLimit *int64 `json:"FileCountLimit,omitempty" xml:"FileCountLimit,omitempty"`
	// The total quota capacity limit. Unit: bytes.
	//
	// Valid values:
	//
	// 	- Minimum value: 10,737,418,240 (10 GiB).
	//
	// 	- Step size: 1,073,741,824 (1 GiB).
	//
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
