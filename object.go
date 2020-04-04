package pipescript

import (
	"container/list"
	"fmt"
)

type aggregatePipeContext struct {
	cp   *ChannelPipe
	done bool
}

type aggregateObjectTransform struct {
	obj    map[string]*aggregatePipeContext
	data   map[string]interface{}
	isDone bool
}

func (a *aggregateObjectTransform) OneToOne() bool {
	return false
}

func (a *aggregateObjectTransform) Close() {
	for _, c := range a.obj {
		c.cp.Close()
	}
}

func (a *aggregateObjectTransform) Next(e *TransformEnv, out *Datapoint) (*Datapoint, error) {
	if a.isDone {

		return nil, nil
	}
	defer a.Close()
	dp, _, err := e.Next(nil)
	if err != nil || dp == nil {
		return nil, err
	}
	out.Timestamp = dp.Timestamp
	for {

		for key, pc := range a.obj {
			sentDP := false
			for !sentDP && !pc.done {
				select {
				case res := <-pc.cp.Receiver:
					if res.Err != nil {
						return nil, res.Err
					}
					if res.DP == nil {
						pc.done = true
					} else {
						a.data[key] = res.DP.Data
					}
				case pc.cp.Sender <- dp: // no need to copy dp before sending, since out iterator is at loc
					sentDP = true
				}
			}
		}
		// Once we sent the null datapoint, the streams are considered finished
		if dp == nil {
			break
		}
		out.Duration = dp.Timestamp + dp.Duration - out.Timestamp

		dp, _, err = e.Next(nil)
		if err != nil {
			return nil, err
		}
	}

	// All the data was sent. Now iterate until all the elements are done processing
	for key, pc := range a.obj {
		for !pc.done {
			res := <-pc.cp.Receiver
			if res.Err != nil {
				return nil, res.Err
			}
			if res.DP == nil {
				pc.done = true
			} else {
				a.data[key] = res.DP.Data
			}
		}
	}
	out.Data = a.data
	a.isDone = true
	return out, nil
}

type oneToOneObjectContext struct {
	cp    *ChannelPipe
	recvd *list.List
	done  bool
}

type oneToOneObjectTransform struct {
	obj      map[string]*oneToOneObjectContext
	isDone   bool
	dataDone bool
	idx      int
}

func (o *oneToOneObjectTransform) OneToOne() bool {
	return true
}

func (o *oneToOneObjectTransform) Close() {
	for _, c := range o.obj {
		c.cp.Close()
	}
}

func (o *oneToOneObjectTransform) Next(e *TransformEnv, out *Datapoint) (*Datapoint, error) {
	if o.isDone {
		return nil, nil
	}

	// Check if we have a result already waiting in the list
	hasNext := false
	for _, pc := range o.obj {
		if pc.recvd.Len() == 0 {
			hasNext = false
			break
		}
	}
	if !hasNext {
		if !o.dataDone {
			dp, _, err := e.Peek(o.idx, nil)
			if err != nil {
				o.Close()
				return nil, err
			}
			o.idx++

			for {
				hasNext = true
				for _, pc := range o.obj {
					sentDP := false
					for !sentDP && !pc.done {
						select {
						case res := <-pc.cp.Receiver:
							if res.Err != nil {
								o.Close()
								return nil, res.Err
							}
							if res.DP == nil {
								pc.done = true
							} else {
								pc.recvd.PushBack(res.DP)
							}
						case pc.cp.Sender <- dp: // no need to copy dp before sending, since out iterator is at loc
							sentDP = true
						}
					}
					if pc.recvd.Len() == 0 {
						if pc.done {
							o.Close()
							o.isDone = true
							return nil, nil // No more datapoints
						}
						hasNext = false
					}
				}
				if dp == nil {
					o.dataDone = true
					break
				}
				if hasNext {
					// We can return the datapoint
					break
				}

				dp, _, err = e.Peek(o.idx, nil)
				if err != nil {
					o.Close()
					return nil, err
				}
				o.idx++
			}
		}
		if o.dataDone && !hasNext {
			// The data is done, but we don't have results for all the keys!
			for _, pc := range o.obj {
				if pc.recvd.Len() == 0 {
					if pc.done {
						o.isDone = true
						o.Close()
						return nil, nil
					}
					// Don't have data for this one. Get it.
					res := <-pc.cp.Receiver
					if res.Err != nil {
						return nil, res.Err
					}
					if res.DP == nil {
						pc.done = true
						o.isDone = true
						o.Close()
						return nil, nil
					}
					pc.recvd.PushBack(res.DP)

				}
			}

		}
	}

	data := make(map[string]interface{})
	for key, pc := range o.obj {
		data[key] = pc.recvd.Remove(pc.recvd.Front()).(*Datapoint).Data
	}
	out.Data = data
	dp, _, err := e.Next(nil)
	if err != nil {
		o.Close()
		return nil, err
	}
	o.idx--
	out.Timestamp = dp.Timestamp
	out.Duration = dp.Duration

	return out, nil
}

func NewObjectTransform(obj map[string]*Pipe) *Transform {
	if len(obj) == 0 {
		return NewConstTransform(make(map[string]interface{}))
	}
	isOneToOne := true
	for _, p := range obj {
		if !p.OneToOne() {
			isOneToOne = false
			break
		}
	}
	oname := "{"
	for k, v := range obj {
		oname += fmt.Sprintf("\"%s\": %s,", k, v.String())
	}
	oname = oname[:len(oname)-1] + "}"

	if isOneToOne {
		return &Transform{
			Name: oname,
			Constructor: func(transform *Transform, consts []interface{}, pipes []*Pipe) (TransformIterator, error) {
				ooc := make(map[string]*oneToOneObjectContext)
				for k, p := range obj {
					ooc[k] = &oneToOneObjectContext{
						cp:    NewChannelPipe(p.Copy()),
						recvd: list.New(),
					}
				}

				return &oneToOneObjectTransform{
					obj: ooc,
				}, nil
			},
		}
	}
	return &Transform{
		Name: oname,
		Constructor: func(transform *Transform, consts []interface{}, pipes []*Pipe) (TransformIterator, error) {

			apc := make(map[string]*aggregatePipeContext)
			vals := make(map[string]interface{})
			for k, p := range obj {
				apc[k] = &aggregatePipeContext{
					cp: NewChannelPipe(p.Copy()),
				}
				vals[k] = nil
			}

			return &aggregateObjectTransform{
				obj:  apc,
				data: vals,
			}, nil
		},
	}

}
