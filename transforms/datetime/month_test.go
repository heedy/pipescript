package datetime

import (
	"testing"

	"github.com/connectordb/pipescript"
)

func TestMonth(t *testing.T) {
	Register()

	pipescript.TestCase{
		Pipescript: "month('lol')",
		ParseError: true,
	}.Run(t)
	pipescript.TestCase{
		Pipescript: "month",
		Input: []pipescript.Datapoint{
			{float64(1454907600), 1},
		},
		Output: []pipescript.Datapoint{
			{1454907600, int64(554)},
		},
	}.Run(t)
}
