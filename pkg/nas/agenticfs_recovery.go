//go:build !windows

package nas

import (
	"context"
	"strings"

	sdk "github.com/alibabacloud-go/nas-20170626/v4/client"
	"github.com/alibabacloud-go/tea/tea"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"k8s.io/klog/v2"
)

// recoverExistingAgenticSpace handles the NAS API's non-idempotent
// CreateAgenticSpace: when the API returns "Path already used" instead of
// replaying the ClientToken, discover the existing space by path.
//
// This is a workaround for a known upstream API issue; remove it once NAS
// replays the ClientToken. Quota compatibility is not checked here because
// GetAgenticSpace does not return the size limit; callers that need a larger
// quota should use ExpandVolume after recovery.
func (c *agenticfsController) recoverExistingAgenticSpace(ctx context.Context, filesystemID, path string) (string, error) {
	klog.InfoS("CreateAgenticSpace: path already exists, discovering existing space",
		"fileSystemId", filesystemID, "path", path)
	return c.discoverAgenticSpaceByPath(ctx, filesystemID, path)
}

func (c *agenticfsController) discoverAgenticSpaceByPath(ctx context.Context, filesystemID, path string) (string, error) {
	nextToken := ""
	matchedID := ""
	seenTokens := make(map[string]struct{})
	for page := 1; page <= spaceListMaxPages; page++ {
		if err := ctx.Err(); err != nil {
			return "", err
		}
		req := &sdk.DescribeAgenticSpacesRequest{
			FileSystemId: tea.String(filesystemID),
			MaxResults:   tea.Int64(spaceListMaxResults),
		}
		if nextToken != "" {
			req.NextToken = tea.String(nextToken)
		}
		listResp, err := c.nasClient.DescribeAgenticSpaces(ctx, req)
		if err != nil {
			return "", apiStatusError("nas:DescribeAgenticSpaces failed while discovering existing AgenticSpace", err)
		}
		if listResp == nil || listResp.Body == nil || listResp.Body.AgenticSpaces == nil {
			return "", status.Errorf(codes.Aborted, "nas:DescribeAgenticSpaces: malformed response on discovery page %d; retry later", page)
		}
		for _, space := range listResp.Body.AgenticSpaces.AgenticSpace {
			if space == nil || tea.StringValue(space.FileSystemId) != filesystemID {
				return "", status.Errorf(codes.Aborted, "nas:DescribeAgenticSpaces: missing record or mismatched filesystem on discovery page %d; refusing to reuse", page)
			}
			if strings.TrimSuffix(tea.StringValue(space.FileSystemPath), "/") != strings.TrimSuffix(path, "/") {
				continue
			}
			spaceID := tea.StringValue(space.AgenticSpaceId)
			if spaceID == "" {
				return "", status.Errorf(codes.Aborted, "nas:DescribeAgenticSpaces: empty AgenticSpaceId for path %s; refusing to reuse", path)
			}
			if matchedID != "" && matchedID != spaceID {
				return "", status.Errorf(codes.Aborted, "nas:DescribeAgenticSpaces: conflicting AgenticSpaces for path %s; refusing to choose", path)
			}
			matchedID = spaceID
		}
		nextToken = tea.StringValue(listResp.Body.NextToken)
		if nextToken == "" {
			if matchedID == "" {
				return "", status.Errorf(codes.Aborted, "nas:CreateAgenticSpace reported Path already used for %s but no matching AgenticSpace found via DescribeAgenticSpaces; retry after the space becomes visible", path)
			}
			klog.FromContext(ctx).Info("Discovered existing AgenticSpace",
				"fileSystemId", filesystemID, "agenticSpaceId", matchedID, "path", path)
			return matchedID, nil
		}
		if _, repeated := seenTokens[nextToken]; repeated {
			return "", status.Error(codes.Aborted, "nas:DescribeAgenticSpaces: repeated pagination token while discovering existing AgenticSpace; retry later")
		}
		seenTokens[nextToken] = struct{}{}
	}
	return "", status.Errorf(codes.Aborted, "nas:DescribeAgenticSpaces: did not converge in %d pages while discovering AgenticSpace for path %s", spaceListMaxPages, path)
}
