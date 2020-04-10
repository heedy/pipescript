package strings

import (
	"net/url"
	"strings"

	"github.com/heedy/pipescript"
)

var Urldomain = &pipescript.Transform{
	Name:        "urldomain",
	Description: "Returns the domain name/host that is used in the given url",
	InputSchema: map[string]interface{}{
		"type": "string",
	},
	OutputSchema: map[string]interface{}{
		"type": "string",
	},
	Constructor: pipescript.NewBasic(nil, func(dp *pipescript.Datapoint, args []*pipescript.Datapoint, consts []interface{}, pipes []*pipescript.Pipe, out *pipescript.Datapoint) (*pipescript.Datapoint, error) {
		ds, err := dp.String()
		if err != nil {
			out.Data = ""
			return out, nil
		}
		u, err := url.Parse(ds)
		if err != nil {
			out.Data = ""
			return out, nil
		}
		out.Data = strings.TrimPrefix(u.Host, "www.")
		return out, nil
	}),
}
