package core

import (
	"testing"

	"github.com/connectordb/pipescript"
	"github.com/stretchr/testify/require"
)

func TestIdentityError(t *testing.T) {
	c := pipescript.ConstantScript("hi")
	c.Constant = false
	_, err := IdentityTransform.Generator("$", []*pipescript.Script{c})
	require.Error(t, err)
}

func TestIdentity(t *testing.T) {
	pipescript.TestCase{
		Pipescript: "$ | $",
		Input: []pipescript.Datapoint{
			{1, 1},
			{2, true},
			{3, "hi"},
		},
		Output: []pipescript.Datapoint{
			{1, 1},
			{2, true},
			{3, "hi"},
		},
		SecondaryInput: []pipescript.Datapoint{
			{4, 4},
			{5, 5},
		},
		SecondaryOutput: []pipescript.Datapoint{
			{4, 4},
			{5, 5},
		},
	}.Run(t)

	pipescript.TestCase{
		Pipescript: "$(1)",
		Input: []pipescript.Datapoint{
			{1, []string{"hi", "ho"}},
			{2, map[string]interface{}{"hello": "world", "1": map[string]interface{}{"wee": "mo"}}},
		},
		Output: []pipescript.Datapoint{
			{1, "ho"},
			{2, map[string]interface{}{"wee": "mo"}},
		},
	}.Run(t)

	pipescript.TestCase{
		Pipescript:  "$[1]",
		OutputError: true,
		Input: []pipescript.Datapoint{
			{1, []string{"hi", "ho"}},
			{2, map[string]string{"hello": "world", "1": "bar"}},
		},
		Output: []pipescript.Datapoint{
			{1, "ho"},
			{2, "bar"},
		},
		SecondaryInput: []pipescript.Datapoint{
			{4, map[string]string{"hello": "world", "1": "bar"}},
			{5, 5}, // This should give an error
		},
		SecondaryOutput: []pipescript.Datapoint{
			{4, "bar"},
		},
	}.Run(t)

	pipescript.TestCase{
		Pipescript: "$[$]", // The transform requries a constant
		ParseError: true,
	}.Run(t)

	pipescript.TestCase{
		Pipescript: "$[1,2]", // Identity only accepts one arg
		ParseError: true,
	}.Run(t)

	pipescript.TestCase{
		Pipescript:  "$ 1", // Bash-like usage
		OutputError: true,
		Input: []pipescript.Datapoint{
			{1, []string{"hi", "ho"}},
			{2, map[string]string{"hello": "world", "1": "bar"}},
		},
		Output: []pipescript.Datapoint{
			{1, "ho"},
			{2, "bar"},
		},
		SecondaryInput: []pipescript.Datapoint{
			{4, map[string]string{"hello": "world", "1": "bar"}},
			{5, 5}, // This should give an error
		},
		SecondaryOutput: []pipescript.Datapoint{
			{4, "bar"},
		},
	}.Run(t)

}


func TestIdentityD(t *testing.T) {
	// d is same as identity. Make sure it works.
	pipescript.TestCase{
		Pipescript: "d | d",
		Input: []pipescript.Datapoint{
			{1, 1},
			{2, true},
			{3, "hi"},
		},
		Output: []pipescript.Datapoint{
			{1, 1},
			{2, true},
			{3, "hi"},
		},
		SecondaryInput: []pipescript.Datapoint{
			{4, 4},
			{5, 5},
		},
		SecondaryOutput: []pipescript.Datapoint{
			{4, 4},
			{5, 5},
		},
	}.Run(t)
}