package misc

import (
	"errors"
	"math"

	"github.com/heedy/pipescript"
	"github.com/heedy/pipescript/resources"
)

// EarthRadius is the earth radius in meters
var EarthRadius = float64(6371000)

// Radians is the multiplication constant to convert degrees to radians
var Radians = math.Pi / 180.0

var Distance = &pipescript.Transform{
	Name:          "distance",
	Description:   "Returns distance in meters from given latitude/longitude coordinates to datapoint",
	Documentation: string(resources.MustAsset("docs/transforms/distance.md")),
	Args: []pipescript.TransformArg{
		{
			Description: "Latitude",
			Type:        pipescript.ConstArgType,
			Schema: map[string]interface{}{
				"type": "number",
			},
		},
		{
			Description: "Longitude",
			Type:        pipescript.ConstArgType,
			Schema: map[string]interface{}{
				"type": "number",
			},
		},
	},
	InputSchema: map[string]interface{}{
		"type": "object",
		"properties": map[string]interface{}{
			"latitude": map[string]interface{}{
				"type": "number",
			},
			"longitude": map[string]interface{}{
				"type": "number",
			},
		},
		"required": []string{"latitude", "longitude"},
	},
	OutputSchema: map[string]interface{}{
		"type": "number",
	},
	Constructor: pipescript.NewBasic(func(consts []interface{}, pipes []*pipescript.Pipe) ([]interface{}, []*pipescript.Pipe, error) {
		c2 := make([]interface{}, 2)
		f, ok := pipescript.Float(consts[0])
		if !ok {
			return nil, nil, errors.New("latitude must be a number")
		}
		c2[0] = f * Radians
		f, ok = pipescript.Float(consts[1])
		if !ok {
			return nil, nil, errors.New("longitude must be a number")
		}
		c2[1] = f * Radians
		return c2, pipes, nil
	}, func(dp *pipescript.Datapoint, args []*pipescript.Datapoint, consts []interface{}, pipes []*pipescript.Pipe, out *pipescript.Datapoint) (*pipescript.Datapoint, error) {
		data, ok := dp.Data.(map[string]interface{})
		if !ok {
			return nil, errors.New("Distance can only be found to objects with latitude and longitude keys")
		}
		v, ok := data["latitude"]
		if !ok {
			return nil, errors.New("Could not find latitude in datapoint")
		}
		lat, ok := v.(float64)
		if !ok {
			return nil, errors.New("Latitude must be a number")
		}

		v, ok = data["longitude"]
		if !ok {
			return nil, errors.New("Could not find longitude in datapoint")
		}
		long, ok := v.(float64)
		if !ok {
			return nil, errors.New("Longitude must be a number")
		}

		// Convert our distances to Radians
		lat = lat * Radians
		long = long * Radians

		arglat := consts[0].(float64)
		arglong := consts[1].(float64)

		// The radian distances between chosen point and datapoint
		dlat := lat - arglat
		dlong := long - arglong

		// Now we compute the distance using haverside formula
		a := math.Sin(dlat/2)*math.Sin(dlat/2) +
			math.Cos(lat)*math.Cos(arglat)*math.Sin(dlong/2)*math.Sin(dlong/2)

		c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
		out.Data = EarthRadius * c

		return out, nil
	}),
}
