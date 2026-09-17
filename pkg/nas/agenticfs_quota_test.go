//go:build !windows

package nas

import (
	"encoding/json"
	"testing"

	sdk "github.com/alibabacloud-go/nas-20170626/v4/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"k8s.io/apimachinery/pkg/api/resource"
)

func TestPrepareAgenticSpaceExpansionPreservesCloudSnapshot(t *testing.T) {
	for _, tt := range []struct {
		name          string
		fileLimit     *int64
		fileUsage     int64
		wantFileLimit *int64
		wantWarning   bool
		wantOmitted   bool
	}{
		{"valid", tea.Int64(10000), 0, tea.Int64(10000), false, false},
		{"missing", nil, 0, nil, false, false},
		{"zero", tea.Int64(0), 0, nil, false, false},
		{"below minimum", tea.Int64(9999), 0, nil, true, true},
		{"above maximum", tea.Int64(maxAgenticSpaceFileCountLimit + 1), 0, nil, true, true},
		{"usage above unchanged limit", tea.Int64(10000), 10001, tea.Int64(10000), true, false},
	} {
		t.Run(tt.name, func(t *testing.T) {
			resp := getSpaceResp(20*GiB, 10000, 0, tt.fileUsage)
			resp.Body.AgenticSpace.Quota.FileCountLimit = tt.fileLimit
			before, err := json.Marshal(resp)
			require.NoError(t, err)
			logger, _ := newLogCapture(t)
			req, err := prepareAgenticSpaceExpansion(logger, "fs", "space", 30*GiB, resp)
			require.NoError(t, err)
			assert.Equal(t, &sdk.SetAgenticSpaceQuotaRequest{
				FileSystemId:   tea.String("fs"),
				AgenticSpaceId: tea.String("space"),
				SizeLimit:      tea.Int64(30 * GiB),
				FileCountLimit: tt.wantFileLimit,
			}, req)
			if req.FileCountLimit != nil {
				assert.NotSame(t, resp.Body.AgenticSpace.Quota.FileCountLimit, req.FileCountLimit)
				*req.FileCountLimit = 12345
			}
			after, err := json.Marshal(resp)
			require.NoError(t, err)
			assert.Equal(t, string(before), string(after), "request construction must not mutate the cloud snapshot")
			logs := logText(logger)
			if tt.wantWarning {
				assert.Contains(t, logs, "WARNING")
			} else {
				assert.NotContains(t, logs, "WARNING")
			}
			if tt.wantOmitted {
				assert.Contains(t, logs, "OMITTED")
			}
		})
	}
}

func TestPrepareAgenticSpaceExpansionRejectsIncompleteOrUnsafeInput(t *testing.T) {
	for _, tt := range []struct {
		name     string
		mutate   func(*sdk.GetAgenticSpaceResponse) *sdk.GetAgenticSpaceResponse
		size     int64
		wantCode codes.Code
	}{
		{"nil response", func(*sdk.GetAgenticSpaceResponse) *sdk.GetAgenticSpaceResponse { return nil }, 30 * GiB, codes.Internal},
		{"nil quota", func(r *sdk.GetAgenticSpaceResponse) *sdk.GetAgenticSpaceResponse {
			r.Body.AgenticSpace.Quota = nil
			return r
		}, 30 * GiB, codes.Internal},
		{"missing usage", func(r *sdk.GetAgenticSpaceResponse) *sdk.GetAgenticSpaceResponse {
			r.Body.AgenticSpace.SpaceUsage = nil
			return r
		}, 30 * GiB, codes.Internal},
		{"shrink", func(r *sdk.GetAgenticSpaceResponse) *sdk.GetAgenticSpaceResponse { return r }, 10 * GiB, codes.OutOfRange},
		{"below used bytes", func(r *sdk.GetAgenticSpaceResponse) *sdk.GetAgenticSpaceResponse {
			r.Body.AgenticSpace.SpaceUsage = tea.Int64(40 * GiB)
			return r
		}, 30 * GiB, codes.OutOfRange},
		{"below minimum size", func(r *sdk.GetAgenticSpaceResponse) *sdk.GetAgenticSpaceResponse {
			r.Body.AgenticSpace.Quota.SizeLimit = tea.Int64(GiB)
			return r
		}, 5 * GiB, codes.InvalidArgument},
	} {
		t.Run(tt.name, func(t *testing.T) {
			logger, _ := newLogCapture(t)
			req, err := prepareAgenticSpaceExpansion(logger, "fs", "space", tt.size, tt.mutate(getSpaceResp(20*GiB, 10000, 0, 0)))
			assert.Nil(t, req, "unsafe snapshots must never produce a write request")
			assert.Equal(t, tt.wantCode, status.Code(err))
		})
	}
}

// Check successful parsing against invariants independently of error wording.
// Seed cases run in the regular suite; fuzzing exercises rounding and cap edges.
func FuzzComputeAgenticSpaceSizeLimit(f *testing.F) {
	f.Add(int64(10*GiB), int64(0), "")
	f.Add(int64(10*GiB+1), int64(0), "10.5Gi")
	f.Add(int64(0), int64(20*GiB), "10Gi")
	f.Add(int64(-1), int64(0), "invalid")
	f.Add(int64(maxAgenticSpaceSizeLimit), int64(0), "1024000Gi")
	f.Fuzz(func(t *testing.T, required, limit int64, capParam string) {
		if len(capParam) > 128 {
			t.Skip()
		}
		got, err := computeAgenticSpaceSizeLimit(&csi.CapacityRange{RequiredBytes: required, LimitBytes: limit}, capParam)
		if err != nil {
			return
		}
		require.Positive(t, got)
		assert.Zero(t, got%GiB)
		assert.LessOrEqual(t, got, maxAgenticSpaceSizeLimit)
		if required > 0 {
			assert.GreaterOrEqual(t, got, required)
		}
		if limit > 0 {
			assert.LessOrEqual(t, got, limit)
		}
		if capParam != "" {
			q, err := resource.ParseQuantity(capParam)
			require.NoError(t, err)
			assert.LessOrEqual(t, got, q.Value())
		}
	})
}
