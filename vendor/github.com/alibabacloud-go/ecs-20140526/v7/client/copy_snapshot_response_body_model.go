// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopySnapshotResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CopySnapshotResponseBody
	GetRequestId() *string
	SetSnapshotId(v string) *CopySnapshotResponseBody
	GetSnapshotId() *string
}

type CopySnapshotResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// C8B26B44-0189-443E-9816-D951F596****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the new snapshot.
	//
	// example:
	//
	// s-bp17441ohwka0yui****
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
}

func (s CopySnapshotResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CopySnapshotResponseBody) GoString() string {
	return s.String()
}

func (s *CopySnapshotResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CopySnapshotResponseBody) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *CopySnapshotResponseBody) SetRequestId(v string) *CopySnapshotResponseBody {
	s.RequestId = &v
	return s
}

func (s *CopySnapshotResponseBody) SetSnapshotId(v string) *CopySnapshotResponseBody {
	s.SnapshotId = &v
	return s
}

func (s *CopySnapshotResponseBody) Validate() error {
	return dara.Validate(s)
}
