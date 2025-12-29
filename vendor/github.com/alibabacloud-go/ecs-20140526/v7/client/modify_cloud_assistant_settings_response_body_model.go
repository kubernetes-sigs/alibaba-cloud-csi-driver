// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCloudAssistantSettingsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyCloudAssistantSettingsResponseBody
	GetRequestId() *string
}

type ModifyCloudAssistantSettingsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyCloudAssistantSettingsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyCloudAssistantSettingsResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyCloudAssistantSettingsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyCloudAssistantSettingsResponseBody) SetRequestId(v string) *ModifyCloudAssistantSettingsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyCloudAssistantSettingsResponseBody) Validate() error {
	return dara.Validate(s)
}
