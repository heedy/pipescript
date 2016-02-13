package stats

import "github.com/connectordb/pipescript"

type meanTransform struct {
	Sum float64
	Tot int64
}

func (t *meanTransform) Copy() (pipescript.TransformInstance, error) {
	return &meanTransform{t.Sum, t.Tot}, nil
}

func (t *meanTransform) Next(ti *pipescript.TransformIterator) (*pipescript.Datapoint, error) {
	te := ti.Next()
	if te.IsFinished() {
		t.Sum = 0
		t.Tot = 0
		return te.Get()
	}
	v, err := te.Datapoint.Float()
	if err != nil {
		return nil, err
	}
	t.Sum += v
	t.Tot++
	return te.Set(t.Sum / float64(t.Tot))
}

var Mean = pipescript.Transform{
	Name:        "mean",
	Description: "Returns the average of all datapoints that have passed through it",
	OneToOne:    true,
	Generator: func(name string, args []*pipescript.Script) (*pipescript.TransformInitializer, error) {
		return &pipescript.TransformInitializer{Transform: &meanTransform{0, 0}}, nil
	},
}
