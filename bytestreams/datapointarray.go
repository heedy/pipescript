package bytestreams

import (
	"encoding/json"
	"io"
	"io/ioutil"

	"github.com/heedy/pipescript"
)

func NewArrayReader(r io.Reader) (*pipescript.DatapointArrayIterator, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	var dpa []pipescript.Datapoint
	err = json.Unmarshal(b, &dpa)
	return pipescript.NewDatapointArrayIterator(dpa), err
}

//DatapointReader
type DatapointReader struct {
	dec *json.Decoder
}

// Next allows us to conform to the DatapointIterator interface
func (r *DatapointReader) Next(out *pipescript.Datapoint) (*pipescript.Datapoint, error) {
	if r.dec.More() {
		return out, r.dec.Decode(out)
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
