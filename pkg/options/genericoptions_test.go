package options

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/rest"
)

func TestGetRestConfigForCRD(t *testing.T) {
	cfg := &rest.Config{}
	cfg.ContentType = runtime.ContentTypeProtobuf

	assert.Equal(t, GetRestConfigForCRD(*cfg).ContentType, runtime.ContentTypeJSON)
	assert.Equal(t, cfg.ContentType, runtime.ContentTypeProtobuf)
}
