package solution

import (
	"github.com/kyokomi/emoji/v2"
)

func GetMessage() string {
	resp := emoji.Sprintf("Hello :world_map:!")
	return resp
}
