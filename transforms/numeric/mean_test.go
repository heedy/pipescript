package numeric

import (
	"testing"

	"github.com/heedy/pipescript"
)

func TestMean(t *testing.T) {
	Mean.Register()
	pipescript.TestCase{
		Pipescript: "mean",
		Input: []pipescript.Datapoint{
			{Timestamp: 1, Data: 1},
			{Timestamp: 2, Data: 3},
			{Timestamp: 3, Data: 5},
		},
		Output: []pipescript.Datapoint{
			{Timestamp: 1, Duration: 2, Data: float64(3)},
		},
	}.Run(t)
}
