package testing

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/testing/protocmp"
)

func AssertProtoEqual(t *testing.T, expected, actual interface{}) {
	t.Helper()
	if diff := cmp.Diff(expected, actual, protocmp.Transform()); diff != "" {
		t.Errorf("unexpected difference:\n%v", diff)
	}
}
