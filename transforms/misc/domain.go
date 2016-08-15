package misc

import (
	"net/url"

	"github.com/connectordb/pipescript"
)

type domainTransform struct{}

func (t *domainTransform) Copy() (pipescript.TransformInstance, error) {
	return &domainTransform{}, nil
}

func (t *domainTransform) Next(ti *pipescript.TransformIterator) (*pipescript.Datapoint, error) {
	te := ti.Next()
	if te.IsFinished() {
		return te.Get()
	}

	s, err := te.Datapoint.DataString()
	if err != nil {
		return nil, err
	}

	u, err := url.Parse(s)
	if err != nil {
		return te.Set("")
	}

	return te.Set(u.Host)
}

var Domain = pipescript.Transform{
	Name:         "domain",
	Description:  "Returns the domain name/host that is used in the given url",
	OutputSchema: `{"type": "string"}`,
	OneToOne:     true,
	Stateless:    true,
	Generator: func(name string, args []*pipescript.Script) (*pipescript.TransformInitializer, error) {
		return &pipescript.TransformInitializer{Transform: &domainTransform{}}, nil
	},
}
