package pipescript

var NotTransform = &Transform{
	Name:        "not",
	Description: "Boolean not",

	Constructor: NewBasic(nil, func(dp *Datapoint, args []*Datapoint, consts []interface{}, pipes []*Pipe, out *Datapoint) (*Datapoint, error) {
		b, err := dp.Bool()
		out.Data = !b
		return out, err
	}),
}

var AndTransform = &Transform{
	Name:        "and",
	Description: "logical and",

	Args: []TransformArg{
		TransformArg{
			Description: "Value to and against data",
			Type:        TransformArgType,
		},
		TransformArg{
			Description: "Value to and against data",
			Type:        TransformArgType,
		},
	},
	Constructor: NewArgBasic(func(args []*Datapoint, consts []interface{}, pipes []*Pipe, out *Datapoint) (*Datapoint, error) {
		f1, err := args[0].Bool()
		if err == nil {
			var f2 bool
			f2, err = args[1].Bool()

			out.Data = f1 && f2
		}
		return out, err
	}),
}

var OrTransform = &Transform{
	Name:        "or",
	Description: "logical or",

	Args: []TransformArg{
		TransformArg{
			Description: "Value to or against data",
			Type:        TransformArgType,
		},
		TransformArg{
			Description: "Value to and against data",
			Type:        TransformArgType,
		},
	},
	Constructor: NewArgBasic(func(args []*Datapoint, consts []interface{}, pipes []*Pipe, out *Datapoint) (*Datapoint, error) {
		f1, err := args[0].Bool()
		if err == nil {
			var f2 bool
			f2, err = args[1].Bool()

			out.Data = f1 || f2
		}
		return out, err
	}),
}
