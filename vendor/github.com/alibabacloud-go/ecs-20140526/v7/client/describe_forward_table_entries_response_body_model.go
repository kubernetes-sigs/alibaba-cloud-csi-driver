// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeForwardTableEntriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetForwardTableEntries(v *DescribeForwardTableEntriesResponseBodyForwardTableEntries) *DescribeForwardTableEntriesResponseBody
	GetForwardTableEntries() *DescribeForwardTableEntriesResponseBodyForwardTableEntries
	SetPageNumber(v int32) *DescribeForwardTableEntriesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeForwardTableEntriesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeForwardTableEntriesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeForwardTableEntriesResponseBody
	GetTotalCount() *int32
}

type DescribeForwardTableEntriesResponseBody struct {
	ForwardTableEntries *DescribeForwardTableEntriesResponseBodyForwardTableEntries `json:"ForwardTableEntries,omitempty" xml:"ForwardTableEntries,omitempty" type:"Struct"`
	PageNumber          *int32                                                      `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize            *int32                                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId           *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount          *int32                                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeForwardTableEntriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeForwardTableEntriesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeForwardTableEntriesResponseBody) GetForwardTableEntries() *DescribeForwardTableEntriesResponseBodyForwardTableEntries {
	return s.ForwardTableEntries
}

func (s *DescribeForwardTableEntriesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeForwardTableEntriesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeForwardTableEntriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeForwardTableEntriesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeForwardTableEntriesResponseBody) SetForwardTableEntries(v *DescribeForwardTableEntriesResponseBodyForwardTableEntries) *DescribeForwardTableEntriesResponseBody {
	s.ForwardTableEntries = v
	return s
}

func (s *DescribeForwardTableEntriesResponseBody) SetPageNumber(v int32) *DescribeForwardTableEntriesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeForwardTableEntriesResponseBody) SetPageSize(v int32) *DescribeForwardTableEntriesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeForwardTableEntriesResponseBody) SetRequestId(v string) *DescribeForwardTableEntriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeForwardTableEntriesResponseBody) SetTotalCount(v int32) *DescribeForwardTableEntriesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeForwardTableEntriesResponseBody) Validate() error {
	if s.ForwardTableEntries != nil {
		if err := s.ForwardTableEntries.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeForwardTableEntriesResponseBodyForwardTableEntries struct {
	ForwardTableEntry []*DescribeForwardTableEntriesResponseBodyForwardTableEntriesForwardTableEntry `json:"ForwardTableEntry,omitempty" xml:"ForwardTableEntry,omitempty" type:"Repeated"`
}

func (s DescribeForwardTableEntriesResponseBodyForwardTableEntries) String() string {
	return dara.Prettify(s)
}

func (s DescribeForwardTableEntriesResponseBodyForwardTableEntries) GoString() string {
	return s.String()
}

func (s *DescribeForwardTableEntriesResponseBodyForwardTableEntries) GetForwardTableEntry() []*DescribeForwardTableEntriesResponseBodyForwardTableEntriesForwardTableEntry {
	return s.ForwardTableEntry
}

func (s *DescribeForwardTableEntriesResponseBodyForwardTableEntries) SetForwardTableEntry(v []*DescribeForwardTableEntriesResponseBodyForwardTableEntriesForwardTableEntry) *DescribeForwardTableEntriesResponseBodyForwardTableEntries {
	s.ForwardTableEntry = v
	return s
}

func (s *DescribeForwardTableEntriesResponseBodyForwardTableEntries) Validate() error {
	if s.ForwardTableEntry != nil {
		for _, item := range s.ForwardTableEntry {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeForwardTableEntriesResponseBodyForwardTableEntriesForwardTableEntry struct {
	ExternalIp     *string `json:"ExternalIp,omitempty" xml:"ExternalIp,omitempty"`
	ExternalPort   *string `json:"ExternalPort,omitempty" xml:"ExternalPort,omitempty"`
	ForwardEntryId *string `json:"ForwardEntryId,omitempty" xml:"ForwardEntryId,omitempty"`
	ForwardTableId *string `json:"ForwardTableId,omitempty" xml:"ForwardTableId,omitempty"`
	InternalIp     *string `json:"InternalIp,omitempty" xml:"InternalIp,omitempty"`
	InternalPort   *string `json:"InternalPort,omitempty" xml:"InternalPort,omitempty"`
	IpProtocol     *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeForwardTableEntriesResponseBodyForwardTableEntriesForwardTableEntry) String() string {
	return dara.Prettify(s)
}

func (s DescribeForwardTableEntriesResponseBodyForwardTableEntriesForwardTableEntry) GoString() string {
	return s.String()
}

func (s *DescribeForwardTableEntriesResponseBodyForwardTableEntriesForwardTableEntry) GetExternalIp() *string {
	return s.ExternalIp
}

func (s *DescribeForwardTableEntriesResponseBodyForwardTableEntriesForwardTableEntry) GetExternalPort() *string {
	return s.ExternalPort
}

func (s *DescribeForwardTableEntriesResponseBodyForwardTableEntriesForwardTableEntry) GetForwardEntryId() *string {
	return s.ForwardEntryId
}

func (s *DescribeForwardTableEntriesResponseBodyForwardTableEntriesForwardTableEntry) GetForwardTableId() *string {
	return s.ForwardTableId
}

func (s *DescribeForwardTableEntriesResponseBodyForwardTableEntriesForwardTableEntry) GetInternalIp() *string {
	return s.InternalIp
}

func (s *DescribeForwardTableEntriesResponseBodyForwardTableEntriesForwardTableEntry) GetInternalPort() *string {
	return s.InternalPort
}

func (s *DescribeForwardTableEntriesResponseBodyForwardTableEntriesForwardTableEntry) GetIpProtocol() *string {
	return s.IpProtocol
}

func (s *DescribeForwardTableEntriesResponseBodyForwardTableEntriesForwardTableEntry) GetStatus() *string {
	return s.Status
}

func (s *DescribeForwardTableEntriesResponseBodyForwardTableEntriesForwardTableEntry) SetExternalIp(v string) *DescribeForwardTableEntriesResponseBodyForwardTableEntriesForwardTableEntry {
	s.ExternalIp = &v
	return s
}

func (s *DescribeForwardTableEntriesResponseBodyForwardTableEntriesForwardTableEntry) SetExternalPort(v string) *DescribeForwardTableEntriesResponseBodyForwardTableEntriesForwardTableEntry {
	s.ExternalPort = &v
	return s
}

func (s *DescribeForwardTableEntriesResponseBodyForwardTableEntriesForwardTableEntry) SetForwardEntryId(v string) *DescribeForwardTableEntriesResponseBodyForwardTableEntriesForwardTableEntry {
	s.ForwardEntryId = &v
	return s
}

func (s *DescribeForwardTableEntriesResponseBodyForwardTableEntriesForwardTableEntry) SetForwardTableId(v string) *DescribeForwardTableEntriesResponseBodyForwardTableEntriesForwardTableEntry {
	s.ForwardTableId = &v
	return s
}

func (s *DescribeForwardTableEntriesResponseBodyForwardTableEntriesForwardTableEntry) SetInternalIp(v string) *DescribeForwardTableEntriesResponseBodyForwardTableEntriesForwardTableEntry {
	s.InternalIp = &v
	return s
}

func (s *DescribeForwardTableEntriesResponseBodyForwardTableEntriesForwardTableEntry) SetInternalPort(v string) *DescribeForwardTableEntriesResponseBodyForwardTableEntriesForwardTableEntry {
	s.InternalPort = &v
	return s
}

func (s *DescribeForwardTableEntriesResponseBodyForwardTableEntriesForwardTableEntry) SetIpProtocol(v string) *DescribeForwardTableEntriesResponseBodyForwardTableEntriesForwardTableEntry {
	s.IpProtocol = &v
	return s
}

func (s *DescribeForwardTableEntriesResponseBodyForwardTableEntriesForwardTableEntry) SetStatus(v string) *DescribeForwardTableEntriesResponseBodyForwardTableEntriesForwardTableEntry {
	s.Status = &v
	return s
}

func (s *DescribeForwardTableEntriesResponseBodyForwardTableEntriesForwardTableEntry) Validate() error {
	return dara.Validate(s)
}
