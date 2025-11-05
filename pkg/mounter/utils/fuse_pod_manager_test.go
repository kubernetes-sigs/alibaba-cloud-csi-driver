package utils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	corev1 "k8s.io/api/core/v1"
	apiserrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/fake"
	clientgotesting "k8s.io/client-go/testing"
	"k8s.io/klog/v2/ktesting"
)

func Test_ExtractFuseContainerConfig(t *testing.T) {
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
	config := ExtractFuseContainerConfig(configmap, OssFsType)
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

type testFuse struct{}

func (t testFuse) Name() string {
	return "test"
}
func (t testFuse) PodTemplateSpec(c *FusePodContext, target string) (*corev1.PodTemplateSpec, error) {
	return &corev1.PodTemplateSpec{
		Spec: corev1.PodSpec{
			NodeName: c.NodeName,
		},
	}, nil
}
func (t testFuse) AddDefaultMountOptions(options []string) []string {
	return options
}

func TestCreate(t *testing.T) {
	_, ctx := ktesting.NewTestContext(t)
	client := fake.NewSimpleClientset()
	client.PrependReactor("create", "pods", func(action clientgotesting.Action) (bool, runtime.Object, error) {
		pod := action.(clientgotesting.CreateAction).GetObject().(*corev1.Pod)
		pod.ResourceVersion = "1"
		return false, pod, nil
	})
	fpm := NewFusePodManager(testFuse{}, client)

	go func() {
		// slimulate kubelet
		for {
			time.Sleep(100 * time.Millisecond)
			pods, err := client.CoreV1().Pods("test-fuse").List(ctx, metav1.ListOptions{})
			if apiserrors.IsNotFound(err) {
				continue
			}
			require.NoError(t, err)
			require.Equal(t, 1, len(pods.Items))
			assert.Equal(t, "test-node", pods.Items[0].Spec.NodeName)
			pod := pods.Items[0].DeepCopy()
			pod.Status.Phase = corev1.PodRunning
			pod.Status.Conditions = []corev1.PodCondition{
				{
					Type:   corev1.PodReady,
					Status: corev1.ConditionTrue,
				},
			}
			_, err = client.CoreV1().Pods("test-fuse").UpdateStatus(ctx, pod, metav1.UpdateOptions{})
			require.NoError(t, err)
			return
		}
	}()

	_, err := fpm.Create(&FusePodContext{
		Context:   ctx,
		Namespace: "test-fuse",
		VolumeId:  "test-volume",
		NodeName:  "test-node",
	}, "/run/test-fuse")
	require.NoError(t, err)
}

func TestDelete(t *testing.T) {
	_, ctx := ktesting.NewTestContext(t)
	client := fake.NewSimpleClientset(&corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-fuse-pod",
			Namespace: "test-fuse",
			Labels:    map[string]string{FuseVolumeIdLabelKey: "test-volume"},
		},
	})

	fpm := NewFusePodManager(testFuse{}, client)
	err := fpm.Delete(&FusePodContext{
		Context:   ctx,
		Namespace: "test-fuse",
		VolumeId:  "test-volume",
		NodeName:  "test-node",
	})
	require.NoError(t, err)
	pods, err := client.CoreV1().Pods("test-fuse").List(ctx, metav1.ListOptions{})
	require.NoError(t, err)
	assert.Empty(t, pods.Items)
}
