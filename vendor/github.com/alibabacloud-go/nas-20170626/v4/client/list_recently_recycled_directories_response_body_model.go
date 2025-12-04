// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRecentlyRecycledDirectoriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEntries(v []*ListRecentlyRecycledDirectoriesResponseBodyEntries) *ListRecentlyRecycledDirectoriesResponseBody
	GetEntries() []*ListRecentlyRecycledDirectoriesResponseBodyEntries
	SetNextToken(v string) *ListRecentlyRecycledDirectoriesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListRecentlyRecycledDirectoriesResponseBody
	GetRequestId() *string
}

type ListRecentlyRecycledDirectoriesResponseBody struct {
	// The information about the directories that are recently deleted.
	Entries []*ListRecentlyRecycledDirectoriesResponseBodyEntries `json:"Entries,omitempty" xml:"Entries,omitempty" type:"Repeated"`
	// A pagination token.
	//
	// If not all directories are returned in a query, the return value of the NextToken parameter is not empty. In this case, you can specify a valid value for the NextToken parameter to continue the query.
	//
	// example:
	//
	// 1256****25
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9E15E394-38A6-457A-A62A-D9797C9A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListRecentlyRecycledDirectoriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRecentlyRecycledDirectoriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListRecentlyRecycledDirectoriesResponseBody) GetEntries() []*ListRecentlyRecycledDirectoriesResponseBodyEntries {
	return s.Entries
}

func (s *ListRecentlyRecycledDirectoriesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListRecentlyRecycledDirectoriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRecentlyRecycledDirectoriesResponseBody) SetEntries(v []*ListRecentlyRecycledDirectoriesResponseBodyEntries) *ListRecentlyRecycledDirectoriesResponseBody {
	s.Entries = v
	return s
}

func (s *ListRecentlyRecycledDirectoriesResponseBody) SetNextToken(v string) *ListRecentlyRecycledDirectoriesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListRecentlyRecycledDirectoriesResponseBody) SetRequestId(v string) *ListRecentlyRecycledDirectoriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRecentlyRecycledDirectoriesResponseBody) Validate() error {
	if s.Entries != nil {
		for _, item := range s.Entries {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListRecentlyRecycledDirectoriesResponseBodyEntries struct {
	// The ID of the directory.
	//
	// example:
	//
	// 04***08
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The time when the directory was last deleted.
	//
	// example:
	//
	// 2021-05-30T10:08:08Z
	LastDeleteTime *string `json:"LastDeleteTime,omitempty" xml:"LastDeleteTime,omitempty"`
	// The name of the directory.
	//
	// example:
	//
	// b
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The absolute path to the directory.
	//
	// example:
	//
	// /a/b
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s ListRecentlyRecycledDirectoriesResponseBodyEntries) String() string {
	return dara.Prettify(s)
}

func (s ListRecentlyRecycledDirectoriesResponseBodyEntries) GoString() string {
	return s.String()
}

func (s *ListRecentlyRecycledDirectoriesResponseBodyEntries) GetFileId() *string {
	return s.FileId
}

func (s *ListRecentlyRecycledDirectoriesResponseBodyEntries) GetLastDeleteTime() *string {
	return s.LastDeleteTime
}

func (s *ListRecentlyRecycledDirectoriesResponseBodyEntries) GetName() *string {
	return s.Name
}

func (s *ListRecentlyRecycledDirectoriesResponseBodyEntries) GetPath() *string {
	return s.Path
}

func (s *ListRecentlyRecycledDirectoriesResponseBodyEntries) SetFileId(v string) *ListRecentlyRecycledDirectoriesResponseBodyEntries {
	s.FileId = &v
	return s
}

func (s *ListRecentlyRecycledDirectoriesResponseBodyEntries) SetLastDeleteTime(v string) *ListRecentlyRecycledDirectoriesResponseBodyEntries {
	s.LastDeleteTime = &v
	return s
}

func (s *ListRecentlyRecycledDirectoriesResponseBodyEntries) SetName(v string) *ListRecentlyRecycledDirectoriesResponseBodyEntries {
	s.Name = &v
	return s
}

func (s *ListRecentlyRecycledDirectoriesResponseBodyEntries) SetPath(v string) *ListRecentlyRecycledDirectoriesResponseBodyEntries {
	s.Path = &v
	return s
}

func (s *ListRecentlyRecycledDirectoriesResponseBodyEntries) Validate() error {
	return dara.Validate(s)
}
