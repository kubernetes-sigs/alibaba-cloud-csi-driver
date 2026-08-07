// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgenticSpaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAzone(v string) *CreateAgenticSpaceRequest
	GetAzone() *string
	SetClientToken(v string) *CreateAgenticSpaceRequest
	GetClientToken() *string
	SetDescription(v string) *CreateAgenticSpaceRequest
	GetDescription() *string
	SetDryRun(v bool) *CreateAgenticSpaceRequest
	GetDryRun() *bool
	SetFileSystemId(v string) *CreateAgenticSpaceRequest
	GetFileSystemId() *string
	SetFileSystemPath(v string) *CreateAgenticSpaceRequest
	GetFileSystemPath() *string
	SetQuota(v *CreateAgenticSpaceRequestQuota) *CreateAgenticSpaceRequest
	GetQuota() *CreateAgenticSpaceRequestQuota
}

type CreateAgenticSpaceRequest struct {
	// The zone ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou-i
	Azone *string `json:"Azone,omitempty" xml:"Azone,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but make sure that the token is unique among different requests.
	//
	// The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// > If you do not specify this parameter, the system automatically uses the RequestId of the API request as the ClientToken. The RequestId may differ for each API request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the AgenticSpace.
	//
	// Limits:
	//
	// - The description must be 2 to 128 characters in length.
	//
	// - The description must start with a letter or a Chinese character and cannot start with `http://` or `https://`.
	//
	// - The description can contain digits, colons (:), underscores (_), or hyphens (-).
	//
	// example:
	//
	// AgenticSpace Description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to perform a dry run for this request. A dry run checks parameter validity and dependencies without actually modifying the instance or incurring fees.
	//
	// Valid values:
	//
	// - true: sends a dry run request without modifying the protocol service. The system checks whether required parameters are specified, whether the request format is valid, and whether business limit dependencies are met. If the check fails, the corresponding error is returned. If the check succeeds, HTTP status code 200 is returned.
	//
	// - false (default): sends a normal request. After the check is passed, the protocol service is directly modified.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The file system ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 31a8e4****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The absolute path of the file. Only first-level directories are supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// /path/
	FileSystemPath *string `json:"FileSystemPath,omitempty" xml:"FileSystemPath,omitempty"`
	// The quota information.
	//
	// This parameter is required.
	Quota *CreateAgenticSpaceRequestQuota `json:"Quota,omitempty" xml:"Quota,omitempty" type:"Struct"`
}

func (s CreateAgenticSpaceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAgenticSpaceRequest) GoString() string {
	return s.String()
}

func (s *CreateAgenticSpaceRequest) GetAzone() *string {
	return s.Azone
}

func (s *CreateAgenticSpaceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateAgenticSpaceRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateAgenticSpaceRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateAgenticSpaceRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *CreateAgenticSpaceRequest) GetFileSystemPath() *string {
	return s.FileSystemPath
}

func (s *CreateAgenticSpaceRequest) GetQuota() *CreateAgenticSpaceRequestQuota {
	return s.Quota
}

func (s *CreateAgenticSpaceRequest) SetAzone(v string) *CreateAgenticSpaceRequest {
	s.Azone = &v
	return s
}

func (s *CreateAgenticSpaceRequest) SetClientToken(v string) *CreateAgenticSpaceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateAgenticSpaceRequest) SetDescription(v string) *CreateAgenticSpaceRequest {
	s.Description = &v
	return s
}

func (s *CreateAgenticSpaceRequest) SetDryRun(v bool) *CreateAgenticSpaceRequest {
	s.DryRun = &v
	return s
}

func (s *CreateAgenticSpaceRequest) SetFileSystemId(v string) *CreateAgenticSpaceRequest {
	s.FileSystemId = &v
	return s
}

func (s *CreateAgenticSpaceRequest) SetFileSystemPath(v string) *CreateAgenticSpaceRequest {
	s.FileSystemPath = &v
	return s
}

func (s *CreateAgenticSpaceRequest) SetQuota(v *CreateAgenticSpaceRequestQuota) *CreateAgenticSpaceRequest {
	s.Quota = v
	return s
}

func (s *CreateAgenticSpaceRequest) Validate() error {
	if s.Quota != nil {
		if err := s.Quota.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAgenticSpaceRequestQuota struct {
	// The maximum number of files allowed by the quota. Valid values:
	//
	// - Minimum value: 10000.
	//
	// - Maximum value: 100000000.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	FileCountLimit *int64 `json:"FileCountLimit,omitempty" xml:"FileCountLimit,omitempty"`
	// The total capacity limit of the quota. Unit: bytes.
	//
	// Valid values:
	//
	// - Minimum value: 10737418240 (10 GiB).
	//
	// - Maximum value: 1099511627776000 (1024000 GiB).
	//
	// - Increment: 1073741824 (1 GiB).
	//
	// This parameter is required.
	//
	// example:
	//
	// 10737418240
	SizeLimit *int64 `json:"SizeLimit,omitempty" xml:"SizeLimit,omitempty"`
}

func (s CreateAgenticSpaceRequestQuota) String() string {
	return dara.Prettify(s)
}

func (s CreateAgenticSpaceRequestQuota) GoString() string {
	return s.String()
}

func (s *CreateAgenticSpaceRequestQuota) GetFileCountLimit() *int64 {
	return s.FileCountLimit
}

func (s *CreateAgenticSpaceRequestQuota) GetSizeLimit() *int64 {
	return s.SizeLimit
}

func (s *CreateAgenticSpaceRequestQuota) SetFileCountLimit(v int64) *CreateAgenticSpaceRequestQuota {
	s.FileCountLimit = &v
	return s
}

func (s *CreateAgenticSpaceRequestQuota) SetSizeLimit(v int64) *CreateAgenticSpaceRequestQuota {
	s.SizeLimit = &v
	return s
}

func (s *CreateAgenticSpaceRequestQuota) Validate() error {
	return dara.Validate(s)
}
