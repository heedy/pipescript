package strings

import (
	"strings"

	"github.com/connectordb/pipescript"
	"github.com/connectordb/pipescript/resources"
)

type containsTransform struct {
	s string
}

func (t *containsTransform) Copy() (pipescript.TransformInstance, error) {
	return &containsTransform{t.s}, nil
}

func (t *containsTransform) Next(ti *pipescript.TransformIterator) (*pipescript.Datapoint, error) {
	te := ti.Next()
	if te.IsFinished() {
		return te.Get()
	}

	s, err := te.Datapoint.DataString()
	if err != nil {
		return nil, err
	}

	return te.Set(strings.Contains(s, t.s))
}

var Contains = pipescript.Transform{
	Name:          "contains",
	Description:   "Returns true if the given string is found in the datapoint string",
	OutputSchema:  `{"type": "boolean"}`,
	Documentation: string(resources.MustAsset("docs/transforms/contains.md")),
	OneToOne:      true,
	Stateless:     true,
	Args: []pipescript.TransformArg{
		{
			Description: "The string to search for",
			Constant:    true,
		},
	},
	Generator: func(name string, args []*pipescript.Script) (*pipescript.TransformInitializer, error) {
		dp, err := args[0].GetConstant()
		if err != nil {
			return nil, err
		}
		s, err := dp.DataString()
		if err != nil {
			return nil, err
		}
		return &pipescript.TransformInitializer{Transform: &containsTransform{s}}, nil
	},
}
