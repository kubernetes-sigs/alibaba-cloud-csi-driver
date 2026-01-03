// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSnapshotGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOperationProgressSet(v *DeleteSnapshotGroupResponseBodyOperationProgressSet) *DeleteSnapshotGroupResponseBody
	GetOperationProgressSet() *DeleteSnapshotGroupResponseBodyOperationProgressSet
	SetRequestId(v string) *DeleteSnapshotGroupResponseBody
	GetRequestId() *string
}

type DeleteSnapshotGroupResponseBody struct {
	// Details about the delete operation.
	OperationProgressSet *DeleteSnapshotGroupResponseBodyOperationProgressSet `json:"OperationProgressSet,omitempty" xml:"OperationProgressSet,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 6EDE885A-FDC1-4FAE-BC44-6EACAEA6CC6E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSnapshotGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSnapshotGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSnapshotGroupResponseBody) GetOperationProgressSet() *DeleteSnapshotGroupResponseBodyOperationProgressSet {
	return s.OperationProgressSet
}

func (s *DeleteSnapshotGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSnapshotGroupResponseBody) SetOperationProgressSet(v *DeleteSnapshotGroupResponseBodyOperationProgressSet) *DeleteSnapshotGroupResponseBody {
	s.OperationProgressSet = v
	return s
}

func (s *DeleteSnapshotGroupResponseBody) SetRequestId(v string) *DeleteSnapshotGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSnapshotGroupResponseBody) Validate() error {
	if s.OperationProgressSet != nil {
		if err := s.OperationProgressSet.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteSnapshotGroupResponseBodyOperationProgressSet struct {
	OperationProgress []*DeleteSnapshotGroupResponseBodyOperationProgressSetOperationProgress `json:"OperationProgress,omitempty" xml:"OperationProgress,omitempty" type:"Repeated"`
}

func (s DeleteSnapshotGroupResponseBodyOperationProgressSet) String() string {
	return dara.Prettify(s)
}

func (s DeleteSnapshotGroupResponseBodyOperationProgressSet) GoString() string {
	return s.String()
}

func (s *DeleteSnapshotGroupResponseBodyOperationProgressSet) GetOperationProgress() []*DeleteSnapshotGroupResponseBodyOperationProgressSetOperationProgress {
	return s.OperationProgress
}

func (s *DeleteSnapshotGroupResponseBodyOperationProgressSet) SetOperationProgress(v []*DeleteSnapshotGroupResponseBodyOperationProgressSetOperationProgress) *DeleteSnapshotGroupResponseBodyOperationProgressSet {
	s.OperationProgress = v
	return s
}

func (s *DeleteSnapshotGroupResponseBodyOperationProgressSet) Validate() error {
	if s.OperationProgress != nil {
		for _, item := range s.OperationProgress {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DeleteSnapshotGroupResponseBodyOperationProgressSetOperationProgress struct {
	// The error code. This parameter is empty when the operation is successful.
	//
	// For information about error codes and error messages, visit the [API error center](https://error-center.aliyun.com/status/product/Ecs).
	//
	// example:
	//
	// 400
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message. This parameter is empty when the operation is successful.
	//
	// For information about error codes and error messages, visit the [API error center](https://error-center.aliyun.com/status/product/Ecs).
	//
	// example:
	//
	// testErrorMsg
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// Indicates whether the operation was successful. If the operation was successful, a value of Success is returned. If the operation failed, an error code and an error message are returned.
	//
	// example:
	//
	// Success
	OperationStatus *string `json:"OperationStatus,omitempty" xml:"OperationStatus,omitempty"`
	// Details about the resources.
	RelatedItemSet *DeleteSnapshotGroupResponseBodyOperationProgressSetOperationProgressRelatedItemSet `json:"RelatedItemSet,omitempty" xml:"RelatedItemSet,omitempty" type:"Struct"`
}

func (s DeleteSnapshotGroupResponseBodyOperationProgressSetOperationProgress) String() string {
	return dara.Prettify(s)
}

func (s DeleteSnapshotGroupResponseBodyOperationProgressSetOperationProgress) GoString() string {
	return s.String()
}

func (s *DeleteSnapshotGroupResponseBodyOperationProgressSetOperationProgress) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteSnapshotGroupResponseBodyOperationProgressSetOperationProgress) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *DeleteSnapshotGroupResponseBodyOperationProgressSetOperationProgress) GetOperationStatus() *string {
	return s.OperationStatus
}

func (s *DeleteSnapshotGroupResponseBodyOperationProgressSetOperationProgress) GetRelatedItemSet() *DeleteSnapshotGroupResponseBodyOperationProgressSetOperationProgressRelatedItemSet {
	return s.RelatedItemSet
}

func (s *DeleteSnapshotGroupResponseBodyOperationProgressSetOperationProgress) SetErrorCode(v string) *DeleteSnapshotGroupResponseBodyOperationProgressSetOperationProgress {
	s.ErrorCode = &v
	return s
}

func (s *DeleteSnapshotGroupResponseBodyOperationProgressSetOperationProgress) SetErrorMsg(v string) *DeleteSnapshotGroupResponseBodyOperationProgressSetOperationProgress {
	s.ErrorMsg = &v
	return s
}

func (s *DeleteSnapshotGroupResponseBodyOperationProgressSetOperationProgress) SetOperationStatus(v string) *DeleteSnapshotGroupResponseBodyOperationProgressSetOperationProgress {
	s.OperationStatus = &v
	return s
}

func (s *DeleteSnapshotGroupResponseBodyOperationProgressSetOperationProgress) SetRelatedItemSet(v *DeleteSnapshotGroupResponseBodyOperationProgressSetOperationProgressRelatedItemSet) *DeleteSnapshotGroupResponseBodyOperationProgressSetOperationProgress {
	s.RelatedItemSet = v
	return s
}

func (s *DeleteSnapshotGroupResponseBodyOperationProgressSetOperationProgress) Validate() error {
	if s.RelatedItemSet != nil {
		if err := s.RelatedItemSet.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteSnapshotGroupResponseBodyOperationProgressSetOperationProgressRelatedItemSet struct {
	RelatedItem []*DeleteSnapshotGroupResponseBodyOperationProgressSetOperationProgressRelatedItemSetRelatedItem `json:"RelatedItem,omitempty" xml:"RelatedItem,omitempty" type:"Repeated"`
}

func (s DeleteSnapshotGroupResponseBodyOperationProgressSetOperationProgressRelatedItemSet) String() string {
	return dara.Prettify(s)
}

func (s DeleteSnapshotGroupResponseBodyOperationProgressSetOperationProgressRelatedItemSet) GoString() string {
	return s.String()
}

func (s *DeleteSnapshotGroupResponseBodyOperationProgressSetOperationProgressRelatedItemSet) GetRelatedItem() []*DeleteSnapshotGroupResponseBodyOperationProgressSetOperationProgressRelatedItemSetRelatedItem {
	return s.RelatedItem
}

func (s *DeleteSnapshotGroupResponseBodyOperationProgressSetOperationProgressRelatedItemSet) SetRelatedItem(v []*DeleteSnapshotGroupResponseBodyOperationProgressSetOperationProgressRelatedItemSetRelatedItem) *DeleteSnapshotGroupResponseBodyOperationProgressSetOperationProgressRelatedItemSet {
	s.RelatedItem = v
	return s
}

func (s *DeleteSnapshotGroupResponseBodyOperationProgressSetOperationProgressRelatedItemSet) Validate() error {
	if s.RelatedItem != nil {
		for _, item := range s.RelatedItem {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DeleteSnapshotGroupResponseBodyOperationProgressSetOperationProgressRelatedItemSetRelatedItem struct {
	// The name of the resource.
	//
	// example:
	//
	// SnapshotId
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the resource.
	//
	// example:
	//
	// s-j6c9lpuyxo2uxxnx****
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DeleteSnapshotGroupResponseBodyOperationProgressSetOperationProgressRelatedItemSetRelatedItem) String() string {
	return dara.Prettify(s)
}

func (s DeleteSnapshotGroupResponseBodyOperationProgressSetOperationProgressRelatedItemSetRelatedItem) GoString() string {
	return s.String()
}

func (s *DeleteSnapshotGroupResponseBodyOperationProgressSetOperationProgressRelatedItemSetRelatedItem) GetName() *string {
	return s.Name
}

func (s *DeleteSnapshotGroupResponseBodyOperationProgressSetOperationProgressRelatedItemSetRelatedItem) GetValue() *string {
	return s.Value
}

func (s *DeleteSnapshotGroupResponseBodyOperationProgressSetOperationProgressRelatedItemSetRelatedItem) SetName(v string) *DeleteSnapshotGroupResponseBodyOperationProgressSetOperationProgressRelatedItemSetRelatedItem {
	s.Name = &v
	return s
}

func (s *DeleteSnapshotGroupResponseBodyOperationProgressSetOperationProgressRelatedItemSetRelatedItem) SetValue(v string) *DeleteSnapshotGroupResponseBodyOperationProgressSetOperationProgressRelatedItemSetRelatedItem {
	s.Value = &v
	return s
}

func (s *DeleteSnapshotGroupResponseBodyOperationProgressSetOperationProgressRelatedItemSetRelatedItem) Validate() error {
	return dara.Validate(s)
}
