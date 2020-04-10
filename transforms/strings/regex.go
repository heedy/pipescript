package strings

import (
	"errors"
	"regexp"

	"github.com/heedy/pipescript"
	"github.com/heedy/pipescript/resources"
)

var Regex = &pipescript.Transform{
	Name:          "regex",
	Description:   "Returns true if the given regular expression matches the data string",
	Documentation: string(resources.MustAsset("docs/transforms/regex.md")),
	InputSchema: map[string]interface{}{
		"type": "string",
	},
	OutputSchema: map[string]interface{}{
		"type": "boolean",
	},
	Args: []pipescript.TransformArg{
		{
			Description: "The regular expression to use",
			Type:        pipescript.ConstArgType,
			Schema: map[string]interface{}{
				"type": "string",
			},
		},
	},
	Constructor: pipescript.NewBasic(func(consts []interface{}, pipes []*pipescript.Pipe) ([]interface{}, []*pipescript.Pipe, error) {
		c2 := make([]interface{}, 1)
		c, ok := pipescript.String(consts[0])
		if !ok {
			return nil, nil, errors.New("regex must be a string")
		}
		_, err := regexp.Compile(c)
		if err != nil {
			return nil, nil, err
		}
		c2[0] = c
		return c2, pipes, nil

	}, func(dp *pipescript.Datapoint, args []*pipescript.Datapoint, consts []interface{}, pipes []*pipescript.Pipe, out *pipescript.Datapoint) (*pipescript.Datapoint, error) {
		d, err := dp.String()
		if err != nil {
			out.Data = false
			return out, nil
		}
		re, _ := regexp.Compile(consts[0].(string))
		out.Data = re.MatchString(d)
		return out, nil
	}),
}
