package proto

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
)

func TestMarshalCycle(t *testing.T) {
	req := &ExpandVolumeRequest{
		Volume: "test",
	}
	data, err := proto.Marshal(req)
	assert.NoError(t, err)

	newReq := &ExpandVolumeRequest{}
	err = proto.Unmarshal(data, newReq)
	assert.NoError(t, err)

	assert.Equal(t, req.Volume, newReq.Volume)
}
