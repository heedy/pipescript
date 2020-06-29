package interpolators

import (
	"github.com/heedy/pipescript"
	"github.com/heedy/pipescript/datasets"
)

type closestInterpolator struct {
	reference *pipescript.BufferIterator
	stream    pipescript.Iterator
	prev      *pipescript.Datapoint
	next      *pipescript.Datapoint
}

func (c *closestInterpolator) Next(dp *pipescript.Datapoint) (*pipescript.Datapoint, error) {

	ref, err := c.reference.Next()
	if err != nil || ref == nil {
		return nil, err
	}
	if c.next == nil || c.prev.Timestamp >= ref.Timestamp {
		// We're on the last datapoint of the stream
		*dp = *c.prev
		return dp, nil
	}

	// Now next is not nil, and prev's timestamp is less than current reference
	for c.next.Timestamp <= ref.Timestamp {
		pd := c.prev
		c.prev = c.next
		c.next, err = c.stream.Next(pd)
		if c.next == nil || err != nil {
			*dp = *c.prev
			return dp, err
		}
	}

	// At this point, prev has timestamp <= ref, and next has timestamp>=ref. We check which is closer to ref
	if ref.Timestamp-c.prev.Timestamp > c.next.Timestamp-ref.Timestamp {
		// The next datapoint is closer. Since future datapoints can't be closer to prev, shift by one more datapoint
		pd := c.prev
		c.prev = c.next
		c.next, err = c.stream.Next(pd)

	}
	*dp = *c.prev

	return dp, err

}

var Closest = &datasets.Interpolator{
	Name:        "closest",
	Description: "Returns the datapoint with the closest timestamp to the reference timestamp",
	Constructor: func(name string, options map[string]interface{}, reference *pipescript.BufferIterator, stream pipescript.Iterator) (pipescript.Iterator, error) {
		dp1, err := stream.Next(&pipescript.Datapoint{})
		if err != nil || dp1 == nil {
			return pipescript.EmptyIterator{}, nil
		}
		dp2, err := stream.Next(&pipescript.Datapoint{})
		return &closestInterpolator{
			reference: reference,
			stream:    stream,
			prev:      dp1,
			next:      dp2,
		}, err
	},
}
