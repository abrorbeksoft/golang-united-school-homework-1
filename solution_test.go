package solution

import (
	"strings"
	"testing"
)

func TestGetMessage(t *testing.T) {
	got := GetMessage()
	want := string([]byte{72, 101, 108, 108, 111, 32, 58, 119, 111, 114, 100, 95, 109, 97, 112, 58, 33})
	if !strings.EqualFold(got, want) {
		t.Errorf("Unexpected result:\n\tExpected: %q\n\tGot: %q", want, got)
	}

}
