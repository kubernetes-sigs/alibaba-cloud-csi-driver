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

package nas

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDoMount(t *testing.T) {

	nfsServer := "0.0.0.0"
	nfsPath := "/test"
	nfsVers := "2.0"
	mountOptions := ""
	mountPoint := ".tmp"
	volumeID := "testtsettest"
	err := DoMount(MountProtocolNFS, nfsServer, nfsPath, nfsVers, mountOptions, mountPoint, volumeID, "podUID", "false")
	assert.NotNil(t, err)

}

func TestCheckNfsPathMounted(t *testing.T) {

	mountPoint := "/data"
	path := "tmp"

	result := CheckNfsPathMounted(mountPoint, path)
	assert.False(t, result)

}

func TestGetNfsDetails(t *testing.T) {

	nfsServerString := "0.0.0.0"
	nfsServer, nfsPath := GetNfsDetails(nfsServerString)
	assert.Equal(t, "0.0.0.0", nfsServer)
	assert.Equal(t, "/", nfsPath)

	nfsServerString = "0.0.0.0:/test"
	nfsServer, nfsPath = GetNfsDetails(nfsServerString)
	assert.Equal(t, "0.0.0.0", nfsServer)
	assert.Equal(t, "/test", nfsPath)

	nfsServerString = "0.0.0.0,0.0.0.1:/test/"
	nfsServer, nfsPath = GetNfsDetails(nfsServerString)
	assert.Equal(t, "0.0.0.0", nfsServer)
	assert.Equal(t, "/", nfsPath)

}

func TestWaitTimeout(t *testing.T) {

	var wg sync.WaitGroup
	wg.Add(1)

	go func(*sync.WaitGroup) {
		time.Sleep(2 * time.Second)
		wg.Done()
	}(&wg)

	result := waitTimeout(&wg, 1)
	assert.True(t, result)

	var wg1 sync.WaitGroup
	wg1.Add(1)

	go func(*sync.WaitGroup) {
		time.Sleep(1 * time.Second)
		wg1.Done()
	}(&wg1)

	result = waitTimeout(&wg1, 2)
	assert.False(t, result)
}

func TestParseMountFlags(t *testing.T) {

	mntOptions1 := []string{"mnt=/test", "vers=3.0"}

	ver, result := ParseMountFlags(mntOptions1)

	assert.Equal(t, "3", ver)
	assert.Equal(t, "mnt=/test", result)

	mntOptions2 := []string{"mnt=/test", "vers=3"}

	ver, result = ParseMountFlags(mntOptions2)

	assert.Equal(t, "3", ver)
	assert.Equal(t, "mnt=/test", result)

	mntOptions3 := []string{"mnt=/test", "vers=4.0"}

	ver, result = ParseMountFlags(mntOptions3)

	assert.Equal(t, "4.0", ver)
	assert.Equal(t, "mnt=/test", result)

	mntOptions4 := []string{"mnt=/test", "vers=4.1"}

	ver, result = ParseMountFlags(mntOptions4)

	assert.Equal(t, "4.1", ver)
	assert.Equal(t, "mnt=/test", result)
}
