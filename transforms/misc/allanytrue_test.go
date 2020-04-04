package misc

import (
	"testing"

	"github.com/heedy/pipescript"
)

func TestAnyTrue(t *testing.T) {
	Anytrue.Register()
	pipescript.TestCase{
		Pipescript: "anytrue",
		Input: []pipescript.Datapoint{
			{Timestamp: 1, Data: false},
			{Timestamp: 2, Data: false},
			{Timestamp: 3, Data: true},
		},
		Output: []pipescript.Datapoint{
			{Timestamp: 3, Data: true},
		},
	}.Run(t)
	pipescript.TestCase{
		Pipescript: "anytrue",
		Input: []pipescript.Datapoint{
			{Timestamp: 1, Data: false},
			{Timestamp: 2, Data: false},
			{Timestamp: 3, Data: false},
		},
		Output: []pipescript.Datapoint{
			{Timestamp: 1, Duration: 2, Data: false},
		},
	}.Run(t)
}
func TestAllTrue(t *testing.T) {
	Alltrue.Register()
	pipescript.TestCase{
		Pipescript: "alltrue",
		Input: []pipescript.Datapoint{
			{Timestamp: 1, Data: true},
			{Timestamp: 2, Data: false},
			{Timestamp: 3, Data: true},
		},
		Output: []pipescript.Datapoint{
			{Timestamp: 2, Data: false},
		},
	}.Run(t)
	pipescript.TestCase{
		Pipescript: "alltrue",
		Input: []pipescript.Datapoint{
			{Timestamp: 1, Data: true},
			{Timestamp: 2, Data: true},
			{Timestamp: 3, Data: true},
		},
		Output: []pipescript.Datapoint{
			{Timestamp: 1, Duration: 2, Data: true},
		},
	}.Run(t)
}
