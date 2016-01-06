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

func eqScript(a1, a2 *Script) (*Script, error) {
	pe, err := NewPipelineElement([]*Script{a1, a2}, eqTransform{})

	return &Script{
		input:      pe,
		output:     pe,
		IsOneToOne: true,
		Constant:   a1.Constant && a2.Constant,
	}, err
}

// neqTransform returns whether its two arguments are not equal
type neqTransform struct{}

func (t neqTransform) Copy() TransformInstance {
	return eqTransform{}
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

func neqScript(a1, a2 *Script) (*Script, error) {
	pe, err := NewPipelineElement([]*Script{a1, a2}, neqTransform{})

	return &Script{
		input:      pe,
		output:     pe,
		IsOneToOne: true,
		Constant:   a1.Constant && a2.Constant,
	}, err
}

// ltTransform returns whether a < b
type ltTransform struct{}

func (t ltTransform) Copy() TransformInstance {
	return eqTransform{}
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

func ltScript(a1, a2 *Script) (*Script, error) {
	pe, err := NewPipelineElement([]*Script{a1, a2}, ltTransform{})

	return &Script{
		input:      pe,
		output:     pe,
		IsOneToOne: true,
		Constant:   a1.Constant && a2.Constant,
	}, err
}

// gtScript uses ltTransform with args reversed
func gtScript(a1, a2 *Script) (*Script, error) {
	pe, err := NewPipelineElement([]*Script{a2, a1}, ltTransform{})

	return &Script{
		input:      pe,
		output:     pe,
		IsOneToOne: true,
		Constant:   a1.Constant && a2.Constant,
	}, err
}

// lteTransform returns whether a <= b
type lteTransform struct{}

func (t lteTransform) Copy() TransformInstance {
	return eqTransform{}
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

func lteScript(a1, a2 *Script) (*Script, error) {
	pe, err := NewPipelineElement([]*Script{a1, a2}, lteTransform{})

	return &Script{
		input:      pe,
		output:     pe,
		IsOneToOne: true,
		Constant:   a1.Constant && a2.Constant,
	}, err
}

// gteScript uses lteTransform with args reversed
func gteScript(a1, a2 *Script) (*Script, error) {
	pe, err := NewPipelineElement([]*Script{a2, a1}, lteTransform{})

	return &Script{
		input:      pe,
		output:     pe,
		IsOneToOne: true,
		Constant:   a1.Constant && a2.Constant,
	}, err
}
