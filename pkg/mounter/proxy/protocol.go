package proxy

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"k8s.io/klog/v2"
)

const (
	MaxMsgSize = 1 << 28 // 256MB
	MessageEnd = '\n'
)

type Method string

const (
	Mount Method = "mount"
	Ping  Method = "ping"
)

type Header struct {
	Method Method `json:"method,omitempty"`
}

type Request struct {
	Header Header `json:"header"`
	Body   any    `json:"body,omitempty"`
}

type Response struct {
	Seq   int64  `json:"seq,omitempty"`
	Error string `json:"error,omitempty"`
}

func (r *Response) ToError() error {
	if r.Error == "" {
		return nil
	}
	return errors.New(r.Error)
}

type MountRequest struct {
	Source      string            `json:"source,omitempty"`
	Target      string            `json:"target,omitempty"`
	Fstype      string            `json:"fstype,omitempty"`
	Options     []string          `json:"options,omitempty"`
	MountFlags  []string          `json:"mountFlags,omitempty"`
	Secrets     map[string]string `json:"secrets,omitempty"`
	MetricsPath string            `json:"metricsPath,omitempty"`
	VolumeID    string            `json:"volumeID,omitempty"`
}

func ReadMsg(r io.Reader, msg any) error {
	lr := io.LimitedReader{R: r, N: MaxMsgSize}
	dec := json.NewDecoder(&lr)
	err := dec.Decode(msg)
	if err != nil {
		if lr.N <= 0 {
			return fmt.Errorf("message too large")
		}
		return fmt.Errorf("read msg: %w", err)
	}

	var p [32]byte
	n, err := io.MultiReader(dec.Buffered(), &lr).Read(p[:])
	if err != nil {
		return fmt.Errorf("read msg end: %w", err)
	}
	if n == 0 || p[0] != MessageEnd {
		return errors.New("no message end after json")
	}
	if n > 1 {
		klog.V(1).InfoS("extra bytes after message end", "bytes", p[1:n])
	}
	return nil
}
