package v1beta1

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/json"
	"k8s.io/client-go/dynamic"
)

// GetCnfsObject get cnfs's object
func GetCnfsObject(client dynamic.Interface, name string) (*ContainerNetworkFileSystem, error) {
	return NewCNFSGetter(client).GetCNFS(context.TODO(), name)
}

type CNFSGetter interface {
	GetCNFS(ctx context.Context, name string) (*ContainerNetworkFileSystem, error)
}

func NewCNFSGetter(client dynamic.Interface) CNFSGetter {
	return &dynamicClientCNFSGetter{client: client}
}

type dynamicClientCNFSGetter struct {
	client dynamic.Interface
}

func (g *dynamicClientCNFSGetter) GetCNFS(ctx context.Context, name string) (*ContainerNetworkFileSystem, error) {
	utd, err := g.client.Resource(GVR).Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	data, err := utd.MarshalJSON()
	if err != nil {
		return nil, err
	}
	var cnfs ContainerNetworkFileSystem
	if err := json.Unmarshal(data, &cnfs); err != nil {
		return nil, err
	}
	return &cnfs, nil
}
