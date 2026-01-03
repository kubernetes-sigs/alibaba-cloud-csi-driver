// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePrefixListAssociationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *DescribePrefixListAssociationsResponseBody
	GetNextToken() *string
	SetPrefixListAssociations(v *DescribePrefixListAssociationsResponseBodyPrefixListAssociations) *DescribePrefixListAssociationsResponseBody
	GetPrefixListAssociations() *DescribePrefixListAssociationsResponseBodyPrefixListAssociations
	SetRequestId(v string) *DescribePrefixListAssociationsResponseBody
	GetRequestId() *string
}

type DescribePrefixListAssociationsResponseBody struct {
	// The query token that is returned in this call. If the return value is empty, no more data is returned.
	//
	// example:
	//
	// AAAAAdDWBF2****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Details about the resources that are associated with the prefix list.
	PrefixListAssociations *DescribePrefixListAssociationsResponseBodyPrefixListAssociations `json:"PrefixListAssociations,omitempty" xml:"PrefixListAssociations,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 38793DB8-A4B2-4AEC-BFD3-111234E9188D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePrefixListAssociationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePrefixListAssociationsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePrefixListAssociationsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribePrefixListAssociationsResponseBody) GetPrefixListAssociations() *DescribePrefixListAssociationsResponseBodyPrefixListAssociations {
	return s.PrefixListAssociations
}

func (s *DescribePrefixListAssociationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePrefixListAssociationsResponseBody) SetNextToken(v string) *DescribePrefixListAssociationsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribePrefixListAssociationsResponseBody) SetPrefixListAssociations(v *DescribePrefixListAssociationsResponseBodyPrefixListAssociations) *DescribePrefixListAssociationsResponseBody {
	s.PrefixListAssociations = v
	return s
}

func (s *DescribePrefixListAssociationsResponseBody) SetRequestId(v string) *DescribePrefixListAssociationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePrefixListAssociationsResponseBody) Validate() error {
	if s.PrefixListAssociations != nil {
		if err := s.PrefixListAssociations.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribePrefixListAssociationsResponseBodyPrefixListAssociations struct {
	PrefixListAssociation []*DescribePrefixListAssociationsResponseBodyPrefixListAssociationsPrefixListAssociation `json:"PrefixListAssociation,omitempty" xml:"PrefixListAssociation,omitempty" type:"Repeated"`
}

func (s DescribePrefixListAssociationsResponseBodyPrefixListAssociations) String() string {
	return dara.Prettify(s)
}

func (s DescribePrefixListAssociationsResponseBodyPrefixListAssociations) GoString() string {
	return s.String()
}

func (s *DescribePrefixListAssociationsResponseBodyPrefixListAssociations) GetPrefixListAssociation() []*DescribePrefixListAssociationsResponseBodyPrefixListAssociationsPrefixListAssociation {
	return s.PrefixListAssociation
}

func (s *DescribePrefixListAssociationsResponseBodyPrefixListAssociations) SetPrefixListAssociation(v []*DescribePrefixListAssociationsResponseBodyPrefixListAssociationsPrefixListAssociation) *DescribePrefixListAssociationsResponseBodyPrefixListAssociations {
	s.PrefixListAssociation = v
	return s
}

func (s *DescribePrefixListAssociationsResponseBodyPrefixListAssociations) Validate() error {
	if s.PrefixListAssociation != nil {
		for _, item := range s.PrefixListAssociation {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePrefixListAssociationsResponseBodyPrefixListAssociationsPrefixListAssociation struct {
	// The ID of the resource.
	//
	// example:
	//
	// sg-bp11ujym6xsff6l0****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of the resource.
	//
	// example:
	//
	// securitygroup
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s DescribePrefixListAssociationsResponseBodyPrefixListAssociationsPrefixListAssociation) String() string {
	return dara.Prettify(s)
}

func (s DescribePrefixListAssociationsResponseBodyPrefixListAssociationsPrefixListAssociation) GoString() string {
	return s.String()
}

func (s *DescribePrefixListAssociationsResponseBodyPrefixListAssociationsPrefixListAssociation) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribePrefixListAssociationsResponseBodyPrefixListAssociationsPrefixListAssociation) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribePrefixListAssociationsResponseBodyPrefixListAssociationsPrefixListAssociation) SetResourceId(v string) *DescribePrefixListAssociationsResponseBodyPrefixListAssociationsPrefixListAssociation {
	s.ResourceId = &v
	return s
}

func (s *DescribePrefixListAssociationsResponseBodyPrefixListAssociationsPrefixListAssociation) SetResourceType(v string) *DescribePrefixListAssociationsResponseBodyPrefixListAssociationsPrefixListAssociation {
	s.ResourceType = &v
	return s
}

func (s *DescribePrefixListAssociationsResponseBodyPrefixListAssociationsPrefixListAssociation) Validate() error {
	return dara.Validate(s)
}
