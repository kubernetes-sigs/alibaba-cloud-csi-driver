package om

import (
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGlobalConfigSet_Defaults(t *testing.T) {
	_ = os.Unsetenv(MessageFileLines)
	_ = os.Unsetenv(IssueBlockReference)

	GlobalConfigSet()

	assert.Equal(t, 20, GlobalConfigVar.MessageFileTailLines, "expected default MessageFileTailLines=20")
	assert.False(t, GlobalConfigVar.IssueBlockReference, "expected default IssueBlockReference=false")
}

func TestGlobalConfigSet_MessageFileLinesEnv(t *testing.T) {
	_ = os.Setenv(MessageFileLines, "50")
	defer func() { _ = os.Unsetenv(MessageFileLines) }()
	_ = os.Unsetenv(IssueBlockReference)

	GlobalConfigSet()

	assert.Equal(t, 50, GlobalConfigVar.MessageFileTailLines, "expected MessageFileTailLines=50")
}

func TestGlobalConfigSet_MessageFileLinesInvalidEnv(t *testing.T) {
	_ = os.Setenv(MessageFileLines, "notanumber")
	defer func() { _ = os.Unsetenv(MessageFileLines) }()

	GlobalConfigSet()

	// On invalid value, it should fall back to default 20
	assert.Equal(t, 20, GlobalConfigVar.MessageFileTailLines, "expected fallback MessageFileTailLines=20")
}

func TestGlobalConfigSet_MessageFileLinesLarge(t *testing.T) {
	_ = os.Setenv(MessageFileLines, strconv.Itoa(501))
	defer func() { _ = os.Unsetenv(MessageFileLines) }()

	GlobalConfigSet()

	// Large values should still be set (just warned)
	assert.Equal(t, 501, GlobalConfigVar.MessageFileTailLines, "expected MessageFileTailLines=501")
}

func TestGlobalConfigSet_IssueBlockReferenceTrue(t *testing.T) {
	_ = os.Unsetenv(MessageFileLines)
	_ = os.Setenv(IssueBlockReference, "true")
	defer func() { _ = os.Unsetenv(IssueBlockReference) }()

	GlobalConfigSet()

	assert.True(t, GlobalConfigVar.IssueBlockReference, "expected IssueBlockReference=true")
}

func TestGlobalConfigSet_IssueBlockReferenceFalse(t *testing.T) {
	_ = os.Unsetenv(MessageFileLines)
	_ = os.Setenv(IssueBlockReference, "false")
	defer func() { _ = os.Unsetenv(IssueBlockReference) }()

	GlobalConfigSet()

	assert.False(t, GlobalConfigVar.IssueBlockReference, "expected IssueBlockReference=false")
}

func TestConstants(t *testing.T) {
	assert.NotEmpty(t, SysLog, "SysLog constant should not be empty")
	assert.NotEmpty(t, IssueBlockReference, "IssueBlockReference constant should not be empty")
	assert.NotEmpty(t, MessageFileLines, "MessageFileLines constant should not be empty")
}
