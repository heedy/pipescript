package datasets

import (
	"testing"

	"github.com/heedy/pipescript"
	"github.com/heedy/pipescript/transforms/numeric"
)

func TestTransformInterpolator(t *testing.T) {
	numeric.Count.Register()
	TestCase{
		Interpolator: "d", // Use a PipeScript
		Reference: []pipescript.Datapoint{
			{Timestamp: 0.1},
			{Timestamp: 1.3},
			{Timestamp: 2.01},
			{Timestamp: 3},
			{Timestamp: 4.1},
			{Timestamp: 4},
			{Timestamp: 6},
			{Timestamp: 8},
		},
		Stream: []pipescript.Datapoint{
			{Timestamp: 1, Data: "1"},
			{Timestamp: 2, Data: "2"},
			{Timestamp: 4, Data: "3"},
			{Timestamp: 4, Data: "4"},
			{Timestamp: 5, Data: "5"},
		},
		Output: []pipescript.Datapoint{
			{Timestamp: 0.1, Data: nil},
			{Timestamp: 1, Data: "1"},
			{Timestamp: 2, Data: "2"},
			{Timestamp: 3, Data: nil},
			{Timestamp: 4, Data: "4"},
			{Timestamp: 4, Data: nil},
			{Timestamp: 5, Data: "5"},
			{Timestamp: 8, Data: nil},
		},
	}.Run(t)

	TestCase{
		Interpolator: "count", // Use a PipeScript
		Options: map[string]interface{}{
			"run_on": "dt",
		},
		Reference: []pipescript.Datapoint{
			{Timestamp: 0.1, Duration: 0.5},
			{Timestamp: 0.6, Duration: 0.5},
			{Timestamp: 2.01, Duration: 1.2},
			{Timestamp: 3.4, Duration: 0},
			{Timestamp: 3.5, Duration: 2},
			{Timestamp: 4, Duration: 0},
		},
		Stream: []pipescript.Datapoint{
			{Timestamp: 1, Data: 1},
			{Timestamp: 2, Data: 1},
			{Timestamp: 3, Data: 1},
			{Timestamp: 4, Data: 1},
			{Timestamp: 5, Data: 1},
		},
		Output: []pipescript.Datapoint{
			{Data: int64(0)},
			{Timestamp: 1, Data: int64(1)},
			{Timestamp: 3, Duration: 0, Data: int64(1)},
			{Data: int64(0)},
			{Timestamp: 4, Duration: 1, Data: int64(2)},
			{Data: int64(0)},
		},
	}.Run(t)
}
