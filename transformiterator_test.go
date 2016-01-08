package pipescript

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// This transform assumes that the input is a sequence of increasing integers 1,2,3,4...
// The second argument is assumed to be a constant 1
type testPeekTransform struct{}

func (t testPeekTransform) Copy() TransformInstance {
	return testPeekTransform{}
}

func (t testPeekTransform) Next(ti *TransformIterator) (*Datapoint, error) {
	te := ti.Next()
	if te.IsFinished() {
		return te.Get()
	}
	return te.Get() //UNFINISHED
}

func TestTransformIterator(t *testing.T) {
	assert.NoError(t, Transform{
		Name:        "testiterator",
		Description: "I am testing!",

		Generator: func(name string, args []*Script) (*TransformInitializer, error) {
			return &TransformInitializer{Args: args, Transform: testPeekTransform{}}, nil
		},
	}.Register())

}
