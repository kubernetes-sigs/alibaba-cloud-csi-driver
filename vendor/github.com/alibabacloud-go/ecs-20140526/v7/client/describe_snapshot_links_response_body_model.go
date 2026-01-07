// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSnapshotLinksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *DescribeSnapshotLinksResponseBody
	GetNextToken() *string
	SetPageNumber(v int32) *DescribeSnapshotLinksResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeSnapshotLinksResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeSnapshotLinksResponseBody
	GetRequestId() *string
	SetSnapshotLinks(v *DescribeSnapshotLinksResponseBodySnapshotLinks) *DescribeSnapshotLinksResponseBody
	GetSnapshotLinks() *DescribeSnapshotLinksResponseBodySnapshotLinks
	SetTotalCount(v int32) *DescribeSnapshotLinksResponseBody
	GetTotalCount() *int32
}

type DescribeSnapshotLinksResponseBody struct {
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
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
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the snapshot chains.
	SnapshotLinks *DescribeSnapshotLinksResponseBodySnapshotLinks `json:"SnapshotLinks,omitempty" xml:"SnapshotLinks,omitempty" type:"Struct"`
	// The total number of snapshot chains.
	//
	// > When using the `MaxResults` and `NextToken` parameters for a paginated query, the returned `TotalCount` parameter value is invalid.
	//
	// example:
	//
	// 9
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSnapshotLinksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSnapshotLinksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSnapshotLinksResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeSnapshotLinksResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSnapshotLinksResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSnapshotLinksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSnapshotLinksResponseBody) GetSnapshotLinks() *DescribeSnapshotLinksResponseBodySnapshotLinks {
	return s.SnapshotLinks
}

func (s *DescribeSnapshotLinksResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeSnapshotLinksResponseBody) SetNextToken(v string) *DescribeSnapshotLinksResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeSnapshotLinksResponseBody) SetPageNumber(v int32) *DescribeSnapshotLinksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSnapshotLinksResponseBody) SetPageSize(v int32) *DescribeSnapshotLinksResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeSnapshotLinksResponseBody) SetRequestId(v string) *DescribeSnapshotLinksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSnapshotLinksResponseBody) SetSnapshotLinks(v *DescribeSnapshotLinksResponseBodySnapshotLinks) *DescribeSnapshotLinksResponseBody {
	s.SnapshotLinks = v
	return s
}

func (s *DescribeSnapshotLinksResponseBody) SetTotalCount(v int32) *DescribeSnapshotLinksResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeSnapshotLinksResponseBody) Validate() error {
	if s.SnapshotLinks != nil {
		if err := s.SnapshotLinks.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSnapshotLinksResponseBodySnapshotLinks struct {
	SnapshotLink []*DescribeSnapshotLinksResponseBodySnapshotLinksSnapshotLink `json:"SnapshotLink,omitempty" xml:"SnapshotLink,omitempty" type:"Repeated"`
}

func (s DescribeSnapshotLinksResponseBodySnapshotLinks) String() string {
	return dara.Prettify(s)
}

func (s DescribeSnapshotLinksResponseBodySnapshotLinks) GoString() string {
	return s.String()
}

func (s *DescribeSnapshotLinksResponseBodySnapshotLinks) GetSnapshotLink() []*DescribeSnapshotLinksResponseBodySnapshotLinksSnapshotLink {
	return s.SnapshotLink
}

func (s *DescribeSnapshotLinksResponseBodySnapshotLinks) SetSnapshotLink(v []*DescribeSnapshotLinksResponseBodySnapshotLinksSnapshotLink) *DescribeSnapshotLinksResponseBodySnapshotLinks {
	s.SnapshotLink = v
	return s
}

func (s *DescribeSnapshotLinksResponseBodySnapshotLinks) Validate() error {
	if s.SnapshotLink != nil {
		for _, item := range s.SnapshotLink {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSnapshotLinksResponseBodySnapshotLinksSnapshotLink struct {
	// The type of the snapshot chain. Valid values:
	//
	// 	- standard: standard snapshot chain.
	//
	// 	- archive: archive snapshot chain.
	//
	// 	- flash: instant access snapshot chain.
	//
	// example:
	//
	// standard
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// i-bp1h6jmbefj2cyqs****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the instance.
	//
	// example:
	//
	// testInstanceName
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// Indicates whether the instant access feature is enabled. Valid values:
	//
	// 	- true: The instant access feature is enabled. The feature can be enabled only for Enterprise SSDs (ESSDs).
	//
	// 	- false: The instant access feature is disabled. The snapshot is a standard snapshot for which the instant access feature is disabled.
	//
	// >  This parameter is no longer used. By default, standard snapshots of ESSDs are upgraded to instant access snapshots free of charge without the need for additional configurations. For more information, see [Use the instant access feature](https://help.aliyun.com/document_detail/193667.html).
	//
	// example:
	//
	// false
	InstantAccess *bool `json:"InstantAccess,omitempty" xml:"InstantAccess,omitempty"`
	// The region ID of the source disk.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the snapshot chain.
	//
	// example:
	//
	// sl-2ze0y1jwzpb1geqx****
	SnapshotLinkId *string `json:"SnapshotLinkId,omitempty" xml:"SnapshotLinkId,omitempty"`
	// The ID of the source disk. This parameter is retained even if the source disk is deleted.
	//
	// example:
	//
	// d-bp1d6tsvznfghy7y****
	SourceDiskId *string `json:"SourceDiskId,omitempty" xml:"SourceDiskId,omitempty"`
	// The name of the source disk.
	//
	// example:
	//
	// testSourceDiskName
	SourceDiskName *string `json:"SourceDiskName,omitempty" xml:"SourceDiskName,omitempty"`
	// The capacity of the source disk. Unit: GiB.
	//
	// example:
	//
	// 40
	SourceDiskSize *int32 `json:"SourceDiskSize,omitempty" xml:"SourceDiskSize,omitempty"`
	// The type of the source disk. Valid values:
	//
	// 	- system: system disk
	//
	// 	- data: data disk
	//
	// example:
	//
	// data
	SourceDiskType *string `json:"SourceDiskType,omitempty" xml:"SourceDiskType,omitempty"`
	// The total number of snapshots.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The total size of all snapshots in the snapshot chain. Unit: byte.
	//
	// example:
	//
	// 2097152
	TotalSize *int64 `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s DescribeSnapshotLinksResponseBodySnapshotLinksSnapshotLink) String() string {
	return dara.Prettify(s)
}

func (s DescribeSnapshotLinksResponseBodySnapshotLinksSnapshotLink) GoString() string {
	return s.String()
}

func (s *DescribeSnapshotLinksResponseBodySnapshotLinksSnapshotLink) GetCategory() *string {
	return s.Category
}

func (s *DescribeSnapshotLinksResponseBodySnapshotLinksSnapshotLink) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeSnapshotLinksResponseBodySnapshotLinksSnapshotLink) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeSnapshotLinksResponseBodySnapshotLinksSnapshotLink) GetInstantAccess() *bool {
	return s.InstantAccess
}

func (s *DescribeSnapshotLinksResponseBodySnapshotLinksSnapshotLink) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSnapshotLinksResponseBodySnapshotLinksSnapshotLink) GetSnapshotLinkId() *string {
	return s.SnapshotLinkId
}

func (s *DescribeSnapshotLinksResponseBodySnapshotLinksSnapshotLink) GetSourceDiskId() *string {
	return s.SourceDiskId
}

func (s *DescribeSnapshotLinksResponseBodySnapshotLinksSnapshotLink) GetSourceDiskName() *string {
	return s.SourceDiskName
}

func (s *DescribeSnapshotLinksResponseBodySnapshotLinksSnapshotLink) GetSourceDiskSize() *int32 {
	return s.SourceDiskSize
}

func (s *DescribeSnapshotLinksResponseBodySnapshotLinksSnapshotLink) GetSourceDiskType() *string {
	return s.SourceDiskType
}

func (s *DescribeSnapshotLinksResponseBodySnapshotLinksSnapshotLink) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeSnapshotLinksResponseBodySnapshotLinksSnapshotLink) GetTotalSize() *int64 {
	return s.TotalSize
}

func (s *DescribeSnapshotLinksResponseBodySnapshotLinksSnapshotLink) SetCategory(v string) *DescribeSnapshotLinksResponseBodySnapshotLinksSnapshotLink {
	s.Category = &v
	return s
}

func (s *DescribeSnapshotLinksResponseBodySnapshotLinksSnapshotLink) SetInstanceId(v string) *DescribeSnapshotLinksResponseBodySnapshotLinksSnapshotLink {
	s.InstanceId = &v
	return s
}

func (s *DescribeSnapshotLinksResponseBodySnapshotLinksSnapshotLink) SetInstanceName(v string) *DescribeSnapshotLinksResponseBodySnapshotLinksSnapshotLink {
	s.InstanceName = &v
	return s
}

func (s *DescribeSnapshotLinksResponseBodySnapshotLinksSnapshotLink) SetInstantAccess(v bool) *DescribeSnapshotLinksResponseBodySnapshotLinksSnapshotLink {
	s.InstantAccess = &v
	return s
}

func (s *DescribeSnapshotLinksResponseBodySnapshotLinksSnapshotLink) SetRegionId(v string) *DescribeSnapshotLinksResponseBodySnapshotLinksSnapshotLink {
	s.RegionId = &v
	return s
}

func (s *DescribeSnapshotLinksResponseBodySnapshotLinksSnapshotLink) SetSnapshotLinkId(v string) *DescribeSnapshotLinksResponseBodySnapshotLinksSnapshotLink {
	s.SnapshotLinkId = &v
	return s
}

func (s *DescribeSnapshotLinksResponseBodySnapshotLinksSnapshotLink) SetSourceDiskId(v string) *DescribeSnapshotLinksResponseBodySnapshotLinksSnapshotLink {
	s.SourceDiskId = &v
	return s
}

func (s *DescribeSnapshotLinksResponseBodySnapshotLinksSnapshotLink) SetSourceDiskName(v string) *DescribeSnapshotLinksResponseBodySnapshotLinksSnapshotLink {
	s.SourceDiskName = &v
	return s
}

func (s *DescribeSnapshotLinksResponseBodySnapshotLinksSnapshotLink) SetSourceDiskSize(v int32) *DescribeSnapshotLinksResponseBodySnapshotLinksSnapshotLink {
	s.SourceDiskSize = &v
	return s
}

func (s *DescribeSnapshotLinksResponseBodySnapshotLinksSnapshotLink) SetSourceDiskType(v string) *DescribeSnapshotLinksResponseBodySnapshotLinksSnapshotLink {
	s.SourceDiskType = &v
	return s
}

func (s *DescribeSnapshotLinksResponseBodySnapshotLinksSnapshotLink) SetTotalCount(v int32) *DescribeSnapshotLinksResponseBodySnapshotLinksSnapshotLink {
	s.TotalCount = &v
	return s
}

func (s *DescribeSnapshotLinksResponseBodySnapshotLinksSnapshotLink) SetTotalSize(v int64) *DescribeSnapshotLinksResponseBodySnapshotLinksSnapshotLink {
	s.TotalSize = &v
	return s
}

func (s *DescribeSnapshotLinksResponseBodySnapshotLinksSnapshotLink) Validate() error {
	return dara.Validate(s)
}
