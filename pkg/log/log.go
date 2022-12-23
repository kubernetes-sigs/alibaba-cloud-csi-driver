package log

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"path/filepath"
	"time"
)

// LogfilePrefix prefix of log file
const LogfilePrefix = "/var/log/alicloud/"

var Log *logrus.Logger

func NewLogger(driver string) *logrus.Logger {
	if Log != nil {
		return Log
	}
	filepath := filepath.Join(LogfilePrefix, fmt.Sprintf("%s.log", driver))
	writer, _ := rotatelogs.New(
		filepath+"-%Y-%m-%d",
		rotatelogs.WithLinkName(filepath),
		rotatelogs.WithMaxAge(30*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	writerMap := lfshook.WriterMap{
		logrus.InfoLevel:  writer,
		logrus.ErrorLevel: writer,
	}

	Log = logrus.New()
	Log.Hooks.Add(lfshook.NewHook(
		writerMap,
		&logrus.JSONFormatter{},
	))

	return Log
}
