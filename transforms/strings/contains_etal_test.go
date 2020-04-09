package strings

import (
	"testing"

	"github.com/heedy/pipescript"
)

func TestContains(t *testing.T) {
	Register()
	pipescript.TestCase{
		Pipescript: "contains('hi')",
		Input: []pipescript.Datapoint{
			{Timestamp: 1, Data: "Hello World!"},
			{Timestamp: 2, Data: " high"},
			{Timestamp: 3, Data: 34},
		},
		Output: []pipescript.Datapoint{
			{Timestamp: 1, Data: false},
			{Timestamp: 2, Data: true},
			{Timestamp: 3, Data: false},
		},
	}.Run(t)
}

func TestStartswith(t *testing.T) {
	Register()
	pipescript.TestCase{
		Pipescript: "startswith('hi')",
		Input: []pipescript.Datapoint{
			{Timestamp: 1, Data: "Hello World!"},
			{Timestamp: 2, Data: "hiyo"},
			{Timestamp: 3, Data: "yohi"},
		},
		Output: []pipescript.Datapoint{
			{Timestamp: 1, Data: false},
			{Timestamp: 2, Data: true},
			{Timestamp: 3, Data: false},
		},
	}.Run(t)
}

func TestEndswith(t *testing.T) {
	Register()
	pipescript.TestCase{
		Pipescript: "endswith('hi')",
		Input: []pipescript.Datapoint{
			{Timestamp: 1, Data: "Hello World!"},
			{Timestamp: 2, Data: "hiyo"},
			{Timestamp: 3, Data: "yohi"},
		},
		Output: []pipescript.Datapoint{
			{Timestamp: 1, Data: false},
			{Timestamp: 2, Data: false},
			{Timestamp: 3, Data: true},
		},
	}.Run(t)
}
