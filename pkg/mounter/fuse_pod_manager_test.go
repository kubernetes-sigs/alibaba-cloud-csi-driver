package mounter

import (
	"testing"

	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
)

func TestParseConfigLine(t *testing.T) {
	tests := []struct {
		name     string
		line     string
		config   *FuseContainerConfig
		expected *FuseContainerConfig
		wantErr  bool
	}{
		{
			name:     "empty line",
			line:     "",
			config:   &FuseContainerConfig{},
			expected: &FuseContainerConfig{},
			wantErr:  false,
		},
		{
			name:     "nil config",
			line:     "key=value",
			config:   nil,
			expected: nil,
			wantErr:  false,
		},
		{
			name:     "empty key",
			line:     "=value",
			config:   &FuseContainerConfig{},
			expected: &FuseContainerConfig{},
			wantErr:  true,
		},
		{
			name:     "valid dbglevel",
			line:     "dbglevel=debug",
			config:   &FuseContainerConfig{},
			expected: &FuseContainerConfig{Dbglevel: "debug"},
			wantErr:  false,
		},
		{
			name:     "invalid dbglevel",
			line:     "dbglevel=invalid",
			config:   &FuseContainerConfig{},
			expected: &FuseContainerConfig{},
			wantErr:  true,
		},
		{
			name:     "image config",
			line:     "image=nginx",
			config:   &FuseContainerConfig{},
			expected: &FuseContainerConfig{},
			wantErr:  true,
		},
		{
			name:     "image-tag config",
			line:     "image-tag=latest",
			config:   &FuseContainerConfig{},
			expected: &FuseContainerConfig{ImageTag: "latest"},
			wantErr:  false,
		},
		{
			name:     "cpu-request config",
			line:     "cpu-request=500m",
			config:   &FuseContainerConfig{Resources: corev1.ResourceRequirements{Requests: make(corev1.ResourceList)}},
			expected: &FuseContainerConfig{Resources: corev1.ResourceRequirements{Requests: corev1.ResourceList{"cpu": resource.MustParse("500m")}}},
			wantErr:  false,
		},
		{
			name:     "cpu-limit config",
			line:     "cpu-limit=1",
			config:   &FuseContainerConfig{Resources: corev1.ResourceRequirements{Limits: make(corev1.ResourceList)}},
			expected: &FuseContainerConfig{Resources: corev1.ResourceRequirements{Limits: corev1.ResourceList{"cpu": resource.MustParse("1")}}},
			wantErr:  false,
		},
		{
			name:     "invalid cpu-request",
			line:     "cpu-request=invalid",
			config:   &FuseContainerConfig{},
			expected: &FuseContainerConfig{},
			wantErr:  true,
		},
		{
			name:     "annotations config",
			line:     `annotations={"key":"value"}`,
			config:   &FuseContainerConfig{},
			expected: &FuseContainerConfig{Annotations: map[string]string{"key": "value"}},
			wantErr:  false,
		},
		{
			name:     "invalid",
			line:     "labels=invalid",
			config:   &FuseContainerConfig{},
			expected: &FuseContainerConfig{},
			wantErr:  true,
		},
		{
			name:     "customKey",
			line:     "customKey=customValue",
			config:   &FuseContainerConfig{},
			expected: &FuseContainerConfig{Extra: map[string]string{"customKey": "customValue"}},
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := parseConfigLine(tt.line, tt.config)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.expected, tt.config)
		})
	}
}

func Test_extractFuseContainerConfig(t *testing.T) {
	configmap := &corev1.ConfigMap{
		Data: map[string]string{
			"fuse-ossfs": `
				image=ossfs:latest
				cpu-request=100m
				cpu-limit=1
				memory-request=500Mi
				memory-limit=2Gi
				dbglevel=info
				mime-support=false
				annotations={"anno1": "val1", "anno2": "val2"}
				labels={"label1": "val1", "label2": "val2"}
			`,
		},
	}
	config := extractFuseContainerConfig(configmap, "ossfs")
	expected := FuseContainerConfig{
		Resources: corev1.ResourceRequirements{
			Limits: corev1.ResourceList{
				corev1.ResourceCPU:    resource.MustParse("1"),
				corev1.ResourceMemory: resource.MustParse("2Gi"),
			},
			Requests: corev1.ResourceList{
				corev1.ResourceCPU:    resource.MustParse("100m"),
				corev1.ResourceMemory: resource.MustParse("500Mi"),
			},
		},
		Image:    "",
		Dbglevel: "info",
		Extra: map[string]string{
			"mime-support": "false",
		},
		Annotations: map[string]string{
			"anno1": "val1",
			"anno2": "val2",
		},
		Labels: map[string]string{
			"label1": "val1",
			"label2": "val2",
		},
	}

	assert.Equal(t, expected, config)
}
