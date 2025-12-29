// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceConsoleOutputResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConsoleOutput(v string) *GetInstanceConsoleOutputResponseBody
	GetConsoleOutput() *string
	SetInstanceId(v string) *GetInstanceConsoleOutputResponseBody
	GetInstanceId() *string
	SetLastUpdateTime(v string) *GetInstanceConsoleOutputResponseBody
	GetLastUpdateTime() *string
	SetRequestId(v string) *GetInstanceConsoleOutputResponseBody
	GetRequestId() *string
}

type GetInstanceConsoleOutputResponseBody struct {
	// The Base64-encoded command output of the instance.
	//
	// example:
	//
	// V2VsY29tZSB0byBDZW50T1MgCgpDaGVja2luZyBmaWxlc3lzdGVtcwpDaGVja2luZyBhbGwgZmlsZSBzeXN0ZW1zLgpbL3NiaW4vZnNjay5leHQ0ICgxKSAtLSAvXSBmc2NrLmV4dDQgLWEgL2Rldi92ZGExIAovZGV2L3ZkYTE6IGNsZWFuLCAzMjAxNi8yNjIxNDQwIGZpbGVzLCA0NDc5NzQvMTA0ODU1MDQgYmxvY2tzCgpFbnRlcmluZyBub24taW50ZXJhY3RpdmUgc3RhcnR1cApDYWxsaW5nIHRoZSBzeXN0ZW0gYWN0aXZpdHkgZGF0YSBjb2xsZWN0b3IgKHNhZGMpLi4uIAoKQnJpbmdpbmcgdXAgaW50ZXJmYWNlIGV0aDA6ICAKRGV0ZXJtaW5pbmcgSVAgaW5mb3JtYXRpb24gZm9yIGV0aDAuLi4gZG9uZS4KCmFsaXl1bi1zZXJ2aWNlIHN0YXJ0L3J1bm5pbmcsIHByb2Nlc3MgMTczMwpmaW5pc2hlZAoKQ2VudE9TIHJlbGVhc2UgNi44IChGaW5hbCkKS2VybmVsIDIuNi4zMi02OTYuMy4yLmVsNi5pNjg2IG9uIGFuIGk2ODYKCmlaMnplZDk2ZTQ2MmF5cjBxeioqKioqIGxvZ2luOg==
	ConsoleOutput *string `json:"ConsoleOutput,omitempty" xml:"ConsoleOutput,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// i-bp1c1xhsrac2coiw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The time when the last log entry was generated in the Linux kernel. The time follows the ISO 8601 standard in the yyyy-MM-ddThh:mmZ format. The time is displayed in UTC+8.
	//
	// example:
	//
	// 2018-03-22 10:04:57
	LastUpdateTime *string `json:"LastUpdateTime,omitempty" xml:"LastUpdateTime,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetInstanceConsoleOutputResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceConsoleOutputResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceConsoleOutputResponseBody) GetConsoleOutput() *string {
	return s.ConsoleOutput
}

func (s *GetInstanceConsoleOutputResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetInstanceConsoleOutputResponseBody) GetLastUpdateTime() *string {
	return s.LastUpdateTime
}

func (s *GetInstanceConsoleOutputResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInstanceConsoleOutputResponseBody) SetConsoleOutput(v string) *GetInstanceConsoleOutputResponseBody {
	s.ConsoleOutput = &v
	return s
}

func (s *GetInstanceConsoleOutputResponseBody) SetInstanceId(v string) *GetInstanceConsoleOutputResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceConsoleOutputResponseBody) SetLastUpdateTime(v string) *GetInstanceConsoleOutputResponseBody {
	s.LastUpdateTime = &v
	return s
}

func (s *GetInstanceConsoleOutputResponseBody) SetRequestId(v string) *GetInstanceConsoleOutputResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceConsoleOutputResponseBody) Validate() error {
	return dara.Validate(s)
}
