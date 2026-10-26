//go:build !windows

// Copyright 2026 The Kubernetes Authors.
// SPDX-License-Identifier: Apache-2.0

package nas

import (
	"context"
	"testing"
	"time"

	sdk "github.com/alibabacloud-go/nas-20170626/v4/client"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestAgenticfsAccessPointActiveStatusCasing(t *testing.T) {
	for _, state := range []accessPointStatus{"active", "Active", "ACTIVE"} {
		t.Run(string(state), func(t *testing.T) {
			ctrl := &agenticfsController{
				clock: newAgenticfsPollClock(), apPollInterval: time.Second, apPollTimeout: -time.Second,
				nasClient: agenticfsDescribeClient{describe: func(context.Context, string, string) (*sdk.DescribeAccessPointResponse, error) {
					return agenticfsDescribeResponse(state, "ap.example.com"), nil
				}},
			}
			server := ""
			require.NoError(t, ctrl.waitAccessPointActive(t.Context(), "fs", "ap", &server))
			require.Equal(t, "ap.example.com", server)
		})
	}
}

func TestAgenticfsAccessPointSelectionStatusCasing(t *testing.T) {
	for _, state := range []accessPointStatus{"active", "Active", "inactive", "Inactive", "deleting", "Deleting"} {
		t.Run(string(state), func(t *testing.T) {
			fake := newFakeNasClientV2()
			fake.listPages = []*sdk.ListAccessPointsResponseBody{apPage(apItem(testAgenticFsAccessPointID, state, testAgenticFsAPDomain))}
			active := state == "active" || state == "Active"
			if active {
				fake.listPages[0].AccessPoints = append(apPage(apItem("ap-pending", "Pending", "pending.example.com")).AccessPoints, fake.listPages[0].AccessPoints...)
			}
			ctrl := newAgenticfsCtrl(t, fake)
			id, _, err := ctrl.findReusableAccessPoint(t.Context(), "fs", testAgenticFsAgenticSpaceID)
			if active {
				require.NoError(t, err)
				require.Equal(t, testAgenticFsAccessPointID, id)
			} else {
				require.Equal(t, codes.Aborted, status.Code(err))
				require.Empty(t, id)
			}
			require.Empty(t, fake.createAccessPointReqs)
		})
	}
}
