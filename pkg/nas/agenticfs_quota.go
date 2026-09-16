//go:build !windows

package nas

import (
	"fmt"
	"math"
	"strconv"

	sdk "github.com/alibabacloud-go/nas-20170626/v4/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/container-storage-interface/spec/lib/go/csi"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/klog/v2"
)

const (
	defaultAgenticSpaceFileCountLimit int64 = 1000000
	minAgenticSpaceSizeLimit          int64 = 10 * GiB
	maxAgenticSpaceSizeLimit          int64 = 1099511627776000
	minAgenticSpaceFileCountLimit     int64 = 10000
	maxAgenticSpaceFileCountLimit     int64 = 100000000
)

// prepareAgenticSpaceExpansion validates the cloud snapshot and builds a new
// request without mutating it. It performs no API calls: the controller reads
// once, evaluates this policy, and only then sends SetAgenticSpaceQuota.
func prepareAgenticSpaceExpansion(logger klog.Logger, filesystemID, agenticSpaceID string, sizeLimit int64, getResp *sdk.GetAgenticSpaceResponse) (*sdk.SetAgenticSpaceQuotaRequest, error) {
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
	// Preserve readable file-count limits; omission retains the existing API policy.
	fileCountLimit := tea.Int64Value(space.Quota.FileCountLimit)
	sendsFileCountLimit := space.Quota.FileCountLimit != nil && fileCountLimit != 0
	var fileCountLimitReq *int64
	if sendsFileCountLimit {
		fileCountLimitReq = tea.Int64(fileCountLimit)
	}

	if sizeLimit < currentSizeLimit {
		return nil, status.Errorf(codes.OutOfRange,
			"shrinking an AgenticSpace is not supported: requested %d bytes, current size limit is %d bytes", sizeLimit, currentSizeLimit)
	}
	if sizeLimit < spaceUsage {
		return nil, status.Errorf(codes.OutOfRange,
			"requested size limit %d bytes is below the %d bytes already used by agenticspace %s", sizeLimit, spaceUsage, agenticSpaceID)
	}
	if sendsFileCountLimit && fileCountLimit < fileCountUsage {
		logger.Info("WARNING: agenticspace file count limit is below its current usage; this expansion does not modify the file count limit, proceeding",
			"agenticSpaceId", agenticSpaceID, "fileCountLimit", fileCountLimit, "fileCountUsage", fileCountUsage)
	}
	if sendsFileCountLimit {
		if err := validateAgenticSpaceFileCountLimit(fileCountLimit); err != nil {
			logger.Info("WARNING: the file count limit read back from nas:GetAgenticSpace is outside the OpenAPI range; this expansion does not modify it, so the field is OMITTED from SetAgenticSpaceQuota and the size expansion proceeds",
				"agenticSpaceId", agenticSpaceID, "fileCountLimit", fileCountLimit,
				"range", fmt.Sprintf("[%d, %d]", minAgenticSpaceFileCountLimit, maxAgenticSpaceFileCountLimit),
				"validationError", err.Error())
			fileCountLimitReq = nil
		}
	}
	if err := validateAgenticSpaceSizeLimit(sizeLimit); err != nil {
		return nil, err
	}
	return &sdk.SetAgenticSpaceQuotaRequest{
		FileSystemId:   tea.String(filesystemID),
		AgenticSpaceId: tea.String(agenticSpaceID),
		SizeLimit:      tea.Int64(sizeLimit),
		FileCountLimit: fileCountLimitReq,
	}, nil
}

// The optional StorageClass value is a cap, or a fallback when capacity is absent.
// Check both the requested bytes and the rounded quota against it.
func computeAgenticSpaceSizeLimit(cr *csi.CapacityRange, sizeLimitParam string) (int64, error) {
	var cap int64
	if sizeLimitParam != "" {
		q, err := resource.ParseQuantity(sizeLimitParam)
		if err != nil {
			return 0, status.Errorf(codes.InvalidArgument, "invalid parameters.agenticSpaceSizeLimit %q: %v", sizeLimitParam, err)
		}
		cap = q.Value()
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
	sizeLimit, err := roundUpToGiBChecked(bytes, cr.GetLimitBytes())
	if err != nil {
		return 0, err
	}
	if cap > 0 && sizeLimit > cap {
		return 0, status.Errorf(codes.InvalidArgument,
			"capacity %d bytes rounded up to the 1 GiB boundary (%d bytes) exceeds the cap parameters.agenticSpaceSizeLimit=%q (%d bytes)",
			bytes, sizeLimit, sizeLimitParam, cap)
	}
	return sizeLimit, nil
}

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

func roundUpToGiBChecked(bytes, limitBytes int64) (int64, error) {
	if bytes > math.MaxInt64-GiB+1 {
		return 0, status.Errorf(codes.OutOfRange, "capacity %d bytes overflows the 1 GiB rounding", bytes)
	}
	sizeLimit := roundUpToGiB(bytes)
	if sizeLimit > maxAgenticSpaceSizeLimit {
		return 0, status.Errorf(codes.InvalidArgument,
			"capacity %d bytes exceeds the AgenticSpace maximum of %d bytes (1024000 GiB)", sizeLimit, maxAgenticSpaceSizeLimit)
	}
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

// Use only on creation: expansion treats an invalid cloud-echoed file limit as a
// warning instead of rejecting an unrelated size increase.
func validateAgenticSpaceQuota(sizeLimit, fileCountLimit int64) error {
	if err := validateAgenticSpaceSizeLimit(sizeLimit); err != nil {
		return err
	}
	return validateAgenticSpaceFileCountLimit(fileCountLimit)
}

func validateAgenticSpaceSizeLimit(sizeLimit int64) error {
	if sizeLimit < minAgenticSpaceSizeLimit {
		return status.Errorf(codes.InvalidArgument,
			"size limit %d bytes is below the minimum %d bytes (10 GiB)", sizeLimit, minAgenticSpaceSizeLimit)
	}
	return nil
}

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
