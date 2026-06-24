package cloud

import (
	"errors"

	"github.com/alibabacloud-go/tea/tea"
	"k8s.io/utils/ptr"
)

// ErrorCodeV2 returns the Alibaba error code carried by a v2 SDK error
// (*tea.SDKError), or "" if err is not (and does not wrap) one. The code is read
// straight from the SDK error, so callers can match it without depending on the
// error having been through a logging decorator.
func ErrorCodeV2(err error) string {
	if sdk, ok := errors.AsType[*tea.SDKError](err); ok {
		return ptr.Deref(sdk.Code, "")
	}
	return ""
}
