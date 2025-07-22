package server

import (
	"io"
	"sync/atomic"
)

type switchableWriter struct {
	target atomic.Pointer[io.Writer]
}

func NewSwitchableWriter(initial io.Writer) *switchableWriter {
	sw := &switchableWriter{}
	sw.target.Store(&initial)
	return sw
}

func (w *switchableWriter) Write(p []byte) (n int, err error) {
	target := *w.target.Load()
	return target.Write(p)
}

func (w *switchableWriter) SwitchTarget(newTarget io.Writer) {
	w.target.Store(&newTarget)
}
