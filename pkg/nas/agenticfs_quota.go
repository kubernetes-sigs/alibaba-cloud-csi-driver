//go:build !windows

package nas

import (
	"math"
	"strconv"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"k8s.io/apimachinery/pkg/api/resource"
)

const (
	defaultAgenticSpaceFileCountLimit int64 = 1000000
	minAgenticSpaceSizeLimit          int64 = 10 * GiB
	maxAgenticSpaceSizeLimit          int64 = 1099511627776000
	minAgenticSpaceFileCountLimit     int64 = 10000
	maxAgenticSpaceFileCountLimit     int64 = 100000000
)

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
