package log

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

func Infof(storageType, reason, message string, args ...interface{}) {
	logrus.Infof(msgFormat(storageType, reason, message), args...)
}

func Warningf(storageType, reason, message string, args ...interface{}) {
	logrus.Warningf(msgFormat(storageType, reason, message), args...)
}

func Errorf(storageType, reason, message string, args ...interface{}) {
	logrus.Errorf(msgFormat(storageType, reason, message), args...)
}

func Fatalf(storageType, reason, message string, args ...interface{}) {
	logrus.Fatalf(msgFormat(storageType, reason, message), args...)
}

func msgFormat(storageType, reason, message string) string {
	msg := fmt.Sprintf("[%s], Code: %s, Message: %s", storageType, reason, message)
	if storageType == "" {
		msg = fmt.Sprintf("Code: %s, Message: %s", reason, message)
	}
	return msg
}
