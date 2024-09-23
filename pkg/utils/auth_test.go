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
	"gopkg.in/h2non/gock.v1"
	"k8s.io/klog/v2"
)

func TestGetAccessControl(t *testing.T) {
	testAccessKey := "testkey"
	testAccessKeySecret := "testvalue"
	os.Setenv("ACCESS_KEY_ID", testAccessKey)
	os.Setenv("ACCESS_KEY_SECRET", testAccessKeySecret)
	ac := GetAccessControl()
	assert.Equal(t, testAccessKey, ac.AccessKeyID)
	assert.Equal(t, testAccessKeySecret, ac.AccessKeySecret)
	assert.Empty(t, ac.StsToken)
	os.Unsetenv("ACCESS_KEY_ID")
	os.Unsetenv("ACCESS_KEY_SECRET")
	ac = GetAccessControl()
	assert.Empty(t, ac.AccessKeyID)
	assert.Empty(t, ac.AccessKeySecret)
	assert.Empty(t, ac.StsToken)
}

func TestGetOIDCToken(t *testing.T) {
	defer gock.Off()
	testExamples := []struct {
		regionId    string
		ownerId     string
		expectKeyId string
		fatalError  bool
		newProvider bool
	}{
		{
			regionId:    "cn-test1",
			ownerId:     "owner-test1",
			expectKeyId: "",
			fatalError:  false,
			newProvider: true,
		},
		{
			regionId:    "",
			ownerId:     "owner-test1",
			expectKeyId: "",
			fatalError:  false,
			newProvider: false,
		},
		{
			regionId:    "",
			ownerId:     "owner-test1",
			expectKeyId: "",
			fatalError:  true,
			newProvider: true,
		},
		{
			regionId:    "",
			ownerId:     "",
			expectKeyId: "",
			fatalError:  true,
			newProvider: true,
		},
	}
	defer func() { klog.OsExit = os.Exit }()
	klog.OsExit = func(c int) { panic(c) }
	for _, test := range testExamples {
		if test.newProvider {
			oidcProvider = nil
		}
		os.Setenv("USE_OIDC_AUTH_INNER", "true")
		// os.Setenv("REGION_ID", test.regionId)
		// os.Setenv("ACCOUNT_ID", test.ownerId)
		gock.New("http://100.100.100.200").Get("/latest/meta-data/region-id").Reply(200).BodyString(test.regionId)
		gock.New("http://100.100.100.200").Get("/latest/meta-data/owner-account-id").Reply(200).BodyString(test.ownerId)
		if test.fatalError {
			assert.Panics(t, func() { getOIDCToken() })
		} else {
			ac := getOIDCToken()
			if ac.AccessKeyID != "" {
				assert.Equal(t, test.expectKeyId, ac.AccessKeyID)
			}
		}
	}
}
