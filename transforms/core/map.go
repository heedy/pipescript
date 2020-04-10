package core

import (
	"github.com/heedy/pipescript"
	"github.com/heedy/pipescript/resources"
)

type mapChannel struct {
	cp   *pipescript.ChannelPipe
	done bool
}

var Map = &pipescript.Transform{
	Name:          "map",
	Description:   "Splits the timeseries by the first arg, and returns an object where each key is the result of running the pipe in the transform in the second arg on the split series",
	Documentation: string(resources.MustAsset("docs/transforms/map.md")),
	Args: []pipescript.TransformArg{
		{
			Description: "The value to split on. This must be something that can be converted to string.",
			Type:        pipescript.TransformArgType,
			Schema: map[string]interface{}{
				"type": "boolean",
			},
		},
		{
			Description: "The transform to instantiate for each different value of the first argument.",
			Type:        pipescript.PipeArgType,
		},
	},
	Constructor: pipescript.NewAggregator(func(e *pipescript.TransformEnv, consts []interface{}, pipes []*pipescript.Pipe, out *pipescript.Datapoint) (*pipescript.Datapoint, error) {
		// Make the output map
		data := make(map[string]interface{})

		cp := make(map[string]*mapChannel)

		// The pipes need to be closed
		defer func() {
			for _, p := range cp {
				p.cp.Close()
			}
		}()

		argarray := make([]*pipescript.Datapoint, 1)
		// While there is new data, keep adding it to the appropriate pipe
		dp, args, err := e.Next(argarray)
		if err != nil || dp == nil {
			return nil, err
		}
		dp2 := dp
		out.Timestamp = dp.Timestamp
		for dp != nil {
			dp2 = dp
			key := args[0].ToString()
			p, ok := cp[key]
			if !ok {
				p = &mapChannel{pipescript.NewChannelPipe(pipes[0].Copy()), false}
				cp[key] = p
			}
			sentDP := false
			for !sentDP && !p.done {
				select {
				case res := <-p.cp.Receiver:
					if res.Err != nil {
						return nil, res.Err
					}
					if res.DP == nil {
						p.done = true
					} else {
						data[key] = res.DP.Data
					}
				case p.cp.Sender <- dp:
					sentDP = true
				}
			}

			dp, args, err = e.Next(argarray)
			if err != nil {
				return nil, err
			}
		}
		// All the data was sent. Now send the null datapoint, signaling end of stream
		for key, p := range cp {
			sentDP := false
			for !sentDP && !p.done {
				select {
				case res := <-p.cp.Receiver:
					if res.Err != nil {
						return nil, res.Err
					}
					if res.DP == nil {
						p.done = true
					} else {
						data[key] = res.DP.Data
					}
				case p.cp.Sender <- nil:
					sentDP = true
				}
			}
		}
		// Finally, empty all the pipes of data
		for key, p := range cp {
			for !p.done {
				res := <-p.cp.Receiver
				if res.Err != nil {
					return nil, res.Err
				}
				if res.DP == nil {
					p.done = true
				} else {
					data[key] = res.DP.Data
				}
			}
		}

		out.Duration = dp2.Timestamp + dp2.Duration - out.Timestamp
		out.Data = data
		return out, nil
	}),
}
