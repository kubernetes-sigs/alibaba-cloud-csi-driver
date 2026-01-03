// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableDiskEncryptionByDefaultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisableDiskEncryptionByDefaultResponseBody
	GetRequestId() *string
}

type DisableDiskEncryptionByDefaultResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableDiskEncryptionByDefaultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableDiskEncryptionByDefaultResponseBody) GoString() string {
	return s.String()
}

func (s *DisableDiskEncryptionByDefaultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableDiskEncryptionByDefaultResponseBody) SetRequestId(v string) *DisableDiskEncryptionByDefaultResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableDiskEncryptionByDefaultResponseBody) Validate() error {
	return dara.Validate(s)
}
