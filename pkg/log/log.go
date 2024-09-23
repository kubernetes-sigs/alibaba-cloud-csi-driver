package log

import (
	"io"
	"runtime"
	"strings"

	"github.com/go-logr/logr"
	"github.com/sirupsen/logrus"
)

type noopFormatter struct{}

func (f noopFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	return nil, nil
}

type logrHook struct {
	sink logr.LogSink

	v int
}

func (h logrHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

const (
	maximumCallerDepth int = 25

	// Shortest path to the first non-logrus calling function:
	// runtime.Callers
	// github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/log.getCallerDepth
	// github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/log.logrHook.Fire
	// github.com/sirupsen/logrus.LevelHooks.Fire
	// github.com/sirupsen/logrus.(*Entry).fireHooks
	// github.com/sirupsen/logrus.(*Entry).log
	// github.com/sirupsen/logrus.(*Entry).Log
	knownLogrusFrames int = 7
)

// getCallerDepth retrieves the depth of the first non-logrus calling function
func getCallerDepth() int {
	pcs := make([]uintptr, maximumCallerDepth)
	n := runtime.Callers(knownLogrusFrames, pcs)
	frames := runtime.CallersFrames(pcs[:n])

	more := n > 0
	for i := knownLogrusFrames; more; i++ {
		var frame runtime.Frame
		frame, more = frames.Next()
		if !strings.HasPrefix(frame.Function, "github.com/sirupsen/logrus") {
			return i - 2 // remove runtime.Callers frame and self
		}
	}
	return 0
}

// mimics https://github.com/go-logr/logr/blob/88d98bdb5beceebfb6bb935f2b175c722e18d119/sloghandler.go#L59
func (h *logrHook) Fire(entry *logrus.Entry) error {
	level := 0
	// follow https://github.com/kubernetes/community/blob/master/contributors/devel/sig-instrumentation/logging.md#what-method-to-use
	switch entry.Level {
	case logrus.WarnLevel:
		level = 1
	case logrus.InfoLevel:
		level = 2
	case logrus.DebugLevel:
		level = 4
	case logrus.TraceLevel:
		level = 5
	}
	level += h.v
	if !h.sink.Enabled(level) {
		return nil
	}

	kvList := make([]any, 0, len(entry.Data)*2+2)
	kvList = append(kvList, "level", entry.Level)
	for k, v := range entry.Data {
		kvList = append(kvList, k, v)
	}

	s := h.sink
	if withCallDepth, ok := s.(logr.CallDepthLogSink); ok {
		s = withCallDepth.WithCallDepth(getCallerDepth() - 1) // remove runtimeInfo.CallDepth frames
	}

	if entry.Level <= logrus.ErrorLevel {
		s.Error(nil, entry.Message, kvList...)
	} else {
		s.Info(level, entry.Message, kvList...)
	}

	return nil
}

func RedirectLogrusToLogr(logrusLogger *logrus.Logger, logrLogger logr.Logger) {
	logrusLogger.AddHook(&logrHook{sink: logrLogger.GetSink(), v: logrLogger.GetV()})
	logrusLogger.SetOutput(io.Discard)
	logrusLogger.SetFormatter(noopFormatter{})
	logrusLogger.SetReportCaller(false) // we do this in logrHook
}
