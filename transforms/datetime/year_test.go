package datetime

import (
	"testing"

	"github.com/connectordb/pipescript"
)

func TestYear(t *testing.T) {
	Register()

	pipescript.TestCase{
		Pipescript: "year('lol')",
		ParseError: true,
	}.Run(t)
	pipescript.TestCase{
		Pipescript: "year",
		Input: []pipescript.Datapoint{
			{1454821200, 1},
		},
		Output: []pipescript.Datapoint{
			{1454821200, int64(2016)},
		},
	}.Run(t)
}
