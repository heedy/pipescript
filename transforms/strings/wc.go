package strings

import (
	"strings"

	"github.com/heedy/pipescript"
	"github.com/heedy/pipescript/resources"
)

var Wc = &pipescript.Transform{
	Name:          "wc",
	Description:   "Returns the number of words in the given string",
	Documentation: string(resources.MustAsset("docs/transforms/wc.md")),
	InputSchema: map[string]interface{}{
		"type": "string",
	},
	OutputSchema: map[string]interface{}{
		"type": "number",
	},
	Constructor: pipescript.NewBasic(nil, func(dp *pipescript.Datapoint, args []*pipescript.Datapoint, consts []interface{}, pipes []*pipescript.Pipe, out *pipescript.Datapoint) (*pipescript.Datapoint, error) {
		ds, err := dp.String()
		if err != nil {
			return nil, err
		}
		out.Data = len(strings.Fields(ds))

		return out, nil
	}),
}
