package pipescript

var LtTransform = &Transform{
	Name:        "lt",
	Description: "returns true if the data of the incoming stream is less than the value of the first arg",

	Args: []TransformArg{
		TransformArg{
			Description: "Value to check against data",
			Type:        TransformArgType,
		},
		TransformArg{
			Description: "Value to check against data",
			Type:        TransformArgType,
		},
	},
	Constructor: NewArgBasic(func(args []*Datapoint, consts []interface{}, pipes []*Pipe, out *Datapoint) (*Datapoint, error) {
		f1, err := args[0].Float()
		if err == nil {
			var f2 float64
			f2, err = args[1].Float()

			out.Data = f1 < f2
		}
		return out, err
	}),
}

var GtTransform = &Transform{
	Name:        "gt",
	Description: "returns true if the data of the incoming stream is greater than the value of the first arg",

	Args: []TransformArg{
		TransformArg{
			Description: "Value to check against data",
			Type:        TransformArgType,
		},
		TransformArg{
			Description: "Value to check against data",
			Type:        TransformArgType,
		},
	},
	Constructor: NewArgBasic(func(args []*Datapoint, consts []interface{}, pipes []*Pipe, out *Datapoint) (*Datapoint, error) {
		f1, err := args[0].Float()
		if err == nil {
			var f2 float64
			f2, err = args[1].Float()

			out.Data = f1 > f2
		}
		return out, err
	}),
}

var LteTransform = &Transform{
	Name:        "lte",
	Description: "returns true if the data of the incoming stream is less than or equal to the value of the first arg",

	Args: []TransformArg{
		TransformArg{
			Description: "Value to check against data",
			Type:        TransformArgType,
		},
		TransformArg{
			Description: "Value to check against data",
			Type:        TransformArgType,
		},
	},
	Constructor: NewArgBasic(func(args []*Datapoint, consts []interface{}, pipes []*Pipe, out *Datapoint) (*Datapoint, error) {
		f1, err := args[0].Float()
		if err == nil {
			var f2 float64
			f2, err = args[1].Float()

			out.Data = f1 <= f2
		}
		return out, err
	}),
}

var GteTransform = &Transform{
	Name:        "gte",
	Description: "returns true if the data of the incoming stream is greater than or equal to  the value of the first arg",

	Args: []TransformArg{
		TransformArg{
			Description: "Value to check against data",
			Type:        TransformArgType,
		},
		TransformArg{
			Description: "Value to check against data",
			Type:        TransformArgType,
		},
	},
	Constructor: NewArgBasic(func(args []*Datapoint, consts []interface{}, pipes []*Pipe, out *Datapoint) (*Datapoint, error) {
		f1, err := args[0].Float()
		if err == nil {
			var f2 float64
			f2, err = args[1].Float()

			out.Data = f1 >= f2
		}
		return out, err
	}),
}

var EqTransform = &Transform{
	Name:        "eq",
	Description: "returns true if the data of the incoming stream is not equal to the value of the first arg",

	Args: []TransformArg{
		TransformArg{
			Description: "Value to check against data",
			Type:        TransformArgType,
		},
		TransformArg{
			Description: "Value to check against data",
			Type:        TransformArgType,
		},
	},
	Constructor: NewArgBasic(func(args []*Datapoint, consts []interface{}, pipes []*Pipe, out *Datapoint) (*Datapoint, error) {
		out.Data = Equal(args[0].Data, args[1].Data)
		return out, nil
	}),
}

var NeTransform = &Transform{
	Name:        "ne",
	Description: "returns true if the data of the incoming stream is not equal to the value of the first arg",

	Args: []TransformArg{
		TransformArg{
			Description: "Value to check against data",
			Type:        TransformArgType,
		},
		TransformArg{
			Description: "Value to check against data",
			Type:        TransformArgType,
		},
	},
	Constructor: NewArgBasic(func(args []*Datapoint, consts []interface{}, pipes []*Pipe, out *Datapoint) (*Datapoint, error) {
		out.Data = !Equal(args[0].Data, args[1].Data)
		return out, nil
	}),
}
