package bytestreams

import (
	"encoding/json"
	"io"

	"github.com/heedy/pipescript"
)

//JSONDatapointReader
type JSONDatapointReader struct {
	dec *json.Decoder
	gen *DatapointGenerator
	dp  *pipescript.Datapoint
}

// Next allows us to conform to the DatapointIterator interface
func (r *JSONDatapointReader) Next(out *pipescript.Datapoint) (*pipescript.Datapoint, error) {
	if r.dp != nil {
		out.Timestamp = r.dp.Timestamp
		out.Duration = r.dp.Duration
		out.Data = r.dp.Data
		r.dp = nil
		return out, nil
	}
	if r.dec.More() {
		// There is another datapoint
		dp := make(map[string]interface{})
		err := r.dec.Decode(&dp)
		if err != nil {
			return nil, err
		}
		return r.gen.Generate(dp)
	}
	return nil, nil
}

// NewJSONDatapointReader creates a new readed for JSON files, using the optional
// timestamp key hint and optional disabling of timestamp
func NewJSONDatapointReader(r io.Reader, timestamphint string, disabletimestamp bool) (*JSONDatapointReader, error) {
	dec := json.NewDecoder(r)
	_, err := dec.Token() // Read starting value
	if err != nil {
		return nil, err
	}

	// If timestamps are disabled, just read that shit without playing with timestamps
	if disabletimestamp || !dec.More() {
		return &JSONDatapointReader{dec, EmptyDatapointGenerator(), nil}, nil
	}

	dp := make(map[string]interface{})
	err = dec.Decode(&dp)
	if err != nil {
		return nil, err
	}
	gen, err := NewDatapointGenerator(dp, timestamphint)
	if err != nil {
		return nil, err
	}
	// Since generator succeeded, the datapoint will be read correctly
	dpt, _ := gen.Generate(dp)

	return &JSONDatapointReader{dec, gen, dpt}, nil
}
