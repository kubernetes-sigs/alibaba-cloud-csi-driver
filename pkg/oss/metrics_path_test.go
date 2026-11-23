package oss

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMetricsPathPrefixByRuntimeType(t *testing.T) {
	assert.Equal(t, agentMetricsPathPrefix, getMetricsPathPrefixByRuntimeType(RuntimeTypeMicroVM))
	assert.Equal(t, metricsPathPrefix, getMetricsPathPrefixByRuntimeType(RuntimeTypeRunD))
	assert.Equal(t, metricsPathPrefix, getMetricsPathPrefixByRuntimeType(RuntimeTypeRunC))
}
