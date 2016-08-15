package core

import (
	"errors"

	"github.com/connectordb/pipescript"
	"github.com/connectordb/pipescript/resources"
)

type iwhileTransformStruct struct {
	script *pipescript.Script
	iter   *pipescript.SingleDatapointIterator
}

func (t iwhileTransformStruct) Copy() (pipescript.TransformInstance, error) {
	news, err := t.script.Copy()
	if err != nil {
		return nil, err
	}
	iter := &pipescript.SingleDatapointIterator{}
	news.SetInput(iter)
	return iwhileTransformStruct{news, iter}, nil
}

func (t iwhileTransformStruct) Next(ti *pipescript.TransformIterator) (dp *pipescript.Datapoint, err error) {
	te := ti.Next()
	if te.IsFinished() {
		// Clear the internal script
		t.iter.Set(nil, nil)
		t.script.Next()
		return te.Get()
	}

	v, err := te.Args[0].Bool()
	if err != nil {
		return nil, err
	}
	if !v {
		// Reset the script
		t.iter.Set(nil, nil)
		t.script.Next()
	}
	t.iter.Set(te.Datapoint, nil)
	return t.script.Next()

}

// IWhile is same as while, but returns all intermediate values
var IWhile = pipescript.Transform{
	Name:          "iwhile",
	Description:   "Equivalent to a while loop that runs while the first argument is true. Restarts the loop when the argument is false. The i version returns all intermediate results.",
	Documentation: string(resources.MustAsset("docs/transforms/while.md")),
	OneToOne:      true,
	Args: []pipescript.TransformArg{
		{
			Description: "The statement to check for truth",
		},
		{
			Description: "pipe to run, and to reset when the first arg is false",
			Hijacked:    true,
		},
	},
	Generator: func(name string, args []*pipescript.Script) (*pipescript.TransformInitializer, error) {
		if args[1].Peek {
			return nil, errors.New("iwhile cannot be used with transforms that peek")
		}

		iter := &pipescript.SingleDatapointIterator{}
		args[1].SetInput(iter)
		return &pipescript.TransformInitializer{Args: args[0:1], Transform: iwhileTransformStruct{args[1], iter}}, nil
	},
}
