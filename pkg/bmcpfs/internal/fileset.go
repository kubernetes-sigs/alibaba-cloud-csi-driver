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
	"context"
	"errors"
	"fmt"

	nasclient "github.com/alibabacloud-go/nas-20170626/v4/client"
	"github.com/alibabacloud-go/tea/tea"
	"k8s.io/klog/v2"
)

// CPFSFileSetManager is the interface for managing CPFS filesets
type CPFSFileSetManager interface {
	CreateFileSet(ctx context.Context, fsID, pvName, fileSystemPath string, countLimit, sizeBytes int64, deleteProtection bool) (string, error)
	DeleteFileSet(ctx context.Context, fsID, fileSetID string) error
	DescribeFileSets(ctx context.Context, fsID string) ([]*nasclient.DescribeFilesetsResponseBodyEntriesEntrie, error)
	DescribeFilesetByFsetID(ctx context.Context, fsID, fileSetID string) (*nasclient.DescribeFilesetsResponseBodyEntriesEntrie, error)
	DescribeFilesetByFilePath(ctx context.Context, fsID, FilesystemPath string) (*nasclient.DescribeFilesetsResponseBodyEntriesEntrie, error)
}

type NasClient interface {
	CreateFileset(request *nasclient.CreateFilesetRequest) (*nasclient.CreateFilesetResponse, error)
	DeleteFileset(request *nasclient.DeleteFilesetRequest) (*nasclient.DeleteFilesetResponse, error)
	DescribeFilesets(request *nasclient.DescribeFilesetsRequest) (*nasclient.DescribeFilesetsResponse, error)
}

type cpfsFileSetManager struct {
	client NasClient
}

// ErrFilesetNotFound signals that the requested fileset does not exist.
var ErrFilesetNotFound = errors.New("fileset not found")

// NewCPFSFileSetManager creates a new CPFS fileset manager
func NewCPFSFileSetManager(client NasClient) CPFSFileSetManager {
	return &cpfsFileSetManager{
		client: client,
	}
}

func (m *cpfsFileSetManager) CreateFileSet(ctx context.Context, fsID, pvName, fileSystemPath string, countLimit, sizeBytes int64, deleteProtection bool) (string, error) {
	quota := &nasclient.CreateFilesetRequestQuota{
		SizeLimit: tea.Int64(sizeBytes),
	}
	if countLimit != 0 {
		quota.FileCountLimit = tea.Int64(countLimit)
	}
	request := &nasclient.CreateFilesetRequest{
		FileSystemId:       tea.String(fsID),
		ClientToken:        tea.String(pvName),
		DeletionProtection: tea.Bool(deleteProtection),
		FileSystemPath:     tea.String(fileSystemPath),
		Quota:              quota,
	}
	klog.InfoS("before create fileset", "request", *request)
	response, err := m.client.CreateFileset(request)
	if err != nil {
		return "", fmt.Errorf("create fileset %s/%s failed: %v", fsID, pvName, err)
	}
	if response == nil || response.Body == nil || response.Body.FsetId == nil {
		return "", fmt.Errorf("create fileset %s/%s failed: empty response", fsID, fileSystemPath)
	}
	klog.V(4).Infof("create fileset %s/%s successfully", fsID, fileSystemPath)
	return *response.Body.FsetId, nil
}

// DeleteFileSet deletes a fileset by fileSystemId and fileSetId
func (m *cpfsFileSetManager) DeleteFileSet(ctx context.Context, fsID, fileSetID string) error {
	request := &nasclient.DeleteFilesetRequest{
		FileSystemId: tea.String(fsID),
		FsetId:       tea.String(fileSetID),
	}

	response, err := m.client.DeleteFileset(request)
	if err != nil {
		return fmt.Errorf("delete fileset %s/%s failed: %v", fsID, fileSetID, err)
	}

	if response == nil || response.Body == nil {
		return fmt.Errorf("delete fileset %s/%s failed: empty response", fsID, fileSetID)
	}

	klog.V(4).Infof("delete fileset %s/%s successfully", fsID, fileSetID)
	return nil
}

// DescribeFileSets describes all filesets in a filesystem by fileSystemId
func (m *cpfsFileSetManager) DescribeFileSets(ctx context.Context, fsID string) ([]*nasclient.DescribeFilesetsResponseBodyEntriesEntrie, error) {
	var allFileSets []*nasclient.DescribeFilesetsResponseBodyEntriesEntrie
	var nextToken *string

	// Handle pagination - iterate through all pages to collect all filesets
	for {
		request := &nasclient.DescribeFilesetsRequest{
			FileSystemId: tea.String(fsID),
			NextToken:    nextToken,
		}

		response, err := m.client.DescribeFilesets(request)
		if err != nil {
			return nil, fmt.Errorf("describe filesets for filesystem %s failed: %v", fsID, err)
		}

		if response == nil || response.Body == nil {
			return nil, fmt.Errorf("describe filesets for filesystem %s failed: empty response", fsID)
		}

		// Collect filesets from current page
		if response.Body.Entries != nil && len(response.Body.Entries.Entrie) > 0 {
			allFileSets = append(allFileSets, response.Body.Entries.Entrie...)
		}

		// Check if there are more pages
		if response.Body.NextToken == nil || *response.Body.NextToken == "" {
			// No more pages, break the loop
			break
		}

		// Move to the next page
		nextToken = response.Body.NextToken
	}

	klog.V(4).Infof("describe filesets for filesystem %s successfully, total count: %d", fsID, len(allFileSets))
	return allFileSets, nil
}

// DescribeFileset describes a specific fileset by fileSystemId and fileSetId
func (m *cpfsFileSetManager) DescribeFilesetByFsetID(ctx context.Context, fsID, fileSetID string) (*nasclient.DescribeFilesetsResponseBodyEntriesEntrie, error) {
	request := &nasclient.DescribeFilesetsRequest{
		FileSystemId: tea.String(fsID),
		Filters: []*nasclient.DescribeFilesetsRequestFilters{
			{
				Key:   tea.String("FsetIds"),
				Value: tea.String(fileSetID),
			},
		},
	}

	response, err := m.client.DescribeFilesets(request)
	if err != nil {
		return nil, fmt.Errorf("describe fileset %s/%s failed: %v", fsID, fileSetID, err)
	}

	if response == nil || response.Body == nil {
		return nil, fmt.Errorf("describe fileset %s/%s failed: empty response", fsID, fileSetID)
	}

	// Check if fileset exists in the response
	if response.Body.Entries == nil || len(response.Body.Entries.Entrie) == 0 {
		return nil, fmt.Errorf("fileset %s not found in filesystem %s: %w", fileSetID, fsID, ErrFilesetNotFound)
	}

	// Return the first (and should be only) fileset
	fileset := response.Body.Entries.Entrie[0]
	klog.V(4).Infof("describe fileset %s/%s successfully", fsID, fileSetID)
	return fileset, nil
}

// DescribeFilesetByFilePath describes a specific fileset by fileSystemId and filesystemPath
func (m *cpfsFileSetManager) DescribeFilesetByFilePath(ctx context.Context, fsID, filesystemPath string) (*nasclient.DescribeFilesetsResponseBodyEntriesEntrie, error) {
	request := &nasclient.DescribeFilesetsRequest{
		FileSystemId: tea.String(fsID),
		Filters: []*nasclient.DescribeFilesetsRequestFilters{
			{
				Key:   tea.String("FileSystemPath"),
				Value: tea.String(filesystemPath),
			},
		},
	}

	response, err := m.client.DescribeFilesets(request)
	if err != nil {
		return nil, fmt.Errorf("describe fileset by filepath %s/%s failed: %v", fsID, filesystemPath, err)
	}

	if response == nil || response.Body == nil {
		return nil, fmt.Errorf("describe fileset by filepath %s/%s failed: empty response", fsID, filesystemPath)
	}

	// Check if fileset exists in the response
	if response.Body.Entries == nil || len(response.Body.Entries.Entrie) == 0 {
		return nil, fmt.Errorf("fileset with path %s not found in filesystem %s: %w", filesystemPath, fsID, ErrFilesetNotFound)
	}

	// Return the first (and should be only) fileset
	fileset := response.Body.Entries.Entrie[0]
	klog.V(4).Infof("describe fileset by filepath %s/%s successfully", fsID, filesystemPath)
	return fileset, nil
}
