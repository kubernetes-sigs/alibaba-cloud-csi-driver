package metadata

import (
	"context"
	"fmt"

	eflo_controller20221215 "github.com/alibabacloud-go/eflo-controller-20221215/v3/client"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud"
	"k8s.io/klog/v2"
	"k8s.io/utils/ptr"
)

type EfloMetadata struct {
	nodeType *eflo_controller20221215.DescribeNodeTypeResponseBody
}

func getEfloNodeType(logger klog.Logger, c cloud.EFLOInterface, nodeID string) (string, error) {
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
		logger.Info("DescribeNode response has nil NodeType", "resp", resp)
		return "", nil
	}
	nodeType := *resp.Body.NodeType
	logger.V(1).Info("EFLO DescribeNode", "nodeType", nodeType, "requestID", ptr.Deref(resp.Body.RequestId, ""))
	return nodeType, nil
}

func NewEfloMetadata(ctx context.Context, c cloud.EFLOInterface, nodeID string) (*EfloMetadata, error) {
	logger := klog.FromContext(ctx)
	nodeType, err := getEfloNodeType(logger, c, nodeID)
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
	logger.V(1).Info("EFLO DescribeNodeType", "requestID", ptr.Deref(resp.Body.RequestId, ""))
	return &EfloMetadata{nodeType: resp.Body}, nil

}

func (m *EfloMetadata) GetAny(_ *mcontext, key MetadataKey) (any, error) {
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
	mPre       middleware
}

func (f *EfloFetcher) ID() fetcherID { return efloFetcherID }

func (f *EfloFetcher) FetchFor(ctx *mcontext, key MetadataKey) (middleware, error) {
	switch key {
	case diskQuantity:
	default:
		return nil, ErrUnknownMetadataKey
	}
	instanceId, err := f.mPre.GetAny(ctx, InstanceID)
	if err != nil {
		return nil, fmt.Errorf("instance ID is not available: %w", err)
	}

	p, err := NewEfloMetadata(ctx, f.efloClient, instanceId.(string))
	if err != nil {
		return nil, err
	}
	return newImmutable(p, "EFLO"), nil
}
