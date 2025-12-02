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

package internal

import (
	nasclient "github.com/alibabacloud-go/nas-20170626/v4/client"
)

// MockNasClient is a mock implementation of NasClient interface
// This is a simplified mock for testing purposes
type MockNasClient struct {
	CreateFilesetFunc    func(request *nasclient.CreateFilesetRequest) (*nasclient.CreateFilesetResponse, error)
	DeleteFilesetFunc    func(request *nasclient.DeleteFilesetRequest) (*nasclient.DeleteFilesetResponse, error)
	DescribeFilesetsFunc func(request *nasclient.DescribeFilesetsRequest) (*nasclient.DescribeFilesetsResponse, error)
}

// CreateFileset mocks the CreateFileset method
func (m *MockNasClient) CreateFileset(request *nasclient.CreateFilesetRequest) (*nasclient.CreateFilesetResponse, error) {
	if m.CreateFilesetFunc != nil {
		return m.CreateFilesetFunc(request)
	}
	return nil, nil
}

// DeleteFileset mocks the DeleteFileset method
func (m *MockNasClient) DeleteFileset(request *nasclient.DeleteFilesetRequest) (*nasclient.DeleteFilesetResponse, error) {
	if m.DeleteFilesetFunc != nil {
		return m.DeleteFilesetFunc(request)
	}
	return nil, nil
}

// DescribeFilesets mocks the DescribeFilesets method
func (m *MockNasClient) DescribeFilesets(request *nasclient.DescribeFilesetsRequest) (*nasclient.DescribeFilesetsResponse, error) {
	if m.DescribeFilesetsFunc != nil {
		return m.DescribeFilesetsFunc(request)
	}
	return nil, nil
}

// NewMockNasClient creates a new mock NAS client
func NewMockNasClient() *MockNasClient {
	return &MockNasClient{}
}
