package disk

import (
	"testing"

	crdv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Test_filterCRD(t *testing.T) {
	snapshotCRDNames := map[string]string{
		"crd1": "",
		"crd2": "",
	}

	t.Run("CRD is nil", func(t *testing.T) {
		crd := (*crdv1.CustomResourceDefinition)(nil)
		volumeGroupSnapshotEnable := true
		expected := true

		result := filterCRD(snapshotCRDNames, crd, volumeGroupSnapshotEnable)

		if result != expected {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})

	t.Run("CRD name not in snapshotCRDNames", func(t *testing.T) {
		crd := &crdv1.CustomResourceDefinition{
			ObjectMeta: metav1.ObjectMeta{
				Name: "crd3",
			},
		}
		volumeGroupSnapshotEnable := true
		expected := true

		result := filterCRD(snapshotCRDNames, crd, volumeGroupSnapshotEnable)

		if result != expected {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})

	t.Run("volumeGroupSnapshotEnable is false and CRD has only one version not equal to v1beta1", func(t *testing.T) {
		crd := &crdv1.CustomResourceDefinition{
			ObjectMeta: metav1.ObjectMeta{
				Name: "crd1",
			},
			Spec: crdv1.CustomResourceDefinitionSpec{
				Versions: []crdv1.CustomResourceDefinitionVersion{
					{
						Name: "v1",
					},
				},
			},
		}
		volumeGroupSnapshotEnable := false
		expected := true

		result := filterCRD(snapshotCRDNames, crd, volumeGroupSnapshotEnable)

		if result != expected {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})

	t.Run("volumeGroupSnapshotEnable is false and CRD has multiple versions", func(t *testing.T) {
		crd := &crdv1.CustomResourceDefinition{
			ObjectMeta: metav1.ObjectMeta{
				Name: "crd1",
			},
			Spec: crdv1.CustomResourceDefinitionSpec{
				Versions: []crdv1.CustomResourceDefinitionVersion{
					{
						Name: "v1beta1",
					},
					{
						Name: "v1",
					},
				},
			},
		}
		volumeGroupSnapshotEnable := false
		expected := false

		result := filterCRD(snapshotCRDNames, crd, volumeGroupSnapshotEnable)

		if result != expected {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})

	t.Run("CRD has shortnames", func(t *testing.T) {
		crd := &crdv1.CustomResourceDefinition{
			ObjectMeta: metav1.ObjectMeta{
				Name: "crd1",
			},
			Spec: crdv1.CustomResourceDefinitionSpec{
				Versions: []crdv1.CustomResourceDefinitionVersion{
					{
						Name: "v1",
					},
				},
				Names: crdv1.CustomResourceDefinitionNames{
					ShortNames: []string{"crd1"},
				},
			},
		}
		volumeGroupSnapshotEnable := true
		expected := true

		result := filterCRD(snapshotCRDNames, crd, volumeGroupSnapshotEnable)

		if result != expected {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})

	t.Run("CRD has no shortname", func(t *testing.T) {
		crd := &crdv1.CustomResourceDefinition{
			ObjectMeta: metav1.ObjectMeta{
				Name: "crd1",
			},
			Spec: crdv1.CustomResourceDefinitionSpec{
				Versions: []crdv1.CustomResourceDefinitionVersion{
					{
						Name: "v1",
					},
				},
			},
		}
		volumeGroupSnapshotEnable := true
		expected := false

		result := filterCRD(snapshotCRDNames, crd, volumeGroupSnapshotEnable)

		if result != expected {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})
}
