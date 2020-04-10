package core

import (
	"github.com/heedy/pipescript"
	"github.com/heedy/pipescript/resources"
)

type rememberTransform struct {
	remembered *pipescript.Datapoint
}

func (t *rememberTransform) Copy() (pipescript.TransformInstance, error) {
	return &rememberTransform{t.remembered}, nil
}

func (t *rememberTransform) Next(ti *pipescript.TransformIterator) (*pipescript.Datapoint, error) {
	te := ti.Next()
	if te.IsFinished() {
		t.remembered = nil
		return te.Get()
	}
	result, err := te.Args[0].Bool()
	if err != nil {
		return nil, err
	}
	if t.remembered == nil || result {
		t.remembered = te.Args[1]
	}
	return t.remembered, nil
}

var Remember = pipescript.Transform{
	Name:          "remember",
	Description:   "Behaves as a single-datapoint memory cell, which is reset when its first argument is true.",
	Documentation: string(resources.MustAsset("docs/transforms/remember.md")),
	OneToOne:      true,
	Args: []pipescript.TransformArg{
		{
			Description: "The statement to check for truth. If true, it will remember the current datapoint",
		},
		{
			Description: "Optional, if this is set, then the result of this argument is stored instead of the datapoint",
			Optional:    true,
			Default:     nil, // THIS HAS TO BE SET IN init.go to identity
		},
	},
	Generator: func(name string, args []*pipescript.Script) (*pipescript.TransformInitializer, error) {
		return &pipescript.TransformInitializer{Args: args, Transform: &rememberTransform{nil}}, nil
	},
}
