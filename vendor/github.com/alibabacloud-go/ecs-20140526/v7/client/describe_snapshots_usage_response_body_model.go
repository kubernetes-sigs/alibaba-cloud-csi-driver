// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSnapshotsUsageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeSnapshotsUsageResponseBody
	GetRequestId() *string
	SetSnapshotCount(v int32) *DescribeSnapshotsUsageResponseBody
	GetSnapshotCount() *int32
	SetSnapshotSize(v int64) *DescribeSnapshotsUsageResponseBody
	GetSnapshotSize() *int64
}

type DescribeSnapshotsUsageResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of snapshots stored in the current region.
	//
	// example:
	//
	// 5
	SnapshotCount *int32 `json:"SnapshotCount,omitempty" xml:"SnapshotCount,omitempty"`
	// The total size of snapshots stored in the current region. Unit: bytes.
	//
	// example:
	//
	// 122
	SnapshotSize *int64 `json:"SnapshotSize,omitempty" xml:"SnapshotSize,omitempty"`
}

func (s DescribeSnapshotsUsageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSnapshotsUsageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSnapshotsUsageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSnapshotsUsageResponseBody) GetSnapshotCount() *int32 {
	return s.SnapshotCount
}

func (s *DescribeSnapshotsUsageResponseBody) GetSnapshotSize() *int64 {
	return s.SnapshotSize
}

func (s *DescribeSnapshotsUsageResponseBody) SetRequestId(v string) *DescribeSnapshotsUsageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSnapshotsUsageResponseBody) SetSnapshotCount(v int32) *DescribeSnapshotsUsageResponseBody {
	s.SnapshotCount = &v
	return s
}

func (s *DescribeSnapshotsUsageResponseBody) SetSnapshotSize(v int64) *DescribeSnapshotsUsageResponseBody {
	s.SnapshotSize = &v
	return s
}

func (s *DescribeSnapshotsUsageResponseBody) Validate() error {
	return dara.Validate(s)
}
