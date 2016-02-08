package bytestreams

import (
	"encoding/json"
	"testing"

	"github.com/connectordb/pipescript"
	"github.com/stretchr/testify/require"
)

func TestJsonReader(t *testing.T) {
	timestamps := []float64{1000, 1500, 2001, 2500, 3000}

	dpb := make([]pipescript.Datapoint, 5)
	var dpc pipescript.Datapoint

	for i := 0; i < 5; i++ {
		dpb[i] = pipescript.Datapoint{Timestamp: timestamps[i], Data: float64(i)}
	}

	dpa := pipescript.NewDatapointArrayIterator(dpb)

	jr, err := NewJsonReader(dpa, "", "\n", "", "", "")
	require.NoError(t, err)

	dec := json.NewDecoder(jr)
	for i := 0; i < 5; i++ {
		require.NoError(t, dec.Decode(&dpc))
		require.EqualValues(t, dpb[i], dpc)
	}

}
