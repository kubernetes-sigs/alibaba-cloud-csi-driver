// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageSharePermissionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccounts(v *DescribeImageSharePermissionResponseBodyAccounts) *DescribeImageSharePermissionResponseBody
	GetAccounts() *DescribeImageSharePermissionResponseBodyAccounts
	SetImageId(v string) *DescribeImageSharePermissionResponseBody
	GetImageId() *string
	SetPageNumber(v int32) *DescribeImageSharePermissionResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeImageSharePermissionResponseBody
	GetPageSize() *int32
	SetRegionId(v string) *DescribeImageSharePermissionResponseBody
	GetRegionId() *string
	SetRequestId(v string) *DescribeImageSharePermissionResponseBody
	GetRequestId() *string
	SetShareGroups(v *DescribeImageSharePermissionResponseBodyShareGroups) *DescribeImageSharePermissionResponseBody
	GetShareGroups() *DescribeImageSharePermissionResponseBodyShareGroups
	SetTotalCount(v int32) *DescribeImageSharePermissionResponseBody
	GetTotalCount() *int32
}

type DescribeImageSharePermissionResponseBody struct {
	// The Alibaba Cloud accounts.
	Accounts *DescribeImageSharePermissionResponseBodyAccounts `json:"Accounts,omitempty" xml:"Accounts,omitempty" type:"Struct"`
	// The ID of the custom image.
	//
	// example:
	//
	// m-bp1caf3yicx5jlfl****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
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
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the custom image.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The shared groups.
	ShareGroups *DescribeImageSharePermissionResponseBodyShareGroups `json:"ShareGroups,omitempty" xml:"ShareGroups,omitempty" type:"Struct"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeImageSharePermissionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageSharePermissionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeImageSharePermissionResponseBody) GetAccounts() *DescribeImageSharePermissionResponseBodyAccounts {
	return s.Accounts
}

func (s *DescribeImageSharePermissionResponseBody) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeImageSharePermissionResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeImageSharePermissionResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeImageSharePermissionResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeImageSharePermissionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeImageSharePermissionResponseBody) GetShareGroups() *DescribeImageSharePermissionResponseBodyShareGroups {
	return s.ShareGroups
}

func (s *DescribeImageSharePermissionResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeImageSharePermissionResponseBody) SetAccounts(v *DescribeImageSharePermissionResponseBodyAccounts) *DescribeImageSharePermissionResponseBody {
	s.Accounts = v
	return s
}

func (s *DescribeImageSharePermissionResponseBody) SetImageId(v string) *DescribeImageSharePermissionResponseBody {
	s.ImageId = &v
	return s
}

func (s *DescribeImageSharePermissionResponseBody) SetPageNumber(v int32) *DescribeImageSharePermissionResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeImageSharePermissionResponseBody) SetPageSize(v int32) *DescribeImageSharePermissionResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeImageSharePermissionResponseBody) SetRegionId(v string) *DescribeImageSharePermissionResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeImageSharePermissionResponseBody) SetRequestId(v string) *DescribeImageSharePermissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeImageSharePermissionResponseBody) SetShareGroups(v *DescribeImageSharePermissionResponseBodyShareGroups) *DescribeImageSharePermissionResponseBody {
	s.ShareGroups = v
	return s
}

func (s *DescribeImageSharePermissionResponseBody) SetTotalCount(v int32) *DescribeImageSharePermissionResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeImageSharePermissionResponseBody) Validate() error {
	if s.Accounts != nil {
		if err := s.Accounts.Validate(); err != nil {
			return err
		}
	}
	if s.ShareGroups != nil {
		if err := s.ShareGroups.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeImageSharePermissionResponseBodyAccounts struct {
	Account []*DescribeImageSharePermissionResponseBodyAccountsAccount `json:"Account,omitempty" xml:"Account,omitempty" type:"Repeated"`
}

func (s DescribeImageSharePermissionResponseBodyAccounts) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageSharePermissionResponseBodyAccounts) GoString() string {
	return s.String()
}

func (s *DescribeImageSharePermissionResponseBodyAccounts) GetAccount() []*DescribeImageSharePermissionResponseBodyAccountsAccount {
	return s.Account
}

func (s *DescribeImageSharePermissionResponseBodyAccounts) SetAccount(v []*DescribeImageSharePermissionResponseBodyAccountsAccount) *DescribeImageSharePermissionResponseBodyAccounts {
	s.Account = v
	return s
}

func (s *DescribeImageSharePermissionResponseBodyAccounts) Validate() error {
	if s.Account != nil {
		for _, item := range s.Account {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeImageSharePermissionResponseBodyAccountsAccount struct {
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 1234567890
	AliyunId *string `json:"AliyunId,omitempty" xml:"AliyunId,omitempty"`
	// The time when the image was shared. The time follows the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-01-01T12:05:00Z
	SharedTime *string `json:"SharedTime,omitempty" xml:"SharedTime,omitempty"`
}

func (s DescribeImageSharePermissionResponseBodyAccountsAccount) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageSharePermissionResponseBodyAccountsAccount) GoString() string {
	return s.String()
}

func (s *DescribeImageSharePermissionResponseBodyAccountsAccount) GetAliyunId() *string {
	return s.AliyunId
}

func (s *DescribeImageSharePermissionResponseBodyAccountsAccount) GetSharedTime() *string {
	return s.SharedTime
}

func (s *DescribeImageSharePermissionResponseBodyAccountsAccount) SetAliyunId(v string) *DescribeImageSharePermissionResponseBodyAccountsAccount {
	s.AliyunId = &v
	return s
}

func (s *DescribeImageSharePermissionResponseBodyAccountsAccount) SetSharedTime(v string) *DescribeImageSharePermissionResponseBodyAccountsAccount {
	s.SharedTime = &v
	return s
}

func (s *DescribeImageSharePermissionResponseBodyAccountsAccount) Validate() error {
	return dara.Validate(s)
}

type DescribeImageSharePermissionResponseBodyShareGroups struct {
	ShareGroup []*DescribeImageSharePermissionResponseBodyShareGroupsShareGroup `json:"ShareGroup,omitempty" xml:"ShareGroup,omitempty" type:"Repeated"`
}

func (s DescribeImageSharePermissionResponseBodyShareGroups) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageSharePermissionResponseBodyShareGroups) GoString() string {
	return s.String()
}

func (s *DescribeImageSharePermissionResponseBodyShareGroups) GetShareGroup() []*DescribeImageSharePermissionResponseBodyShareGroupsShareGroup {
	return s.ShareGroup
}

func (s *DescribeImageSharePermissionResponseBodyShareGroups) SetShareGroup(v []*DescribeImageSharePermissionResponseBodyShareGroupsShareGroup) *DescribeImageSharePermissionResponseBodyShareGroups {
	s.ShareGroup = v
	return s
}

func (s *DescribeImageSharePermissionResponseBodyShareGroups) Validate() error {
	if s.ShareGroup != nil {
		for _, item := range s.ShareGroup {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeImageSharePermissionResponseBodyShareGroupsShareGroup struct {
	// The shared group.
	//
	// example:
	//
	// all
	Group *string `json:"Group,omitempty" xml:"Group,omitempty"`
}

func (s DescribeImageSharePermissionResponseBodyShareGroupsShareGroup) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageSharePermissionResponseBodyShareGroupsShareGroup) GoString() string {
	return s.String()
}

func (s *DescribeImageSharePermissionResponseBodyShareGroupsShareGroup) GetGroup() *string {
	return s.Group
}

func (s *DescribeImageSharePermissionResponseBodyShareGroupsShareGroup) SetGroup(v string) *DescribeImageSharePermissionResponseBodyShareGroupsShareGroup {
	s.Group = &v
	return s
}

func (s *DescribeImageSharePermissionResponseBodyShareGroupsShareGroup) Validate() error {
	return dara.Validate(s)
}
