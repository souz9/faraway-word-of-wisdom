package quotes

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestQuotesLoad(t *testing.T) {
	var qs Quotes
	err := qs.Load("testdata/quotes.txt")
	require.NoError(t, err)
	require.Equal(t, Quotes{
		"One",
		"Two",
		"Three",
	}, qs)
}
