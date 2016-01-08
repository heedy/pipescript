package core

import "github.com/connectordb/pipescript"

type tTransformStruct struct{}

func (t tTransformStruct) Copy() pipescript.TransformInstance {
	return tTransformStruct{}
}

func (t tTransformStruct) Next(ti *pipescript.TransformIterator) (*pipescript.Datapoint, error) {
	te := ti.Next()
	if te.IsFinished() {
		return te.Get()
	}
	return te.Set(te.Datapoint.Timestamp)
}

var tTransform = pipescript.Transform{
	Name:         "t",
	Description:  "The current datapoint's timestamp in floating point unix seconds",
	OutputSchema: `{"type": "number"}`,
	OneToOne:     true,
	Stateless:    true,
	Generator: func(name string, args []*pipescript.Script) (*pipescript.TransformInitializer, error) {
		return &pipescript.TransformInitializer{Transform: tTransformStruct{}}, nil
	},
}
