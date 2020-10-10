package pipescript

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func BenchmarkFloat(b *testing.B) {
	v := float64(3454.4)
	for n := 0; n < b.N; n++ {
		Float(v)
	}
}

func TestEqual(t *testing.T) {
	require.True(t, Equal(map[string]interface{}{"a": 1}, map[string]interface{}{"a": 1.0}))
	require.False(t, Equal(map[string]interface{}{"a": 1}, map[string]interface{}{"a": true}))
	require.False(t, Equal(map[string]interface{}{"a": 1}, nil))

	require.True(t, Equal([]interface{}{1, "hi"}, []interface{}{1, "hi"}))
	require.False(t, Equal([]interface{}{1, "hi"}, []interface{}{1, 3}))
}
