package builder

import (
	gojson "encoding/json"
	"fmt"
	"io"
	"reflect"
	"runtime"
	"unsafe"

	"sigs.k8s.io/json"
)

type parserState int

const (
	stateLookingForObj parserState = iota
	stateLookingForItem
	stateLookingForKeyValueSep
	stateLookingForItemSep
	stateLookingForValue
	stateEnd
)

type FastObjParser struct {
	input []byte
	pos   int

	state parserState
}

func NewFastObjParser(input []byte) FastObjParser {
	return FastObjParser{
		input: input,
		state: stateLookingForObj,
	}
}

var whitespace = [256]bool{
	' ':  true,
	'\t': true,
	'\n': true,
	'\r': true,
}

func isWhitespace(c byte) bool {
	return whitespace[c]
}

func (p *FastObjParser) getValue(startPos int) ([]byte, error) {
	foundRootValue := false
	isQuoted := false
	isEscaped := false
	level := 0
	i := startPos
Loop:
	for ; i < len(p.input); i++ {
		if isQuoted {
			// Skip escaped character
			if isEscaped {
				isEscaped = false
				continue
			}

			switch p.input[i] {
			case '\\':
				isEscaped = true
			case '"':
				isQuoted = false
			}

			continue
		}

		// Skip whitespace
		if isWhitespace(p.input[i]) {
			continue
		}

		// If we are at the top level and find the next object, we are done
		if level == 0 && foundRootValue {
			switch p.input[i] {
			case ',', '}', ']', ':', '{', '[':
				break Loop
			}
		}

		switch p.input[i] {
		// Keep track of the nesting level
		case '{':
			level++
		case '}':
			level--
		case '[':
			level++
		case ']':
			level--

		// Start of a string
		case '"':
			isQuoted = true
		}

		foundRootValue = true
	}

	if level != 0 {
		return nil, fmt.Errorf("expected '}' or ']' but reached end of input")
	}

	if isQuoted {
		return nil, fmt.Errorf("expected '\"' but reached end of input")
	}

	if !foundRootValue {
		return nil, fmt.Errorf("expected value but reached end of input")
	}

	return p.input[startPos:i], nil
}

func (p *FastObjParser) Parse() ([]byte, error) {
	for {
		if p.pos >= len(p.input) {
			return nil, io.EOF
		}

		// Skip whitespace
		if isWhitespace(p.input[p.pos]) {
			p.pos++
			continue
		}

		switch p.state {
		case stateLookingForObj:
			if p.input[p.pos] != '{' {
				return nil, fmt.Errorf("expected '{' at position %d", p.pos)
			}

			p.state = stateLookingForItem

		case stateLookingForItem:
			if p.input[p.pos] == '}' {
				p.state = stateEnd
				return nil, io.EOF
			}

			strSlice, err := p.getValue(p.pos)
			if err != nil {
				return nil, err
			}

			p.pos += len(strSlice)
			p.state = stateLookingForKeyValueSep
			return strSlice, nil

		case stateLookingForKeyValueSep:
			if p.input[p.pos] != ':' {
				return nil, fmt.Errorf("expected ':' at position %d", p.pos)
			}

			p.state = stateLookingForValue

		case stateLookingForValue:
			valueSlice, err := p.getValue(p.pos)
			if err != nil {
				return nil, err
			}

			p.pos += len(valueSlice)
			p.state = stateLookingForItemSep
			return valueSlice, nil

		case stateLookingForItemSep:
			if p.input[p.pos] == ',' {
				p.state = stateLookingForItem
			} else if p.input[p.pos] == '}' {
				p.state = stateEnd
			} else {
				return nil, fmt.Errorf("expected ',' or '}' at position %d", p.pos)
			}

		case stateEnd:
			return nil, io.EOF
		}

		p.pos++
	}
}

func UnmarshalString(input []byte) (string, error) {
	var v string
	// No need to enable case sensitivity or int preservation here, as we are only unmarshalling strings.
	if err := gojson.Unmarshal(input, (*string)(noescape(unsafe.Pointer(&v)))); err != nil {
		return "", err
	}

	runtime.KeepAlive(v)

	return v, nil
}

func UnmarshalInterface(input []byte) (interface{}, error) {
	var v interface{}
	if err := json.UnmarshalCaseSensitivePreserveInts(input, (*interface{})(noescape(unsafe.Pointer(&v)))); err != nil {
		return "", err
	}

	runtime.KeepAlive(v)

	return v, nil
}

// Create a read-only byte array from a string
func StringToReadOnlyByteSlice(s string) []byte {
	// Get StringHeader from string
	stringHeader := (*reflect.StringHeader)(unsafe.Pointer(&s))

	// Construct SliceHeader with capacity equal to the length
	sliceHeader := reflect.SliceHeader{Data: stringHeader.Data, Len: stringHeader.Len, Cap: stringHeader.Len}

	// Convert SliceHeader to a byte slice
	return *(*[]byte)(unsafe.Pointer(&sliceHeader))
}
