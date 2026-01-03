// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClustersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusters(v *DescribeClustersResponseBodyClusters) *DescribeClustersResponseBody
	GetClusters() *DescribeClustersResponseBodyClusters
	SetRequestId(v string) *DescribeClustersResponseBody
	GetRequestId() *string
}

type DescribeClustersResponseBody struct {
	Clusters  *DescribeClustersResponseBodyClusters `json:"Clusters,omitempty" xml:"Clusters,omitempty" type:"Struct"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeClustersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeClustersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClustersResponseBody) GetClusters() *DescribeClustersResponseBodyClusters {
	return s.Clusters
}

func (s *DescribeClustersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeClustersResponseBody) SetClusters(v *DescribeClustersResponseBodyClusters) *DescribeClustersResponseBody {
	s.Clusters = v
	return s
}

func (s *DescribeClustersResponseBody) SetRequestId(v string) *DescribeClustersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeClustersResponseBody) Validate() error {
	if s.Clusters != nil {
		if err := s.Clusters.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeClustersResponseBodyClusters struct {
	Cluster []*DescribeClustersResponseBodyClustersCluster `json:"Cluster,omitempty" xml:"Cluster,omitempty" type:"Repeated"`
}

func (s DescribeClustersResponseBodyClusters) String() string {
	return dara.Prettify(s)
}

func (s DescribeClustersResponseBodyClusters) GoString() string {
	return s.String()
}

func (s *DescribeClustersResponseBodyClusters) GetCluster() []*DescribeClustersResponseBodyClustersCluster {
	return s.Cluster
}

func (s *DescribeClustersResponseBodyClusters) SetCluster(v []*DescribeClustersResponseBodyClustersCluster) *DescribeClustersResponseBodyClusters {
	s.Cluster = v
	return s
}

func (s *DescribeClustersResponseBodyClusters) Validate() error {
	if s.Cluster != nil {
		for _, item := range s.Cluster {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeClustersResponseBodyClustersCluster struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DescribeClustersResponseBodyClustersCluster) String() string {
	return dara.Prettify(s)
}

func (s DescribeClustersResponseBodyClustersCluster) GoString() string {
	return s.String()
}

func (s *DescribeClustersResponseBodyClustersCluster) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeClustersResponseBodyClustersCluster) SetClusterId(v string) *DescribeClustersResponseBodyClustersCluster {
	s.ClusterId = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersCluster) Validate() error {
	return dara.Validate(s)
}
