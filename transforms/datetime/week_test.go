package datetime

import (
	"testing"

	"github.com/connectordb/pipescript"
)

func TestWeek(t *testing.T) {
	Register()

	pipescript.TestCase{
		Pipescript: "week('lol')",
		ParseError: true,
	}.Run(t)
	pipescript.TestCase{
		Pipescript: "week('UTC')",
		Input: []pipescript.Datapoint{
			{float64(1454907600), 1},
			{1454821200, 1},
		},
		Output: []pipescript.Datapoint{
			{1454907600, int64(2405)},
			{1454821200, int64(2404)},
		},
	}.Run(t)
}
