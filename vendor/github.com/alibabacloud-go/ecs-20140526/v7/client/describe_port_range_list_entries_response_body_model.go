// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePortRangeListEntriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEntries(v []*DescribePortRangeListEntriesResponseBodyEntries) *DescribePortRangeListEntriesResponseBody
	GetEntries() []*DescribePortRangeListEntriesResponseBodyEntries
	SetRequestId(v string) *DescribePortRangeListEntriesResponseBody
	GetRequestId() *string
}

type DescribePortRangeListEntriesResponseBody struct {
	// Port list entries.
	Entries []*DescribePortRangeListEntriesResponseBodyEntries `json:"Entries,omitempty" xml:"Entries,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 882304EC-5CE2-5860-98ED-3FA1D8D74A0C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePortRangeListEntriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePortRangeListEntriesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePortRangeListEntriesResponseBody) GetEntries() []*DescribePortRangeListEntriesResponseBodyEntries {
	return s.Entries
}

func (s *DescribePortRangeListEntriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePortRangeListEntriesResponseBody) SetEntries(v []*DescribePortRangeListEntriesResponseBodyEntries) *DescribePortRangeListEntriesResponseBody {
	s.Entries = v
	return s
}

func (s *DescribePortRangeListEntriesResponseBody) SetRequestId(v string) *DescribePortRangeListEntriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePortRangeListEntriesResponseBody) Validate() error {
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

type DescribePortRangeListEntriesResponseBodyEntries struct {
	// The description of the port range.
	//
	// example:
	//
	// Description information of PortRangeList
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The port range.
	//
	// example:
	//
	// 80/80
	PortRange *string `json:"PortRange,omitempty" xml:"PortRange,omitempty"`
}

func (s DescribePortRangeListEntriesResponseBodyEntries) String() string {
	return dara.Prettify(s)
}

func (s DescribePortRangeListEntriesResponseBodyEntries) GoString() string {
	return s.String()
}

func (s *DescribePortRangeListEntriesResponseBodyEntries) GetDescription() *string {
	return s.Description
}

func (s *DescribePortRangeListEntriesResponseBodyEntries) GetPortRange() *string {
	return s.PortRange
}

func (s *DescribePortRangeListEntriesResponseBodyEntries) SetDescription(v string) *DescribePortRangeListEntriesResponseBodyEntries {
	s.Description = &v
	return s
}

func (s *DescribePortRangeListEntriesResponseBodyEntries) SetPortRange(v string) *DescribePortRangeListEntriesResponseBodyEntries {
	s.PortRange = &v
	return s
}

func (s *DescribePortRangeListEntriesResponseBodyEntries) Validate() error {
	return dara.Validate(s)
}
