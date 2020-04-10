package datasets

import (
	"errors"

	"github.com/heedy/pipescript"
	"github.com/mitchellh/mapstructure"
)

type timestampIterator struct {
	Until    float64
	NextDP   *pipescript.Datapoint
	Iterator pipescript.Iterator
	Done     bool
}

func (t *timestampIterator) Next(out *pipescript.Datapoint) (*pipescript.Datapoint, error) {
	if t.NextDP != nil {
		if t.NextDP.Timestamp >= t.Until {
			return nil, nil
		}
		out.Timestamp = t.NextDP.Timestamp
		out.Duration = t.NextDP.Duration
		out.Data = t.NextDP.Data
		t.NextDP = nil
		return out, nil
	}
	if t.Done {
		return nil, nil
	}
	dp, err := t.Iterator.Next(out)
	if err != nil || dp == nil {
		if dp == nil && err == nil {
			t.Done = true
		}
		return dp, err
	}
	if dp.Timestamp >= t.Until {
		t.NextDP = &pipescript.Datapoint{
			Timestamp: dp.Timestamp,
			Data:      dp.Data,
			Duration:  dp.Duration,
		}
		return nil, nil
	}
	return dp, nil
}

// IterateUntil clears the datapoints before the Until time
func (t *timestampIterator) IterateUntil() error {
	if !t.Done || t.NextDP != nil {
		dp := &pipescript.Datapoint{}
		for {
			dp2, err := t.Next(dp)
			if err != nil {
				return err
			}
			if dp2 == nil {
				return nil
			}
		}
	}
	return nil
}

type tTransformInterpolator struct {
	ref  *pipescript.BufferIterator
	pipe *pipescript.Pipe
	tsi  timestampIterator
	opt  TIOptions
}

func (ti *tTransformInterpolator) Next(out *pipescript.Datapoint) (*pipescript.Datapoint, error) {
	err := ti.tsi.IterateUntil() // Move the iterator to the start of the stream for this datapoint
	if err != nil {
		return nil, err
	}

	ref, err := ti.ref.Next()
	if ref == nil || err != nil {
		return ref, err
	}

	ti.tsi.Until = ref.Timestamp
	p := ti.pipe.Copy()
	p.InputIterator(&ti.tsi)
	dp, err := p.Last(out)
	if dp == nil {
		out.Timestamp = ref.Timestamp
		out.Duration = ref.Duration
		out.Data = nil
		return out, err
	}
	return dp, err
}

type dtTransformInterpolator struct {
	ref  *pipescript.BufferIterator
	pipe *pipescript.Pipe
	tsi  timestampIterator
	opt  TIOptions
}

func (ti *dtTransformInterpolator) Next(out *pipescript.Datapoint) (*pipescript.Datapoint, error) {
	ref, err := ti.ref.Next()
	if ref == nil || err != nil {
		return ref, err
	}
	ti.tsi.Until = ref.Timestamp
	err = ti.tsi.IterateUntil()
	if err != nil {
		return nil, err
	}

	ti.tsi.Until = ref.Timestamp + ref.Duration
	p := ti.pipe.Copy()
	p.InputIterator(&ti.tsi)
	dp, err := p.Last(out)
	if dp == nil {
		out.Timestamp = ref.Timestamp
		out.Duration = ref.Duration
		out.Data = nil
		return out, err
	}
	return dp, err
}

type TIOptions struct {
	Transform string `mapstructure:"transform"`
	RunOn     string `mapstructure:"run_on"`
}

var TransformInterpolator = &Interpolator{
	Name:        "transform",
	Description: "Allows using pipescript transforms to interpolate between values",
	Options: map[string]interface{}{
		"transform": map[string]interface{}{
			"type": "string",
		},
		"run_on": map[string]interface{}{
			"type":    "string",
			"enum":    []interface{}{"t", "dt"},
			"default": "t",
		},
		"required": []interface{}{"transform", "run_on"},
	},
	Constructor: func(name string, options map[string]interface{}, reference *pipescript.BufferIterator, stream pipescript.Iterator) (pipescript.Iterator, error) {
		var opt TIOptions
		err := mapstructure.Decode(options, &opt)
		if err != nil {
			return nil, err
		}
		if opt.Transform == "" {
			return nil, errors.New("Empty transform")
		}
		p, err := pipescript.Parse(opt.Transform)
		if err != nil {
			return nil, err
		}
		tsi := timestampIterator{
			Until:    -9999999999,
			Iterator: stream,
		}
		switch opt.RunOn {
		case "t":
			return &tTransformInterpolator{
				ref:  reference,
				pipe: p,
				tsi:  tsi,
				opt:  opt,
			}, nil
		case "dt":
			return &dtTransformInterpolator{
				ref:  reference,
				pipe: p,
				tsi:  tsi,
				opt:  opt,
			}, nil
		default:
			return nil, errors.New("Unrecognized transform run_on value")
		}
	},
}

func init() {
	// Registering initializes the json schema of options
	TransformInterpolator.Register()
}
