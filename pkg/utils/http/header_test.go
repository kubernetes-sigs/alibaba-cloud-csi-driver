package http

import (
	"net/http"
	"os"
	"testing"

	"github.com/alibabacloud-go/tea/tea"
	"github.com/stretchr/testify/assert"
	"k8s.io/klog/v2"
)

func TestParseHttpHeaderEnv(t *testing.T) {
	t.Setenv("ALIBABA_CLOUD_HTTP_HEADERS", "X-key1: value1\nx-key2: value2")
	t.Setenv("ECS_HEADERS", "x-key3: value3\nx-key2: value-ecs")
	t.Setenv("X-ACSPROXY-ASCM-CONTEXT", "my-context")
	headers := MustParseHeaderEnv("ECS_HEADERS")
	assert.Equal(t, http.Header{
		"X-Key1":                  {"value1"},
		"X-Key2":                  {"value-ecs"},
		"X-Key3":                  {"value3"},
		"X-Acsproxy-Ascm-Context": {"my-context"},
	}, headers)
}

func TestParseHttpHeaderEnvInvalid(t *testing.T) {
	defer func() { klog.OsExit = os.Exit }()
	klog.OsExit = func(c int) { panic(c) }

	t.Setenv("ALIBABA_CLOUD_HTTP_HEADERS", "not-http-header")
	assert.Panics(t, func() {
		MustParseHeaderEnv("ALIBABA_CLOUD_HTTP_HEADERS")
	})
}

func TestParseHttpHeaderEnvEmpty(t *testing.T) {
	assert.Equal(t, http.Header{}, MustParseHeaderEnv("ECS_HEADERS"))
}

func TestToV2SDKHeader(t *testing.T) {
	assert.Equal(t, map[string]*string{
		"X-Key1": tea.String("value1"),
		"X-Key2": tea.String("value2"),
	}, MustToV2SDKHeaders(http.Header{
		"X-Key1":  {"value1"},
		"X-Key2":  {"value2"},
		"useless": {},
	}))
}

func TestToV2SDKHeaderMulti(t *testing.T) {
	defer func() { klog.OsExit = os.Exit }()
	klog.OsExit = func(c int) { panic(c) }

	assert.Panics(t, func() {
		MustToV2SDKHeaders(http.Header{
			"X-Key1": {"value1", "value2"},
		})
	})
}
