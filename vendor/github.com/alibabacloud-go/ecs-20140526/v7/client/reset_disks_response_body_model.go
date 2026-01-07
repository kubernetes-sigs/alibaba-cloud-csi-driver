// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetDisksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOperationProgressSet(v *ResetDisksResponseBodyOperationProgressSet) *ResetDisksResponseBody
	GetOperationProgressSet() *ResetDisksResponseBodyOperationProgressSet
	SetRequestId(v string) *ResetDisksResponseBody
	GetRequestId() *string
}

type ResetDisksResponseBody struct {
	// Details about the rollback operation.
	OperationProgressSet *ResetDisksResponseBodyOperationProgressSet `json:"OperationProgressSet,omitempty" xml:"OperationProgressSet,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 3D66C85C-AA97-4A00-B0ED-2D9A80FE782C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResetDisksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResetDisksResponseBody) GoString() string {
	return s.String()
}

func (s *ResetDisksResponseBody) GetOperationProgressSet() *ResetDisksResponseBodyOperationProgressSet {
	return s.OperationProgressSet
}

func (s *ResetDisksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResetDisksResponseBody) SetOperationProgressSet(v *ResetDisksResponseBodyOperationProgressSet) *ResetDisksResponseBody {
	s.OperationProgressSet = v
	return s
}

func (s *ResetDisksResponseBody) SetRequestId(v string) *ResetDisksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResetDisksResponseBody) Validate() error {
	if s.OperationProgressSet != nil {
		if err := s.OperationProgressSet.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ResetDisksResponseBodyOperationProgressSet struct {
	OperationProgress []*ResetDisksResponseBodyOperationProgressSetOperationProgress `json:"OperationProgress,omitempty" xml:"OperationProgress,omitempty" type:"Repeated"`
}

func (s ResetDisksResponseBodyOperationProgressSet) String() string {
	return dara.Prettify(s)
}

func (s ResetDisksResponseBodyOperationProgressSet) GoString() string {
	return s.String()
}

func (s *ResetDisksResponseBodyOperationProgressSet) GetOperationProgress() []*ResetDisksResponseBodyOperationProgressSetOperationProgress {
	return s.OperationProgress
}

func (s *ResetDisksResponseBodyOperationProgressSet) SetOperationProgress(v []*ResetDisksResponseBodyOperationProgressSetOperationProgress) *ResetDisksResponseBodyOperationProgressSet {
	s.OperationProgress = v
	return s
}

func (s *ResetDisksResponseBodyOperationProgressSet) Validate() error {
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

type ResetDisksResponseBodyOperationProgressSetOperationProgress struct {
	// The error code that is returned if the request failed. This parameter is empty if the request is successful.
	//
	// For information about error codes and error messages, see [Service error codes](https://error-center.alibabacloud.com/status/product/Ecs).
	//
	// example:
	//
	// 400
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message that is returned if the request failed. This parameter is empty if the request is successful.
	//
	// For information about error codes and error messages, see [Service error codes](https://error-center.alibabacloud.com/status/product/Ecs).
	//
	// example:
	//
	// testErrorMsg
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// Indicates whether the request is successful. If the request is successful, Success is returned. If the request failed, an error code and an error message are returned.
	//
	// example:
	//
	// Success
	OperationStatus *string `json:"OperationStatus,omitempty" xml:"OperationStatus,omitempty"`
	// Details about the resources.
	RelatedItemSet *ResetDisksResponseBodyOperationProgressSetOperationProgressRelatedItemSet `json:"RelatedItemSet,omitempty" xml:"RelatedItemSet,omitempty" type:"Struct"`
}

func (s ResetDisksResponseBodyOperationProgressSetOperationProgress) String() string {
	return dara.Prettify(s)
}

func (s ResetDisksResponseBodyOperationProgressSetOperationProgress) GoString() string {
	return s.String()
}

func (s *ResetDisksResponseBodyOperationProgressSetOperationProgress) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ResetDisksResponseBodyOperationProgressSetOperationProgress) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *ResetDisksResponseBodyOperationProgressSetOperationProgress) GetOperationStatus() *string {
	return s.OperationStatus
}

func (s *ResetDisksResponseBodyOperationProgressSetOperationProgress) GetRelatedItemSet() *ResetDisksResponseBodyOperationProgressSetOperationProgressRelatedItemSet {
	return s.RelatedItemSet
}

func (s *ResetDisksResponseBodyOperationProgressSetOperationProgress) SetErrorCode(v string) *ResetDisksResponseBodyOperationProgressSetOperationProgress {
	s.ErrorCode = &v
	return s
}

func (s *ResetDisksResponseBodyOperationProgressSetOperationProgress) SetErrorMsg(v string) *ResetDisksResponseBodyOperationProgressSetOperationProgress {
	s.ErrorMsg = &v
	return s
}

func (s *ResetDisksResponseBodyOperationProgressSetOperationProgress) SetOperationStatus(v string) *ResetDisksResponseBodyOperationProgressSetOperationProgress {
	s.OperationStatus = &v
	return s
}

func (s *ResetDisksResponseBodyOperationProgressSetOperationProgress) SetRelatedItemSet(v *ResetDisksResponseBodyOperationProgressSetOperationProgressRelatedItemSet) *ResetDisksResponseBodyOperationProgressSetOperationProgress {
	s.RelatedItemSet = v
	return s
}

func (s *ResetDisksResponseBodyOperationProgressSetOperationProgress) Validate() error {
	if s.RelatedItemSet != nil {
		if err := s.RelatedItemSet.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ResetDisksResponseBodyOperationProgressSetOperationProgressRelatedItemSet struct {
	RelatedItem []*ResetDisksResponseBodyOperationProgressSetOperationProgressRelatedItemSetRelatedItem `json:"RelatedItem,omitempty" xml:"RelatedItem,omitempty" type:"Repeated"`
}

func (s ResetDisksResponseBodyOperationProgressSetOperationProgressRelatedItemSet) String() string {
	return dara.Prettify(s)
}

func (s ResetDisksResponseBodyOperationProgressSetOperationProgressRelatedItemSet) GoString() string {
	return s.String()
}

func (s *ResetDisksResponseBodyOperationProgressSetOperationProgressRelatedItemSet) GetRelatedItem() []*ResetDisksResponseBodyOperationProgressSetOperationProgressRelatedItemSetRelatedItem {
	return s.RelatedItem
}

func (s *ResetDisksResponseBodyOperationProgressSetOperationProgressRelatedItemSet) SetRelatedItem(v []*ResetDisksResponseBodyOperationProgressSetOperationProgressRelatedItemSetRelatedItem) *ResetDisksResponseBodyOperationProgressSetOperationProgressRelatedItemSet {
	s.RelatedItem = v
	return s
}

func (s *ResetDisksResponseBodyOperationProgressSetOperationProgressRelatedItemSet) Validate() error {
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

type ResetDisksResponseBodyOperationProgressSetOperationProgressRelatedItemSetRelatedItem struct {
	// The resource name.
	//
	// example:
	//
	// SnapshotId
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The resource ID.
	//
	// example:
	//
	// s-j6cdofbycydvg7ey****
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ResetDisksResponseBodyOperationProgressSetOperationProgressRelatedItemSetRelatedItem) String() string {
	return dara.Prettify(s)
}

func (s ResetDisksResponseBodyOperationProgressSetOperationProgressRelatedItemSetRelatedItem) GoString() string {
	return s.String()
}

func (s *ResetDisksResponseBodyOperationProgressSetOperationProgressRelatedItemSetRelatedItem) GetName() *string {
	return s.Name
}

func (s *ResetDisksResponseBodyOperationProgressSetOperationProgressRelatedItemSetRelatedItem) GetValue() *string {
	return s.Value
}

func (s *ResetDisksResponseBodyOperationProgressSetOperationProgressRelatedItemSetRelatedItem) SetName(v string) *ResetDisksResponseBodyOperationProgressSetOperationProgressRelatedItemSetRelatedItem {
	s.Name = &v
	return s
}

func (s *ResetDisksResponseBodyOperationProgressSetOperationProgressRelatedItemSetRelatedItem) SetValue(v string) *ResetDisksResponseBodyOperationProgressSetOperationProgressRelatedItemSetRelatedItem {
	s.Value = &v
	return s
}

func (s *ResetDisksResponseBodyOperationProgressSetOperationProgressRelatedItemSetRelatedItem) Validate() error {
	return dara.Validate(s)
}
