package mounter

import (
	"reflect"
	"testing"

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
			`,
		},
	}
	config := extractFuseContainerConfig(configmap, FuseOssfs)
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
		Image: "ossfs:latest",
		Extra: map[string]string{
			"dbglevel": "info",
		},
	}
	if !reflect.DeepEqual(config, expected) {
		t.Fail()
	}
}
