// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKeyPairResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetKeyPairFingerPrint(v string) *CreateKeyPairResponseBody
	GetKeyPairFingerPrint() *string
	SetKeyPairId(v string) *CreateKeyPairResponseBody
	GetKeyPairId() *string
	SetKeyPairName(v string) *CreateKeyPairResponseBody
	GetKeyPairName() *string
	SetPrivateKeyBody(v string) *CreateKeyPairResponseBody
	GetPrivateKeyBody() *string
	SetRequestId(v string) *CreateKeyPairResponseBody
	GetRequestId() *string
}

type CreateKeyPairResponseBody struct {
	// The fingerprint of the key pair. The message-digest algorithm 5 (MD5) is used based on the public key fingerprint format defined in RFC 4716. For more information, see [RFC 4716](https://tools.ietf.org/html/rfc4716).
	//
	// example:
	//
	// 89:f0:ba:62:ac:b8:aa:e1:61:5e:fd:81:69:86:6d:6b:f0:c0:5a:**
	KeyPairFingerPrint *string `json:"KeyPairFingerPrint,omitempty" xml:"KeyPairFingerPrint,omitempty"`
	// The ID of the key pair.
	//
	// example:
	//
	// ssh-bp67acfmxazb4p****
	KeyPairId *string `json:"KeyPairId,omitempty" xml:"KeyPairId,omitempty"`
	// The name of the key pair.
	//
	// example:
	//
	// testKeyPairName
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// The private key of the key pair. The private key is encoded with PEM in the PKCS#8 format.
	//
	// example:
	//
	// MIIEpAIBAAKCAQEAtReyMzLIcBH78EV2zj****
	PrivateKeyBody *string `json:"PrivateKeyBody,omitempty" xml:"PrivateKeyBody,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateKeyPairResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateKeyPairResponseBody) GoString() string {
	return s.String()
}

func (s *CreateKeyPairResponseBody) GetKeyPairFingerPrint() *string {
	return s.KeyPairFingerPrint
}

func (s *CreateKeyPairResponseBody) GetKeyPairId() *string {
	return s.KeyPairId
}

func (s *CreateKeyPairResponseBody) GetKeyPairName() *string {
	return s.KeyPairName
}

func (s *CreateKeyPairResponseBody) GetPrivateKeyBody() *string {
	return s.PrivateKeyBody
}

func (s *CreateKeyPairResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateKeyPairResponseBody) SetKeyPairFingerPrint(v string) *CreateKeyPairResponseBody {
	s.KeyPairFingerPrint = &v
	return s
}

func (s *CreateKeyPairResponseBody) SetKeyPairId(v string) *CreateKeyPairResponseBody {
	s.KeyPairId = &v
	return s
}

func (s *CreateKeyPairResponseBody) SetKeyPairName(v string) *CreateKeyPairResponseBody {
	s.KeyPairName = &v
	return s
}

func (s *CreateKeyPairResponseBody) SetPrivateKeyBody(v string) *CreateKeyPairResponseBody {
	s.PrivateKeyBody = &v
	return s
}

func (s *CreateKeyPairResponseBody) SetRequestId(v string) *CreateKeyPairResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateKeyPairResponseBody) Validate() error {
	return dara.Validate(s)
}
