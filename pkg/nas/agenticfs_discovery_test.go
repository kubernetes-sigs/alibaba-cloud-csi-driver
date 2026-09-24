//go:build !windows

package nas

import (
	"context"

	sdk "github.com/alibabacloud-go/nas-20170626/v4/client"
	"github.com/alibabacloud-go/tea/tea"
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
