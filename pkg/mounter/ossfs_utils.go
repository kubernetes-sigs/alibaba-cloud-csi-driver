package mounter

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/alibabacloud-go/tea/tea"
	log "github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
)

// keep consitent with RAM response
var secretRefKeysToParse []string = []string{
	"AccessKeyId",
	"AccessKeySecret",
	"Expiration",
	"SecurityToken",
}

func getPasswdSecretVolume(secretRef, nodeName, volumeId string) (secret *corev1.SecretVolumeSource) {
	passwdFilename := "passwd-ossfs"
	if secretRef == "" {
		secret = &corev1.SecretVolumeSource{
			SecretName: OssfsCredentialSecretName,
			Items: []corev1.KeyToPath{
				{
					Key:  fmt.Sprintf("%s.%s", nodeName, volumeId),
					Path: passwdFilename,
					Mode: tea.Int32(0600),
				},
			},
			Optional: tea.Bool(true),
		}
	} else {
		items := []corev1.KeyToPath{}
		for _, key := range secretRefKeysToParse {
			item := corev1.KeyToPath{
				Key:  key,
				Path: fmt.Sprintf("%s/%s", passwdFilename, key),
				Mode: tea.Int32(0600),
			}
			items = append(items, item)
		}
		secret = &corev1.SecretVolumeSource{
			SecretName: secretRef,
			Items:      items,
			Optional:   tea.Bool(true),
		}
	}
	return
}

func SetupOssfsCredentialSecret(ctx context.Context, clientset kubernetes.Interface, node, volumeId, bucket, akId, akSecret string) error {
	key := fmt.Sprintf("%s.%s", node, volumeId)
	value := fmt.Sprintf("%s:%s:%s", bucket, akId, akSecret)
	secretClient := clientset.CoreV1().Secrets(fuseMountNamespace)
	secret, err := secretClient.Get(ctx, OssfsCredentialSecretName, metav1.GetOptions{})
	if err != nil {
		// if secret not found, create it
		if errors.IsNotFound(err) {
			secret = new(corev1.Secret)
			secret.Name = OssfsCredentialSecretName
			secret.Namespace = fuseMountNamespace
			secret.Data = map[string][]byte{key: []byte(value)}
			_, err = secretClient.Create(ctx, secret, metav1.CreateOptions{})
			if err == nil {
				log.WithField("volumeId", volumeId).Infof("created secret %s to add credentials", OssfsCredentialSecretName)
			}
			return err
		}
		return err
	}
	if string(secret.Data[key]) == value {
		return nil
	}
	// patch secret
	patch := corev1.Secret{
		Data: map[string][]byte{
			key: []byte(value),
		},
	}
	patchData, err := json.Marshal(patch)
	if err != nil {
		return err
	}
	_, err = secretClient.Patch(ctx, OssfsCredentialSecretName, types.StrategicMergePatchType, patchData, metav1.PatchOptions{})
	if err == nil {
		log.WithField("volumeId", volumeId).Infof("patched secret %s", OssfsCredentialSecretName)
	}
	return err
}

func CleanupOssfsCredentialSecret(ctx context.Context, clientset kubernetes.Interface, node, volumeId string) error {
	key := fmt.Sprintf("%s.%s", node, volumeId)
	secretClient := clientset.CoreV1().Secrets(fuseMountNamespace)
	secret, err := secretClient.Get(ctx, OssfsCredentialSecretName, metav1.GetOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			return nil
		}
		return err
	}
	_, exists := secret.Data[key]
	if !exists {
		return nil
	}
	// patch secret
	patch := corev1.Secret{
		Data: map[string][]byte{
			key: nil,
		},
	}
	patchData, err := json.Marshal(patch)
	if err != nil {
		return err
	}
	_, err = secretClient.Patch(ctx, OssfsCredentialSecretName, types.StrategicMergePatchType, patchData, metav1.PatchOptions{})
	if err == nil {
		log.WithField("volumeId", volumeId).Infof("patched secret %s to remove credentials", OssfsCredentialSecretName)
	}
	return err
}
