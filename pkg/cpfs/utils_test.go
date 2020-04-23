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

package cpfs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCpfsDetails(t *testing.T) {

	cpfsServersString := "0.0.0.0:/test,0.0.0.1,0.0.0.2"
	cpfsServer, cpfsFileSystem, cpfsPath := GetCpfsDetails(cpfsServersString)
	assert.Equal(t, "0.0.0.0", cpfsServer)
	assert.Equal(t, "test", cpfsFileSystem)
	assert.Equal(t, "/", cpfsPath)

	cpfsServersStringTest := "0.0.0.1"
	cpfsServer, cpfsFileSystem, cpfsPath = GetCpfsDetails(cpfsServersStringTest)
	assert.Equal(t, "", cpfsServer)
	assert.Equal(t, "", cpfsFileSystem)
	assert.Equal(t, "", cpfsPath)

}
