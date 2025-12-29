// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNetworkInterfacePermissionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNetworkInterfacePermissions(v *DescribeNetworkInterfacePermissionsResponseBodyNetworkInterfacePermissions) *DescribeNetworkInterfacePermissionsResponseBody
	GetNetworkInterfacePermissions() *DescribeNetworkInterfacePermissionsResponseBodyNetworkInterfacePermissions
	SetPageNumber(v int32) *DescribeNetworkInterfacePermissionsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeNetworkInterfacePermissionsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeNetworkInterfacePermissionsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeNetworkInterfacePermissionsResponseBody
	GetTotalCount() *int32
}

type DescribeNetworkInterfacePermissionsResponseBody struct {
	// Details about the ENI permissions.
	NetworkInterfacePermissions *DescribeNetworkInterfacePermissionsResponseBodyNetworkInterfacePermissions `json:"NetworkInterfacePermissions,omitempty" xml:"NetworkInterfacePermissions,omitempty" type:"Struct"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0FCD3DEF-63D3-4605-A818-805C8BD7DB87
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeNetworkInterfacePermissionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkInterfacePermissionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNetworkInterfacePermissionsResponseBody) GetNetworkInterfacePermissions() *DescribeNetworkInterfacePermissionsResponseBodyNetworkInterfacePermissions {
	return s.NetworkInterfacePermissions
}

func (s *DescribeNetworkInterfacePermissionsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeNetworkInterfacePermissionsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeNetworkInterfacePermissionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeNetworkInterfacePermissionsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeNetworkInterfacePermissionsResponseBody) SetNetworkInterfacePermissions(v *DescribeNetworkInterfacePermissionsResponseBodyNetworkInterfacePermissions) *DescribeNetworkInterfacePermissionsResponseBody {
	s.NetworkInterfacePermissions = v
	return s
}

func (s *DescribeNetworkInterfacePermissionsResponseBody) SetPageNumber(v int32) *DescribeNetworkInterfacePermissionsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeNetworkInterfacePermissionsResponseBody) SetPageSize(v int32) *DescribeNetworkInterfacePermissionsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeNetworkInterfacePermissionsResponseBody) SetRequestId(v string) *DescribeNetworkInterfacePermissionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNetworkInterfacePermissionsResponseBody) SetTotalCount(v int32) *DescribeNetworkInterfacePermissionsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeNetworkInterfacePermissionsResponseBody) Validate() error {
	if s.NetworkInterfacePermissions != nil {
		if err := s.NetworkInterfacePermissions.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeNetworkInterfacePermissionsResponseBodyNetworkInterfacePermissions struct {
	NetworkInterfacePermission []*DescribeNetworkInterfacePermissionsResponseBodyNetworkInterfacePermissionsNetworkInterfacePermission `json:"NetworkInterfacePermission,omitempty" xml:"NetworkInterfacePermission,omitempty" type:"Repeated"`
}

func (s DescribeNetworkInterfacePermissionsResponseBodyNetworkInterfacePermissions) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkInterfacePermissionsResponseBodyNetworkInterfacePermissions) GoString() string {
	return s.String()
}

func (s *DescribeNetworkInterfacePermissionsResponseBodyNetworkInterfacePermissions) GetNetworkInterfacePermission() []*DescribeNetworkInterfacePermissionsResponseBodyNetworkInterfacePermissionsNetworkInterfacePermission {
	return s.NetworkInterfacePermission
}

func (s *DescribeNetworkInterfacePermissionsResponseBodyNetworkInterfacePermissions) SetNetworkInterfacePermission(v []*DescribeNetworkInterfacePermissionsResponseBodyNetworkInterfacePermissionsNetworkInterfacePermission) *DescribeNetworkInterfacePermissionsResponseBodyNetworkInterfacePermissions {
	s.NetworkInterfacePermission = v
	return s
}

func (s *DescribeNetworkInterfacePermissionsResponseBodyNetworkInterfacePermissions) Validate() error {
	if s.NetworkInterfacePermission != nil {
		for _, item := range s.NetworkInterfacePermission {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeNetworkInterfacePermissionsResponseBodyNetworkInterfacePermissionsNetworkInterfacePermission struct {
	// The ID of the Alibaba Cloud partner (a certified ISV) or individual user.
	//
	// example:
	//
	// 1234567890
	AccountId *int64 `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The ID of ENI N.
	//
	// example:
	//
	// eni-bp14v2sdd3v8htln****
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
	// The ID of the ENI permission.
	//
	// example:
	//
	// eni-perm-bp1cs4lwn56lfb****
	NetworkInterfacePermissionId *string `json:"NetworkInterfacePermissionId,omitempty" xml:"NetworkInterfacePermissionId,omitempty"`
	// The ENI permission.
	//
	// example:
	//
	// InstanceAttach
	Permission *string `json:"Permission,omitempty" xml:"Permission,omitempty"`
	// The status of the ENI permission. Valid values:
	//
	// 	- Pending: The permission is being granted.
	//
	// 	- Granted: The permission is granted.
	//
	// 	- Revoking: The permission is being revoked.
	//
	// 	- Revoked: The permission is revoked.
	//
	// example:
	//
	// Granted
	PermissionState *string `json:"PermissionState,omitempty" xml:"PermissionState,omitempty"`
	// The name of the Alibaba Cloud service.
	//
	// example:
	//
	// Elastic Compute Service
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s DescribeNetworkInterfacePermissionsResponseBodyNetworkInterfacePermissionsNetworkInterfacePermission) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkInterfacePermissionsResponseBodyNetworkInterfacePermissionsNetworkInterfacePermission) GoString() string {
	return s.String()
}

func (s *DescribeNetworkInterfacePermissionsResponseBodyNetworkInterfacePermissionsNetworkInterfacePermission) GetAccountId() *int64 {
	return s.AccountId
}

func (s *DescribeNetworkInterfacePermissionsResponseBodyNetworkInterfacePermissionsNetworkInterfacePermission) GetNetworkInterfaceId() *string {
	return s.NetworkInterfaceId
}

func (s *DescribeNetworkInterfacePermissionsResponseBodyNetworkInterfacePermissionsNetworkInterfacePermission) GetNetworkInterfacePermissionId() *string {
	return s.NetworkInterfacePermissionId
}

func (s *DescribeNetworkInterfacePermissionsResponseBodyNetworkInterfacePermissionsNetworkInterfacePermission) GetPermission() *string {
	return s.Permission
}

func (s *DescribeNetworkInterfacePermissionsResponseBodyNetworkInterfacePermissionsNetworkInterfacePermission) GetPermissionState() *string {
	return s.PermissionState
}

func (s *DescribeNetworkInterfacePermissionsResponseBodyNetworkInterfacePermissionsNetworkInterfacePermission) GetServiceName() *string {
	return s.ServiceName
}

func (s *DescribeNetworkInterfacePermissionsResponseBodyNetworkInterfacePermissionsNetworkInterfacePermission) SetAccountId(v int64) *DescribeNetworkInterfacePermissionsResponseBodyNetworkInterfacePermissionsNetworkInterfacePermission {
	s.AccountId = &v
	return s
}

func (s *DescribeNetworkInterfacePermissionsResponseBodyNetworkInterfacePermissionsNetworkInterfacePermission) SetNetworkInterfaceId(v string) *DescribeNetworkInterfacePermissionsResponseBodyNetworkInterfacePermissionsNetworkInterfacePermission {
	s.NetworkInterfaceId = &v
	return s
}

func (s *DescribeNetworkInterfacePermissionsResponseBodyNetworkInterfacePermissionsNetworkInterfacePermission) SetNetworkInterfacePermissionId(v string) *DescribeNetworkInterfacePermissionsResponseBodyNetworkInterfacePermissionsNetworkInterfacePermission {
	s.NetworkInterfacePermissionId = &v
	return s
}

func (s *DescribeNetworkInterfacePermissionsResponseBodyNetworkInterfacePermissionsNetworkInterfacePermission) SetPermission(v string) *DescribeNetworkInterfacePermissionsResponseBodyNetworkInterfacePermissionsNetworkInterfacePermission {
	s.Permission = &v
	return s
}

func (s *DescribeNetworkInterfacePermissionsResponseBodyNetworkInterfacePermissionsNetworkInterfacePermission) SetPermissionState(v string) *DescribeNetworkInterfacePermissionsResponseBodyNetworkInterfacePermissionsNetworkInterfacePermission {
	s.PermissionState = &v
	return s
}

func (s *DescribeNetworkInterfacePermissionsResponseBodyNetworkInterfacePermissionsNetworkInterfacePermission) SetServiceName(v string) *DescribeNetworkInterfacePermissionsResponseBodyNetworkInterfacePermissionsNetworkInterfacePermission {
	s.ServiceName = &v
	return s
}

func (s *DescribeNetworkInterfacePermissionsResponseBodyNetworkInterfacePermissionsNetworkInterfacePermission) Validate() error {
	return dara.Validate(s)
}
