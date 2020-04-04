package strings

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTokenize(t *testing.T) {

	require.Equal(t, Tokenize("Hello World!"), []string{"hello", "world"})
	require.Equal(t, Tokenize("I like,trains"), []string{"i", "like", "trains"})
	require.Equal(t, Tokenize("WTF I dunno :)"), []string{"wtf", "i", "dunno"})
	require.Equal(t, Tokenize("wouldn't have it's"), []string{"wouldn't", "have", "it's"})
}
