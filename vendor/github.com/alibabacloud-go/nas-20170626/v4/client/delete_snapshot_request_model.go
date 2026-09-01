// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSnapshotRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSnapshotId(v string) *DeleteSnapshotRequest
	GetSnapshotId() *string
}

type DeleteSnapshotRequest struct {
	// The snapshot ID. After a snapshot is successfully created on an Advanced Extreme NAS file system by calling [CreateSnapshot](https://www.alibabacloud.com/help/en/nas/developer-reference/api-nas-2017-06-26-createsnapshot), call [DescribeSnapshots](https://www.alibabacloud.com/help/en/nas/developer-reference/api-nas-2017-06-26-describesnapshots) (with FileSystemType set to extreme) to query the snapshot list and obtain the snapshot ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// s-extreme-snapsho****
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
}

func (s DeleteSnapshotRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSnapshotRequest) GoString() string {
	return s.String()
}

func (s *DeleteSnapshotRequest) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *DeleteSnapshotRequest) SetSnapshotId(v string) *DeleteSnapshotRequest {
	s.SnapshotId = &v
	return s
}

func (s *DeleteSnapshotRequest) Validate() error {
	return dara.Validate(s)
}
