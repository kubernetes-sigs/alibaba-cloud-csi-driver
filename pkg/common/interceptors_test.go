package common

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestGetMethodName(t *testing.T) {
	tests := []struct {
		name     string
		arg      string
		expected string
	}{
		{
			name:     "Regular full name",
			arg:      "/csi.v1.Identity/GetPluginInfo",
			expected: "GetPluginInfo",
		},
		{
			name:     "Full name without '/'",
			arg:      "NodePublishVolume",
			expected: "",
		},
		{
			name:     "Full name with '/' as the last character",
			arg:      "NodePublishVolume/",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := getMethodName(tt.arg)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestGetCodeFromError(t *testing.T) {
	tests := []struct {
		name     string
		arg      error
		expected codes.Code
	}{
		{
			name:     "RPC unimplemented error",
			arg:      status.Error(codes.Unimplemented, "method not implemented"),
			expected: codes.Unimplemented,
		},
		{
			name:     "Nil error",
			arg:      nil,
			expected: codes.OK,
		},
		{
			name:     "Plain go error",
			arg:      fmt.Errorf("error message"),
			expected: codes.Unknown,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := getCodeFromError(tt.arg)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestGetPodStartTime(t *testing.T) {
	tests := []struct {
		name     string
		arg      *v1.Pod
		expected time.Time
		wantErr  bool
	}{
		{
			name: "Pod with start time",
			arg: &v1.Pod{
				Status: v1.PodStatus{
					StartTime: &metav1.Time{Time: time.Unix(100, 0)},
				},
			},
			expected: time.Unix(100, 0),
			wantErr:  false,
		},
		{
			name:     "Pod without start time",
			arg:      &v1.Pod{},
			expected: time.Time{},
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := getPodStartTime(tt.arg)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, actual)
			}
		})
	}
}
