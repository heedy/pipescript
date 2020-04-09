package misc

import (
	"testing"

	"github.com/heedy/pipescript"
)

func TestLength(t *testing.T) {
	Register()
	pipescript.TestCase{
		Pipescript: "length",
		Input: []pipescript.Datapoint{
			{Timestamp: 1, Data: "hi"},
			{Timestamp: 2, Data: []interface{}{1, 2, 3, 4, 5}},
			{Timestamp: 3, Data: map[string]interface{}{"hi": nil, "hr": nil, "ree": nil}},
		},
		Output: []pipescript.Datapoint{
			{Timestamp: 1, Data: 2},
			{Timestamp: 2, Data: 5},
			{Timestamp: 3, Data: 3},
		},
	}.Run(t)
	pipescript.TestCase{
		Pipescript:  "length",
		OutputError: true,
		Input: []pipescript.Datapoint{
			{Timestamp: 1, Data: 2},
		},
	}.Run(t)
}
