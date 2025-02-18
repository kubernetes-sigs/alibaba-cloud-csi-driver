/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http:// www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package dfs

import (
	"context"
	"testing"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
)

func Test_nodeServer_setPodVscAttachedCondition(t *testing.T) {
	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: "default",
			Name:      "test",
		},
		Status: corev1.PodStatus{
			Conditions: []corev1.PodCondition{
				{
					Type:   corev1.PodScheduled,
					Status: corev1.ConditionTrue,
				},
			},
		},
	}
	clientset := fake.NewSimpleClientset(pod)

	n := &nodeServer{kubeClient: clientset}
	req := &csi.NodePublishVolumeRequest{
		VolumeContext: map[string]string{
			utils.PodNameKey:      pod.Name,
			utils.PodNamespaceKey: pod.Namespace,
		},
	}
	ctx, _ := utils.WithPodInfo(context.Background(), clientset, req)

	{
		err := n.setPodVscAttachedCondition(ctx, req)
		assert.NoError(t, err)
		pod1, err := clientset.CoreV1().Pods(pod.Namespace).Get(context.Background(), pod.Name, metav1.GetOptions{})
		assert.NoError(t, err)
		assert.Len(t, pod1.Status.Conditions, 2)
		assertConditionFound(t, pod1, corev1.PodScheduled, corev1.ConditionTrue)
		assertConditionFound(t, pod1, ConditionVscAttached, corev1.ConditionTrue)
	}

	{
		err := n.setPodVscAttachedCondition(ctx, req)
		assert.NoError(t, err)
		pod1, err := clientset.CoreV1().Pods(pod.Namespace).Get(context.Background(), pod.Name, metav1.GetOptions{})
		assert.NoError(t, err)
		assert.Len(t, pod1.Status.Conditions, 2)
		assertConditionFound(t, pod1, corev1.PodScheduled, corev1.ConditionTrue)
		assertConditionFound(t, pod1, ConditionVscAttached, corev1.ConditionTrue)
	}
}

func assertConditionFound(t *testing.T, pod *corev1.Pod, conditionType corev1.PodConditionType, conditionStatus corev1.ConditionStatus) {
	t.Helper()
	for _, cond := range pod.Status.Conditions {
		if cond.Type == conditionType && cond.Status == conditionStatus {
			return
		}
	}
	assert.Fail(t, "Condition not found: \"%s=%s\"", conditionType, conditionStatus)
}
