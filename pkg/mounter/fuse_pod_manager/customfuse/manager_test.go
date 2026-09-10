package customfuse

import (
	"context"
	"testing"
	"time"

	mounterutils "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/kubernetes/fake"
	clientgotesting "k8s.io/client-go/testing"
	"k8s.io/klog/v2/ktesting"
)

const testSyncTimeout = 200 * time.Millisecond

func testConfigMap(data map[string]string) *corev1.ConfigMap {
	return &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{Name: configMapName, Namespace: configMapNamespace},
		Data:       data,
	}
}

// forbiddenConfigMaps makes every list and watch of the configmap fail the way a
// missing RBAC rule does, so the informer never syncs and keeps retrying.
func forbiddenConfigMaps(client *fake.Clientset) *fake.Clientset {
	forbidden := apierrors.NewForbidden(schema.GroupResource{Resource: "configmaps"}, configMapName, nil)
	for _, verb := range []string{"list", "watch"} {
		client.PrependReactor(verb, "configmaps", func(clientgotesting.Action) (bool, runtime.Object, error) {
			return true, nil, forbidden
		})
	}
	return client
}

func TestStartWithoutClient(t *testing.T) {
	_, ctx := ktesting.NewTestContext(t)

	require.NoError(t, NewCustomFuse(utils.Config{}, nil).Start(ctx))
}

func TestStartLoadsConfigMap(t *testing.T) {
	_, baseCtx := ktesting.NewTestContext(t)
	ctx, cancel := context.WithCancel(baseCtx)
	defer cancel()

	client := fake.NewSimpleClientset(testConfigMap(map[string]string{
		"fuse-" + mounterutils.CustomFuseType: "image=fuse-image\n",
	}))
	f := NewCustomFuse(utils.Config{}, client)
	f.syncTimeout = testSyncTimeout

	require.NoError(t, f.Start(ctx))
	assert.Equal(t, "fuse-image", f.resolveConfig(mounterutils.CustomFuseType).Image)
}

// A list the RBAC rule does not authorize never starts succeeding, so Start has to
// give up on its own instead of blocking until the kubelet kills the container.
func TestStartReportsForbiddenConfigMap(t *testing.T) {
	_, baseCtx := ktesting.NewTestContext(t)
	ctx, cancel := context.WithCancel(baseCtx)
	defer cancel()

	f := NewCustomFuse(utils.Config{}, forbiddenConfigMaps(fake.NewSimpleClientset()))
	f.syncTimeout = testSyncTimeout

	start := time.Now()
	err := f.Start(ctx)
	require.Error(t, err)
	assert.Less(t, time.Since(start), 5*time.Second, "Start outlived the sync timeout")
	assert.ErrorContains(t, err, "did not sync")
	assert.ErrorContains(t, err, "list/watch")
	assert.ErrorContains(t, err, "AuthorizeWithSelectors")
}

// A caller that goes away must be reported as cancellation rather than as the RBAC
// problem, which is what the sync-timeout message claims.
func TestStartReportsCancelledContext(t *testing.T) {
	_, baseCtx := ktesting.NewTestContext(t)
	ctx, cancel := context.WithCancel(baseCtx)

	f := NewCustomFuse(utils.Config{}, forbiddenConfigMaps(fake.NewSimpleClientset()))
	f.syncTimeout = time.Hour
	cancel()

	err := f.Start(ctx)
	require.Error(t, err)
	assert.ErrorIs(t, err, context.Canceled)
}

func TestResolveConfig(t *testing.T) {
	f := NewCustomFuse(utils.Config{ConfigMap: map[string]string{
		"fuse-" + mounterutils.CustomFuseType: "image=default-image\n",
	}}, nil)
	require.Equal(t, "default-image", f.defaultConfig.Image)

	assert.Equal(t, f.defaultConfig, f.resolveConfig(mounterutils.CustomFuseType), "before any configmap is seen")

	f.updateFromConfigMap(testConfigMap(map[string]string{
		"fuse-" + mounterutils.CustomFuseType: "image=configmap-image\nmemory-limit=2Gi\n",
	}))
	config := f.resolveConfig(mounterutils.CustomFuseType)
	assert.Equal(t, "configmap-image", config.Image)
	assert.Equal(t, "2Gi", config.Resources.Limits.Memory().String())

	// Without an image the configmap entry cannot start a container, so the whole
	// entry is dropped instead of overriding the default with a half of it.
	f.updateFromConfigMap(testConfigMap(map[string]string{
		"fuse-" + mounterutils.CustomFuseType: "memory-limit=2Gi\n",
	}))
	assert.Equal(t, f.defaultConfig, f.resolveConfig(mounterutils.CustomFuseType))

	f.clearConfigMap()
	assert.Equal(t, f.defaultConfig, f.resolveConfig(mounterutils.CustomFuseType), "after the configmap is deleted")
}
