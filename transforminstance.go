package pipescript

// TransformInstance is the interface which underlies each transform. Copy() should return an exact copy
// of this transform instance as it is right now.
type TransformInstance interface {
	Next(dp DatapointPeekIterator, args []*Datapoint) (*Datapoint, error)
	Copy() TransformInstance
}
