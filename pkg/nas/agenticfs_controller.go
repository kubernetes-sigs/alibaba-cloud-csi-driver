//go:build !windows

package nas

import (
	"context"
	"fmt"
	"strings"
	"time"

	sdk "github.com/alibabacloud-go/nas-20170626/v4/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/metadata"
	cnfsv1beta1 "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cnfs/v1beta1"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/nas/interfaces"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/nas/internal"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/klog/v2"
	"k8s.io/utils/clock"
)

// agenticfsController orchestrates one PVC -> one AgenticSpace -> one AccessPoint.
// The outer controllerServer owns per-volume locking and PV lookup. This mode
// depends only on NAS, CNFS lookup and a clock; it does not retain RPC-local state.
//
// Creation only moves forward: replay the space ClientToken, discover/reuse its
// accesspoint, then wait for Active. Only DeleteVolume destroys resources.
// Abandoned provisioning still requires external reconciliation.
type agenticfsController struct {
	nasClient  interfaces.NasClientV2Interface
	cnfsGetter cnfsv1beta1.CNFSGetter
	clock      clock.Clock
	region     string

	apPollInterval time.Duration
	apPollTimeout  time.Duration
}

var _ internal.Controller = (*agenticfsController)(nil)

func newAgenticfsController(config *internal.ControllerConfig) (internal.Controller, error) {
	region, err := config.Metadata.Get(metadata.RegionID)
	if err != nil {
		return nil, fmt.Errorf("failed to get region ID: %w", err)
	}
	nasClient, err := config.NasClientFactory.V2(region)
	if err != nil {
		return nil, err
	}
	return &agenticfsController{
		nasClient:      nasClient,
		cnfsGetter:     config.CNFSGetter,
		clock:          clock.RealClock{},
		region:         region,
		apPollInterval: defaultApPollInterval,
		apPollTimeout:  defaultApPollTimeout,
	}, nil
}

func (c *agenticfsController) VolumeAs() string {
	return agenticFsVolumeAs
}

func (c *agenticfsController) CreateVolume(ctx context.Context, req *csi.CreateVolumeRequest) (resp *csi.CreateVolumeResponse, retErr error) {
	logger := klog.FromContext(ctx)
	args, err := c.getAgenticfsVolumeOptions(ctx, req)
	if err != nil {
		return nil, err
	}

	var agenticSpaceID, accesspointID string
	// Observe failures only after validation. Never roll back a resource that the
	// next request can recover via its ClientToken or accesspoint discovery.
	defer func() {
		if retErr != nil {
			c.reportCreateVolumeFailure(logger, args.FileSystemID, agenticSpaceID, args.FileSystemPath, accesspointID, retErr)
		}
	}()

	agenticSpaceID, err = c.createAgenticSpace(ctx, args)
	if err != nil {
		return nil, err
	}
	logger.V(2).Info(resourceCreatedLogPrefix+": agenticspace created, not delivered yet",
		"fileSystemId", args.FileSystemID,
		"agenticSpaceId", agenticSpaceID,
		"fileSystemPath", args.FileSystemPath)

	accesspointID, server, err := c.ensureAccessPoint(ctx, args, agenticSpaceID)
	if err != nil {
		return nil, err
	}
	if err := c.waitAccessPointActive(ctx, args.FileSystemID, accesspointID, &server); err != nil {
		return nil, err
	}
	if server == "" {
		return nil, status.Errorf(codes.Internal,
			"nas:DescribeAccesspoint: accesspoint %s is Active but the OpenAPI returned no domain to mount", accesspointID)
	}
	return &csi.CreateVolumeResponse{Volume: args.volume(agenticSpaceID, accesspointID, server)}, nil
}

func (c *agenticfsController) DeleteVolume(ctx context.Context, req *csi.DeleteVolumeRequest, pv *corev1.PersistentVolume) (*csi.DeleteVolumeResponse, error) {
	attributes := pv.Spec.CSI.VolumeAttributes
	filesystemID := agenticfsFilesystemID(attributes)
	if filesystemID == "" {
		return nil, status.Errorf(codes.InvalidArgument, "missing %s in volume attributes", filesystemIDKey)
	}
	agenticSpaceID := attributes[vcKeyAgenticSpaceId]
	if agenticSpaceID == "" {
		return nil, status.Errorf(codes.InvalidArgument,
			"missing agenticSpaceId in volume attributes (read fileSystemId=%q, accesspointId=%q, volumeId=%q): refusing to report success because the AgenticSpace would leak; follow the manual cleanup procedure in examples/nas/agenticfs/README.md §4.1",
			filesystemID, attributes[vcKeyAccesspointId], req.VolumeId)
	}

	accesspointIDs, spaceGone, err := c.accessPointsForDeletion(ctx, filesystemID, agenticSpaceID, attributes[vcKeyAccesspointId])
	if err != nil {
		return nil, err
	}
	if spaceGone {
		return &csi.DeleteVolumeResponse{}, nil
	}
	// Include recovery residue, not only the AP saved in the PV. All accesspoints
	// must be confirmed absent before DeleteAgenticSpace can proceed.
	for _, accesspointID := range accesspointIDs {
		if err := c.deleteAccessPoint(ctx, filesystemID, accesspointID); err != nil {
			return nil, err
		}
	}
	if err := c.deleteAgenticSpace(ctx, filesystemID, agenticSpaceID, req.VolumeId); err != nil {
		return nil, err
	}
	return &csi.DeleteVolumeResponse{}, nil
}

func (c *agenticfsController) ControllerExpandVolume(ctx context.Context, req *csi.ControllerExpandVolumeRequest, pv *corev1.PersistentVolume) (*csi.ControllerExpandVolumeResponse, error) {
	attributes := pv.Spec.CSI.VolumeAttributes
	filesystemID := agenticfsFilesystemID(attributes)
	if filesystemID == "" {
		return nil, status.Errorf(codes.InvalidArgument, "missing %s in volume attributes", filesystemIDKey)
	}
	agenticSpaceID := attributes[vcKeyAgenticSpaceId]
	if agenticSpaceID == "" {
		return nil, status.Error(codes.InvalidArgument, "missing agenticSpaceId in volume attributes")
	}
	sizeLimit, err := computeExpandSizeLimit(req.GetCapacityRange())
	if err != nil {
		return nil, err
	}

	space, err := c.nasClient.GetAgenticSpace(ctx, &sdk.GetAgenticSpaceRequest{
		FileSystemId:   tea.String(filesystemID),
		AgenticSpaceId: tea.String(agenticSpaceID),
	})
	if err != nil {
		return nil, apiStatusError("nas:GetAgenticSpace", err)
	}
	quota, err := prepareAgenticSpaceExpansion(klog.FromContext(ctx), filesystemID, agenticSpaceID, sizeLimit, space)
	if err != nil {
		return nil, err
	}
	if _, err := c.nasClient.SetAgenticSpaceQuota(ctx, quota); err != nil {
		return nil, apiStatusError("nas:SetAgenticSpaceQuota", err)
	}
	return &csi.ControllerExpandVolumeResponse{CapacityBytes: sizeLimit}, nil
}

// The token and path are derived from the same validated name on every attempt.
// A failed/malformed response is not proof that NAS did not create the space.
func (c *agenticfsController) createAgenticSpace(ctx context.Context, args *agenticfsVolumeArgs) (string, error) {
	resp, err := c.nasClient.CreateAgenticSpace(ctx, &sdk.CreateAgenticSpaceRequest{
		FileSystemId:   tea.String(args.FileSystemID),
		FileSystemPath: tea.String(args.FileSystemPath),
		Azone:          tea.String(args.ZoneID),
		ClientToken:    tea.String(args.Name),
		Quota: &sdk.CreateAgenticSpaceRequestQuota{
			SizeLimit:      tea.Int64(args.SizeLimit),
			FileCountLimit: tea.Int64(args.FileCountLimit),
		},
	})
	if err != nil {
		// WORKAROUND: NAS CreateAgenticSpace does not honor ClientToken for
		// idempotency. When the same ClientToken+Path is retried (e.g. after a
		// gRPC timeout where the server accepted but the client lost the
		// response), the API returns InvalidArgument with message "Path already
		// used" instead of returning the existing AgenticSpaceId.
		//
		// We fall back to discovering the space via ListAccesspoints. This adds
		// an extra API call on the error-recovery path only (not on first
		// creation). Remove this workaround once NAS supports proper ClientToken
		// idempotency for CreateAgenticSpace.
		if strings.Contains(err.Error(), "Path already used") {
			klog.InfoS("CreateAgenticSpace: path already exists, discovering existing space",
				"fileSystemId", args.FileSystemID, "path", args.FileSystemPath)
			return c.discoverAgenticSpaceByPath(ctx, args.FileSystemID, args.FileSystemPath)
		}
		return "", apiStatusError("nas:CreateAgenticSpace", err)
	}
	if resp == nil || resp.Body == nil {
		return "", status.Error(codes.Internal, "nas:CreateAgenticSpace: malformed response (nil body), AgenticSpaceId unavailable")
	}
	id := tea.StringValue(resp.Body.AgenticSpaceId)
	if id == "" {
		return "", status.Error(codes.Internal, "nas:CreateAgenticSpace: empty AgenticSpaceId in response")
	}
	return id, nil
}

func (c *agenticfsController) discoverAgenticSpaceByPath(ctx context.Context, filesystemID, path string) (string, error) {
	listResp, err := c.nasClient.ListAccesspoints(ctx, &sdk.ListAccessPointsRequest{
		FileSystemId: tea.String(filesystemID),
		MaxResults:   tea.Int32(100),
	})
	if err != nil {
		return "", fmt.Errorf("nas:ListAccesspoints failed while discovering existing AgenticSpace: %w", err)
	}
	for _, ap := range listResp.Body.AccessPoints {
		if tea.StringValue(ap.RootPath) == path || tea.StringValue(ap.RootPath)+"/" == path || path+"/" == tea.StringValue(ap.RootPath) {
			spaceID := tea.StringValue(ap.AgenticSpaceId)
			if spaceID != "" {
				klog.InfoS("Discovered existing AgenticSpace via AccessPoint",
					"agenticSpaceId", spaceID, "accessPointId", tea.StringValue(ap.AccessPointId), "path", path)
				return spaceID, nil
			}
		}
	}
	return "", fmt.Errorf("nas:CreateAgenticSpace reported Path already used for %s but no matching AgenticSpace found via ListAccesspoints", path)
}

func (c *agenticfsController) deleteAgenticSpace(ctx context.Context, filesystemID, agenticSpaceID, volumeID string) error {
	_, err := c.nasClient.DeleteAgenticSpace(ctx, &sdk.DeleteAgenticSpaceRequest{
		FileSystemId:   tea.String(filesystemID),
		AgenticSpaceId: tea.String(agenticSpaceID),
		ClientToken:    tea.String(volumeID),
	})
	if err == nil {
		return nil
	}
	if !isAgenticSpaceNotFoundError(err) {
		if isNotFoundError(err) {
			return status.Errorf(codes.Aborted,
				"nas:DeleteAgenticSpace: %v; an absent accesspoint does not confirm agenticspace %s was deleted, retrying", err, agenticSpaceID)
		}
		return apiStatusError("nas:DeleteAgenticSpace", err)
	}
	klog.FromContext(ctx).Info("agenticspace already deleted", "agenticSpaceId", agenticSpaceID)
	return nil
}
