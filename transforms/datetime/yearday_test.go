package datetime

import (
	"testing"
	"time"

	"github.com/connectordb/pipescript"
)

func TestYearday(t *testing.T) {
	Register()

	pipescript.TestCase{
		Pipescript: "yearday('lol')",
		ParseError: true,
	}.Run(t)
	pipescript.TestCase{
		Pipescript: "yearday('UTC')",
		Input: []pipescript.Datapoint{
			{float64(time.Now().Unix()), 1},
		},
		Output: []pipescript.Datapoint{
			{float64(time.Now().Unix()), time.Now().YearDay()},
		},
	}.Run(t)
}
