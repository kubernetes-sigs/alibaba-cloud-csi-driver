// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLaunchTemplateDefaultVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLaunchTemplateId(v string) *ModifyLaunchTemplateDefaultVersionResponseBody
	GetLaunchTemplateId() *string
	SetRequestId(v string) *ModifyLaunchTemplateDefaultVersionResponseBody
	GetRequestId() *string
}

type ModifyLaunchTemplateDefaultVersionResponseBody struct {
	// The ID of the launch template. For more information, see [DescribeLaunchTemplates](https://help.aliyun.com/document_detail/73759.html).
	//
	// example:
	//
	// lt-bp1apo0bbbkuy0rj****
	LaunchTemplateId *string `json:"LaunchTemplateId,omitempty" xml:"LaunchTemplateId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyLaunchTemplateDefaultVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyLaunchTemplateDefaultVersionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyLaunchTemplateDefaultVersionResponseBody) GetLaunchTemplateId() *string {
	return s.LaunchTemplateId
}

func (s *ModifyLaunchTemplateDefaultVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyLaunchTemplateDefaultVersionResponseBody) SetLaunchTemplateId(v string) *ModifyLaunchTemplateDefaultVersionResponseBody {
	s.LaunchTemplateId = &v
	return s
}

func (s *ModifyLaunchTemplateDefaultVersionResponseBody) SetRequestId(v string) *ModifyLaunchTemplateDefaultVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyLaunchTemplateDefaultVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
