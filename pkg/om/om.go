package om

import (
	"os"
	"strconv"
	"strings"
	"time"

	"k8s.io/klog/v2"
)

const (
	// SysLog log file
	SysLog = "/var/log/messages"
	// IssueBlockReference tag
	IssueBlockReference = "ISSUE_BLOCK_REFERENCE"
	// MessageFileLines tag
	MessageFileLines = "MESSAGE_FILE_LINES"
)

var (
	// GlobalConfigVar var
	GlobalConfigVar GlobalConfig
)

// GlobalConfig save global values for om
type GlobalConfig struct {
	MessageFileTailLines int
	IssueBlockReference  bool
}

// StorageOM storage Operation and Maintenance
func StorageOM() {
	GlobalConfigSet()

	for {
		// fix block volume reference not removed issue;
		// error message: The device %q is still referenced from other Pods;
		if GlobalConfigVar.IssueBlockReference {
			CheckMessageFileIssue()
		}

		// loop interval time
		time.Sleep(time.Duration(time.Second * 10))
	}
}

// CheckMessageFileIssue check/fix issues from message file
func CheckMessageFileIssue() {
	// got the last few lines of message file
	lines := ReadFileLinesFromHost(SysLog)

	// loop in lines
	for _, line := range lines {
		// Fix Block Volume Reference Issue;
		if GlobalConfigVar.IssueBlockReference && strings.Contains(line, "is still referenced from other Pods") {
			if FixReferenceMountIssue(line) {

			}
		}
	}
}

// GlobalConfigSet set Global Config
func GlobalConfigSet() {
	GlobalConfigVar.MessageFileTailLines = 20
	messageFileLine := os.Getenv(MessageFileLines)
	if messageFileLine != "" {
		lineNum, err := strconv.Atoi(messageFileLine)
		if err != nil {
			klog.Errorf("OM GlobalConfigSet: MessageFileLines error format: %s", messageFileLine)
		} else {
			GlobalConfigVar.MessageFileTailLines = lineNum
			if GlobalConfigVar.MessageFileTailLines > 500 {
				klog.Warningf("OM GlobalConfigSet: MessageFileLines too large: %s", messageFileLine)
			}
		}
	}

	GlobalConfigVar.IssueBlockReference = false
	blockRef := os.Getenv(IssueBlockReference)
	if blockRef == "true" {
		GlobalConfigVar.IssueBlockReference = true
	}
}
