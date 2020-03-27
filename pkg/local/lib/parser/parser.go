/*

Copyright 2017 Google Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

*/

package parser

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"

	pb "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/local/lib/proto"
)

const separator = "<:SEP:>"

// VolumeType is volume type
type VolumeType byte

var volumeTypeKeys = []byte("mMoOrRsSpviIlcVtTe")

// types
const (
	VolumeTypeMirrored                  VolumeType = 'm'
	VolumeTypeMirroredWithoutSync       VolumeType = 'M'
	VolumeTypeOrigin                    VolumeType = 'o'
	VolumeTypeOriginWithMergingSnapshot VolumeType = 'O'
	VolumeTypeRAID                      VolumeType = 'r'
	VolumeTypeRAIDWithoutSync           VolumeType = 'R'
	VolumeTypeSnapshot                  VolumeType = 's'
	VolumeTypeMergingSnapshot           VolumeType = 'S'
	VolumeTypePVMove                    VolumeType = 'p'
	VolumeTypeVirtualMirror             VolumeType = 'v'
	VolumeTypeVirtualRaidImage          VolumeType = 'i'
	VolumeTypeRaidImageOutOfSync        VolumeType = 'I'
	VolumeTypeMirrorLog                 VolumeType = 'l'
	VolumeTypeUnderConversion           VolumeType = 'c'
	VolumeTypeThin                      VolumeType = 'V'
	VolumeTypeThinPool                  VolumeType = 't'
	VolumeTypeThinPoolData              VolumeType = 'T'
	VolumeTypeRaidOrThinPoolMetadata    VolumeType = 'e'
)

func (t VolumeType) toProto() pb.LogicalVolume_Attributes_Type {
	idx := bytes.IndexByte(volumeTypeKeys, byte(t))
	if idx == -1 {
		return pb.LogicalVolume_Attributes_MALFORMED_TYPE
	}
	return pb.LogicalVolume_Attributes_Type(idx + 1)
}

// VolumePermissions is volume permissions
type VolumePermissions rune

var volumePermissonsKeys = []byte("wrR")

// permissions
const (
	VolumePermissionsWriteable          VolumePermissions = 'w'
	VolumePermissionsReadOnly           VolumePermissions = 'r'
	VolumePermissionsReadOnlyActivation VolumePermissions = 'R'
)

func (t VolumePermissions) toProto() pb.LogicalVolume_Attributes_Permissions {
	idx := bytes.IndexByte(volumePermissonsKeys, byte(t))
	if idx == -1 {
		return pb.LogicalVolume_Attributes_MALFORMED_PERMISSIONS
	}
	return pb.LogicalVolume_Attributes_Permissions(idx + 1)
}

// VolumeAllocation is volume allocation policy
type VolumeAllocation rune

var volumeAllocationKeys = []byte("acilnACILN")

// allocations
const (
	VolumeAllocationAnywhere         VolumeAllocation = 'a'
	VolumeAllocationContiguous       VolumeAllocation = 'c'
	VolumeAllocationInherited        VolumeAllocation = 'i'
	VolumeAllocationCling            VolumeAllocation = 'l'
	VolumeAllocationNormal           VolumeAllocation = 'n'
	VolumeAllocationAnywhereLocked   VolumeAllocation = 'A'
	VolumeAllocationContiguousLocked VolumeAllocation = 'C'
	VolumeAllocationInheritedLocked  VolumeAllocation = 'I'
	VolumeAllocationClingLocked      VolumeAllocation = 'L'
	VolumeAllocationNormalLocked     VolumeAllocation = 'N'
)

func (t VolumeAllocation) toProto() pb.LogicalVolume_Attributes_Allocation {
	idx := bytes.IndexByte(volumeAllocationKeys, byte(t))
	if idx == -1 {
		return pb.LogicalVolume_Attributes_MALFORMED_ALLOCATION
	}
	return pb.LogicalVolume_Attributes_Allocation(idx + 1)
}

// VolumeFixedMinor is volume fixed minor
type VolumeFixedMinor rune

// fixed minor
const (
	VolumeFixedMinorEnabled  VolumeFixedMinor = 'm'
	VolumeFixedMinorDisabled VolumeFixedMinor = '-'
)

func (t VolumeFixedMinor) toProto() bool {
	return t == VolumeFixedMinorEnabled
}

// VolumeState is volume state
type VolumeState rune

var volumeStateKeys = []byte("asISmMdi")

// states
const (
	VolumeStateActive                               VolumeState = 'a'
	VolumeStateSuspended                            VolumeState = 's'
	VolumeStateInvalidSnapshot                      VolumeState = 'I'
	VolumeStateInvalidSuspendedSnapshot             VolumeState = 'S'
	VolumeStateSnapshotMergeFailed                  VolumeState = 'm'
	VolumeStateSuspendedSnapshotMergeFailed         VolumeState = 'M'
	VolumeStateMappedDevicePresentWithoutTables     VolumeState = 'd'
	VolumeStateMappedDevicePresentWithInactiveTable VolumeState = 'i'
)

func (t VolumeState) toProto() pb.LogicalVolume_Attributes_State {
	idx := bytes.IndexByte(volumeStateKeys, byte(t))
	if idx == -1 {
		return pb.LogicalVolume_Attributes_MALFORMED_STATE
	}
	return pb.LogicalVolume_Attributes_State(idx + 1)
}

// VolumeOpen is volume open
type VolumeOpen rune

// open
const (
	VolumeOpenIsOpen    VolumeOpen = 'o'
	VolumeOpenIsNotOpen VolumeOpen = '-'
)

func (t VolumeOpen) toProto() bool {
	return t == VolumeOpenIsOpen
}

// VolumeTargetType is volume taget type
type VolumeTargetType rune

var volumeTargetTypeKeys = []byte("mrstuv")

// target type
const (
	VolumeTargetTypeMirror   VolumeTargetType = 'm'
	VolumeTargetTypeRAID     VolumeTargetType = 'r'
	VolumeTargetTypeSnapshot VolumeTargetType = 's'
	VolumeTargetTypeThin     VolumeTargetType = 't'
	VolumeTargetTypeUnknown  VolumeTargetType = 'u'
	VolumeTargetTypeVirtual  VolumeTargetType = 'v'
)

func (t VolumeTargetType) toProto() pb.LogicalVolume_Attributes_TargetType {
	idx := bytes.IndexByte(volumeTargetTypeKeys, byte(t))
	if idx == -1 {
		return pb.LogicalVolume_Attributes_MALFORMED_TARGET
	}
	return pb.LogicalVolume_Attributes_TargetType(idx + 1)
}

// VolumeZeroing is volume zeroing
type VolumeZeroing rune

// zeroing
const (
	VolumeZeroingIsZeroing    VolumeZeroing = 'z'
	VolumeZeroingIsNonZeroing VolumeZeroing = '-'
)

func (t VolumeZeroing) toProto() bool {
	return t == VolumeZeroingIsZeroing
}

// VolumeHealth is volume health
type VolumeHealth rune

// health
const (
	VolumeHealthOK              VolumeHealth = '-'
	VolumeHealthPartial         VolumeHealth = 'p'
	VolumeHealthRefreshNeeded   VolumeHealth = 'r'
	VolumeHealthMismatchesExist VolumeHealth = 'm'
	VolumeHealthWritemostly     VolumeHealth = 'w'
)

func (t VolumeHealth) toProto() pb.LogicalVolume_Attributes_Health {
	idx := bytes.IndexByte(volumeTargetTypeKeys, byte(t))
	if idx == -1 {
		return pb.LogicalVolume_Attributes_MALFORMED_HEALTH
	}
	return pb.LogicalVolume_Attributes_Health(idx + 1)
}

// VolumeActivationSkipped is volume activation
type VolumeActivationSkipped rune

// activation
const (
	VolumeActivationSkippedIsSkipped    VolumeActivationSkipped = 's'
	VolumeActivationSkippedIsNotSkipped VolumeActivationSkipped = '-'
)

func (t VolumeActivationSkipped) toProto() bool {
	return t == VolumeActivationSkippedIsSkipped
}

// LVAttributes is attributes
type LVAttributes struct {
	Type              VolumeType
	Permissions       VolumePermissions
	Allocation        VolumeAllocation
	FixedMinor        VolumeFixedMinor
	State             VolumeState
	Open              VolumeOpen
	TargetType        VolumeTargetType
	Zeroing           VolumeZeroing
	Health            VolumeHealth
	ActivationSkipped VolumeActivationSkipped
}

// ToProto returns lvm.LogicalVolume.Attributes representation of struct
func (a LVAttributes) ToProto() *pb.LogicalVolume_Attributes {
	return &pb.LogicalVolume_Attributes{
		Type:              a.Type.toProto(),
		Permissions:       a.Permissions.toProto(),
		Allocation:        a.Allocation.toProto(),
		FixedMinor:        a.FixedMinor.toProto(),
		State:             a.State.toProto(),
		Open:              a.Open.toProto(),
		TargetType:        a.TargetType.toProto(),
		Zeroing:           a.Zeroing.toProto(),
		Health:            a.Health.toProto(),
		ActivationSkipped: a.ActivationSkipped.toProto(),
	}
}

// LV is a logical volume
type LV struct {
	Name               string
	Size               uint64
	UUID               string
	Attributes         LVAttributes
	CopyPercent        string
	ActualDevMajNumber uint32
	ActualDevMinNumber uint32
	Tags               []string
}

// VG is volume group
type VG struct {
	Name     string
	Size     uint64
	FreeSize uint64
	UUID     string
	Tags     []string
	PvCount  uint64
}

// ToProto returns lvm.LogicalVolume representation of struct
func (lv LV) ToProto() *pb.LogicalVolume {
	return &pb.LogicalVolume{
		Name:                 lv.Name,
		Size:                 lv.Size,
		Uuid:                 lv.UUID,
		Attributes:           lv.Attributes.ToProto(),
		CopyPercent:          lv.CopyPercent,
		ActualDevMajorNumber: lv.ActualDevMajNumber,
		ActualDevMinorNumber: lv.ActualDevMinNumber,
		Tags:                 lv.Tags,
	}
}

// ToProto to proto
func (vg VG) ToProto() *pb.VolumeGroup {
	return &pb.VolumeGroup{
		Name:     vg.Name,
		Size:     vg.Size,
		FreeSize: vg.FreeSize,
		Uuid:     vg.UUID,
		Tags:     vg.Tags,
		PvCount:  vg.PvCount,
	}
}

func parse(line string, numComponents int) (map[string]string, error) {
	components := strings.Split(line, separator)
	if len(components) != numComponents {
		return nil, fmt.Errorf("expected %d components, got %d", numComponents, len(components))
	}

	fields := map[string]string{}
	for _, c := range components {
		idx := strings.Index(c, "=")
		if idx == -1 {
			return nil, fmt.Errorf("failed to parse component '%s'", c)
		}
		key := c[0:idx]
		value := c[idx+1:]
		if len(value) < 2 {
			return nil, fmt.Errorf("failed to parse component '%s'", c)
		}
		if value[0] != '\'' || value[len(value)-1] != '\'' {
			return nil, fmt.Errorf("failed to parse component '%s'", c)
		}
		value = value[1 : len(value)-1]
		fields[key] = value
	}

	return fields, nil
}

// ParseLV parses a line from lvs
func ParseLV(line string) (*LV, error) {
	// lvs --units=b --separator="<:SEP:>" --nosuffix --noheadings -o lv_name,lv_size,lv_uuid,lv_attr,copy_percent,lv_kernel_major,lv_kernel_minor,lv_tags --nameprefixes -a
	// todo: devices, lv_ancestors, lv_descendants, lv_major, lv_minor, mirror_log, modules, move_pv, origin, region_size
	//       seg_count, seg_size, seg_start, seg_tags, segtype, snap_percent, stripes, stripe_size
	fields, err := parse(line, 8)
	if err != nil {
		return nil, err
	}

	size, err := strconv.ParseUint(fields["LVM2_LV_SIZE"], 10, 64)
	if err != nil {
		return nil, err
	}

	kernelMajNumber, err := strconv.ParseUint(fields["LVM2_LV_KERNEL_MAJOR"], 10, 32)
	if err != nil {
		return nil, err
	}

	kernelMinNumber, err := strconv.ParseUint(fields["LVM2_LV_KERNEL_MINOR"], 10, 32)
	if err != nil {
		return nil, err
	}

	attrs, err := parseAttrs(fields["LVM2_LV_ATTR"])
	if err != nil {
		return nil, err
	}

	return &LV{
		Name:               fields["LVM2_LV_NAME"],
		Size:               size,
		UUID:               fields["LVM2_LV_UUID"],
		Attributes:         *attrs,
		CopyPercent:        fields["LVM2_COPY_PERCENT"],
		ActualDevMajNumber: uint32(kernelMajNumber),
		ActualDevMinNumber: uint32(kernelMinNumber),
		Tags:               strings.Split(fields["LVM2_LV_TAGS"], ","),
	}, nil
}

// ParseVG parse volume group
func ParseVG(line string) (*VG, error) {
	// vgs --units=b --separator="<:SEP:>" --nosuffix --noheadings -o vg_name,vg_size,vg_free,vg_uuid,vg_tags,pv_count --nameprefixes -a
	fields, err := parse(line, 6)
	if err != nil {
		return nil, err
	}

	size, err := strconv.ParseUint(fields["LVM2_VG_SIZE"], 10, 64)
	if err != nil {
		return nil, err
	}

	freeSize, err := strconv.ParseUint(fields["LVM2_VG_FREE"], 10, 64)
	if err != nil {
		return nil, err
	}
	pvCount, err := strconv.ParseUint(fields["LVM2_PV_COUNT"], 10, 64)
	if err != nil {
		return nil, err
	}
	return &VG{
		Name:     fields["LVM2_VG_NAME"],
		Size:     size,
		FreeSize: freeSize,
		UUID:     fields["LVM2_VG_UUID"],
		Tags:     strings.Split(fields["LVM2_VG_TAGS"], ","),
		PvCount:  pvCount,
	}, nil
}

func parseAttrs(attrs string) (*LVAttributes, error) {
	if len(attrs) != 10 {
		return nil, fmt.Errorf("incorrect attrs block size, expected 10, got %d in %s", len(attrs), attrs)
	}

	ret := &LVAttributes{}
	ret.Type = VolumeType(attrs[0])
	ret.Permissions = VolumePermissions(attrs[1])
	ret.Allocation = VolumeAllocation(attrs[2])
	ret.FixedMinor = VolumeFixedMinor(attrs[3])
	ret.State = VolumeState(attrs[4])
	ret.Open = VolumeOpen(attrs[5])
	ret.TargetType = VolumeTargetType(attrs[6])
	ret.Zeroing = VolumeZeroing(attrs[7])
	ret.Health = VolumeHealth(attrs[8])
	ret.ActivationSkipped = VolumeActivationSkipped(attrs[9])

	return ret, nil
}
