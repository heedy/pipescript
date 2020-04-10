package strings

import (
	"errors"
	"regexp"
	"strings"

	"github.com/heedy/pipescript"
	"github.com/heedy/pipescript/resources"
)

func constString(consts []interface{}, pipes []*pipescript.Pipe) ([]interface{}, []*pipescript.Pipe, error) {
	c2 := make([]interface{}, 1)
	c, ok := pipescript.String(consts[0])
	if !ok {
		return nil, nil, errors.New("argument must be a string")
	}
	_, err := regexp.Compile(c)
	if err != nil {
		return nil, nil, err
	}
	c2[0] = c
	return c2, pipes, nil

}

var Contains = &pipescript.Transform{
	Name:          "contains",
	Description:   "Returns true if the given string is found in the datapoint string",
	Documentation: string(resources.MustAsset("docs/transforms/contains.md")),
	Args: []pipescript.TransformArg{
		{
			Description: "The substring to check for",
			Type:        pipescript.ConstArgType,
		},
	},
	Constructor: pipescript.NewBasic(constString, func(dp *pipescript.Datapoint, args []*pipescript.Datapoint, consts []interface{}, pipes []*pipescript.Pipe, out *pipescript.Datapoint) (*pipescript.Datapoint, error) {
		d, err := dp.String()
		if err != nil {
			out.Data = false
			return out, nil
		}

		out.Data = strings.Contains(d, consts[0].(string))
		return out, nil
	}),
}

var Startswith = &pipescript.Transform{
	Name:        "startswith",
	Description: "Returns true if datapoint string starts with the substring given in arg",
	Args: []pipescript.TransformArg{
		{
			Description: "The substring to check for",
			Type:        pipescript.ConstArgType,
			Schema: map[string]interface{}{
				"type": "string",
			},
		},
	},
	InputSchema: map[string]interface{}{
		"type": "string",
	},
	OutputSchema: map[string]interface{}{
		"type": "boolean",
	},
	Constructor: pipescript.NewBasic(constString, func(dp *pipescript.Datapoint, args []*pipescript.Datapoint, consts []interface{}, pipes []*pipescript.Pipe, out *pipescript.Datapoint) (*pipescript.Datapoint, error) {
		d, err := dp.String()
		if err != nil {
			out.Data = false
			return out, nil
		}

		out.Data = strings.HasPrefix(d, consts[0].(string))
		return out, nil
	}),
}

var Endswith = &pipescript.Transform{
	Name:        "endswith",
	Description: "Returns true if datapoint string ends with the substring given in arg",
	Args: []pipescript.TransformArg{
		{
			Description: "The substring to check for",
			Type:        pipescript.ConstArgType,
			Schema: map[string]interface{}{
				"type": "string",
			},
		},
	},
	InputSchema: map[string]interface{}{
		"type": "string",
	},
	OutputSchema: map[string]interface{}{
		"type": "boolean",
	},
	Constructor: pipescript.NewBasic(constString, func(dp *pipescript.Datapoint, args []*pipescript.Datapoint, consts []interface{}, pipes []*pipescript.Pipe, out *pipescript.Datapoint) (*pipescript.Datapoint, error) {
		d, err := dp.String()
		if err != nil {
			out.Data = false
			return out, nil
		}

		out.Data = strings.HasSuffix(d, consts[0].(string))
		return out, nil
	}),
}
