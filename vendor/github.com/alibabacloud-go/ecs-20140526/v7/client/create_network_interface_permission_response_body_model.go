// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNetworkInterfacePermissionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNetworkInterfacePermission(v *CreateNetworkInterfacePermissionResponseBodyNetworkInterfacePermission) *CreateNetworkInterfacePermissionResponseBody
	GetNetworkInterfacePermission() *CreateNetworkInterfacePermissionResponseBodyNetworkInterfacePermission
	SetRequestId(v string) *CreateNetworkInterfacePermissionResponseBody
	GetRequestId() *string
}

type CreateNetworkInterfacePermissionResponseBody struct {
	// Details about permissions on the ENI.
	NetworkInterfacePermission *CreateNetworkInterfacePermissionResponseBodyNetworkInterfacePermission `json:"NetworkInterfacePermission,omitempty" xml:"NetworkInterfacePermission,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 0FCD3DEF-63D3-4605-A818-805C8BD7DB87
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateNetworkInterfacePermissionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateNetworkInterfacePermissionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateNetworkInterfacePermissionResponseBody) GetNetworkInterfacePermission() *CreateNetworkInterfacePermissionResponseBodyNetworkInterfacePermission {
	return s.NetworkInterfacePermission
}

func (s *CreateNetworkInterfacePermissionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateNetworkInterfacePermissionResponseBody) SetNetworkInterfacePermission(v *CreateNetworkInterfacePermissionResponseBodyNetworkInterfacePermission) *CreateNetworkInterfacePermissionResponseBody {
	s.NetworkInterfacePermission = v
	return s
}

func (s *CreateNetworkInterfacePermissionResponseBody) SetRequestId(v string) *CreateNetworkInterfacePermissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateNetworkInterfacePermissionResponseBody) Validate() error {
	if s.NetworkInterfacePermission != nil {
		if err := s.NetworkInterfacePermission.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateNetworkInterfacePermissionResponseBodyNetworkInterfacePermission struct {
	// The ID of the Alibaba Cloud partner (a certified ISV).
	//
	// example:
	//
	// 1234567890
	AccountId *int64 `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The ID of the ENI.
	//
	// example:
	//
	// eni-bp14v2sdd3v8htln****
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
	// The ID of the permission on the ENI.
	//
	// example:
	//
	// eni-perm-bp1cs4lwn56lfb****
	NetworkInterfacePermissionId *string `json:"NetworkInterfacePermissionId,omitempty" xml:"NetworkInterfacePermissionId,omitempty"`
	// The permission on the ENI.
	//
	// example:
	//
	// InstanceAttach
	Permission *string `json:"Permission,omitempty" xml:"Permission,omitempty"`
	// The state of the permission on the ENI. Valid values:
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

func (s CreateNetworkInterfacePermissionResponseBodyNetworkInterfacePermission) String() string {
	return dara.Prettify(s)
}

func (s CreateNetworkInterfacePermissionResponseBodyNetworkInterfacePermission) GoString() string {
	return s.String()
}

func (s *CreateNetworkInterfacePermissionResponseBodyNetworkInterfacePermission) GetAccountId() *int64 {
	return s.AccountId
}

func (s *CreateNetworkInterfacePermissionResponseBodyNetworkInterfacePermission) GetNetworkInterfaceId() *string {
	return s.NetworkInterfaceId
}

func (s *CreateNetworkInterfacePermissionResponseBodyNetworkInterfacePermission) GetNetworkInterfacePermissionId() *string {
	return s.NetworkInterfacePermissionId
}

func (s *CreateNetworkInterfacePermissionResponseBodyNetworkInterfacePermission) GetPermission() *string {
	return s.Permission
}

func (s *CreateNetworkInterfacePermissionResponseBodyNetworkInterfacePermission) GetPermissionState() *string {
	return s.PermissionState
}

func (s *CreateNetworkInterfacePermissionResponseBodyNetworkInterfacePermission) GetServiceName() *string {
	return s.ServiceName
}

func (s *CreateNetworkInterfacePermissionResponseBodyNetworkInterfacePermission) SetAccountId(v int64) *CreateNetworkInterfacePermissionResponseBodyNetworkInterfacePermission {
	s.AccountId = &v
	return s
}

func (s *CreateNetworkInterfacePermissionResponseBodyNetworkInterfacePermission) SetNetworkInterfaceId(v string) *CreateNetworkInterfacePermissionResponseBodyNetworkInterfacePermission {
	s.NetworkInterfaceId = &v
	return s
}

func (s *CreateNetworkInterfacePermissionResponseBodyNetworkInterfacePermission) SetNetworkInterfacePermissionId(v string) *CreateNetworkInterfacePermissionResponseBodyNetworkInterfacePermission {
	s.NetworkInterfacePermissionId = &v
	return s
}

func (s *CreateNetworkInterfacePermissionResponseBodyNetworkInterfacePermission) SetPermission(v string) *CreateNetworkInterfacePermissionResponseBodyNetworkInterfacePermission {
	s.Permission = &v
	return s
}

func (s *CreateNetworkInterfacePermissionResponseBodyNetworkInterfacePermission) SetPermissionState(v string) *CreateNetworkInterfacePermissionResponseBodyNetworkInterfacePermission {
	s.PermissionState = &v
	return s
}

func (s *CreateNetworkInterfacePermissionResponseBodyNetworkInterfacePermission) SetServiceName(v string) *CreateNetworkInterfacePermissionResponseBodyNetworkInterfacePermission {
	s.ServiceName = &v
	return s
}

func (s *CreateNetworkInterfacePermissionResponseBodyNetworkInterfacePermission) Validate() error {
	return dara.Validate(s)
}
