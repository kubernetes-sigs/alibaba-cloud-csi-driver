// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFileSystemsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFileSystems(v *DescribeFileSystemsResponseBodyFileSystems) *DescribeFileSystemsResponseBody
	GetFileSystems() *DescribeFileSystemsResponseBodyFileSystems
	SetPageNumber(v int32) *DescribeFileSystemsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeFileSystemsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeFileSystemsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeFileSystemsResponseBody
	GetTotalCount() *int32
}

type DescribeFileSystemsResponseBody struct {
	FileSystems *DescribeFileSystemsResponseBodyFileSystems `json:"FileSystems,omitempty" xml:"FileSystems,omitempty" type:"Struct"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 035B3A3A-E514-4B41-B906-5D906CFB****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of file systems.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeFileSystemsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeFileSystemsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemsResponseBody) GetFileSystems() *DescribeFileSystemsResponseBodyFileSystems {
	return s.FileSystems
}

func (s *DescribeFileSystemsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeFileSystemsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeFileSystemsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeFileSystemsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeFileSystemsResponseBody) SetFileSystems(v *DescribeFileSystemsResponseBodyFileSystems) *DescribeFileSystemsResponseBody {
	s.FileSystems = v
	return s
}

func (s *DescribeFileSystemsResponseBody) SetPageNumber(v int32) *DescribeFileSystemsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeFileSystemsResponseBody) SetPageSize(v int32) *DescribeFileSystemsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeFileSystemsResponseBody) SetRequestId(v string) *DescribeFileSystemsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFileSystemsResponseBody) SetTotalCount(v int32) *DescribeFileSystemsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeFileSystemsResponseBody) Validate() error {
	if s.FileSystems != nil {
		if err := s.FileSystems.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeFileSystemsResponseBodyFileSystems struct {
	FileSystem []*DescribeFileSystemsResponseBodyFileSystemsFileSystem `json:"FileSystem,omitempty" xml:"FileSystem,omitempty" type:"Repeated"`
}

func (s DescribeFileSystemsResponseBodyFileSystems) String() string {
	return dara.Prettify(s)
}

func (s DescribeFileSystemsResponseBodyFileSystems) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemsResponseBodyFileSystems) GetFileSystem() []*DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	return s.FileSystem
}

func (s *DescribeFileSystemsResponseBodyFileSystems) SetFileSystem(v []*DescribeFileSystemsResponseBodyFileSystemsFileSystem) *DescribeFileSystemsResponseBodyFileSystems {
	s.FileSystem = v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystems) Validate() error {
	if s.FileSystem != nil {
		for _, item := range s.FileSystem {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeFileSystemsResponseBodyFileSystemsFileSystem struct {
	AccessPointCount     *string                                                           `json:"AccessPointCount,omitempty" xml:"AccessPointCount,omitempty"`
	AutoSnapshotPolicyId *string                                                           `json:"AutoSnapshotPolicyId,omitempty" xml:"AutoSnapshotPolicyId,omitempty"`
	Bandwidth            *int64                                                            `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	Capacity             *int64                                                            `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	ChargeType           *string                                                           `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	CreateTime           *string                                                           `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description          *string                                                           `json:"Description,omitempty" xml:"Description,omitempty"`
	EncryptType          *int32                                                            `json:"EncryptType,omitempty" xml:"EncryptType,omitempty"`
	ExpiredTime          *string                                                           `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	FileSystemId         *string                                                           `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	FileSystemType       *string                                                           `json:"FileSystemType,omitempty" xml:"FileSystemType,omitempty"`
	KMSKeyId             *string                                                           `json:"KMSKeyId,omitempty" xml:"KMSKeyId,omitempty"`
	Ldap                 *DescribeFileSystemsResponseBodyFileSystemsFileSystemLdap         `json:"Ldap,omitempty" xml:"Ldap,omitempty" type:"Struct"`
	MeteredArchiveSize   *int64                                                            `json:"MeteredArchiveSize,omitempty" xml:"MeteredArchiveSize,omitempty"`
	MeteredIASize        *int64                                                            `json:"MeteredIASize,omitempty" xml:"MeteredIASize,omitempty"`
	MeteredSize          *int64                                                            `json:"MeteredSize,omitempty" xml:"MeteredSize,omitempty"`
	MountTargets         *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargets `json:"MountTargets,omitempty" xml:"MountTargets,omitempty" type:"Struct"`
	Options              *DescribeFileSystemsResponseBodyFileSystemsFileSystemOptions      `json:"Options,omitempty" xml:"Options,omitempty" type:"Struct"`
	Packages             *DescribeFileSystemsResponseBodyFileSystemsFileSystemPackages     `json:"Packages,omitempty" xml:"Packages,omitempty" type:"Struct"`
	ProtocolType         *string                                                           `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	QuorumVswId          *string                                                           `json:"QuorumVswId,omitempty" xml:"QuorumVswId,omitempty"`
	// example:
	//
	// LRS
	RedundancyType       *string                                                                   `json:"RedundancyType,omitempty" xml:"RedundancyType,omitempty"`
	RedundancyVSwitchIds *DescribeFileSystemsResponseBodyFileSystemsFileSystemRedundancyVSwitchIds `json:"RedundancyVSwitchIds,omitempty" xml:"RedundancyVSwitchIds,omitempty" type:"Struct"`
	RegionId             *string                                                                   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId      *string                                                                   `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Status               *string                                                                   `json:"Status,omitempty" xml:"Status,omitempty"`
	StorageType          *string                                                                   `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	SupportedFeatures    *DescribeFileSystemsResponseBodyFileSystemsFileSystemSupportedFeatures    `json:"SupportedFeatures,omitempty" xml:"SupportedFeatures,omitempty" type:"Struct"`
	Tags                 *DescribeFileSystemsResponseBodyFileSystemsFileSystemTags                 `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	Version              *string                                                                   `json:"Version,omitempty" xml:"Version,omitempty"`
	VpcId                *string                                                                   `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VscTarget            *string                                                                   `json:"VscTarget,omitempty" xml:"VscTarget,omitempty"`
	VswIds               *DescribeFileSystemsResponseBodyFileSystemsFileSystemVswIds               `json:"VswIds,omitempty" xml:"VswIds,omitempty" type:"Struct"`
	ZoneId               *string                                                                   `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystem) String() string {
	return dara.Prettify(s)
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystem) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) GetAccessPointCount() *string {
	return s.AccessPointCount
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) GetAutoSnapshotPolicyId() *string {
	return s.AutoSnapshotPolicyId
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) GetBandwidth() *int64 {
	return s.Bandwidth
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) GetCapacity() *int64 {
	return s.Capacity
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) GetChargeType() *string {
	return s.ChargeType
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) GetDescription() *string {
	return s.Description
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) GetEncryptType() *int32 {
	return s.EncryptType
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) GetExpiredTime() *string {
	return s.ExpiredTime
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) GetFileSystemType() *string {
	return s.FileSystemType
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) GetKMSKeyId() *string {
	return s.KMSKeyId
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) GetLdap() *DescribeFileSystemsResponseBodyFileSystemsFileSystemLdap {
	return s.Ldap
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) GetMeteredArchiveSize() *int64 {
	return s.MeteredArchiveSize
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) GetMeteredIASize() *int64 {
	return s.MeteredIASize
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) GetMeteredSize() *int64 {
	return s.MeteredSize
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) GetMountTargets() *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargets {
	return s.MountTargets
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) GetOptions() *DescribeFileSystemsResponseBodyFileSystemsFileSystemOptions {
	return s.Options
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) GetPackages() *DescribeFileSystemsResponseBodyFileSystemsFileSystemPackages {
	return s.Packages
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) GetProtocolType() *string {
	return s.ProtocolType
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) GetQuorumVswId() *string {
	return s.QuorumVswId
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) GetRedundancyType() *string {
	return s.RedundancyType
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) GetRedundancyVSwitchIds() *DescribeFileSystemsResponseBodyFileSystemsFileSystemRedundancyVSwitchIds {
	return s.RedundancyVSwitchIds
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) GetStatus() *string {
	return s.Status
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) GetStorageType() *string {
	return s.StorageType
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) GetSupportedFeatures() *DescribeFileSystemsResponseBodyFileSystemsFileSystemSupportedFeatures {
	return s.SupportedFeatures
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) GetTags() *DescribeFileSystemsResponseBodyFileSystemsFileSystemTags {
	return s.Tags
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) GetVersion() *string {
	return s.Version
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) GetVscTarget() *string {
	return s.VscTarget
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) GetVswIds() *DescribeFileSystemsResponseBodyFileSystemsFileSystemVswIds {
	return s.VswIds
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetAccessPointCount(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.AccessPointCount = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetAutoSnapshotPolicyId(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.AutoSnapshotPolicyId = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetBandwidth(v int64) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.Bandwidth = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetCapacity(v int64) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.Capacity = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetChargeType(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.ChargeType = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetCreateTime(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.CreateTime = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetDescription(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.Description = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetEncryptType(v int32) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.EncryptType = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetExpiredTime(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetFileSystemId(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.FileSystemId = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetFileSystemType(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.FileSystemType = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetKMSKeyId(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.KMSKeyId = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetLdap(v *DescribeFileSystemsResponseBodyFileSystemsFileSystemLdap) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.Ldap = v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetMeteredArchiveSize(v int64) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.MeteredArchiveSize = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetMeteredIASize(v int64) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.MeteredIASize = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetMeteredSize(v int64) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.MeteredSize = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetMountTargets(v *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargets) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.MountTargets = v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetOptions(v *DescribeFileSystemsResponseBodyFileSystemsFileSystemOptions) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.Options = v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetPackages(v *DescribeFileSystemsResponseBodyFileSystemsFileSystemPackages) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.Packages = v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetProtocolType(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.ProtocolType = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetQuorumVswId(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.QuorumVswId = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetRedundancyType(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.RedundancyType = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetRedundancyVSwitchIds(v *DescribeFileSystemsResponseBodyFileSystemsFileSystemRedundancyVSwitchIds) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.RedundancyVSwitchIds = v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetRegionId(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.RegionId = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetResourceGroupId(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetStatus(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.Status = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetStorageType(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.StorageType = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetSupportedFeatures(v *DescribeFileSystemsResponseBodyFileSystemsFileSystemSupportedFeatures) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.SupportedFeatures = v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetTags(v *DescribeFileSystemsResponseBodyFileSystemsFileSystemTags) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.Tags = v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetVersion(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.Version = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetVpcId(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.VpcId = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetVscTarget(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.VscTarget = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetVswIds(v *DescribeFileSystemsResponseBodyFileSystemsFileSystemVswIds) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.VswIds = v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetZoneId(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.ZoneId = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) Validate() error {
	if s.Ldap != nil {
		if err := s.Ldap.Validate(); err != nil {
			return err
		}
	}
	if s.MountTargets != nil {
		if err := s.MountTargets.Validate(); err != nil {
			return err
		}
	}
	if s.Options != nil {
		if err := s.Options.Validate(); err != nil {
			return err
		}
	}
	if s.Packages != nil {
		if err := s.Packages.Validate(); err != nil {
			return err
		}
	}
	if s.RedundancyVSwitchIds != nil {
		if err := s.RedundancyVSwitchIds.Validate(); err != nil {
			return err
		}
	}
	if s.SupportedFeatures != nil {
		if err := s.SupportedFeatures.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	if s.VswIds != nil {
		if err := s.VswIds.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeFileSystemsResponseBodyFileSystemsFileSystemLdap struct {
	BindDN     *string `json:"BindDN,omitempty" xml:"BindDN,omitempty"`
	SearchBase *string `json:"SearchBase,omitempty" xml:"SearchBase,omitempty"`
	URI        *string `json:"URI,omitempty" xml:"URI,omitempty"`
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystemLdap) String() string {
	return dara.Prettify(s)
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystemLdap) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemLdap) GetBindDN() *string {
	return s.BindDN
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemLdap) GetSearchBase() *string {
	return s.SearchBase
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemLdap) GetURI() *string {
	return s.URI
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemLdap) SetBindDN(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystemLdap {
	s.BindDN = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemLdap) SetSearchBase(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystemLdap {
	s.SearchBase = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemLdap) SetURI(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystemLdap {
	s.URI = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemLdap) Validate() error {
	return dara.Validate(s)
}

type DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargets struct {
	MountTarget []*DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget `json:"MountTarget,omitempty" xml:"MountTarget,omitempty" type:"Repeated"`
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargets) String() string {
	return dara.Prettify(s)
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargets) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargets) GetMountTarget() []*DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget {
	return s.MountTarget
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargets) SetMountTarget(v []*DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget) *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargets {
	s.MountTarget = v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargets) Validate() error {
	if s.MountTarget != nil {
		for _, item := range s.MountTarget {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget struct {
	AccessGroupName            *string                                                                                       `json:"AccessGroupName,omitempty" xml:"AccessGroupName,omitempty"`
	ClientMasterNodes          *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetClientMasterNodes `json:"ClientMasterNodes,omitempty" xml:"ClientMasterNodes,omitempty" type:"Struct"`
	DualStackMountTargetDomain *string                                                                                       `json:"DualStackMountTargetDomain,omitempty" xml:"DualStackMountTargetDomain,omitempty"`
	MountTargetDomain          *string                                                                                       `json:"MountTargetDomain,omitempty" xml:"MountTargetDomain,omitempty"`
	NetworkType                *string                                                                                       `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	Status                     *string                                                                                       `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags                       *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetTags              `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	VpcId                      *string                                                                                       `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VswId                      *string                                                                                       `json:"VswId,omitempty" xml:"VswId,omitempty"`
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget) String() string {
	return dara.Prettify(s)
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget) GetAccessGroupName() *string {
	return s.AccessGroupName
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget) GetClientMasterNodes() *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetClientMasterNodes {
	return s.ClientMasterNodes
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget) GetDualStackMountTargetDomain() *string {
	return s.DualStackMountTargetDomain
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget) GetMountTargetDomain() *string {
	return s.MountTargetDomain
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget) GetNetworkType() *string {
	return s.NetworkType
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget) GetStatus() *string {
	return s.Status
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget) GetTags() *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetTags {
	return s.Tags
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget) GetVswId() *string {
	return s.VswId
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget) SetAccessGroupName(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget {
	s.AccessGroupName = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget) SetClientMasterNodes(v *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetClientMasterNodes) *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget {
	s.ClientMasterNodes = v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget) SetDualStackMountTargetDomain(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget {
	s.DualStackMountTargetDomain = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget) SetMountTargetDomain(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget {
	s.MountTargetDomain = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget) SetNetworkType(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget {
	s.NetworkType = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget) SetStatus(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget {
	s.Status = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget) SetTags(v *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetTags) *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget {
	s.Tags = v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget) SetVpcId(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget {
	s.VpcId = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget) SetVswId(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget {
	s.VswId = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget) Validate() error {
	if s.ClientMasterNodes != nil {
		if err := s.ClientMasterNodes.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetClientMasterNodes struct {
	ClientMasterNode []*DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetClientMasterNodesClientMasterNode `json:"ClientMasterNode,omitempty" xml:"ClientMasterNode,omitempty" type:"Repeated"`
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetClientMasterNodes) String() string {
	return dara.Prettify(s)
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetClientMasterNodes) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetClientMasterNodes) GetClientMasterNode() []*DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetClientMasterNodesClientMasterNode {
	return s.ClientMasterNode
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetClientMasterNodes) SetClientMasterNode(v []*DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetClientMasterNodesClientMasterNode) *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetClientMasterNodes {
	s.ClientMasterNode = v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetClientMasterNodes) Validate() error {
	if s.ClientMasterNode != nil {
		for _, item := range s.ClientMasterNode {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetClientMasterNodesClientMasterNode struct {
	DefaultPasswd *string `json:"DefaultPasswd,omitempty" xml:"DefaultPasswd,omitempty"`
	EcsId         *string `json:"EcsId,omitempty" xml:"EcsId,omitempty"`
	EcsIp         *string `json:"EcsIp,omitempty" xml:"EcsIp,omitempty"`
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetClientMasterNodesClientMasterNode) String() string {
	return dara.Prettify(s)
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetClientMasterNodesClientMasterNode) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetClientMasterNodesClientMasterNode) GetDefaultPasswd() *string {
	return s.DefaultPasswd
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetClientMasterNodesClientMasterNode) GetEcsId() *string {
	return s.EcsId
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetClientMasterNodesClientMasterNode) GetEcsIp() *string {
	return s.EcsIp
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetClientMasterNodesClientMasterNode) SetDefaultPasswd(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetClientMasterNodesClientMasterNode {
	s.DefaultPasswd = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetClientMasterNodesClientMasterNode) SetEcsId(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetClientMasterNodesClientMasterNode {
	s.EcsId = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetClientMasterNodesClientMasterNode) SetEcsIp(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetClientMasterNodesClientMasterNode {
	s.EcsIp = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetClientMasterNodesClientMasterNode) Validate() error {
	return dara.Validate(s)
}

type DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetTags struct {
	Tag []*DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetTags) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetTags) GetTag() []*DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetTagsTag {
	return s.Tag
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetTags) SetTag(v []*DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetTagsTag) *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetTags {
	s.Tag = v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetTags) Validate() error {
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

type DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetTagsTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetTagsTag) GetKey() *string {
	return s.Key
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetTagsTag) GetValue() *string {
	return s.Value
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetTagsTag) SetKey(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetTagsTag) SetValue(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetTagsTag {
	s.Value = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetTagsTag) Validate() error {
	return dara.Validate(s)
}

type DescribeFileSystemsResponseBodyFileSystemsFileSystemOptions struct {
	EnableABE    *bool `json:"EnableABE,omitempty" xml:"EnableABE,omitempty"`
	EnableOplock *bool `json:"EnableOplock,omitempty" xml:"EnableOplock,omitempty"`
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystemOptions) String() string {
	return dara.Prettify(s)
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystemOptions) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemOptions) GetEnableABE() *bool {
	return s.EnableABE
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemOptions) GetEnableOplock() *bool {
	return s.EnableOplock
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemOptions) SetEnableABE(v bool) *DescribeFileSystemsResponseBodyFileSystemsFileSystemOptions {
	s.EnableABE = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemOptions) SetEnableOplock(v bool) *DescribeFileSystemsResponseBodyFileSystemsFileSystemOptions {
	s.EnableOplock = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemOptions) Validate() error {
	return dara.Validate(s)
}

type DescribeFileSystemsResponseBodyFileSystemsFileSystemPackages struct {
	Package []*DescribeFileSystemsResponseBodyFileSystemsFileSystemPackagesPackage `json:"Package,omitempty" xml:"Package,omitempty" type:"Repeated"`
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystemPackages) String() string {
	return dara.Prettify(s)
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystemPackages) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemPackages) GetPackage() []*DescribeFileSystemsResponseBodyFileSystemsFileSystemPackagesPackage {
	return s.Package
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemPackages) SetPackage(v []*DescribeFileSystemsResponseBodyFileSystemsFileSystemPackagesPackage) *DescribeFileSystemsResponseBodyFileSystemsFileSystemPackages {
	s.Package = v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemPackages) Validate() error {
	if s.Package != nil {
		for _, item := range s.Package {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeFileSystemsResponseBodyFileSystemsFileSystemPackagesPackage struct {
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	PackageId   *string `json:"PackageId,omitempty" xml:"PackageId,omitempty"`
	PackageType *string `json:"PackageType,omitempty" xml:"PackageType,omitempty"`
	Size        *int64  `json:"Size,omitempty" xml:"Size,omitempty"`
	StartTime   *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystemPackagesPackage) String() string {
	return dara.Prettify(s)
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystemPackagesPackage) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemPackagesPackage) GetExpiredTime() *string {
	return s.ExpiredTime
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemPackagesPackage) GetPackageId() *string {
	return s.PackageId
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemPackagesPackage) GetPackageType() *string {
	return s.PackageType
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemPackagesPackage) GetSize() *int64 {
	return s.Size
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemPackagesPackage) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemPackagesPackage) SetExpiredTime(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystemPackagesPackage {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemPackagesPackage) SetPackageId(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystemPackagesPackage {
	s.PackageId = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemPackagesPackage) SetPackageType(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystemPackagesPackage {
	s.PackageType = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemPackagesPackage) SetSize(v int64) *DescribeFileSystemsResponseBodyFileSystemsFileSystemPackagesPackage {
	s.Size = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemPackagesPackage) SetStartTime(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystemPackagesPackage {
	s.StartTime = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemPackagesPackage) Validate() error {
	return dara.Validate(s)
}

type DescribeFileSystemsResponseBodyFileSystemsFileSystemRedundancyVSwitchIds struct {
	RedundancyVSwitchId []*string `json:"RedundancyVSwitchId,omitempty" xml:"RedundancyVSwitchId,omitempty" type:"Repeated"`
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystemRedundancyVSwitchIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystemRedundancyVSwitchIds) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemRedundancyVSwitchIds) GetRedundancyVSwitchId() []*string {
	return s.RedundancyVSwitchId
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemRedundancyVSwitchIds) SetRedundancyVSwitchId(v []*string) *DescribeFileSystemsResponseBodyFileSystemsFileSystemRedundancyVSwitchIds {
	s.RedundancyVSwitchId = v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemRedundancyVSwitchIds) Validate() error {
	return dara.Validate(s)
}

type DescribeFileSystemsResponseBodyFileSystemsFileSystemSupportedFeatures struct {
	SupportedFeature []*string `json:"SupportedFeature,omitempty" xml:"SupportedFeature,omitempty" type:"Repeated"`
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystemSupportedFeatures) String() string {
	return dara.Prettify(s)
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystemSupportedFeatures) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemSupportedFeatures) GetSupportedFeature() []*string {
	return s.SupportedFeature
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemSupportedFeatures) SetSupportedFeature(v []*string) *DescribeFileSystemsResponseBodyFileSystemsFileSystemSupportedFeatures {
	s.SupportedFeature = v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemSupportedFeatures) Validate() error {
	return dara.Validate(s)
}

type DescribeFileSystemsResponseBodyFileSystemsFileSystemTags struct {
	Tag []*DescribeFileSystemsResponseBodyFileSystemsFileSystemTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystemTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystemTags) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemTags) GetTag() []*DescribeFileSystemsResponseBodyFileSystemsFileSystemTagsTag {
	return s.Tag
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemTags) SetTag(v []*DescribeFileSystemsResponseBodyFileSystemsFileSystemTagsTag) *DescribeFileSystemsResponseBodyFileSystemsFileSystemTags {
	s.Tag = v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemTags) Validate() error {
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

type DescribeFileSystemsResponseBodyFileSystemsFileSystemTagsTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystemTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystemTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemTagsTag) GetKey() *string {
	return s.Key
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemTagsTag) GetValue() *string {
	return s.Value
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemTagsTag) SetKey(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystemTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemTagsTag) SetValue(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystemTagsTag {
	s.Value = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemTagsTag) Validate() error {
	return dara.Validate(s)
}

type DescribeFileSystemsResponseBodyFileSystemsFileSystemVswIds struct {
	VswId []*string `json:"VswId,omitempty" xml:"VswId,omitempty" type:"Repeated"`
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystemVswIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystemVswIds) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemVswIds) GetVswId() []*string {
	return s.VswId
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemVswIds) SetVswId(v []*string) *DescribeFileSystemsResponseBodyFileSystemsFileSystemVswIds {
	s.VswId = v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemVswIds) Validate() error {
	return dara.Validate(s)
}
