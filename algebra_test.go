package pipescript

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAlgebra(t *testing.T) {
	testcases := []struct {
		pipeline string
		output   interface{}
	}{
		{"5+5", float64(10)},
		{"'hello'+'world'", "helloworld"},
		{"23 - 3", float64(20)},
		{"23-3", float64(20)}, // This caused error since lexer checked negative numbers
		{"5*3", float64(15)},
		{"15/3", float64(5)},
		{"4%2", float64(0)},
		{"5%2", float64(1)},
		{"3^2", float64(9)},
	}

	for i := range testcases {
		s, err := Parse(testcases[i].pipeline)
		require.NoError(t, err, testcases[i].pipeline)
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
