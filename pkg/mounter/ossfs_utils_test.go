package mounter

import (
	"testing"

	"github.com/alibabacloud-go/tea/tea"
	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
)

func Test_getPasswdSecretVolume(t *testing.T) {
	tests := []struct {
		name          string
		secretRef     string
		nodeName      string
		volumeId      string
		expectedName  string
		expectedItems []corev1.KeyToPath
	}{
		{
			name:         "TestEmptySecretRef",
			secretRef:    "",
			nodeName:     "node1",
			volumeId:     "vol1",
			expectedName: OssfsCredentialSecretName,
			expectedItems: []corev1.KeyToPath{
				{
					Key:  "node1.vol1",
					Path: "passwd-ossfs",
					Mode: tea.Int32(0600),
				},
			},
		},
		{
			name:         "TestNonEmptySecretRef",
			secretRef:    "my-secret",
			nodeName:     "node2",
			volumeId:     "vol2",
			expectedName: "my-secret",
			expectedItems: []corev1.KeyToPath{
				{
					Key:  "AccessKeyId",
					Path: "passwd-ossfs/AccessKeyId",
					Mode: tea.Int32(0600),
				},
				{
					Key:  "AccessKeySecret",
					Path: "passwd-ossfs/AccessKeySecret",
					Mode: tea.Int32(0600),
				},
				{
					Key:  "Expiration",
					Path: "passwd-ossfs/Expiration",
					Mode: tea.Int32(0600),
				},
				{
					Key:  "SecurityToken",
					Path: "passwd-ossfs/SecurityToken",
					Mode: tea.Int32(0600),
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			secret := getPasswdSecretVolume(tt.secretRef, tt.nodeName, tt.volumeId)

			assert.NotNil(t, secret)
			assert.Equal(t, tt.expectedName, secret.SecretName)
			assert.Equal(t, tt.expectedItems, secret.Items)
		})
	}
}
