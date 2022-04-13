package solution

import (
	"strings"
	"testing"
)

func TestGetMessage(t *testing.T) {
	got := GetMessage()
	want := string([]byte{72, 101, 108, 108, 111, 32, 240, 159, 151, 186, 239, 184, 143, 32, 33})

	if !strings.EqualFold(want, got) {
		t.Errorf("Unexpected result! GOT %q want %s", got, want)
	}

}
