// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeKeyPairsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetKeyPairs(v *DescribeKeyPairsResponseBodyKeyPairs) *DescribeKeyPairsResponseBody
	GetKeyPairs() *DescribeKeyPairsResponseBodyKeyPairs
	SetPageNumber(v int32) *DescribeKeyPairsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeKeyPairsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeKeyPairsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeKeyPairsResponseBody
	GetTotalCount() *int32
}

type DescribeKeyPairsResponseBody struct {
	// The information of the key pairs.
	KeyPairs *DescribeKeyPairsResponseBodyKeyPairs `json:"KeyPairs,omitempty" xml:"KeyPairs,omitempty" type:"Struct"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of key pairs.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeKeyPairsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeKeyPairsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeKeyPairsResponseBody) GetKeyPairs() *DescribeKeyPairsResponseBodyKeyPairs {
	return s.KeyPairs
}

func (s *DescribeKeyPairsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeKeyPairsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeKeyPairsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeKeyPairsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeKeyPairsResponseBody) SetKeyPairs(v *DescribeKeyPairsResponseBodyKeyPairs) *DescribeKeyPairsResponseBody {
	s.KeyPairs = v
	return s
}

func (s *DescribeKeyPairsResponseBody) SetPageNumber(v int32) *DescribeKeyPairsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeKeyPairsResponseBody) SetPageSize(v int32) *DescribeKeyPairsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeKeyPairsResponseBody) SetRequestId(v string) *DescribeKeyPairsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeKeyPairsResponseBody) SetTotalCount(v int32) *DescribeKeyPairsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeKeyPairsResponseBody) Validate() error {
	if s.KeyPairs != nil {
		if err := s.KeyPairs.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeKeyPairsResponseBodyKeyPairs struct {
	KeyPair []*DescribeKeyPairsResponseBodyKeyPairsKeyPair `json:"KeyPair,omitempty" xml:"KeyPair,omitempty" type:"Repeated"`
}

func (s DescribeKeyPairsResponseBodyKeyPairs) String() string {
	return dara.Prettify(s)
}

func (s DescribeKeyPairsResponseBodyKeyPairs) GoString() string {
	return s.String()
}

func (s *DescribeKeyPairsResponseBodyKeyPairs) GetKeyPair() []*DescribeKeyPairsResponseBodyKeyPairsKeyPair {
	return s.KeyPair
}

func (s *DescribeKeyPairsResponseBodyKeyPairs) SetKeyPair(v []*DescribeKeyPairsResponseBodyKeyPairsKeyPair) *DescribeKeyPairsResponseBodyKeyPairs {
	s.KeyPair = v
	return s
}

func (s *DescribeKeyPairsResponseBodyKeyPairs) Validate() error {
	if s.KeyPair != nil {
		for _, item := range s.KeyPair {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeKeyPairsResponseBodyKeyPairsKeyPair struct {
	// The time when the key pair was created.
	//
	// example:
	//
	// 2023-09-04T08:33Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The fingerprint of the key pair.
	//
	// example:
	//
	// ABC1234567
	KeyPairFingerPrint *string `json:"KeyPairFingerPrint,omitempty" xml:"KeyPairFingerPrint,omitempty"`
	// The name of the key pair.
	//
	// example:
	//
	// testKeyPairName
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// The content of the public key.
	//
	// example:
	//
	// ssh-rsa****
	PublicKey *string `json:"PublicKey,omitempty" xml:"PublicKey,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-amnhr7u7c7hj****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags of the key pair.
	Tags *DescribeKeyPairsResponseBodyKeyPairsKeyPairTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
}

func (s DescribeKeyPairsResponseBodyKeyPairsKeyPair) String() string {
	return dara.Prettify(s)
}

func (s DescribeKeyPairsResponseBodyKeyPairsKeyPair) GoString() string {
	return s.String()
}

func (s *DescribeKeyPairsResponseBodyKeyPairsKeyPair) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeKeyPairsResponseBodyKeyPairsKeyPair) GetKeyPairFingerPrint() *string {
	return s.KeyPairFingerPrint
}

func (s *DescribeKeyPairsResponseBodyKeyPairsKeyPair) GetKeyPairName() *string {
	return s.KeyPairName
}

func (s *DescribeKeyPairsResponseBodyKeyPairsKeyPair) GetPublicKey() *string {
	return s.PublicKey
}

func (s *DescribeKeyPairsResponseBodyKeyPairsKeyPair) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeKeyPairsResponseBodyKeyPairsKeyPair) GetTags() *DescribeKeyPairsResponseBodyKeyPairsKeyPairTags {
	return s.Tags
}

func (s *DescribeKeyPairsResponseBodyKeyPairsKeyPair) SetCreationTime(v string) *DescribeKeyPairsResponseBodyKeyPairsKeyPair {
	s.CreationTime = &v
	return s
}

func (s *DescribeKeyPairsResponseBodyKeyPairsKeyPair) SetKeyPairFingerPrint(v string) *DescribeKeyPairsResponseBodyKeyPairsKeyPair {
	s.KeyPairFingerPrint = &v
	return s
}

func (s *DescribeKeyPairsResponseBodyKeyPairsKeyPair) SetKeyPairName(v string) *DescribeKeyPairsResponseBodyKeyPairsKeyPair {
	s.KeyPairName = &v
	return s
}

func (s *DescribeKeyPairsResponseBodyKeyPairsKeyPair) SetPublicKey(v string) *DescribeKeyPairsResponseBodyKeyPairsKeyPair {
	s.PublicKey = &v
	return s
}

func (s *DescribeKeyPairsResponseBodyKeyPairsKeyPair) SetResourceGroupId(v string) *DescribeKeyPairsResponseBodyKeyPairsKeyPair {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeKeyPairsResponseBodyKeyPairsKeyPair) SetTags(v *DescribeKeyPairsResponseBodyKeyPairsKeyPairTags) *DescribeKeyPairsResponseBodyKeyPairsKeyPair {
	s.Tags = v
	return s
}

func (s *DescribeKeyPairsResponseBodyKeyPairsKeyPair) Validate() error {
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeKeyPairsResponseBodyKeyPairsKeyPairTags struct {
	Tag []*DescribeKeyPairsResponseBodyKeyPairsKeyPairTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeKeyPairsResponseBodyKeyPairsKeyPairTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeKeyPairsResponseBodyKeyPairsKeyPairTags) GoString() string {
	return s.String()
}

func (s *DescribeKeyPairsResponseBodyKeyPairsKeyPairTags) GetTag() []*DescribeKeyPairsResponseBodyKeyPairsKeyPairTagsTag {
	return s.Tag
}

func (s *DescribeKeyPairsResponseBodyKeyPairsKeyPairTags) SetTag(v []*DescribeKeyPairsResponseBodyKeyPairsKeyPairTagsTag) *DescribeKeyPairsResponseBodyKeyPairsKeyPairTags {
	s.Tag = v
	return s
}

func (s *DescribeKeyPairsResponseBodyKeyPairsKeyPairTags) Validate() error {
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

type DescribeKeyPairsResponseBodyKeyPairsKeyPairTagsTag struct {
	// The tag key of the key pair.
	//
	// example:
	//
	// TestKey
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value of the key pair.
	//
	// example:
	//
	// TestValue
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeKeyPairsResponseBodyKeyPairsKeyPairTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeKeyPairsResponseBodyKeyPairsKeyPairTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeKeyPairsResponseBodyKeyPairsKeyPairTagsTag) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeKeyPairsResponseBodyKeyPairsKeyPairTagsTag) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeKeyPairsResponseBodyKeyPairsKeyPairTagsTag) SetTagKey(v string) *DescribeKeyPairsResponseBodyKeyPairsKeyPairTagsTag {
	s.TagKey = &v
	return s
}

func (s *DescribeKeyPairsResponseBodyKeyPairsKeyPairTagsTag) SetTagValue(v string) *DescribeKeyPairsResponseBodyKeyPairsKeyPairTagsTag {
	s.TagValue = &v
	return s
}

func (s *DescribeKeyPairsResponseBodyKeyPairsKeyPairTagsTag) Validate() error {
	return dara.Validate(s)
}
