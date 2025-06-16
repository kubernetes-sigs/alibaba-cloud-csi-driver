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
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

func getDevices() []string {
	devices := []string{}
	files, _ := ioutil.ReadDir("/dev")
	for _, file := range files {
		if !file.IsDir() && strings.HasPrefix(file.Name(), "vd") {
			devices = append(devices, fmt.Sprintf("/dev/%s", file.Name()))
		}
	}
	return devices
}

var stat = os.Stat

func calcNewDevices(old, new []string) []string {
	var devicePaths []string
	for _, d := range new {
		var isNew = true
		for _, a := range old {
			if d == a {
				isNew = false
			}
		}
		if isNew {
			info, err := stat(d)
			if err != nil {
				log.Errorf("stat new device %s error: %s", d, err)
				continue
			}
			if info.Mode()&os.ModeDevice != 0 && info.Mode()&os.ModeCharDevice == 0 {
				devicePaths = append(devicePaths, d)
			}
		}
	}

	return devicePaths
}
