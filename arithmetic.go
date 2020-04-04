package pipescript

import "math"

var NegTransform = &Transform{
	Name:        "neg",
	Description: "Negation of numbers",

	Constructor: NewBasic(nil, func(dp *Datapoint, args []*Datapoint, consts []interface{}, pipes []*Pipe, out *Datapoint) (*Datapoint, error) {
		b, ok := dp.Data.(bool)
		if ok {
			out.Data = !b
			return out, nil
		}
		f, err := dp.Float()
		out.Data = -f
		return out, err
	}),
}

var AddTransform = &Transform{
	Name:        "add",
	Description: "add",

	Args: []TransformArg{
		TransformArg{
			Description: "Value to add to the datapoint",
			Type:        TransformArgType,
		},
		TransformArg{
			Description: "Value to add to the datapoint",
			Type:        TransformArgType,
		},
	},
	Constructor: NewArgBasic(func(args []*Datapoint, consts []interface{}, pipes []*Pipe, out *Datapoint) (*Datapoint, error) {
		f1, err := args[0].Float()
		if err == nil {
			var f2 float64
			f2, err = args[1].Float()

			out.Data = f1 + f2
		}
		return out, err
	}),
}

var SubTransform = &Transform{
	Name:        "sub",
	Description: "subtracts first argument from the datapoint",

	Args: []TransformArg{
		TransformArg{
			Description: "Value to subtract from the datapoint",
			Type:        TransformArgType,
		},
		TransformArg{
			Description: "Value to add to the datapoint",
			Type:        TransformArgType,
		},
	},
	Constructor: NewArgBasic(func(args []*Datapoint, consts []interface{}, pipes []*Pipe, out *Datapoint) (*Datapoint, error) {
		f1, err := args[0].Float()
		if err == nil {
			var f2 float64
			f2, err = args[1].Float()

			out.Data = f1 - f2
		}
		return out, err
	}),
}

var MulTransform = &Transform{
	Name:        "sub",
	Description: "multiplies datapoint by arg",

	Args: []TransformArg{
		TransformArg{
			Description: "Value to multiply",
			Type:        TransformArgType,
		},
		TransformArg{
			Description: "Value to add to the datapoint",
			Type:        TransformArgType,
		},
	},
	Constructor: NewArgBasic(func(args []*Datapoint, consts []interface{}, pipes []*Pipe, out *Datapoint) (*Datapoint, error) {
		f1, err := args[0].Float()
		if err == nil {
			var f2 float64
			f2, err = args[1].Float()

			out.Data = f1 * f2
		}
		return out, err
	}),
}

var DivTransform = &Transform{
	Name:        "sub",
	Description: "divides datapoint by arg",

	Args: []TransformArg{
		TransformArg{
			Description: "denominator",
			Type:        TransformArgType,
		},
		TransformArg{
			Description: "Value to add to the datapoint",
			Type:        TransformArgType,
		},
	},
	Constructor: NewArgBasic(func(args []*Datapoint, consts []interface{}, pipes []*Pipe, out *Datapoint) (*Datapoint, error) {
		f1, err := args[0].Float()
		if err == nil {
			var f2 float64
			f2, err = args[1].Float()

			out.Data = f1 / f2
		}
		return out, err
	}),
}

var ModTransform = &Transform{
	Name:        "mod",
	Description: "modulo",

	Args: []TransformArg{
		TransformArg{
			Description: "mod by this",
			Type:        TransformArgType,
		},
		TransformArg{
			Description: "Value to add to the datapoint",
			Type:        TransformArgType,
		},
	},
	Constructor: NewArgBasic(func(args []*Datapoint, consts []interface{}, pipes []*Pipe, out *Datapoint) (*Datapoint, error) {
		f1, err := args[0].Int()
		if err == nil {
			var f2 int64
			f2, err = args[1].Int()

			out.Data = f1 % f2
		}
		return out, err
	}),
}

var PowTransform = &Transform{
	Name:        "pow",
	Description: "sets datapoint to arg's power",

	Args: []TransformArg{
		TransformArg{
			Description: "Value to multiply",
			Type:        TransformArgType,
		},
		TransformArg{
			Description: "Value to add to the datapoint",
			Type:        TransformArgType,
		},
	},
	Constructor: NewArgBasic(func(args []*Datapoint, consts []interface{}, pipes []*Pipe, out *Datapoint) (*Datapoint, error) {
		f1, err := args[0].Float()
		if err == nil {
			var f2 float64
			f2, err = args[1].Float()

			out.Data = math.Pow(f1, f2)
		}
		return out, err
	}),
}
