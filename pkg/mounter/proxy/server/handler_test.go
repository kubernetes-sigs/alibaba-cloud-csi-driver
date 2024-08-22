package server

import (
	"encoding/json"
	"testing"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy"
	"github.com/stretchr/testify/assert"
)

func Test_parseRawRequest(t *testing.T) {
	req := &proxy.Request{
		Header: proxy.Header{
			Method: "test",
		},
		Body: map[string]any{
			"key": map[string]any{
				"in_key": "value",
			},
		},
	}

	data, err := json.Marshal(req)
	assert.NoError(t, err)

	raw, err := parseRawRequest(append(data, '\n'))
	assert.NoError(t, err)
	assert.NotNil(t, raw)

	parsed := &proxy.Request{
		Header: raw.Header,
		Body:   make(map[string]string),
	}

	err = json.Unmarshal(raw.Body, &parsed.Body)
	assert.NoError(t, err)

	assert.Equal(t, req, parsed)
}
