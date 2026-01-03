// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetDisksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDisk(v []*ResetDisksRequestDisk) *ResetDisksRequest
	GetDisk() []*ResetDisksRequestDisk
	SetDryRun(v bool) *ResetDisksRequest
	GetDryRun() *bool
	SetOwnerAccount(v string) *ResetDisksRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ResetDisksRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ResetDisksRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ResetDisksRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ResetDisksRequest
	GetResourceOwnerId() *int64
}

type ResetDisksRequest struct {
	// The disks that you want to roll back.
	//
	// This parameter is required.
	Disk []*ResetDisksRequestDisk `json:"Disk,omitempty" xml:"Disk,omitempty" type:"Repeated"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- true: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and resource state limits. If the request fails the dry run, an error code is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- false: performs a dry run and performs the actual request. If the request passes the dry run, the rollback operation is performed.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	DryRun       *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ResetDisksRequest) String() string {
	return dara.Prettify(s)
}

func (s ResetDisksRequest) GoString() string {
	return s.String()
}

func (s *ResetDisksRequest) GetDisk() []*ResetDisksRequestDisk {
	return s.Disk
}

func (s *ResetDisksRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ResetDisksRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ResetDisksRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ResetDisksRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ResetDisksRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ResetDisksRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ResetDisksRequest) SetDisk(v []*ResetDisksRequestDisk) *ResetDisksRequest {
	s.Disk = v
	return s
}

func (s *ResetDisksRequest) SetDryRun(v bool) *ResetDisksRequest {
	s.DryRun = &v
	return s
}

func (s *ResetDisksRequest) SetOwnerAccount(v string) *ResetDisksRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ResetDisksRequest) SetOwnerId(v int64) *ResetDisksRequest {
	s.OwnerId = &v
	return s
}

func (s *ResetDisksRequest) SetRegionId(v string) *ResetDisksRequest {
	s.RegionId = &v
	return s
}

func (s *ResetDisksRequest) SetResourceOwnerAccount(v string) *ResetDisksRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ResetDisksRequest) SetResourceOwnerId(v int64) *ResetDisksRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ResetDisksRequest) Validate() error {
	if s.Disk != nil {
		for _, item := range s.Disk {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ResetDisksRequestDisk struct {
	// The ID of the disk that you want to roll back. You can specify up to 10 disk IDs.
	//
	// example:
	//
	// d-j6cf7l0ewidb78lq****
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// The ID of the disk snapshot that is contained in the instance snapshot. You can specify up to 10 disk snapshot IDs.
	//
	// example:
	//
	// s-j6cdofbycydvg7ey****
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
}

func (s ResetDisksRequestDisk) String() string {
	return dara.Prettify(s)
}

func (s ResetDisksRequestDisk) GoString() string {
	return s.String()
}

func (s *ResetDisksRequestDisk) GetDiskId() *string {
	return s.DiskId
}

func (s *ResetDisksRequestDisk) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *ResetDisksRequestDisk) SetDiskId(v string) *ResetDisksRequestDisk {
	s.DiskId = &v
	return s
}

func (s *ResetDisksRequestDisk) SetSnapshotId(v string) *ResetDisksRequestDisk {
	s.SnapshotId = &v
	return s
}

func (s *ResetDisksRequestDisk) Validate() error {
	return dara.Validate(s)
}
