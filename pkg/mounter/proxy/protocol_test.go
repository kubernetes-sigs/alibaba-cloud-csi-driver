package proxy

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReadMsg(t *testing.T) {
	req := Request{
		Header: Header{
			Method: "test",
		},
		Body: map[string]any{
			"key": map[string]any{
				"in_key": "value",
			},
		},
	}

	data, err := json.Marshal(&req)
	assert.NoError(t, err)

	var parsed Request
	err = ReadMsg(bytes.NewReader(append(data, '\n')), &parsed)
	assert.NoError(t, err)
	assert.Equal(t, req, parsed)
}

func TestReadMsgErrors(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr string // substring of expected error
	}{
		{
			name:    "invalid json",
			input:   `{bad}`,
			wantErr: "read msg",
		},
		{
			name:    "no message end after json",
			input:   `{"header":{"method":"test"}}`,
			wantErr: "read msg end: EOF",
		},
		{
			name:    "empty input",
			input:   "",
			wantErr: "read msg",
		},
		{
			name:    "only newline",
			input:   "\n",
			wantErr: "read msg",
		},
		{
			name:    "wrong delimiter after json",
			input:   `{"header":{"method":"test"}}X`,
			wantErr: "no message end",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var parsed Request
			err := ReadMsg(strings.NewReader(tt.input), &parsed)
			require.Error(t, err)
			assert.Contains(t, err.Error(), tt.wantErr)
		})
	}
}

func TestReadMsgExtraBytesAfterDelimiter(t *testing.T) {
	// JSON followed by \n and extra bytes — ReadMsg should succeed,
	// logging extra bytes at V(1).
	var parsed Request
	err := ReadMsg(strings.NewReader(`{"header":{"method":"test"}}`+"\nextra"), &parsed)
	require.NoError(t, err)
	assert.Equal(t, Header{Method: "test"}, parsed.Header)
}
