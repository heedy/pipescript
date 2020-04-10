package core

import (
	"github.com/heedy/pipescript"
	"github.com/heedy/pipescript/resources"
)

type ttrueTransform struct{}

func (t *ttrueTransform) Copy() (pipescript.TransformInstance, error) {
	return &ttrueTransform{}, nil
}

func (t *ttrueTransform) Next(ti *pipescript.TransformIterator) (*pipescript.Datapoint, error) {
	var start float64
	var end float64

	// First, we loop until we get a true
	te := ti.Next()
	for {
		if te.IsFinished() {
			return te.Get()
		}
		b, err := te.Datapoint.Bool()
		if err != nil {
			return nil, err
		}
		if b {
			start = te.Datapoint.Timestamp
			break
		}
		te = ti.Next()
	}

	// if this is the end of the stream, we're done
	te = ti.Next()
	if te.IsFinished() {
		return te.Get()
	}

	// Otherwise, loop until we get false or end of stream
	for {
		end = te.Datapoint.Timestamp
		b, err := te.Datapoint.Bool()
		if err != nil {
			return nil, err
		}
		if !b {
			return te.Set(end - start)
		}

		te = ti.Next()
		if te.IsFinished() {
			return &pipescript.Datapoint{Timestamp: end, Data: end - start}, nil
		}
	}
}

var Ttrue = pipescript.Transform{
	Name:          "ttrue",
	Description:   "The time period for which a boolean stream is true before turning false",
	Documentation: string(resources.MustAsset("docs/transforms/ttrue.md")),
	Generator: func(name string, args []*pipescript.Script) (*pipescript.TransformInitializer, error) {
		return &pipescript.TransformInitializer{Transform: &ttrueTransform{}}, nil
	},
}
