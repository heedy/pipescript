package pipescript

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConstant(t *testing.T) {
	cs := ConstantScript(1337)

	require.True(t, cs.Constant)
	require.True(t, cs.IsOneToOne)

	testarray := []Datapoint{
		Datapoint{1, 1}, Datapoint{2, 2}, Datapoint{3, 3}, Datapoint{4, 4},
	}
	cs.SetInput(NewDatapointArrayIterator(testarray))

	for i := 1; i < 5; i++ {
		dp, err := cs.Next()
		require.NoError(t, err)
		require.EqualValues(t, &Datapoint{float64(i), 1337}, dp)
	}

	dp, err := cs.Next()
	require.NoError(t, err)
	require.Nil(t, dp)
}

func TestAndTransform(t *testing.T) {
	s, err := andScript(ConstantScript(true), ConstantScript(true))
	require.NoError(t, err)

	require.True(t, s.Constant)
	require.True(t, s.IsOneToOne)
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
	one.IsOneToOne = false
	s, err = andScript(one, ConstantScript("false"))
	require.Error(t, err)
}

func TestOrTransform(t *testing.T) {
	s, err := orScript(ConstantScript(true), ConstantScript(true))
	require.NoError(t, err)

	require.True(t, s.Constant)
	require.True(t, s.IsOneToOne)
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
	one.IsOneToOne = false
	s, err = orScript(one, ConstantScript("false"))
	require.Error(t, err)
}
