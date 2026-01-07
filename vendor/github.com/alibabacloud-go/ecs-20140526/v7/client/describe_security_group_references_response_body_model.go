// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSecurityGroupReferencesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeSecurityGroupReferencesResponseBody
	GetRequestId() *string
	SetSecurityGroupReferences(v *DescribeSecurityGroupReferencesResponseBodySecurityGroupReferences) *DescribeSecurityGroupReferencesResponseBody
	GetSecurityGroupReferences() *DescribeSecurityGroupReferencesResponseBodySecurityGroupReferences
}

type DescribeSecurityGroupReferencesResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Details about the references to the specified security groups.
	SecurityGroupReferences *DescribeSecurityGroupReferencesResponseBodySecurityGroupReferences `json:"SecurityGroupReferences,omitempty" xml:"SecurityGroupReferences,omitempty" type:"Struct"`
}

func (s DescribeSecurityGroupReferencesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityGroupReferencesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSecurityGroupReferencesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSecurityGroupReferencesResponseBody) GetSecurityGroupReferences() *DescribeSecurityGroupReferencesResponseBodySecurityGroupReferences {
	return s.SecurityGroupReferences
}

func (s *DescribeSecurityGroupReferencesResponseBody) SetRequestId(v string) *DescribeSecurityGroupReferencesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSecurityGroupReferencesResponseBody) SetSecurityGroupReferences(v *DescribeSecurityGroupReferencesResponseBodySecurityGroupReferences) *DescribeSecurityGroupReferencesResponseBody {
	s.SecurityGroupReferences = v
	return s
}

func (s *DescribeSecurityGroupReferencesResponseBody) Validate() error {
	if s.SecurityGroupReferences != nil {
		if err := s.SecurityGroupReferences.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSecurityGroupReferencesResponseBodySecurityGroupReferences struct {
	SecurityGroupReference []*DescribeSecurityGroupReferencesResponseBodySecurityGroupReferencesSecurityGroupReference `json:"SecurityGroupReference,omitempty" xml:"SecurityGroupReference,omitempty" type:"Repeated"`
}

func (s DescribeSecurityGroupReferencesResponseBodySecurityGroupReferences) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityGroupReferencesResponseBodySecurityGroupReferences) GoString() string {
	return s.String()
}

func (s *DescribeSecurityGroupReferencesResponseBodySecurityGroupReferences) GetSecurityGroupReference() []*DescribeSecurityGroupReferencesResponseBodySecurityGroupReferencesSecurityGroupReference {
	return s.SecurityGroupReference
}

func (s *DescribeSecurityGroupReferencesResponseBodySecurityGroupReferences) SetSecurityGroupReference(v []*DescribeSecurityGroupReferencesResponseBodySecurityGroupReferencesSecurityGroupReference) *DescribeSecurityGroupReferencesResponseBodySecurityGroupReferences {
	s.SecurityGroupReference = v
	return s
}

func (s *DescribeSecurityGroupReferencesResponseBodySecurityGroupReferences) Validate() error {
	if s.SecurityGroupReference != nil {
		for _, item := range s.SecurityGroupReference {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSecurityGroupReferencesResponseBodySecurityGroupReferencesSecurityGroupReference struct {
	// Details about the security groups whose rules reference the specified security group.
	ReferencingSecurityGroups *DescribeSecurityGroupReferencesResponseBodySecurityGroupReferencesSecurityGroupReferenceReferencingSecurityGroups `json:"ReferencingSecurityGroups,omitempty" xml:"ReferencingSecurityGroups,omitempty" type:"Struct"`
	// The ID of the specified security group.
	//
	// example:
	//
	// sg-bp67acfmxazb4p****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
}

func (s DescribeSecurityGroupReferencesResponseBodySecurityGroupReferencesSecurityGroupReference) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityGroupReferencesResponseBodySecurityGroupReferencesSecurityGroupReference) GoString() string {
	return s.String()
}

func (s *DescribeSecurityGroupReferencesResponseBodySecurityGroupReferencesSecurityGroupReference) GetReferencingSecurityGroups() *DescribeSecurityGroupReferencesResponseBodySecurityGroupReferencesSecurityGroupReferenceReferencingSecurityGroups {
	return s.ReferencingSecurityGroups
}

func (s *DescribeSecurityGroupReferencesResponseBodySecurityGroupReferencesSecurityGroupReference) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *DescribeSecurityGroupReferencesResponseBodySecurityGroupReferencesSecurityGroupReference) SetReferencingSecurityGroups(v *DescribeSecurityGroupReferencesResponseBodySecurityGroupReferencesSecurityGroupReferenceReferencingSecurityGroups) *DescribeSecurityGroupReferencesResponseBodySecurityGroupReferencesSecurityGroupReference {
	s.ReferencingSecurityGroups = v
	return s
}

func (s *DescribeSecurityGroupReferencesResponseBodySecurityGroupReferencesSecurityGroupReference) SetSecurityGroupId(v string) *DescribeSecurityGroupReferencesResponseBodySecurityGroupReferencesSecurityGroupReference {
	s.SecurityGroupId = &v
	return s
}

func (s *DescribeSecurityGroupReferencesResponseBodySecurityGroupReferencesSecurityGroupReference) Validate() error {
	if s.ReferencingSecurityGroups != nil {
		if err := s.ReferencingSecurityGroups.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSecurityGroupReferencesResponseBodySecurityGroupReferencesSecurityGroupReferenceReferencingSecurityGroups struct {
	ReferencingSecurityGroup []*DescribeSecurityGroupReferencesResponseBodySecurityGroupReferencesSecurityGroupReferenceReferencingSecurityGroupsReferencingSecurityGroup `json:"ReferencingSecurityGroup,omitempty" xml:"ReferencingSecurityGroup,omitempty" type:"Repeated"`
}

func (s DescribeSecurityGroupReferencesResponseBodySecurityGroupReferencesSecurityGroupReferenceReferencingSecurityGroups) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityGroupReferencesResponseBodySecurityGroupReferencesSecurityGroupReferenceReferencingSecurityGroups) GoString() string {
	return s.String()
}

func (s *DescribeSecurityGroupReferencesResponseBodySecurityGroupReferencesSecurityGroupReferenceReferencingSecurityGroups) GetReferencingSecurityGroup() []*DescribeSecurityGroupReferencesResponseBodySecurityGroupReferencesSecurityGroupReferenceReferencingSecurityGroupsReferencingSecurityGroup {
	return s.ReferencingSecurityGroup
}

func (s *DescribeSecurityGroupReferencesResponseBodySecurityGroupReferencesSecurityGroupReferenceReferencingSecurityGroups) SetReferencingSecurityGroup(v []*DescribeSecurityGroupReferencesResponseBodySecurityGroupReferencesSecurityGroupReferenceReferencingSecurityGroupsReferencingSecurityGroup) *DescribeSecurityGroupReferencesResponseBodySecurityGroupReferencesSecurityGroupReferenceReferencingSecurityGroups {
	s.ReferencingSecurityGroup = v
	return s
}

func (s *DescribeSecurityGroupReferencesResponseBodySecurityGroupReferencesSecurityGroupReferenceReferencingSecurityGroups) Validate() error {
	if s.ReferencingSecurityGroup != nil {
		for _, item := range s.ReferencingSecurityGroup {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSecurityGroupReferencesResponseBodySecurityGroupReferencesSecurityGroupReferenceReferencingSecurityGroupsReferencingSecurityGroup struct {
	// The ID of the Alibaba Cloud account to which the security group whose rules reference the specified security group belongs.
	//
	// example:
	//
	// 123456****
	AliUid *string `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The ID of the security group whose rules reference the specified security group.
	//
	// example:
	//
	// sg-bp67acfmxazb4j****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
}

func (s DescribeSecurityGroupReferencesResponseBodySecurityGroupReferencesSecurityGroupReferenceReferencingSecurityGroupsReferencingSecurityGroup) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityGroupReferencesResponseBodySecurityGroupReferencesSecurityGroupReferenceReferencingSecurityGroupsReferencingSecurityGroup) GoString() string {
	return s.String()
}

func (s *DescribeSecurityGroupReferencesResponseBodySecurityGroupReferencesSecurityGroupReferenceReferencingSecurityGroupsReferencingSecurityGroup) GetAliUid() *string {
	return s.AliUid
}

func (s *DescribeSecurityGroupReferencesResponseBodySecurityGroupReferencesSecurityGroupReferenceReferencingSecurityGroupsReferencingSecurityGroup) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *DescribeSecurityGroupReferencesResponseBodySecurityGroupReferencesSecurityGroupReferenceReferencingSecurityGroupsReferencingSecurityGroup) SetAliUid(v string) *DescribeSecurityGroupReferencesResponseBodySecurityGroupReferencesSecurityGroupReferenceReferencingSecurityGroupsReferencingSecurityGroup {
	s.AliUid = &v
	return s
}

func (s *DescribeSecurityGroupReferencesResponseBodySecurityGroupReferencesSecurityGroupReferenceReferencingSecurityGroupsReferencingSecurityGroup) SetSecurityGroupId(v string) *DescribeSecurityGroupReferencesResponseBodySecurityGroupReferencesSecurityGroupReferenceReferencingSecurityGroupsReferencingSecurityGroup {
	s.SecurityGroupId = &v
	return s
}

func (s *DescribeSecurityGroupReferencesResponseBodySecurityGroupReferencesSecurityGroupReferenceReferencingSecurityGroupsReferencingSecurityGroup) Validate() error {
	return dara.Validate(s)
}
