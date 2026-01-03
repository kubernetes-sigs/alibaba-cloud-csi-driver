// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceScreenshotResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetInstanceScreenshotResponseBody
	GetInstanceId() *string
	SetRequestId(v string) *GetInstanceScreenshotResponseBody
	GetRequestId() *string
	SetScreenshot(v string) *GetInstanceScreenshotResponseBody
	GetScreenshot() *string
}

type GetInstanceScreenshotResponseBody struct {
	// The instance ID.
	//
	// example:
	//
	// i-bp1gbz20g229bvu5****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The Base64-encoded instance screenshot in the JPG format.
	//
	// example:
	//
	// iVBORw0KGgoA...AAABJRU5ErkJggg==
	Screenshot *string `json:"Screenshot,omitempty" xml:"Screenshot,omitempty"`
}

func (s GetInstanceScreenshotResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceScreenshotResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceScreenshotResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetInstanceScreenshotResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInstanceScreenshotResponseBody) GetScreenshot() *string {
	return s.Screenshot
}

func (s *GetInstanceScreenshotResponseBody) SetInstanceId(v string) *GetInstanceScreenshotResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceScreenshotResponseBody) SetRequestId(v string) *GetInstanceScreenshotResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceScreenshotResponseBody) SetScreenshot(v string) *GetInstanceScreenshotResponseBody {
	s.Screenshot = &v
	return s
}

func (s *GetInstanceScreenshotResponseBody) Validate() error {
	return dara.Validate(s)
}
