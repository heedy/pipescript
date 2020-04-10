package numeric

import (
	"errors"
	"fmt"
	"math"

	"github.com/heedy/pipescript"
	"github.com/heedy/pipescript/resources"
)

var Bucket = pipescript.Transform{
	Name:          "bucket",
	Description:   "Puts numbers into custom-sized buckets. Useful for histograms.",
	Documentation: string(resources.MustAsset("docs/transforms/bucket.md")),
	Args: []pipescript.TransformArg{
		{
			Description: "The size of each bucket",
			Optional:    true,
			Default:     pipescript.MustPipe(pipescript.NewConstTransform(10), nil),
			Type:        pipescript.ConstArgType,
			Schema: map[string]interface{}{
				"type": "number",
			},
		},
		{
			Description: "Start location for bucketing",
			Optional:    true,
			Default:     pipescript.MustPipe(pipescript.NewConstTransform(0), nil),
			Type:        pipescript.ConstArgType,
			Schema: map[string]interface{}{
				"type": "number",
			},
		},
	},
	InputSchema: map[string]interface{}{
		"type": "number",
	},
	OutputSchema: map[string]interface{}{
		"type": "string",
	},
	Constructor: pipescript.NewBasic(func(consts []interface{}, pipes []*pipescript.Pipe) ([]interface{}, []*pipescript.Pipe, error) {
		f, ok := pipescript.Float(consts[0])
		if !ok {
			return nil, nil, errors.New("bucket size must be a number")
		}
		consts[0] = f
		f, ok = pipescript.Float(consts[1])
		if !ok {
			return nil, nil, errors.New("bucket start location must be a number")
		}
		consts[1] = f
		return consts, pipes, nil
	}, func(dp *pipescript.Datapoint, args []*pipescript.Datapoint, consts []interface{}, pipes []*pipescript.Pipe, out *pipescript.Datapoint) (*pipescript.Datapoint, error) {
		f, err := dp.Float()
		if err != nil {
			return nil, err
		}
		bucketstart := consts[1].(float64)
		bucketsize := consts[0].(float64)
		// The current bucket number
		bucketnum := math.Floor((f - bucketstart) / bucketsize)
		// We avoid a bit of floating point issues when doing it this way
		bucketStartLocation := bucketstart + bucketnum*bucketsize
		bucketEndLocation := bucketstart + (bucketnum+1)*bucketsize
		out.Data = fmt.Sprintf("[%g,%g)", bucketStartLocation, bucketEndLocation)
		return out, nil
	}),
}
