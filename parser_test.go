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
}

func TestParser(t *testing.T) {
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
