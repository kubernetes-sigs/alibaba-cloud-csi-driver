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
	SetName(v string) *SendFileRequest
	GetName() *string
	SetNodeIdList(v []*string) *SendFileRequest
	GetNodeIdList() []*string
	SetOverwrite(v bool) *SendFileRequest
	GetOverwrite() *bool
	SetTargetDir(v string) *SendFileRequest
	GetTargetDir() *string
	SetTimeout(v int32) *SendFileRequest
	GetTimeout() *int32
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
	// #!/bin/bash echo "Current User is :" echo $(ps | grep "$$" | awk \\"{print $2}\\") -------- oss://bucketName/objectName
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The content type of the file. Valid values:
	//
	// PlainText Base64 Default value: PlainText.
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
	// The user group of the file. This parameter takes effect only on Linux instances. Default value: root. The value can be up to 64 characters in length.
	//
	// Note If you want to use a non-root user group, make sure that the user group exists in the instances.
	//
	// example:
	//
	// test
	FileGroup *string `json:"FileGroup,omitempty" xml:"FileGroup,omitempty"`
	// The permissions on the file. This parameter takes effect only on Linux instances. You can configure this parameter in the same way as you configure the chmod command.
	//
	// Default value: 0644: the owner of the file has the read and write permission. The user group of the file and other users have read-only permission.
	//
	// example:
	//
	// 0644
	FileMode *string `json:"FileMode,omitempty" xml:"FileMode,omitempty"`
	// The owner of the file. This parameter takes effect only on Linux instances. Default value: root.
	//
	// example:
	//
	// root
	FileOwner *string `json:"FileOwner,omitempty" xml:"FileOwner,omitempty"`
	// The file name. The name can be up to 255 characters in length and can contain any characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// file.txt
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The node list.
	//
	// This parameter is required.
	NodeIdList []*string `json:"NodeIdList,omitempty" xml:"NodeIdList,omitempty" type:"Repeated"`
	// Specifies whether to overwrite file with the same name in the destination directory.
	//
	// 	- true
	//
	// 	- false
	//
	// Default value: false.
	//
	// example:
	//
	// True
	Overwrite *bool `json:"Overwrite,omitempty" xml:"Overwrite,omitempty"`
	// The directory in the Lingjun node to which the file is sent. If the specified directory does not exist, the system creates the directory automatically.
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
	// 	- If the specified timeout period is less than 10 seconds, the system sets the timeout period to 10 seconds to ensure that the file can be sent.
	//
	// Default value: 60.
	//
	// example:
	//
	// 600
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
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

func (s *SendFileRequest) GetName() *string {
	return s.Name
}

func (s *SendFileRequest) GetNodeIdList() []*string {
	return s.NodeIdList
}

func (s *SendFileRequest) GetOverwrite() *bool {
	return s.Overwrite
}

func (s *SendFileRequest) GetTargetDir() *string {
	return s.TargetDir
}

func (s *SendFileRequest) GetTimeout() *int32 {
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

func (s *SendFileRequest) SetName(v string) *SendFileRequest {
	s.Name = &v
	return s
}

func (s *SendFileRequest) SetNodeIdList(v []*string) *SendFileRequest {
	s.NodeIdList = v
	return s
}

func (s *SendFileRequest) SetOverwrite(v bool) *SendFileRequest {
	s.Overwrite = &v
	return s
}

func (s *SendFileRequest) SetTargetDir(v string) *SendFileRequest {
	s.TargetDir = &v
	return s
}

func (s *SendFileRequest) SetTimeout(v int32) *SendFileRequest {
	s.Timeout = &v
	return s
}

func (s *SendFileRequest) Validate() error {
	return dara.Validate(s)
}
