// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeZonesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeZonesResponseBody
	GetRequestId() *string
	SetZones(v *DescribeZonesResponseBodyZones) *DescribeZonesResponseBody
	GetZones() *DescribeZonesResponseBodyZones
}

type DescribeZonesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// A70BEE5D-76D3-49FB-B58F-1F398211****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The queried zones.
	Zones *DescribeZonesResponseBodyZones `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Struct"`
}

func (s DescribeZonesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeZonesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeZonesResponseBody) GetZones() *DescribeZonesResponseBodyZones {
	return s.Zones
}

func (s *DescribeZonesResponseBody) SetRequestId(v string) *DescribeZonesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeZonesResponseBody) SetZones(v *DescribeZonesResponseBodyZones) *DescribeZonesResponseBody {
	s.Zones = v
	return s
}

func (s *DescribeZonesResponseBody) Validate() error {
	if s.Zones != nil {
		if err := s.Zones.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeZonesResponseBodyZones struct {
	Zone []*DescribeZonesResponseBodyZonesZone `json:"Zone,omitempty" xml:"Zone,omitempty" type:"Repeated"`
}

func (s DescribeZonesResponseBodyZones) String() string {
	return dara.Prettify(s)
}

func (s DescribeZonesResponseBodyZones) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBodyZones) GetZone() []*DescribeZonesResponseBodyZonesZone {
	return s.Zone
}

func (s *DescribeZonesResponseBodyZones) SetZone(v []*DescribeZonesResponseBodyZonesZone) *DescribeZonesResponseBodyZones {
	s.Zone = v
	return s
}

func (s *DescribeZonesResponseBodyZones) Validate() error {
	if s.Zone != nil {
		for _, item := range s.Zone {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeZonesResponseBodyZonesZone struct {
	// This parameter is reserved. You can ignore this parameter.
	Capacity *DescribeZonesResponseBodyZonesZoneCapacity `json:"Capacity,omitempty" xml:"Capacity,omitempty" type:"Struct"`
	// The details about file system types.
	InstanceTypes *DescribeZonesResponseBodyZonesZoneInstanceTypes `json:"InstanceTypes,omitempty" xml:"InstanceTypes,omitempty" type:"Struct"`
	// This parameter is reserved. You can ignore this parameter.
	Performance *DescribeZonesResponseBodyZonesZonePerformance `json:"Performance,omitempty" xml:"Performance,omitempty" type:"Struct"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeZonesResponseBodyZonesZone) String() string {
	return dara.Prettify(s)
}

func (s DescribeZonesResponseBodyZonesZone) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBodyZonesZone) GetCapacity() *DescribeZonesResponseBodyZonesZoneCapacity {
	return s.Capacity
}

func (s *DescribeZonesResponseBodyZonesZone) GetInstanceTypes() *DescribeZonesResponseBodyZonesZoneInstanceTypes {
	return s.InstanceTypes
}

func (s *DescribeZonesResponseBodyZonesZone) GetPerformance() *DescribeZonesResponseBodyZonesZonePerformance {
	return s.Performance
}

func (s *DescribeZonesResponseBodyZonesZone) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeZonesResponseBodyZonesZone) SetCapacity(v *DescribeZonesResponseBodyZonesZoneCapacity) *DescribeZonesResponseBodyZonesZone {
	s.Capacity = v
	return s
}

func (s *DescribeZonesResponseBodyZonesZone) SetInstanceTypes(v *DescribeZonesResponseBodyZonesZoneInstanceTypes) *DescribeZonesResponseBodyZonesZone {
	s.InstanceTypes = v
	return s
}

func (s *DescribeZonesResponseBodyZonesZone) SetPerformance(v *DescribeZonesResponseBodyZonesZonePerformance) *DescribeZonesResponseBodyZonesZone {
	s.Performance = v
	return s
}

func (s *DescribeZonesResponseBodyZonesZone) SetZoneId(v string) *DescribeZonesResponseBodyZonesZone {
	s.ZoneId = &v
	return s
}

func (s *DescribeZonesResponseBodyZonesZone) Validate() error {
	if s.Capacity != nil {
		if err := s.Capacity.Validate(); err != nil {
			return err
		}
	}
	if s.InstanceTypes != nil {
		if err := s.InstanceTypes.Validate(); err != nil {
			return err
		}
	}
	if s.Performance != nil {
		if err := s.Performance.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeZonesResponseBodyZonesZoneCapacity struct {
	Protocol []*string `json:"Protocol,omitempty" xml:"Protocol,omitempty" type:"Repeated"`
}

func (s DescribeZonesResponseBodyZonesZoneCapacity) String() string {
	return dara.Prettify(s)
}

func (s DescribeZonesResponseBodyZonesZoneCapacity) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBodyZonesZoneCapacity) GetProtocol() []*string {
	return s.Protocol
}

func (s *DescribeZonesResponseBodyZonesZoneCapacity) SetProtocol(v []*string) *DescribeZonesResponseBodyZonesZoneCapacity {
	s.Protocol = v
	return s
}

func (s *DescribeZonesResponseBodyZonesZoneCapacity) Validate() error {
	return dara.Validate(s)
}

type DescribeZonesResponseBodyZonesZoneInstanceTypes struct {
	InstanceType []*DescribeZonesResponseBodyZonesZoneInstanceTypesInstanceType `json:"InstanceType,omitempty" xml:"InstanceType,omitempty" type:"Repeated"`
}

func (s DescribeZonesResponseBodyZonesZoneInstanceTypes) String() string {
	return dara.Prettify(s)
}

func (s DescribeZonesResponseBodyZonesZoneInstanceTypes) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBodyZonesZoneInstanceTypes) GetInstanceType() []*DescribeZonesResponseBodyZonesZoneInstanceTypesInstanceType {
	return s.InstanceType
}

func (s *DescribeZonesResponseBodyZonesZoneInstanceTypes) SetInstanceType(v []*DescribeZonesResponseBodyZonesZoneInstanceTypesInstanceType) *DescribeZonesResponseBodyZonesZoneInstanceTypes {
	s.InstanceType = v
	return s
}

func (s *DescribeZonesResponseBodyZonesZoneInstanceTypes) Validate() error {
	if s.InstanceType != nil {
		for _, item := range s.InstanceType {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeZonesResponseBodyZonesZoneInstanceTypesInstanceType struct {
	// The protocol type.
	//
	// 	- If the FileSystemType parameter is set to standard, the protocol type is nfs or smb.
	//
	// 	- If the FileSystemType parameter is set to extreme, the protocol type is nfs.
	//
	// 	- If the FileSystemType parameter is set to cpfs, the protocol type is cpfs.
	//
	// example:
	//
	// nfs
	ProtocolType *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	// The storage type.
	//
	// 	- If the FileSystemType parameter is set to standard, the storage type is Performance or Capacity.
	//
	// 	- If the FileSystemType parameter is set to extreme, the storage type is standard or advance.
	//
	// 	- If the FileSystemType parameter is set to cpfs, the storage type is advance_100 (100 MB/s/TiB baseline) or advance_200 (200 MB/s/TiB baseline).
	//
	// example:
	//
	// Capacity
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
}

func (s DescribeZonesResponseBodyZonesZoneInstanceTypesInstanceType) String() string {
	return dara.Prettify(s)
}

func (s DescribeZonesResponseBodyZonesZoneInstanceTypesInstanceType) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBodyZonesZoneInstanceTypesInstanceType) GetProtocolType() *string {
	return s.ProtocolType
}

func (s *DescribeZonesResponseBodyZonesZoneInstanceTypesInstanceType) GetStorageType() *string {
	return s.StorageType
}

func (s *DescribeZonesResponseBodyZonesZoneInstanceTypesInstanceType) SetProtocolType(v string) *DescribeZonesResponseBodyZonesZoneInstanceTypesInstanceType {
	s.ProtocolType = &v
	return s
}

func (s *DescribeZonesResponseBodyZonesZoneInstanceTypesInstanceType) SetStorageType(v string) *DescribeZonesResponseBodyZonesZoneInstanceTypesInstanceType {
	s.StorageType = &v
	return s
}

func (s *DescribeZonesResponseBodyZonesZoneInstanceTypesInstanceType) Validate() error {
	return dara.Validate(s)
}

type DescribeZonesResponseBodyZonesZonePerformance struct {
	Protocol []*string `json:"Protocol,omitempty" xml:"Protocol,omitempty" type:"Repeated"`
}

func (s DescribeZonesResponseBodyZonesZonePerformance) String() string {
	return dara.Prettify(s)
}

func (s DescribeZonesResponseBodyZonesZonePerformance) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBodyZonesZonePerformance) GetProtocol() []*string {
	return s.Protocol
}

func (s *DescribeZonesResponseBodyZonesZonePerformance) SetProtocol(v []*string) *DescribeZonesResponseBodyZonesZonePerformance {
	s.Protocol = v
	return s
}

func (s *DescribeZonesResponseBodyZonesZonePerformance) Validate() error {
	return dara.Validate(s)
}
