package metadata

import (
	"fmt"

	eflo_controller20221215 "github.com/alibabacloud-go/eflo-controller-20221215/v3/client"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud"
	"k8s.io/klog/v2"
	"k8s.io/utils/ptr"
)

type EfloMetadata struct {
	nodeType *eflo_controller20221215.DescribeNodeTypeResponseBody
}

func getEfloNodeType(c cloud.EFLOInterface, nodeID string) (string, error) {
	req := &eflo_controller20221215.DescribeNodeRequest{
		NodeId: &nodeID,
	}
	resp, err := c.DescribeNode(req)
	if err != nil {
		return "", fmt.Errorf("DescribeNode failed: %w", err)
	}
	if resp == nil || resp.Body == nil {
		return "", fmt.Errorf("DescribeNode returned nil response, resp: %v", resp)
	}
	if resp.Body.NodeType == nil {
		klog.InfoS("DescribeNode response has nil NodeType", "resp", resp)
		return "", nil
	}
	nodeType := *resp.Body.NodeType
	klog.V(1).InfoS("EFLO DescribeNode", "nodeType", nodeType, "requestID", ptr.Deref(resp.Body.RequestId, ""))
	return nodeType, nil
}

func NewEfloMetadata(c cloud.EFLOInterface, nodeID string) (*EfloMetadata, error) {
	nodeType, err := getEfloNodeType(c, nodeID)
	if err != nil {
		return nil, err
	}
	if nodeType == "" {
		return nil, nil // Some LingJun instances don't have nodeType
	}

	req := &eflo_controller20221215.DescribeNodeTypeRequest{
		NodeType: &nodeType,
	}
	resp, err := c.DescribeNodeType(req)
	if err != nil {
		return nil, fmt.Errorf("DescribeNodeType failed: %w", err)
	}
	if resp == nil || resp.Body == nil {
		return nil, fmt.Errorf("DescribeNodeType returned invalid response: %+v", resp)
	}
	klog.V(1).InfoS("EFLO DescribeNodeType", "requestID", ptr.Deref(resp.Body.RequestId, ""))
	return &EfloMetadata{nodeType: resp.Body}, nil

}

func (m *EfloMetadata) GetAny(key MetadataKey) (any, error) {
	if m == nil {
		return "", ErrUnknownMetadataKey
	}
	switch key {
	case diskQuantity:
		if m.nodeType.DiskQuantity != nil {
			return *m.nodeType.DiskQuantity, nil
		}
	}
	return "", ErrUnknownMetadataKey
}

type EfloFetcher struct {
	efloClient cloud.EFLOInterface
	mPre       MetadataProvider
}

func (f *EfloFetcher) FetchFor(key MetadataKey) (middleware, error) {
	switch key {
	case diskQuantity:
	default:
		return nil, ErrUnknownMetadataKey
	}
	instanceId, err := f.mPre.Get(InstanceID)
	if err != nil {
		return nil, fmt.Errorf("instance ID is not available: %w", err)
	}

	p, err := NewEfloMetadata(f.efloClient, instanceId)
	if err != nil {
		return nil, err
	}
	return newImmutable(p, "EFLO"), nil
}
