package misc

import (
	"errors"

	"github.com/heedy/pipescript"
)

var Length = &pipescript.Transform{
	Name:        "length",
	Description: "Returns the length of an array/string, or number of object keys.",
	Constructor: pipescript.NewBasic(nil, func(dp *pipescript.Datapoint, args []*pipescript.Datapoint, consts []interface{}, pipes []*pipescript.Pipe, out *pipescript.Datapoint) (*pipescript.Datapoint, error) {
		switch v := dp.Data.(type) {
		case string:
			out.Data = len(v)
		case []interface{}:
			out.Data = len(v)
		case map[string]interface{}:
			out.Data = len(v)
		default:
			return nil, errors.New("Could not find length of data")
		}

		return out, nil
	}),
}
