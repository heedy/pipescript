package misc

import (
	"testing"

	"github.com/heedy/pipescript"
)

func TestContains(t *testing.T) {
	Register()
	pipescript.TestCase{
		Pipescript: "domain",
		Input: []pipescript.Datapoint{
			{1, "http://google.com/hi"},
			{2, "random"},
			{3, 34},
		},
		Output: []pipescript.Datapoint{
			{1, "google.com"},
			{2, ""},
			{3, ""},
		},
		SecondaryInput: []pipescript.Datapoint{
			{1, "https://golang.org/pkg/net/url/#URL.EscapedPath"},
			{2, "https://www.heedy.io"},
		},
		SecondaryOutput: []pipescript.Datapoint{
			{1, "golang.org"},
			{2, "heedy.io"},
		},
	}.Run(t)
}
