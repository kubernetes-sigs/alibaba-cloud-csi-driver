package lvm

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

// Control verbose output of all LVM CLI commands
var Verbose bool

// isInsufficientSpace returns true if the error is due to insufficient space
func isInsufficientSpace(err error) bool {
	return strings.Contains(err.Error(), "insufficient free space")
}

// MaxSize states that all available space should be used by the
// create operation.
const MaxSize uint64 = 0

type simpleError string

func (s simpleError) Error() string { return string(s) }

const ErrNoSpace = simpleError("lvm: not enough free space")
const ErrPhysicalVolumeNotFound = simpleError("lvm: physical volume not found")
const ErrVolumeGroupNotFound = simpleError("lvm: volume group not found")

var vgnameRegexp = regexp.MustCompile("^[A-Za-z0-9_+.][A-Za-z0-9_+.-]*$")

const ErrInvalidVGName = simpleError("lvm: Name contains invalid character, valid set includes: [A-Za-z0-9_+.-]")

var lvnameRegexp = regexp.MustCompile("^[A-Za-z0-9_+.][A-Za-z0-9_+.-]*$")

const ErrInvalidLVName = simpleError("lvm: Name contains invalid character, valid set includes: [A-Za-z0-9_+.-]")

var tagRegexp = regexp.MustCompile("^[A-Za-z0-9_+.][A-Za-z0-9_+.-]*$")

const ErrTagInvalidLength = simpleError("lvm: Tag length must be between 1 and 1024 characters")
const ErrTagHasInvalidChars = simpleError("lvm: Tag must consist of only [A-Za-z0-9_+.-] and cannot start with a '-'")

type PhysicalVolume struct {
	dev string
}

// Remove removes the physical volume.
func (pv *PhysicalVolume) Remove() error {
	if err := run("pvremove", nil, pv.dev); err != nil {
		return err
	}
	return nil
}

// Check runs the pvck command on the physical volume.
func (pv *PhysicalVolume) Check() error {
	if err := run("pvck", nil, pv.dev); err != nil {
		return err
	}
	return nil
}

type VolumeGroup struct {
	name string
}

func (vg *VolumeGroup) Name() string {
	return vg.name
}

// Check runs the vgck command on the volume group.
func (vg *VolumeGroup) Check() error {
	if err := run("vgck", nil, vg.name); err != nil {
		return err
	}
	return nil
}

// BytesTotal returns the current size in bytes of the volume group.
func (vg *VolumeGroup) BytesTotal() (uint64, error) {
	result := new(vgsOutput)
	if err := run("vgs", result, "--options=vg_size", vg.name); err != nil {
		if IsVolumeGroupNotFound(err) {
			return 0, ErrVolumeGroupNotFound
		}
		return 0, err
	}
	for _, report := range result.Report {
		for _, vg := range report.Vg {
			return vg.VgSize, nil
		}
	}
	return 0, ErrVolumeGroupNotFound
}

// BytesFree returns the unallocated space in bytes of the volume group.
func (vg *VolumeGroup) BytesFree() (uint64, error) {
	result := new(vgsOutput)
	if err := run("vgs", result, "--options=vg_free", vg.name); err != nil {
		if IsVolumeGroupNotFound(err) {
			return 0, ErrVolumeGroupNotFound
		}
		return 0, err
	}
	for _, report := range result.Report {
		for _, vg := range report.Vg {
			return vg.VgFree, nil
		}
	}
	return 0, ErrVolumeGroupNotFound
}

// ExtentSize returns the size in bytes of a single extent.
func (vg *VolumeGroup) ExtentSize() (uint64, error) {
	result := new(vgsOutput)
	if err := run("vgs", result, "--options=vg_extent_size", vg.name); err != nil {
		if IsVolumeGroupNotFound(err) {
			return 0, ErrVolumeGroupNotFound
		}
		return 0, err
	}
	for _, report := range result.Report {
		for _, vg := range report.Vg {
			return vg.VgExtentSize, nil
		}
	}
	return 0, ErrVolumeGroupNotFound
}

// ExtentCount returns the number of extents.
func (vg *VolumeGroup) ExtentCount() (uint64, error) {
	result := new(vgsOutput)
	if err := run("vgs", result, "--options=vg_extent_count", vg.name); err != nil {
		if IsVolumeGroupNotFound(err) {
			return 0, ErrVolumeGroupNotFound
		}
		return 0, err
	}
	for _, report := range result.Report {
		for _, vg := range report.Vg {
			return vg.VgExtentCount, nil
		}
	}
	return 0, ErrVolumeGroupNotFound
}

// ExtentFreeCount returns the number of free extents.
func (vg *VolumeGroup) ExtentFreeCount() (uint64, error) {
	result := new(vgsOutput)
	if err := run("vgs", result, "--options=vg_free_count", vg.name); err != nil {
		if IsVolumeGroupNotFound(err) {
			return 0, ErrVolumeGroupNotFound
		}
		return 0, err
	}
	for _, report := range result.Report {
		for _, vg := range report.Vg {
			return vg.VgFreeExtentCount, nil
		}
	}
	return 0, ErrVolumeGroupNotFound
}

// CreateLogicalVolume creates a logical volume of the given device
// and size.
//
// The actual size may be larger than asked for as the smallest
// increment is the size of an extent on the volume group in question.
//
// If sizeInBytes is zero the entire available space is allocated.
func (vg *VolumeGroup) CreateLogicalVolume(name string, sizeInBytes uint64, tags []string) (*LogicalVolume, error) {
	if err := ValidateLogicalVolumeName(name); err != nil {
		return nil, err
	}
	// Validate the tag.
	var args []string
	for _, tag := range tags {
		if tag != "" {
			if err := ValidateTag(tag); err != nil {
				return nil, err
			}
			args = append(args, "--add-tag="+tag)
		}
	}
	args = append(args, fmt.Sprintf("--size=%db", sizeInBytes))
	args = append(args, "--name="+name)
	args = append(args, vg.name)
	if err := run("lvcreate", nil, args...); err != nil {
		if isInsufficientSpace(err) {
			return nil, ErrNoSpace
		}
		return nil, err
	}
	return &LogicalVolume{name, sizeInBytes, vg}, nil
}

// ValidateLogicalVolumeName validates a volume group name. A valid volume
// group name can consist of a limited range of characters only. The allowed
// characters are [A-Za-z0-9_+.-].
func ValidateLogicalVolumeName(name string) error {
	if !lvnameRegexp.MatchString(name) {
		return ErrInvalidLVName
	}
	return nil
}

const ErrLogicalVolumeNotFound = simpleError("lvm: logical volume not found")

type lvsOutput struct {
	Report []struct {
		Lv []struct {
			Name   string `json:"lv_name"`
			VgName string `json:"vg_name"`
			LvPath string `json:"lv_path"`
			LvSize uint64 `json:"lv_size,string"`
			LvTags string `json:"lv_tags"`
		} `json:"lv"`
	} `json:"report"`
}

func IsLogicalVolumeNotFound(err error) bool {
	const prefix = "Failed to find logical volume"
	lines := strings.Split(err.Error(), "\n")
	if len(lines) == 0 {
		return false
	}
	for _, line := range lines {
		if strings.HasPrefix(line, prefix) {
			return true
		}
	}
	return false
}

// LookupLogicalVolume looks up the logical volume in the volume group
// with the given name.
func (vg *VolumeGroup) LookupLogicalVolume(name string) (*LogicalVolume, error) {
	result := new(lvsOutput)
	if err := run("lvs", result, "--options=lv_name,lv_size,vg_name", vg.Name()); err != nil {
		if IsLogicalVolumeNotFound(err) {
			return nil, ErrLogicalVolumeNotFound
		}
		return nil, err
	}
	for _, report := range result.Report {
		for _, lv := range report.Lv {
			if lv.VgName != vg.Name() {
				continue
			}
			if lv.Name != name {
				continue
			}
			return &LogicalVolume{lv.Name, lv.LvSize, vg}, nil
		}
	}
	return nil, ErrLogicalVolumeNotFound
}

// ListLogicalVolumes returns the names of the logical volumes in this volume group.
func (vg *VolumeGroup) ListLogicalVolumeNames() ([]string, error) {
	var names []string
	result := new(lvsOutput)
	if err := run("lvs", result, "--options=lv_name,vg_name", vg.name); err != nil {
		return nil, err
	}
	for _, report := range result.Report {
		for _, lv := range report.Lv {
			if lv.VgName == vg.name {
				names = append(names, lv.Name)
			}
		}
	}
	return names, nil
}

func IsPhysicalVolumeNotFound(err error) bool {
	return isPhysicalVolumeNotFound(err) ||
		isNoPhysicalVolumeLabel(err)
}

func isPhysicalVolumeNotFound(err error) bool {
	const prefix = "Failed to find device"
	lines := strings.Split(err.Error(), "\n")
	if len(lines) == 0 {
		return false
	}
	for _, line := range lines {
		if strings.HasPrefix(line, prefix) {
			return true
		}
	}
	return false
}

func isNoPhysicalVolumeLabel(err error) bool {
	const prefix = "No physical volume label read from"
	lines := strings.Split(err.Error(), "\n")
	if len(lines) == 0 {
		return false
	}
	for _, line := range lines {
		if strings.HasPrefix(line, prefix) {
			return true
		}
	}
	return false
}

func IsVolumeGroupNotFound(err error) bool {
	const prefix = "Volume group"
	const suffix = "not found"
	lines := strings.Split(err.Error(), "\n")
	if len(lines) == 0 {
		return false
	}
	for _, line := range lines {
		if strings.HasPrefix(line, prefix) && strings.HasSuffix(line, suffix) {
			return true
		}
	}
	return false
}

// ListPhysicalVolumeNames returns the names of the physical volumes in this volume group.
func (vg *VolumeGroup) ListPhysicalVolumeNames() ([]string, error) {
	var names []string
	result := new(pvsOutput)
	if err := run("pvs", result, "--options=pv_name,vg_name"); err != nil {
		return nil, err
	}
	for _, report := range result.Report {
		for _, pv := range report.Pv {
			if pv.VgName == vg.name {
				names = append(names, pv.Name)
			}
		}
	}
	return names, nil
}

// Tags returns the volume group tags.
func (vg *VolumeGroup) Tags() ([]string, error) {
	result := new(vgsOutput)
	if err := run("vgs", result, "--options=vg_tags", vg.name); err != nil {
		if IsVolumeGroupNotFound(err) {
			return nil, ErrVolumeGroupNotFound
		}

		return nil, err
	}
	for _, report := range result.Report {
		for _, vg := range report.Vg {
			var tags []string
			for _, tag := range strings.Split(vg.VgTags, ",") {
				tag = strings.TrimSpace(tag)
				if tag != "" {
					tags = append(tags, tag)
				}
			}
			return tags, nil
		}
	}
	return nil, ErrVolumeGroupNotFound
}

// Remove removes the volume group from disk.
func (vg *VolumeGroup) Remove() error {
	if err := run("vgremove", nil, "-f", vg.name); err != nil {
		return err
	}
	return nil
}

type LogicalVolume struct {
	name        string
	sizeInBytes uint64
	vg          *VolumeGroup
}

func (lv *LogicalVolume) Name() string {
	return lv.name
}

func (lv *LogicalVolume) SizeInBytes() uint64 {
	return lv.sizeInBytes
}

// Path returns the device path for the logical volume.
func (lv *LogicalVolume) Path() (string, error) {
	result := new(lvsOutput)
	if err := run("lvs", result, "--options=lv_path", lv.vg.name+"/"+lv.name); err != nil {
		if IsLogicalVolumeNotFound(err) {
			return "", ErrLogicalVolumeNotFound
		}
		return "", err
	}
	for _, report := range result.Report {
		for _, lv := range report.Lv {
			return lv.LvPath, nil
		}
	}
	return "", ErrLogicalVolumeNotFound
}

// Tags returns the volume group tags.
func (lv *LogicalVolume) Tags() ([]string, error) {
	result := new(lvsOutput)
	if err := run("lvs", result, "--options=lv_tags", lv.vg.name+"/"+lv.name); err != nil {
		if IsLogicalVolumeNotFound(err) {
			return nil, ErrLogicalVolumeNotFound
		}
		return nil, err
	}
	for _, report := range result.Report {
		for _, lv := range report.Lv {
			var tags []string
			for _, tag := range strings.Split(lv.LvTags, ",") {
				tag = strings.TrimSpace(tag)
				if tag != "" {
					tags = append(tags, tag)
				}
			}
			return tags, nil
		}
	}
	return nil, ErrLogicalVolumeNotFound
}

func (lv *LogicalVolume) Remove() error {
	if err := run("lvremove", nil, "-f", lv.vg.name+"/"+lv.name); err != nil {
		return err
	}
	return nil
}

// PVScan runs the `pvscan --cache <dev>` command. It scans for the
// device at `dev` and adds it to the LVM metadata cache if `lvmetad`
// is running. If `dev` is an empty string, it scans all devices.
func PVScan(dev string) error {
	args := []string{"--cache"}
	if dev != "" {
		args = append(args, dev)
	}
	return run("pvscan", nil, args...)
}

// VGScan runs the `vgscan --cache <name>` command. It scans for the
// volume group and adds it to the LVM metadata cache if `lvmetad`
// is running. If `name` is an empty string, it scans all volume groups.
func VGScan(name string) error {
	args := []string{"--cache"}
	if name != "" {
		args = append(args, name)
	}
	return run("vgscan", nil, args...)
}

// CreateVolumeGroup creates a new volume group.
func CreateVolumeGroup(
	name string,
	pvs []*PhysicalVolume,
	tags []string) (*VolumeGroup, error) {
	var args []string
	if err := ValidateVolumeGroupName(name); err != nil {
		return nil, err
	}
	for _, tag := range tags {
		if tag != "" {
			if err := ValidateTag(tag); err != nil {
				return nil, err
			}
			args = append(args, "--add-tag="+tag)
		}
	}
	args = append(args, name)
	for _, pv := range pvs {
		args = append(args, pv.dev)
	}
	if err := run("vgcreate", nil, args...); err != nil {
		return nil, err
	}
	// Perform a best-effort scan to trigger a lvmetad cache refresh.
	// We ignore errors as for better or worse, the volume group now exists.
	// Without this lvmetad can fail to pickup newly created volume groups.
	// See https://bugzilla.redhat.com/show_bug.cgi?id=837599
	if err := PVScan(""); err != nil {
		log.Printf("error during pvscan: %v", err)
	}
	if err := VGScan(""); err != nil {
		log.Printf("error during vgscan: %v", err)
	}
	return &VolumeGroup{name}, nil
}

// ValidateVolumeGroupName validates a volume group name. A valid volume group
// name can consist of a limited range of characters only. The allowed
// characters are [A-Za-z0-9_+.-].
func ValidateVolumeGroupName(name string) error {
	if !vgnameRegexp.MatchString(name) {
		return ErrInvalidVGName
	}
	return nil
}

// ValidateTag validates a tag. LVM tags are strings of up to 1024
// characters. LVM tags cannot start with a hyphen. A valid tag can consist of
// a limited range of characters only. The allowed characters are
// [A-Za-z0-9_+.-]. As of the Red Hat Enterprise Linux 6.1 release, the list of
// allowed characters was extended, and tags can contain the /, =, !, :, #, and
// & characters.
// See https://access.redhat.com/documentation/en-us/red_hat_enterprise_linux/7/html/logical_volume_manager_administration/lvm_tags
func ValidateTag(tag string) error {
	if len(tag) > 1024 {
		return ErrTagInvalidLength
	}
	if !tagRegexp.MatchString(tag) {
		return ErrTagHasInvalidChars
	}
	return nil
}

type vgsOutput struct {
	Report []struct {
		Vg []struct {
			Name              string `json:"vg_name"`
			UUID              string `json:"vg_uuid"`
			VgSize            uint64 `json:"vg_size,string"`
			VgFree            uint64 `json:"vg_free,string"`
			VgExtentSize      uint64 `json:"vg_extent_size,string"`
			VgExtentCount     uint64 `json:"vg_extent_count,string"`
			VgFreeExtentCount uint64 `json:"vg_free_count,string"`
			VgTags            string `json:"vg_tags"`
		} `json:"vg"`
	} `json:"report"`
}

// LookupVolumeGroup returns the volume group with the given name.
func LookupVolumeGroup(name string) (*VolumeGroup, error) {
	result := new(vgsOutput)
	if err := run("vgs", result, "--options=vg_name", name); err != nil {
		if IsVolumeGroupNotFound(err) {
			return nil, ErrVolumeGroupNotFound
		}
		return nil, err
	}
	for _, report := range result.Report {
		for _, vg := range report.Vg {
			return &VolumeGroup{vg.Name}, nil
		}
	}
	return nil, ErrVolumeGroupNotFound
}

// ListVolumeGroupNames returns the names of the list of volume groups. This
// does not normally scan for devices. To scan for devices, use the `Scan()`
// function.
func ListVolumeGroupNames() ([]string, error) {
	result := new(vgsOutput)
	if err := run("vgs", result); err != nil {
		return nil, err
	}
	var names []string
	for _, report := range result.Report {
		for _, vg := range report.Vg {
			names = append(names, vg.Name)
		}
	}
	return names, nil
}

// ListVolumeGroupUUIDs returns the UUIDs of the list of volume groups. This
// does not normally scan for devices. To scan for devices, use the `Scan()`
// function.
func ListVolumeGroupUUIDs() ([]string, error) {
	result := new(vgsOutput)
	if err := run("vgs", result, "--options=vg_uuid"); err != nil {
		return nil, err
	}
	var uuids []string
	for _, report := range result.Report {
		for _, vg := range report.Vg {
			uuids = append(uuids, vg.UUID)
		}
	}
	return uuids, nil
}

// CreatePhysicalVolume creates a physical volume of the given device.
func CreatePhysicalVolume(dev string) (*PhysicalVolume, error) {
	if err := run("pvcreate", nil, dev); err != nil {
		return nil, fmt.Errorf("lvm: CreatePhysicalVolume: %v", err)
	}
	return &PhysicalVolume{dev}, nil
}

type pvsOutput struct {
	Report []struct {
		Pv []struct {
			Name   string `json:"pv_name"`
			VgName string `json:"vg_name"`
		} `json:"pv"`
	} `json:"report"`
}

// ListPhysicalVolumes lists all physical volumes.
func ListPhysicalVolumes() ([]*PhysicalVolume, error) {
	result := new(pvsOutput)
	if err := run("pvs", result); err != nil {
		return nil, err
	}
	var pvs []*PhysicalVolume
	for _, report := range result.Report {
		for _, pv := range report.Pv {
			pvs = append(pvs, &PhysicalVolume{pv.Name})
		}
	}
	return pvs, nil
}

// LookupPhysicalVolume returns a physical volume with the given name.
func LookupPhysicalVolume(name string) (*PhysicalVolume, error) {
	result := new(pvsOutput)
	if err := run("pvs", result, "--options=pv_name", name); err != nil {
		if IsPhysicalVolumeNotFound(err) {
			return nil, ErrPhysicalVolumeNotFound
		}
		return nil, err
	}
	for _, report := range result.Report {
		for _, pv := range report.Pv {
			return &PhysicalVolume{pv.Name}, nil
		}
	}
	return nil, ErrPhysicalVolumeNotFound
}

// Extent sizing for linear logical volumes:
// https://github.com/Jajcus/lvm2/blob/266d6564d7a72fcff5b25367b7a95424ccf8089e/lib/metadata/metadata.c#L983

func run(cmd string, v interface{}, extraArgs ...string) error {
	var args []string
	if v != nil {
		args = append(args, "--reportformat=json")
		args = append(args, "--units=b")
		args = append(args, "--nosuffix")
	}
	args = append(args, extraArgs...)
	c := exec.Command(cmd, args...)
	log.Printf("Executing: %v", c)
	stdout, stderr := new(bytes.Buffer), new(bytes.Buffer)
	c.Stdout = stdout
	c.Stderr = stderr
	if err := c.Run(); err != nil {
		errstr := ignoreWarnings(stderr.String())
		log.Print("stdout: " + stdout.String())
		log.Print("stderr: " + errstr)
		return errors.New(errstr)
	}
	stdoutbuf := stdout.Bytes()
	stderrbuf := stderr.Bytes()
	errstr := ignoreWarnings(string(stderrbuf))
	//log.Printf("stdout: " + string(stdoutbuf))
	if errstr != "" {
		log.Printf("stderr: " + errstr)
	}
	if v != nil {
		if err := json.Unmarshal(stdoutbuf, v); err != nil {
			return fmt.Errorf("%v: [%v]", err, string(stdoutbuf))
		}
	}
	return nil
}

func ignoreWarnings(str string) string {
	lines := strings.Split(str, "\n")
	result := make([]string, 0, len(lines))
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "WARNING") {
			log.Printf(line)
			continue
		}
		// Ignore warnings of the kind:
		// "File descriptor 13 (pipe:[120900]) leaked on vgs invocation. Parent PID 2: ./csilvm"
		// For some reason lvm2 decided to complain if there are open file descriptors
		// that it didn't create when it exits. This doesn't play nice with the fact
		// that csilvm gets launched by e.g., mesos-agent.
		if strings.HasPrefix(line, "File descriptor") {
			log.Printf(line)
			continue
		}
		result = append(result, line)
	}
	return strings.TrimSpace(strings.Join(result, "\n"))
}
