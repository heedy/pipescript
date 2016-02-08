package datetime

import (
	"testing"

	"github.com/connectordb/pipescript"
)

func TestDay(t *testing.T) {
	Register()

	pipescript.TestCase{
		Pipescript: "day('lol')",
		ParseError: true,
	}.Run(t)
	pipescript.TestCase{
		Pipescript: "day('UTC')",
		Input: []pipescript.Datapoint{
			{float64(1454907600), 1},
			{1454821200, 1},
		},
		Output: []pipescript.Datapoint{
			{1454907600, int64(16839)},
			{1454821200, int64(16838)},
		},
	}.Run(t)
}
