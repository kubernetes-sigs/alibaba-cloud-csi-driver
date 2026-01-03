// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportKeyPairResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetKeyPairFingerPrint(v string) *ImportKeyPairResponseBody
	GetKeyPairFingerPrint() *string
	SetKeyPairName(v string) *ImportKeyPairResponseBody
	GetKeyPairName() *string
	SetRequestId(v string) *ImportKeyPairResponseBody
	GetRequestId() *string
}

type ImportKeyPairResponseBody struct {
	// The fingerprint of the key pair. The MD5 message-digest algorithm is used based on the public key fingerprint format defined in RFC 4716.
	//
	// example:
	//
	// 89:f0:ba:62:ac:b8:aa:e1:61:5e:fd:81:69:86:6d:6b:f0:c0:5a:**
	KeyPairFingerPrint *string `json:"KeyPairFingerPrint,omitempty" xml:"KeyPairFingerPrint,omitempty"`
	// The name of the key pair.
	//
	// example:
	//
	// testKeyPairName
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ImportKeyPairResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ImportKeyPairResponseBody) GoString() string {
	return s.String()
}

func (s *ImportKeyPairResponseBody) GetKeyPairFingerPrint() *string {
	return s.KeyPairFingerPrint
}

func (s *ImportKeyPairResponseBody) GetKeyPairName() *string {
	return s.KeyPairName
}

func (s *ImportKeyPairResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ImportKeyPairResponseBody) SetKeyPairFingerPrint(v string) *ImportKeyPairResponseBody {
	s.KeyPairFingerPrint = &v
	return s
}

func (s *ImportKeyPairResponseBody) SetKeyPairName(v string) *ImportKeyPairResponseBody {
	s.KeyPairName = &v
	return s
}

func (s *ImportKeyPairResponseBody) SetRequestId(v string) *ImportKeyPairResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImportKeyPairResponseBody) Validate() error {
	return dara.Validate(s)
}
