//go:build !windows

package nas

import (
	"context"
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
	"sync/atomic"
	"time"
	"unicode"

	sdk "github.com/alibabacloud-go/nas-20170626/v4/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/metadata"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/wrap"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/nas/cloud"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/nas/interfaces"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/nas/internal"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/klog/v2"

	mounterutils "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
)

// AgenticFS is a serverless NAS offering where one filesystem manages up to 500k isolated
// AgenticSpaces. Each PVC provisions one AgenticSpace plus one AccessPoint bound to it. This
// controller implements volumeAs="Agentic" and touches the control plane only.
//
// CreateAgenticSpace is idempotent by ClientToken=<PV name>; CreateAccessPoint is not, so every
// attempt lists first and reuses one - otherwise each retry orphans an accesspoint, and
// DeleteAgenticSpace requires them all detached. A failure never deletes the space: the compensation
// deletes only the accesspoint this call created, only on a terminal gRPC code, and reports the
// leftover through the log prefixes below, a contract with an external reaper.
const (
	// Use the same canonical value for CSI volumeAs, CNFS spec.type and NAS StorageType.
	agenticFsVolumeAs = cloud.StorageTypeAgentic

	// cnfsSpecTypeAgenticFS is the cross-repo contract value for CNFS spec.type.
	cnfsSpecTypeAgenticFS = cloud.StorageTypeAgentic

	// StorageClass parameter keys.
	paramContainerNetworkFileSystem = "containerNetworkFileSystem"
	paramAgenticSpaceSizeLimit      = "agenticSpaceSizeLimit"
	paramAgenticSpaceFileCountLimit = "agenticSpaceFileCountLimit"

	// The AP domain rides in "server", the fstype in "mountProtocol", and the filesystem ID must use the
	// package-wide filesystemIDKey spelling getNASIDFromMapOrServer looks up.
	vcKeyServer         = "server"
	vcKeyPath           = "path"
	vcKeyMountProtocol  = "mountProtocol"
	vcKeyOptions        = "options"
	vcKeyAccesspointId  = "accesspointId"
	vcKeyAgenticSpaceId = "agenticSpaceId"

	// An earlier, never released build wrote this spelling; accepted on read only, for rolling upgrades.
	vcKeyFilesystemIdLegacy = "filesystemId"

	// Written explicitly: controllerserver.go does not forward it, and a "server" key does not trigger doMount's alinas branch.
	mountProtocolAlinas = "alinas"

	// Controller-side fallbacks: preserve every caller-supplied key and append
	// these options only when absent. The node validates TLS/RAM without fixing them.
	defaultAgenticFsMountOptions = "tls,vers=3,ram"

	accessPointStatusActive   = "Active"
	accessPointStatusDeleting = "Deleting"
	// Provisioned but not yet mountable, so it must not be treated as reusable.
	accessPointStatusInactive = "Inactive"

	// Observability only: a tag match does not prove which AgenticSpace an accesspoint is bound to.
	apTagKeyVolumeID = "csi.alibabacloud.com/volume-id"

	apListFilterAgenticSpaceId = "AgenticSpaceId"

	apListMaxResults int32 = 100 // SDK maximum
	// Hitting the cap means the AgenticSpaceId filter was not applied server-side; Aborted so the caller retries.
	apListMaxPages = 50

	defaultAgenticSpaceFileCountLimit int64 = 1000000

	// AgenticSpace quota limits enforced by the OpenAPI, mirrored from the SDK request models.
	minAgenticSpaceSizeLimit      int64 = 10 * GiB // 10737418240 (10 GiB)
	maxAgenticSpaceSizeLimit      int64 = 1099511627776000
	minAgenticSpaceFileCountLimit int64 = 10000
	maxAgenticSpaceFileCountLimit int64 = 100000000

	maxVolumeNameLen = 64

	defaultApPollInterval = 3 * time.Second
	// Deliberately smaller than external-provisioner's default --timeout=60s so the driver decides the timeout: it
	// classifies the failure as retryable DeadlineExceeded, letting the next attempt reuse the space and accesspoint.
	defaultApPollTimeout = 45 * time.Second

	// Bounds the compensating DeleteAccesspoint, which compensateCreateVolume enforces itself because ctx
	// never reaches the SDK. Must stay strictly greater than nasAPICallBound, or the outcome is a coin flip.
	compensationTimeout = 15 * time.Second

	// driverRPCBudget is the driver's model of how long a CreateVolume RPC may run before the
	// sidecar stops listening; it gates the compensating delete so the delete is only attempted
	// while its verdict can still be observed. It mirrors external-provisioner's 60s upstream
	// default, not this chart's --timeout=150s: a budget that holds under 60s also holds under 150s.
	driverRPCBudget = 60 * time.Second

	// Mirrors connTimeout in pkg/nas/cloud; duplicated as a literal and pinned by TestAgenticfsConstants.
	nasAPICallBound = 10 * time.Second

	// Marks "a billable resource now exists but has not been delivered yet": the reconciliation key for
	// a crash window no error path can report. Not a cleanup signal - most of these end in a delivered volume.
	resourceCreatedLogPrefix = "agenticfs-resource-created"

	// Marks "kept on purpose for the next attempt"; carries the same contract fields as the orphan line.
	retainedForRetryLogPrefix = "agenticfs-resource-retained-for-retry"

	// orphanLogPrefix is the confirmed-leak line and the primary hook for an external reaper. It
	// always carries eight contract fields: fileSystemId, agenticSpaceId, fileSystemPath,
	// accesspointId, region, volumeHandle, reason and cause. fileSystemPath finds a space whose ID
	// was never read back; region is required because fileSystemId is only region-scoped.
	orphanLogPrefix = "agenticfs-orphan-resource"

	// Not a substring of orphanLogPrefix on purpose, so a reaper counts each leak exactly once.
	orphanResolvedLogPrefix = "agenticfs-orphan-resolved"

	// Mirrors the wrapper pkg/nas/cloud's wait() puts on every rate-limiter failure: the only signal
	// distinguishing "the delete was never issued" from "it was issued and failed".
	nasRateLimiterWaitPrefix = "error while waiting for rate limiter"
)

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
		config:         config,
		nasClient:      nasClient,
		region:         region,
		apPollInterval: defaultApPollInterval,
		apPollTimeout:  defaultApPollTimeout,
		compTimeout:    compensationTimeout,
		rpcBudget:      driverRPCBudget,
	}, nil
}

type agenticfsController struct {
	nasClient interfaces.NasClientV2Interface
	config    *internal.ControllerConfig

	// On the struct so every orphan line carries it: fileSystemId is only region-scoped.
	region string

	// AP creation is asynchronous; these tune the poll loop and are overridable in tests.
	apPollInterval time.Duration
	apPollTimeout  time.Duration

	// A field only so tests can shrink it; production always uses compensationTimeout.
	compTimeout time.Duration

	// The delete runs only while "elapsed + compTimeout <= rpcBudget". A field for the same reason.
	rpcBudget time.Duration
}

func (c *agenticfsController) VolumeAs() string {
	return agenticFsVolumeAs
}

func (c *agenticfsController) CreateVolume(ctx context.Context, req *csi.CreateVolumeRequest) (resp *csi.CreateVolumeResponse, retErr error) {
	logger := klog.FromContext(ctx)

	// Stamped at handler entry: inside compensateCreateVolume it would read elapsed as 0 and open the
	// budget gate on exactly the path it exists to close, and ctx.Deadline() fails with no deadline set.
	rpcStarted := time.Now()

	// 0. req.Name ends up in FileSystemPath, ClientToken and AccessPointName, so validate it first.
	if err := validateVolumeName(req.Name); err != nil {
		return nil, err
	}
	parameters := req.Parameters

	// 1. fileSystemId wins when both sources are set. The direct path does not pre-check the type:
	// CreateAgenticSpace rejects a non-AgenticFS filesystem precisely, without an extra round trip.
	filesystemId := parameters[filesystemIDKey]
	if filesystemId == "" {
		var err error
		if filesystemId, err = c.filesystemIDFromCNFS(ctx, parameters[paramContainerNetworkFileSystem]); err != nil {
			return nil, err
		}
	}

	// Placement comes from the StorageClass, so one filesystem can back StorageClasses in different zones.
	zoneId := parameters[ZoneID]
	if zoneId == "" {
		return nil, status.Errorf(codes.InvalidArgument,
			"storageclass parameters.%s is empty; it is required as CreateAgenticSpace.Azone", ZoneID)
	}
	vpcId := parameters[VpcID]
	if vpcId == "" {
		return nil, status.Errorf(codes.InvalidArgument,
			"storageclass parameters.%s is empty; it is required as CreateAccessPoint.VpcId", VpcID)
	}
	vswId := parameters[VSwitchID]
	if vswId == "" {
		return nil, status.Errorf(codes.InvalidArgument,
			"storageclass parameters.%s is empty; it is required as CreateAccessPoint.VswId", VSwitchID)
	}

	// 2. Before any cloud call, so an invalid StorageClass never creates a billable resource.
	sizeLimit, err := computeAgenticSpaceSizeLimit(req.GetCapacityRange(), parameters[paramAgenticSpaceSizeLimit])
	if err != nil {
		return nil, err
	}
	fileCountLimit, err := computeAgenticSpaceFileCountLimit(parameters[paramAgenticSpaceFileCountLimit])
	if err != nil {
		return nil, err
	}
	if err := validateAgenticSpaceQuota(sizeLimit, fileCountLimit); err != nil {
		return nil, err
	}

	agenticSpaceId := ""
	createdAccesspointId := ""

	// From here on a failure can leave billable resources with no DeleteVolume ever called for them. The
	// compensation deletes only the accesspoint this call created, and only on a terminal gRPC code.
	defer func() {
		if retErr == nil {
			return
		}
		c.compensateCreateVolume(ctx, logger, filesystemId, agenticSpaceId, "/"+req.Name, createdAccesspointId, retErr, rpcStarted)
	}()

	// 3. ClientToken = req.Name makes this idempotent: a retry replays it and gets the same AgenticSpaceId.
	// The request has no RegionId field; region comes from the client's GlobalParameters.
	spaceResp, err := c.nasClient.CreateAgenticSpace(ctx, &sdk.CreateAgenticSpaceRequest{
		FileSystemId: tea.String(filesystemId),
		// The vendor model documents a trailing slash; whether the OpenAPI requires it cannot be settled offline.
		FileSystemPath: tea.String("/" + req.Name), // first-level directory; PV name is pvc-<uuid>
		Azone:          tea.String(zoneId),
		ClientToken:    tea.String(req.Name),
		Quota: &sdk.CreateAgenticSpaceRequestQuota{
			SizeLimit:      tea.Int64(sizeLimit),
			FileCountLimit: tea.Int64(fileCountLimit),
		},
	})
	if err != nil {
		// An ambiguous failure (reset, 5xx) may have created a billable space whose ID was never read back.
		// The defer reports it keyed by fileSystemPath, the only key still available here.
		return nil, apiStatusError("nas:CreateAgenticSpace", err)
	}
	// A malformed response and an empty AgenticSpaceId are the same failure mode: a billable space may
	// exist but its ID was never read back. Retryable Internal, because the replay is the only recovery.
	malformedSpaceResp := spaceResp == nil || spaceResp.Body == nil
	if !malformedSpaceResp {
		agenticSpaceId = tea.StringValue(spaceResp.Body.AgenticSpaceId)
	}
	if agenticSpaceId == "" {
		msg := "nas:CreateAgenticSpace: empty AgenticSpaceId in response"
		if malformedSpaceResp {
			msg = "nas:CreateAgenticSpace: malformed response (nil body), AgenticSpaceId unavailable"
		}
		return nil, status.Error(codes.Internal, msg)
	}

	// Emitted pre-emptively on the success path: a crash between here and the final response runs no defer
	// and returns no error. V(2) because it fires on every successful provision.
	logger.V(2).Info(resourceCreatedLogPrefix+": agenticspace created, not delivered yet",
		"fileSystemId", filesystemId,
		"agenticSpaceId", agenticSpaceId,
		"fileSystemPath", "/"+req.Name)

	// 4. CreateAccessPoint has no ClientToken, so this lookup is the only thing that keeps a retried
	// CreateVolume from stacking duplicate accesspoints on the same space.
	accesspointId, server, err := c.findReusableAccessPoint(ctx, filesystemId, agenticSpaceId)
	if err != nil {
		return nil, err
	}

	// 5. EnabledRam must be true; AgenticFS does not support AccessGroup/RootDirectory/Owner*/Permission/Posix*.
	if accesspointId == "" {
		apResp, err := c.nasClient.CreateAccesspoint(ctx, &sdk.CreateAccessPointRequest{
			FileSystemId:    tea.String(filesystemId),
			AgenticSpaceId:  tea.String(agenticSpaceId),
			VswId:           tea.String(vswId),
			VpcId:           tea.String(vpcId),
			AccessPointName: tea.String(req.Name),
			EnabledRam:      tea.Bool(true),
			Tag: []*sdk.CreateAccessPointRequestTag{{
				Key:   tea.String(apTagKeyVolumeID),
				Value: tea.String(req.Name),
			}},
		})
		if err != nil {
			if isNotFoundError(err) {
				// Almost always a stale AgenticSpace id, e.g. a replay whose space an earlier attempt compensated
				// away. Terminal on purpose: retrying cannot fix it and recreating the PVC yields a fresh ClientToken.
				return nil, status.Errorf(codes.InvalidArgument,
					"nas:CreateAccesspoint: agenticspace %s was rejected as missing (%v); it may have been deleted by an earlier failed attempt replaying the same ClientToken - delete and recreate the PVC",
					agenticSpaceId, err)
			}
			return nil, apiStatusError("nas:CreateAccesspoint", err)
		}
		if apResp == nil || apResp.Body == nil || apResp.Body.AccessPoint == nil {
			return nil, status.Error(codes.Internal, "nas:CreateAccesspoint: empty AccessPoint in response")
		}
		// tea.StringValue turns nil into ""; an empty ID would be reported as a volume nobody can ever delete.
		accesspointId = tea.StringValue(apResp.Body.AccessPoint.AccessPointId)
		if accesspointId == "" {
			return nil, status.Error(codes.Internal, "nas:CreateAccesspoint: empty AccessPointId in response")
		}
		// The single ID the compensation is allowed to delete; the reuse branch leaves it empty.
		createdAccesspointId = accesspointId
		server = tea.StringValue(apResp.Body.AccessPoint.AccessPointDomain)
	} else {
		logger.Info("reusing the accesspoint already bound to the agenticspace",
			"agenticSpaceId", agenticSpaceId, "accesspointId", accesspointId)
	}

	// 6. Creation is asynchronous; wait for the AccessPoint to become Active.
	if err := c.waitAccessPointActive(ctx, filesystemId, accesspointId, &server); err != nil {
		return nil, err
	}
	// An Active accesspoint without a domain cannot be mounted and the node side cannot recover it.
	if server == "" {
		return nil, status.Errorf(codes.Internal,
			"nas:DescribeAccesspoint: accesspoint %s is Active but the OpenAPI returned no domain to mount", accesspointId)
	}

	// 7. The generic wrapper overwrites "options" and always adds "volumeAs"; enforceAgenticFsMountOptions
	// runs after it. Do not write accesspoint / authType / sandboxId: the phase-1 node side cannot use them.
	volumeContext := map[string]string{
		vcKeyServer:         server,
		vcKeyPath:           "/",
		vcKeyMountProtocol:  mountProtocolAlinas,
		vcKeyOptions:        mergeAgenticFsMountOptions(parameters[vcKeyOptions], defaultAgenticFsMountOptions),
		filesystemIDKey:     filesystemId,
		vcKeyAccesspointId:  accesspointId,
		vcKeyAgenticSpaceId: agenticSpaceId,
	}
	return &csi.CreateVolumeResponse{
		Volume: &csi.Volume{
			VolumeId: req.Name,
			// A replay does not re-check the quota against CapacityBytes, but req.Name is pvc-<uuid>: every replay
			// comes from the same PVC, and a recreated PVC gets a fresh space rather than a replay onto an old one.
			CapacityBytes: sizeLimit,
			VolumeContext: volumeContext,
		},
	}, nil
}

// Resolves the filesystem a CNFS created, for a StorageClass that does not name one with fileSystemId.
func (c *agenticfsController) filesystemIDFromCNFS(ctx context.Context, cnfsName string) (string, error) {
	if cnfsName == "" {
		return "", status.Errorf(codes.InvalidArgument,
			"storageclass parameters.%s and parameters.%s are both empty; one of them must name the AgenticFS filesystem",
			filesystemIDKey, paramContainerNetworkFileSystem)
	}
	cnfs, err := c.config.CNFSGetter.GetCNFS(ctx, cnfsName)
	if err != nil {
		if apierrors.IsNotFound(err) {
			return "", status.Errorf(codes.InvalidArgument, "CNFS not found: %s", cnfsName)
		}
		return "", status.Errorf(codes.Internal, "failed to get CNFS %s: %v", cnfsName, err)
	}

	// AgenticFS is distinguished by StorageType only; its FilesystemType is "standard" and must NOT be used.
	if cnfs.Spec.StorageType != cnfsSpecTypeAgenticFS {
		return "", status.Errorf(codes.InvalidArgument,
			"CNFS %s spec.type is %q, expected %q", cnfsName, cnfs.Spec.StorageType, cnfsSpecTypeAgenticFS)
	}
	if cnfs.Status.FsAttributes.StorageType != cloud.StorageTypeAgentic {
		return "", status.Errorf(codes.InvalidArgument,
			"CNFS %s status.fsAttributes.storageType is %q, expected %q (AgenticFS not ready or mismatched)",
			cnfsName, cnfs.Status.FsAttributes.StorageType, cloud.StorageTypeAgentic)
	}

	// Server is intentionally NOT read: mounting goes through the AP domain, not the filesystem server.
	filesystemId := cnfs.Status.FsAttributes.FilesystemID
	if filesystemId == "" {
		return "", status.Errorf(codes.InvalidArgument, "CNFS %s status.fsAttributes.filesystemId is empty", cnfsName)
	}
	return filesystemId, nil
}

// Completes missing options after controllerserver.go has overwritten "options".
// Caller-supplied values always win. Other volume modes are unchanged.
func enforceAgenticFsMountOptions(volumeAs string, volumeContext map[string]string) {
	if volumeAs != agenticFsVolumeAs || volumeContext == nil {
		return
	}
	current := volumeContext[vcKeyOptions]
	merged := mergeAgenticFsMountOptions(current, defaultAgenticFsMountOptions)
	if merged == current {
		// Left as it was, not even re-joined: re-joining could reorder it and change the default path.
		return
	}
	volumeContext[vcKeyOptions] = merged
}

// Uses the same parser and key extraction as the node side, so "present" means the same on both sides.
func mergeAgenticFsMountOptions(options, forced string) string {
	base := mounterutils.SplitMountOptions(options)
	// "none" suppresses optional defaults, not the AgenticFS required options.
	if len(base) == 1 && strings.EqualFold(base[0], "none") {
		base = nil
	}
	merged := mounterutils.MergeMountOptions(base, mounterutils.SplitMountOptions(forced))
	return strings.Join(merged, ",")
}

// Returns an accesspoint already bound to agenticSpaceId, or empty strings. The filter is evaluated
// server-side, so a hit is authoritative - a tag lookup cannot prove it.
func (c *agenticfsController) findReusableAccessPoint(ctx context.Context, filesystemId, agenticSpaceId string) (accesspointId, domain string, err error) {
	logger := klog.FromContext(ctx)
	accesspoints, err := c.listAccessPointsOfSpace(ctx, filesystemId, agenticSpaceId)
	if err != nil {
		return "", "", apiStatusError("nas:ListAccesspoints", err)
	}

	var usable []*sdk.ListAccessPointsResponseBodyAccessPoints
	var unavailable []*sdk.ListAccessPointsResponseBodyAccessPoints
	for _, ap := range accesspoints {
		if tea.StringValue(ap.AccessPointId) == "" {
			continue
		}
		// Neither can be mounted: an Inactive one would poll until the timeout, a Deleting one disappears.
		switch tea.StringValue(ap.Status) {
		case accessPointStatusDeleting, accessPointStatusInactive:
			unavailable = append(unavailable, ap)
			continue
		}
		usable = append(usable, ap)
	}
	if len(usable) == 0 {
		if len(unavailable) > 0 {
			// None can be mounted and creating a second one would stack them. Aborted converges without destroying
			// anything: a Deleting accesspoint disappears, an Inactive one eventually becomes Active.
			described := make([]string, 0, len(unavailable))
			ids := make([]string, 0, len(unavailable))
			for _, ap := range unavailable {
				id := tea.StringValue(ap.AccessPointId)
				described = append(described, fmt.Sprintf("%s(%s)", id, tea.StringValue(ap.Status)))
				ids = append(ids, id)
			}
			logger.Info("WARNING: none of the accesspoints of the agenticspace is usable yet, retrying",
				"agenticSpaceId", agenticSpaceId, "accesspoints", strings.Join(described, ", "))
			// An Inactive-only space can retry forever, so the message must tell an operator what to do.
			return "", "", status.Errorf(codes.Aborted,
				"nas:ListAccesspoints: the accesspoint(s) of agenticspace %s are not usable yet: %s; "+
					"activate or delete accesspoint %s in the NAS console, or delete the PVC to release the agenticspace",
				agenticSpaceId, strings.Join(described, ", "), strings.Join(ids, ", "))
		}
		return "", "", nil
	}

	picked := usable[0]
	if tea.StringValue(picked.Status) != accessPointStatusActive {
		for _, ap := range usable {
			if tea.StringValue(ap.Status) == accessPointStatusActive {
				picked = ap
				break
			}
		}
	}
	if len(usable) > 1 {
		// Should not happen, but an interrupted older build can leave several; DeleteVolume removes the set.
		logger.Info("more than one accesspoint is bound to the agenticspace, reusing one of them",
			"agenticSpaceId", agenticSpaceId, "count", len(usable), "reused", tea.StringValue(picked.AccessPointId))
	}
	return tea.StringValue(picked.AccessPointId), tea.StringValue(picked.DomainName), nil
}

// Follows NextToken to the end: a missed accesspoint means a duplicate gets created, or a leftover
// blocks DeleteAgenticSpace. Errors are unwrapped so callers can tell NotFound from a failure.
func (c *agenticfsController) listAccessPointsOfSpace(ctx context.Context, filesystemId, agenticSpaceId string) ([]*sdk.ListAccessPointsResponseBodyAccessPoints, error) {
	var accesspoints []*sdk.ListAccessPointsResponseBodyAccessPoints
	nextToken := ""
	for page := 1; ; page++ {
		if page > apListMaxPages {
			return nil, status.Errorf(codes.Aborted,
				"nas:ListAccesspoints: gave up after %d pages of %d while enumerating the accesspoints of agenticspace %s on filesystem %s; retry later",
				apListMaxPages, apListMaxResults, agenticSpaceId, filesystemId)
		}
		listReq := &sdk.ListAccessPointsRequest{
			FileSystemId: tea.String(filesystemId),
			Filters: []*sdk.ListAccessPointsRequestFilters{{
				Name:  tea.String(apListFilterAgenticSpaceId),
				Value: tea.String(agenticSpaceId),
			}},
			MaxResults: tea.Int32(apListMaxResults),
		}
		if nextToken != "" {
			listReq.NextToken = tea.String(nextToken)
		}
		resp, err := c.nasClient.ListAccesspoints(ctx, listReq)
		if err != nil {
			return nil, err
		}
		if resp == nil || resp.Body == nil {
			// A nil response or body is malformed, never "genuinely none": acting on it would stack a duplicate.
			return nil, status.Errorf(codes.Aborted,
				"nas:ListAccesspoints: page %d of the agenticspace %s enumeration returned a malformed response (nil body); refusing to act on a partial list",
				page, agenticSpaceId)
		}
		// Re-verify the server-side filter: if the server ignored Filters it would hand back other volumes'
		// accesspoints, and mounting one would point VolumeContext[server] at someone else's accesspoint.
		for _, ap := range resp.Body.AccessPoints {
			got := tea.StringValue(ap.AgenticSpaceId)
			if got == agenticSpaceId {
				continue
			}
			// AgenticSpaceId is optional on the wire, so an API version that omits it wedges both cases on Aborted.

			if got == "" {
				return nil, status.Errorf(codes.Aborted,
					"nas:ListAccesspoints: accesspoint %s came back without an AgenticSpaceId, so it cannot be proven to belong to agenticspace %q; refusing to act (this usually means the API version does not return the field)",
					tea.StringValue(ap.AccessPointId), agenticSpaceId)
			}
			return nil, status.Errorf(codes.Aborted,
				"nas:ListAccesspoints: the server-side AgenticSpaceId filter did not take effect: accesspoint %s belongs to agenticspace %q, not the requested %q; refusing to act on an accesspoint of another volume",
				tea.StringValue(ap.AccessPointId), got, agenticSpaceId)
		}
		accesspoints = append(accesspoints, resp.Body.AccessPoints...)
		nextToken = tea.StringValue(resp.Body.NextToken)
		if nextToken == "" {
			return accesspoints, nil
		}
	}
}

// The returned error is already a gRPC status: retryable for a timeout, permanent for a rejected request.
func (c *agenticfsController) waitAccessPointActive(ctx context.Context, filesystemId, accesspointId string, server *string) error {
	return c.pollAccessPoint(ctx, filesystemId, accesspointId, apPoll{
		waiting: "to become Active",
		observe: func(ap *sdk.DescribeAccessPointResponseBodyAccessPoint) {
			if server != nil && *server == "" {
				*server = tea.StringValue(ap.DomainName)
			}
		},
		done: func(ap *sdk.DescribeAccessPointResponseBodyAccessPoint, notFound bool) bool {
			// NotFound right after CreateAccesspoint is the read-after-write window, not a permanent failure.
			return !notFound && ap != nil && tea.StringValue(ap.Status) == accessPointStatusActive
		},
		timeout: func(lastStatus string) error {
			return status.Errorf(codes.DeadlineExceeded,
				"timed out waiting for accesspoint %s to become Active after %s (last status: %s)",
				accesspointId, c.apPollTimeout, lastStatus)
		},
	})
}

// DeleteAccesspoint returning only means the API accepted it, and DeleteAgenticSpace rejects a space with
// accesspoints attached.
func (c *agenticfsController) waitAccessPointGone(ctx context.Context, filesystemId, accesspointId string) error {
	return c.pollAccessPoint(ctx, filesystemId, accesspointId, apPoll{
		waiting: "to be deleted",
		done: func(_ *sdk.DescribeAccessPointResponseBodyAccessPoint, notFound bool) bool {
			return notFound
		},
		timeout: func(lastStatus string) error {
			// Aborted, not Internal: the accesspoint is converging towards gone, so the caller should retry.
			return status.Errorf(codes.Aborted,
				"accesspoint %s is still being deleted after %s (last status: %s); DeleteAgenticSpace requires all accesspoints of the space to be detached, retrying",
				accesspointId, c.apPollTimeout, lastStatus)
		},
	})
}

type apPoll struct {
	waiting string
	observe func(ap *sdk.DescribeAccessPointResponseBodyAccessPoint)
	// ap is nil when the API returned an error; notFound tells that apart from other failures.
	done    func(ap *sdk.DescribeAccessPointResponseBodyAccessPoint, notFound bool) bool
	timeout func(lastStatus string) error
}

// Checks cancellation BEFORE spending an API call and reuses one timer (time.After leaks one per round).
func (c *agenticfsController) pollAccessPoint(ctx context.Context, filesystemId, accesspointId string, p apPoll) error {
	logger := klog.FromContext(ctx)
	deadline := time.Now().Add(c.apPollTimeout)
	timer := time.NewTimer(c.apPollInterval)
	defer timer.Stop()
	lastStatus := "unknown"
	for {
		select {
		case <-ctx.Done():
			return status.Errorf(codes.DeadlineExceeded, "waiting for accesspoint %s %s: %v", accesspointId, p.waiting, ctx.Err())
		default:
		}

		var ap *sdk.DescribeAccessPointResponseBodyAccessPoint
		resp, err := c.nasClient.DescribeAccesspoint(ctx, filesystemId, accesspointId)
		notFound := isNotFoundError(err)
		switch {
		case err == nil:
			if resp != nil && resp.Body != nil {
				ap = resp.Body.AccessPoint
			}
			if ap != nil {
				lastStatus = tea.StringValue(ap.Status)
				if p.observe != nil {
					p.observe(ap)
				}
			}
		case notFound:
			lastStatus = "NotFound"
		case isPermanentAPIError(err):
			return status.Errorf(codes.InvalidArgument, "nas:DescribeAccesspoint: %v", err)
		default:
			logger.V(2).Info("transient DescribeAccesspoint error, retrying within the poll budget",
				"accesspointId", accesspointId, "waiting", p.waiting, "error", err.Error())
		}

		if p.done(ap, notFound) {
			return nil
		}
		if time.Now().After(deadline) {
			return p.timeout(lastStatus)
		}
		logger.V(4).Info("accesspoint not in the target state yet, retrying",
			"accesspointId", accesspointId, "waiting", p.waiting, "status", lastStatus)

		// Correct only on Go >= 1.23, whose timer semantics guarantee Reset/Stop never deliver a stale value.
		timer.Reset(c.apPollInterval)
		select {
		case <-ctx.Done():
			return status.Errorf(codes.DeadlineExceeded, "waiting for accesspoint %s %s: %v", accesspointId, p.waiting, ctx.Err())
		case <-timer.C:
		}
	}
}

// The compensating action for a CreateVolume that failed after the AgenticSpace was created. It never
// deletes the space (DeleteAgenticSpace needs every accesspoint detached, which this does not wait
// for), deletes only the accesspoint this call created, and fires only on a terminal gRPC code.
func (c *agenticfsController) compensateCreateVolume(ctx context.Context, logger klog.Logger, filesystemId, agenticSpaceId, fileSystemPath, createdAccesspointId string, cause error, rpcStarted time.Time) {
	code := status.Code(cause)

	// One latch for every reporting exit, so "exactly one of {orphan, resolved}" holds.
	var reported atomic.Bool
	// logErr is attached at Error level; cause is always the CreateVolume failure, in the "cause" key.
	reportLeak := func(reason string, logErr error) {
		if !reported.CompareAndSwap(false, true) {
			return
		}
		if spaceMayExist(agenticSpaceId, code) {
			c.logOrphanResource(logger, filesystemId, agenticSpaceId, fileSystemPath, createdAccesspointId, reason, cause, logErr)
			return
		}
		// A prefix-less diagnostic, not a confirmed leak; see spaceMayExist for the unconfirmed premise.
		c.logCompensationDiagnostic(logger, filesystemId, agenticSpaceId, fileSystemPath, createdAccesspointId,
			"terminal CreateVolume failure with no AgenticSpace in existence: nothing to reap",
			reason, cause, "code", code)
	}
	reportResolved := func(reason string) {
		if !reported.CompareAndSwap(false, true) {
			return
		}
		c.logOrphanResolved(logger, filesystemId, agenticSpaceId, fileSystemPath, createdAccesspointId, reason, cause)
	}

	if !isTerminalCompensationCode(code) {
		// codes.Unknown is final to external-provisioner but retryable here. Unreachable today; logged so it cannot leak.
		switch code {
		case codes.Internal, codes.Aborted, codes.DeadlineExceeded, codes.Unavailable, codes.Canceled:
			// The retryable codes this controller actually emits.
		default:
			// Carries "reason" like every other compensation line, so a reaper sees a single schema.
			c.logCompensationDiagnostic(logger, filesystemId, agenticSpaceId, fileSystemPath, createdAccesspointId,
				fmt.Sprintf("unclassified gRPC code %s; treating as retryable", code),
				fmt.Sprintf("CreateVolume failed with gRPC code %s, which this controller never emits and which is not in isTerminalCompensationCode's terminal set; it is treated as RETRYABLE, so no compensating delete was sent and the accesspoint/agenticspace are kept for the next attempt. If external-provisioner actually treats this code as final, the resources leak silently - that is what this line is for", code),
				cause, "code", code)
		}
		// Leave both for the next attempt. Deliberately not the orphan prefix - cleaning these up destroys a live space.

		logger.Info(retainedForRetryLogPrefix+": CreateVolume failed with a retryable code, the accesspoint/agenticspace are kept for the next attempt - NOT an orphan, do not clean them up by hand",
			c.compensationFields(filesystemId, agenticSpaceId, fileSystemPath, createdAccesspointId,
				fmt.Sprintf("CreateVolume failed with retryable code %s; the agenticspace and the accesspoint are retained for the next attempt, which reuses them via the CreateAgenticSpace ClientToken replay and findReusableAccessPoint", code),
				cause)...)
		return
	}

	// Terminal: the provisioner will not retry, so delete exactly the accesspoint this call created.

	if createdAccesspointId == "" {
		// Checked before the budget gate on purpose: when both hold, "no accesspoint exists" is the accurate reason.
		reportLeak(fmt.Sprintf("CreateVolume failed terminally with code %s; no accesspoint was created by this call so no compensating delete was sent, the agenticspace is left to the reaper", code), nil)
		return
	}

	// A terminal verdict can land long after the sidecar's deadline with ctx.Err() still nil.
	if open, skipReason := c.compensationWindowOpen(ctx, rpcStarted); !open {
		reportLeak(fmt.Sprintf("CreateVolume failed terminally with code %s and %s; the compensating nas:DeleteAccesspoint was NOT ATTEMPTED, the accesspoint and the agenticspace are left to the reaper", code, skipReason), nil)
		return
	}

	// No tier prefix: the orphan or resolved line comes only once the outcome is known, so a reaper counts one.
	logger.Info("compensating: about to delete the accesspoint this call created (the agenticspace is left to the reaper by design)",
		c.compensationFields(filesystemId, agenticSpaceId, fileSystemPath, createdAccesspointId,
			fmt.Sprintf("CreateVolume failed terminally with code %s; deleting the accesspoint this call created", code),
			cause)...)

	// Neither inherits the (usually expired) request deadline nor can be cancelled, and is bounded.
	cleanupCtx, cancel := context.WithTimeout(context.WithoutCancel(ctx), c.compTimeout)
	defer cancel()

	// ctx never reaches the SDK, so this select is what caps the per-volume lock hold.
	waitBudget := c.compTimeout - nasAPICallBound
	if waitBudget <= 0 {
		// Degenerate only in tests; clamped to the full budget rather than 0 so the delete stays reachable.
		waitBudget = c.compTimeout
	}
	done := make(chan compensateOutcome, 1)
	go func() {
		// Deliberately not nil: an early return added later would otherwise be reported as "reconciled".
		outcome := compensateOutcome{err: errCompensateDeleteNoOutcome}
		// Detached and can outlive the RPC; reports before sending the sentinel the select relies on.
		defer func() {
			r := recover()
			if r != nil {
				outcome.err = errCompensateDeletePanicked
				outcome.panicValue = r
			}
			if r != nil {
				func() {
					defer func() { _ = recover() }()
					// A no-op when the select already reported, which is the double count the latch exists to stop.
					reportLeak(fmt.Sprintf("CreateVolume failed terminally with code %s and the compensating nas:DeleteAccesspoint PANICKED (%v); the accesspoint and the agenticspace are both left behind", code, r), nil)
				}()
			}
			done <- outcome
		}()
		// waitCtx, not cleanupCtx, reaches limiter.Wait, so no token can be granted late enough to outlive this.
		waitCtx, cancelWait := context.WithTimeout(cleanupCtx, waitBudget)
		defer cancelWait()
		if err := waitCtx.Err(); err != nil {
			// No wire request exists; reported distinctly from "issued and failed", not guessed from a string.
			outcome.err = fmt.Errorf("%w: %v", errCompensateDeleteNeverSent, err)
			return
		}
		err := c.nasClient.DeleteAccesspoint(waitCtx, filesystemId, createdAccesspointId)
		outcome.err = err
		// "Sent" means "got past the rate limiter", which is what deleteNeverSent recognises.
		outcome.sent = !deleteNeverSent(err)
	}()
	select {
	case outcome := <-done:
		switch {
		case errors.Is(outcome.err, errCompensateDeletePanicked):
			// Already reported by the goroutine's recover above.
		case !outcome.sent:
			// Nothing is in flight and nothing was deleted, so the reason must not claim a failed delete.
			reportLeak(fmt.Sprintf("CreateVolume failed terminally with code %s and the compensating nas:DeleteAccesspoint was NEVER SENT (%v) - it never got past the wait budget / shared rate limiter, so no wire request exists; the accesspoint and the agenticspace are both left behind for the reaper", code, outcome.err), outcome.err)
		case outcome.err != nil && !isNotFoundError(outcome.err):
			reportLeak(fmt.Sprintf("CreateVolume failed terminally with code %s and the compensating nas:DeleteAccesspoint was sent and FAILED (%v); the accesspoint and the agenticspace are both left behind", code, outcome.err), outcome.err)
		default:
			// Without the resolved line the orphan stream cannot self-reconcile.
			outcomeText := "succeeded"
			if outcome.err != nil {
				outcomeText = fmt.Sprintf("answered NotFound (%v), the accesspoint was already gone", outcome.err)
			}
			reportResolved(fmt.Sprintf("CreateVolume failed terminally with code %s and the compensating nas:DeleteAccesspoint %s; the accesspoint is reconciled, the agenticspace is still left to the reaper by design", code, outcomeText))
		}
	case <-cleanupCtx.Done():
		// Neither can outlive this function, so the lock is never released while a delete of this volume is in flight.
		reportLeak(fmt.Sprintf("CreateVolume failed terminally with code %s and the compensating nas:DeleteAccesspoint did not return within %s (%v); the delete may still be in flight OR may never have been sent at all (the goroutine can still be queued on the shared rate limiter when the bound fires, in which case limiter.Wait returns context.DeadlineExceeded and the wire request is never issued). Neither case can outlive this call: the wait phase is capped at %s and the wire phase at %s - reconcile against the accesspoint's actual cloud state before assuming either", code, c.compTimeout, cleanupCtx.Err(), waitBudget, nasAPICallBound), cleanupCtx.Err())
	}
}

// Deliberately a superset of what is reachable today, so a code that becomes reachable later is still compensated.
func isTerminalCompensationCode(code codes.Code) bool {
	switch code {
	case codes.InvalidArgument, codes.OutOfRange, codes.FailedPrecondition,
		codes.PermissionDenied, codes.Unauthenticated, codes.Unimplemented:
		return true
	default:
		return false
	}
}

// Sent after the goroutine recovers, so the select does not emit a second line for the same accesspoint.
var errCompensateDeletePanicked = errors.New("compensating DeleteAccesspoint panicked")

// The wait phase expired before DeleteAccesspoint was called, so no wire request was ever made.
var errCompensateDeleteNeverSent = errors.New("compensating DeleteAccesspoint was never sent")

// Unreachable: every path overwrites it. Makes a future early return loud instead of "reconciled".

var errCompensateDeleteNoOutcome = errors.New("compensating DeleteAccesspoint produced no outcome")

// A struct, not a bare error: a limiter.Wait error and an HTTP error look alike but mean opposites.
type compensateOutcome struct {
	err error
	// True only once the call got past the rate limiter, i.e. a wire request may exist and be in flight.
	sent bool
	// Kept for diagnostics; the sentinel err is what the select branches on.
	panicValue any
}

// pkg/nas/cloud's wait() wraps every rate-limiter failure before the request is built - the only signal.
func deleteNeverSent(err error) bool {
	return err != nil && strings.Contains(err.Error(), nasRateLimiterWaitPrefix)
}

func causeText(cause error) string {
	if cause == nil {
		return "<nil>"
	}
	return cause.Error()
}

// Only an ambiguous failure can have created a space whose ID never came back, and those are all retryable.
func spaceMayExist(agenticSpaceId string, code codes.Code) bool {
	return agenticSpaceId != "" || !isTerminalCompensationCode(code)
}

// Elapsed comes from rpcStarted because ctx.Err() can still be nil long after the provisioner's deadline.
func (c *agenticfsController) compensationWindowOpen(ctx context.Context, rpcStarted time.Time) (bool, string) {
	if err := ctx.Err(); err != nil {
		return false, fmt.Sprintf("its request context is already done (%v), so external-provisioner has cancelled or timed out this RPC and will never observe the terminal verdict; skipping the delete also releases the per-volume lock at once instead of holding it for another %s", err, c.compTimeout)
	}
	budget := c.rpcBudget
	if budget <= 0 {
		budget = driverRPCBudget
	}
	elapsed := time.Since(rpcStarted)
	if remaining := budget - elapsed; remaining < c.compTimeout {
		return false, fmt.Sprintf("the driver-side RPC budget is exhausted (%s elapsed of %s, %s remaining, and the compensating delete alone needs up to %s), so external-provisioner has already given up on this RPC and will never observe the terminal verdict; the delete is NOT ATTEMPTED because its whole premise - a terminal verdict the provisioner still observes, so that giving up means the accesspoint should be cleaned up for it - is void for a caller that is no longer listening, and running it now would only destroy a resource nobody is left to receive; skipping it also releases the per-volume lock at once instead of holding it for another %s", elapsed.Round(time.Millisecond), budget, remaining.Round(time.Millisecond), c.compTimeout, c.compTimeout)
	}
	return true, ""
}

// fileSystemPath is the only key that survives when the AgenticSpaceId was never read back.
func (c *agenticfsController) orphanLogFields(filesystemId, agenticSpaceId, fileSystemPath, accesspointId string) []any {
	return []any{
		"fileSystemId", filesystemId,
		"agenticSpaceId", agenticSpaceId,
		"fileSystemPath", fileSystemPath,
		"accesspointId", accesspointId,
		"region", c.region,
		"volumeHandle", strings.TrimPrefix(fileSystemPath, "/"),
	}
}

func (c *agenticfsController) compensationFields(filesystemId, agenticSpaceId, fileSystemPath, accesspointId, reason string, cause error, extra ...any) []any {
	fields := append(c.orphanLogFields(filesystemId, agenticSpaceId, fileSystemPath, accesspointId),
		"reason", reason, "cause", causeText(cause))
	return append(fields, extra...)
}

// The confirmed-leak line a reaper greps for; exactly one per leak, via the reported latch.
func (c *agenticfsController) logOrphanResource(logger klog.Logger, filesystemId, agenticSpaceId, fileSystemPath, accesspointId, reason string, cause, logErr error) {
	logger.Error(logErr, orphanLogPrefix+": a CreateVolume failure left cloud resources behind",
		c.compensationFields(filesystemId, agenticSpaceId, fileSystemPath, accesspointId, reason, cause)...)
}

func (c *agenticfsController) logOrphanResolved(logger klog.Logger, filesystemId, agenticSpaceId, fileSystemPath, accesspointId, reason string, cause error) {
	logger.Info(orphanResolvedLogPrefix+": the compensating delete reconciled the accesspoint this call created",
		c.compensationFields(filesystemId, agenticSpaceId, fileSystemPath, accesspointId, reason, cause)...)
}

// Full contract schema but no tier prefix, where an Error-level orphan line would be a false positive.
func (c *agenticfsController) logCompensationDiagnostic(logger klog.Logger, filesystemId, agenticSpaceId, fileSystemPath, accesspointId, msg, reason string, cause error, extra ...any) {
	logger.Info(msg,
		c.compensationFields(filesystemId, agenticSpaceId, fileSystemPath, accesspointId, reason, cause, extra...)...)
}

func (c *agenticfsController) DeleteVolume(ctx context.Context, req *csi.DeleteVolumeRequest, pv *corev1.PersistentVolume) (*csi.DeleteVolumeResponse, error) {
	logger := klog.FromContext(ctx)
	attributes := pv.Spec.CSI.VolumeAttributes
	filesystemId := agenticfsFilesystemID(attributes)
	if filesystemId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "missing %s in volume attributes", filesystemIDKey)
	}
	agenticSpaceId := attributes[vcKeyAgenticSpaceId]
	if agenticSpaceId == "" {
		// Reporting success would leak a billable AgenticSpace once the PV is gone; echoes every ID read.
		return nil, status.Errorf(codes.InvalidArgument,
			"missing agenticSpaceId in volume attributes (read fileSystemId=%q, accesspointId=%q, volumeId=%q): refusing to report success because the AgenticSpace would leak; follow the manual cleanup procedure in examples/nas/agenticfs/README.md §4.1",
			filesystemId, attributes[vcKeyAccesspointId], req.VolumeId)
	}

	// An interrupted CreateVolume can have left more than the one in the PV, and deleting one would wedge DeleteVolume.
	listed, err := c.listAccessPointsOfSpace(ctx, filesystemId, agenticSpaceId)
	switch {
	case err == nil:
	case isNotFoundError(err):
		// An absent accesspoint (or an unqualified NotFound from a list) does not
		// prove the billable space is gone. Verify it with a space-scoped call;
		// if it still exists, retry rather than act on an incomplete AP listing.
		code := strings.ToLower(apiErrorCode(err))
		if !strings.Contains(code, "agenticspace") && !strings.Contains(code, "filesystem") {
			_, getErr := c.nasClient.GetAgenticSpace(ctx, &sdk.GetAgenticSpaceRequest{
				FileSystemId:   tea.String(filesystemId),
				AgenticSpaceId: tea.String(agenticSpaceId),
			})
			if !isAgenticSpaceNotFoundError(getErr) {
				if getErr != nil && !isNotFoundError(getErr) {
					return nil, apiStatusError("nas:GetAgenticSpace", getErr)
				}
				return nil, status.Errorf(codes.Aborted,
					"nas:ListAccesspoints: %v; agenticspace %s is not confirmed absent (GetAgenticSpace: %v), retrying the accesspoint listing",
					err, agenticSpaceId, getErr)
			}
			err = getErr
		}
		gone := "agenticspace"
		if strings.Contains(strings.ToLower(apiErrorCode(err)), "filesystem") {
			gone = "filesystem"
		}
		logger.Info(gone+" is gone, treating the volume as already deleted",
			"agenticSpaceId", agenticSpaceId, "errorCode", apiErrorCode(err))
		return &csi.DeleteVolumeResponse{}, nil
	default:
		return nil, apiStatusError("nas:ListAccesspoints", err)
	}

	accesspointIds := make([]string, 0, len(listed)+1)
	addAccesspointId := func(id string) {
		if id == "" {
			return
		}
		for _, existing := range accesspointIds {
			if existing == id {
				return
			}
		}
		accesspointIds = append(accesspointIds, id)
	}
	// First, because the list API can lag a very recently created accesspoint that must still be deleted.
	addAccesspointId(attributes[vcKeyAccesspointId])
	for _, ap := range listed {
		addAccesspointId(tea.StringValue(ap.AccessPointId))
	}
	if len(accesspointIds) == 0 {
		// ListAccesspoints can lag a recently created accesspoint, so this may self-heal on a retry.
		return nil, status.Errorf(codes.Aborted,
			"missing accesspointId in volume attributes and agenticspace %s has no accesspoint listed yet: the listing may lag a recently created accesspoint, retrying; if it persists follow the manual cleanup procedure in examples/nas/agenticfs/README.md §4.1",
			agenticSpaceId)
	}
	if attributes[vcKeyAccesspointId] == "" {
		logger.Info("missing accesspointId in volume attributes, falling back to the accesspoints listed for the agenticspace",
			"agenticSpaceId", agenticSpaceId, "accesspointIds", accesspointIds)
	}

	// Every accesspoint must be detached (and gone) before the AgenticSpace can be deleted.
	for _, accesspointId := range accesspointIds {
		err := c.nasClient.DeleteAccesspoint(ctx, filesystemId, accesspointId)
		if err != nil && !isNotFoundError(err) {
			return nil, apiStatusError("nas:DeleteAccesspoint", err)
		}
		if err != nil {
			// One still mid-Deleting makes the following DeleteAgenticSpace fail for several rounds.
			logger.Info("accesspoint already deleted", "accesspointId", accesspointId)
		}
		if err := c.waitAccessPointGone(ctx, filesystemId, accesspointId); err != nil {
			return nil, err
		}
	}

	if _, err := c.nasClient.DeleteAgenticSpace(ctx, &sdk.DeleteAgenticSpaceRequest{
		FileSystemId:   tea.String(filesystemId),
		AgenticSpaceId: tea.String(agenticSpaceId),
		ClientToken:    tea.String(req.VolumeId),
	}); err != nil {
		if !isAgenticSpaceNotFoundError(err) {
			if isNotFoundError(err) {
				return nil, status.Errorf(codes.Aborted,
					"nas:DeleteAgenticSpace: %v; an absent accesspoint does not confirm agenticspace %s was deleted, retrying", err, agenticSpaceId)
			}
			return nil, apiStatusError("nas:DeleteAgenticSpace", err)
		}
		logger.Info("agenticspace already deleted", "agenticSpaceId", agenticSpaceId)
	}

	return &csi.DeleteVolumeResponse{}, nil
}

func (c *agenticfsController) ControllerExpandVolume(ctx context.Context, req *csi.ControllerExpandVolumeRequest, pv *corev1.PersistentVolume) (*csi.ControllerExpandVolumeResponse, error) {
	logger := klog.FromContext(ctx)
	attributes := pv.Spec.CSI.VolumeAttributes
	filesystemId := agenticfsFilesystemID(attributes)
	if filesystemId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "missing %s in volume attributes", filesystemIDKey)
	}
	agenticSpaceId := attributes[vcKeyAgenticSpaceId]
	if agenticSpaceId == "" {
		return nil, status.Error(codes.InvalidArgument, "missing agenticSpaceId in volume attributes")
	}

	sizeLimit, err := computeExpandSizeLimit(req.GetCapacityRange())
	if err != nil {
		return nil, err
	}

	getResp, err := c.nasClient.GetAgenticSpace(ctx, &sdk.GetAgenticSpaceRequest{
		FileSystemId:   tea.String(filesystemId),
		AgenticSpaceId: tea.String(agenticSpaceId),
	})
	if err != nil {
		return nil, apiStatusError("nas:GetAgenticSpace", err)
	}
	// A nil would make a guard below evaluate against 0. A nil or 0 FileCountLimit is simply omitted.
	if getResp == nil || getResp.Body == nil || getResp.Body.AgenticSpace == nil ||
		getResp.Body.AgenticSpace.Quota == nil ||
		getResp.Body.AgenticSpace.Quota.SizeLimit == nil ||
		getResp.Body.AgenticSpace.SpaceUsage == nil ||
		getResp.Body.AgenticSpace.FileCountUsage == nil {
		return nil, status.Error(codes.Internal,
			"nas:GetAgenticSpace: incomplete quota/usage in response; refusing to guess because SetAgenticSpaceQuota overwrites the quota and the usage guards cannot run against a missing value")
	}
	space := getResp.Body.AgenticSpace
	currentSizeLimit := tea.Int64Value(space.Quota.SizeLimit)
	spaceUsage := tea.Int64Value(space.SpaceUsage)
	fileCountUsage := tea.Int64Value(space.FileCountUsage)
	// The OpenAPI minimum is 10000, so 0 is never legal. Omitting assumes the API preserves it - undocumented.
	fileCountLimit := tea.Int64Value(space.Quota.FileCountLimit)
	sendsFileCountLimit := space.Quota.FileCountLimit != nil && fileCountLimit != 0
	var fileCountLimitReq *int64
	if sendsFileCountLimit {
		fileCountLimitReq = tea.Int64(fileCountLimit)
	}

	// A clear OutOfRange beats an opaque cloud error, and OutOfRange is permanent so the resizer stops.
	if sizeLimit < currentSizeLimit {
		return nil, status.Errorf(codes.OutOfRange,
			"shrinking an AgenticSpace is not supported: requested %d bytes, current size limit is %d bytes", sizeLimit, currentSizeLimit)
	}
	if sizeLimit < spaceUsage {
		return nil, status.Errorf(codes.OutOfRange,
			"requested size limit %d bytes is below the %d bytes already used by agenticspace %s", sizeLimit, spaceUsage, agenticSpaceId)
	}
	// OutOfRange is terminal and would permanently block an otherwise valid grow, hence a warning.
	if sendsFileCountLimit && fileCountLimit < fileCountUsage {
		logger.Info("WARNING: agenticspace file count limit is below its current usage; this expansion does not modify the file count limit, proceeding",
			"agenticSpaceId", agenticSpaceId, "fileCountLimit", fileCountLimit, "fileCountUsage", fileCountUsage)
	}

	if sendsFileCountLimit {
		if err := validateAgenticSpaceFileCountLimit(fileCountLimit); err != nil {
			logger.Info("WARNING: the file count limit read back from nas:GetAgenticSpace is outside the OpenAPI range; this expansion does not modify it, so the field is OMITTED from SetAgenticSpaceQuota and the size expansion proceeds",
				"agenticSpaceId", agenticSpaceId, "fileCountLimit", fileCountLimit,
				"range", fmt.Sprintf("[%d, %d]", minAgenticSpaceFileCountLimit, maxAgenticSpaceFileCountLimit),
				"validationError", err.Error())
			fileCountLimitReq = nil
		}
	}

	// Reuses the create-path clause so the two cannot drift; unlike the file count clause this stays terminal.
	if err := validateAgenticSpaceSizeLimit(sizeLimit); err != nil {
		return nil, err
	}

	// Flat SizeLimit/FileCountLimit, not a nested Quota and not SetDirQuota; FileCountLimit stays nil when unread.
	if _, err := c.nasClient.SetAgenticSpaceQuota(ctx, &sdk.SetAgenticSpaceQuotaRequest{
		FileSystemId:   tea.String(filesystemId),
		AgenticSpaceId: tea.String(agenticSpaceId),
		SizeLimit:      tea.Int64(sizeLimit),
		FileCountLimit: fileCountLimitReq,
	}); err != nil {
		return nil, apiStatusError("nas:SetAgenticSpaceQuota", err)
	}
	return &csi.ControllerExpandVolumeResponse{CapacityBytes: sizeLimit}, nil
}

// Not a fallback: external-provisioner always fills required_bytes, so one would ignore an administrator's cap.
func computeAgenticSpaceSizeLimit(cr *csi.CapacityRange, sizeLimitParam string) (int64, error) {
	var cap int64
	if sizeLimitParam != "" {
		q, err := resource.ParseQuantity(sizeLimitParam)
		if err != nil {
			return 0, status.Errorf(codes.InvalidArgument, "invalid parameters.agenticSpaceSizeLimit %q: %v", sizeLimitParam, err)
		}
		cap = q.Value()
		// A non-positive cap is a misconfiguration, not "no cap": dropping it would re-create the hazard below.
		if cap <= 0 {
			return 0, status.Errorf(codes.InvalidArgument,
				"parameters.agenticSpaceSizeLimit %q must be a positive capacity", sizeLimitParam)
		}
	}

	bytes := cr.GetRequiredBytes()
	switch {
	case bytes <= 0:
		if cap <= 0 {
			return 0, status.Error(codes.InvalidArgument,
				"no capacity specified: neither PVC capacity nor parameters.agenticSpaceSizeLimit is set")
		}
		bytes = cap
	case cap > 0 && bytes > cap:
		return 0, status.Errorf(codes.InvalidArgument,
			"requested capacity %d bytes exceeds the cap parameters.agenticSpaceSizeLimit=%q (%d bytes); lower the PVC request or raise the StorageClass cap",
			bytes, sizeLimitParam, cap)
	}
	return roundUpToGiBChecked(bytes, cr.GetLimitBytes())
}

// Applies the same rounding, minimum, maximum and limit_bytes checks as the create path.
func computeExpandSizeLimit(cr *csi.CapacityRange) (int64, error) {
	bytes := cr.GetRequiredBytes()
	if bytes <= 0 {
		return 0, status.Error(codes.InvalidArgument, "capacity_range.required_bytes must be set to expand an AgenticSpace")
	}
	sizeLimit, err := roundUpToGiBChecked(bytes, cr.GetLimitBytes())
	if err != nil {
		return 0, err
	}
	if sizeLimit < minAgenticSpaceSizeLimit {
		return 0, status.Errorf(codes.InvalidArgument,
			"expanded size limit %d bytes is below the minimum %d bytes (10 GiB)", sizeLimit, minAgenticSpaceSizeLimit)
	}
	return sizeLimit, nil
}

// Rejects overflow and the OpenAPI maximum: a cloud rejection comes back retryable and loops forever.
func roundUpToGiBChecked(bytes, limitBytes int64) (int64, error) {
	if bytes > math.MaxInt64-GiB+1 {
		return 0, status.Errorf(codes.OutOfRange, "capacity %d bytes overflows the 1 GiB rounding", bytes)
	}
	sizeLimit := roundUpToGiB(bytes)
	if sizeLimit > maxAgenticSpaceSizeLimit {
		return 0, status.Errorf(codes.InvalidArgument,
			"capacity %d bytes exceeds the AgenticSpace maximum of %d bytes (1024000 GiB)", sizeLimit, maxAgenticSpaceSizeLimit)
	}
	// Rounding up can push the result over limit_bytes, which CSI requires it to stay within.
	if limitBytes > 0 && sizeLimit > limitBytes {
		return 0, status.Errorf(codes.OutOfRange,
			"capacity %d bytes rounded up to the 1 GiB boundary (%d bytes) exceeds capacity_range.limit_bytes (%d bytes)",
			bytes, sizeLimit, limitBytes)
	}
	return sizeLimit, nil
}

func computeAgenticSpaceFileCountLimit(param string) (int64, error) {
	if param == "" {
		return defaultAgenticSpaceFileCountLimit, nil
	}
	v, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		return 0, status.Errorf(codes.InvalidArgument, "invalid parameters.agenticSpaceFileCountLimit %q: %v", param, err)
	}
	return v, nil
}

// The expand path must not use it: a value echoed by nas:GetAgenticSpace would wedge an unrelated expansion.

func validateAgenticSpaceQuota(sizeLimit, fileCountLimit int64) error {
	if err := validateAgenticSpaceSizeLimit(sizeLimit); err != nil {
		return err
	}
	return validateAgenticSpaceFileCountLimit(fileCountLimit)
}

// Terminal on both paths: the size limit always comes from the caller, never a cloud echo.
func validateAgenticSpaceSizeLimit(sizeLimit int64) error {
	if sizeLimit < minAgenticSpaceSizeLimit {
		return status.Errorf(codes.InvalidArgument,
			"size limit %d bytes is below the minimum %d bytes (10 GiB)", sizeLimit, minAgenticSpaceSizeLimit)
	}
	return nil
}

// Whether a violation is TERMINAL depends on where the value came from - see validateAgenticSpaceQuota.
func validateAgenticSpaceFileCountLimit(fileCountLimit int64) error {
	if fileCountLimit < minAgenticSpaceFileCountLimit || fileCountLimit > maxAgenticSpaceFileCountLimit {
		return status.Errorf(codes.InvalidArgument,
			"file count limit %d is out of range [%d, %d]", fileCountLimit, minAgenticSpaceFileCountLimit, maxAgenticSpaceFileCountLimit)
	}
	return nil
}

func roundUpToGiB(bytes int64) int64 {
	return (bytes + GiB - 1) / GiB * GiB
}

// Guards req.Name, which is interpolated into FileSystemPath, ClientToken and AccessPointName.
func validateVolumeName(name string) error {
	if name == "" {
		return status.Error(codes.InvalidArgument, "volume name is empty")
	}
	// A space would make FileSystemPath and ClientToken differ from a later attempt's, breaking the replay.
	if strings.TrimSpace(name) != name {
		return status.Errorf(codes.InvalidArgument,
			"volume name %q must not have leading or trailing whitespace", name)
	}
	if len(name) > maxVolumeNameLen {
		// len counts bytes, so the message says "bytes"; the loop below rejects any non-ASCII byte anyway.
		return status.Errorf(codes.InvalidArgument,
			"volume name %q is longer than %d bytes", name, maxVolumeNameLen)
	}
	if strings.ContainsAny(name, `/\`) {
		return status.Errorf(codes.InvalidArgument, "volume name %q must not contain '/' or '\\'", name)
	}
	// An embedded ".." is inert: req.Name is one first-level segment and '/' is already rejected.
	if name == "." || name == ".." {
		return status.Errorf(codes.InvalidArgument, "volume name %q must not be '.' or '..'", name)
	}
	for _, r := range name {
		if r > unicode.MaxASCII || !unicode.IsPrint(r) {
			return status.Errorf(codes.InvalidArgument, "volume name %q must only contain printable ASCII characters", name)
		}
	}
	// With '/' and '\' rejected, "/"+name cannot escape the first level.
	return nil
}

// Reads the canonical filesystemIDKey the node side looks up, plus the legacy spelling on read.
func agenticfsFilesystemID(attributes map[string]string) string {
	if id := attributes[filesystemIDKey]; id != "" {
		return id
	}
	return attributes[vcKeyFilesystemIdLegacy]
}

// A v2 SDK failure arrives as briefAliError, which unwraps to a *teaError exposing ErrorCode().
type aliErrorCode interface {
	ErrorCode() string
}

// NAS pads error codes, so one trailing newline defeats every exact comparison downstream.
func apiErrorCode(err error) string {
	var e aliErrorCode
	if errors.As(err, &e) {
		return strings.TrimSpace(e.ErrorCode())
	}
	return ""
}

// Exact, not a ".NotFound" suffix test, which would swallow other verdicts.
func isNotFoundError(err error) bool {
	if err == nil {
		return false
	}
	if errors.Is(err, wrap.ErrorCode("NotFound")) {
		return true
	}
	switch strings.ToLower(apiErrorCode(err)) {
	case "notfound",
		"invalidagenticspaceid.notfound",
		"invalidaccesspoint.notfound",
		"invalidaccesspointid.notfound",
		"invalidfilesystem.notfound",
		"invalidfilesystemid.notfound":
		return true
	}
	return false
}

// Only for space-scoped Get/Delete calls: an unqualified NotFound refers to
// the requested space, but an explicitly missing accesspoint never does.
func isAgenticSpaceNotFoundError(err error) bool {
	switch strings.ToLower(apiErrorCode(err)) {
	case "invalidaccesspoint.notfound", "invalidaccesspointid.notfound":
		return false
	}
	return isNotFoundError(err)
}

// Reads the code only, never the message; the region list lives server side, so none is hardcoded.

// Whole codes, lower-cased against apiErrorCode's trimmed value. InvalidArgument, so the provisioner stops.
var permanentAPIErrorCodes = []string{
	"invalidregionid.notfound",
	"invalidzoneid.notfound",
	"forbidden.regiondisabled",
	"forbidden.region",
}

// Permanent regardless of suffix: the request's own parameters are wrong, or the region/account is off.
var permanentAPIErrorPrefixes = []string{
	"invalidparameter",
	"invalidparam",
	"invalidregion",
	"invalidzone",
	"forbidden.",
}

func isPermanentAPIError(err error) bool {
	code := strings.ToLower(apiErrorCode(err))
	if code == "" {
		return false
	}
	for _, c := range permanentAPIErrorCodes {
		if code == c {
			return true
		}
	}
	for _, p := range permanentAPIErrorPrefixes {
		if strings.HasPrefix(code, p) {
			return true
		}
	}
	// Each rule needs its structural second half, so a state-class *.NotSupported stays retryable.
	if (strings.Contains(code, "region") || strings.Contains(code, "zone")) && strings.HasSuffix(code, "notsupported") {
		return true
	}
	if strings.HasPrefix(code, "unsupported") && (strings.Contains(code, "region") || strings.Contains(code, "zone")) {
		return true
	}
	// Callers that treat NotFound as success check isNotFoundError first; this keeps one whitelist.

	if isNotFoundError(err) {
		return true
	}
	return false
}

// Passes through errors this package already classified; leaves everything else retryable.
func apiStatusError(op string, err error) error {
	if err == nil {
		return nil
	}
	if _, ok := status.FromError(err); ok {
		return err
	}
	switch {
	case errors.Is(err, context.Canceled), errors.Is(err, context.DeadlineExceeded):
		return status.Errorf(codes.DeadlineExceeded, "%s: %v", op, err)
	case isPermanentAPIError(err):
		return status.Errorf(codes.InvalidArgument, "%s: permanent OpenAPI error: %v", op, err)
	default:
		return status.Errorf(codes.Internal, "%s: %v", op, err)
	}
}
