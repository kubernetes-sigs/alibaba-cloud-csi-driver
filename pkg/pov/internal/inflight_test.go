package internal

import (
	"testing"
)

type testRequest struct {
	volumeId string
	expResp  bool
	delete   bool
}

func TestInFlight(t *testing.T) {
	testCases := []struct {
		name     string
		requests []testRequest
	}{
		{
			name: "success normal",
			requests: []testRequest{
				{

					volumeId: "random-vol-name",
					expResp:  true,
				},
			},
		},
		{
			name: "success adding request with different volumeId",
			requests: []testRequest{
				{
					volumeId: "random-vol-foobar",
					expResp:  true,
				},
				{
					volumeId: "random-vol-name-foobar",
					expResp:  true,
				},
			},
		},
		{
			name: "failed adding request with same volumeId",
			requests: []testRequest{
				{
					volumeId: "random-vol-name-foobar",
					expResp:  true,
				},
				{
					volumeId: "random-vol-name-foobar",
					expResp:  false,
				},
			},
		},

		{
			name: "success add, delete, add copy",
			requests: []testRequest{
				{
					volumeId: "random-vol-name",
					expResp:  true,
				},
				{
					volumeId: "random-vol-name",
					expResp:  false,
					delete:   true,
				},
				{
					volumeId: "random-vol-name",
					expResp:  true,
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			db := NewInFlight()
			for _, r := range tc.requests {
				var resp bool
				if r.delete {
					db.Delete(r.volumeId)
				} else {
					resp = db.Insert(r.volumeId)
				}
				if r.expResp != resp {
					t.Fatalf("expected insert to be %+v, got %+v", r.expResp, resp)
				}
			}
		})

	}
}
