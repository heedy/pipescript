package pipescript

import (
	"errors"

	"github.com/connectordb/duck"
)

// eqTransform returns whether its two arguments are equal
type eqTransform struct{}

func (t eqTransform) Copy() TransformInstance {
	return eqTransform{}
}

func (t eqTransform) Next(ti *TransformIterator) (*Datapoint, error) {
	te := ti.Next()
	if te.IsFinished() {
		return te.Get()
	}
	cmp, ok := duck.Equal(te.Args[0].Data, te.Args[1].Data)
	if !ok {
		return nil, errors.New("Could not compare two objects")
	}
	return te.Set(cmp)
}

// neqTransform returns whether its two arguments are not equal
type neqTransform struct{}

func (t neqTransform) Copy() TransformInstance {
	return neqTransform{}
}

func (t neqTransform) Next(ti *TransformIterator) (*Datapoint, error) {
	te := ti.Next()
	if te.IsFinished() {
		return te.Get()
	}
	cmp, ok := duck.Equal(te.Args[0].Data, te.Args[1].Data)
	if !ok {
		return nil, errors.New("Could not compare two objects")
	}
	return te.Set(!cmp)
}

// ltTransform returns whether a < b
type ltTransform struct{}

func (t ltTransform) Copy() TransformInstance {
	return ltTransform{}
}

func (t ltTransform) Next(ti *TransformIterator) (*Datapoint, error) {
	te := ti.Next()
	if te.IsFinished() {
		return te.Get()
	}
	cmp, ok := duck.Lt(te.Args[0].Data, te.Args[1].Data)
	if !ok {
		return nil, errors.New("Values given to comparison not numerical!")
	}
	return te.Set(cmp)
}

// lteTransform returns whether a <= b
type lteTransform struct{}

func (t lteTransform) Copy() TransformInstance {
	return lteTransform{}
}

func (t lteTransform) Next(ti *TransformIterator) (*Datapoint, error) {
	te := ti.Next()
	if te.IsFinished() {
		return te.Get()
	}
	cmp, ok := duck.Lte(te.Args[0].Data, te.Args[1].Data)
	if !ok {
		return nil, errors.New("Values given to comparison not numerical!")
	}
	return te.Set(cmp)
}

// comparisonScript generates a comparison between two values using the given compasions string
// valid values for comparison are
//	==
//	!=
//	<
//	>
//	<=
//	>=
func comparisonScript(comparison string, a1, a2 *Script) (*Script, error) {
	var ti TransformInstance

	args := []*Script{a1, a2}

	switch comparison {
	case "==":
		ti = eqTransform{}
	case "!=":
		ti = neqTransform{}
	case "<":
		ti = ltTransform{}
	case "<=":
		ti = lteTransform{}
	case ">":
		ti = ltTransform{}
		// We flip the arguments
		args = []*Script{a2, a1}
	case ">=":
		ti = lteTransform{}
		// We flip the arguments
		args = []*Script{a2, a1}
	default:
		return nil, errors.New("Invalid comparison")
	}

	pe, err := NewPipelineElement(args, ti)

	return &Script{
		input:      pe,
		output:     pe,
		OneToOne: true,
		Constant:   a1.Constant && a2.Constant,
		Stateless:  true,
	}, err
}
