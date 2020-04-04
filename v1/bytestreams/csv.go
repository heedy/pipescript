package bytestreams

import (
	"encoding/csv"
	"errors"
	"io"

	"github.com/heedy/pipescript"
)

//CSVDatapointReader is a reader for CSV files
type CSVDatapointReader struct {
	reader *csv.Reader
	gen    *DatapointGenerator
	dp     *pipescript.Datapoint
	keys   []string
}

func getCSVMap(keys []string, data []string) (map[string]interface{}, error) {
	if len(data) != len(keys) {
		return nil, errors.New("CSV file does not have equal row lengths")
	}
	datapoint := make(map[string]interface{})

	for i := range data {
		datapoint[keys[i]] = data[i]
	}
	return datapoint, nil
}

// Next allows us to conform to the DatapointIterator interface
func (r *CSVDatapointReader) Next() (*pipescript.Datapoint, error) {
	if r.dp != nil {
		dp := r.dp
		r.dp = nil
		return dp, nil
	}

	data, err := r.reader.Read()
	if err != nil {
		if err == io.EOF {
			// EOF means we're done
			return nil, nil
		}
		return nil, err
	}

	dp, err := getCSVMap(r.keys, data)
	if err != nil {
		return nil, err
	}
	return r.gen.Generate(dp)
}

// NewCSVDatapointReader creates a new reader for CSV files, using the optional
// timestamp key hint and optional disabling of timestamp
func NewCSVDatapointReader(r io.Reader, timestamphint string, disabletimestamp bool) (*CSVDatapointReader, error) {
	rdr := csv.NewReader(r)

	keys, err := rdr.Read()
	if err != nil {
		return nil, err
	}

	data, err := rdr.Read()
	if err != nil && err != io.EOF {
		return nil, err
	}

	// If timestamps are disabled, just read that shit without playing with timestamps
	if disabletimestamp || err == io.EOF {
		return &CSVDatapointReader{rdr, EmptyDatapointGenerator(), nil, keys}, nil
	}

	dp, err := getCSVMap(keys, data)
	if err != nil {
		return nil, err
	}
	gen, err := NewDatapointGenerator(dp, timestamphint)
	if err != nil {
		return nil, err
	}
	// Since generator succeeded, the datapoint will be read correctly
	dpt, _ := gen.Generate(dp)

	return &CSVDatapointReader{rdr, gen, dpt, keys}, nil
}
