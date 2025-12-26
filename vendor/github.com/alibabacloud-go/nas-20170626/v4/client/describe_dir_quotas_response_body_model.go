// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDirQuotasResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDirQuotaInfos(v []*DescribeDirQuotasResponseBodyDirQuotaInfos) *DescribeDirQuotasResponseBody
	GetDirQuotaInfos() []*DescribeDirQuotasResponseBodyDirQuotaInfos
	SetPageNumber(v int32) *DescribeDirQuotasResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDirQuotasResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeDirQuotasResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeDirQuotasResponseBody
	GetTotalCount() *int32
}

type DescribeDirQuotasResponseBody struct {
	// The queried directory quotas.
	DirQuotaInfos []*DescribeDirQuotasResponseBodyDirQuotaInfos `json:"DirQuotaInfos,omitempty" xml:"DirQuotaInfos,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 5BC5CB97-9F28-42FE-84A4-0CD0DF42****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of directories.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDirQuotasResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDirQuotasResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDirQuotasResponseBody) GetDirQuotaInfos() []*DescribeDirQuotasResponseBodyDirQuotaInfos {
	return s.DirQuotaInfos
}

func (s *DescribeDirQuotasResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDirQuotasResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDirQuotasResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDirQuotasResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeDirQuotasResponseBody) SetDirQuotaInfos(v []*DescribeDirQuotasResponseBodyDirQuotaInfos) *DescribeDirQuotasResponseBody {
	s.DirQuotaInfos = v
	return s
}

func (s *DescribeDirQuotasResponseBody) SetPageNumber(v int32) *DescribeDirQuotasResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDirQuotasResponseBody) SetPageSize(v int32) *DescribeDirQuotasResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDirQuotasResponseBody) SetRequestId(v string) *DescribeDirQuotasResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDirQuotasResponseBody) SetTotalCount(v int32) *DescribeDirQuotasResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDirQuotasResponseBody) Validate() error {
	if s.DirQuotaInfos != nil {
		for _, item := range s.DirQuotaInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDirQuotasResponseBodyDirQuotaInfos struct {
	// The inode number of the directory.
	//
	// example:
	//
	// 1123
	DirInode *string `json:"DirInode,omitempty" xml:"DirInode,omitempty"`
	// The absolute path of a directory.
	//
	// example:
	//
	// /data/sub1
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The status of the quota created for the directory. Valid values: Initializing and Normal. The Initializing state indicates that the quota is being created. The Normal state indicates that the quota is created.
	//
	// example:
	//
	// Normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The information about quotas for all users.
	UserQuotaInfos []*DescribeDirQuotasResponseBodyDirQuotaInfosUserQuotaInfos `json:"UserQuotaInfos,omitempty" xml:"UserQuotaInfos,omitempty" type:"Repeated"`
}

func (s DescribeDirQuotasResponseBodyDirQuotaInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeDirQuotasResponseBodyDirQuotaInfos) GoString() string {
	return s.String()
}

func (s *DescribeDirQuotasResponseBodyDirQuotaInfos) GetDirInode() *string {
	return s.DirInode
}

func (s *DescribeDirQuotasResponseBodyDirQuotaInfos) GetPath() *string {
	return s.Path
}

func (s *DescribeDirQuotasResponseBodyDirQuotaInfos) GetStatus() *string {
	return s.Status
}

func (s *DescribeDirQuotasResponseBodyDirQuotaInfos) GetUserQuotaInfos() []*DescribeDirQuotasResponseBodyDirQuotaInfosUserQuotaInfos {
	return s.UserQuotaInfos
}

func (s *DescribeDirQuotasResponseBodyDirQuotaInfos) SetDirInode(v string) *DescribeDirQuotasResponseBodyDirQuotaInfos {
	s.DirInode = &v
	return s
}

func (s *DescribeDirQuotasResponseBodyDirQuotaInfos) SetPath(v string) *DescribeDirQuotasResponseBodyDirQuotaInfos {
	s.Path = &v
	return s
}

func (s *DescribeDirQuotasResponseBodyDirQuotaInfos) SetStatus(v string) *DescribeDirQuotasResponseBodyDirQuotaInfos {
	s.Status = &v
	return s
}

func (s *DescribeDirQuotasResponseBodyDirQuotaInfos) SetUserQuotaInfos(v []*DescribeDirQuotasResponseBodyDirQuotaInfosUserQuotaInfos) *DescribeDirQuotasResponseBodyDirQuotaInfos {
	s.UserQuotaInfos = v
	return s
}

func (s *DescribeDirQuotasResponseBodyDirQuotaInfos) Validate() error {
	if s.UserQuotaInfos != nil {
		for _, item := range s.UserQuotaInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDirQuotasResponseBodyDirQuotaInfosUserQuotaInfos struct {
	// The maximum number of files that a user can create in the directory.
	//
	// example:
	//
	// 10000
	FileCountLimit *int64 `json:"FileCountLimit,omitempty" xml:"FileCountLimit,omitempty"`
	// The total number of files that a user has created in the directory.
	//
	// example:
	//
	// 5100
	FileCountReal *int64 `json:"FileCountReal,omitempty" xml:"FileCountReal,omitempty"`
	// The type of the quota. Valid values: Accounting and Enforcement.
	//
	// example:
	//
	// Accounting
	QuotaType *string `json:"QuotaType,omitempty" xml:"QuotaType,omitempty"`
	// The maximum size of files that a user can create in the directory. Unit: GiB.
	//
	// example:
	//
	// 1024
	SizeLimit *int64 `json:"SizeLimit,omitempty" xml:"SizeLimit,omitempty"`
	// The total size of files that a user has created in the directory. Unit: GiB.
	//
	// example:
	//
	// 800
	SizeReal *int64 `json:"SizeReal,omitempty" xml:"SizeReal,omitempty"`
	// The total size of files that a user has created in the directory. Unit: bytes.
	//
	// example:
	//
	// 858995833870
	SizeRealInByte *int64 `json:"SizeRealInByte,omitempty" xml:"SizeRealInByte,omitempty"`
	// The ID of the user that you specify to create a quota for the directory. The value depends on the value of the UserType parameter. Valid values: Uid and Gid.
	//
	// example:
	//
	// 500
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The type of user. Valid values: Uid, Gid, and AllUsers.
	//
	// 	- If Uid or Gid is returned, a value is returned for UserId.
	//
	// 	- If AllUsers is returned, UserId is empty.
	//
	// example:
	//
	// Uid
	UserType *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
}

func (s DescribeDirQuotasResponseBodyDirQuotaInfosUserQuotaInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeDirQuotasResponseBodyDirQuotaInfosUserQuotaInfos) GoString() string {
	return s.String()
}

func (s *DescribeDirQuotasResponseBodyDirQuotaInfosUserQuotaInfos) GetFileCountLimit() *int64 {
	return s.FileCountLimit
}

func (s *DescribeDirQuotasResponseBodyDirQuotaInfosUserQuotaInfos) GetFileCountReal() *int64 {
	return s.FileCountReal
}

func (s *DescribeDirQuotasResponseBodyDirQuotaInfosUserQuotaInfos) GetQuotaType() *string {
	return s.QuotaType
}

func (s *DescribeDirQuotasResponseBodyDirQuotaInfosUserQuotaInfos) GetSizeLimit() *int64 {
	return s.SizeLimit
}

func (s *DescribeDirQuotasResponseBodyDirQuotaInfosUserQuotaInfos) GetSizeReal() *int64 {
	return s.SizeReal
}

func (s *DescribeDirQuotasResponseBodyDirQuotaInfosUserQuotaInfos) GetSizeRealInByte() *int64 {
	return s.SizeRealInByte
}

func (s *DescribeDirQuotasResponseBodyDirQuotaInfosUserQuotaInfos) GetUserId() *string {
	return s.UserId
}

func (s *DescribeDirQuotasResponseBodyDirQuotaInfosUserQuotaInfos) GetUserType() *string {
	return s.UserType
}

func (s *DescribeDirQuotasResponseBodyDirQuotaInfosUserQuotaInfos) SetFileCountLimit(v int64) *DescribeDirQuotasResponseBodyDirQuotaInfosUserQuotaInfos {
	s.FileCountLimit = &v
	return s
}

func (s *DescribeDirQuotasResponseBodyDirQuotaInfosUserQuotaInfos) SetFileCountReal(v int64) *DescribeDirQuotasResponseBodyDirQuotaInfosUserQuotaInfos {
	s.FileCountReal = &v
	return s
}

func (s *DescribeDirQuotasResponseBodyDirQuotaInfosUserQuotaInfos) SetQuotaType(v string) *DescribeDirQuotasResponseBodyDirQuotaInfosUserQuotaInfos {
	s.QuotaType = &v
	return s
}

func (s *DescribeDirQuotasResponseBodyDirQuotaInfosUserQuotaInfos) SetSizeLimit(v int64) *DescribeDirQuotasResponseBodyDirQuotaInfosUserQuotaInfos {
	s.SizeLimit = &v
	return s
}

func (s *DescribeDirQuotasResponseBodyDirQuotaInfosUserQuotaInfos) SetSizeReal(v int64) *DescribeDirQuotasResponseBodyDirQuotaInfosUserQuotaInfos {
	s.SizeReal = &v
	return s
}

func (s *DescribeDirQuotasResponseBodyDirQuotaInfosUserQuotaInfos) SetSizeRealInByte(v int64) *DescribeDirQuotasResponseBodyDirQuotaInfosUserQuotaInfos {
	s.SizeRealInByte = &v
	return s
}

func (s *DescribeDirQuotasResponseBodyDirQuotaInfosUserQuotaInfos) SetUserId(v string) *DescribeDirQuotasResponseBodyDirQuotaInfosUserQuotaInfos {
	s.UserId = &v
	return s
}

func (s *DescribeDirQuotasResponseBodyDirQuotaInfosUserQuotaInfos) SetUserType(v string) *DescribeDirQuotasResponseBodyDirQuotaInfosUserQuotaInfos {
	s.UserType = &v
	return s
}

func (s *DescribeDirQuotasResponseBodyDirQuotaInfosUserQuotaInfos) Validate() error {
	return dara.Validate(s)
}
