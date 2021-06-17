package pipescript

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParserConst(t *testing.T) {
	cases := []struct {
		script string
		output interface{}
	}{
		{"true", true},
		{"false", false},
		{"not (0.1 < 1)", false},
		{"not 0.1 < 1", false},
		{"(not 1) < 1", true},

		// Check operator prescedence
		{"5+5/10", float64(5.5)},
		{"(5+5)/10", float64(1)},
		{"[5+5]/10", float64(1)},
		{"5-5*10", float64(-45)},
		{"(5-6)*-10", float64(10)},
		{"[5-6]*-10", float64(10)},
		{"true and 1 - 1", false},
		{"false or 1 - 1", false},
		{"(true and 1) - 1", float64(0)},
		{"(false or 1) - 1", float64(0)},
		{"true and 5==5", true},
		{"-1 + 2", float64(1)},
		{"-(1 + 2)", float64(-3)},
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
	}

	for _, c := range cases {
		p, err := Parse(c.script)
		require.NoError(t, err, c.script)
		v, err := p.GetConst()
		require.NoError(t, err, "(%s) (%s)", c.script, p.String())
		require.Equal(t, c.output, v, "(%s) (%s)", c.script, p.String())
	}
}

func TestParser(t *testing.T) {
	// Here we perform more advanced pipes to make sure everything works as it should in the parser
	// We assume all built-in functions are available
	TestCase{
		Pipescript: "d==1",
		Input: []Datapoint{
			{Timestamp: 1, Data: 1},
			{Timestamp: 2, Data: 2},
		},
		Output: []Datapoint{
			{Timestamp: 1, Data: true},
			{Timestamp: 2, Data: false},
		},
	}.Run(t)
	TestCase{
		Pipescript: "d!=1",
		Input: []Datapoint{
			{Timestamp: 1, Data: 1},
			{Timestamp: 2, Data: 2},
		},
		Output: []Datapoint{
			{Timestamp: 1, Data: false},
			{Timestamp: 2, Data: true},
		},
	}.Run(t)

	TestCase{
		Pipescript: "d >= 3",
		Input: []Datapoint{
			{Timestamp: 1, Data: 1},
			{Timestamp: 4, Data: 1.0},
			{Timestamp: 5, Data: 3},
			{Timestamp: 6, Data: 2.0},
			{Timestamp: 7, Data: 3.14},
		},
		Output: []Datapoint{
			{Timestamp: 1, Data: false},
			{Timestamp: 4, Data: false},
			{Timestamp: 5, Data: true},
			{Timestamp: 6, Data: false},
			{Timestamp: 7, Data: true},
		},
	}.Run(t)

	TestCase{
		Pipescript: "d(1):(d >= 3)",
		Input: []Datapoint{
			{Timestamp: 1, Data: 1},
			{Timestamp: 4, Data: 1.0},
			{Timestamp: 5, Data: 3, Duration: 2},
			{Timestamp: 6, Data: 2.0},
			{Timestamp: 7, Data: 3.14},
		},
		Output: []Datapoint{
			{Timestamp: 4, Data: false},
			{Timestamp: 5, Data: true, Duration: 2},
			{Timestamp: 6, Data: false},
			{Timestamp: 7, Data: true},
		},
	}.Run(t)

	TestCase{
		// This tests order of prescedence: ":" pipes are high prescedence, and will be executed first
		Pipescript: "d['test']",
		Input: []Datapoint{
			{Timestamp: 1, Data: map[string]interface{}{"test": 4}},
			{Timestamp: 3, Data: map[string]interface{}{"test": 3}},
		},
		Output: []Datapoint{
			{Timestamp: 1, Data: 4},
			{Timestamp: 3, Data: 3},
		},
	}.Run(t)

	TestCase{
		// This tests order of prescedence: ":" pipes are high prescedence, and will be executed first
		Pipescript: "d['test'] - d[2-1]:d['test']",
		Input: []Datapoint{
			{Timestamp: 1, Data: map[string]interface{}{"test": 4}},
			{Timestamp: 2, Data: map[string]interface{}{"test": 8}, Duration: 1},
			{Timestamp: 3, Data: map[string]interface{}{"test": 3}},
		},
		Output: []Datapoint{
			{Timestamp: 1, Data: float64(-4)},
			{Timestamp: 2, Data: float64(5), Duration: 1},
		},
	}.Run(t)

	TestCase{
		Pipescript: "t",
		Input: []Datapoint{
			{Timestamp: 1, Data: 1},
			{Timestamp: 4, Data: 1.0},
			{Timestamp: 5, Data: 3, Duration: 6},
			{Timestamp: 6, Data: 2.0},
			{Timestamp: 7, Data: 3.14},
		},
		Output: []Datapoint{
			{Timestamp: 1, Data: float64(1)},
			{Timestamp: 4, Data: float64(4)},
			{Timestamp: 5, Data: float64(5), Duration: 6},
			{Timestamp: 6, Data: float64(6)},
			{Timestamp: 7, Data: float64(7)},
		},
	}.Run(t)

	TestCase{
		Pipescript: "dt",
		Input: []Datapoint{
			{Timestamp: 1, Data: 1, Duration: 5},
			{Timestamp: 4, Data: 1.0, Duration: 4},
			{Timestamp: 5, Data: 3, Duration: 3},
			{Timestamp: 6, Data: 2.0, Duration: 2},
			{Timestamp: 7, Data: 3.14, Duration: 1},
		},
		Output: []Datapoint{
			{Timestamp: 1, Data: float64(5), Duration: 5},
			{Timestamp: 4, Data: float64(4), Duration: 4},
			{Timestamp: 5, Data: float64(3), Duration: 3},
			{Timestamp: 6, Data: float64(2), Duration: 2},
			{Timestamp: 7, Data: float64(1), Duration: 1},
		},
	}.Run(t)

}

func TestObject(t *testing.T) {
	TestCase{
		Pipescript: "{'a':d,'b':d[1]}",
		Input: []Datapoint{
			{Timestamp: 1, Data: 1},
			{Timestamp: 4, Data: 2},
			{Timestamp: 5, Data: 3},
			{Timestamp: 6, Data: 4},
			{Timestamp: 7, Data: 5},
		},
		Output: []Datapoint{
			{Timestamp: 1, Data: map[string]interface{}{"a": 1, "b": 2}},
			{Timestamp: 4, Data: map[string]interface{}{"a": 2, "b": 3}},
			{Timestamp: 5, Data: map[string]interface{}{"a": 3, "b": 4}},
			{Timestamp: 6, Data: map[string]interface{}{"a": 4, "b": 5}},
		},
	}.Run(t)
}
