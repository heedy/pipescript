package strings

import (
	"regexp"

	"github.com/connectordb/pipescript"
	"github.com/connectordb/pipescript/resources"
)

type regexTransform struct {
	re *regexp.Regexp
}

func (t *regexTransform) Copy() (pipescript.TransformInstance, error) {
	return &regexTransform{t.re}, nil
}

func (t *regexTransform) Next(ti *pipescript.TransformIterator) (*pipescript.Datapoint, error) {
	te := ti.Next()
	if te.IsFinished() {
		return te.Get()
	}

	s, err := te.Datapoint.DataString()
	if err != nil {
		return nil, err
	}

	return te.Set(t.re.MatchString(s))
}

var Regex = pipescript.Transform{
	Name:          "regex",
	Description:   "Returns true if the given regular expression matches the data string",
	OutputSchema:  `{"type": "boolean"}`,
	Documentation: string(resources.MustAsset("docs/transforms/regex.md")),
	OneToOne:      true,
	Stateless:     true,
	Args: []pipescript.TransformArg{
		{
			Description: "The regular expression to use",
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
		re, err := regexp.Compile(s)
		if err != nil {
			return nil, err
		}
		return &pipescript.TransformInitializer{Transform: &regexTransform{re}}, nil
	},
}
