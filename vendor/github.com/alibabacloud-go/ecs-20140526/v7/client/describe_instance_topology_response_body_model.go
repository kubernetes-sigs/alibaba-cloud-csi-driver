// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceTopologyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeInstanceTopologyResponseBody
	GetRequestId() *string
	SetTopologys(v *DescribeInstanceTopologyResponseBodyTopologys) *DescribeInstanceTopologyResponseBody
	GetTopologys() *DescribeInstanceTopologyResponseBodyTopologys
}

type DescribeInstanceTopologyResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Details about the topology.
	Topologys *DescribeInstanceTopologyResponseBodyTopologys `json:"Topologys,omitempty" xml:"Topologys,omitempty" type:"Struct"`
}

func (s DescribeInstanceTopologyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceTopologyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTopologyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceTopologyResponseBody) GetTopologys() *DescribeInstanceTopologyResponseBodyTopologys {
	return s.Topologys
}

func (s *DescribeInstanceTopologyResponseBody) SetRequestId(v string) *DescribeInstanceTopologyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceTopologyResponseBody) SetTopologys(v *DescribeInstanceTopologyResponseBodyTopologys) *DescribeInstanceTopologyResponseBody {
	s.Topologys = v
	return s
}

func (s *DescribeInstanceTopologyResponseBody) Validate() error {
	if s.Topologys != nil {
		if err := s.Topologys.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeInstanceTopologyResponseBodyTopologys struct {
	Topology []*DescribeInstanceTopologyResponseBodyTopologysTopology `json:"Topology,omitempty" xml:"Topology,omitempty" type:"Repeated"`
}

func (s DescribeInstanceTopologyResponseBodyTopologys) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceTopologyResponseBodyTopologys) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTopologyResponseBodyTopologys) GetTopology() []*DescribeInstanceTopologyResponseBodyTopologysTopology {
	return s.Topology
}

func (s *DescribeInstanceTopologyResponseBodyTopologys) SetTopology(v []*DescribeInstanceTopologyResponseBodyTopologysTopology) *DescribeInstanceTopologyResponseBodyTopologys {
	s.Topology = v
	return s
}

func (s *DescribeInstanceTopologyResponseBodyTopologys) Validate() error {
	if s.Topology != nil {
		for _, item := range s.Topology {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInstanceTopologyResponseBodyTopologysTopology struct {
	// The ID of the host where the ECS instance resides. This parameter is encrypted and cannot match the ID of the ECS instance. However, if the values of this parameter for different ECS instances are the same, the ECS instances reside on the same host.
	//
	// example:
	//
	// ZWNobyBo****
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// i-bp67acfmxazb4p****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeInstanceTopologyResponseBodyTopologysTopology) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceTopologyResponseBodyTopologysTopology) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTopologyResponseBodyTopologysTopology) GetHostId() *string {
	return s.HostId
}

func (s *DescribeInstanceTopologyResponseBodyTopologysTopology) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstanceTopologyResponseBodyTopologysTopology) SetHostId(v string) *DescribeInstanceTopologyResponseBodyTopologysTopology {
	s.HostId = &v
	return s
}

func (s *DescribeInstanceTopologyResponseBodyTopologysTopology) SetInstanceId(v string) *DescribeInstanceTopologyResponseBodyTopologysTopology {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceTopologyResponseBodyTopologysTopology) Validate() error {
	return dara.Validate(s)
}
