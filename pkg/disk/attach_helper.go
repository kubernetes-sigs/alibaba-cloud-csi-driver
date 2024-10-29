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

package disk

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"slices"

	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
)

func getNoSerialDevices() (sets.Set[string], error) {
	return getNoSerialDevicesFromSysfs(DefaultDeviceManager.SysfsPath + "/block")
}

func deviceHasSerialNumber(devicePath string) (bool, error) {
	return deviceNameHasSerialNumber(DefaultDeviceManager.SysfsPath+"/block", filepath.Base(devicePath))
}

func getNoSerialDevicesFromSysfs(sysfsBlockPath string) (sets.Set[string], error) {
	devices := sets.New[string]()
	entries, err := os.ReadDir(sysfsBlockPath)
	if err != nil {
		return nil, fmt.Errorf("failed to readdir %s: %v", sysfsBlockPath, err)
	}
	for _, entry := range entries {
		hasSerial, err := deviceNameHasSerialNumber(sysfsBlockPath, entry.Name())
		if err != nil {
			return nil, err
		}
		if !hasSerial {
			devices.Insert(fmt.Sprintf("/dev/%s", entry.Name()))
		}
	}
	klog.V(4).InfoS("got no serial devices", "devices", devices)
	return devices, nil
}

func deviceNameHasSerialNumber(sysfsBlockPath, name string) (bool, error) {
	serial, err := os.ReadFile(fmt.Sprintf("%s/%s/serial", sysfsBlockPath, name))
	if errors.Is(err, fs.ErrNotExist) {
		serial, err = os.ReadFile(fmt.Sprintf("%s/%s/device/serial", sysfsBlockPath, name))
	}
	if errors.Is(err, fs.ErrNotExist) {
		return false, nil
	}
	if err != nil {
		return false, fmt.Errorf("failed to read serial for %s: %v", name, err)
	}
	return slices.ContainsFunc(serial, func(b byte) bool { return b != '\n' }), nil
}
