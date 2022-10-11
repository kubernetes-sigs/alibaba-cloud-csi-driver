package local

import (
	"context"
	"testing"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/local/types"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
)

func makeNode(node string, nodeLabels map[string]string) *v1.Node {
	return &v1.Node{
		ObjectMeta: metav1.ObjectMeta{
			Name:   node,
			Labels: nodeLabels,
		},
	}
}

func TestIsPmemSupported(t *testing.T) {
	tables := []struct {
		name             string
		nodeLabels       map[string]string
		expectPmemEnable bool
		expectPmemType   string
		emptyNode        bool
	}{
		{
			name: "node1",
			nodeLabels: map[string]string{
				types.PmemNodeLable: types.PmemLVMType,
			},
			expectPmemEnable: true,
			expectPmemType:   types.PmemLVMType,
			emptyNode:        false,
		},
		{
			name:             "node2",
			nodeLabels:       map[string]string{},
			expectPmemEnable: false,
			expectPmemType:   "",
			emptyNode:        false,
		},
		{
			name:             "node3",
			nodeLabels:       map[string]string{},
			expectPmemEnable: false,
			expectPmemType:   "",
			emptyNode:        true,
		},
	}

	defer func() { log.StandardLogger().ExitFunc = nil }()
	var fatal bool
	log.StandardLogger().ExitFunc = func(int) { fatal = true }

	for _, table := range tables {
		t.Run(table.name, func(t *testing.T) {
			node := &v1.Node{}
			if !table.emptyNode {
				node = makeNode(table.name, table.nodeLabels)
			}
			client := fake.NewSimpleClientset()
			client.CoreV1().Nodes().Create(context.Background(), node, metav1.CreateOptions{})
			aPmemEnable, aPmemType := IsPmemSupported(table.name, client)
			assert.Equal(t, table.expectPmemEnable, aPmemEnable)
			assert.Equal(t, table.expectPmemType, aPmemType)
			assert.Equal(t, table.emptyNode, fatal)
		})
	}
}
