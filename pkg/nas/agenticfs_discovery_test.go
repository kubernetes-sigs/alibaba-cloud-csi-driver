//go:build !windows

package nas

import (
	"context"
	"strconv"
	"testing"

	sdk "github.com/alibabacloud-go/nas-20170626/v4/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (f *fakeNasClientV2) DescribeAgenticSpaces(_ context.Context, req *sdk.DescribeAgenticSpacesRequest) (*sdk.DescribeAgenticSpacesResponse, error) {
	f.describeAgenticSpacesReqs = append(f.describeAgenticSpacesReqs, req)
	f.callOrder = append(f.callOrder, "DescribeAgenticSpaces")
	if f.describeAgenticSpacesErr != nil {
		return nil, f.describeAgenticSpacesErr
	}
	if len(f.describeAgenticSpacesResponses) == 0 {
		return spacePage(""), nil
	}
	idx := min(len(f.describeAgenticSpacesReqs)-1, len(f.describeAgenticSpacesResponses)-1)
	return f.describeAgenticSpacesResponses[idx], nil
}

func spaceRecord(path, id string) *sdk.DescribeAgenticSpacesResponseBodyAgenticSpacesAgenticSpace {
	return &sdk.DescribeAgenticSpacesResponseBodyAgenticSpacesAgenticSpace{
		FileSystemId:   tea.String(testAgenticFsFilesystemID),
		FileSystemPath: tea.String(path), AgenticSpaceId: tea.String(id),
	}
}

func spacePage(next string, spaces ...*sdk.DescribeAgenticSpacesResponseBodyAgenticSpacesAgenticSpace) *sdk.DescribeAgenticSpacesResponse {
	return &sdk.DescribeAgenticSpacesResponse{Body: &sdk.DescribeAgenticSpacesResponseBody{
		AgenticSpaces: &sdk.DescribeAgenticSpacesResponseBodyAgenticSpaces{AgenticSpace: spaces},
		NextToken:     tea.String(next),
	}}
}

func TestAgenticfsDiscoveryFollowsNextToken(t *testing.T) {
	for _, emptyFirstPage := range []bool{false, true} {
		t.Run(strconv.FormatBool(emptyFirstPage), func(t *testing.T) {
			first := spacePage("second-page")
			if !emptyFirstPage {
				first = spacePage("second-page", spaceRecord("/unrelated/", "agentic-other"))
			}
			fake := newFakeNasClientV2()
			fake.describeAgenticSpacesResponses = []*sdk.DescribeAgenticSpacesResponse{
				first, spacePage("", spaceRecord("/wanted", testAgenticFsAgenticSpaceID)),
			}
			id, err := newAgenticfsCtrl(t, fake).discoverAgenticSpaceByPath(context.Background(), testAgenticFsFilesystemID, "/wanted/")
			require.NoError(t, err)
			assert.Equal(t, testAgenticFsAgenticSpaceID, id)
			require.Len(t, fake.describeAgenticSpacesReqs, 2)
			assert.Empty(t, tea.StringValue(fake.describeAgenticSpacesReqs[0].NextToken))
			assert.Equal(t, "second-page", tea.StringValue(fake.describeAgenticSpacesReqs[1].NextToken))
			for _, req := range fake.describeAgenticSpacesReqs {
				assert.Equal(t, testAgenticFsFilesystemID, tea.StringValue(req.FileSystemId))
				assert.EqualValues(t, 100, tea.Int64Value(req.MaxResults))
				assert.Empty(t, req.Filters)
			}
			assert.Empty(t, fake.listAccessPointsReqs)
		})
	}
}

func TestAgenticfsDiscoveryRejectsMalformedResponses(t *testing.T) {
	for _, shape := range []string{"nil response", "nil body", "nil space list", "nil record"} {
		t.Run(shape, func(t *testing.T) {
			fake := newFakeNasClientV2()
			switch shape {
			case "nil response":
				fake.describeAgenticSpacesResponses = []*sdk.DescribeAgenticSpacesResponse{nil}
			case "nil body":
				fake.describeAgenticSpacesResponses = []*sdk.DescribeAgenticSpacesResponse{{}}
			case "nil space list":
				fake.describeAgenticSpacesResponses = []*sdk.DescribeAgenticSpacesResponse{{Body: &sdk.DescribeAgenticSpacesResponseBody{}}}
			case "nil record":
				fake.describeAgenticSpacesResponses = []*sdk.DescribeAgenticSpacesResponse{spacePage("", nil)}
			}
			require.NotPanics(t, func() {
				id, err := newAgenticfsCtrl(t, fake).discoverAgenticSpaceByPath(context.Background(), testAgenticFsFilesystemID, "/wanted/")
				assert.Empty(t, id)
				assert.Equal(t, codes.Aborted, status.Code(err))
			})
		})
	}
}

func TestAgenticfsDiscoveryStopsOnPaginationCycle(t *testing.T) {
	fake := newFakeNasClientV2()
	fake.describeAgenticSpacesResponses = []*sdk.DescribeAgenticSpacesResponse{spacePage("cycle")}
	id, err := newAgenticfsCtrl(t, fake).discoverAgenticSpaceByPath(context.Background(), testAgenticFsFilesystemID, "/wanted/")
	assert.Empty(t, id)
	assert.Equal(t, codes.Aborted, status.Code(err))
	assert.Len(t, fake.describeAgenticSpacesReqs, 2)
}

func TestAgenticfsDiscoveryBoundsPagination(t *testing.T) {
	fake := newFakeNasClientV2()
	for i := range 50 {
		fake.describeAgenticSpacesResponses = append(fake.describeAgenticSpacesResponses, spacePage(strconv.Itoa(i+1)))
	}
	id, err := newAgenticfsCtrl(t, fake).discoverAgenticSpaceByPath(context.Background(), testAgenticFsFilesystemID, "/wanted/")
	assert.Empty(t, id)
	assert.Equal(t, codes.Aborted, status.Code(err))
	assert.Len(t, fake.describeAgenticSpacesReqs, 50)
}

func TestAgenticfsDiscoveryHonorsCancellation(t *testing.T) {
	fake := newFakeNasClientV2()
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	id, err := newAgenticfsCtrl(t, fake).discoverAgenticSpaceByPath(ctx, testAgenticFsFilesystemID, "/wanted/")
	assert.Empty(t, id)
	assert.ErrorIs(t, err, context.Canceled)
	assert.Empty(t, fake.listAccessPointsReqs)
	assert.Empty(t, fake.describeAgenticSpacesReqs)
}

func TestAgenticfsDiscoveryRejectsForeignFilesystem(t *testing.T) {
	fake := newFakeNasClientV2()
	foreign := spaceRecord("/wanted/", testAgenticFsAgenticSpaceID)
	foreign.FileSystemId = tea.String("another-filesystem")
	fake.describeAgenticSpacesResponses = []*sdk.DescribeAgenticSpacesResponse{spacePage("", foreign)}
	id, err := newAgenticfsCtrl(t, fake).discoverAgenticSpaceByPath(context.Background(), testAgenticFsFilesystemID, "/wanted/")
	assert.Empty(t, id)
	assert.Equal(t, codes.Aborted, status.Code(err))
}

func TestAgenticfsDiscoveryRejectsConflictingSpacesAcrossPages(t *testing.T) {
	fake := newFakeNasClientV2()
	fake.describeAgenticSpacesResponses = []*sdk.DescribeAgenticSpacesResponse{
		spacePage("next", spaceRecord("/wanted/", testAgenticFsAgenticSpaceID)),
		spacePage("", spaceRecord("/wanted/", "another-space")),
	}
	id, err := newAgenticfsCtrl(t, fake).discoverAgenticSpaceByPath(context.Background(), testAgenticFsFilesystemID, "/wanted/")
	assert.Empty(t, id)
	assert.Equal(t, codes.Aborted, status.Code(err))
	assert.Len(t, fake.describeAgenticSpacesReqs, 2)
}

func TestAgenticfsCreateVolumeRecoversSpaceWithoutAccessPoint(t *testing.T) {
	fake := newFakeNasClientV2()
	fake.createAgenticSpaceErr = status.Error(codes.InvalidArgument, "Path already used")
	fake.describeAgenticSpacesResponses = []*sdk.DescribeAgenticSpacesResponse{
		spacePage("", spaceRecord("/"+testAgenticFsPVName+"/", testAgenticFsAgenticSpaceID)),
	}
	resp, err := newAgenticfsCtrl(t, fake).CreateVolume(context.Background(), agenticfsCreateReq(testAgenticFsPVName, 20*GiB, nil))
	require.NoError(t, err)
	assert.Equal(t, testAgenticFsAgenticSpaceID, resp.Volume.VolumeContext[vcKeyAgenticSpaceId])
	require.Len(t, fake.createAccessPointReqs, 1)
	assert.Equal(t, testAgenticFsAgenticSpaceID, tea.StringValue(fake.createAccessPointReqs[0].AgenticSpaceId))
	assert.Equal(t, []string{"CreateAgenticSpace", "DescribeAgenticSpaces", "ListAccesspoints", "CreateAccesspoint", "DescribeAccesspoint"}, fake.callOrder)
	assert.Empty(t, fake.deleteAgenticSpaceReqs)
	assert.Empty(t, fake.deleteAccessPointIDs)
}

func TestAgenticfsDiscoveryDoesNotAcceptPartialMatch(t *testing.T) {
	fake := newFakeNasClientV2()
	fake.describeAgenticSpacesResponses = []*sdk.DescribeAgenticSpacesResponse{
		spacePage("next", spaceRecord("/wanted/", testAgenticFsAgenticSpaceID)), nil,
	}
	id, err := newAgenticfsCtrl(t, fake).discoverAgenticSpaceByPath(context.Background(), testAgenticFsFilesystemID, "/wanted/")
	assert.Empty(t, id)
	assert.Equal(t, codes.Aborted, status.Code(err))
	assert.Len(t, fake.describeAgenticSpacesReqs, 2)
}

func TestAgenticfsCreateVolumeDoesNotDiscoverWithoutPathConflict(t *testing.T) {
	for _, createErr := range []error{nil, status.Error(codes.Internal, "temporary creation failure")} {
		fake := newFakeNasClientV2()
		fake.createAgenticSpaceErr = createErr
		_, err := newAgenticfsCtrl(t, fake).CreateVolume(context.Background(), agenticfsCreateReq(testAgenticFsPVName, 20*GiB, nil))
		if createErr == nil {
			require.NoError(t, err)
		} else {
			require.Error(t, err)
			assert.Empty(t, fake.listAccessPointsReqs)
		}
		assert.Empty(t, fake.describeAgenticSpacesReqs)
	}
}
