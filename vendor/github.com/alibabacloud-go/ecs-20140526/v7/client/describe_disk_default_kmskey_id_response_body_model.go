// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDiskDefaultKMSKeyIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetKMSKeyId(v string) *DescribeDiskDefaultKMSKeyIdResponseBody
	GetKMSKeyId() *string
	SetRequestId(v string) *DescribeDiskDefaultKMSKeyIdResponseBody
	GetRequestId() *string
}

type DescribeDiskDefaultKMSKeyIdResponseBody struct {
	// The ID of the KMS key.
	//
	// example:
	//
	// 0e478b7a-4262-4802-b8cb-00d3fb40****
	KMSKeyId *string `json:"KMSKeyId,omitempty" xml:"KMSKeyId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDiskDefaultKMSKeyIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiskDefaultKMSKeyIdResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDiskDefaultKMSKeyIdResponseBody) GetKMSKeyId() *string {
	return s.KMSKeyId
}

func (s *DescribeDiskDefaultKMSKeyIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDiskDefaultKMSKeyIdResponseBody) SetKMSKeyId(v string) *DescribeDiskDefaultKMSKeyIdResponseBody {
	s.KMSKeyId = &v
	return s
}

func (s *DescribeDiskDefaultKMSKeyIdResponseBody) SetRequestId(v string) *DescribeDiskDefaultKMSKeyIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDiskDefaultKMSKeyIdResponseBody) Validate() error {
	return dara.Validate(s)
}
