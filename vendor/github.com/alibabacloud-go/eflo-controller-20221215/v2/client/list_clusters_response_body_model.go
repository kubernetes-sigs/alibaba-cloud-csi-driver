// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClustersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusters(v []*ListClustersResponseBodyClusters) *ListClustersResponseBody
	GetClusters() []*ListClustersResponseBodyClusters
	SetNextToken(v string) *ListClustersResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListClustersResponseBody
	GetRequestId() *string
}

type ListClustersResponseBody struct {
	// The clusters.
	Clusters []*ListClustersResponseBodyClusters `json:"Clusters,omitempty" xml:"Clusters,omitempty" type:"Repeated"`
	// The returned pagination token which can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// f4f9a292c17072a2
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2FE2B22C-CF9D-59DE-BF63-DC9B9B33A9D1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListClustersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListClustersResponseBody) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBody) GetClusters() []*ListClustersResponseBodyClusters {
	return s.Clusters
}

func (s *ListClustersResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListClustersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListClustersResponseBody) SetClusters(v []*ListClustersResponseBodyClusters) *ListClustersResponseBody {
	s.Clusters = v
	return s
}

func (s *ListClustersResponseBody) SetNextToken(v string) *ListClustersResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListClustersResponseBody) SetRequestId(v string) *ListClustersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListClustersResponseBody) Validate() error {
	if s.Clusters != nil {
		for _, item := range s.Clusters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListClustersResponseBodyClusters struct {
	// The cluster description.
	//
	// example:
	//
	// PPU-cluster2 bz
	ClusterDescription *string `json:"ClusterDescription,omitempty" xml:"ClusterDescription,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// i137590131672134915401
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The cluster name.
	//
	// example:
	//
	// cnp_test_cluster
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The cluster type.
	//
	// Valid values:
	//
	// 	- AckEdgePro
	//
	// 	- ExclusiveBareCluster
	//
	// 	- Lite
	//
	// example:
	//
	// AckEdgePro
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// The component information.
	//
	// example:
	//
	// {}
	Components interface{} `json:"Components,omitempty" xml:"Components,omitempty"`
	// The IP type of the computing network.
	//
	// example:
	//
	// IPv4
	ComputingIpVersion *string `json:"ComputingIpVersion,omitempty" xml:"ComputingIpVersion,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 1672134938
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The cluster number.
	//
	// example:
	//
	// B1
	HpnZone *string `json:"HpnZone,omitempty" xml:"HpnZone,omitempty"`
	// The number of nodes.
	//
	// example:
	//
	// 12
	NodeCount *int64 `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	// The number of node groups.
	//
	// example:
	//
	// 2
	NodeGroupCount *int64 `json:"NodeGroupCount,omitempty" xml:"NodeGroupCount,omitempty"`
	// The cluster status.
	//
	// Valid values:
	//
	// 	- running
	//
	// 	- expanding
	//
	// 	- shrinking
	//
	// 	- initializing
	//
	// example:
	//
	// initializing
	OperatingState *string `json:"OperatingState,omitempty" xml:"OperatingState,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-aek2ajbjoloa23q
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags.
	Tags []*ListClustersResponseBodyClustersTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The job ID.
	//
	// example:
	//
	// i156365121663149566024
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The update time.
	//
	// example:
	//
	// 1672134968
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The virtual private cloud (VPC) ID.
	//
	// example:
	//
	// vpc-0jlx4hol2bjboafzmffvd
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListClustersResponseBodyClusters) String() string {
	return dara.Prettify(s)
}

func (s ListClustersResponseBodyClusters) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBodyClusters) GetClusterDescription() *string {
	return s.ClusterDescription
}

func (s *ListClustersResponseBodyClusters) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListClustersResponseBodyClusters) GetClusterName() *string {
	return s.ClusterName
}

func (s *ListClustersResponseBodyClusters) GetClusterType() *string {
	return s.ClusterType
}

func (s *ListClustersResponseBodyClusters) GetComponents() interface{} {
	return s.Components
}

func (s *ListClustersResponseBodyClusters) GetComputingIpVersion() *string {
	return s.ComputingIpVersion
}

func (s *ListClustersResponseBodyClusters) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListClustersResponseBodyClusters) GetHpnZone() *string {
	return s.HpnZone
}

func (s *ListClustersResponseBodyClusters) GetNodeCount() *int64 {
	return s.NodeCount
}

func (s *ListClustersResponseBodyClusters) GetNodeGroupCount() *int64 {
	return s.NodeGroupCount
}

func (s *ListClustersResponseBodyClusters) GetOperatingState() *string {
	return s.OperatingState
}

func (s *ListClustersResponseBodyClusters) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListClustersResponseBodyClusters) GetTags() []*ListClustersResponseBodyClustersTags {
	return s.Tags
}

func (s *ListClustersResponseBodyClusters) GetTaskId() *string {
	return s.TaskId
}

func (s *ListClustersResponseBodyClusters) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListClustersResponseBodyClusters) GetVpcId() *string {
	return s.VpcId
}

func (s *ListClustersResponseBodyClusters) SetClusterDescription(v string) *ListClustersResponseBodyClusters {
	s.ClusterDescription = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetClusterId(v string) *ListClustersResponseBodyClusters {
	s.ClusterId = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetClusterName(v string) *ListClustersResponseBodyClusters {
	s.ClusterName = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetClusterType(v string) *ListClustersResponseBodyClusters {
	s.ClusterType = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetComponents(v interface{}) *ListClustersResponseBodyClusters {
	s.Components = v
	return s
}

func (s *ListClustersResponseBodyClusters) SetComputingIpVersion(v string) *ListClustersResponseBodyClusters {
	s.ComputingIpVersion = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetCreateTime(v string) *ListClustersResponseBodyClusters {
	s.CreateTime = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetHpnZone(v string) *ListClustersResponseBodyClusters {
	s.HpnZone = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetNodeCount(v int64) *ListClustersResponseBodyClusters {
	s.NodeCount = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetNodeGroupCount(v int64) *ListClustersResponseBodyClusters {
	s.NodeGroupCount = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetOperatingState(v string) *ListClustersResponseBodyClusters {
	s.OperatingState = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetResourceGroupId(v string) *ListClustersResponseBodyClusters {
	s.ResourceGroupId = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetTags(v []*ListClustersResponseBodyClustersTags) *ListClustersResponseBodyClusters {
	s.Tags = v
	return s
}

func (s *ListClustersResponseBodyClusters) SetTaskId(v string) *ListClustersResponseBodyClusters {
	s.TaskId = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetUpdateTime(v string) *ListClustersResponseBodyClusters {
	s.UpdateTime = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetVpcId(v string) *ListClustersResponseBodyClusters {
	s.VpcId = &v
	return s
}

func (s *ListClustersResponseBodyClusters) Validate() error {
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

type ListClustersResponseBodyClustersTags struct {
	// The tag key.
	//
	// example:
	//
	// aa_key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// aa_value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListClustersResponseBodyClustersTags) String() string {
	return dara.Prettify(s)
}

func (s ListClustersResponseBodyClustersTags) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBodyClustersTags) GetKey() *string {
	return s.Key
}

func (s *ListClustersResponseBodyClustersTags) GetValue() *string {
	return s.Value
}

func (s *ListClustersResponseBodyClustersTags) SetKey(v string) *ListClustersResponseBodyClustersTags {
	s.Key = &v
	return s
}

func (s *ListClustersResponseBodyClustersTags) SetValue(v string) *ListClustersResponseBodyClustersTags {
	s.Value = &v
	return s
}

func (s *ListClustersResponseBodyClustersTags) Validate() error {
	return dara.Validate(s)
}
