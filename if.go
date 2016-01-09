package pipescript

import "fmt"

type ifTransformStruct struct{}

func (t ifTransformStruct) Copy() (TransformInstance, error) {
	return ifTransformStruct{}, nil
}

func (t ifTransformStruct) Next(ti *TransformIterator) (dp *Datapoint, err error) {
	var te *TransformEnvironment

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

var ifTransform = Transform{
	Name:        "if",
	Description: "A datapoint filter - filters the datapoints where its argument is false. This is PipeScript's if statement.",
	Args: []TransformArg{
		{
			Description: "The statement to check for truth",
		},
	},
	Generator: func(name string, args []*Script) (*TransformInitializer, error) {
		return &TransformInitializer{Args: args, Transform: ifTransformStruct{}}, nil
	},
}
