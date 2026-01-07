// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportImageRequest interface {
  dara.Model
  String() string
  GoString() string
  SetDryRun(v bool) *ExportImageRequest
  GetDryRun() *bool 
  SetImageFormat(v string) *ExportImageRequest
  GetImageFormat() *string 
  SetImageId(v string) *ExportImageRequest
  GetImageId() *string 
  SetOSSBucket(v string) *ExportImageRequest
  GetOSSBucket() *string 
  SetOSSPrefix(v string) *ExportImageRequest
  GetOSSPrefix() *string 
  SetOwnerId(v int64) *ExportImageRequest
  GetOwnerId() *int64 
  SetRegionId(v string) *ExportImageRequest
  GetRegionId() *string 
  SetResourceOwnerAccount(v string) *ExportImageRequest
  GetResourceOwnerAccount() *string 
  SetResourceOwnerId(v int64) *ExportImageRequest
  GetResourceOwnerId() *int64 
  SetRoleName(v string) *ExportImageRequest
  GetRoleName() *string 
}

type ExportImageRequest struct {
  DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
  // The format in which you want to export the custom image. Valid values:
  // 
  // 	- raw
  // 
  // 	- vhd
  // 
  // 	- qcow2
  // 
  // 	- vmdk
  // 
  // 	- vdi
  // 
  // Default value: raw.
  // 
  // example:
  // 
  // raw
  ImageFormat *string `json:"ImageFormat,omitempty" xml:"ImageFormat,omitempty"`
  // The custom image ID.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // m-bp67acfmxazb4p****
  ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
  // The OSS bucket in which you want to store the exported custom image.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // testexportImage
  OSSBucket *string `json:"OSSBucket,omitempty" xml:"OSSBucket,omitempty"`
  // The prefix for the name of the OSS object. The prefix must be 1 to 30 characters in length and can contain digits and letters.
  // 
  // example:
  // 
  // EcsExport
  OSSPrefix *string `json:"OSSPrefix,omitempty" xml:"OSSPrefix,omitempty"`
  OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
  // The region ID of the custom image. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // cn-hangzhou
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
  ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
  ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
  // The name of the RAM role that you want to use to export the custom image.
  // 
  // example:
  // 
  // AliyunECSImageExportDefaultRole
  RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
}

func (s ExportImageRequest) String() string {
  return dara.Prettify(s)
}

func (s ExportImageRequest) GoString() string {
  return s.String()
}

func (s *ExportImageRequest) GetDryRun() *bool  {
  return s.DryRun
}

func (s *ExportImageRequest) GetImageFormat() *string  {
  return s.ImageFormat
}

func (s *ExportImageRequest) GetImageId() *string  {
  return s.ImageId
}

func (s *ExportImageRequest) GetOSSBucket() *string  {
  return s.OSSBucket
}

func (s *ExportImageRequest) GetOSSPrefix() *string  {
  return s.OSSPrefix
}

func (s *ExportImageRequest) GetOwnerId() *int64  {
  return s.OwnerId
}

func (s *ExportImageRequest) GetRegionId() *string  {
  return s.RegionId
}

func (s *ExportImageRequest) GetResourceOwnerAccount() *string  {
  return s.ResourceOwnerAccount
}

func (s *ExportImageRequest) GetResourceOwnerId() *int64  {
  return s.ResourceOwnerId
}

func (s *ExportImageRequest) GetRoleName() *string  {
  return s.RoleName
}

func (s *ExportImageRequest) SetDryRun(v bool) *ExportImageRequest {
  s.DryRun = &v
  return s
}

func (s *ExportImageRequest) SetImageFormat(v string) *ExportImageRequest {
  s.ImageFormat = &v
  return s
}

func (s *ExportImageRequest) SetImageId(v string) *ExportImageRequest {
  s.ImageId = &v
  return s
}

func (s *ExportImageRequest) SetOSSBucket(v string) *ExportImageRequest {
  s.OSSBucket = &v
  return s
}

func (s *ExportImageRequest) SetOSSPrefix(v string) *ExportImageRequest {
  s.OSSPrefix = &v
  return s
}

func (s *ExportImageRequest) SetOwnerId(v int64) *ExportImageRequest {
  s.OwnerId = &v
  return s
}

func (s *ExportImageRequest) SetRegionId(v string) *ExportImageRequest {
  s.RegionId = &v
  return s
}

func (s *ExportImageRequest) SetResourceOwnerAccount(v string) *ExportImageRequest {
  s.ResourceOwnerAccount = &v
  return s
}

func (s *ExportImageRequest) SetResourceOwnerId(v int64) *ExportImageRequest {
  s.ResourceOwnerId = &v
  return s
}

func (s *ExportImageRequest) SetRoleName(v string) *ExportImageRequest {
  s.RoleName = &v
  return s
}

func (s *ExportImageRequest) Validate() error {
  return dara.Validate(s)
}

