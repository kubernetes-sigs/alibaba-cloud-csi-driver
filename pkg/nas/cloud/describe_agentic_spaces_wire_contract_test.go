package cloud

import (
	"context"
	"net/http"
	"testing"

	sdk "github.com/alibabacloud-go/nas-20170626/v4/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/wrap"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/time/rate"
)

func TestDescribeAgenticSpacesWireContract(t *testing.T) {
	body := `{"RequestId":"spaces-request","NextToken":"page-3","AgenticSpaces":{"AgenticSpace":[{"FileSystemId":"fs-test","AgenticSpaceId":"agentic-test","FileSystemPath":"/wanted/"}]}}`
	sdkClient, capture := newWireCaptureNasClient(t, http.StatusOK, body)
	client := &NasClientV2{client: sdkClient, limiter: rate.NewLimiter(rate.Inf, 1)}
	req := &sdk.DescribeAgenticSpacesRequest{
		FileSystemId: tea.String("fs-test"), MaxResults: tea.Int64(100), NextToken: tea.String("page-2"),
	}

	resp, err := client.DescribeAgenticSpaces(t.Context(), req)
	require.NoError(t, err)
	call := capture.last(t)
	assert.Equal(t, "DescribeAgenticSpaces", call.Action)
	assert.Equal(t, "2017-06-26", call.Version)
	assert.Equal(t, "fs-test", call.Query.Get("FileSystemId"))
	assert.Equal(t, "100", call.Query.Get("MaxResults"))
	assert.Equal(t, "page-2", call.Query.Get("NextToken"))
	assert.ElementsMatch(t, []string{"FileSystemId", "MaxResults", "NextToken", "RegionId"}, queryKeys(call.Query))
	require.NotNil(t, resp.Body)
	require.NotNil(t, resp.Body.AgenticSpaces)
	require.Len(t, resp.Body.AgenticSpaces.AgenticSpace, 1)
	space := resp.Body.AgenticSpaces.AgenticSpace[0]
	assert.Equal(t, "fs-test", tea.StringValue(space.FileSystemId))
	assert.Equal(t, "agentic-test", tea.StringValue(space.AgenticSpaceId))
	assert.Equal(t, "/wanted/", tea.StringValue(space.FileSystemPath))
	assert.Equal(t, "page-3", tea.StringValue(resp.Body.NextToken))
}

func TestDescribeAgenticSpacesPreservesPermissionError(t *testing.T) {
	sdkClient, _ := newWireCaptureNasClient(t, http.StatusForbidden, `{"Code":"Forbidden.RAM","Message":"not authorized","RequestId":"denied-request"}`)
	client := &NasClientV2{client: sdkClient, limiter: rate.NewLimiter(rate.Inf, 1)}
	_, err := client.DescribeAgenticSpaces(t.Context(), &sdk.DescribeAgenticSpacesRequest{FileSystemId: tea.String("fs-test")})
	require.ErrorIs(t, err, wrap.ErrorCode("Forbidden.RAM"))
}

func TestDescribeAgenticSpacesHonorsCancellation(t *testing.T) {
	sdkClient, capture := newWireCaptureNasClient(t, http.StatusOK, `{}`)
	client := &NasClientV2{client: sdkClient, limiter: rate.NewLimiter(rate.Inf, 1)}
	ctx, cancel := context.WithCancel(t.Context())
	cancel()
	_, err := client.DescribeAgenticSpaces(ctx, &sdk.DescribeAgenticSpacesRequest{FileSystemId: tea.String("fs-test")})
	require.ErrorIs(t, err, context.Canceled)
	assert.Empty(t, capture.all())
}
