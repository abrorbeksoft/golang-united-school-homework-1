package solution

import (
	"testing"

	emoji "github.com/kyokomi/emoji"
)

func TestGetMessage(t *testing.T) {
	got := GetMessage()
	want := emoji.Sprint("Hello :word_map:!")

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
