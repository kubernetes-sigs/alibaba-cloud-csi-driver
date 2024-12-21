package utils

import (
	"os"
	"testing"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
)

func Test_ExtractFuseContainerConfig(t *testing.T) {
	dir := t.TempDir()
	assert.NoError(t, os.WriteFile(dir+"/fuse-ossfs", []byte(`
image=ossfs:latest
cpu-request=100m
cpu-limit=1
memory-request=500Mi
memory-limit=2Gi
dbglevel=info
mime-support=false
annotations={"anno1": "val1", "anno2": "val2"}
labels={"label1": "val1", "label2": "val2"}
`), 0644))

	configmap := utils.Config{Path: dir}
	config := ExtractFuseContainerConfig(&configmap, OssFsType)
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
