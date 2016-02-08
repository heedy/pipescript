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
		Pipescript: "dayhour('UTC')",
		Input: []pipescript.Datapoint{
			{float64(time.Now().Unix()), 1},
		},
		Output: []pipescript.Datapoint{
			{float64(time.Now().Unix()), time.Now().UTC().Hour()},
		},
	}.Run(t)
}
