// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProtocolMountTargetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *GetProtocolMountTargetRequest
	GetClientToken() *string
	SetExportId(v string) *GetProtocolMountTargetRequest
	GetExportId() *string
	SetFileSystemId(v string) *GetProtocolMountTargetRequest
	GetFileSystemId() *string
	SetMaxResults(v int64) *GetProtocolMountTargetRequest
	GetMaxResults() *int64
	SetNextToken(v string) *GetProtocolMountTargetRequest
	GetNextToken() *string
	SetProtocolServiceId(v string) *GetProtocolMountTargetRequest
	GetProtocolServiceId() *string
}

type GetProtocolMountTargetRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but make sure that the token is unique among different requests.
	//
	// The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// > If you do not specify this parameter, the system automatically uses the RequestId of the API request as the ClientToken. The RequestId may vary for each API request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The export directory ID of the protocol service. **Required**.
	//
	// **How to obtain**:
	//
	// - Call [CreateProtocolMountTarget](https://www.alibabacloud.com/help/en/cpfs/cpfsonecs/developer-reference/api-nas-2017-06-26-createprotocolmounttarget-cpfs) to create an export directory and obtain the ExportId from the response.
	//
	// - Or call [DescribeProtocolMountTarget](https://www.alibabacloud.com/help/en/cpfs/cpfsonecs/developer-reference/api-nas-2017-06-26-describeprotocolmounttarget-cpfs) to query the list of export directories and obtain the ExportId from the response.
	//
	// This parameter is required.
	//
	// example:
	//
	// exp-19abf5beab8d****
	ExportId *string `json:"ExportId,omitempty" xml:"ExportId,omitempty"`
	// The file system ID.
	//
	// - CPFS: The ID must start with `cpfs-`, such as cpfs-125487\\*\\*\\*\\*.
	//
	// - CPFS for Lingjun: The ID must start with `bmcpfs-`, such as bmcpfs-0015\\*\\*\\*\\*.
	//
	// - CPFS SE: The ID must start with `cpfsse-`, such as cpfsse-022c71b134\\*\\*\\*\\*.
	//
	// This parameter is required.
	//
	// example:
	//
	// cpfs-099394bd928c****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The maximum number of results to return per query.
	//
	// Valid values: 10 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// If the response is truncated, you can use NextToken to send a subsequent request to retrieve the content after the current truncation point.
	//
	// example:
	//
	// M18xMA==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The protocol service ID. **Required**.
	//
	// **How to obtain**: Call the [DescribeProtocolService](https://www.alibabacloud.com/help/en/cpfs/cpfsonecs/developer-reference/api-nas-2017-06-26-describeprotocolservice-cpfs) operation to query the list of protocol services and obtain the ProtocolServiceId from the response.
	//
	// This parameter is required.
	//
	// example:
	//
	// ptc-197ed6a00f2b****
	ProtocolServiceId *string `json:"ProtocolServiceId,omitempty" xml:"ProtocolServiceId,omitempty"`
}

func (s GetProtocolMountTargetRequest) String() string {
	return dara.Prettify(s)
}

func (s GetProtocolMountTargetRequest) GoString() string {
	return s.String()
}

func (s *GetProtocolMountTargetRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *GetProtocolMountTargetRequest) GetExportId() *string {
	return s.ExportId
}

func (s *GetProtocolMountTargetRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *GetProtocolMountTargetRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *GetProtocolMountTargetRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *GetProtocolMountTargetRequest) GetProtocolServiceId() *string {
	return s.ProtocolServiceId
}

func (s *GetProtocolMountTargetRequest) SetClientToken(v string) *GetProtocolMountTargetRequest {
	s.ClientToken = &v
	return s
}

func (s *GetProtocolMountTargetRequest) SetExportId(v string) *GetProtocolMountTargetRequest {
	s.ExportId = &v
	return s
}

func (s *GetProtocolMountTargetRequest) SetFileSystemId(v string) *GetProtocolMountTargetRequest {
	s.FileSystemId = &v
	return s
}

func (s *GetProtocolMountTargetRequest) SetMaxResults(v int64) *GetProtocolMountTargetRequest {
	s.MaxResults = &v
	return s
}

func (s *GetProtocolMountTargetRequest) SetNextToken(v string) *GetProtocolMountTargetRequest {
	s.NextToken = &v
	return s
}

func (s *GetProtocolMountTargetRequest) SetProtocolServiceId(v string) *GetProtocolMountTargetRequest {
	s.ProtocolServiceId = &v
	return s
}

func (s *GetProtocolMountTargetRequest) Validate() error {
	return dara.Validate(s)
}
