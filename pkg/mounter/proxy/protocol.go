package proxy

import (
	"errors"
)

const (
	MaxMsgSize = 1 << 11
	MessageEnd = '\n'
)

type Method string

const (
	Mount Method = "mount"
)

type Header struct {
	Method Method `json:"method,omitempty"`
}

type Request struct {
	Header Header `json:"header"`
	Body   any    `json:"body,omitempty"`
}

type Response struct {
	Error string `json:"error,omitempty"`
}

func (r *Response) ToError() error {
	if r.Error == "" {
		return nil
	}
	return errors.New(r.Error)
}

type MountRequest struct {
	Source     string            `json:"source,omitempty"`
	Target     string            `json:"target,omitempty"`
	Fstype     string            `json:"fstype,omitempty"`
	Recovery   bool              `json:"recovery,omitempty"`
	Options    []string          `json:"options,omitempty"`
	MountFlags []string          `json:"mountFlags,omitempty"`
	Secrets    map[string]string `json:"secrets,omitempty"`
}
