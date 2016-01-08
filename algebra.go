package pipescript

import (
	"errors"
	"math"

	"github.com/connectordb/duck"
)

var errConversion = errors.New("Failed to perform algebraic operation: couldn't convert to float.")

type addTransform struct{}

func (t addTransform) Copy() TransformInstance {
	return addTransform{}
}

func (t addTransform) Next(ti *TransformIterator) (*Datapoint, error) {
	te := ti.Next()
	if te.IsFinished() {
		return te.Get()
	}
	v, ok := duck.Add(te.Args[0].Data, te.Args[1].Data)
	if !ok {
		return nil, errConversion
	}
	return te.Set(v)
}

func addScript(a1, a2 *Script) (*Script, error) {
	pe, ok := NewPipelineElement([]*Script{a1, a2}, addTransform{})

	return &Script{
		input:      pe,
		output:     pe,
		IsOneToOne: true,
		Constant:   a1.Constant && a2.Constant,
	}, ok

}

type subtractTransform struct{}

func (t subtractTransform) Copy() TransformInstance {
	return subtractTransform{}
}

func (t subtractTransform) Next(ti *TransformIterator) (*Datapoint, error) {
	te := ti.Next()
	if te.IsFinished() {
		return te.Get()
	}
	v, ok := duck.Subtract(te.Args[0].Data, te.Args[1].Data)
	if !ok {
		return nil, errConversion
	}
	return te.Set(v)
}

func subtractScript(a1, a2 *Script) (*Script, error) {
	pe, ok := NewPipelineElement([]*Script{a1, a2}, subtractTransform{})

	return &Script{
		input:      pe,
		output:     pe,
		IsOneToOne: true,
		Constant:   a1.Constant && a2.Constant,
	}, ok

}

type mulTransform struct{}

func (t mulTransform) Copy() TransformInstance {
	return mulTransform{}
}

func (t mulTransform) Next(ti *TransformIterator) (*Datapoint, error) {
	te := ti.Next()
	if te.IsFinished() {
		return te.Get()
	}
	v, ok := duck.Multiply(te.Args[0].Data, te.Args[1].Data)
	if !ok {
		return nil, errConversion
	}
	return te.Set(v)
}

func mulScript(a1, a2 *Script) (*Script, error) {
	pe, ok := NewPipelineElement([]*Script{a1, a2}, mulTransform{})

	return &Script{
		input:      pe,
		output:     pe,
		IsOneToOne: true,
		Constant:   a1.Constant && a2.Constant,
	}, ok

}

type divTransform struct{}

func (t divTransform) Copy() TransformInstance {
	return divTransform{}
}

func (t divTransform) Next(ti *TransformIterator) (*Datapoint, error) {
	te := ti.Next()
	if te.IsFinished() {
		return te.Get()
	}
	v, ok := duck.Divide(te.Args[0].Data, te.Args[1].Data)
	if !ok {
		return nil, errConversion
	}
	return te.Set(v)
}

func divScript(a1, a2 *Script) (*Script, error) {
	pe, ok := NewPipelineElement([]*Script{a1, a2}, divTransform{})

	return &Script{
		input:      pe,
		output:     pe,
		IsOneToOne: true,
		Constant:   a1.Constant && a2.Constant,
	}, ok

}

type modTransform struct{}

func (t modTransform) Copy() TransformInstance {
	return modTransform{}
}

func (t modTransform) Next(ti *TransformIterator) (*Datapoint, error) {
	te := ti.Next()
	if te.IsFinished() {
		return te.Get()
	}
	v, ok := duck.Mod(te.Args[0].Data, te.Args[1].Data)
	if !ok {
		return nil, errConversion
	}
	return te.Set(v)
}

func modScript(a1, a2 *Script) (*Script, error) {
	pe, ok := NewPipelineElement([]*Script{a1, a2}, modTransform{})

	return &Script{
		input:      pe,
		output:     pe,
		IsOneToOne: true,
		Constant:   a1.Constant && a2.Constant,
	}, ok

}

type powTransform struct{}

func (t powTransform) Copy() TransformInstance {
	return powTransform{}
}

func (t powTransform) Next(ti *TransformIterator) (*Datapoint, error) {
	te := ti.Next()
	if te.IsFinished() {
		return te.Get()
	}
	v1, ok := duck.Float(te.Args[0].Data)
	if !ok {
		return nil, errConversion
	}
	v2, ok := duck.Float(te.Args[1].Data)
	if !ok {
		return nil, errConversion
	}
	return te.Set(math.Pow(v1, v2))
}

func powScript(a1, a2 *Script) (*Script, error) {
	pe, ok := NewPipelineElement([]*Script{a1, a2}, powTransform{})

	return &Script{
		input:      pe,
		output:     pe,
		IsOneToOne: true,
		Constant:   a1.Constant && a2.Constant,
	}, ok

}
