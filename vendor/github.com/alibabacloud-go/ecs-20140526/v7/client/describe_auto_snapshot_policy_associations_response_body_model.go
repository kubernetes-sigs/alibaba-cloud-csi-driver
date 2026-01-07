// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAutoSnapshotPolicyAssociationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAutoSnapshotPolicyAssociations(v *DescribeAutoSnapshotPolicyAssociationsResponseBodyAutoSnapshotPolicyAssociations) *DescribeAutoSnapshotPolicyAssociationsResponseBody
	GetAutoSnapshotPolicyAssociations() *DescribeAutoSnapshotPolicyAssociationsResponseBodyAutoSnapshotPolicyAssociations
	SetNextToken(v string) *DescribeAutoSnapshotPolicyAssociationsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeAutoSnapshotPolicyAssociationsResponseBody
	GetRequestId() *string
}

type DescribeAutoSnapshotPolicyAssociationsResponseBody struct {
	// The association of automatic snapshot policies.
	AutoSnapshotPolicyAssociations *DescribeAutoSnapshotPolicyAssociationsResponseBodyAutoSnapshotPolicyAssociations `json:"AutoSnapshotPolicyAssociations,omitempty" xml:"AutoSnapshotPolicyAssociations,omitempty" type:"Struct"`
	// The returned pagination token which can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAutoSnapshotPolicyAssociationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoSnapshotPolicyAssociationsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAutoSnapshotPolicyAssociationsResponseBody) GetAutoSnapshotPolicyAssociations() *DescribeAutoSnapshotPolicyAssociationsResponseBodyAutoSnapshotPolicyAssociations {
	return s.AutoSnapshotPolicyAssociations
}

func (s *DescribeAutoSnapshotPolicyAssociationsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeAutoSnapshotPolicyAssociationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAutoSnapshotPolicyAssociationsResponseBody) SetAutoSnapshotPolicyAssociations(v *DescribeAutoSnapshotPolicyAssociationsResponseBodyAutoSnapshotPolicyAssociations) *DescribeAutoSnapshotPolicyAssociationsResponseBody {
	s.AutoSnapshotPolicyAssociations = v
	return s
}

func (s *DescribeAutoSnapshotPolicyAssociationsResponseBody) SetNextToken(v string) *DescribeAutoSnapshotPolicyAssociationsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeAutoSnapshotPolicyAssociationsResponseBody) SetRequestId(v string) *DescribeAutoSnapshotPolicyAssociationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAutoSnapshotPolicyAssociationsResponseBody) Validate() error {
	if s.AutoSnapshotPolicyAssociations != nil {
		if err := s.AutoSnapshotPolicyAssociations.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAutoSnapshotPolicyAssociationsResponseBodyAutoSnapshotPolicyAssociations struct {
	AutoSnapshotPolicyAssociation []*DescribeAutoSnapshotPolicyAssociationsResponseBodyAutoSnapshotPolicyAssociationsAutoSnapshotPolicyAssociation `json:"AutoSnapshotPolicyAssociation,omitempty" xml:"AutoSnapshotPolicyAssociation,omitempty" type:"Repeated"`
}

func (s DescribeAutoSnapshotPolicyAssociationsResponseBodyAutoSnapshotPolicyAssociations) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoSnapshotPolicyAssociationsResponseBodyAutoSnapshotPolicyAssociations) GoString() string {
	return s.String()
}

func (s *DescribeAutoSnapshotPolicyAssociationsResponseBodyAutoSnapshotPolicyAssociations) GetAutoSnapshotPolicyAssociation() []*DescribeAutoSnapshotPolicyAssociationsResponseBodyAutoSnapshotPolicyAssociationsAutoSnapshotPolicyAssociation {
	return s.AutoSnapshotPolicyAssociation
}

func (s *DescribeAutoSnapshotPolicyAssociationsResponseBodyAutoSnapshotPolicyAssociations) SetAutoSnapshotPolicyAssociation(v []*DescribeAutoSnapshotPolicyAssociationsResponseBodyAutoSnapshotPolicyAssociationsAutoSnapshotPolicyAssociation) *DescribeAutoSnapshotPolicyAssociationsResponseBodyAutoSnapshotPolicyAssociations {
	s.AutoSnapshotPolicyAssociation = v
	return s
}

func (s *DescribeAutoSnapshotPolicyAssociationsResponseBodyAutoSnapshotPolicyAssociations) Validate() error {
	if s.AutoSnapshotPolicyAssociation != nil {
		for _, item := range s.AutoSnapshotPolicyAssociation {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAutoSnapshotPolicyAssociationsResponseBodyAutoSnapshotPolicyAssociationsAutoSnapshotPolicyAssociation struct {
	// The ID of the automatic snapshot policy.
	//
	// example:
	//
	// sp-bp12quk7gqhhuu1f****
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" xml:"AutoSnapshotPolicyId,omitempty"`
	// The ID of the cloud disk.
	//
	// example:
	//
	// d-bp67acfmxazb4p****
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
}

func (s DescribeAutoSnapshotPolicyAssociationsResponseBodyAutoSnapshotPolicyAssociationsAutoSnapshotPolicyAssociation) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoSnapshotPolicyAssociationsResponseBodyAutoSnapshotPolicyAssociationsAutoSnapshotPolicyAssociation) GoString() string {
	return s.String()
}

func (s *DescribeAutoSnapshotPolicyAssociationsResponseBodyAutoSnapshotPolicyAssociationsAutoSnapshotPolicyAssociation) GetAutoSnapshotPolicyId() *string {
	return s.AutoSnapshotPolicyId
}

func (s *DescribeAutoSnapshotPolicyAssociationsResponseBodyAutoSnapshotPolicyAssociationsAutoSnapshotPolicyAssociation) GetDiskId() *string {
	return s.DiskId
}

func (s *DescribeAutoSnapshotPolicyAssociationsResponseBodyAutoSnapshotPolicyAssociationsAutoSnapshotPolicyAssociation) SetAutoSnapshotPolicyId(v string) *DescribeAutoSnapshotPolicyAssociationsResponseBodyAutoSnapshotPolicyAssociationsAutoSnapshotPolicyAssociation {
	s.AutoSnapshotPolicyId = &v
	return s
}

func (s *DescribeAutoSnapshotPolicyAssociationsResponseBodyAutoSnapshotPolicyAssociationsAutoSnapshotPolicyAssociation) SetDiskId(v string) *DescribeAutoSnapshotPolicyAssociationsResponseBodyAutoSnapshotPolicyAssociationsAutoSnapshotPolicyAssociation {
	s.DiskId = &v
	return s
}

func (s *DescribeAutoSnapshotPolicyAssociationsResponseBodyAutoSnapshotPolicyAssociationsAutoSnapshotPolicyAssociation) Validate() error {
	return dara.Validate(s)
}
