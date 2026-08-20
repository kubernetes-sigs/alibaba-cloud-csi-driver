// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCpfsAccessPointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyCpfsAccessPointResponseBody
	GetRequestId() *string
}

type ModifyCpfsAccessPointResponseBody struct {
	// example:
	//
	// 70EACC9C-D07A-4A34-ADA4-77506C42B023
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyCpfsAccessPointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyCpfsAccessPointResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyCpfsAccessPointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyCpfsAccessPointResponseBody) SetRequestId(v string) *ModifyCpfsAccessPointResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyCpfsAccessPointResponseBody) Validate() error {
	return dara.Validate(s)
}
