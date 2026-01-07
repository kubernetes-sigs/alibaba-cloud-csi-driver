// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePrefixListAttributesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAddressFamily(v string) *DescribePrefixListAttributesResponseBody
	GetAddressFamily() *string
	SetCreationTime(v string) *DescribePrefixListAttributesResponseBody
	GetCreationTime() *string
	SetDescription(v string) *DescribePrefixListAttributesResponseBody
	GetDescription() *string
	SetEntries(v *DescribePrefixListAttributesResponseBodyEntries) *DescribePrefixListAttributesResponseBody
	GetEntries() *DescribePrefixListAttributesResponseBodyEntries
	SetMaxEntries(v int32) *DescribePrefixListAttributesResponseBody
	GetMaxEntries() *int32
	SetPrefixListId(v string) *DescribePrefixListAttributesResponseBody
	GetPrefixListId() *string
	SetPrefixListName(v string) *DescribePrefixListAttributesResponseBody
	GetPrefixListName() *string
	SetRequestId(v string) *DescribePrefixListAttributesResponseBody
	GetRequestId() *string
}

type DescribePrefixListAttributesResponseBody struct {
	// The IP address family of the prefix list. Valid values:
	//
	// 	- IPv4
	//
	// 	- IPv6
	//
	// example:
	//
	// IPv4
	AddressFamily *string `json:"AddressFamily,omitempty" xml:"AddressFamily,omitempty"`
	// The time when the prefix list was created.
	//
	// example:
	//
	// 2021-02-20T07:11Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The description of the prefix list.
	//
	// example:
	//
	// This is description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Details about the entries in the prefix list.
	Entries *DescribePrefixListAttributesResponseBodyEntries `json:"Entries,omitempty" xml:"Entries,omitempty" type:"Struct"`
	// The maximum number of entries in the prefix list.
	//
	// example:
	//
	// 10
	MaxEntries *int32 `json:"MaxEntries,omitempty" xml:"MaxEntries,omitempty"`
	// The ID of the prefix list.
	//
	// example:
	//
	// pl-x1j1k5ykzqlixdcy****
	PrefixListId *string `json:"PrefixListId,omitempty" xml:"PrefixListId,omitempty"`
	// The name of the prefix list.
	//
	// example:
	//
	// PrefixListNameSample
	PrefixListName *string `json:"PrefixListName,omitempty" xml:"PrefixListName,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 38793DB8-A4B2-4AEC-BFD3-111234E9188D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePrefixListAttributesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePrefixListAttributesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePrefixListAttributesResponseBody) GetAddressFamily() *string {
	return s.AddressFamily
}

func (s *DescribePrefixListAttributesResponseBody) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribePrefixListAttributesResponseBody) GetDescription() *string {
	return s.Description
}

func (s *DescribePrefixListAttributesResponseBody) GetEntries() *DescribePrefixListAttributesResponseBodyEntries {
	return s.Entries
}

func (s *DescribePrefixListAttributesResponseBody) GetMaxEntries() *int32 {
	return s.MaxEntries
}

func (s *DescribePrefixListAttributesResponseBody) GetPrefixListId() *string {
	return s.PrefixListId
}

func (s *DescribePrefixListAttributesResponseBody) GetPrefixListName() *string {
	return s.PrefixListName
}

func (s *DescribePrefixListAttributesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePrefixListAttributesResponseBody) SetAddressFamily(v string) *DescribePrefixListAttributesResponseBody {
	s.AddressFamily = &v
	return s
}

func (s *DescribePrefixListAttributesResponseBody) SetCreationTime(v string) *DescribePrefixListAttributesResponseBody {
	s.CreationTime = &v
	return s
}

func (s *DescribePrefixListAttributesResponseBody) SetDescription(v string) *DescribePrefixListAttributesResponseBody {
	s.Description = &v
	return s
}

func (s *DescribePrefixListAttributesResponseBody) SetEntries(v *DescribePrefixListAttributesResponseBodyEntries) *DescribePrefixListAttributesResponseBody {
	s.Entries = v
	return s
}

func (s *DescribePrefixListAttributesResponseBody) SetMaxEntries(v int32) *DescribePrefixListAttributesResponseBody {
	s.MaxEntries = &v
	return s
}

func (s *DescribePrefixListAttributesResponseBody) SetPrefixListId(v string) *DescribePrefixListAttributesResponseBody {
	s.PrefixListId = &v
	return s
}

func (s *DescribePrefixListAttributesResponseBody) SetPrefixListName(v string) *DescribePrefixListAttributesResponseBody {
	s.PrefixListName = &v
	return s
}

func (s *DescribePrefixListAttributesResponseBody) SetRequestId(v string) *DescribePrefixListAttributesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePrefixListAttributesResponseBody) Validate() error {
	if s.Entries != nil {
		if err := s.Entries.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribePrefixListAttributesResponseBodyEntries struct {
	Entry []*DescribePrefixListAttributesResponseBodyEntriesEntry `json:"Entry,omitempty" xml:"Entry,omitempty" type:"Repeated"`
}

func (s DescribePrefixListAttributesResponseBodyEntries) String() string {
	return dara.Prettify(s)
}

func (s DescribePrefixListAttributesResponseBodyEntries) GoString() string {
	return s.String()
}

func (s *DescribePrefixListAttributesResponseBodyEntries) GetEntry() []*DescribePrefixListAttributesResponseBodyEntriesEntry {
	return s.Entry
}

func (s *DescribePrefixListAttributesResponseBodyEntries) SetEntry(v []*DescribePrefixListAttributesResponseBodyEntriesEntry) *DescribePrefixListAttributesResponseBodyEntries {
	s.Entry = v
	return s
}

func (s *DescribePrefixListAttributesResponseBodyEntries) Validate() error {
	if s.Entry != nil {
		for _, item := range s.Entry {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePrefixListAttributesResponseBodyEntriesEntry struct {
	// The CIDR block in entry N.
	//
	// example:
	//
	// 192.168.1.0/24
	Cidr *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	// The description in entry N.
	//
	// example:
	//
	// Description Sample 01
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s DescribePrefixListAttributesResponseBodyEntriesEntry) String() string {
	return dara.Prettify(s)
}

func (s DescribePrefixListAttributesResponseBodyEntriesEntry) GoString() string {
	return s.String()
}

func (s *DescribePrefixListAttributesResponseBodyEntriesEntry) GetCidr() *string {
	return s.Cidr
}

func (s *DescribePrefixListAttributesResponseBodyEntriesEntry) GetDescription() *string {
	return s.Description
}

func (s *DescribePrefixListAttributesResponseBodyEntriesEntry) SetCidr(v string) *DescribePrefixListAttributesResponseBodyEntriesEntry {
	s.Cidr = &v
	return s
}

func (s *DescribePrefixListAttributesResponseBodyEntriesEntry) SetDescription(v string) *DescribePrefixListAttributesResponseBodyEntriesEntry {
	s.Description = &v
	return s
}

func (s *DescribePrefixListAttributesResponseBodyEntriesEntry) Validate() error {
	return dara.Validate(s)
}
