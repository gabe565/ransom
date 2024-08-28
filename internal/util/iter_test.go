package util

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAlphabet(t *testing.T) {
	it := Alphabet()
	require.NotNil(t, it)
	letters := slices.Collect(it)
	assert.Len(t, letters, 26)
	for i, letter := range letters {
		assert.Equal(t, string('a'+byte(i)), letter)
	}
}
