// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBlackListClientsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClients(v string) *DescribeBlackListClientsResponseBody
	GetClients() *string
	SetRequestId(v string) *DescribeBlackListClientsResponseBody
	GetRequestId() *string
}

type DescribeBlackListClientsResponseBody struct {
	// The IDs of clients and the status of each client. The parameter value is a JSON string, for example, `{"client1": "EVICTING","client2":"EVICTED"}`.
	//
	// Available client statuses include:
	//
	// 	- EVICTING: The client is being evicted.
	//
	// 	- EVICTED: The client is evicted.
	//
	// 	- ACCEPTING: The write access to the file system is being granted to the client.
	//
	// 	- ACCEPTABLE: The write access to the file system is granted to the client.
	//
	// example:
	//
	// {"client1": "EVICTING","client2":"EVICTED"}
	Clients *string `json:"Clients,omitempty" xml:"Clients,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A70BEE5D-76D3-49FB-B58F-1F398211A5C3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeBlackListClientsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBlackListClientsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBlackListClientsResponseBody) GetClients() *string {
	return s.Clients
}

func (s *DescribeBlackListClientsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBlackListClientsResponseBody) SetClients(v string) *DescribeBlackListClientsResponseBody {
	s.Clients = &v
	return s
}

func (s *DescribeBlackListClientsResponseBody) SetRequestId(v string) *DescribeBlackListClientsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBlackListClientsResponseBody) Validate() error {
	return dara.Validate(s)
}
