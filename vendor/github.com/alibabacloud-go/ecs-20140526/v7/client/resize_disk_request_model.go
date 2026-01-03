// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResizeDiskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ResizeDiskRequest
	GetClientToken() *string
	SetDiskId(v string) *ResizeDiskRequest
	GetDiskId() *string
	SetNewSize(v int32) *ResizeDiskRequest
	GetNewSize() *int32
	SetOwnerAccount(v string) *ResizeDiskRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ResizeDiskRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ResizeDiskRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ResizeDiskRequest
	GetResourceOwnerId() *int64
	SetType(v string) *ResizeDiskRequest
	GetType() *string
}

type ResizeDiskRequest struct {
	// The ID of the order.
	//
	// > This parameter is returned only when you resize subscription disks.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the disk. You can call the [DescribeDisks](https://help.aliyun.com/document_detail/25514.html) operation to query available disk IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// d-bp67acfmxazb4p****
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// The new disk capacity. Unit: GiB. Valid values:
	//
	// 	- System disk:
	//
	//     	- Basic disk (cloud): 20 to 500.
	//
	//     	- ESSD (cloud_essd): The valid values vary based on the performance level of the ESSD.
	//
	//         	- Valid values when SystemDisk.PerformanceLevel is set to PL0: 1 to 2048.
	//
	//         	- Valid values when SystemDisk.PerformanceLevel is set to PL1: 20 to 2048.
	//
	//         	- Valid values when SystemDisk.PerformanceLevel is set to PL2: 461 to 2048.
	//
	//         	- Valid values when SystemDisk.PerformanceLevel is set to PL3: 1261 to 2048.
	//
	//     	- ESSD AutoPL disk: 1 to 2048.
	//
	//     	- Other disk categories: 20 to 2048.
	//
	// 	- Data disk:
	//
	//     	- Ultra disk (cloud_efficiency): 20 to 32768.
	//
	//     	- Standard SSD (cloud_ssd): 20 to 32768.
	//
	//     	- ESSD (cloud_essd): The valid values vary based on the performance level of the ESSD.`` To query the performance level of an ESSD, call the [DescribeDisks](https://help.aliyun.com/document_detail/25514.html) operation to query disk information and check the `PerformanceLevel` value in the response.
	//
	//         	- PL0 ESSD: 1 to 65536.
	//
	//         	- PL1 ESSD: 20 to 65536.
	//
	//         	- PL2 ESSD: 461 to 65536.
	//
	//         	- PL3 ESSD: 1261 to 65536.
	//
	//     	- Basic disk (cloud): 5 to 2000.
	//
	//     	- ESSD AutoPL disk (cloud_auto): 1 to 65536.
	//
	//     	- Standard elastic ephemeral disk (elastic_ephemeral_disk_standard): 64 to 8192.
	//
	//     	- Premium elastic ephemeral disk (elastic_ephemeral_disk_premium): 64 to 8192.
	//
	// >  The new disk capacity must be larger than the original disk capacity. Otherwise, an error is reported.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1900
	NewSize              *int32  `json:"NewSize,omitempty" xml:"NewSize,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The method that you want to use to resize the disk. Specifies whether to check the image used by the instance supports hot migration. Valid values:
	//
	// 	- offline (default): resizes the disk offline. After you resize a disk offline, you must [restart the associated instance](https://help.aliyun.com/document_detail/25440.html) in the ECS console or by calling the [RebootInstance](https://help.aliyun.com/document_detail/25502.html) operation for the resizing operation to take effect.
	//
	// 	- online: resizes the disk online. After you resize a disk online, the resizing operation immediately takes effect. You do not need to restart the associated instance. You can resize ultra disks, standard SSDs, ESSDs, and elastic ephemeral disks online.
	//
	// example:
	//
	// offline
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ResizeDiskRequest) String() string {
	return dara.Prettify(s)
}

func (s ResizeDiskRequest) GoString() string {
	return s.String()
}

func (s *ResizeDiskRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ResizeDiskRequest) GetDiskId() *string {
	return s.DiskId
}

func (s *ResizeDiskRequest) GetNewSize() *int32 {
	return s.NewSize
}

func (s *ResizeDiskRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ResizeDiskRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ResizeDiskRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ResizeDiskRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ResizeDiskRequest) GetType() *string {
	return s.Type
}

func (s *ResizeDiskRequest) SetClientToken(v string) *ResizeDiskRequest {
	s.ClientToken = &v
	return s
}

func (s *ResizeDiskRequest) SetDiskId(v string) *ResizeDiskRequest {
	s.DiskId = &v
	return s
}

func (s *ResizeDiskRequest) SetNewSize(v int32) *ResizeDiskRequest {
	s.NewSize = &v
	return s
}

func (s *ResizeDiskRequest) SetOwnerAccount(v string) *ResizeDiskRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ResizeDiskRequest) SetOwnerId(v int64) *ResizeDiskRequest {
	s.OwnerId = &v
	return s
}

func (s *ResizeDiskRequest) SetResourceOwnerAccount(v string) *ResizeDiskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ResizeDiskRequest) SetResourceOwnerId(v int64) *ResizeDiskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ResizeDiskRequest) SetType(v string) *ResizeDiskRequest {
	s.Type = &v
	return s
}

func (s *ResizeDiskRequest) Validate() error {
	return dara.Validate(s)
}
