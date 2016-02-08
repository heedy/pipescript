package datetime

import (
	"testing"

	"github.com/connectordb/pipescript"
)

func TestHour(t *testing.T) {
	Register()

	pipescript.TestCase{
		Pipescript: "hour('lol')",
		ParseError: true,
	}.Run(t)
	pipescript.TestCase{
		Pipescript: "hour('UTC')",
		Input: []pipescript.Datapoint{
			{float64(1454907600), 1},
		},
		Output: []pipescript.Datapoint{
			{1454907600, int64(404141)}, // Takes time zone into account
		},
	}.Run(t)
}
