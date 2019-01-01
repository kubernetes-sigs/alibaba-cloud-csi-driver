package disk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

var DefaultAttachEntry AttachEntry

func init() {
	DefaultAttachEntry = &attachEntry{}
}

var logPath = "/var/lib/kubelet/plugins/csi-diskplugin/mount-entries"

type AttachEntry interface {
	Get(diskID string) (string, error)
	Add(diskID, devicePath string) error
	Remove(diskID string) error
}

type attachEntry struct {
	sync.RWMutex
}

func (m *attachEntry) Get(diskID string) (string, error) {
	m.RLock()
	defer m.RUnlock()

	dat, err := ioutil.ReadFile(logPath)

	if err != nil {
		return "", err
	}

	var entries map[string]string
	err = json.Unmarshal(dat, &entries)
	if err != nil {
		return "", err
	}

	device, ok := entries[diskID]
	if !ok {
		return "", fmt.Errorf("NOT EXISTS")
	}

	return device, nil
}
func (m *attachEntry) Remove(diskID string) error {
	m.Lock()
	m.Unlock()

	dat, err := ioutil.ReadFile(logPath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		} else {
			return err
		}
	}

	var entries map[string]string
	err = json.Unmarshal(dat, &entries)
	if err != nil {
		return err
	}
	delete(entries, diskID)
	dat, err = json.Marshal(entries)
	return ioutil.WriteFile(logPath, dat, 0644)
}

func (m *attachEntry) Add(diskID, devicePath string) error {
	m.Lock()
	m.Unlock()

	var entries map[string]string

	dat, err := ioutil.ReadFile(logPath)
	if err != nil {
		if os.IsNotExist(err) {
			entries = make(map[string]string)
		} else {
			return err
		}
	} else {
		err = json.Unmarshal(dat, &entries)
		if err != nil {
			return err
		}
	}

	entries[diskID] = devicePath
	dat, err = json.Marshal(entries)

	return ioutil.WriteFile(logPath, dat, 0644)
}
