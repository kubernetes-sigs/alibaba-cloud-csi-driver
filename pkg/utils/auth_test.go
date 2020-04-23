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

package utils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDefaultAK(t *testing.T) {
	testAccessKey := "testkey"
	testAccessKeySecret := "testvalue"
	os.Setenv("ACCESS_KEY_ID", testAccessKey)
	os.Setenv("ACCESS_KEY_SECRET", testAccessKeySecret)
	accessKeyID, accessKeySecret, accessToken := GetDefaultAK()
	assert.Equal(t, testAccessKey, accessKeyID)
	assert.Equal(t, testAccessKeySecret, accessKeySecret)
	assert.Empty(t, accessToken)
	os.Unsetenv("ACCESS_KEY_ID")
	os.Unsetenv("ACCESS_KEY_SECRET")
	accessKeyID, accessKeySecret, accessToken = GetDefaultAK()
	assert.Empty(t, accessKeyID)
	assert.Empty(t, accessKeySecret)
	assert.Empty(t, accessToken)
}
