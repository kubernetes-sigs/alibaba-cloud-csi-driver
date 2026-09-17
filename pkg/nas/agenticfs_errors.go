//go:build !windows

package nas

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/wrap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"k8s.io/klog/v2"
)

const (
	resourceCreatedLogPrefix  = "agenticfs-resource-created"
	retainedForRetryLogPrefix = "agenticfs-resource-retained-for-retry"
	// Historical prefix: requests reconciliation, not proof that a resource is orphaned.
	orphanLogPrefix = "agenticfs-orphan-resource"
)

// A v2 SDK failure unwraps to an error exposing the OpenAPI code.
type aliErrorCode interface {
	ErrorCode() string
}

func apiErrorCode(err error) string {
	var e aliErrorCode
	if errors.As(err, &e) {
		return strings.TrimSpace(e.ErrorCode())
	}
	return ""
}

// Match whole codes, not a .NotFound suffix: other missing resources do not
// establish that the requested filesystem, space or accesspoint is absent.
func isNotFoundError(err error) bool {
	if err == nil {
		return false
	}
	if errors.Is(err, wrap.ErrorCode("NotFound")) {
		return true
	}
	switch strings.ToLower(apiErrorCode(err)) {
	case "notfound",
		"invalidagenticspace.notfound",
		"invalidagenticspaceid.notfound",
		"invalidaccesspoint.notfound",
		"invalidaccesspointid.notfound",
		"invalidfilesystem.notfound",
		"invalidfilesystemid.notfound":
		return true
	}
	return false
}

// Only for space-scoped Get/Delete calls. Fileset/Fset aliases remain local to
// this predicate; a missing accesspoint never proves that its space is absent.
func isAgenticSpaceNotFoundError(err error) bool {
	switch strings.ToLower(apiErrorCode(err)) {
	case "invalidaccesspoint.notfound", "invalidaccesspointid.notfound":
		return false
	case "invalidfilesetid.notfound", "invalidfsetid.notfound":
		return true
	}
	return isNotFoundError(err)
}

var permanentAPIErrorCodes = []string{
	"invalidfilesystempath.invalidcharacters",
	"invalidregionid.notfound",
	"invalidzoneid.notfound",
	"forbidden.regiondisabled",
	"forbidden.region",
}

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
	if (strings.Contains(code, "region") || strings.Contains(code, "zone")) && strings.HasSuffix(code, "notsupported") {
		return true
	}
	if strings.HasPrefix(code, "unsupported") && (strings.Contains(code, "region") || strings.Contains(code, "zone")) {
		return true
	}
	return isNotFoundError(err)
}

// Preserve existing gRPC statuses; unknown cloud errors remain retryable.
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

// Reporting only: error classification affects diagnostics, never resource deletion.
// IDs describe what this call observed, not the full history of the ClientToken.
func (c *agenticfsController) reportCreateVolumeFailure(logger klog.Logger, filesystemId, agenticSpaceId, fileSystemPath, accesspointId string, cause error) {
	logger = logger.WithValues(c.orphanLogFields(filesystemId, agenticSpaceId, fileSystemPath, accesspointId)...).
		WithValues("cause", fmt.Sprint(cause), "code", status.Code(cause))
	if agenticSpaceId == "" {
		logger = logger.WithValues("resourceState", "unknown")
	}
	switch status.Code(cause) {
	case codes.InvalidArgument, codes.OutOfRange, codes.FailedPrecondition,
		codes.PermissionDenied, codes.Unauthenticated, codes.Unimplemented:
		logger.Error(cause, orphanLogPrefix+": provisioning failed; reconcile resources before cleanup",
			"reason", "terminal failure; no resources deleted, missing IDs do not prove absence")
	default:
		logger.Info(retainedForRetryLogPrefix+": resources retained for retry - NOT an orphan",
			"reason", "retryable or unclassified failure; replay ClientToken and rediscover accesspoints")
	}
}

func (c *agenticfsController) orphanLogFields(filesystemId, agenticSpaceId, fileSystemPath, accesspointId string) []any {
	return []any{
		"fileSystemId", filesystemId,
		"agenticSpaceId", agenticSpaceId,
		"fileSystemPath", fileSystemPath,
		"accesspointId", accesspointId,
		"region", c.region,
		"volumeHandle", strings.TrimSuffix(strings.TrimPrefix(fileSystemPath, "/"), "/"),
	}
}
