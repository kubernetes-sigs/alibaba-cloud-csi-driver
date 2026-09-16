//go:build !windows

package nas

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"testing"

	sdk "github.com/alibabacloud-go/nas-20170626/v4/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/wrap"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestAgenticfsDeleteVolumeVerifiesAmbiguousListNotFound(t *testing.T) {
	listErrors := []struct {
		name string
		err  error
	}{
		{"accesspoint", aliErr("InvalidAccessPoint.NotFound")},
		{"accesspointID", aliErr("InvalidAccessPointId.NotFound")},
		{"paddedAccesspoint", aliErr("\tInvalidAccessPointId.NotFound\n")},
		{"caseInsensitiveAccesspoint", aliErr("INVALIDACCESSPOINTID.NOTFOUND")},
		{"unqualified", aliErr("NotFound")},
		{"paddedUnqualified", aliErr("\tNotFound\n")},
		{"wrappedUnqualified", wrap.ErrorCode("NotFound")},
	}
	for _, listError := range listErrors {
		t.Run(listError.name, func(t *testing.T) {
			for _, tt := range []struct {
				name     string
				resp     *sdk.GetAgenticSpaceResponse
				err      error
				wantCode codes.Code
			}{
				{"spaceExists", getSpaceResp(10*GiB, 10000, 0, 0), nil, codes.Aborted},
				{"nilResponse", nil, nil, codes.Aborted},
				{"nilBody", &sdk.GetAgenticSpaceResponse{}, nil, codes.Aborted},
				{"missingSpaceBody", &sdk.GetAgenticSpaceResponse{Body: &sdk.GetAgenticSpaceResponseBody{}}, nil, codes.Aborted},
				{"spaceGone", nil, aliErr("InvalidAgenticSpaceId.NotFound"), codes.OK},
				{"documentedSpaceGone", nil, aliErr("InvalidAgenticSpace.NotFound"), codes.OK},
				{"paddedDocumentedSpaceGone", nil, aliErr("\tInvalidAgenticSpace.NotFound\n"), codes.OK},
				{"filesetAliasGone", nil, aliErr("InvalidFilesetId.NotFound"), codes.OK},
				{"fsetAliasGone", nil, aliErr("InvalidFsetId.NotFound"), codes.OK},
				{"filesystemGone", nil, aliErr("InvalidFileSystemId.NotFound"), codes.OK},
				{"scopedUnqualifiedNotFound", nil, wrap.ErrorCode("NotFound"), codes.OK},
				{"scopedPaddedNotFound", nil, aliErr("\tNotFound\n"), codes.OK},
				{"onlyAccesspointGone", nil, aliErr("InvalidAccessPointId.NotFound"), codes.Aborted},
				{"verificationFails", nil, errors.New("connection reset"), codes.Internal},
				{"verificationTimesOut", nil, context.DeadlineExceeded, codes.DeadlineExceeded},
				{"verificationDenied", nil, aliErr("Forbidden.RAM"), codes.InvalidArgument},
			} {
				t.Run(tt.name, func(t *testing.T) {
					fake := newDeleteFakeNasClientV2()
					fake.listAccessPointsErr = listError.err
					fake.getAgenticSpaceResp = tt.resp
					fake.getAgenticSpaceErr = tt.err
					ctrl := newAgenticfsCtrl(t, fake)

					resp, err := ctrl.DeleteVolume(context.Background(), agenticfsDeleteReq(), agenticfsDeletePV())
					require.Equal(t, tt.wantCode, status.Code(err), "%v", err)
					if tt.wantCode == codes.OK {
						require.NotNil(t, resp)
					} else {
						require.Nil(t, resp, "an unconfirmed deletion must not release the PV")
					}
					assert.Equal(t, []string{"ListAccesspoints", "GetAgenticSpace"}, fake.callOrder)
					require.Len(t, fake.getAgenticSpaceReqs, 1)
					assert.Equal(t, testAgenticFsFilesystemID, tea.StringValue(fake.getAgenticSpaceReqs[0].FileSystemId))
					assert.Equal(t, testAgenticFsAgenticSpaceID, tea.StringValue(fake.getAgenticSpaceReqs[0].AgenticSpaceId))
					assert.Empty(t, fake.deleteAccessPointIDs, "never act on an incomplete listing")
					assert.Empty(t, fake.deleteAgenticSpaceReqs)
				})
			}
		})
	}
}

func TestAgenticfsDeleteVolumeRetriesAfterAccesspointListNotFound(t *testing.T) {
	fake := newDeleteFakeNasClientV2()
	fake.listAccessPointsErr = aliErr("InvalidAccessPointId.NotFound")
	fake.getAgenticSpaceResp = getSpaceResp(10*GiB, 10000, 0, 0)
	ctrl := newAgenticfsCtrl(t, fake)

	resp, err := ctrl.DeleteVolume(context.Background(), agenticfsDeleteReq(), agenticfsDeletePV())
	require.Equal(t, codes.Aborted, status.Code(err))
	require.Nil(t, resp)

	// Once enumeration recovers, all accesspoints (not just the PV's ID) must
	// be removed before deleting the space.
	fake.listAccessPointsErr = nil
	fake.listPages = []*sdk.ListAccessPointsResponseBody{apPage(
		apItem("ap-orphan", accessPointStatusActive, "orphan.example.com"),
	)}
	resp, err = ctrl.DeleteVolume(context.Background(), agenticfsDeleteReq(), agenticfsDeletePV())
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, []string{testAgenticFsAccessPointID, "ap-orphan"}, fake.deleteAccessPointIDs)
	require.Len(t, fake.deleteAgenticSpaceReqs, 1)
	assert.Equal(t, "DeleteAgenticSpace", fake.callOrder[len(fake.callOrder)-1])
}

// Use the error codes documented by GetAgenticSpace/DeleteAgenticSpace, not
// just the legacy InvalidAgenticSpaceId spelling used in earlier fixtures.
func TestAgenticfsDeleteVolumeAcceptsDocumentedMissingSpace(t *testing.T) {
	for _, code := range []string{
		"InvalidAgenticSpace.NotFound",
		"InvalidFilesetId.NotFound",
		"InvalidFsetId.NotFound",
	} {
		for _, variant := range []struct {
			name string
			err  error
		}{
			{"exact", aliErr(code)},
			{"padded", aliErr("\t" + code + "\n")},
			{"upperCase", aliErr(strings.ToUpper(code))},
			{"wrapped", fmt.Errorf("outer: %w", aliErr(code))},
		} {
			t.Run(code+"/"+variant.name, func(t *testing.T) {
				fake := newDeleteFakeNasClientV2()
				fake.deleteAgenticSpaceErr = variant.err
				ctrl := newAgenticfsCtrl(t, fake)

				// Repeated deletion must succeed even when the space remains absent.
				for range 2 {
					resp, err := ctrl.DeleteVolume(context.Background(), agenticfsDeleteReq(), agenticfsDeletePV())
					require.NoError(t, err)
					require.NotNil(t, resp)
				}
				require.Len(t, fake.deleteAgenticSpaceReqs, 2)
				for _, req := range fake.deleteAgenticSpaceReqs {
					assert.Equal(t, testAgenticFsFilesystemID, tea.StringValue(req.FileSystemId))
					assert.Equal(t, testAgenticFsAgenticSpaceID, tea.StringValue(req.AgenticSpaceId))
				}
			})
		}
	}
}

func TestAgenticfsDeleteVolumeListNotFoundResourceScope(t *testing.T) {
	for _, tt := range []struct {
		code     string
		wantCode codes.Code
	}{
		{"InvalidAgenticSpace.NotFound", codes.OK},
		{"InvalidFilesetId.NotFound", codes.Internal},
		{"InvalidFsetId.NotFound", codes.Internal},
	} {
		t.Run(tt.code, func(t *testing.T) {
			fake := newDeleteFakeNasClientV2()
			fake.listAccessPointsErr = aliErr(tt.code)
			ctrl := newAgenticfsCtrl(t, fake)
			resp, err := ctrl.DeleteVolume(context.Background(), agenticfsDeleteReq(), agenticfsDeletePV())
			require.Equal(t, tt.wantCode, status.Code(err), "%v", err)
			if tt.wantCode == codes.OK {
				require.NotNil(t, resp)
			} else {
				require.Nil(t, resp, "an unrelated missing fileset must not release the PV")
			}
			assert.Equal(t, []string{"ListAccesspoints"}, fake.callOrder)
		})
	}
}

func TestIsAgenticSpaceNotFoundError(t *testing.T) {
	for _, tt := range []struct {
		name string
		err  error
		want bool
	}{
		{"nil", nil, false},
		{"documented", aliErr("InvalidAgenticSpace.NotFound"), true},
		{"legacy", aliErr("InvalidAgenticSpaceId.NotFound"), true},
		{"fileset", aliErr("InvalidFilesetId.NotFound"), true},
		{"fset", aliErr("InvalidFsetId.NotFound"), true},
		{"wrappedPaddedAlias", fmt.Errorf("outer: %w", aliErr("\tINVALIDFSETID.NOTFOUND\n")), true},
		{"unqualified", wrap.ErrorCode("NotFound"), true},
		{"filesystem", aliErr("InvalidFileSystem.NotFound"), true},
		{"accesspoint", aliErr("InvalidAccessPoint.NotFound"), false},
		{"accesspointID", aliErr("InvalidAccessPointId.NotFound"), false},
		{"quota", aliErr("Quota.NotFound"), false},
		{"unknownSuffix", aliErr("InvalidFilesetId.NotSupported"), false},
		{"messageOnly", errors.New("InvalidAgenticSpace.NotFound"), false},
	} {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, isAgenticSpaceNotFoundError(tt.err))
		})
	}
}

func TestAgenticfsDeleteVolumeSpaceDeleteRejectsAccesspointNotFound(t *testing.T) {
	for _, code := range []string{
		"InvalidAccessPoint.NotFound",
		"InvalidAccessPointId.NotFound",
		"\tInvalidAccessPointId.NotFound\n",
		"INVALIDACCESSPOINTID.NOTFOUND",
	} {
		t.Run(code, func(t *testing.T) {
			fake := newDeleteFakeNasClientV2()
			fake.deleteAgenticSpaceErr = aliErr(code)
			ctrl := newAgenticfsCtrl(t, fake)

			resp, err := ctrl.DeleteVolume(context.Background(), agenticfsDeleteReq(), agenticfsDeletePV())
			require.Equal(t, codes.Aborted, status.Code(err))
			require.Nil(t, resp)
			require.Len(t, fake.deleteAgenticSpaceReqs, 1)

			fake.deleteAgenticSpaceErr = nil
			resp, err = ctrl.DeleteVolume(context.Background(), agenticfsDeleteReq(), agenticfsDeletePV())
			require.NoError(t, err)
			require.NotNil(t, resp)
			assert.Len(t, fake.deleteAgenticSpaceReqs, 2)
		})
	}
}
