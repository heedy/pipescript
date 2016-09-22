package core

import (
	"fmt"

	"github.com/connectordb/pipescript"
	"github.com/connectordb/pipescript/resources"
)

type newTransform struct {
	seendata map[string]bool
}

func (t *newTransform) Copy() (pipescript.TransformInstance, error) {
	seendata := make(map[string]bool)
	for i, _ := range t.seendata {
		seendata[i] = true
	}
	return &newTransform{seendata}, nil
}

func (t *newTransform) Next(ti *pipescript.TransformIterator) (*pipescript.Datapoint, error) {
	te := ti.Next()
	if te.IsFinished() {
		// Clear the seen datapoints
		t.seendata = make(map[string]bool)

		return te.Get()
	}

	// Convert the key value to string
	v, err := te.Datapoint.DataString()
	if err != nil {
		return nil, err
	}

	//Check if the value exists
	_, ok := t.seendata[v]
	if !ok {
		if len(t.seendata) >= SplitMax {
			return nil, fmt.Errorf("Reached maximum split amount %d.", SplitMax)
		}

		t.seendata[v] = true
	}

	return te.Set(!ok)
}

// New checks if the datapoint was not yet seen
var New = pipescript.Transform{
	Name:          "new",
	Description:   `Returns true only when the given data was not yet seen.`,
	OutputSchema:  `{"type": "boolean"}`,
	Documentation: string(resources.MustAsset("docs/transforms/new.md")),
	OneToOne:      true,
	Generator: func(name string, args []*pipescript.Script) (*pipescript.TransformInitializer, error) {
		seendata := make(map[string]bool)
		return &pipescript.TransformInitializer{
			Transform: &newTransform{seendata},
		}, nil
	},
}
