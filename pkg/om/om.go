package om

import (
	log "github.com/sirupsen/logrus"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	// SYS_LOG log file
	SYS_LOG = "/var/log/messages"
	// NsenterCmd is nsenter mount command
	NsenterCmd = "/nsenter --mount=/proc/1/ns/mnt"
	// IssueMessageFile tag
	IssueMessageFile = "ISSUE_MESSAGE_FILE"
	// IssueBlockReference tag
	IssueBlockReference = "ISSUE_BLOCK_REFERENCE"
	// MessageFileLines tag
	MessageFileLines = "MESSAGE_FILE_LINES"
)

var (
	GlobalConfigVar GlobalConfig
)

// GlobalConfig save global values for om
type GlobalConfig struct {
	IssueMessageFile     bool
	MessageFileTailLines int
	IssueBlockReference  bool
}

// StorageOM, storage Operation and Maintenance
func StorageOM() {
	GlobalConfigSet()

	for {
		// fix block volume reference not removed issue;
		// error message: The device %q is still referenced from other Pods;
		if GlobalConfigVar.IssueMessageFile {
			CheckMessageFileIssue()
		}

		// loop interval time
		time.Sleep(time.Duration(time.Second * 10))
	}
}

// CheckMessageFileIssue check/fix issues from message file
func CheckMessageFileIssue() {
	// got the last few lines of message file
	lines := ReadFileLinesFromHost(SYS_LOG)

	// loop in lines
	for _, line := range lines {
		// fix block volume reference not removed issue;
		// error message: The device %q is still referenced from other Pods;
		if GlobalConfigVar.IssueBlockReference && strings.Contains(line, "is still referenced from other Pods") {
			if FixReferenceMountIssue(line) {

			}
		}
		// TODO: CheckOrphanPodIssue()
	}
}

// GlobalConfigSet set Global Config
func GlobalConfigSet() {
	GlobalConfigVar.IssueMessageFile = false
	messageFile := os.Getenv(IssueMessageFile)
	if messageFile == "true" {
		GlobalConfigVar.IssueMessageFile = true
	}

	GlobalConfigVar.MessageFileTailLines = 20
	messageFileLine := os.Getenv(MessageFileLines)
	if messageFileLine != "" {
		lineNum, err := strconv.Atoi(messageFileLine)
		if err != nil {
			log.Errorf("OM GlobalConfigSet: MessageFileLines error format: %s", messageFileLine)
		} else {
			GlobalConfigVar.MessageFileTailLines = lineNum
			if GlobalConfigVar.MessageFileTailLines > 500 {
				log.Warnf("OM GlobalConfigSet: MessageFileLines too large: %s", messageFileLine)
			}
		}
	}

	GlobalConfigVar.IssueBlockReference = false
	blockRef := os.Getenv(IssueBlockReference)
	if blockRef == "true" {
		GlobalConfigVar.IssueBlockReference = true
	}
}
