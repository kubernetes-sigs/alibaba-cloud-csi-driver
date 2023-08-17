/*
Copyright 2019 The Kubernetes Authors.

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

package mem

import (
	"encoding/json"
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	log "github.com/sirupsen/logrus"
)

const (
	// MetadataURL is metadata server url
	MetadataURL = "http://100.100.100.200/latest/meta-data/"
	// InstanceID is the instance id tag
	InstanceID = "instance-id"
)

// GetPvNameFormMntPoint get pv name
func GetPvNameFormMntPoint(mntPath string) string {
	if mntPath == "" {
		return ""
	}
	if strings.HasSuffix(mntPath, "/mount") {
		tmpPath := mntPath[0 : len(mntPath)-6]
		pvName := filepath.Base(tmpPath)
		return pvName
	}
	return ""
}

func pickRegionForKMEM(size string) (string, error) {
	return "bind:3", nil
}

func createNameSpace(region string) error {
	createCmd := fmt.Sprintf("%s ndctl create-namespace -r %s --mode=devdax", NsenterCmd, region)
	_, err := utils.Run(createCmd)
	if err != nil {
		log.Errorf("Create NameSpace for region %s error: %v", region, err)
		return err
	}
	log.Infof("Create NameSpace for region %s successful", region)
	return nil
}

func checkKMEMNamespaceValid(region string) (string, error) {
	regions, err := getRegionNamespaceInfo(region)
	if err != nil {
		return "", err
	}
	namespaceMode := regions.Regions[0].Namespaces[0].Mode
	if namespaceMode != "devdax" {
		log.Errorf("KMEM namespace mode %s wrong", namespaceMode)
		return "", errors.New("KMEM namespace wrong mode" + namespaceMode)
	}
	return regions.Regions[0].Namespaces[0].CharDev, nil
}

func getRegionNamespaceInfo(region string) (*PmemRegions, error) {
	listCmd := fmt.Sprintf("%s ndctl list -RN -r %s", NsenterCmd, region)

	out, err := utils.Run(listCmd)
	if err != nil {
		log.Errorf("List NameSpace for region %s error: %v", region, err)
		return nil, err
	}
	regions := &PmemRegions{}
	err = json.Unmarshal(([]byte)(out), regions)
	if len(regions.Regions) == 0 {
		log.Errorf("list Namespace for region %s get 0 region", region)
		return nil, errors.New("list Namespace get 0 region by " + region)
	}

	if len(regions.Regions[0].Namespaces) != 1 {
		log.Errorf("list Namespace for region %s get 0 or multi namespaces", region)
		return nil, errors.New("list Namespace for region get 0 or multi namespaces" + region)
	}
	return regions, nil
}

// GetRegions ...
func GetRegions() (*PmemRegions, error) {
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

func checkKMEMCreated(chardev string) (bool, error) {
	listCmd := fmt.Sprintf("%s daxctl list", NsenterCmd)
	out, err := utils.Run(listCmd)
	if err != nil {
		log.Errorf("List daxctl error: %v", err)
		return false, err
	}
	memList := []*DaxctrlMem{}
	err = json.Unmarshal(([]byte)(out), &memList)
	if err != nil {
		return false, err
	}
	for _, mem := range memList {
		if mem.Chardev == chardev && mem.Mode == "system-ram" {
			return true, nil
		}
	}
	return false, nil
}

func makeNamespaceMemory(chardev string) error {
	makeCmd := fmt.Sprintf("%s daxctl reconfigure-device -m system-ram %s", NsenterCmd, chardev)
	_, err := utils.Run(makeCmd)
	return err
}
