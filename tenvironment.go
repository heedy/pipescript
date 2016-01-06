package pipescript

// TransformEnvironment contains a single instance of data. This information is NOT to be modified
// by individual transforms!
type TransformEnvironment struct {
	Datapoint *Datapoint
	Args      []*Datapoint
	Error     error
}

// IsFinished is true if the iterator is done or of the error is non-nil
func (t *TransformEnvironment) IsFinished() bool {
	return t.Error != nil || t.Datapoint == nil
}

// Get returns the current datapoint and error if any exists
func (t *TransformEnvironment) Get() (*Datapoint, error) {
	return t.Datapoint, t.Error
}

// Set returns a datapoint formatted with the given value as data
func (t *TransformEnvironment) Set(val interface{}) (*Datapoint, error) {
	if t.Datapoint != nil {
		return &Datapoint{t.Datapoint.Timestamp, val}, t.Error
	}
	return nil, t.Error
}
