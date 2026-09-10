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
	// The information about directories on which delete operations were recently performed.
	Entries []*ListRecentlyRecycledDirectoriesResponseBodyEntries `json:"Entries,omitempty" xml:"Entries,omitempty" type:"Repeated"`
	// The pagination token for the next page.
	//
	// If the query results are not completely returned, the NextToken parameter is returned with a value. You can specify the NextToken value in the next request to continue the query.
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
	// The directory ID.
	//
	// example:
	//
	// 04***08
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The most recent time when a delete operation was performed on the directory. The time follows the ISO 8601 standard in UTC. Format: yyyy-MM-ddTHH:mm:ssZ.
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
	// The absolute path of the directory.
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
