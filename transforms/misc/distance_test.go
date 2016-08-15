package misc

import (
	"testing"

	"github.com/connectordb/pipescript"
)

func TestDistance(t *testing.T) {
	Register()
	pipescript.TestCase{
		Pipescript: "distance(40.424454, -86.911356)",
		Input: []pipescript.Datapoint{
			{1, map[string]interface{}{"latitude": 40.424095, "longitude": -86.907388}},
			{2, map[string]interface{}{"latitude": 40.425353, "longitude": -86.913997}},
		},
		Output: []pipescript.Datapoint{
			{1, float64(338.24967084131407)},
			{2, float64(244.88709902527467)},
		},
	}.Run(t)
}
