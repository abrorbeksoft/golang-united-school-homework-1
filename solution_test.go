package solution

import (
	"strings"
	"testing"
)

func TestGetMessage(t *testing.T) {
	got := GetMessage()
	want := string([]rune{72, 101, 108, 108, 111, 32, 128506, 65039, 32, 33})

	if strings.EqualFold(got, want) {
		t.Errorf("got %s want %s", got, want)
	}
}
