package solution

import (
	"fmt"
	"testing"

	emoji "github.com/kyokomi/emoji/v2"
)

func TestGetMessage(t *testing.T) {
	got := GetMessage()
	fmt.Println(got)
	want := emoji.Sprint("Hello :map:!")

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
