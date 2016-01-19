package datetime

import (
	"testing"
	"time"

	"github.com/connectordb/pipescript"
)

func TestWeekday(t *testing.T) {
	Register()

	pipescript.TestCase{
		Pipescript: "weekday('lol')",
		ParseError: true,
	}.Run(t)
	pipescript.TestCase{
		Pipescript: "weekday",
		Input: []pipescript.Datapoint{
			{float64(time.Now().Unix()), 1},
		},
		Output: []pipescript.Datapoint{
			{float64(time.Now().Unix()), time.Now().Weekday().String()},
		},
	}.Run(t)
}
