package pipescript

import "fmt"

// TransformGenerator creates a new TransformInstance from the generator. The args given are only the
// known constants - all other elements are nil
type TransformGenerator func(name string, args []*Datapoint) (TransformInstance, error)

// TransformArg represents an argument passed into the transform function
type TransformArg struct {
	Description string      `json:"description"`       // A description of what the arg represents
	Optional    bool        `json:"optional"`          // Whether the arg is optional
	Default     interface{} `json:"default,omitempty"` // If the arg is optional, what is its default value
	Constant    bool        `json:"constant"`          // If the argument must be a constant (ie, not part of a transform)
}

// Transform is the struct which holds the name, docstring, and generator for a transform function
type Transform struct {
	Name         string         `json:"name"`              // The name of the transform
	Description  string         `json:"description"`       // The description of the transform - a docstring
	InputSchema  string         `json:"ischema,omitempty"` // The schema of the input datapoint that the given transform expects (optional)
	OutputSchema string         `json:"oschema,omitempty"` // The schema of the output data that this transform gives (optional).
	Args         []TransformArg `json:"args"`              // The arguments that the transform accepts
	OneToOne     bool           `json:"one_to_one"`        //Whether or not the transform is one to one

	Generator TransformGenerator `json:"-"` // The generator function of the transform
}

var (
	// TransformRegistry is the map of all the transforms that are currently registered.
	// Do not manually add/remove elements from this map.
	// Use Transform.Register to insert new transforms.
	TransformRegistry = make(map[string]Transform)
)

// Register registers the transform with the system. Note that it is not threadsafe. Register is
// assumed to be run once at the startup of the query system. Adding functions during runtime is
// not supported.
func (t Transform) Register() error {
	if t.Name == "" || t.Generator == nil {
		err := fmt.Errorf("Attempted to register invalid transform: '%s'", t.Name)
		return err
	}
	_, ok := TransformRegistry[t.Name]
	if ok {
		err := fmt.Errorf("A transform with the name '%s' already exists.", t.Name)
		return err
	}

	TransformRegistry[t.Name] = t

	return nil
}
