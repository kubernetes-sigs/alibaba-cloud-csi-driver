package mounter

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
)

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
		Image: "",
		Extra: map[string]string{
			"dbglevel":     "info",
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
	// if !reflect.DeepEqual(config, expected) {
	// 	t.Fail()
	// }

	// invalid value
	configmap = &corev1.ConfigMap{
		Data: map[string]string{
			"fuse-ossfs": `
				annotations={"anno1": "val1", "anno2": "invalid@#$%^"}
				labels={"label1": "val1", "label2": "invalid@#$%^"}
			`,
		},
	}
	config = extractFuseContainerConfig(configmap, "ossfs")
	expected = FuseContainerConfig{}
	if !reflect.DeepEqual(config.Annotations, expected.Annotations) ||
		!reflect.DeepEqual(config.Labels, expected.Labels) {
		t.Fail()
	}
}
