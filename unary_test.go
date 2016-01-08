package pipescript

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnary(t *testing.T) {
	testcases := []struct {
		pipeline string
		output   interface{}
	}{
		{"not true", false},
		{"not false", true},
		{"-true", false},
		{"-false", true},
		{"- 50", float64(-50)},
		{"- '0.54'", float64(-0.54)},
	}

	for i := range testcases {
		s, err := Parse(testcases[i].pipeline)
		require.NoError(t, err)
		s2, err := s.Copy()
		require.NoError(t, err)
		bv, err := s.GetConstant()
		require.NoError(t, err)
		require.NotNil(t, bv)
		require.Equal(t, testcases[i].output, bv.Data, testcases[i].pipeline)
		bv, err = s2.GetConstant()
		require.NoError(t, err)
		require.NotNil(t, bv)
		require.Equal(t, testcases[i].output, bv.Data, testcases[i].pipeline)
	}
}
