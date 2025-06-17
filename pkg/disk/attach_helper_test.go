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
	"io/fs"
	"os"
	"slices"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetDevice(t *testing.T) {

	devices := getDevices()
	assert.NotNil(t, devices)

}

func TestCalNewDevices(t *testing.T) {

	old := []string{"a", "b", "c", "d"}
	new := []string{"b", "c", "d", "e", "f", "g"}

	mockStat(t, []string{"b", "c", "d", "e"}, []string{"f"})

	result := calcNewDevices(old, new)
	assert.Equal(t, []string{"e"}, result)
}

type mockFileInfo struct {
	mode os.FileMode
}

func (m *mockFileInfo) Mode() os.FileMode  { return m.mode }
func (m *mockFileInfo) IsDir() bool        { return false }
func (m *mockFileInfo) Name() string       { return "" }
func (m *mockFileInfo) Size() int64        { return 0 }
func (m *mockFileInfo) Sys() any           { return nil }
func (m *mockFileInfo) ModTime() time.Time { return time.Time{} }

func mockStat(t *testing.T, block, char []string) {
	stat = func(path string) (os.FileInfo, error) {
		if slices.Contains(block, path) {
			return &mockFileInfo{mode: os.ModeDevice}, nil
		}
		if slices.Contains(char, path) {
			return &mockFileInfo{mode: os.ModeDevice & os.ModeCharDevice}, nil
		}
		return nil, fs.ErrNotExist
	}
	t.Cleanup(func() { stat = os.Stat })
}
