package local

import (
	"context"
	"github.com/pkg/errors"
	"io/ioutil"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// addToSecretFromFile will add data to the secret from a file, attached using the provided key.
func addToSecretFromFile(secret *corev1.Secret, key string, path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return errors.WithStack(err)
	}
	secret.Data[key] = data
	return nil
}

// addToSecretFromData will add data to the secret at the provided key.
func addToSecretFromData(secret *corev1.Secret, key string, rawData []byte) {
	secret.Data[key] = rawData
}

// prepareSecret will construct a new Secret struct with the provided metadata.
func prepareSecret(
	namespace string,
	name string,
	labels map[string]string,
	annotations map[string]string,
) *corev1.Secret {
	newSecret := corev1.Secret{}
	newSecret.Name = name
	newSecret.Namespace = namespace
	newSecret.Labels = labels
	newSecret.Annotations = annotations
	newSecret.Data = map[string][]byte{}
	return &newSecret
}

// createSecret will create the provided secret on the Kubernetes cluster.
func createSecret(newSecret *corev1.Secret) error {
	client := newKubeClient()
	_, err := client.CoreV1().Secrets(newSecret.Namespace).Create(context.Background(), newSecret, metav1.CreateOptions{})
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// getSecret will get a Kubernetes secret by name in the provided namespace.
func getSecret(namespace string, name string) (*corev1.Secret, error) {
	client := newKubeClient()
	secret, err := client.CoreV1().Secrets(namespace).Get(context.Background(), name, metav1.GetOptions{})
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return secret, nil
}

// listSecrets will list all the secrets that match the provided filters in the provided namespace.
func listSecrets(namespace string, filters metav1.ListOptions) ([]corev1.Secret, error) {
	client := newKubeClient()
	resp, err := client.CoreV1().Secrets(namespace).List(context.Background(), filters)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return resp.Items, nil
}

// deleteSecret will delete the secret in the provided namespace that has the provided name.
func deleteSecret(namespace string, secretName string) error {
	client := newKubeClient()
	err := client.CoreV1().Secrets(namespace).Delete(context.Background(), secretName, metav1.DeleteOptions{})
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}
