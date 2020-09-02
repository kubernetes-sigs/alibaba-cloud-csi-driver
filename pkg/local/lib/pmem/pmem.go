package pmem

import (
	"encoding/json"
	"errors"
	"fmt"
	pb "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/local/lib/proto"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	log "github.com/sirupsen/logrus"
	"os/exec"
	"path/filepath"
	"strings"
)

const (
	// PmemVolumeGroupNameRegion0 tag
	PmemVolumeGroupNameRegion0 = "pmemvgregion0"
	// PmemVolumeGroupNameRegion1 tag
	PmemVolumeGroupNameRegion1 = "pmemvgregion1"
	// PmemRegionNameDefault tag
	PmemRegionNameDefault = "region0"
)

// PmemRegions list all regions
type PmemRegions struct {
	Regions []PmemRegion `json:"regions"`
}

// PmemRegion define on pmem region
type PmemRegion struct {
	Dev               string          `json:"dev"`
	Size              int64           `json:"size,omitempty"`
	AvailableSize     int64           `json:"available_size,omitempty"`
	MaxAvailableExent int64           `json:"max_available_extent,omitempty"`
	RegionType        string          `json:"type,omitempty"`
	IsetId            int64           `json:"iset_id,omitempty"`
	PersistenceDomain string          `json:"persistence_domain,omitempty"`
	Namespaces        []PmemNameSpace `json:"namespaces,omitempty"`
}

// PmemNameSpace define one pmem namespaces
type PmemNameSpace struct {
	Dev        string `json:"dev,omitempty"`
	Mode       string `json:"mode,omitempty"`
	MapType    string `json:"map,omitempty"`
	Size       int64  `json:"size,omitempty"`
	Uuid       string `json:"uuid,omitempty"`
	SectorSize int64  `json:"sector_size,omitempty"`
	Align      int64  `json:"align,omitempty"`
	BlockDev   string `json:"blockdev,omitempty"`
	CharDev    string `json:"chardev,omitempty"`
	Name       string `json:"name,omitempty"`
}

const (
	// NsenterCmd is the nsenter command
	NsenterCmd = "/nsenter --mount=/proc/1/ns/mnt --ipc=/proc/1/ns/ipc --net=/proc/1/ns/net --uts=/proc/1/ns/uts "
)

// ErrParse is an error that is returned when parse operation fails
var ErrParse = errors.New("Cannot parse output of blkid")

func MaintainPMEM(pmemType string) error {
	regions, err := getRegions()
	if err != nil {
		log.Errorf("Get pmem regions error: %v", err)
		return err
	}
	if regions == nil || len(regions.Regions) == 0 {
		log.Infof("No pmem Regions found, exit")
		return nil
	}
	log.Infof("Get Regions Info: %v", regions)

	if pmemType == "lvm" {
		if err := MaintainLVM(regions); err != nil {
			log.Errorf("MaintainLVM got error: %v", err)
			return err
		}
		log.Infof("MaintainLVM Successful")
	} else if pmemType == "direct" {
		if err := MaintainDirect(regions); err != nil {
			log.Errorf("MaintainDirect got error: %v", err)
			return err
		}
		log.Infof("MaintainDirect Successful")
	} else {
		log.Warnf("Set pmem type %s is not supported", pmemType)
	}

	return nil
}

func MaintainDirect(regions *PmemRegions) error {
	return nil
}

func MaintainLVM(regions *PmemRegions) error {
	// Create Namespaces if not exist
	for _, region := range regions.Regions {
		if len(region.Namespaces) == 0 {
			if err := createNameSpace(region.Dev); err != nil {
				log.Errorf("Create NameSpace error for region: %s", region.Dev)
				return errors.New("Create NameSpace error for region: " + region.Dev)
			}
		}
	}
	// Get Regions
	regions, err := getRegions()
	if err != nil {
		log.Errorf("Get regions error after Check Namespace: %v", err)
		return err
	}

	// Create VolumeGroup
	for _, region := range regions.Regions {
		deviceList := []string{}
		needCreate := true
		for _, namespace := range region.Namespaces {
			devicePath := filepath.Join("/dev", namespace.BlockDev)
			if namespace.Mode == "devdax" {
				devicePath = filepath.Join("/dev", namespace.CharDev)
			}
			if checkNameSpaceUsed(devicePath) {
				log.Warnf("NameSpace heen used: %v, %s", namespace, devicePath)
				needCreate = false
				break
			}
			deviceList = append(deviceList, devicePath)
		}

		if len(deviceList) > 0 && needCreate {
			if region.Dev == "region0" {
				if err := createPmemVG(deviceList, PmemVolumeGroupNameRegion0); err != nil {
					return err
				}

			} else if region.Dev == "region1" {
				if err := createPmemVG(deviceList, PmemVolumeGroupNameRegion1); err != nil {
					return err
				}
			} else {
				return errors.New("Create volumegroup with not supported region " + region.Dev)
			}
		}
	}
	return nil
}

func createNameSpace(region string) error {
	createCmd := fmt.Sprintf("%s ndctl create-namespace -r %s", NsenterCmd, region)
	_, err := utils.Run(createCmd)
	if err != nil {
		log.Errorf("Create NameSpace for region %s error: %v", region, err)
		return err
	}
	log.Infof("Create NameSpace for region %s successful", region)
	return nil
}

// device used in pv
// device used in block
func checkNameSpaceUsed(devicePath string) bool {
	pvCheckCmd := fmt.Sprintf("%s pvs %s 2>&1 | grep -v \"Failed to \" | grep /dev | awk '{print $2}' | wc -l", NsenterCmd, devicePath)
	out, err := utils.Run(pvCheckCmd)
	if err == nil && strings.TrimSpace(out) != "0" {
		log.Infof("NameSpace %s used for pv", devicePath)
		return true
	}

	out, err = checkFSType(devicePath)
	if err == nil && strings.TrimSpace(out) != "" {
		log.Infof("NameSpace %s format as %s", devicePath, out)
		return true
	}
	return false
}

func createPmemVG(deviceList []string, vgName string) error {
	localDeviceStr := strings.Join(deviceList, " ")
	vgAddCmd := fmt.Sprintf("%s vgcreate --force %s %s", NsenterCmd, vgName, localDeviceStr)
	_, err := utils.Run(vgAddCmd)
	if err != nil {
		log.Errorf("Create VG (%v) with PV (%v) error: %s", vgName, localDeviceStr, err.Error())
		return err
	}

	log.Infof("Successful add Local Disks to VG (%s): %s", vgName, localDeviceStr)
	return nil
}

func getRegions() (*PmemRegions, error) {
	regions := &PmemRegions{}
	getRegionCmd := fmt.Sprintf("%s ndctl list -RN", NsenterCmd)
	regionOut, err := utils.Run(getRegionCmd)
	if err != nil {
		return regions, err
	}
	err = json.Unmarshal(([]byte)(regionOut), regions)
	if err != nil {
		if strings.HasPrefix(regionOut, "[") {
			regionList := []PmemRegion{}
			err = json.Unmarshal(([]byte)(regionOut), &regionList)
			if err != nil {
				return regions, err
			}
			regions.Regions = regionList
		} else {
			return regions, err
		}
	}

	return regions, nil
}

func GetSetCapacity(ns *PmemNameSpace) int64 {
	expect := (ns.Size + ns.Align) * 4096 / 4032
	return expect
}

func GetNameSpace(namespaceName string) (*PmemNameSpace, error) {
	namespace := &PmemNameSpace{}
	namespaceList := []*PmemNameSpace{}
	getRegionCmd := fmt.Sprintf("%s ndctl list -n %s", NsenterCmd, namespaceName)
	regionOut, err := utils.Run(getRegionCmd)
	if err != nil {
		return namespace, err
	}
	err = json.Unmarshal(([]byte)(regionOut), &namespaceList)
	if err != nil {
		return namespace, err
	}
	if len(namespaceList) == 1 {
		return namespaceList[0], nil
	}
	return namespace, fmt.Errorf("namespace found error")
}

func (pns *PmemNameSpace) ToProto() *pb.NameSpace {
	new := &pb.NameSpace{}
	new.CharDev = pns.CharDev
	new.Name = pns.Name
	new.Dev = pns.Dev
	new.Mode = pns.Mode
	new.Size = pns.Size
	new.Uuid = pns.Uuid
	new.Align = pns.Align
	new.MapType = pns.MapType
	new.SectorSize = pns.SectorSize
	return new
}

func ListNameSpace() ([]*pb.NameSpace, error) {
	regions, err := getRegions()
	if err != nil {
		return nil, err
	}

	namespaces := []*pb.NameSpace{}
	for _, region := range regions.Regions {
		for _, ns := range region.Namespaces {
			namespaces = append(namespaces, ns.ToProto())
		}
	}
	return namespaces, nil
}

func checkFSType(devicePath string) (string, error) {
	// We use `file -bsL` to determine whether any filesystem type is detected.
	// If a filesystem is detected (ie., the output is not "data", we use
	// `blkid` to determine what the filesystem is. We use `blkid` as `file`
	// has inconvenient output.
	// We do *not* use `lsblk` as that requires udev to be up-to-date which
	// is often not the case when a device is erased using `dd`.
	output, err := exec.Command("file", "-bsL", devicePath).CombinedOutput()
	if err != nil {
		return "", err
	}
	if strings.TrimSpace(string(output)) == "data" {
		return "", nil
	}
	output, err = exec.Command("blkid", "-c", "/dev/null", "-o", "export", devicePath).CombinedOutput()
	if err != nil {
		return "", err
	}

	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		fields := strings.Split(strings.TrimSpace(line), "=")
		if len(fields) != 2 {
			return "", ErrParse
		}
		if fields[0] == "TYPE" {
			return fields[1], nil
		}
	}
	return "", ErrParse
}
