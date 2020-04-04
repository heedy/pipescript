package bytestreams

import (
	"errors"
	"time"

	"github.com/heedy/pipescript"
	"github.com/jinzhu/now"
)

func getTimestamp(data interface{}) (float64, error) {
	s, ok := data.(string)
	if !ok {
		return 0, errors.New("Timestamp was not string")
	}
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		t, err = now.Parse(s)
		if err != nil {
			return 0, err
		}

	}
	return float64(t.UnixNano()) * 1e-9, nil
}

func findTimestampKey(sampledata map[string]interface{}) (string, error) {
	for key := range sampledata {
		_, err := getTimestamp(sampledata[key])
		if err == nil {
			return key, nil
		}
	}
	return "", errors.New("Timestamp not found in dataset")
}

// DatapointGenerator converts a map of data to a datapoint, making sure that
// the associated timestamp is correctly parsed and converted. It modifies the map during execution
type DatapointGenerator struct {
	key string
}

// Generate performs the generation of datapoint
func (t *DatapointGenerator) Generate(data map[string]interface{}) (*pipescript.Datapoint, error) {
	if t.key == "" {
		// if we don't have a timestamp key, just return datapoint with 0 timestamp
		return &pipescript.Datapoint{Data: data}, nil
	}
	v, ok := data[t.key]
	if !ok {
		return nil, errors.New("Did not find timestamp field")
	}
	ts, err := getTimestamp(v)
	if err != nil {
		return nil, err
	}

	// Return the entire thing in the data field
	return &pipescript.Datapoint{
		Data:      data,
		Timestamp: ts,
	}, nil

}

// NewDatapointGenerator creates a new generator
func NewDatapointGenerator(sampledata map[string]interface{}, keyhint string) (*DatapointGenerator, error) {
	if keyhint != "" {
		v, ok := sampledata[keyhint]
		if !ok {
			return nil, errors.New("Given timestamp key not found")
		}
		_, err := getTimestamp(v)
		return &DatapointGenerator{keyhint}, err
	}
	key, err := findTimestampKey(sampledata)
	return &DatapointGenerator{key}, err
}

// EmptyDatapointGenerator returns datapoints without timestamps. For use on datasets which do not
// have timestamps defined
func EmptyDatapointGenerator() *DatapointGenerator {
	return &DatapointGenerator{}
}
