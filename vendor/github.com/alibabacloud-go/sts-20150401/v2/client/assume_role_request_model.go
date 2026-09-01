// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssumeRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDurationSeconds(v int64) *AssumeRoleRequest
	GetDurationSeconds() *int64
	SetExternalId(v string) *AssumeRoleRequest
	GetExternalId() *string
	SetPolicy(v string) *AssumeRoleRequest
	GetPolicy() *string
	SetRoleArn(v string) *AssumeRoleRequest
	GetRoleArn() *string
	SetRoleSessionName(v string) *AssumeRoleRequest
	GetRoleSessionName() *string
	SetSourceIdentity(v string) *AssumeRoleRequest
	GetSourceIdentity() *string
}

type AssumeRoleRequest struct {
	// The validity period of the STS token. Unit: seconds.
	//
	// Minimum value: 900. Maximum value: the value of the `MaxSessionDuration` parameter. Default value: 3600.
	//
	// You can call the CreateRole or UpdateRole operation to configure the `MaxSessionDuration` parameter. For more information, see [CreateRole](https://help.aliyun.com/document_detail/28710.html) or [UpdateRole](https://help.aliyun.com/document_detail/28712.html).
	//
	// example:
	//
	// 3600
	DurationSeconds *int64 `json:"DurationSeconds,omitempty" xml:"DurationSeconds,omitempty"`
	// The external ID of the RAM role.
	//
	// This parameter is provided by an external party and is used to prevent the confused deputy problem. For more information, see [Use ExternalId to prevent the confused deputy problem](https://help.aliyun.com/document_detail/2361741.html).
	//
	// The value must be 2 to 1,224 characters in length and can contain letters, digits, and the following special characters: `= , . @ : / - _`. The regular expression for this parameter is `[\\w+=,.@:\\/-]*`.
	//
	// example:
	//
	// abcd1234
	ExternalId *string `json:"ExternalId,omitempty" xml:"ExternalId,omitempty"`
	// The policy that specifies the permissions of the returned STS token. You can use this parameter to grant the STS token fewer permissions than the permissions granted to the RAM role.
	//
	// 	- If you specify this parameter, the permissions of the returned STS token are the permissions that are included in the value of this parameter and owned by the RAM role.
	//
	// 	- If you do not specify this parameter, the returned STS token has all the permissions of the RAM role.
	//
	// The value must be 1 to 2,048 characters in length.
	//
	// For more information about policy elements and sample policies, see [Policy elements](https://help.aliyun.com/document_detail/93738.html) and [Overview of sample policies](https://help.aliyun.com/document_detail/210969.html).
	//
	// example:
	//
	// {"Statement": [{"Action": ["*"],"Effect": "Allow","Resource": ["*"]}],"Version":"1"}
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the RAM role.
	//
	// The trusted entity of the RAM role is an Alibaba Cloud account. For more information, see [Create a RAM role for a trusted Alibaba Cloud account](https://help.aliyun.com/document_detail/93691.html) or [CreateRole](https://help.aliyun.com/document_detail/28710.html).
	//
	// Format: `acs:ram::<account_id>:role/<role_name>`.
	//
	// You can view the ARN in the RAM console or by calling operations. The following items describe the validity periods of storage addresses:
	//
	// 	- For more information about how to view the ARN in the RAM console, see [How do I find the ARN of the RAM role?](https://help.aliyun.com/document_detail/39744.html)
	//
	// 	- For more information about how to view the ARN by calling operations, see [ListRoles](https://help.aliyun.com/document_detail/28713.html) or [GetRole](https://help.aliyun.com/document_detail/28711.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// acs:ram::123456789012****:role/adminrole
	RoleArn *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	// The custom name of the role session.
	//
	// Set this parameter based on your business requirements. In most cases, you can set this parameter to the identity of the API caller. For example, you can specify a username. You can specify `RoleSessionName` to identify API callers that assume the same RAM role in ActionTrail logs. This allows you to track the users that perform the operations.
	//
	// The value must be 2 to 64 characters in length and can contain letters, digits, and the following special characters: `. @ - _`.
	//
	// This parameter is required.
	//
	// example:
	//
	// alice
	RoleSessionName *string `json:"RoleSessionName,omitempty" xml:"RoleSessionName,omitempty"`
	SourceIdentity  *string `json:"SourceIdentity,omitempty" xml:"SourceIdentity,omitempty"`
}

func (s AssumeRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s AssumeRoleRequest) GoString() string {
	return s.String()
}

func (s *AssumeRoleRequest) GetDurationSeconds() *int64 {
	return s.DurationSeconds
}

func (s *AssumeRoleRequest) GetExternalId() *string {
	return s.ExternalId
}

func (s *AssumeRoleRequest) GetPolicy() *string {
	return s.Policy
}

func (s *AssumeRoleRequest) GetRoleArn() *string {
	return s.RoleArn
}

func (s *AssumeRoleRequest) GetRoleSessionName() *string {
	return s.RoleSessionName
}

func (s *AssumeRoleRequest) GetSourceIdentity() *string {
	return s.SourceIdentity
}

func (s *AssumeRoleRequest) SetDurationSeconds(v int64) *AssumeRoleRequest {
	s.DurationSeconds = &v
	return s
}

func (s *AssumeRoleRequest) SetExternalId(v string) *AssumeRoleRequest {
	s.ExternalId = &v
	return s
}

func (s *AssumeRoleRequest) SetPolicy(v string) *AssumeRoleRequest {
	s.Policy = &v
	return s
}

func (s *AssumeRoleRequest) SetRoleArn(v string) *AssumeRoleRequest {
	s.RoleArn = &v
	return s
}

func (s *AssumeRoleRequest) SetRoleSessionName(v string) *AssumeRoleRequest {
	s.RoleSessionName = &v
	return s
}

func (s *AssumeRoleRequest) SetSourceIdentity(v string) *AssumeRoleRequest {
	s.SourceIdentity = &v
	return s
}

func (s *AssumeRoleRequest) Validate() error {
	return dara.Validate(s)
}
