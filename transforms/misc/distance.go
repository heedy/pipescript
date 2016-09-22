package misc

import (
	"errors"
	"math"

	"github.com/connectordb/duck"
	"github.com/connectordb/pipescript"
	"github.com/connectordb/pipescript/resources"
)

type distanceTransform struct {
	lat  float64
	long float64
}

func (t *distanceTransform) Copy() (pipescript.TransformInstance, error) {
	return &distanceTransform{t.lat, t.long}, nil
}

// EarthRadius is the earth radius in meters
var EarthRadius = float64(6371000)

// Radians is the multiplication cosntant to convert degrees to radians
var Radians = math.Pi / 180.0

func (t *distanceTransform) Next(ti *pipescript.TransformIterator) (*pipescript.Datapoint, error) {
	te := ti.Next()
	if te.IsFinished() {
		return te.Get()
	}

	s, err := te.Datapoint.Get("latitude")
	if err != nil {
		return nil, err
	}
	lat, ok := duck.Float(s)
	if !ok {
		return nil, errors.New("Could not convert latitude to float")
	}
	s, err = te.Datapoint.Get("longitude")
	if err != nil {
		return nil, err
	}
	long, ok := duck.Float(s)
	if !ok {
		return nil, errors.New("Could not convert longitude to float")
	}

	// https://en.wikipedia.org/wiki/Haversine_formula
	// http://www.movable-type.co.uk/scripts/latlong.html

	// Convert our distances to Radians
	lat = lat * Radians
	long = long * Radians

	// The radian distances between chosen point and datapoint
	dlat := lat - t.lat
	dlong := long - t.long

	// Now we compute the distance using haverside formula
	a := math.Sin(dlat/2)*math.Sin(dlat/2) +
		math.Cos(lat)*math.Cos(t.lat)*math.Sin(dlong/2)*math.Sin(dlong/2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return te.Set(EarthRadius * c)
}

var Distance = pipescript.Transform{
	Name:          "distance",
	Description:   "Returns distance in meters from given latitude/longitude coordinates to datapoint",
	OutputSchema:  `{"type": "boolean"}`,
	Documentation: string(resources.MustAsset("docs/transforms/distance.md")),
	OneToOne:      true,
	Stateless:     true,
	Args: []pipescript.TransformArg{
		{
			Description: "Latitude",
			Constant:    true,
		},
		{
			Description: "Longitude",
			Constant:    true,
		},
	},
	Generator: func(name string, args []*pipescript.Script) (*pipescript.TransformInitializer, error) {
		dp, err := args[0].GetConstant()
		if err != nil {
			return nil, err
		}
		lat, err := dp.Float()
		if err != nil {
			return nil, err
		}
		dp, err = args[1].GetConstant()
		if err != nil {
			return nil, err
		}
		long, err := dp.Float()
		if err != nil {
			return nil, err
		}
		return &pipescript.TransformInitializer{Transform: &distanceTransform{lat * Radians, long * Radians}}, nil
	},
}
