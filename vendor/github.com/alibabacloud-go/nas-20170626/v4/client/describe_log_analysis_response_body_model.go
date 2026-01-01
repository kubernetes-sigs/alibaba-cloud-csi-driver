// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLogAnalysisResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAnalyses(v *DescribeLogAnalysisResponseBodyAnalyses) *DescribeLogAnalysisResponseBody
	GetAnalyses() *DescribeLogAnalysisResponseBodyAnalyses
	SetCode(v string) *DescribeLogAnalysisResponseBody
	GetCode() *string
	SetPageNumber(v int32) *DescribeLogAnalysisResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeLogAnalysisResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeLogAnalysisResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeLogAnalysisResponseBody
	GetTotalCount() *int32
}

type DescribeLogAnalysisResponseBody struct {
	// The collection of log dump information.
	Analyses *DescribeLogAnalysisResponseBodyAnalyses `json:"Analyses,omitempty" xml:"Analyses,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of log dump entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C84F77AF-3DE5-48F1-B19B-37FCBE24****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of log dump entries in the region.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeLogAnalysisResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogAnalysisResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLogAnalysisResponseBody) GetAnalyses() *DescribeLogAnalysisResponseBodyAnalyses {
	return s.Analyses
}

func (s *DescribeLogAnalysisResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeLogAnalysisResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeLogAnalysisResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeLogAnalysisResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLogAnalysisResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeLogAnalysisResponseBody) SetAnalyses(v *DescribeLogAnalysisResponseBodyAnalyses) *DescribeLogAnalysisResponseBody {
	s.Analyses = v
	return s
}

func (s *DescribeLogAnalysisResponseBody) SetCode(v string) *DescribeLogAnalysisResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeLogAnalysisResponseBody) SetPageNumber(v int32) *DescribeLogAnalysisResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeLogAnalysisResponseBody) SetPageSize(v int32) *DescribeLogAnalysisResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeLogAnalysisResponseBody) SetRequestId(v string) *DescribeLogAnalysisResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLogAnalysisResponseBody) SetTotalCount(v int32) *DescribeLogAnalysisResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeLogAnalysisResponseBody) Validate() error {
	if s.Analyses != nil {
		if err := s.Analyses.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLogAnalysisResponseBodyAnalyses struct {
	Analysis []*DescribeLogAnalysisResponseBodyAnalysesAnalysis `json:"Analysis,omitempty" xml:"Analysis,omitempty" type:"Repeated"`
}

func (s DescribeLogAnalysisResponseBodyAnalyses) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogAnalysisResponseBodyAnalyses) GoString() string {
	return s.String()
}

func (s *DescribeLogAnalysisResponseBodyAnalyses) GetAnalysis() []*DescribeLogAnalysisResponseBodyAnalysesAnalysis {
	return s.Analysis
}

func (s *DescribeLogAnalysisResponseBodyAnalyses) SetAnalysis(v []*DescribeLogAnalysisResponseBodyAnalysesAnalysis) *DescribeLogAnalysisResponseBodyAnalyses {
	s.Analysis = v
	return s
}

func (s *DescribeLogAnalysisResponseBodyAnalyses) Validate() error {
	if s.Analysis != nil {
		for _, item := range s.Analysis {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLogAnalysisResponseBodyAnalysesAnalysis struct {
	// The ID of the file system.
	//
	// example:
	//
	// 0c7154xxxx
	MetaKey *string `json:"MetaKey,omitempty" xml:"MetaKey,omitempty"`
	// The log dump information of the file system.
	MetaValue *DescribeLogAnalysisResponseBodyAnalysesAnalysisMetaValue `json:"MetaValue,omitempty" xml:"MetaValue,omitempty" type:"Struct"`
}

func (s DescribeLogAnalysisResponseBodyAnalysesAnalysis) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogAnalysisResponseBodyAnalysesAnalysis) GoString() string {
	return s.String()
}

func (s *DescribeLogAnalysisResponseBodyAnalysesAnalysis) GetMetaKey() *string {
	return s.MetaKey
}

func (s *DescribeLogAnalysisResponseBodyAnalysesAnalysis) GetMetaValue() *DescribeLogAnalysisResponseBodyAnalysesAnalysisMetaValue {
	return s.MetaValue
}

func (s *DescribeLogAnalysisResponseBodyAnalysesAnalysis) SetMetaKey(v string) *DescribeLogAnalysisResponseBodyAnalysesAnalysis {
	s.MetaKey = &v
	return s
}

func (s *DescribeLogAnalysisResponseBodyAnalysesAnalysis) SetMetaValue(v *DescribeLogAnalysisResponseBodyAnalysesAnalysisMetaValue) *DescribeLogAnalysisResponseBodyAnalysesAnalysis {
	s.MetaValue = v
	return s
}

func (s *DescribeLogAnalysisResponseBodyAnalysesAnalysis) Validate() error {
	if s.MetaValue != nil {
		if err := s.MetaValue.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLogAnalysisResponseBodyAnalysesAnalysisMetaValue struct {
	// The name of the dedicated Logstore that is used to store NAS operation logs.
	//
	// example:
	//
	// nas-nfs
	Logstore *string `json:"Logstore,omitempty" xml:"Logstore,omitempty"`
	// The name of the project where the dedicated Logstore resides.
	//
	// example:
	//
	// nas-1746495857602745-cn-hangzhou
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
	// The region where the dedicated Logstore resides.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The role that is used by NAS to access Simple Log Service.
	//
	// example:
	//
	// acs:ram::162165525211xxxx:role/aliyunnaslogarchiverole
	RoleArn *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
}

func (s DescribeLogAnalysisResponseBodyAnalysesAnalysisMetaValue) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogAnalysisResponseBodyAnalysesAnalysisMetaValue) GoString() string {
	return s.String()
}

func (s *DescribeLogAnalysisResponseBodyAnalysesAnalysisMetaValue) GetLogstore() *string {
	return s.Logstore
}

func (s *DescribeLogAnalysisResponseBodyAnalysesAnalysisMetaValue) GetProject() *string {
	return s.Project
}

func (s *DescribeLogAnalysisResponseBodyAnalysesAnalysisMetaValue) GetRegion() *string {
	return s.Region
}

func (s *DescribeLogAnalysisResponseBodyAnalysesAnalysisMetaValue) GetRoleArn() *string {
	return s.RoleArn
}

func (s *DescribeLogAnalysisResponseBodyAnalysesAnalysisMetaValue) SetLogstore(v string) *DescribeLogAnalysisResponseBodyAnalysesAnalysisMetaValue {
	s.Logstore = &v
	return s
}

func (s *DescribeLogAnalysisResponseBodyAnalysesAnalysisMetaValue) SetProject(v string) *DescribeLogAnalysisResponseBodyAnalysesAnalysisMetaValue {
	s.Project = &v
	return s
}

func (s *DescribeLogAnalysisResponseBodyAnalysesAnalysisMetaValue) SetRegion(v string) *DescribeLogAnalysisResponseBodyAnalysesAnalysisMetaValue {
	s.Region = &v
	return s
}

func (s *DescribeLogAnalysisResponseBodyAnalysesAnalysisMetaValue) SetRoleArn(v string) *DescribeLogAnalysisResponseBodyAnalysesAnalysisMetaValue {
	s.RoleArn = &v
	return s
}

func (s *DescribeLogAnalysisResponseBodyAnalysesAnalysisMetaValue) Validate() error {
	return dara.Validate(s)
}
