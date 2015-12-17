package pipescript

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDatapointConversions(t *testing.T) {
	// The testing for duck conversions is done mainly in duck. We do token tests of each function
	//	to make sure things work as expected.
	d := Datapoint{2, 1337.0}

	intv, err := d.Int()
	assert.NoError(t, err)
	assert.Equal(t, intv, int64(1337))

	floatv, err := d.Float()
	assert.NoError(t, err)
	assert.Equal(t, floatv, 1337.0)

	strv, err := d.String()
	assert.NoError(t, err)
	assert.Equal(t, strv, "1337")

	//Now let's get those errors to happen
	d = Datapoint{2, []int{2, 3, 4}}

	_, err = d.Int()
	assert.Error(t, err)
	_, err = d.Float()
	assert.Error(t, err)
	_, err = d.String()
	assert.Error(t, err)

}

func TestDatapointGetSet(t *testing.T) {
	d := Datapoint{2, map[string]interface{}{"hello": "world"}}

	// Check getting existing
	v, err := d.Get("hello")
	assert.NoError(t, err)
	assert.Equal(t, "world", v)

	// Nonexisting
	_, err = d.Get("foo")
	assert.Error(t, err)

	// Set existing
	assert.NoError(t, d.Set("hello", "hi"))
	v, err = d.Get("hello")
	assert.NoError(t, err)
	assert.Equal(t, "hi", v)

	// Set new
	assert.NoError(t, d.Set("foo", "bar"))
	v, err = d.Get("foo")
	assert.NoError(t, err)
	assert.Equal(t, "bar", v)

	//Non-struct values can't be set
	d = Datapoint{2, 3}
	assert.Error(t, d.Set("foo", "bar"))
}

func TestCopy(t *testing.T) {
	d := Datapoint{2, map[string]interface{}{"hello": "world"}}
	d2 := d.Copy()

	v, err := d2.Get("hello")
	assert.NoError(t, err)
	assert.Equal(t, "world", v)

	//Make sure that the values are decoupled
	assert.NoError(t, d2.Set("hello", "hi"))
	v, err = d.Get("hello")
	assert.NoError(t, err)
	assert.Equal(t, "world", v)
}
