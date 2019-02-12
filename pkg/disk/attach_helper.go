package disk

import (
	"fmt"
	"io/ioutil"
	"strings"
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
			devicePaths = append(devicePaths, d)
		}
	}

	return devicePaths
}
