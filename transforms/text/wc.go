package text

import (
	"github.com/connectordb/pipescript"
	"github.com/connectordb/pipescript/resources"
)

type wcTransform struct {
}

func (t *wcTransform) Copy() (pipescript.TransformInstance, error) {
	return &wcTransform{}, nil
}

func (t *wcTransform) Next(ti *pipescript.TransformIterator) (*pipescript.Datapoint, error) {
	te := ti.Next()
	if te.IsFinished() {
		return te.Get()
	}
	s, err := te.Datapoint.DataString()
	if err != nil {
		return nil, err
	}

	return te.Set(len(Tokenize(s)))
}

var Wc = pipescript.Transform{
	Name:          "wc",
	Description:   "Returns the number of words in the given text",
	Documentation: string(resources.MustAsset("docs/transforms/wc.md")),
	OutputSchema:  `{"type": "integer"}`,
	InputSchema:   `{"type": "string"}`,
	OneToOne:      true,
	Stateless:     true,
	Generator: func(name string, args []*pipescript.Script) (*pipescript.TransformInitializer, error) {
		return &pipescript.TransformInitializer{Transform: &wcTransform{}}, nil
	},
}
