//go:build !windows

package customfuse

import (
	"context"
	"testing"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/features"
	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	fake "k8s.io/client-go/kubernetes/fake"
)

func TestResolveAutoCapacity(t *testing.T) {
	enableFeatureGate := func(t *testing.T) {
		t.Helper()
		assert.NoError(t, features.FunctionalMutableFeatureGate.SetFromMap(map[string]bool{
			string(features.CustomFuseAutoCapacity): true,
		}))
		t.Cleanup(func() {
			assert.NoError(t, features.FunctionalMutableFeatureGate.ResetFeatureValueToDefault(features.CustomFuseAutoCapacity))
		})
	}

	t.Run("feature gate disabled", func(t *testing.T) {
		client := fake.NewSimpleClientset(&corev1.PersistentVolume{
			ObjectMeta: metav1.ObjectMeta{Name: "pv-1"},
			Spec: corev1.PersistentVolumeSpec{
				Capacity: corev1.ResourceList{corev1.ResourceStorage: resource.MustParse("100Gi")},
			},
		})
		got := resolveAutoCapacity(context.Background(), client, "pv-1", nil)
		assert.Equal(t, "", got)
	})

	t.Run("volumeAttributes already has capacity", func(t *testing.T) {
		enableFeatureGate(t)
		client := fake.NewSimpleClientset(&corev1.PersistentVolume{
			ObjectMeta: metav1.ObjectMeta{Name: "pv-1"},
			Spec: corev1.PersistentVolumeSpec{
				Capacity: corev1.ResourceList{corev1.ResourceStorage: resource.MustParse("100Gi")},
			},
		})
		got := resolveAutoCapacity(context.Background(), client, "pv-1", map[string]string{"capacity": "50"})
		assert.Equal(t, "", got)
	})

	t.Run("PV found with capacity", func(t *testing.T) {
		enableFeatureGate(t)
		client := fake.NewSimpleClientset(&corev1.PersistentVolume{
			ObjectMeta: metav1.ObjectMeta{Name: "pv-100g"},
			Spec: corev1.PersistentVolumeSpec{
				Capacity: corev1.ResourceList{corev1.ResourceStorage: resource.MustParse("100Gi")},
			},
		})
		got := resolveAutoCapacity(context.Background(), client, "pv-100g", nil)
		assert.Equal(t, "100Gi", got)
	})

	t.Run("PV not found", func(t *testing.T) {
		enableFeatureGate(t)
		client := fake.NewSimpleClientset()
		got := resolveAutoCapacity(context.Background(), client, "nonexistent", nil)
		assert.Equal(t, "", got)
	})

	t.Run("PV has no capacity field", func(t *testing.T) {
		enableFeatureGate(t)
		client := fake.NewSimpleClientset(&corev1.PersistentVolume{
			ObjectMeta: metav1.ObjectMeta{Name: "pv-nocap"},
			Spec:       corev1.PersistentVolumeSpec{},
		})
		got := resolveAutoCapacity(context.Background(), client, "pv-nocap", nil)
		assert.Equal(t, "", got)
	})

	t.Run("PV capacity less than 1Gi returns Quantity string", func(t *testing.T) {
		enableFeatureGate(t)
		client := fake.NewSimpleClientset(&corev1.PersistentVolume{
			ObjectMeta: metav1.ObjectMeta{Name: "pv-small"},
			Spec: corev1.PersistentVolumeSpec{
				Capacity: corev1.ResourceList{corev1.ResourceStorage: resource.MustParse("500Mi")},
			},
		})
		got := resolveAutoCapacity(context.Background(), client, "pv-small", nil)
		assert.Equal(t, "500Mi", got)
	})

	t.Run("nil volContext treated as no capacity", func(t *testing.T) {
		enableFeatureGate(t)
		client := fake.NewSimpleClientset(&corev1.PersistentVolume{
			ObjectMeta: metav1.ObjectMeta{Name: "pv-1"},
			Spec: corev1.PersistentVolumeSpec{
				Capacity: corev1.ResourceList{corev1.ResourceStorage: resource.MustParse("50Gi")},
			},
		})
		got := resolveAutoCapacity(context.Background(), client, "pv-1", nil)
		assert.Equal(t, "50Gi", got)
	})
}
