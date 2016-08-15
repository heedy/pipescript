package core

import (
	"errors"

	"github.com/connectordb/pipescript"
	"github.com/connectordb/pipescript/resources"
)

type whileTransformStruct struct {
	script *pipescript.Script
	iter   *pipescript.SingleDatapointIterator
}

func (t whileTransformStruct) Copy() (pipescript.TransformInstance, error) {
	news, err := t.script.Copy()
	if err != nil {
		return nil, err
	}
	iter := &pipescript.SingleDatapointIterator{}
	news.SetInput(iter)
	return whileTransformStruct{news, iter}, nil
}

// Next in the whileTransform works by peeking forward one datapoint. If the *next* argument is false, it means
// that we return the result of the current datapoint. If not,
func (t whileTransformStruct) Next(ti *pipescript.TransformIterator) (dp *pipescript.Datapoint, err error) {
	// Reset the internal script
	t.iter.Set(nil, nil)
	t.script.Next()

	// When the code gets here, at all times we are either on the first datapoint or on a false, or at the end of the stream
	for {

		te := ti.Next()
		if te.IsFinished() {
			return te.Get()
		}
		// Add the current datapoint to the script.
		t.iter.Set(te.Datapoint, nil)
		dp, err = t.script.Next()
		if err != nil {
			return dp, err
		}

		// Check the next datapoint. If it is false, or the end of the stream, return the value
		te = ti.Peek(0)
		if te.IsFinished() {
			return dp, nil
		}

		v, err := te.Args[0].Bool()
		if err != nil {
			return nil, err
		}

		if !v {
			return dp, nil
		}
	}

}

// While performs a while loop
var While = pipescript.Transform{
	Name:          "while",
	Description:   "Equivalent to a while loop that runs while the first argument is true. Restarts the loop when the argument is false.",
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
			return nil, errors.New("while cannot be used with transforms that peek")
		}

		iter := &pipescript.SingleDatapointIterator{}
		args[1].SetInput(iter)
		return &pipescript.TransformInitializer{Args: args[0:1], Transform: whileTransformStruct{args[1], iter}}, nil
	},
}
