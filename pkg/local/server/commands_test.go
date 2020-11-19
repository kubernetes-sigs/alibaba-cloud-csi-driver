package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStr2ASCII(t *testing.T) {
	testStrs := map[string]string{
		"test1s912": "1161011151161115912",
		"$est2LIA":  "361011151162767365",
		"%st3/$#":   "371151163473635",
	}
	for k, v := range testStrs {
		assert.Equal(t, v, str2ASCII(k))
	}
}