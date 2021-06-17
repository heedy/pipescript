package core

import (
	"testing"

	"github.com/heedy/pipescript"
)

func TestMap(t *testing.T) {
	Map.Register()
	I.Register()

	pipescript.TestCase{
		// This tests order of prescedence: ":" pipes are high prescedence, and will be executed first
		Pipescript: "map d > 5 (i+1)",
		Input: []pipescript.Datapoint{
			{Timestamp: 1, Data: 4},
			{Timestamp: 2, Data: 6},
			{Timestamp: 3, Data: 8},
		},
		Output: []pipescript.Datapoint{
			{Timestamp: 1, Duration: 2, Data: map[string]interface{}{"false": float64(1), "true": float64(2)}},
		},
	}.Run(t)

	pipescript.TestCase{
		// This tests order of prescedence: ":" pipes are high prescedence, and will be executed first
		Pipescript: "map(d > 5,i+1)",
		Input: []pipescript.Datapoint{
			{Timestamp: 1, Data: 4},
			{Timestamp: 2, Data: 6},
			{Timestamp: 3, Data: 8},
		},
		Output: []pipescript.Datapoint{
			{Timestamp: 1, Duration: 2, Data: map[string]interface{}{"false": float64(1), "true": float64(2)}},
		},
	}.Run(t)
}
