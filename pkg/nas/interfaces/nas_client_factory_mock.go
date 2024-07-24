package interfaces

import "github.com/golang/mock/gomock"

type MockNasClientFactory struct {
	V1Client *MockNasV1Interface
	V2Client *MockNasClientV2Interface
}

func NewMockNasClientFactory(ctrl *gomock.Controller) *MockNasClientFactory {
	return &MockNasClientFactory{
		V1Client: NewMockNasV1Interface(ctrl),
		V2Client: NewMockNasClientV2Interface(ctrl),
	}
}

func (n *MockNasClientFactory) V1(string) (NasV1Interface, error) {
	return n.V1Client, nil
}

func (n *MockNasClientFactory) V2(string) (NasClientV2Interface, error) {
	return n.V2Client, nil
}
