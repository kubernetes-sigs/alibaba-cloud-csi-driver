// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateActivationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetActivationCode(v string) *CreateActivationResponseBody
	GetActivationCode() *string
	SetActivationId(v string) *CreateActivationResponseBody
	GetActivationId() *string
	SetRequestId(v string) *CreateActivationResponseBody
	GetRequestId() *string
}

type CreateActivationResponseBody struct {
	// The value of the activation code. The value is returned only once after the CreateActivation operation is called and cannot be queried afterward. Properly save the return value.
	//
	// example:
	//
	// a-hz0ch3SwhOlE1234+Xo32lAZC****
	ActivationCode *string `json:"ActivationCode,omitempty" xml:"ActivationCode,omitempty"`
	// The ID of the activation code.
	//
	// example:
	//
	// 4ECEEE12-56F1-4FBC-9AB1-890F1234****
	ActivationId *string `json:"ActivationId,omitempty" xml:"ActivationId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 4ECEEE12-56F1-4FBC-9AB1-890F1234****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateActivationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateActivationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateActivationResponseBody) GetActivationCode() *string {
	return s.ActivationCode
}

func (s *CreateActivationResponseBody) GetActivationId() *string {
	return s.ActivationId
}

func (s *CreateActivationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateActivationResponseBody) SetActivationCode(v string) *CreateActivationResponseBody {
	s.ActivationCode = &v
	return s
}

func (s *CreateActivationResponseBody) SetActivationId(v string) *CreateActivationResponseBody {
	s.ActivationId = &v
	return s
}

func (s *CreateActivationResponseBody) SetRequestId(v string) *CreateActivationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateActivationResponseBody) Validate() error {
	return dara.Validate(s)
}
