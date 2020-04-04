package pipescript

import (
	"encoding/json"
	"fmt"
	"sync"
)

type Iterator interface {
	Next(*Datapoint) (*Datapoint, error)
}

// ArgType gives the types of args that a trasnform can accept
type ArgType int

const (
	ConstArgType ArgType = iota
	TransformArgType
	OneToOnePipeArgType
	PipeArgType
)

func (a ArgType) String() string {
	switch a {
	case ConstArgType:
		return "const"
	case TransformArgType:
		return "transform"
	case OneToOnePipeArgType:
		return "one_to_one_pipe"
	case PipeArgType:
		return "pipe"
	}
	return ""
}

func (a ArgType) MarshalJSON() ([]byte, error) {
	return json.Marshal(a.String())
}

func (a *ArgType) UnmarshalJSON(b []byte) error {
	var av string
	err := json.Unmarshal(b, &av)
	if err != nil {
		return err
	}
	switch av {
	case "const":
		*a = ConstArgType

	case "transform":
		*a = TransformArgType
	case "one_to_one_pipe":
		*a = OneToOnePipeArgType
	case "pipe":
		*a = PipeArgType
	default:
		return fmt.Errorf("Transform arg type '%s' not recognized", av)
	}
	return nil
}

type TransformArg struct {
	Description string  `json:"description"`       // A description of what the arg represents
	Optional    bool    `json:"optional"`          // Whether the arg is optional
	Default     *Pipe   `json:"default,omitempty"` // If the arg is optional, what is its default value
	Type        ArgType `json:"arg_type"`          // The type expected of the arg
}

type Transform struct {
	Name          string         `json:"name"`              // The name of the transform
	Description   string         `json:"description"`       // A single line description of the transform
	Documentation string         `json:"documentation"`     // Full markdown documentation of the transform
	InputSchema   string         `json:"ischema,omitempty"` // The schema of the input datapoint that the given transform expects (optional)
	OutputSchema  string         `json:"oschema,omitempty"` // The schema of the output data that this transform gives (optional).
	Args          []TransformArg `json:"args"`              // The arguments that the transform accepts

	Constructor TransformConstructor `json:"-"` // The function that constructs a transform
}

type TransformConstructor func(transform *Transform, consts []interface{}, pipes []*Pipe) (TransformIterator, error)

type TransformIterator interface {
	Next(*TransformEnv, *Datapoint) (*Datapoint, error)
	OneToOne() bool
}

var (
	// TransformRegistry is the map of all the transforms that are currently registered.
	// Do not manually add/remove elements from this map.
	// Use Transform.Register to insert new transforms.
	TransformRegistry = map[string]*Transform{
		"$":  Identity,
		"t":  T,
		"dt": DT,
	}

	// RegistryLock enables adding/deleting transforms during runtime. It is exported, since some
	// applications (heedy) might want to print out the registry
	RegistryLock = &sync.RWMutex{}
)

// Unregister removes the given named transform from the registry
func Unregister(name string) {
	RegistryLock.Lock()
	delete(TransformRegistry, name)
	RegistryLock.Unlock()
}

// Register registers the transform with the system.
func (t *Transform) Register() error {
	if t.Name == "" || t.Constructor == nil {
		return fmt.Errorf("Attempted to register invalid transform: '%s'", t.Name)
	}
	hadOptional := false
	for i := range t.Args {
		if t.Args[i].Optional {
			hadOptional = true
		} else if hadOptional {
			return fmt.Errorf("Transform '%s' has required arg after optional args", t.Name)
		}
	}

	RegistryLock.Lock()
	defer RegistryLock.Unlock()

	TransformRegistry[t.Name] = t

	return nil
}
