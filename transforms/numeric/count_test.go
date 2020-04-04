package numeric

import (
	"testing"

	"github.com/heedy/pipescript"
)

func TestCount(t *testing.T) {
	Count.Register()
	pipescript.TestCase{
		Pipescript: "count",
		Input: []pipescript.Datapoint{
			{Timestamp: 1, Data: int64(2)},
			{Timestamp: 2, Data: 3},
			{Timestamp: 3, Data: 4},
		},
		Output: []pipescript.Datapoint{
			{Timestamp: 1, Duration: 2, Data: int64(3)},
		},
	}.Run(t)
}
