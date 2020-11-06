package lib

import (
	"errors"
	"fmt"
	utilexec "k8s.io/utils/exec"
	k8smount "k8s.io/utils/mount"
	"path/filepath"
	"strings"

	log "github.com/sirupsen/logrus"
)

const (
	// PmemVolumeGroupNameRegion0 tag
	PmemVolumeGroupNameRegion0 = "pmemvgregion0"
	// PmemVolumeGroupNameRegion1 tag
	PmemVolumeGroupNameRegion1 = "pmemvgregion1"
	// PmemRegionNameDefault tag
	PmemRegionNameDefault = "region0"
	// fsckErrorsCorrected tag
	fsckErrorsCorrected = 1
	// fsckErrorsUncorrected tag
	fsckErrorsUncorrected = 4
)

// PmemRegions list all regions
type PmemRegions struct {
	Regions []PmemRegion `json:"regions"`
}

// DaxctrlMem list all mems
type DaxctrlMem struct {
	Chardev    string `json:"chardev"`
	Size       int64  `json:"size"`
	TargetNode int    `json:"target_node"`
	Mode       string `json:"mode"`
	Movable    bool   `json:"movable"`
}

// PmemRegion define on pmem region
type PmemRegion struct {
	Dev               string          `json:"dev"`
	Size              int64           `json:"size,omitempty"`
	AvailableSize     int64           `json:"available_size,omitempty"`
	MaxAvailableExent int64           `json:"max_available_extent,omitempty"`
	RegionType        string          `json:"type,omitempty"`
	IsetID            int64           `json:"iset_id,omitempty"`
	PersistenceDomain string          `json:"persistence_domain,omitempty"`
	Namespaces        []PmemNameSpace `json:"namespaces,omitempty"`
}

// PmemNameSpace define one pmem namespaces
type PmemNameSpace struct {
	Dev        string `json:"dev,omitempty"`
	Mode       string `json:"mode,omitempty"`
	MapType    string `json:"map,omitempty"`
	Size       int64  `json:"size,omitempty"`
	UUID       string `json:"uuid,omitempty"`
	SectorSize int64  `json:"sectorsize,omitempty"`
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

// MaintainPMEM build pmem device
func MaintainPMEM(pmemType string, mounter k8smount.Interface) error {
	regions, err := GetRegions()
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
	} else if pmemType == "kmem" {
		if err := MaintainKMEM(regions); err != nil {
			log.Errorf("MaintainKMEM got error: %v", err)
			return err
		}
	} else if pmemType == "projquota" {
		if err := MaintainProjQuota(regions, mounter); err != nil {
			log.Errorf("MaintainProjQuota got error: %v", err)
			return err
		}
	} else {
		log.Warnf("Set pmem type %s is not supported", pmemType)
	}

	return nil
}

// MaintainDirect direct pmem
func MaintainDirect(regions *PmemRegions) error {
	return nil
}

// MaintainKMEM create kmem type namespace
func MaintainKMEM(regions *PmemRegions) error {
	for _, region := range regions.Regions {
		if len(region.Namespaces) == 0 {
			err := createNameSpace(region.Dev, "kmem")
			if err != nil {
				log.Errorf("Create kmem NameSpace error for region: %s", region.Dev)
				return errors.New("Create NameSpace error for region: " + region.Dev)
			}
		}
		chardev, err := checkKMEMNamespaceValid(region.Dev)
		if err != nil {
			return err
		}
		created, err := checkKMEMCreated(chardev)
		if err != nil {
			return err
		}
		if !created {
			err = makeNamespaceMemory(chardev)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// MaintainProjQuota maintain project quota file system
func MaintainProjQuota(regions *PmemRegions, mounter k8smount.Interface) error {
	for _, region := range regions.Regions {
		if len(region.Namespaces) == 0 {
			err := createNameSpace(region.Dev, "lvm")
			if err != nil {
				log.Errorf("Create projQuota namespace error for region: %s", region.Dev)
				return errors.New("Create projQuota namespace error for region " + region.Dev)
			}
		}
		mntOptions := []string{"prjquota", "shared"}
		fsType := "ext4"
		mkfsOptions := strings.Split("-O project,quota", " ")
		devicePath, namespaceName, err := checkProjQuotaNamespaceValid(region.Dev)
		if err != nil {
			return err
		}
		namespaceFullPath := fmt.Sprintf("/mnt/quotapath.%s", namespaceName)
		err = EnsureFolder(namespaceFullPath)
		if err != nil {
			return err
		}
		diskMounter := &k8smount.SafeFormatAndMount{Interface: mounter, Exec: utilexec.New()}
		err = formatAndMount(diskMounter, devicePath, namespaceFullPath, fsType, mkfsOptions, mntOptions)
		if err != nil {
			return err
		}
	}
	return nil
}

// MaintainLVM lvm pmem
func MaintainLVM(regions *PmemRegions) error {
	// Create Namespaces if not exist
	for _, region := range regions.Regions {
		if len(region.Namespaces) == 0 {
			if err := createNameSpace(region.Dev, "lvm"); err != nil {
				log.Errorf("Create lvm NameSpace error for region: %s", region.Dev)
				return errors.New("Create NameSpace error for region: " + region.Dev)
			}
		}
	}
	// Get Regions
	regions, err := GetRegions()
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
