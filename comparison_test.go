package pipescript

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestComparison(t *testing.T) {
	testcases := []struct {
		pipeline string
		output   bool
	}{
		{"true==true", true},
		{"false==false", true},
		{"false!=false", false},
		{"false<true", true},
		{"'hello'=='hello'", true},
		{"'hello'!='hello'", false},
		{"'hello'=='Hello'", false},
		{"'hello'!='Hello'", true},
		{"5==5", true},
		{"6 ==5", false},
		{"6!=5", true},
		{" 5 != 5 ", false},
		{"5 < 5", false},
		{"5 < 6", true},
		{"20 < 135", true},
		{"5 > 5", false},
		{"5 > 6", false},
		{"200 > 135", true},
		{"5.3 <= 5.3", true},
		{"5.3 >= 5.3", true},
		{"5.2 <= 5.3", true},
		{"5.2 >= 5.3", false},
		{"5.3 <= 5", false},
		{"5.3 >= 5", true},
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
