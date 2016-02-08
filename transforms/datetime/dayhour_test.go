package datetime

import (
	"testing"
	"time"

	"github.com/connectordb/pipescript"
)

func TestDayhour(t *testing.T) {
	Register()

	pipescript.TestCase{
		Pipescript: "dayhour('lol')",
		ParseError: true,
	}.Run(t)
	pipescript.TestCase{
		Pipescript: "dayhour",
		Input: []pipescript.Datapoint{
			{float64(time.Now().Unix()), 1},
		},
		Output: []pipescript.Datapoint{
			{float64(time.Now().Unix()), time.Now().Hour()},
		},
	}.Run(t)
}
