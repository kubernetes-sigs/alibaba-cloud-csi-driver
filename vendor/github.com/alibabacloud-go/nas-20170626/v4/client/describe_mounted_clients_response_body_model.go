// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMountedClientsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClients(v *DescribeMountedClientsResponseBodyClients) *DescribeMountedClientsResponseBody
	GetClients() *DescribeMountedClientsResponseBodyClients
	SetPageNumber(v int32) *DescribeMountedClientsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeMountedClientsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeMountedClientsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeMountedClientsResponseBody
	GetTotalCount() *int32
}

type DescribeMountedClientsResponseBody struct {
	// The queried clients.
	Clients *DescribeMountedClientsResponseBodyClients `json:"Clients,omitempty" xml:"Clients,omitempty" type:"Struct"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of IP addresses returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A70BEE5D-76D3-49FB-B58F-1F398211****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of IP addresses.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeMountedClientsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMountedClientsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMountedClientsResponseBody) GetClients() *DescribeMountedClientsResponseBodyClients {
	return s.Clients
}

func (s *DescribeMountedClientsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeMountedClientsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeMountedClientsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMountedClientsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeMountedClientsResponseBody) SetClients(v *DescribeMountedClientsResponseBodyClients) *DescribeMountedClientsResponseBody {
	s.Clients = v
	return s
}

func (s *DescribeMountedClientsResponseBody) SetPageNumber(v int32) *DescribeMountedClientsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeMountedClientsResponseBody) SetPageSize(v int32) *DescribeMountedClientsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeMountedClientsResponseBody) SetRequestId(v string) *DescribeMountedClientsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMountedClientsResponseBody) SetTotalCount(v int32) *DescribeMountedClientsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeMountedClientsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeMountedClientsResponseBodyClients struct {
	Client []*DescribeMountedClientsResponseBodyClientsClient `json:"Client,omitempty" xml:"Client,omitempty" type:"Repeated"`
}

func (s DescribeMountedClientsResponseBodyClients) String() string {
	return dara.Prettify(s)
}

func (s DescribeMountedClientsResponseBodyClients) GoString() string {
	return s.String()
}

func (s *DescribeMountedClientsResponseBodyClients) GetClient() []*DescribeMountedClientsResponseBodyClientsClient {
	return s.Client
}

func (s *DescribeMountedClientsResponseBodyClients) SetClient(v []*DescribeMountedClientsResponseBodyClientsClient) *DescribeMountedClientsResponseBodyClients {
	s.Client = v
	return s
}

func (s *DescribeMountedClientsResponseBodyClients) Validate() error {
	return dara.Validate(s)
}

type DescribeMountedClientsResponseBodyClientsClient struct {
	// The IP address of the client.
	//
	// example:
	//
	// 10.10.10.1
	ClientIP *string `json:"ClientIP,omitempty" xml:"ClientIP,omitempty"`
}

func (s DescribeMountedClientsResponseBodyClientsClient) String() string {
	return dara.Prettify(s)
}

func (s DescribeMountedClientsResponseBodyClientsClient) GoString() string {
	return s.String()
}

func (s *DescribeMountedClientsResponseBodyClientsClient) GetClientIP() *string {
	return s.ClientIP
}

func (s *DescribeMountedClientsResponseBodyClientsClient) SetClientIP(v string) *DescribeMountedClientsResponseBodyClientsClient {
	s.ClientIP = &v
	return s
}

func (s *DescribeMountedClientsResponseBodyClientsClient) Validate() error {
	return dara.Validate(s)
}
