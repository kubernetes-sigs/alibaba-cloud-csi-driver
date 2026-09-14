// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCpfsAccessPointMountedClientsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMountedClient(v []*DescribeCpfsAccessPointMountedClientsResponseBodyMountedClient) *DescribeCpfsAccessPointMountedClientsResponseBody
	GetMountedClient() []*DescribeCpfsAccessPointMountedClientsResponseBodyMountedClient
	SetPageNumber(v int32) *DescribeCpfsAccessPointMountedClientsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeCpfsAccessPointMountedClientsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeCpfsAccessPointMountedClientsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeCpfsAccessPointMountedClientsResponseBody
	GetTotalCount() *int32
}

type DescribeCpfsAccessPointMountedClientsResponseBody struct {
	// The clients mounted to the access point.
	MountedClient []*DescribeCpfsAccessPointMountedClientsResponseBodyMountedClient `json:"MountedClient,omitempty" xml:"MountedClient,omitempty" type:"Repeated"`
	// The page number of the list.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of results on each page. Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of clients mounted to the access point.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeCpfsAccessPointMountedClientsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCpfsAccessPointMountedClientsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCpfsAccessPointMountedClientsResponseBody) GetMountedClient() []*DescribeCpfsAccessPointMountedClientsResponseBodyMountedClient {
	return s.MountedClient
}

func (s *DescribeCpfsAccessPointMountedClientsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeCpfsAccessPointMountedClientsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCpfsAccessPointMountedClientsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCpfsAccessPointMountedClientsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeCpfsAccessPointMountedClientsResponseBody) SetMountedClient(v []*DescribeCpfsAccessPointMountedClientsResponseBodyMountedClient) *DescribeCpfsAccessPointMountedClientsResponseBody {
	s.MountedClient = v
	return s
}

func (s *DescribeCpfsAccessPointMountedClientsResponseBody) SetPageNumber(v int32) *DescribeCpfsAccessPointMountedClientsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeCpfsAccessPointMountedClientsResponseBody) SetPageSize(v int32) *DescribeCpfsAccessPointMountedClientsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeCpfsAccessPointMountedClientsResponseBody) SetRequestId(v string) *DescribeCpfsAccessPointMountedClientsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCpfsAccessPointMountedClientsResponseBody) SetTotalCount(v int32) *DescribeCpfsAccessPointMountedClientsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeCpfsAccessPointMountedClientsResponseBody) Validate() error {
	if s.MountedClient != nil {
		for _, item := range s.MountedClient {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCpfsAccessPointMountedClientsResponseBodyMountedClient struct {
	// The type of the client mounted to the access point. Valid values:
	//
	// - vsc
	//
	// - tcp
	//
	// example:
	//
	// vsc
	ChannelType *string `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	// The ID of the client mounted to the access point.
	//
	// example:
	//
	// vsc-8vb864o3ppwfvh****
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The IP address of the client mounted to the access point.
	//
	// example:
	//
	// 219.145.34.210
	ClientIp *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
}

func (s DescribeCpfsAccessPointMountedClientsResponseBodyMountedClient) String() string {
	return dara.Prettify(s)
}

func (s DescribeCpfsAccessPointMountedClientsResponseBodyMountedClient) GoString() string {
	return s.String()
}

func (s *DescribeCpfsAccessPointMountedClientsResponseBodyMountedClient) GetChannelType() *string {
	return s.ChannelType
}

func (s *DescribeCpfsAccessPointMountedClientsResponseBodyMountedClient) GetClientId() *string {
	return s.ClientId
}

func (s *DescribeCpfsAccessPointMountedClientsResponseBodyMountedClient) GetClientIp() *string {
	return s.ClientIp
}

func (s *DescribeCpfsAccessPointMountedClientsResponseBodyMountedClient) SetChannelType(v string) *DescribeCpfsAccessPointMountedClientsResponseBodyMountedClient {
	s.ChannelType = &v
	return s
}

func (s *DescribeCpfsAccessPointMountedClientsResponseBodyMountedClient) SetClientId(v string) *DescribeCpfsAccessPointMountedClientsResponseBodyMountedClient {
	s.ClientId = &v
	return s
}

func (s *DescribeCpfsAccessPointMountedClientsResponseBodyMountedClient) SetClientIp(v string) *DescribeCpfsAccessPointMountedClientsResponseBodyMountedClient {
	s.ClientIp = &v
	return s
}

func (s *DescribeCpfsAccessPointMountedClientsResponseBodyMountedClient) Validate() error {
	return dara.Validate(s)
}
