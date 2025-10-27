// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNfsAclResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAcl(v *DescribeNfsAclResponseBodyAcl) *DescribeNfsAclResponseBody
	GetAcl() *DescribeNfsAclResponseBodyAcl
	SetRequestId(v string) *DescribeNfsAclResponseBody
	GetRequestId() *string
}

type DescribeNfsAclResponseBody struct {
	// The information about the ACL feature.
	Acl *DescribeNfsAclResponseBodyAcl `json:"Acl,omitempty" xml:"Acl,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// A1098673-1746-505E-A5F1-08527B7EDBDF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeNfsAclResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeNfsAclResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNfsAclResponseBody) GetAcl() *DescribeNfsAclResponseBodyAcl {
	return s.Acl
}

func (s *DescribeNfsAclResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeNfsAclResponseBody) SetAcl(v *DescribeNfsAclResponseBodyAcl) *DescribeNfsAclResponseBody {
	s.Acl = v
	return s
}

func (s *DescribeNfsAclResponseBody) SetRequestId(v string) *DescribeNfsAclResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNfsAclResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeNfsAclResponseBodyAcl struct {
	// Indicates whether the NFS ACL feature is enabled.
	//
	// 	- true: The NFS ACL feature is enabled.
	//
	// 	- false: The NFS ACL feature is disabled.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
}

func (s DescribeNfsAclResponseBodyAcl) String() string {
	return dara.Prettify(s)
}

func (s DescribeNfsAclResponseBodyAcl) GoString() string {
	return s.String()
}

func (s *DescribeNfsAclResponseBodyAcl) GetEnabled() *bool {
	return s.Enabled
}

func (s *DescribeNfsAclResponseBodyAcl) SetEnabled(v bool) *DescribeNfsAclResponseBodyAcl {
	s.Enabled = &v
	return s
}

func (s *DescribeNfsAclResponseBodyAcl) Validate() error {
	return dara.Validate(s)
}
