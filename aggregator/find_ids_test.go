package aggregator

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/stat"
	"github.com/PaulioRandall/scarlet-go/token"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFindIds_1(t *testing.T) {
	// Check it is a type of SequenceFinder.
	var _ SequenceFinder = findIds
}

func TestFindIds_2(t *testing.T) {
	// Check it works when a single ID token is the only input.

	in := []token.Token{
		token.NewToken(token.ID, "abc", 0, 0, 3),
	}

	n, k, e := findIds(in)

	require.Nil(t, e)
	assert.Equal(t, stat.OK, k,
		"Expected kind `%s`, got `%s`", stat.OK.Name(), k.Name())
	assert.Equal(t, 1, n)
}

func TestFindIds_3(t *testing.T) {
	// Check it works when a multiple delimiter separated ID tokens are input.

	in := []token.Token{
		token.NewToken(token.ID, "abc", 0, 0, 3),
		token.NewToken(token.ID_DELIM, ",", 0, 3, 4),
		token.NewToken(token.ID, "efg", 0, 4, 7),
		token.NewToken(token.ID_DELIM, ",", 0, 7, 8),
		token.NewToken(token.ID, "xyz", 0, 8, 11),
	}

	n, k, e := findIds(in)

	require.Nil(t, e)
	assert.Equal(t, stat.OK, k,
		"Expected kind `%s`, got `%s`", stat.OK.Name(), k.Name())
	assert.Equal(t, 5, n)
}

func TestFindIds_4(t *testing.T) {
	// Check it works when a  delimiter separated ID tokens are followed by other
	// tokens.

	in := []token.Token{
		token.NewToken(token.ID, "abc", 0, 0, 3),
		token.NewToken(token.ID_DELIM, ",", 0, 3, 4),
		token.NewToken(token.ID, "efg", 0, 4, 7),
		token.NewToken(token.ASSIGN, ":=", 0, 7, 9),
		token.NewToken(token.STR_LITERAL, "`xyz`", 0, 9, 14),
	}

	n, k, e := findIds(in)

	require.Nil(t, e)
	assert.Equal(t, stat.OK, k,
		"Expected kind `%s`, got `%s`", stat.OK.Name(), k.Name())
	assert.Equal(t, 3, n)
}

func TestFindIds_5(t *testing.T) {
	// Check zero values are returned when the input does not lead with ID tokens.

	in := []token.Token{
		token.NewToken(token.ASSIGN, ":=", 0, 0, 2),
		token.NewToken(token.ID, "abc", 0, 2, 5),
	}

	n, k, e := findIds(in)

	require.Nil(t, e)
	assert.Equal(t, stat.UNDEFINED, k,
		"Expected kind `%s`, got `%s`", stat.UNDEFINED.Name(), k.Name())
	assert.Equal(t, 0, n)
}
