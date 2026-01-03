// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDiskEncryptionByDefaultStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEncrypted(v bool) *DescribeDiskEncryptionByDefaultStatusResponseBody
	GetEncrypted() *bool
	SetRequestId(v string) *DescribeDiskEncryptionByDefaultStatusResponseBody
	GetRequestId() *string
}

type DescribeDiskEncryptionByDefaultStatusResponseBody struct {
	// Indicates whether account-level default encryption of EBS resources is enabled in the region. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// Default value: false.
	//
	// example:
	//
	// false
	Encrypted *bool `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDiskEncryptionByDefaultStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiskEncryptionByDefaultStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDiskEncryptionByDefaultStatusResponseBody) GetEncrypted() *bool {
	return s.Encrypted
}

func (s *DescribeDiskEncryptionByDefaultStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDiskEncryptionByDefaultStatusResponseBody) SetEncrypted(v bool) *DescribeDiskEncryptionByDefaultStatusResponseBody {
	s.Encrypted = &v
	return s
}

func (s *DescribeDiskEncryptionByDefaultStatusResponseBody) SetRequestId(v string) *DescribeDiskEncryptionByDefaultStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDiskEncryptionByDefaultStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
