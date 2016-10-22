package core

import (
	"fmt"
	"math"

	"github.com/connectordb/pipescript"
	"github.com/connectordb/pipescript/resources"
)

type bucketTransform struct {
	bucketstart float64
	bucketsize  float64
}

func (t bucketTransform) Copy() (pipescript.TransformInstance, error) {
	return bucketTransform{
		bucketsize:  t.bucketsize,
		bucketstart: t.bucketstart,
	}, nil
}

func (t bucketTransform) Next(ti *pipescript.TransformIterator) (*pipescript.Datapoint, error) {
	te := ti.Next()
	if te.IsFinished() {
		return te.Get()
	}
	f, err := te.Datapoint.Float()
	if err != nil {
		return nil, err
	}

	// The current bucket number
	bucketnum := math.Floor((f - t.bucketstart) / t.bucketsize)
	// We avoid a bit of floating point issues when doing it this way
	bucketStartLocation := t.bucketstart + bucketnum*t.bucketsize
	bucketEndLocation := t.bucketstart + (bucketnum+1)*t.bucketsize
	return te.Set(fmt.Sprintf("[%g,%g)", bucketStartLocation, bucketEndLocation))
}

var Bucket = pipescript.Transform{
	Name:          "bucket",
	Description:   "Puts numbers into custom-sized buckets. Useful for histograms.",
	Documentation: string(resources.MustAsset("docs/transforms/bucket.md")),
	OutputSchema:  `{"type": "string"}`,
	OneToOne:      true,
	Stateless:     true,
	Args: []pipescript.TransformArg{
		{
			Description: "The size of each bucket (float)",
			Optional:    true,
			Default:     10,
			Constant:    true,
		},
		{
			Description: "Start location for bucketing",
			Optional:    true,
			Default:     0,
			Constant:    true,
		},
	},

	Generator: func(name string, args []*pipescript.Script) (*pipescript.TransformInitializer, error) {
		dp, err := args[0].GetConstant()
		if err != nil {
			return nil, err
		}
		bucketsize, err := dp.Float()
		if err != nil {
			return nil, err
		}
		dp, err = args[1].GetConstant()
		if err != nil {
			return nil, err
		}
		bucketstart, err := dp.Float()
		return &pipescript.TransformInitializer{Transform: &bucketTransform{
			bucketstart: bucketstart,
			bucketsize:  bucketsize,
		}}, err
	},
}
