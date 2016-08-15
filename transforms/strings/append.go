package strings

import (
	"fmt"

	"github.com/connectordb/pipescript"
)

var StringMax = 1000000

type appendTransform struct {
	s string
}

func (t *appendTransform) Copy() (pipescript.TransformInstance, error) {
	return &appendTransform{t.s}, nil
}

func (t *appendTransform) Next(ti *pipescript.TransformIterator) (*pipescript.Datapoint, error) {
	te := ti.Next()
	if te.IsFinished() {
		t.s = ""
		return te.Get()
	}

	s, err := te.Datapoint.DataString()
	if err != nil {
		return nil, err
	}
	t.s = t.s + s
	if len(t.s) > StringMax {
		return nil, fmt.Errorf("Reached maximum string length %d.", StringMax)
	}
	return te.Set(t.s)
}

var Append = pipescript.Transform{
	Name:         "append",
	Description:  "Appends data into one large string",
	OutputSchema: `{"type": "string"}`,
	OneToOne:     true,
	Generator: func(name string, args []*pipescript.Script) (*pipescript.TransformInitializer, error) {
		return &pipescript.TransformInitializer{Transform: &appendTransform{""}}, nil
	},
}
