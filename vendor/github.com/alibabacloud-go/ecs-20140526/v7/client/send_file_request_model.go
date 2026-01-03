// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *SendFileRequest
	GetContent() *string
	SetContentType(v string) *SendFileRequest
	GetContentType() *string
	SetDescription(v string) *SendFileRequest
	GetDescription() *string
	SetFileGroup(v string) *SendFileRequest
	GetFileGroup() *string
	SetFileMode(v string) *SendFileRequest
	GetFileMode() *string
	SetFileOwner(v string) *SendFileRequest
	GetFileOwner() *string
	SetInstanceId(v []*string) *SendFileRequest
	GetInstanceId() []*string
	SetName(v string) *SendFileRequest
	GetName() *string
	SetOverwrite(v bool) *SendFileRequest
	GetOverwrite() *bool
	SetOwnerAccount(v string) *SendFileRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *SendFileRequest
	GetOwnerId() *int64
	SetRegionId(v string) *SendFileRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *SendFileRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *SendFileRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SendFileRequest
	GetResourceOwnerId() *int64
	SetTag(v []*SendFileRequestTag) *SendFileRequest
	GetTag() []*SendFileRequestTag
	SetTargetDir(v string) *SendFileRequest
	GetTargetDir() *string
	SetTimeout(v int64) *SendFileRequest
	GetTimeout() *int64
}

type SendFileRequest struct {
	// The content of the file. The file must not exceed 32 KB in size after it is encoded in Base64.
	//
	// 	- If `ContentType` is set to `PlainText`, the value of Content is in plaintext.
	//
	// 	- If `ContentType` is set to `Base64`, the value of Content is Base64-encoded.
	//
	// This parameter is required.
	//
	// example:
	//
	// #!/bin/bash  echo "Current User is :"  echo $(ps | grep "$$" | awk \\"{print $2}\\")  --------  oss://bucketName/objectName
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The content type of the file. Valid values:
	//
	// 	- PlainText: The file content is not encoded.
	//
	// 	- Base64: The file content is encoded in Base64.
	//
	// Default value: PlainText.
	//
	// example:
	//
	// PlainText
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// The description of the file. The description can be up to 512 characters in length and can contain any characters.
	//
	// example:
	//
	// This is a test file.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The group of the file. This parameter takes effect only on Linux instances. Default value: root. The value can be up to 64 characters in length.
	//
	// >  If you want to use a non-root user group, make sure that the user group exists in the instances.
	//
	// example:
	//
	// test
	FileGroup *string `json:"FileGroup,omitempty" xml:"FileGroup,omitempty"`
	// The permissions on the file. This parameter takes effect only on Linux instances. You can configure this parameter in the same way as you configure the chmod command.
	//
	// Default value: 0644, which indicates that the owner of the file has the read and write permissions on the file and that the user group of the file and other users have the read-only permissions on the file.
	//
	// example:
	//
	// 0644
	FileMode *string `json:"FileMode,omitempty" xml:"FileMode,omitempty"`
	// The owner of the file. This parameter takes effect only on Linux instances. Default value: root. The value can be up to 64 characters in length.
	//
	// >  If you want to use a non-root user, make sure that the user exists in the instances.
	//
	// example:
	//
	// test
	FileOwner *string `json:"FileOwner,omitempty" xml:"FileOwner,omitempty"`
	// The IDs of instances to which to send the file. You can specify up to 50 instance IDs in each request. Valid values of N: 1 to 50.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-bp185dy2o3o6n****
	InstanceId []*string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
	// The name of the file. The name can be up to 255 characters in length and can contain any characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// file.txt
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Specifies whether to overwrite a file in the destination directory if the file has the same name as the sent file.
	//
	// 	- true
	//
	// 	- false
	//
	// Default value: false.
	//
	// example:
	//
	// true
	Overwrite    *bool   `json:"Overwrite,omitempty" xml:"Overwrite,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the instance to which to send the file. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group. When you specify this parameter, take note of the following items:
	//
	// 	- The instance specified by the InstanceId parameter must belong to the specified resource group.
	//
	// 	- If you specify this parameter, you can call the [DescribeSendFileResults](https://help.aliyun.com/document_detail/184117.html) operation to query file sending results in the specified resource group.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tags to add to the file sending task.
	Tag []*SendFileRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The destination directory on the instance to which to send the file. If the specified directory does not exist, the system creates the directory on the instance. The value cannot exceed 255 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// /home
	TargetDir *string `json:"TargetDir,omitempty" xml:"TargetDir,omitempty"`
	// The timeout period for the file sending task. Unit: seconds.
	//
	// 	- A timeout error occurs when a file cannot be sent because the process slows down or because a specific module or Cloud Assistant Agent does not exist.
	//
	// 	- If the specified timeout period is less than 10 seconds, the system sets the timeout period to 10 seconds to ensure that the file can be sent to the instances.
	//
	// Default value: 60.
	//
	// example:
	//
	// 60
	Timeout *int64 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s SendFileRequest) String() string {
	return dara.Prettify(s)
}

func (s SendFileRequest) GoString() string {
	return s.String()
}

func (s *SendFileRequest) GetContent() *string {
	return s.Content
}

func (s *SendFileRequest) GetContentType() *string {
	return s.ContentType
}

func (s *SendFileRequest) GetDescription() *string {
	return s.Description
}

func (s *SendFileRequest) GetFileGroup() *string {
	return s.FileGroup
}

func (s *SendFileRequest) GetFileMode() *string {
	return s.FileMode
}

func (s *SendFileRequest) GetFileOwner() *string {
	return s.FileOwner
}

func (s *SendFileRequest) GetInstanceId() []*string {
	return s.InstanceId
}

func (s *SendFileRequest) GetName() *string {
	return s.Name
}

func (s *SendFileRequest) GetOverwrite() *bool {
	return s.Overwrite
}

func (s *SendFileRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *SendFileRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SendFileRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SendFileRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *SendFileRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SendFileRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SendFileRequest) GetTag() []*SendFileRequestTag {
	return s.Tag
}

func (s *SendFileRequest) GetTargetDir() *string {
	return s.TargetDir
}

func (s *SendFileRequest) GetTimeout() *int64 {
	return s.Timeout
}

func (s *SendFileRequest) SetContent(v string) *SendFileRequest {
	s.Content = &v
	return s
}

func (s *SendFileRequest) SetContentType(v string) *SendFileRequest {
	s.ContentType = &v
	return s
}

func (s *SendFileRequest) SetDescription(v string) *SendFileRequest {
	s.Description = &v
	return s
}

func (s *SendFileRequest) SetFileGroup(v string) *SendFileRequest {
	s.FileGroup = &v
	return s
}

func (s *SendFileRequest) SetFileMode(v string) *SendFileRequest {
	s.FileMode = &v
	return s
}

func (s *SendFileRequest) SetFileOwner(v string) *SendFileRequest {
	s.FileOwner = &v
	return s
}

func (s *SendFileRequest) SetInstanceId(v []*string) *SendFileRequest {
	s.InstanceId = v
	return s
}

func (s *SendFileRequest) SetName(v string) *SendFileRequest {
	s.Name = &v
	return s
}

func (s *SendFileRequest) SetOverwrite(v bool) *SendFileRequest {
	s.Overwrite = &v
	return s
}

func (s *SendFileRequest) SetOwnerAccount(v string) *SendFileRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SendFileRequest) SetOwnerId(v int64) *SendFileRequest {
	s.OwnerId = &v
	return s
}

func (s *SendFileRequest) SetRegionId(v string) *SendFileRequest {
	s.RegionId = &v
	return s
}

func (s *SendFileRequest) SetResourceGroupId(v string) *SendFileRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *SendFileRequest) SetResourceOwnerAccount(v string) *SendFileRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SendFileRequest) SetResourceOwnerId(v int64) *SendFileRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SendFileRequest) SetTag(v []*SendFileRequestTag) *SendFileRequest {
	s.Tag = v
	return s
}

func (s *SendFileRequest) SetTargetDir(v string) *SendFileRequest {
	s.TargetDir = &v
	return s
}

func (s *SendFileRequest) SetTimeout(v int64) *SendFileRequest {
	s.Timeout = &v
	return s
}

func (s *SendFileRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SendFileRequestTag struct {
	// The key of tag N of the file sending task. Valid values of N: 1 to 20. The tag key cannot be an empty string.
	//
	// If a single tag is specified to query resources, up to 1,000 resources that have this tag added can be displayed in the response. If multiple tags are specified to query resources, up to 1,000 resources that have all the tags added can be displayed in the response. To query more than 1,000 resources that have specified tags, call the [ListTagResources](https://help.aliyun.com/document_detail/110425.html) operation.
	//
	// The tag key can be up to 64 characters in length and cannot contain `http://` or `https://`. The tag key cannot start with `acs:` or `aliyun`.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N to add to the file sending task. Valid values of N: 1 to 20. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SendFileRequestTag) String() string {
	return dara.Prettify(s)
}

func (s SendFileRequestTag) GoString() string {
	return s.String()
}

func (s *SendFileRequestTag) GetKey() *string {
	return s.Key
}

func (s *SendFileRequestTag) GetValue() *string {
	return s.Value
}

func (s *SendFileRequestTag) SetKey(v string) *SendFileRequestTag {
	s.Key = &v
	return s
}

func (s *SendFileRequestTag) SetValue(v string) *SendFileRequestTag {
	s.Value = &v
	return s
}

func (s *SendFileRequestTag) Validate() error {
	return dara.Validate(s)
}
