package metadata

import (
	"fmt"

	eflo_controller20221215 "github.com/alibabacloud-go/eflo-controller-20221215/v3/client"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud"
	"k8s.io/utils/ptr"
)

type efloNodeMetadata struct {
	nodeType string // empty string means "no nodeType" (some LingJun instances)
}

func (m *efloNodeMetadata) GetAny(_ *mcontext, key MetadataKey) (any, error) {
	switch key {
	case LingjunNodeType:
		return m.nodeType, nil
	case machineKind:
		return MachineKindLingjun, nil
	}
	return nil, ErrUnknownMetadataKey
}

// EfloNodeFetcher calls DescribeNode to get nodeType
type EfloNodeFetcher struct {
	efloClient cloud.EFLOInterface
	mPre       middleware
}

func (f *EfloNodeFetcher) ID() fetcherID { return efloNodeFetcherID }

func (f *EfloNodeFetcher) FetchFor(ctx *mcontext, key MetadataKey) (middleware, error) {
	switch key {
	case LingjunNodeType, machineKind:
	default:
		return nil, ErrUnknownMetadataKey
	}
	kind, err := f.mPre.GetAny(ctx, machineKind)
	if err == nil && kind != MachineKindLingjun { // skip for non-LingJun instances
		ctx.logger.V(1).Info("skip EFLO metadata fetcher", "machineKind", kind)
		return empty{}, nil
	}
	instanceId, err := f.mPre.GetAny(ctx, InstanceID)
	if err != nil {
		return nil, fmt.Errorf("instance ID is not available: %w", err)
	}

	req := &eflo_controller20221215.DescribeNodeRequest{
		NodeId: new(instanceId.(string)),
	}
	resp, err := f.efloClient.DescribeNode(req)
	if err != nil {
		return nil, fmt.Errorf("DescribeNode failed: %w", err)
	}
	if resp == nil || resp.Body == nil {
		return nil, fmt.Errorf("DescribeNode returned nil response, resp: %v", resp)
	}
	nodeType := ptr.Deref(resp.Body.NodeType, "")
	ctx.logger.V(1).Info("EFLO DescribeNode", "nodeType", nodeType, "requestID", ptr.Deref(resp.Body.RequestId, ""))

	return newImmutable(&efloNodeMetadata{nodeType: nodeType}, "EFLO-Node"), nil
}

type efloNodeTypeMetadata struct {
	nodeType *eflo_controller20221215.DescribeNodeTypeResponseBody
}

func (m *efloNodeTypeMetadata) GetAny(_ *mcontext, key MetadataKey) (any, error) {
	switch key {
	case diskQuantity:
		if m.nodeType != nil && m.nodeType.DiskQuantity != nil {
			return *m.nodeType.DiskQuantity, nil
		}
	}
	return nil, ErrUnknownMetadataKey
}

// EfloNodeTypeFetcher: calls DescribeNodeType to get diskQuantity
type EfloNodeTypeFetcher struct {
	efloClient cloud.EFLOInterface
	mPre       middleware
}

func (f *EfloNodeTypeFetcher) ID() fetcherID { return efloNodeTypeFetcherID }

func (f *EfloNodeTypeFetcher) FetchFor(ctx *mcontext, key MetadataKey) (middleware, error) {
	switch key {
	case diskQuantity:
	default:
		return nil, ErrUnknownMetadataKey
	}

	nodeType, err := f.mPre.GetAny(ctx, LingjunNodeType)
	if err != nil {
		return nil, fmt.Errorf("nodeType is not available: %w", err)
	}
	nodeTypeStr := nodeType.(string)
	if nodeTypeStr == "" {
		ctx.logger.V(1).Info("skip EFLO nodeType fetcher for LingJun node without nodeType")
		return empty{}, nil
	}

	req := &eflo_controller20221215.DescribeNodeTypeRequest{
		NodeType: &nodeTypeStr,
	}
	resp, err := f.efloClient.DescribeNodeType(req)
	if err != nil {
		return nil, fmt.Errorf("DescribeNodeType failed: %w", err)
	}
	if resp == nil || resp.Body == nil {
		return nil, fmt.Errorf("DescribeNodeType returned invalid response: %+v", resp)
	}
	ctx.logger.V(1).Info("EFLO DescribeNodeType", "nodeType", nodeTypeStr, "requestID", ptr.Deref(resp.Body.RequestId, ""))
	return newImmutable(&efloNodeTypeMetadata{nodeType: resp.Body}, "EFLO-NodeType"), nil
}
