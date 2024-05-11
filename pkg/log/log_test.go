package log

import (
	"io"
	"testing"

	"github.com/sirupsen/logrus"
	"k8s.io/klog/v2"
	klogtest "k8s.io/klog/v2/test"
)

func BenchmarkLogrus(b *testing.B) {
	logrus.SetOutput(io.Discard)
	for i := 0; i < b.N; i++ {
		klog.Infof("This is a log message %d", i)
	}
}

func BenchmarkLogrusStructured(b *testing.B) {
	logrus.SetOutput(io.Discard)
	for i := 0; i < b.N; i++ {
		logrus.WithField("index", i).Info("This is a structured log message")
	}
}

func BenchmarkKlog(b *testing.B) {
	klogtest.InitKlog()
	klog.SetOutput(io.Discard)
	for i := 0; i < b.N; i++ {
		klog.Infof("This is a log message %d", i)
	}
}

func BenchmarkKlogStructured(b *testing.B) {
	klogtest.InitKlog()
	klog.SetOutput(io.Discard)
	logger := klog.Background()
	for i := 0; i < b.N; i++ {
		logger.Info("This is a structured log message", "index", i)
	}
}

func setupRedirect(tb testing.TB) {
	klogtest.InitKlog()
	klog.SetOutput(io.Discard)
	RedirectLogrusToLogr(logrus.StandardLogger(), klog.Background())
	tb.Cleanup(func() {
		logrus.StandardLogger().ReplaceHooks(make(logrus.LevelHooks))
	})
}

func BenchmarkLogrusRedirected(b *testing.B) {
	setupRedirect(b)
	for i := 0; i < b.N; i++ {
		klog.Infof("This is a log message %d", i)
	}
}

func BenchmarkLogrusStructuredRedirected(b *testing.B) {
	setupRedirect(b)
	for i := 0; i < b.N; i++ {
		logrus.WithField("index", i).Info("This is a structured log message")
	}
}
