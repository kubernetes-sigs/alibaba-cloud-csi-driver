/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package fieldpath

import (
	"errors"
	"fmt"
	"strconv"

	"sigs.k8s.io/structured-merge-diff/v4/internal/builder"
	"sigs.k8s.io/structured-merge-diff/v4/value"
)

var ErrUnknownPathElementType = errors.New("unknown path element type")

const (
	// Field indicates that the content of this path element is a field's name
	peField byte = 'f'

	// Value indicates that the content of this path element is a field's value
	peValue byte = 'v'

	// Index indicates that the content of this path element is an index in an array
	peIndex byte = 'i'

	// Key indicates that the content of this path element is a key value map
	peKey byte = 'k'

	// Separator separates the type of a path element from the contents
	peSeparator byte = ':'
)

var (
	peFieldSepBytes = []byte{peField, peSeparator}
	peValueSepBytes = []byte{peValue, peSeparator}
	peIndexSepBytes = []byte{peIndex, peSeparator}
	peKeySepBytes   = []byte{peKey, peSeparator}
)

// DeserializePathElement parses a serialized path element
func DeserializePathElement(s string) (PathElement, error) {
	b := builder.StringToReadOnlyByteSlice(s)
	if len(b) < 2 {
		return PathElement{}, errors.New("key must be 2 characters long")
	}
	typeSep0, typeSep1, b := b[0], b[1], b[2:]
	if typeSep1 != peSeparator {
		return PathElement{}, fmt.Errorf("missing colon: %v", s)
	}
	switch typeSep0 {
	case peFieldSepBytes[0]:
		// Slice s rather than convert b, to save on
		// allocations.
		str := s[2:]
		return PathElement{
			FieldName: &str,
		}, nil
	case peValueSepBytes[0]:
		v, err := value.FromJSON(b)
		if err != nil {
			return PathElement{}, err
		}
		return PathElement{Value: &v}, nil
	case peKeySepBytes[0]:
		fields, err := value.FieldListFromJSON(b)
		if err != nil {
			return PathElement{}, err
		}
		fields.Sort()
		return PathElement{Key: &fields}, nil
	case peIndexSepBytes[0]:
		i, err := strconv.Atoi(s[2:])
		if err != nil {
			return PathElement{}, err
		}
		return PathElement{
			Index: &i,
		}, nil
	default:
		return PathElement{}, ErrUnknownPathElementType
	}
}

// SerializePathElement serializes a path element
func SerializePathElement(pe PathElement) (string, error) {
	buf := builder.JSONBuilder{}
	err := serializePathElementToWriter(&buf, pe)
	return buf.String(), err
}

func serializePathElementToWriter(w *builder.JSONBuilder, pe PathElement) error {
	switch {
	case pe.FieldName != nil:
		if _, err := w.Write(peFieldSepBytes); err != nil {
			return err
		}
		w.WriteString(*pe.FieldName)
	case pe.Key != nil:
		if _, err := w.Write(peKeySepBytes); err != nil {
			return err
		}
		w.WriteByte('{')
		nrKeys := len(*pe.Key)
		for i, f := range *pe.Key {
			if err := w.WriteJSON(f.Name); err != nil {
				return err
			}
			w.WriteByte(':')
			if err := w.WriteJSON(f.Value.Unstructured()); err != nil {
				return err
			}
			if i < nrKeys-1 {
				w.WriteByte(',')
			}
		}
		w.WriteByte('}')
	case pe.Value != nil:
		if _, err := w.Write(peValueSepBytes); err != nil {
			return err
		}
		if err := w.WriteJSON((*pe.Value).Unstructured()); err != nil {
			return err
		}
	case pe.Index != nil:
		if _, err := w.Write(peIndexSepBytes); err != nil {
			return err
		}
		w.WriteString(strconv.Itoa(*pe.Index))
	default:
		return errors.New("invalid PathElement")
	}
	return nil
}
