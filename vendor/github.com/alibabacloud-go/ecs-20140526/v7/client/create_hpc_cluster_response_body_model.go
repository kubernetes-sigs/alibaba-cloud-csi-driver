// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHpcClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHpcClusterId(v string) *CreateHpcClusterResponseBody
	GetHpcClusterId() *string
	SetRequestId(v string) *CreateHpcClusterResponseBody
	GetRequestId() *string
}

type CreateHpcClusterResponseBody struct {
	// The ID of cluster.
	//
	// example:
	//
	// hpc-pnlg1ds9rky4****
	HpcClusterId *string `json:"HpcClusterId,omitempty" xml:"HpcClusterId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateHpcClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateHpcClusterResponseBody) GoString() string {
	return s.String()
}

func (s *CreateHpcClusterResponseBody) GetHpcClusterId() *string {
	return s.HpcClusterId
}

func (s *CreateHpcClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateHpcClusterResponseBody) SetHpcClusterId(v string) *CreateHpcClusterResponseBody {
	s.HpcClusterId = &v
	return s
}

func (s *CreateHpcClusterResponseBody) SetRequestId(v string) *CreateHpcClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateHpcClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
