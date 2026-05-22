// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVscsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *DescribeVscsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeVscsResponseBody
	GetRequestId() *string
	SetVscs(v []*DescribeVscsResponseBodyVscs) *DescribeVscsResponseBody
	GetVscs() []*DescribeVscsResponseBodyVscs
}

type DescribeVscsResponseBody struct {
	// example:
	//
	// AAAAAdDWBF2
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 473469C7-AA6F-**-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// VSC
	Vscs []*DescribeVscsResponseBodyVscs `json:"Vscs,omitempty" xml:"Vscs,omitempty" type:"Repeated"`
}

func (s DescribeVscsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVscsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVscsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeVscsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVscsResponseBody) GetVscs() []*DescribeVscsResponseBodyVscs {
	return s.Vscs
}

func (s *DescribeVscsResponseBody) SetNextToken(v string) *DescribeVscsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeVscsResponseBody) SetRequestId(v string) *DescribeVscsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVscsResponseBody) SetVscs(v []*DescribeVscsResponseBodyVscs) *DescribeVscsResponseBody {
	s.Vscs = v
	return s
}

func (s *DescribeVscsResponseBody) Validate() error {
	if s.Vscs != nil {
		for _, item := range s.Vscs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVscsResponseBodyVscs struct {
	// example:
	//
	// ali***-post-cn-j4g45iqze00f
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// i-uf69***21l8zuoizdq
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// rg-aek2zex4ehdyjvq
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// In_use
	Status *string                             `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags   []*DescribeVscsResponseBodyVscsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// VSC ID。
	//
	// example:
	//
	// vsc-hp34ue**g0wmycb27bwal
	VscId *string `json:"VscId,omitempty" xml:"VscId,omitempty"`
	// example:
	//
	// test-vsc
	VscName *string `json:"VscName,omitempty" xml:"VscName,omitempty"`
	// example:
	//
	// Primary
	VscType *string `json:"VscType,omitempty" xml:"VscType,omitempty"`
}

func (s DescribeVscsResponseBodyVscs) String() string {
	return dara.Prettify(s)
}

func (s DescribeVscsResponseBodyVscs) GoString() string {
	return s.String()
}

func (s *DescribeVscsResponseBodyVscs) GetDescription() *string {
	return s.Description
}

func (s *DescribeVscsResponseBodyVscs) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeVscsResponseBodyVscs) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeVscsResponseBodyVscs) GetStatus() *string {
	return s.Status
}

func (s *DescribeVscsResponseBodyVscs) GetTags() []*DescribeVscsResponseBodyVscsTags {
	return s.Tags
}

func (s *DescribeVscsResponseBodyVscs) GetVscId() *string {
	return s.VscId
}

func (s *DescribeVscsResponseBodyVscs) GetVscName() *string {
	return s.VscName
}

func (s *DescribeVscsResponseBodyVscs) GetVscType() *string {
	return s.VscType
}

func (s *DescribeVscsResponseBodyVscs) SetDescription(v string) *DescribeVscsResponseBodyVscs {
	s.Description = &v
	return s
}

func (s *DescribeVscsResponseBodyVscs) SetInstanceId(v string) *DescribeVscsResponseBodyVscs {
	s.InstanceId = &v
	return s
}

func (s *DescribeVscsResponseBodyVscs) SetResourceGroupId(v string) *DescribeVscsResponseBodyVscs {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeVscsResponseBodyVscs) SetStatus(v string) *DescribeVscsResponseBodyVscs {
	s.Status = &v
	return s
}

func (s *DescribeVscsResponseBodyVscs) SetTags(v []*DescribeVscsResponseBodyVscsTags) *DescribeVscsResponseBodyVscs {
	s.Tags = v
	return s
}

func (s *DescribeVscsResponseBodyVscs) SetVscId(v string) *DescribeVscsResponseBodyVscs {
	s.VscId = &v
	return s
}

func (s *DescribeVscsResponseBodyVscs) SetVscName(v string) *DescribeVscsResponseBodyVscs {
	s.VscName = &v
	return s
}

func (s *DescribeVscsResponseBodyVscs) SetVscType(v string) *DescribeVscsResponseBodyVscs {
	s.VscType = &v
	return s
}

func (s *DescribeVscsResponseBodyVscs) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVscsResponseBodyVscsTags struct {
	// example:
	//
	// name
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// example:
	//
	// 15
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeVscsResponseBodyVscsTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeVscsResponseBodyVscsTags) GoString() string {
	return s.String()
}

func (s *DescribeVscsResponseBodyVscsTags) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeVscsResponseBodyVscsTags) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeVscsResponseBodyVscsTags) SetTagKey(v string) *DescribeVscsResponseBodyVscsTags {
	s.TagKey = &v
	return s
}

func (s *DescribeVscsResponseBodyVscsTags) SetTagValue(v string) *DescribeVscsResponseBodyVscsTags {
	s.TagValue = &v
	return s
}

func (s *DescribeVscsResponseBodyVscsTags) Validate() error {
	return dara.Validate(s)
}
