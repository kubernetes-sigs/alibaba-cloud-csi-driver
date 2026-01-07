// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSnapshotGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateSnapshotGroupResponseBody
	GetRequestId() *string
	SetSnapshotGroupId(v string) *CreateSnapshotGroupResponseBody
	GetSnapshotGroupId() *string
}

type CreateSnapshotGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 01ABBD93-1ABB-4D92-B496-1A3D20EC0697
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the snapshot-consistent group.
	//
	// example:
	//
	// ssg-j6ciyh3k52qp7ovm****
	SnapshotGroupId *string `json:"SnapshotGroupId,omitempty" xml:"SnapshotGroupId,omitempty"`
}

func (s CreateSnapshotGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSnapshotGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSnapshotGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSnapshotGroupResponseBody) GetSnapshotGroupId() *string {
	return s.SnapshotGroupId
}

func (s *CreateSnapshotGroupResponseBody) SetRequestId(v string) *CreateSnapshotGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSnapshotGroupResponseBody) SetSnapshotGroupId(v string) *CreateSnapshotGroupResponseBody {
	s.SnapshotGroupId = &v
	return s
}

func (s *CreateSnapshotGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
