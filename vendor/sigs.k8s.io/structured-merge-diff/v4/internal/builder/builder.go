package builder

import (
	"bytes"
	gojson "encoding/json"
	"runtime"
	"unsafe"
)

type ReusableBuilder struct {
	JSONBuilder
}

func (r *ReusableBuilder) Reset() *JSONBuilder {
	r.JSONBuilder.Reset()
	return &r.JSONBuilder
}

func (r *ReusableBuilder) UnsafeString() string {
	b := r.Bytes()
	return *(*string)(unsafe.Pointer(&b))
}

type JSONBuilder struct {
	initialised bool
	bytes.Buffer
	gojson.Encoder
}

type noNewlineWriter struct {
	*bytes.Buffer
}

func (w noNewlineWriter) Write(p []byte) (n int, err error) {
	if len(p) > 0 && p[len(p)-1] == '\n' {
		p = p[:len(p)-1]
	}
	return w.Buffer.Write(p)
}

// noescape hides a pointer from escape analysis. It is the identity function
// but escape analysis doesn't think the output depends on the input.
// noescape is inlined and currently compiles down to zero instructions.
// USE CAREFULLY!
// This was copied from the runtime; see issues 23382 and 7921.
//
//go:nosplit
//go:nocheckptr
func noescape(p unsafe.Pointer) unsafe.Pointer {
	x := uintptr(p)
	return unsafe.Pointer(x ^ 0)
}

func (r *JSONBuilder) WriteJSON(v interface{}) error {
	if !r.initialised {
		r.Encoder = *gojson.NewEncoder(noNewlineWriter{&r.Buffer})
		r.initialised = true
	}

	if err := r.Encoder.Encode((*interface{})(noescape(unsafe.Pointer(&v)))); err != nil {
		return err
	}

	runtime.KeepAlive(v)

	return nil
}

func MarshalString(input string) ([]byte, error) {
	out, err := gojson.Marshal((*string)(noescape(unsafe.Pointer(&input))))
	if err != nil {
		return nil, err
	}

	runtime.KeepAlive(input)

	return out, nil
}

func MarshalInterface(input interface{}) ([]byte, error) {
	out, err := gojson.Marshal((*interface{})(noescape(unsafe.Pointer(&input))))
	if err != nil {
		return nil, err
	}

	runtime.KeepAlive(input)

	return out, nil
}
