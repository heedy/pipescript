package core

import (
	"github.com/connectordb/pipescript"
	"github.com/connectordb/pipescript/resources"
)

type tTransformStruct struct{}

func (t tTransformStruct) Copy() (pipescript.TransformInstance, error) {
	return tTransformStruct{}, nil
}

func (t tTransformStruct) Next(ti *pipescript.TransformIterator) (*pipescript.Datapoint, error) {
	te := ti.Next()
	if te.IsFinished() {
		return te.Get()
	}
	return te.Set(te.Datapoint.Timestamp)
}

var T = pipescript.Transform{
	Name:          "t",
	Description:   "The current datapoint's timestamp in floating point unix seconds",
	OutputSchema:  `{"type": "number"}`,
	Documentation: string(resources.MustAsset("docs/transforms/t.md")),
	OneToOne:      true,
	Stateless:     true,
	Generator: func(name string, args []*pipescript.Script) (*pipescript.TransformInitializer, error) {
		return &pipescript.TransformInitializer{Transform: tTransformStruct{}}, nil
	},
}
