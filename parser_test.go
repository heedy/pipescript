package pipescript_test

import (
	"testing"

	"github.com/connectordb/pipescript"
	"github.com/connectordb/pipescript/transforms"
	"github.com/stretchr/testify/require"
)

func init() {
	// Register the core
	transforms.Register()
}

func TestSyntax(t *testing.T) {
	_, err := pipescript.Parse("-")
	require.Error(t, err)
	_, err = pipescript.Parse("'")
	require.Error(t, err)
	_, err = pipescript.Parse("this_transform_DNE")
	require.Error(t, err)
}

func TestParserConstant(t *testing.T) {
	pipescript.ConstantTestCases{
		// Check builtins
		{"true", true},
		{"false", false},

		// Check operator prescedence
		{"5+5/10", float64(5.5)},
		{"(5+5)/10", float64(1)},
		{"[5+5]/10", float64(1)},
		{"{5+5}/10", float64(1)},
		{"5-5*10", float64(-45)},
		{"(5-6)*-10", float64(10)},
		{"[5-6]*-10", float64(10)},
		{"{5-6}*-10", float64(10)},
		{"true and 1 - 1", false},
		{"false or 1 - 1", false},
		{"(true and 1) - 1", float64(0)},
		{"(false or 1) - 1", float64(0)},
		{"true and 5==5", true},
		{"-1 + 2", float64(1)},
		{"-(1 + 2)", float64(-3)},

		// Test strings
		{"\"❤ ☀ ☆ ☂ ☻ ♞ ☯ ☭ ☢ €\"", "❤ ☀ ☆ ☂ ☻ ♞ ☯ ☭ ☢ €"},
		{"'|'", "|"},
		{"\"string\"", "string"},
		{"'string'", "string"},
		{"'string\\n'", "string\n"},
		{"'string\\t'", "string\t"},
		{"'string\\\\'", "string\\"},
		{"'string\\r'", "string\r"},
		{"'string\"'", "string\""},
		{"'string\\''", "string'"},

		// Test Pipe
		{"5 | true", true},
	}.Run(t)
}

func TestParser(t *testing.T) {
	// Here we perform more advanced pipes to make sure everything works as it should in the parser
	// We assume all built-in functions are available

	pipescript.TestCase{
		Pipescript: "$!=1",
		Input: []pipescript.Datapoint{
			{1, 1},
		},
		Output: []pipescript.Datapoint{
			{1, false},
		},
	}.Run(t)

	pipescript.TestCase{
		Pipescript: "if $ < 5 | $ >= 3",
		Input: []pipescript.Datapoint{
			{1, 1},
			{2, 10},
			{3, 7},
			{4, 1.0},
			{5, 3},
			{6, 2.0},
			{7, 3.14},
		},
		Output: []pipescript.Datapoint{
			{1, false},
			{4, false},
			{5, true},
			{6, false},
			{7, true},
		},
	}.Run(t)

	pipescript.TestCase{
		Pipescript: "if($ < 5):($ >= 3)",
		Input: []pipescript.Datapoint{
			{1, 1},
			{2, 10},
			{3, 7},
			{4, 1.0},
			{5, 3},
			{6, 2.0},
			{7, 3.14},
		},
		Output: []pipescript.Datapoint{
			{1, false},
			{4, false},
			{5, true},
			{6, false},
			{7, true},
		},
	}.Run(t)

	pipescript.TestCase{
		// This tests order of prescedence: ":" pipes are high prescedence, and will be executed first
		Pipescript: "if ($['test']:$ < 5) | $['test']",
		Input: []pipescript.Datapoint{
			{1, map[string]int{"test": 4}},
			{2, map[string]int{"test": 8}},
			{3, map[string]int{"test": 3}},
		},
		Output: []pipescript.Datapoint{
			{1, 4},
			{3, 3},
		},
	}.Run(t)

	pipescript.TestCase{
		// This tests order of prescedence: ":" pipes are high prescedence, and will be executed first
		Pipescript: "if $['test']:$ < 5 | $['test']",
		Input: []pipescript.Datapoint{
			{1, map[string]int{"test": 4}},
			{2, map[string]int{"test": 8}},
			{3, map[string]int{"test": 3}},
		},
		Output: []pipescript.Datapoint{
			{1, 4},
			{3, 3},
		},
	}.Run(t)

	pipescript.TestCase{
		// This tests order of prescedence: ":" pipes are high prescedence, and will be executed first
		Pipescript: "if $:5 > $['test']:$:$:$ | $['test']:$",
		Input: []pipescript.Datapoint{
			{1, map[string]int{"test": 4}},
			{2, map[string]int{"test": 8}},
			{3, map[string]int{"test": 3}},
		},
		Output: []pipescript.Datapoint{
			{1, 4},
			{3, 3},
		},
	}.Run(t)

	// Test multiple arguments
	pipescript.TestCase{
		// This tests order of prescedence: ":" pipes are high prescedence, and will be executed first
		Pipescript: "map($ > 5, count)",
		Input: []pipescript.Datapoint{
			{1, 4},
			{2, 6},
			{3, 8},
		},
		Output: []pipescript.Datapoint{
			{3, map[string]interface{}{"false": int64(1), "true": int64(2)}},
		},
		SecondaryInput: []pipescript.Datapoint{
			{4, 8},
			{5, 4},
			{6, 6},
		},
		SecondaryOutput: []pipescript.Datapoint{
			{6, map[string]interface{}{"false": int64(1), "true": int64(2)}},
		},
	}.Run(t)
}

func TestIdentifierSubtract(t *testing.T) {
	// Test multiple arguments
	pipescript.TestCase{
		// This is a weird parser issue I got
		Pipescript: "month() - remember(first, month)",
		Input: []pipescript.Datapoint{
			{1, 4},
		},
		Output: []pipescript.Datapoint{
			{1, float64(0)},
		},
	}.Run(t)
	pipescript.TestCase{
		// This is a weird parser issue I got
		Pipescript: "month (- remember(first, month))",
		ParseError: true,
	}.Run(t)
	pipescript.TestCase{
		// This is a weird parser issue I got
		Pipescript: "month-remember(first, month)",
		Input: []pipescript.Datapoint{
			{1, 4},
		},
		Output: []pipescript.Datapoint{
			{1, float64(0)},
		},
	}.Run(t)
	pipescript.TestCase{
		// This is a weird parser issue I got
		Pipescript: "month - remember(first, month)",
		Input: []pipescript.Datapoint{
			{1, 4},
		},
		Output: []pipescript.Datapoint{
			{1, float64(0)},
		},
	}.Run(t)

}
