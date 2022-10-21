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
		provisionOrgMsg: provisionOrgMsg + "\n" + "recommand: Please adjust the size of the cloud disk；\n reference：https://help.aliyun.com/document_detail/25412.html#BlockStorageQuota",
	}

	errMsgAttachSample := map[string]string{
		attachOrgMsg1: attachOrgMsg1 + "\n" + "recommand: Please check the node affinity of the pvc/pv and the pod",
		attachOrgMsg2: attachOrgMsg2,
		attachOrgMsg3: attachOrgMsg3 + "\n" + "recommand: Please change the cloud disk type or current ecs instance type. \n reference: https://help.aliyun.com/document_detail/25378.html",
	}

	for origMsg, wrapperedMsg := range errMsgProvisionSample {
		assert.Equal(t, wrapperedMsg, FindSuggestionByErrorMessage(origMsg, DiskProvision))
	}

	for origMsg, wrapperedMsg := range errMsgAttachSample {
		assert.Equal(t, wrapperedMsg, FindSuggestionByErrorMessage(origMsg, DiskAttachDetach))
	}
}
