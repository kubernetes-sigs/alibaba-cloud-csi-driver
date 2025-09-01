// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserClusterTypesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterTypes(v []*ListUserClusterTypesResponseBodyClusterTypes) *ListUserClusterTypesResponseBody
	GetClusterTypes() []*ListUserClusterTypesResponseBodyClusterTypes
	SetNextToken(v string) *ListUserClusterTypesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListUserClusterTypesResponseBody
	GetRequestId() *string
}

type ListUserClusterTypesResponseBody struct {
	// The list of cluster types. Number of elements in the array: 1 to 100.
	ClusterTypes []*ListUserClusterTypesResponseBodyClusterTypes `json:"ClusterTypes,omitempty" xml:"ClusterTypes,omitempty" type:"Repeated"`
	// NextToken for the next page. Include this value when requesting the next page.
	//
	// example:
	//
	// 3a6b93229825ac667104463b56790c91
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 4FD06DF0-9167-5C6F-A145-F30CA4A15D54
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListUserClusterTypesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUserClusterTypesResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserClusterTypesResponseBody) GetClusterTypes() []*ListUserClusterTypesResponseBodyClusterTypes {
	return s.ClusterTypes
}

func (s *ListUserClusterTypesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListUserClusterTypesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUserClusterTypesResponseBody) SetClusterTypes(v []*ListUserClusterTypesResponseBodyClusterTypes) *ListUserClusterTypesResponseBody {
	s.ClusterTypes = v
	return s
}

func (s *ListUserClusterTypesResponseBody) SetNextToken(v string) *ListUserClusterTypesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListUserClusterTypesResponseBody) SetRequestId(v string) *ListUserClusterTypesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserClusterTypesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListUserClusterTypesResponseBodyClusterTypes struct {
	// The access type of cluster type
	//
	// example:
	//
	// Public
	AccessType *string `json:"AccessType,omitempty" xml:"AccessType,omitempty"`
	// The name of cluster type
	//
	// example:
	//
	// AckEdgePro
	TypeName *string `json:"TypeName,omitempty" xml:"TypeName,omitempty"`
}

func (s ListUserClusterTypesResponseBodyClusterTypes) String() string {
	return dara.Prettify(s)
}

func (s ListUserClusterTypesResponseBodyClusterTypes) GoString() string {
	return s.String()
}

func (s *ListUserClusterTypesResponseBodyClusterTypes) GetAccessType() *string {
	return s.AccessType
}

func (s *ListUserClusterTypesResponseBodyClusterTypes) GetTypeName() *string {
	return s.TypeName
}

func (s *ListUserClusterTypesResponseBodyClusterTypes) SetAccessType(v string) *ListUserClusterTypesResponseBodyClusterTypes {
	s.AccessType = &v
	return s
}

func (s *ListUserClusterTypesResponseBodyClusterTypes) SetTypeName(v string) *ListUserClusterTypesResponseBodyClusterTypes {
	s.TypeName = &v
	return s
}

func (s *ListUserClusterTypesResponseBodyClusterTypes) Validate() error {
	return dara.Validate(s)
}
