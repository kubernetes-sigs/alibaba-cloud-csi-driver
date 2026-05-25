// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImagePipelineExecutionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetImagePipelineExecution(v *DescribeImagePipelineExecutionsResponseBodyImagePipelineExecution) *DescribeImagePipelineExecutionsResponseBody
	GetImagePipelineExecution() *DescribeImagePipelineExecutionsResponseBodyImagePipelineExecution
	SetMaxResults(v int32) *DescribeImagePipelineExecutionsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeImagePipelineExecutionsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeImagePipelineExecutionsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeImagePipelineExecutionsResponseBody
	GetTotalCount() *int32
}

type DescribeImagePipelineExecutionsResponseBody struct {
	ImagePipelineExecution *DescribeImagePipelineExecutionsResponseBodyImagePipelineExecution `json:"ImagePipelineExecution,omitempty" xml:"ImagePipelineExecution,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of `NextToken`.
	//
	// example:
	//
	// AAAAAdDWBF2****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The maximum number of entries per page. Valid values: 1 to 500
	//
	// Default value: 50.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If NextToken is empty, no next page exists. For information about how to use the returned value, see the "Usage notes" section in this topic.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeImagePipelineExecutionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeImagePipelineExecutionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeImagePipelineExecutionsResponseBody) GetImagePipelineExecution() *DescribeImagePipelineExecutionsResponseBodyImagePipelineExecution {
	return s.ImagePipelineExecution
}

func (s *DescribeImagePipelineExecutionsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeImagePipelineExecutionsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeImagePipelineExecutionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeImagePipelineExecutionsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeImagePipelineExecutionsResponseBody) SetImagePipelineExecution(v *DescribeImagePipelineExecutionsResponseBodyImagePipelineExecution) *DescribeImagePipelineExecutionsResponseBody {
	s.ImagePipelineExecution = v
	return s
}

func (s *DescribeImagePipelineExecutionsResponseBody) SetMaxResults(v int32) *DescribeImagePipelineExecutionsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeImagePipelineExecutionsResponseBody) SetNextToken(v string) *DescribeImagePipelineExecutionsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeImagePipelineExecutionsResponseBody) SetRequestId(v string) *DescribeImagePipelineExecutionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeImagePipelineExecutionsResponseBody) SetTotalCount(v int32) *DescribeImagePipelineExecutionsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeImagePipelineExecutionsResponseBody) Validate() error {
	if s.ImagePipelineExecution != nil {
		if err := s.ImagePipelineExecution.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeImagePipelineExecutionsResponseBodyImagePipelineExecution struct {
	ImagePipelineExecutionSet []*DescribeImagePipelineExecutionsResponseBodyImagePipelineExecutionImagePipelineExecutionSet `json:"ImagePipelineExecutionSet,omitempty" xml:"ImagePipelineExecutionSet,omitempty" type:"Repeated"`
}

func (s DescribeImagePipelineExecutionsResponseBodyImagePipelineExecution) String() string {
	return dara.Prettify(s)
}

func (s DescribeImagePipelineExecutionsResponseBodyImagePipelineExecution) GoString() string {
	return s.String()
}

func (s *DescribeImagePipelineExecutionsResponseBodyImagePipelineExecution) GetImagePipelineExecutionSet() []*DescribeImagePipelineExecutionsResponseBodyImagePipelineExecutionImagePipelineExecutionSet {
	return s.ImagePipelineExecutionSet
}

func (s *DescribeImagePipelineExecutionsResponseBodyImagePipelineExecution) SetImagePipelineExecutionSet(v []*DescribeImagePipelineExecutionsResponseBodyImagePipelineExecutionImagePipelineExecutionSet) *DescribeImagePipelineExecutionsResponseBodyImagePipelineExecution {
	s.ImagePipelineExecutionSet = v
	return s
}

func (s *DescribeImagePipelineExecutionsResponseBodyImagePipelineExecution) Validate() error {
	if s.ImagePipelineExecutionSet != nil {
		for _, item := range s.ImagePipelineExecutionSet {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeImagePipelineExecutionsResponseBodyImagePipelineExecutionImagePipelineExecutionSet struct {
	CreationTime    *string                                                                                         `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	ErrorCode       *string                                                                                         `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorInfo       *string                                                                                         `json:"ErrorInfo,omitempty" xml:"ErrorInfo,omitempty"`
	ExecutionId     *string                                                                                         `json:"ExecutionId,omitempty" xml:"ExecutionId,omitempty"`
	ImageId         *string                                                                                         `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	ImagePipelineId *string                                                                                         `json:"ImagePipelineId,omitempty" xml:"ImagePipelineId,omitempty"`
	Message         *string                                                                                         `json:"Message,omitempty" xml:"Message,omitempty"`
	ModifiedTime    *string                                                                                         `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	ResourceGroupId *string                                                                                         `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Status          *string                                                                                         `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags            *DescribeImagePipelineExecutionsResponseBodyImagePipelineExecutionImagePipelineExecutionSetTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
}

func (s DescribeImagePipelineExecutionsResponseBodyImagePipelineExecutionImagePipelineExecutionSet) String() string {
	return dara.Prettify(s)
}

func (s DescribeImagePipelineExecutionsResponseBodyImagePipelineExecutionImagePipelineExecutionSet) GoString() string {
	return s.String()
}

func (s *DescribeImagePipelineExecutionsResponseBodyImagePipelineExecutionImagePipelineExecutionSet) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeImagePipelineExecutionsResponseBodyImagePipelineExecutionImagePipelineExecutionSet) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeImagePipelineExecutionsResponseBodyImagePipelineExecutionImagePipelineExecutionSet) GetErrorInfo() *string {
	return s.ErrorInfo
}

func (s *DescribeImagePipelineExecutionsResponseBodyImagePipelineExecutionImagePipelineExecutionSet) GetExecutionId() *string {
	return s.ExecutionId
}

func (s *DescribeImagePipelineExecutionsResponseBodyImagePipelineExecutionImagePipelineExecutionSet) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeImagePipelineExecutionsResponseBodyImagePipelineExecutionImagePipelineExecutionSet) GetImagePipelineId() *string {
	return s.ImagePipelineId
}

func (s *DescribeImagePipelineExecutionsResponseBodyImagePipelineExecutionImagePipelineExecutionSet) GetMessage() *string {
	return s.Message
}

func (s *DescribeImagePipelineExecutionsResponseBodyImagePipelineExecutionImagePipelineExecutionSet) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *DescribeImagePipelineExecutionsResponseBodyImagePipelineExecutionImagePipelineExecutionSet) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeImagePipelineExecutionsResponseBodyImagePipelineExecutionImagePipelineExecutionSet) GetStatus() *string {
	return s.Status
}

func (s *DescribeImagePipelineExecutionsResponseBodyImagePipelineExecutionImagePipelineExecutionSet) GetTags() *DescribeImagePipelineExecutionsResponseBodyImagePipelineExecutionImagePipelineExecutionSetTags {
	return s.Tags
}

func (s *DescribeImagePipelineExecutionsResponseBodyImagePipelineExecutionImagePipelineExecutionSet) SetCreationTime(v string) *DescribeImagePipelineExecutionsResponseBodyImagePipelineExecutionImagePipelineExecutionSet {
	s.CreationTime = &v
	return s
}

func (s *DescribeImagePipelineExecutionsResponseBodyImagePipelineExecutionImagePipelineExecutionSet) SetErrorCode(v string) *DescribeImagePipelineExecutionsResponseBodyImagePipelineExecutionImagePipelineExecutionSet {
	s.ErrorCode = &v
	return s
}

func (s *DescribeImagePipelineExecutionsResponseBodyImagePipelineExecutionImagePipelineExecutionSet) SetErrorInfo(v string) *DescribeImagePipelineExecutionsResponseBodyImagePipelineExecutionImagePipelineExecutionSet {
	s.ErrorInfo = &v
	return s
}

func (s *DescribeImagePipelineExecutionsResponseBodyImagePipelineExecutionImagePipelineExecutionSet) SetExecutionId(v string) *DescribeImagePipelineExecutionsResponseBodyImagePipelineExecutionImagePipelineExecutionSet {
	s.ExecutionId = &v
	return s
}

func (s *DescribeImagePipelineExecutionsResponseBodyImagePipelineExecutionImagePipelineExecutionSet) SetImageId(v string) *DescribeImagePipelineExecutionsResponseBodyImagePipelineExecutionImagePipelineExecutionSet {
	s.ImageId = &v
	return s
}

func (s *DescribeImagePipelineExecutionsResponseBodyImagePipelineExecutionImagePipelineExecutionSet) SetImagePipelineId(v string) *DescribeImagePipelineExecutionsResponseBodyImagePipelineExecutionImagePipelineExecutionSet {
	s.ImagePipelineId = &v
	return s
}

func (s *DescribeImagePipelineExecutionsResponseBodyImagePipelineExecutionImagePipelineExecutionSet) SetMessage(v string) *DescribeImagePipelineExecutionsResponseBodyImagePipelineExecutionImagePipelineExecutionSet {
	s.Message = &v
	return s
}

func (s *DescribeImagePipelineExecutionsResponseBodyImagePipelineExecutionImagePipelineExecutionSet) SetModifiedTime(v string) *DescribeImagePipelineExecutionsResponseBodyImagePipelineExecutionImagePipelineExecutionSet {
	s.ModifiedTime = &v
	return s
}

func (s *DescribeImagePipelineExecutionsResponseBodyImagePipelineExecutionImagePipelineExecutionSet) SetResourceGroupId(v string) *DescribeImagePipelineExecutionsResponseBodyImagePipelineExecutionImagePipelineExecutionSet {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeImagePipelineExecutionsResponseBodyImagePipelineExecutionImagePipelineExecutionSet) SetStatus(v string) *DescribeImagePipelineExecutionsResponseBodyImagePipelineExecutionImagePipelineExecutionSet {
	s.Status = &v
	return s
}

func (s *DescribeImagePipelineExecutionsResponseBodyImagePipelineExecutionImagePipelineExecutionSet) SetTags(v *DescribeImagePipelineExecutionsResponseBodyImagePipelineExecutionImagePipelineExecutionSetTags) *DescribeImagePipelineExecutionsResponseBodyImagePipelineExecutionImagePipelineExecutionSet {
	s.Tags = v
	return s
}

func (s *DescribeImagePipelineExecutionsResponseBodyImagePipelineExecutionImagePipelineExecutionSet) Validate() error {
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeImagePipelineExecutionsResponseBodyImagePipelineExecutionImagePipelineExecutionSetTags struct {
	Tag []*DescribeImagePipelineExecutionsResponseBodyImagePipelineExecutionImagePipelineExecutionSetTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeImagePipelineExecutionsResponseBodyImagePipelineExecutionImagePipelineExecutionSetTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeImagePipelineExecutionsResponseBodyImagePipelineExecutionImagePipelineExecutionSetTags) GoString() string {
	return s.String()
}

func (s *DescribeImagePipelineExecutionsResponseBodyImagePipelineExecutionImagePipelineExecutionSetTags) GetTag() []*DescribeImagePipelineExecutionsResponseBodyImagePipelineExecutionImagePipelineExecutionSetTagsTag {
	return s.Tag
}

func (s *DescribeImagePipelineExecutionsResponseBodyImagePipelineExecutionImagePipelineExecutionSetTags) SetTag(v []*DescribeImagePipelineExecutionsResponseBodyImagePipelineExecutionImagePipelineExecutionSetTagsTag) *DescribeImagePipelineExecutionsResponseBodyImagePipelineExecutionImagePipelineExecutionSetTags {
	s.Tag = v
	return s
}

func (s *DescribeImagePipelineExecutionsResponseBodyImagePipelineExecutionImagePipelineExecutionSetTags) Validate() error {
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

type DescribeImagePipelineExecutionsResponseBodyImagePipelineExecutionImagePipelineExecutionSetTagsTag struct {
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeImagePipelineExecutionsResponseBodyImagePipelineExecutionImagePipelineExecutionSetTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeImagePipelineExecutionsResponseBodyImagePipelineExecutionImagePipelineExecutionSetTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeImagePipelineExecutionsResponseBodyImagePipelineExecutionImagePipelineExecutionSetTagsTag) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeImagePipelineExecutionsResponseBodyImagePipelineExecutionImagePipelineExecutionSetTagsTag) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeImagePipelineExecutionsResponseBodyImagePipelineExecutionImagePipelineExecutionSetTagsTag) SetTagKey(v string) *DescribeImagePipelineExecutionsResponseBodyImagePipelineExecutionImagePipelineExecutionSetTagsTag {
	s.TagKey = &v
	return s
}

func (s *DescribeImagePipelineExecutionsResponseBodyImagePipelineExecutionImagePipelineExecutionSetTagsTag) SetTagValue(v string) *DescribeImagePipelineExecutionsResponseBodyImagePipelineExecutionImagePipelineExecutionSetTagsTag {
	s.TagValue = &v
	return s
}

func (s *DescribeImagePipelineExecutionsResponseBodyImagePipelineExecutionImagePipelineExecutionSetTagsTag) Validate() error {
	return dara.Validate(s)
}
