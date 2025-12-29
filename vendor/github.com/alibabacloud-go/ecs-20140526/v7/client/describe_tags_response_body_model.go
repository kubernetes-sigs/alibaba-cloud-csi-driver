// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTagsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeTagsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeTagsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeTagsResponseBody
	GetRequestId() *string
	SetTags(v *DescribeTagsResponseBodyTags) *DescribeTagsResponseBody
	GetTags() *DescribeTagsResponseBodyTags
	SetTotalCount(v int32) *DescribeTagsResponseBody
	GetTotalCount() *int32
}

type DescribeTagsResponseBody struct {
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
	// The ID of the request.
	//
	// example:
	//
	// B04B8CF3-4489-432D-83BA-6F128E4F2295
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The tags that match all filter conditions.
	Tags *DescribeTagsResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The total number of tags.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeTagsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTagsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeTagsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeTagsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTagsResponseBody) GetTags() *DescribeTagsResponseBodyTags {
	return s.Tags
}

func (s *DescribeTagsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeTagsResponseBody) SetPageNumber(v int32) *DescribeTagsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeTagsResponseBody) SetPageSize(v int32) *DescribeTagsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeTagsResponseBody) SetRequestId(v string) *DescribeTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTagsResponseBody) SetTags(v *DescribeTagsResponseBodyTags) *DescribeTagsResponseBody {
	s.Tags = v
	return s
}

func (s *DescribeTagsResponseBody) SetTotalCount(v int32) *DescribeTagsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeTagsResponseBody) Validate() error {
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeTagsResponseBodyTags struct {
	Tag []*DescribeTagsResponseBodyTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeTagsResponseBodyTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagsResponseBodyTags) GoString() string {
	return s.String()
}

func (s *DescribeTagsResponseBodyTags) GetTag() []*DescribeTagsResponseBodyTagsTag {
	return s.Tag
}

func (s *DescribeTagsResponseBodyTags) SetTag(v []*DescribeTagsResponseBodyTagsTag) *DescribeTagsResponseBodyTags {
	s.Tag = v
	return s
}

func (s *DescribeTagsResponseBodyTags) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeTagsResponseBodyTagsTag struct {
	// The number of resource types.
	ResourceTypeCount *DescribeTagsResponseBodyTagsTagResourceTypeCount `json:"ResourceTypeCount,omitempty" xml:"ResourceTypeCount,omitempty" type:"Struct"`
	// The tag key.
	//
	// example:
	//
	// TestKey
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
	//
	// example:
	//
	// TestValue
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeTagsResponseBodyTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagsResponseBodyTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeTagsResponseBodyTagsTag) GetResourceTypeCount() *DescribeTagsResponseBodyTagsTagResourceTypeCount {
	return s.ResourceTypeCount
}

func (s *DescribeTagsResponseBodyTagsTag) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeTagsResponseBodyTagsTag) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeTagsResponseBodyTagsTag) SetResourceTypeCount(v *DescribeTagsResponseBodyTagsTagResourceTypeCount) *DescribeTagsResponseBodyTagsTag {
	s.ResourceTypeCount = v
	return s
}

func (s *DescribeTagsResponseBodyTagsTag) SetTagKey(v string) *DescribeTagsResponseBodyTagsTag {
	s.TagKey = &v
	return s
}

func (s *DescribeTagsResponseBodyTagsTag) SetTagValue(v string) *DescribeTagsResponseBodyTagsTag {
	s.TagValue = &v
	return s
}

func (s *DescribeTagsResponseBodyTagsTag) Validate() error {
	if s.ResourceTypeCount != nil {
		if err := s.ResourceTypeCount.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeTagsResponseBodyTagsTagResourceTypeCount struct {
	// The number of dedicated hosts to which the tag is added.
	//
	// example:
	//
	// 1
	Ddh *int32 `json:"Ddh,omitempty" xml:"Ddh,omitempty"`
	// The number of disks to which the tag is added.
	//
	// example:
	//
	// 15
	Disk *int32 `json:"Disk,omitempty" xml:"Disk,omitempty"`
	// The number of ENIs to which the tag is added.
	//
	// example:
	//
	// 5
	Eni *int32 `json:"Eni,omitempty" xml:"Eni,omitempty"`
	// The number of images to which the tag is added.
	//
	// example:
	//
	// 6
	Image *int32 `json:"Image,omitempty" xml:"Image,omitempty"`
	// The number of instances to which the tag is added.
	//
	// example:
	//
	// 45
	Instance *int32 `json:"Instance,omitempty" xml:"Instance,omitempty"`
	// The number of key pairs to which the tag is added.
	//
	// example:
	//
	// 17
	KeyPair *int32 `json:"KeyPair,omitempty" xml:"KeyPair,omitempty"`
	// The number of launch templates to which the tag is added.
	//
	// example:
	//
	// 6
	LaunchTemplate *int32 `json:"LaunchTemplate,omitempty" xml:"LaunchTemplate,omitempty"`
	// The number of reserved instances to which the tag is added.
	//
	// example:
	//
	// 4
	ReservedInstance *int32 `json:"ReservedInstance,omitempty" xml:"ReservedInstance,omitempty"`
	// The number of security groups to which the tag is added.
	//
	// example:
	//
	// 4
	Securitygroup *int32 `json:"Securitygroup,omitempty" xml:"Securitygroup,omitempty"`
	// The number of snapshots to which the tag is added.
	//
	// example:
	//
	// 15
	Snapshot *int32 `json:"Snapshot,omitempty" xml:"Snapshot,omitempty"`
	// The number of automatic snapshot policies to which the tag is added.
	//
	// example:
	//
	// 4
	SnapshotPolicy *int32 `json:"SnapshotPolicy,omitempty" xml:"SnapshotPolicy,omitempty"`
	// The number of storage volumes to which the tag is added.
	//
	// example:
	//
	// 6
	Volume *int32 `json:"Volume,omitempty" xml:"Volume,omitempty"`
}

func (s DescribeTagsResponseBodyTagsTagResourceTypeCount) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagsResponseBodyTagsTagResourceTypeCount) GoString() string {
	return s.String()
}

func (s *DescribeTagsResponseBodyTagsTagResourceTypeCount) GetDdh() *int32 {
	return s.Ddh
}

func (s *DescribeTagsResponseBodyTagsTagResourceTypeCount) GetDisk() *int32 {
	return s.Disk
}

func (s *DescribeTagsResponseBodyTagsTagResourceTypeCount) GetEni() *int32 {
	return s.Eni
}

func (s *DescribeTagsResponseBodyTagsTagResourceTypeCount) GetImage() *int32 {
	return s.Image
}

func (s *DescribeTagsResponseBodyTagsTagResourceTypeCount) GetInstance() *int32 {
	return s.Instance
}

func (s *DescribeTagsResponseBodyTagsTagResourceTypeCount) GetKeyPair() *int32 {
	return s.KeyPair
}

func (s *DescribeTagsResponseBodyTagsTagResourceTypeCount) GetLaunchTemplate() *int32 {
	return s.LaunchTemplate
}

func (s *DescribeTagsResponseBodyTagsTagResourceTypeCount) GetReservedInstance() *int32 {
	return s.ReservedInstance
}

func (s *DescribeTagsResponseBodyTagsTagResourceTypeCount) GetSecuritygroup() *int32 {
	return s.Securitygroup
}

func (s *DescribeTagsResponseBodyTagsTagResourceTypeCount) GetSnapshot() *int32 {
	return s.Snapshot
}

func (s *DescribeTagsResponseBodyTagsTagResourceTypeCount) GetSnapshotPolicy() *int32 {
	return s.SnapshotPolicy
}

func (s *DescribeTagsResponseBodyTagsTagResourceTypeCount) GetVolume() *int32 {
	return s.Volume
}

func (s *DescribeTagsResponseBodyTagsTagResourceTypeCount) SetDdh(v int32) *DescribeTagsResponseBodyTagsTagResourceTypeCount {
	s.Ddh = &v
	return s
}

func (s *DescribeTagsResponseBodyTagsTagResourceTypeCount) SetDisk(v int32) *DescribeTagsResponseBodyTagsTagResourceTypeCount {
	s.Disk = &v
	return s
}

func (s *DescribeTagsResponseBodyTagsTagResourceTypeCount) SetEni(v int32) *DescribeTagsResponseBodyTagsTagResourceTypeCount {
	s.Eni = &v
	return s
}

func (s *DescribeTagsResponseBodyTagsTagResourceTypeCount) SetImage(v int32) *DescribeTagsResponseBodyTagsTagResourceTypeCount {
	s.Image = &v
	return s
}

func (s *DescribeTagsResponseBodyTagsTagResourceTypeCount) SetInstance(v int32) *DescribeTagsResponseBodyTagsTagResourceTypeCount {
	s.Instance = &v
	return s
}

func (s *DescribeTagsResponseBodyTagsTagResourceTypeCount) SetKeyPair(v int32) *DescribeTagsResponseBodyTagsTagResourceTypeCount {
	s.KeyPair = &v
	return s
}

func (s *DescribeTagsResponseBodyTagsTagResourceTypeCount) SetLaunchTemplate(v int32) *DescribeTagsResponseBodyTagsTagResourceTypeCount {
	s.LaunchTemplate = &v
	return s
}

func (s *DescribeTagsResponseBodyTagsTagResourceTypeCount) SetReservedInstance(v int32) *DescribeTagsResponseBodyTagsTagResourceTypeCount {
	s.ReservedInstance = &v
	return s
}

func (s *DescribeTagsResponseBodyTagsTagResourceTypeCount) SetSecuritygroup(v int32) *DescribeTagsResponseBodyTagsTagResourceTypeCount {
	s.Securitygroup = &v
	return s
}

func (s *DescribeTagsResponseBodyTagsTagResourceTypeCount) SetSnapshot(v int32) *DescribeTagsResponseBodyTagsTagResourceTypeCount {
	s.Snapshot = &v
	return s
}

func (s *DescribeTagsResponseBodyTagsTagResourceTypeCount) SetSnapshotPolicy(v int32) *DescribeTagsResponseBodyTagsTagResourceTypeCount {
	s.SnapshotPolicy = &v
	return s
}

func (s *DescribeTagsResponseBodyTagsTagResourceTypeCount) SetVolume(v int32) *DescribeTagsResponseBodyTagsTagResourceTypeCount {
	s.Volume = &v
	return s
}

func (s *DescribeTagsResponseBodyTagsTagResourceTypeCount) Validate() error {
	return dara.Validate(s)
}
