// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHpcClustersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHpcClusters(v *DescribeHpcClustersResponseBodyHpcClusters) *DescribeHpcClustersResponseBody
	GetHpcClusters() *DescribeHpcClustersResponseBodyHpcClusters
	SetPageNumber(v int32) *DescribeHpcClustersResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeHpcClustersResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeHpcClustersResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeHpcClustersResponseBody
	GetTotalCount() *int32
}

type DescribeHpcClustersResponseBody struct {
	// The name of the HPC cluster.
	HpcClusters *DescribeHpcClustersResponseBodyHpcClusters `json:"HpcClusters,omitempty" xml:"HpcClusters,omitempty" type:"Struct"`
	// Details about the HPC clusters. The value is an array that consists of the information of each HPC cluster.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The page number.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of HPC clusters.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the HPC cluster.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeHpcClustersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeHpcClustersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHpcClustersResponseBody) GetHpcClusters() *DescribeHpcClustersResponseBodyHpcClusters {
	return s.HpcClusters
}

func (s *DescribeHpcClustersResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeHpcClustersResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeHpcClustersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeHpcClustersResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeHpcClustersResponseBody) SetHpcClusters(v *DescribeHpcClustersResponseBodyHpcClusters) *DescribeHpcClustersResponseBody {
	s.HpcClusters = v
	return s
}

func (s *DescribeHpcClustersResponseBody) SetPageNumber(v int32) *DescribeHpcClustersResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeHpcClustersResponseBody) SetPageSize(v int32) *DescribeHpcClustersResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeHpcClustersResponseBody) SetRequestId(v string) *DescribeHpcClustersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHpcClustersResponseBody) SetTotalCount(v int32) *DescribeHpcClustersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeHpcClustersResponseBody) Validate() error {
	if s.HpcClusters != nil {
		if err := s.HpcClusters.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeHpcClustersResponseBodyHpcClusters struct {
	HpcCluster []*DescribeHpcClustersResponseBodyHpcClustersHpcCluster `json:"HpcCluster,omitempty" xml:"HpcCluster,omitempty" type:"Repeated"`
}

func (s DescribeHpcClustersResponseBodyHpcClusters) String() string {
	return dara.Prettify(s)
}

func (s DescribeHpcClustersResponseBodyHpcClusters) GoString() string {
	return s.String()
}

func (s *DescribeHpcClustersResponseBodyHpcClusters) GetHpcCluster() []*DescribeHpcClustersResponseBodyHpcClustersHpcCluster {
	return s.HpcCluster
}

func (s *DescribeHpcClustersResponseBodyHpcClusters) SetHpcCluster(v []*DescribeHpcClustersResponseBodyHpcClustersHpcCluster) *DescribeHpcClustersResponseBodyHpcClusters {
	s.HpcCluster = v
	return s
}

func (s *DescribeHpcClustersResponseBodyHpcClusters) Validate() error {
	if s.HpcCluster != nil {
		for _, item := range s.HpcCluster {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeHpcClustersResponseBodyHpcClustersHpcCluster struct {
	// The description of the HPC cluster.
	//
	// example:
	//
	// testDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The description of the HPC cluster.
	//
	// example:
	//
	// hpc-bp1a5zr3u7nq9cx****
	HpcClusterId *string `json:"HpcClusterId,omitempty" xml:"HpcClusterId,omitempty"`
	// The name of the HPC cluster.
	//
	// example:
	//
	// testName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeHpcClustersResponseBodyHpcClustersHpcCluster) String() string {
	return dara.Prettify(s)
}

func (s DescribeHpcClustersResponseBodyHpcClustersHpcCluster) GoString() string {
	return s.String()
}

func (s *DescribeHpcClustersResponseBodyHpcClustersHpcCluster) GetDescription() *string {
	return s.Description
}

func (s *DescribeHpcClustersResponseBodyHpcClustersHpcCluster) GetHpcClusterId() *string {
	return s.HpcClusterId
}

func (s *DescribeHpcClustersResponseBodyHpcClustersHpcCluster) GetName() *string {
	return s.Name
}

func (s *DescribeHpcClustersResponseBodyHpcClustersHpcCluster) SetDescription(v string) *DescribeHpcClustersResponseBodyHpcClustersHpcCluster {
	s.Description = &v
	return s
}

func (s *DescribeHpcClustersResponseBodyHpcClustersHpcCluster) SetHpcClusterId(v string) *DescribeHpcClustersResponseBodyHpcClustersHpcCluster {
	s.HpcClusterId = &v
	return s
}

func (s *DescribeHpcClustersResponseBodyHpcClustersHpcCluster) SetName(v string) *DescribeHpcClustersResponseBodyHpcClustersHpcCluster {
	s.Name = &v
	return s
}

func (s *DescribeHpcClustersResponseBodyHpcClustersHpcCluster) Validate() error {
	return dara.Validate(s)
}
