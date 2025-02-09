package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	provisionOrgMsg = "disk size is not supported."
	attachOrgMsg1   = "aaa had volume node affinity conflict"
	attachOrgMsg2   = "xxxxxxxxxxx"
	attachOrgMsg3   = "instance does not support this disk category."
)

func TestFindSuggestionByErrorMessage(t *testing.T) {
	errMsgProvisionSample := map[string]string{
		provisionOrgMsg: provisionOrgMsg + "\n" + "faq: https://help.aliyun.com/document_detail/286495.htm#section-r6k-e7i-l78",
	}

	errMsgAttachSample := map[string]string{
		attachOrgMsg1: attachOrgMsg1 + "\n" + "faq: https://help.aliyun.com/document_detail/286495.htm#section-7o2-zf7-eff",
		attachOrgMsg2: attachOrgMsg2,
		attachOrgMsg3: attachOrgMsg3 + "\n" + "faq: https://help.aliyun.com/document_detail/286495.htm#section-ihn-gds-9mm",
	}

	for origMsg, wrappedMsg := range errMsgProvisionSample {
		assert.Equal(t, wrappedMsg, FindSuggestionByErrorMessage(origMsg, DiskProvision))
	}

	for origMsg, wrappedMsg := range errMsgAttachSample {
		assert.Equal(t, wrappedMsg, FindSuggestionByErrorMessage(origMsg, DiskAttachDetach))
	}
}
