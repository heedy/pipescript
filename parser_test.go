package pipescript

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSyntax(t *testing.T) {
	_, err := Parse("-")
	require.Error(t, err)
	_, err = Parse("'")
	require.Error(t, err)
	_, err = Parse("this_transform_DNE")
	require.Error(t, err)
}

func TestParserConstant(t *testing.T) {
	ConstantTestCases{
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

	TestCase{
		Pipescript: "if $ < 5 | $ >= 3",
		Input: []Datapoint{
			{1, 1},
			{2, 10},
			{3, 7},
			{4, 1.0},
			{5, 3},
			{6, 2.0},
			{7, 3.14},
		},
		Output: []Datapoint{
			{1, false},
			{4, false},
			{5, true},
			{6, false},
			{7, true},
		},
	}.Run(t)

	TestCase{
		Pipescript: "if($ < 5):($ >= 3)",
		Input: []Datapoint{
			{1, 1},
			{2, 10},
			{3, 7},
			{4, 1.0},
			{5, 3},
			{6, 2.0},
			{7, 3.14},
		},
		Output: []Datapoint{
			{1, false},
			{4, false},
			{5, true},
			{6, false},
			{7, true},
		},
	}.Run(t)

	TestCase{
		// This tests order of prescedence: ":" pipes are high prescedence, and will be executed first
		Pipescript: "if ($['test']:$ < 5) | $['test']",
		Input: []Datapoint{
			{1, map[string]int{"test": 4}},
			{2, map[string]int{"test": 8}},
			{3, map[string]int{"test": 3}},
		},
		Output: []Datapoint{
			{1, 4},
			{3, 3},
		},
	}.Run(t)

	TestCase{
		// This tests order of prescedence: ":" pipes are high prescedence, and will be executed first
		Pipescript: "if $['test']:$ < 5 | $['test']",
		Input: []Datapoint{
			{1, map[string]int{"test": 4}},
			{2, map[string]int{"test": 8}},
			{3, map[string]int{"test": 3}},
		},
		Output: []Datapoint{
			{1, 4},
			{3, 3},
		},
	}.Run(t)

}
