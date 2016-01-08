package pipescript

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// This transform assumes that the input is a sequence of increasing integers 1,2,3,4...
// The second argument is assumed to be a constant 1
type testPeekTransform struct{ t *testing.T }

func (t testPeekTransform) Copy() TransformInstance {
	return testPeekTransform{}
}

func (t testPeekTransform) Next(ti *TransformIterator) (*Datapoint, error) {
	require.Equal(t.t, ti.Peek(5).Datapoint.Timestamp, float64(6))
	te := ti.Next()
	require.Equal(t.t, te.Datapoint.Timestamp, float64(1))
	require.Equal(t.t, ti.Peek(0).Datapoint.Timestamp, float64(2))
	require.Equal(t.t, ti.Peek(1).Datapoint.Timestamp, float64(3))
	require.Equal(t.t, ti.Peek(4).Datapoint.Timestamp, float64(6))
	require.Equal(t.t, ti.Peek(3).Datapoint.Timestamp, float64(5))

	require.Nil(t.t, ti.Peek(5).Datapoint)

	require.Equal(t.t, ti.Next().Datapoint.Timestamp, float64(2))
	require.Equal(t.t, ti.Next().Datapoint.Timestamp, float64(3))

	require.Equal(t.t, ti.Peek(2).Datapoint.Timestamp, float64(6))
	require.Nil(t.t, ti.Peek(3).Datapoint)

	require.Equal(t.t, ti.Next().Datapoint.Timestamp, float64(4))
	require.Equal(t.t, ti.Next().Datapoint.Timestamp, float64(5))
	require.Equal(t.t, ti.Next().Datapoint.Timestamp, float64(6))

	require.Nil(t.t, ti.Peek(0).Datapoint)
	require.Nil(t.t, ti.Peek(3).Datapoint)

	require.Nil(t.t, ti.Next().Datapoint)

	require.Nil(t.t, ti.Peek(0).Datapoint)
	require.Nil(t.t, ti.Peek(3).Datapoint)

	return te.Get()
}

func TestTransformIterator(t *testing.T) {
	assert.NoError(t, Transform{
		Name:        "testiterator",
		Description: "I am testing!",

		Generator: func(name string, args []*Script) (*TransformInitializer, error) {
			return &TransformInitializer{Args: args, Transform: testPeekTransform{t}}, nil
		},
	}.Register())

	dpa := NewDatapointArrayIterator([]Datapoint{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}, {6, 6}})

	s, err := Parse("testiterator")
	s.SetInput(dpa)
	require.NoError(t, err)
	s.Next()
}
