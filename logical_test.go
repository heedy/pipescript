package pipescript

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAndTransform(t *testing.T) {
	s, err := andScript(ConstantScript(true), ConstantScript(true))
	require.NoError(t, err)

	require.True(t, s.Constant)
	require.True(t, s.OneToOne)
	vb, err := s.GetConstant()
	require.NoError(t, err)
	v, err := vb.Bool()
	require.NoError(t, err)
	require.True(t, v)

	s, err = andScript(ConstantScript(true), ConstantScript(false))
	require.NoError(t, err)
	vb, err = s.GetConstant()
	require.NoError(t, err)
	v, err = vb.Bool()
	require.NoError(t, err)
	require.False(t, v)

	s, err = andScript(ConstantScript(1.0), ConstantScript("true"))
	require.NoError(t, err)
	vb, err = s.GetConstant()
	require.NoError(t, err)
	v, err = vb.Bool()
	require.NoError(t, err)
	require.True(t, v)

	s, err = andScript(ConstantScript(1.0), ConstantScript("false"))
	require.NoError(t, err)
	vb, err = s.GetConstant()
	require.NoError(t, err)
	v, err = vb.Bool()
	require.NoError(t, err)
	require.False(t, v)

	// Now get error
	s, err = andScript(ConstantScript("pff"), ConstantScript("false"))
	require.NoError(t, err)
	vb, err = s.GetConstant()
	require.Error(t, err)

	one := ConstantScript("true")
	one.OneToOne = false
	s, err = andScript(one, ConstantScript("false"))
	require.Error(t, err)
}

func TestOrTransform(t *testing.T) {
	s, err := orScript(ConstantScript(true), ConstantScript(true))
	require.NoError(t, err)

	require.True(t, s.Constant)
	require.True(t, s.OneToOne)
	vb, err := s.GetConstant()
	require.NoError(t, err)
	v, err := vb.Bool()
	require.NoError(t, err)
	require.True(t, v)

	s, err = orScript(ConstantScript(true), ConstantScript(false))
	require.NoError(t, err)
	vb, err = s.GetConstant()
	require.NoError(t, err)
	v, err = vb.Bool()
	require.NoError(t, err)
	require.True(t, v)

	s, err = orScript(ConstantScript(0), ConstantScript("true"))
	require.NoError(t, err)
	vb, err = s.GetConstant()
	require.NoError(t, err)
	v, err = vb.Bool()
	require.NoError(t, err)
	require.True(t, v)

	s, err = orScript(ConstantScript(0), ConstantScript("false"))
	require.NoError(t, err)
	vb, err = s.GetConstant()
	require.NoError(t, err)
	v, err = vb.Bool()
	require.NoError(t, err)
	require.False(t, v)

	// Now get error
	s, err = orScript(ConstantScript("pff"), ConstantScript("false"))
	require.NoError(t, err)
	vb, err = s.GetConstant()
	require.Error(t, err)

	one := ConstantScript("true")
	one.OneToOne = false
	s, err = orScript(one, ConstantScript("false"))
	require.Error(t, err)
}

func TestLogical(t *testing.T) {
	testcases := []struct {
		pipeline string
		output   bool
	}{
		{"true and true", true},
		{"false and false", false},
		{"true and false", false},
		{"false or true", true},
		{"false or false", false},
		{"true or true", true},
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
