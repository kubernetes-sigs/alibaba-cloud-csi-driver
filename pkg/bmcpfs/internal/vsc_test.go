/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package internal

import (
	"errors"
	"fmt"
	"testing"

	"github.com/alibabacloud-go/tea/tea"
	"github.com/stretchr/testify/assert"
)

func TestAttachNotSupportedError_Error(t *testing.T) {
	err := newAttachNotSupportedError(fmt.Errorf("test error"), "vol-123", "vsc-456")
	expected := "volumeID: vol-123vscId: vsc-456test error"
	assert.Equal(t, expected, err.Error())
}

func TestNewAttachNotSupportedError(t *testing.T) {
	err := NewAttachNotSupportedError(fmt.Errorf("test error"), "vol-123", "vsc-456")
	assert.NotNil(t, err)
	assert.Equal(t, "volumeID: vol-123vscId: vsc-456test error", err.Error())
}

func TestIsAttachNotSupportedError_NilError(t *testing.T) {
	assert.False(t, IsAttachNotSupportedError(nil))
}

func TestIsAttachNotSupportedError_AttachNotSupportedError(t *testing.T) {
	err := newAttachNotSupportedError(fmt.Errorf("test error"), "vol-123", "vsc-456")
	assert.True(t, IsAttachNotSupportedError(err))
}

func TestIsAttachNotSupportedError_SDKErrorWithCorrectCode(t *testing.T) {
	sdkErr := &tea.SDKError{
		Code: tea.String(VscAttachNotSupported),
	}
	assert.True(t, IsAttachNotSupportedError(sdkErr))
}

func TestIsAttachNotSupportedError_SDKErrorWithWrongCode(t *testing.T) {
	sdkErr := &tea.SDKError{
		Code: tea.String("SomeOtherError"),
	}
	assert.False(t, IsAttachNotSupportedError(sdkErr))
}

func TestIsAttachNotSupportedError_GenericError(t *testing.T) {
	err := errors.New("generic error")
	assert.False(t, IsAttachNotSupportedError(err))
}

func TestIsAttachNotSupportedError_SDKErrorWithAttachNotSupportedInMessage(t *testing.T) {
	sdkErr := &tea.SDKError{
		Message: tea.String("Some error with AttachVscTarget.VscAttachNotSupported in message"),
		Code:    tea.String("DifferentCode"),
	}
	// Should be false because the code doesn't match, even if message contains the string
	assert.False(t, IsAttachNotSupportedError(sdkErr))
}
