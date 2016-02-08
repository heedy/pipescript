package datetime

import (
	"testing"
	"time"

	"github.com/connectordb/pipescript"
)

func TestYearMonth(t *testing.T) {
	Register()

	pipescript.TestCase{
		Pipescript: "yearmonth('lol')",
		ParseError: true,
	}.Run(t)
	pipescript.TestCase{
		Pipescript: "yearmonth",
		Input: []pipescript.Datapoint{
			{float64(time.Now().Unix()), 1},
		},
		Output: []pipescript.Datapoint{
			{float64(time.Now().Unix()), time.Now().Month().String()},
		},
	}.Run(t)
}
