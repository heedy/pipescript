package bytestreams

import (
	"encoding/json"
	"io"

	"github.com/connectordb/pipescript"
)

//DatapointReader
type DatapointReader struct {
	dec *json.Decoder
}

// Next allows us to conform to the DatapointIterator interface
func (r *DatapointReader) Next() (*pipescript.Datapoint, error) {
	if r.dec.More() {
		// There is another datapoint
		dp := &pipescript.Datapoint{}
		err := r.dec.Decode(dp)
		return dp, err
	}
	return nil, nil
}

func NewDatapointReader(r io.Reader) (*DatapointReader, error) {
	dec := json.NewDecoder(r)
	_, err := dec.Token() // Read starting value
	if err != nil {
		return nil, err
	}
	return &DatapointReader{dec}, nil
}
