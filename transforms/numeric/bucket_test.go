package numeric

import (
	"testing"

	"github.com/heedy/pipescript"
)

func TestBucket(t *testing.T) {
	Register()
	pipescript.TestCase{
		Pipescript: "bucket",
		Input: []pipescript.Datapoint{
			{Timestamp: 1, Data: 2},
			{Timestamp: 2, Data: 16},
			{Timestamp: 3, Data: 84},
			{Timestamp: 4, Data: -5},
			{Timestamp: 5, Data: 1},
		},
		Output: []pipescript.Datapoint{
			{Timestamp: 1, Data: "[0,10)"},
			{Timestamp: 2, Data: "[10,20)"},
			{Timestamp: 3, Data: "[80,90)"},
			{Timestamp: 4, Data: "[-10,0)"},
			{Timestamp: 5, Data: "[0,10)"},
		},
	}.Run(t)

	pipescript.TestCase{
		Pipescript: "bucket(0.5,0.1)",
		Input: []pipescript.Datapoint{
			{Timestamp: 1, Data: 2},
			{Timestamp: 2, Data: 2.1},
			{Timestamp: 3, Data: -0.4},
		},
		Output: []pipescript.Datapoint{
			{Timestamp: 1, Data: "[1.6,2.1)"},
			{Timestamp: 2, Data: "[2.1,2.6)"},
			{Timestamp: 3, Data: "[-0.4,0.1)"},
		},
	}.Run(t)
}
