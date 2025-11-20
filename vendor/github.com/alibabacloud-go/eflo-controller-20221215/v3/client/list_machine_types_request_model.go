// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMachineTypesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *ListMachineTypesRequest
	GetName() *string
}

type ListMachineTypesRequest struct {
	// The name of the instance type.
	//
	// example:
	//
	// efg1.nvga1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListMachineTypesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMachineTypesRequest) GoString() string {
	return s.String()
}

func (s *ListMachineTypesRequest) GetName() *string {
	return s.Name
}

func (s *ListMachineTypesRequest) SetName(v string) *ListMachineTypesRequest {
	s.Name = &v
	return s
}

func (s *ListMachineTypesRequest) Validate() error {
	return dara.Validate(s)
}
