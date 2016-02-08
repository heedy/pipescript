package core

import (
	"fmt"

	"github.com/connectordb/pipescript"
)

type ifTransformStruct struct{}

func (t ifTransformStruct) Copy() (pipescript.TransformInstance, error) {
	return ifTransformStruct{}, nil
}

func (t ifTransformStruct) Next(ti *pipescript.TransformIterator) (dp *pipescript.Datapoint, err error) {
	var te *pipescript.TransformEnvironment

	//While the if statement is false, loop throguh values
	v := false
	for !v {
		te = ti.Next()
		if te.IsFinished() {
			return te.Get()
		}
		v, err = te.Args[0].Bool()
		if err != nil {
			return nil, fmt.Errorf("Argument given to if can't be converted to boolean")
		}
	}

	// If we got here, the if statement was true - so we pass through our datapoint
	return te.Get()
}

var If = pipescript.Transform{
	Name:        "if",
	Description: "A datapoint filter - filters the datapoints where its argument is false. This is PipeScript's if statement.",
	Args: []pipescript.TransformArg{
		{
			Description: "The statement to check for truth",
		},
	},
	Generator: func(name string, args []*pipescript.Script) (*pipescript.TransformInitializer, error) {
		return &pipescript.TransformInitializer{Args: args, Transform: ifTransformStruct{}}, nil
	},
}
