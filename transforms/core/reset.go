package core

import (
	"errors"

	"github.com/connectordb/pipescript"
)

type resetTransformStruct struct {
	script *pipescript.Script
	iter   *pipescript.SingleDatapointIterator
}

func (t resetTransformStruct) Copy() (pipescript.TransformInstance, error) {
	news, err := t.script.Copy()
	if err != nil {
		return nil, err
	}
	iter := &pipescript.SingleDatapointIterator{}
	news.SetInput(iter)
	return resetTransformStruct{news, iter}, nil
}

func (t resetTransformStruct) Next(ti *pipescript.TransformIterator) (dp *pipescript.Datapoint, err error) {
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
	if v {
		// Reset the script
		t.iter.Set(nil, nil)
		t.script.Next()
	}
	t.iter.Set(te.Datapoint, nil)
	return t.script.Next()

}

var Reset = pipescript.Transform{
	Name:        "reset",
	Description: "Resets the internal state of its second argument when its first argument is true. Returns value of second argument.",
	OneToOne:    true,
	Args: []pipescript.TransformArg{
		{
			Description: "The statement to check for truth",
		},
		{
			Description: "pipe to run, and to reset when the first arg is true",
			Hijacked:    true,
		},
	},
	Generator: func(name string, args []*pipescript.Script) (*pipescript.TransformInitializer, error) {
		if args[1].Peek {
			return nil, errors.New("reset cannot be used with transforms that peek")
		}

		iter := &pipescript.SingleDatapointIterator{}
		args[1].SetInput(iter)
		return &pipescript.TransformInitializer{Args: args[0:1], Transform: resetTransformStruct{args[1], iter}}, nil
	},
}
