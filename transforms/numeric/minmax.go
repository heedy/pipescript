package numeric

import "github.com/heedy/pipescript"

var Min = &pipescript.Transform{
	Name:        "min",
	Description: "Returns the minimum datapoint in the timeseries",
	Args: []pipescript.TransformArg{
		{
			Description: "Optional, if this is set, then the result of this argument is checked for min",
			Optional:    true,
			Default:     pipescript.IdentityPipe,
			Type:        pipescript.TransformArgType,
		},
	},
	Constructor: pipescript.NewAggregator(func(e *pipescript.TransformEnv, consts []interface{}, pipes []*pipescript.Pipe, out *pipescript.Datapoint) (*pipescript.Datapoint, error) {
		args := make([]*pipescript.Datapoint, 1)
		dp, args, err := e.Next(args)
		if err != nil || dp == nil {
			return nil, err
		}
		curmin := dp
		f, err := args[0].Float()
		if err != nil {
			return nil, err
		}
		curval := f
		for {
			dp, args, err = e.Next(args)
			if err != nil {
				return nil, err
			}
			if dp == nil {
				break
			}
			f, err := args[0].Float()
			if err != nil {
				return nil, err
			}
			if f < curval {
				curmin = dp
				curval = f
			}

		}

		out.Data = curmin.Data
		out.Timestamp = curmin.Timestamp
		out.Duration = curmin.Duration
		return out, nil
	}),
}

var Max = &pipescript.Transform{
	Name:        "max",
	Description: "Returns the maximum datapoint in the timeseries",
	Args: []pipescript.TransformArg{
		{
			Description: "Optional, if this is set, then the result of this argument is checked for min",
			Optional:    true,
			Default:     pipescript.IdentityPipe,
			Type:        pipescript.TransformArgType,
		},
	},
	Constructor: pipescript.NewAggregator(func(e *pipescript.TransformEnv, consts []interface{}, pipes []*pipescript.Pipe, out *pipescript.Datapoint) (*pipescript.Datapoint, error) {
		args := make([]*pipescript.Datapoint, 1)
		dp, args, err := e.Next(args)
		if err != nil || dp == nil {
			return nil, err
		}
		curmax := dp
		f, err := args[0].Float()
		if err != nil {
			return nil, err
		}
		curval := f
		for {
			dp, args, err = e.Next(args)
			if err != nil {
				return nil, err
			}
			if dp == nil {
				break
			}
			f, err := args[0].Float()
			if err != nil {
				return nil, err
			}
			if f > curval {
				curmax = dp
				curval = f
			}

		}

		out.Data = curmax.Data
		out.Timestamp = curmax.Timestamp
		out.Duration = curmax.Duration
		return out, nil
	}),
}
