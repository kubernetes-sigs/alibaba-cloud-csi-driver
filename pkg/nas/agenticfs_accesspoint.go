//go:build !windows

package nas

import (
	"context"
	"fmt"
	"strings"
	"time"

	sdk "github.com/alibabacloud-go/nas-20170626/v4/client"
	"github.com/alibabacloud-go/tea/tea"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"k8s.io/klog/v2"
)

const (
	accessPointStatusActive   = "active"
	accessPointStatusDeleting = "deleting"
	accessPointStatusInactive = "inactive"

	// Tags aid observation; only the AgenticSpaceId relationship proves ownership.
	apTagKeyVolumeID                 = "csi.alibabacloud.com/volume-id"
	apListFilterAgenticSpaceId       = "AgenticSpaceId"
	apListMaxResults           int32 = 100
	apListMaxPages                   = 50

	defaultApPollInterval = 3 * time.Second
	defaultApPollTimeout  = 45 * time.Second
)

// Discovery is mandatory: CreateAccessPoint has no ClientToken. Return the ID
// before polling so a later failure can report the resource already observed.
func (c *agenticfsController) ensureAccessPoint(ctx context.Context, args *agenticfsVolumeArgs, agenticSpaceID string) (string, string, error) {
	id, server, err := c.findReusableAccessPoint(ctx, args.FileSystemID, agenticSpaceID)
	if err != nil {
		return "", "", err
	}
	if id != "" {
		klog.FromContext(ctx).Info("reusing the accesspoint already bound to the agenticspace",
			"agenticSpaceId", agenticSpaceID, "accesspointId", id)
		return id, server, nil
	}

	resp, err := c.nasClient.CreateAccesspoint(ctx, &sdk.CreateAccessPointRequest{
		FileSystemId:    tea.String(args.FileSystemID),
		AgenticSpaceId:  tea.String(agenticSpaceID),
		VswId:           tea.String(args.VSwitchID),
		VpcId:           tea.String(args.VpcID),
		AccessPointName: tea.String(args.Name),
		EnabledRam:      tea.Bool(true),
		Tag: []*sdk.CreateAccessPointRequestTag{{
			Key:   tea.String(apTagKeyVolumeID),
			Value: tea.String(args.Name),
		}},
	})
	if err != nil {
		if isNotFoundError(err) {
			return "", "", status.Errorf(codes.InvalidArgument,
				"nas:CreateAccesspoint: agenticspace %s was rejected as missing (%v); it may have been deleted externally while the ClientToken still replays its ID - delete and recreate the PVC",
				agenticSpaceID, err)
		}
		return "", "", apiStatusError("nas:CreateAccesspoint", err)
	}
	if resp == nil || resp.Body == nil || resp.Body.AccessPoint == nil {
		return "", "", status.Error(codes.Internal, "nas:CreateAccesspoint: empty AccessPoint in response")
	}
	id = tea.StringValue(resp.Body.AccessPoint.AccessPointId)
	if id == "" {
		return "", "", status.Error(codes.Internal, "nas:CreateAccesspoint: empty AccessPointId in response")
	}
	return id, tea.StringValue(resp.Body.AccessPoint.AccessPointDomain), nil
}

// Normally a space has zero or one accesspoint: create only when none exists,
// otherwise reuse it or wait for it to become usable. Multiple entries are recovery
// residue, not a placement-selection feature; prefer Active without creating more.
// listAccessPointsOfSpace verifies ownership before any entry can be reused.
func (c *agenticfsController) findReusableAccessPoint(ctx context.Context, filesystemId, agenticSpaceId string) (accesspointId, domain string, err error) {
	logger := klog.FromContext(ctx)
	accesspoints, err := c.listAccessPointsOfSpace(ctx, filesystemId, agenticSpaceId)
	if err != nil {
		return "", "", apiStatusError("nas:ListAccesspoints", err)
	}
	if len(accesspoints) > 1 {
		logger.Info("unexpected multiple accesspoints; recovering an existing one without creating another",
			"agenticSpaceId", agenticSpaceId, "count", len(accesspoints))
	}
	var pending *sdk.ListAccessPointsResponseBodyAccessPoints
	var unavailable []string
	for _, ap := range accesspoints {
		id := tea.StringValue(ap.AccessPointId)
		if id == "" {
			continue
		}
		switch strings.ToLower(tea.StringValue(ap.Status)) {
		case accessPointStatusActive:
			return id, tea.StringValue(ap.DomainName), nil
		case accessPointStatusDeleting, accessPointStatusInactive:
			unavailable = append(unavailable, fmt.Sprintf("%s(%s)", id, tea.StringValue(ap.Status)))
		default:
			if pending == nil {
				pending = ap
			}
		}
	}
	if pending != nil {
		return tea.StringValue(pending.AccessPointId), tea.StringValue(pending.DomainName), nil
	}
	if len(unavailable) > 0 {
		return "", "", status.Errorf(codes.Aborted,
			"nas:ListAccesspoints: the accesspoint(s) of agenticspace %s are not usable yet: %s; "+
				"inspect the NAS console and follow examples/nas/agenticfs/README.md §4.1 for recovery or cleanup; "+
				"deleting a Pending PVC does not guarantee cloud resource cleanup",
			agenticSpaceId, strings.Join(unavailable, ", "))
	}
	return "", "", nil
}

// Follow every page and verify the server-side filter before returning a list.
// Callers must never create or delete based on a partial/unproven result.
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
			return nil, status.Errorf(codes.Aborted,
				"nas:ListAccesspoints: page %d of the agenticspace %s enumeration returned a malformed response (nil body); refusing to act on a partial list",
				page, agenticSpaceId)
		}
		for _, ap := range resp.Body.AccessPoints {
			got := tea.StringValue(ap.AgenticSpaceId)
			if got == agenticSpaceId {
				continue
			}
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

// spaceGone is true only when the listing or a space-scoped verification proves
// absence. Otherwise a successful result contains at least one AP to delete.
func (c *agenticfsController) accessPointsForDeletion(ctx context.Context, filesystemID, agenticSpaceID, knownAccesspointID string) (ids []string, spaceGone bool, err error) {
	logger := klog.FromContext(ctx)
	listed, err := c.listAccessPointsOfSpace(ctx, filesystemID, agenticSpaceID)
	switch {
	case err == nil:
	case isNotFoundError(err):
		code := strings.ToLower(apiErrorCode(err))
		if !strings.Contains(code, "agenticspace") && !strings.Contains(code, "filesystem") {
			_, getErr := c.nasClient.GetAgenticSpace(ctx, &sdk.GetAgenticSpaceRequest{
				FileSystemId:   tea.String(filesystemID),
				AgenticSpaceId: tea.String(agenticSpaceID),
			})
			if !isAgenticSpaceNotFoundError(getErr) {
				if getErr != nil && !isNotFoundError(getErr) {
					return nil, false, apiStatusError("nas:GetAgenticSpace", getErr)
				}
				return nil, false, status.Errorf(codes.Aborted,
					"nas:ListAccesspoints: %v; agenticspace %s is not confirmed absent (GetAgenticSpace: %v), retrying the accesspoint listing",
					err, agenticSpaceID, getErr)
			}
			err = getErr
		}
		gone := "agenticspace"
		if strings.Contains(strings.ToLower(apiErrorCode(err)), "filesystem") {
			gone = "filesystem"
		}
		logger.Info(gone+" is gone, treating the volume as already deleted",
			"agenticSpaceId", agenticSpaceID, "errorCode", apiErrorCode(err))
		return nil, true, nil
	default:
		return nil, false, apiStatusError("nas:ListAccesspoints", err)
	}

	ids = accessPointDeletionOrder(knownAccesspointID, listed)
	if len(ids) == 0 {
		return nil, false, status.Errorf(codes.Aborted,
			"missing accesspointId in volume attributes and agenticspace %s has no accesspoint listed yet: the listing may lag a recently created accesspoint, retrying; if it persists follow the manual cleanup procedure in examples/nas/agenticfs/README.md §4.1",
			agenticSpaceID)
	}
	if knownAccesspointID == "" {
		logger.Info("missing accesspointId in volume attributes, falling back to the accesspoints listed for the agenticspace",
			"agenticSpaceId", agenticSpaceID, "accesspointIds", ids)
	}
	return ids, false, nil
}

// Preserve PV-first ordering (listing can lag) and first occurrence order, while
// avoiding repeated deletes when the PV and listing refer to the same AP.
func accessPointDeletionOrder(knownID string, listed []*sdk.ListAccessPointsResponseBodyAccessPoints) []string {
	ids := make([]string, 0, len(listed)+1)
	seen := make(map[string]struct{}, len(listed)+1)
	add := func(id string) {
		if id == "" {
			return
		}
		if _, ok := seen[id]; ok {
			return
		}
		seen[id] = struct{}{}
		ids = append(ids, id)
	}
	add(knownID)
	for _, ap := range listed {
		add(tea.StringValue(ap.AccessPointId))
	}
	return ids
}

func (c *agenticfsController) deleteAccessPoint(ctx context.Context, filesystemID, accesspointID string) error {
	err := c.nasClient.DeleteAccesspoint(ctx, filesystemID, accesspointID)
	if err != nil && !isNotFoundError(err) {
		return apiStatusError("nas:DeleteAccesspoint", err)
	}
	if err != nil {
		klog.FromContext(ctx).Info("accesspoint already deleted", "accesspointId", accesspointID)
	}
	// An accepted delete or NotFound does not replace the existing disappearance check.
	return c.waitAccessPointGone(ctx, filesystemID, accesspointID)
}

func (c *agenticfsController) waitAccessPointActive(ctx context.Context, filesystemId, accesspointId string, server *string) error {
	return c.pollAccessPoint(ctx, filesystemId, accesspointId, apPoll{
		waiting: "to become Active",
		observe: func(ap *sdk.DescribeAccessPointResponseBodyAccessPoint) {
			if server != nil && *server == "" {
				*server = tea.StringValue(ap.DomainName)
			}
		},
		done: func(ap *sdk.DescribeAccessPointResponseBodyAccessPoint, notFound bool) bool {
			return !notFound && ap != nil && strings.ToLower(tea.StringValue(ap.Status)) == accessPointStatusActive
		},
		timeout: func(lastStatus string) error {
			return status.Errorf(codes.DeadlineExceeded,
				"timed out waiting for accesspoint %s to become Active after %s (last status: %s)",
				accesspointId, c.apPollTimeout, lastStatus)
		},
	})
}

func (c *agenticfsController) waitAccessPointGone(ctx context.Context, filesystemId, accesspointId string) error {
	return c.pollAccessPoint(ctx, filesystemId, accesspointId, apPoll{
		waiting: "to be deleted",
		done: func(_ *sdk.DescribeAccessPointResponseBodyAccessPoint, notFound bool) bool {
			return notFound
		},
		timeout: func(lastStatus string) error {
			return status.Errorf(codes.Aborted,
				"accesspoint %s is still being deleted after %s (last status: %s); DeleteAgenticSpace requires all accesspoints of the space to be detached, retrying",
				accesspointId, c.apPollTimeout, lastStatus)
		},
	})
}

type apPoll struct {
	waiting string
	observe func(ap *sdk.DescribeAccessPointResponseBodyAccessPoint)
	done    func(ap *sdk.DescribeAccessPointResponseBodyAccessPoint, notFound bool) bool
	timeout func(lastStatus string) error
}

// Clock injection follows disk/waitstatus. Preserve polling policy: observe first,
// accept the target state before checking the poll budget, and check cancellation
// before every API call. This deadline does not interrupt SDK credential resolution.
func (c *agenticfsController) pollAccessPoint(ctx context.Context, filesystemId, accesspointId string, p apPoll) error {
	logger := klog.FromContext(ctx)
	deadline := c.clock.Now().Add(c.apPollTimeout)
	timer := c.clock.NewTimer(c.apPollInterval)
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
		if c.clock.Now().After(deadline) {
			return p.timeout(lastStatus)
		}
		logger.V(4).Info("accesspoint not in the target state yet, retrying",
			"accesspointId", accesspointId, "waiting", p.waiting, "status", lastStatus)
		timer.Reset(c.apPollInterval)
		select {
		case <-ctx.Done():
			return status.Errorf(codes.DeadlineExceeded, "waiting for accesspoint %s %s: %v", accesspointId, p.waiting, ctx.Err())
		case <-timer.C():
		}
	}
}
