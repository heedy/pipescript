package core

import (
	"errors"
	"reflect"
	"sort"

	"github.com/connectordb/pipescript"
	"github.com/connectordb/pipescript/resources"
)

// https://groups.google.com/forum/#!topic/golang-nuts/FT7cjmcL7gw
type keyPairValue struct {
	key    string
	value  interface{}
	weight float64
}
type KeyPairList []keyPairValue

func (p KeyPairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p KeyPairList) Len() int           { return len(p) }
func (p KeyPairList) Less(i, j int) bool { return p[i].weight < p[j].weight }

type topTransform struct {
	number int
	script *pipescript.Script
}

func (t *topTransform) Copy() (pipescript.TransformInstance, error) {
	return &topTransform{t.number, t.script}, nil
}

func (t *topTransform) Next(ti *pipescript.TransformIterator) (*pipescript.Datapoint, error) {
	te := ti.Next()
	if te.IsFinished() {
		return te.Get()
	}

	// We only support maps for now in the data of our datapoint
	dmap := reflect.ValueOf(te.Datapoint.Data)
	if dmap.Kind() != reflect.Map {
		return nil, errors.New("Must have object-valued datapoint")
	}

	k := dmap.MapKeys()
	kpl := make(KeyPairList, dmap.Len())

	for i := range k {
		ks, ok := k[i].Interface().(string)
		if !ok {
			return nil, errors.New("Unable to transform non-string keys")
		}
		// Now set up the script - and end it with an if last so only one result is given
		s, err := t.script.Copy()
		if err != nil {
			return nil, err
		}

		pointdata := dmap.MapIndex(k[i]).Interface()

		dpi := pipescript.NewDatapointArrayIterator([]pipescript.Datapoint{
			pipescript.Datapoint{Timestamp: te.Datapoint.Timestamp, Data: pointdata},
		})
		s.SetInput(dpi)

		v, err := s.Next()
		if err != nil {
			return &pipescript.Datapoint{Timestamp: te.Datapoint.Timestamp, Data: nil}, err
		}
		f, err := v.Float()
		if err != nil {
			return nil, err
		}

		kpl[i] = keyPairValue{
			key:    ks,
			value:  pointdata,
			weight: f,
		}

	}

	// We have set up the kpl array. Now sort it
	sort.Sort(kpl)

	resultMap := make(map[string]interface{})

	for i := range kpl {
		if i >= t.number {
			break
		}
		curval := kpl[len(kpl)-i-1]
		resultMap[curval.key] = curval.value
	}

	return te.Set(resultMap)
}

var Top = pipescript.Transform{
	Name:          "top",
	Description:   `Takes a json object, and returns the topn elements`,
	Documentation: string(resources.MustAsset("docs/transforms/top.md")),
	OneToOne:      true,
	Stateless:     true,
	Args: []pipescript.TransformArg{
		{
			Description: "The number of elements to retain of the object",
			Constant:    true,
		},
		{
			Description: "The script to run ($ default) to generate weights over which to sort",
			Hijacked:    true,
			Optional:    true,
			Default:     nil, // THIS HAS TO BE SET IN init.go to identity

		},
	},
	Generator: func(name string, args []*pipescript.Script) (*pipescript.TransformInitializer, error) {
		dp, err := args[0].GetConstant()
		if err != nil {
			return nil, err
		}
		num, err := dp.Int()
		if err != nil {
			return nil, err
		}

		ifl, _ := iflast.Copy() // Shouldn't error
		args[1].Append(ifl)
		return &pipescript.TransformInitializer{Transform: &topTransform{int(num), args[1]}}, nil
	},
}
