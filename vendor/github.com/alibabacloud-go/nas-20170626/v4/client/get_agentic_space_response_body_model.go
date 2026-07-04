// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgenticSpaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAgenticSpace(v *GetAgenticSpaceResponseBodyAgenticSpace) *GetAgenticSpaceResponseBody
	GetAgenticSpace() *GetAgenticSpaceResponseBodyAgenticSpace
	SetRequestId(v string) *GetAgenticSpaceResponseBody
	GetRequestId() *string
}

type GetAgenticSpaceResponseBody struct {
	AgenticSpace *GetAgenticSpaceResponseBodyAgenticSpace `json:"AgenticSpace,omitempty" xml:"AgenticSpace,omitempty" type:"Struct"`
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0D****3E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAgenticSpaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAgenticSpaceResponseBody) GoString() string {
	return s.String()
}

func (s *GetAgenticSpaceResponseBody) GetAgenticSpace() *GetAgenticSpaceResponseBodyAgenticSpace {
	return s.AgenticSpace
}

func (s *GetAgenticSpaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAgenticSpaceResponseBody) SetAgenticSpace(v *GetAgenticSpaceResponseBodyAgenticSpace) *GetAgenticSpaceResponseBody {
	s.AgenticSpace = v
	return s
}

func (s *GetAgenticSpaceResponseBody) SetRequestId(v string) *GetAgenticSpaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAgenticSpaceResponseBody) Validate() error {
	if s.AgenticSpace != nil {
		if err := s.AgenticSpace.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAgenticSpaceResponseBodyAgenticSpace struct {
	// example:
	//
	// agentic-229oypxjgpau2****
	AgenticSpaceId *string `json:"AgenticSpaceId,omitempty" xml:"AgenticSpaceId,omitempty"`
	// example:
	//
	// cn-hangzhou-i
	Azone *string `json:"Azone,omitempty" xml:"Azone,omitempty"`
	// example:
	//
	// 2026-06-10T10:08:08Z
	CreateTimeUtc *string `json:"CreateTimeUtc,omitempty" xml:"CreateTimeUtc,omitempty"`
	// example:
	//
	// AgenticSpace Description。
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 100
	FileCountUsage *int64 `json:"FileCountUsage,omitempty" xml:"FileCountUsage,omitempty"`
	// example:
	//
	// 06229oypxjgox0u****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// example:
	//
	// /test/
	FileSystemPath *string                                       `json:"FileSystemPath,omitempty" xml:"FileSystemPath,omitempty"`
	Quota          *GetAgenticSpaceResponseBodyAgenticSpaceQuota `json:"Quota,omitempty" xml:"Quota,omitempty" type:"Struct"`
	// example:
	//
	// 1024
	SpaceUsage *int64 `json:"SpaceUsage,omitempty" xml:"SpaceUsage,omitempty"`
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 2026-06-10T10:08:08Z
	UpdateTimeUtc *string `json:"UpdateTimeUtc,omitempty" xml:"UpdateTimeUtc,omitempty"`
}

func (s GetAgenticSpaceResponseBodyAgenticSpace) String() string {
	return dara.Prettify(s)
}

func (s GetAgenticSpaceResponseBodyAgenticSpace) GoString() string {
	return s.String()
}

func (s *GetAgenticSpaceResponseBodyAgenticSpace) GetAgenticSpaceId() *string {
	return s.AgenticSpaceId
}

func (s *GetAgenticSpaceResponseBodyAgenticSpace) GetAzone() *string {
	return s.Azone
}

func (s *GetAgenticSpaceResponseBodyAgenticSpace) GetCreateTimeUtc() *string {
	return s.CreateTimeUtc
}

func (s *GetAgenticSpaceResponseBodyAgenticSpace) GetDescription() *string {
	return s.Description
}

func (s *GetAgenticSpaceResponseBodyAgenticSpace) GetFileCountUsage() *int64 {
	return s.FileCountUsage
}

func (s *GetAgenticSpaceResponseBodyAgenticSpace) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *GetAgenticSpaceResponseBodyAgenticSpace) GetFileSystemPath() *string {
	return s.FileSystemPath
}

func (s *GetAgenticSpaceResponseBodyAgenticSpace) GetQuota() *GetAgenticSpaceResponseBodyAgenticSpaceQuota {
	return s.Quota
}

func (s *GetAgenticSpaceResponseBodyAgenticSpace) GetSpaceUsage() *int64 {
	return s.SpaceUsage
}

func (s *GetAgenticSpaceResponseBodyAgenticSpace) GetStatus() *string {
	return s.Status
}

func (s *GetAgenticSpaceResponseBodyAgenticSpace) GetUpdateTimeUtc() *string {
	return s.UpdateTimeUtc
}

func (s *GetAgenticSpaceResponseBodyAgenticSpace) SetAgenticSpaceId(v string) *GetAgenticSpaceResponseBodyAgenticSpace {
	s.AgenticSpaceId = &v
	return s
}

func (s *GetAgenticSpaceResponseBodyAgenticSpace) SetAzone(v string) *GetAgenticSpaceResponseBodyAgenticSpace {
	s.Azone = &v
	return s
}

func (s *GetAgenticSpaceResponseBodyAgenticSpace) SetCreateTimeUtc(v string) *GetAgenticSpaceResponseBodyAgenticSpace {
	s.CreateTimeUtc = &v
	return s
}

func (s *GetAgenticSpaceResponseBodyAgenticSpace) SetDescription(v string) *GetAgenticSpaceResponseBodyAgenticSpace {
	s.Description = &v
	return s
}

func (s *GetAgenticSpaceResponseBodyAgenticSpace) SetFileCountUsage(v int64) *GetAgenticSpaceResponseBodyAgenticSpace {
	s.FileCountUsage = &v
	return s
}

func (s *GetAgenticSpaceResponseBodyAgenticSpace) SetFileSystemId(v string) *GetAgenticSpaceResponseBodyAgenticSpace {
	s.FileSystemId = &v
	return s
}

func (s *GetAgenticSpaceResponseBodyAgenticSpace) SetFileSystemPath(v string) *GetAgenticSpaceResponseBodyAgenticSpace {
	s.FileSystemPath = &v
	return s
}

func (s *GetAgenticSpaceResponseBodyAgenticSpace) SetQuota(v *GetAgenticSpaceResponseBodyAgenticSpaceQuota) *GetAgenticSpaceResponseBodyAgenticSpace {
	s.Quota = v
	return s
}

func (s *GetAgenticSpaceResponseBodyAgenticSpace) SetSpaceUsage(v int64) *GetAgenticSpaceResponseBodyAgenticSpace {
	s.SpaceUsage = &v
	return s
}

func (s *GetAgenticSpaceResponseBodyAgenticSpace) SetStatus(v string) *GetAgenticSpaceResponseBodyAgenticSpace {
	s.Status = &v
	return s
}

func (s *GetAgenticSpaceResponseBodyAgenticSpace) SetUpdateTimeUtc(v string) *GetAgenticSpaceResponseBodyAgenticSpace {
	s.UpdateTimeUtc = &v
	return s
}

func (s *GetAgenticSpaceResponseBodyAgenticSpace) Validate() error {
	if s.Quota != nil {
		if err := s.Quota.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAgenticSpaceResponseBodyAgenticSpaceQuota struct {
	// example:
	//
	// 10000000
	FileCountLimit *int64 `json:"FileCountLimit,omitempty" xml:"FileCountLimit,omitempty"`
	// example:
	//
	// 10737418240
	SizeLimit *int64 `json:"SizeLimit,omitempty" xml:"SizeLimit,omitempty"`
}

func (s GetAgenticSpaceResponseBodyAgenticSpaceQuota) String() string {
	return dara.Prettify(s)
}

func (s GetAgenticSpaceResponseBodyAgenticSpaceQuota) GoString() string {
	return s.String()
}

func (s *GetAgenticSpaceResponseBodyAgenticSpaceQuota) GetFileCountLimit() *int64 {
	return s.FileCountLimit
}

func (s *GetAgenticSpaceResponseBodyAgenticSpaceQuota) GetSizeLimit() *int64 {
	return s.SizeLimit
}

func (s *GetAgenticSpaceResponseBodyAgenticSpaceQuota) SetFileCountLimit(v int64) *GetAgenticSpaceResponseBodyAgenticSpaceQuota {
	s.FileCountLimit = &v
	return s
}

func (s *GetAgenticSpaceResponseBodyAgenticSpaceQuota) SetSizeLimit(v int64) *GetAgenticSpaceResponseBodyAgenticSpaceQuota {
	s.SizeLimit = &v
	return s
}

func (s *GetAgenticSpaceResponseBodyAgenticSpaceQuota) Validate() error {
	return dara.Validate(s)
}
