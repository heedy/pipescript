package datetime

import (
	"testing"
	"time"

	"github.com/connectordb/pipescript"
)

func TestMonthday(t *testing.T) {
	Register()

	pipescript.TestCase{
		Pipescript: "monthday('lol')",
		ParseError: true,
	}.Run(t)
	pipescript.TestCase{
		Pipescript: "monthday",
		Input: []pipescript.Datapoint{
			{float64(time.Now().Unix()), 1},
		},
		Output: []pipescript.Datapoint{
			{float64(time.Now().Unix()), time.Now().Day()},
		},
	}.Run(t)
}
