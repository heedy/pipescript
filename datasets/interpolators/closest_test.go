package interpolators

import (
	"testing"

	"github.com/heedy/pipescript"
	"github.com/heedy/pipescript/datasets"
)

func TestClosestInterpolator(t *testing.T) {
	Closest.Register()
	datasets.TestCase{
		Interpolator: "closest",
		Reference: []pipescript.Datapoint{
			{Timestamp: 1},
			{Timestamp: 2},
			{Timestamp: 3},
			{Timestamp: 4},
			{Timestamp: 5},
		},
		Stream: []pipescript.Datapoint{
			{Timestamp: 1, Data: "1"},
			{Timestamp: 2, Data: "2"},
			{Timestamp: 4.1, Data: "3"},
			{Timestamp: 4.2, Data: "4"},
			{Timestamp: 4.9, Data: "5"},
		},
		Output: []pipescript.Datapoint{
			{Timestamp: 1, Data: "1"},
			{Timestamp: 2, Data: "2"},
			{Timestamp: 2, Data: "2"},
			{Timestamp: 4.1, Data: "3"},
			{Timestamp: 4.9, Data: "5"},
		},
	}.Run(t)
}
