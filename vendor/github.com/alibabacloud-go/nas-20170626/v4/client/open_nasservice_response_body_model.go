// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenNASServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *OpenNASServiceResponseBody
	GetAccessDeniedDetail() *string
	SetOrderId(v string) *OpenNASServiceResponseBody
	GetOrderId() *string
	SetRequestId(v string) *OpenNASServiceResponseBody
	GetRequestId() *string
}

type OpenNASServiceResponseBody struct {
	// The details about the failed permission verification.
	//
	// example:
	//
	// {
	//
	//     "PolicyType": "AccountLevelIdentityBasedPolicy",
	//
	//     "AuthPrincipalOwnerId": "178321033379****",
	//
	//     "EncodedDiagnosticMessage": "AJpt/382mjxDSIYIqa/cUIFvOg9tajlLyN+LJA0C78kWfKIl****",
	//
	//     "AuthPrincipalType": "SubUser",
	//
	//     "AuthPrincipalDisplayName": "21794847602038****",
	//
	//     "NoPermissionType": "ImplicitDeny",
	//
	//     "AuthAction": "nas:OpenNASService"
	//
	//   }
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The order ID.
	//
	// example:
	//
	// 20671870151****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 427DB0B3-9436-4A3C-B2BC-B961F95E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OpenNASServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OpenNASServiceResponseBody) GoString() string {
	return s.String()
}

func (s *OpenNASServiceResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *OpenNASServiceResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *OpenNASServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OpenNASServiceResponseBody) SetAccessDeniedDetail(v string) *OpenNASServiceResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *OpenNASServiceResponseBody) SetOrderId(v string) *OpenNASServiceResponseBody {
	s.OrderId = &v
	return s
}

func (s *OpenNASServiceResponseBody) SetRequestId(v string) *OpenNASServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *OpenNASServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
