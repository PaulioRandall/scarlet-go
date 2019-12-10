package aggregator

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/stat"
	"github.com/PaulioRandall/scarlet-go/token"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStream_SliceBy_1(t *testing.T) {
	// Typicial usage.

	input := []token.Token{
		token.NewToken(token.ID, "abc", 0, 0, 3),
		token.NewToken(token.ASSIGN, ":=", 0, 4, 7),
		token.NewToken(token.ID, "xyz", 0, 8, 11),
		token.NewToken(token.NEWLINE, "\n", 0, 11, 12),
		token.NewToken(token.FUNC, "F", 1, 0, 1),
	}

	s := stream{
		t: input,
	}

	act, k, e := s.SliceBy(func(tok []token.Token) (int, stat.Kind, token.Perror) {
		assert.Equal(t, input, tok)
		return 3, stat.ASSIGN_ID, nil
	})

	require.Nil(t, e)
	assert.Equal(t, stat.ASSIGN_ID, k)
	assert.Equal(t, input[:3], act)
}

func TestStream_SliceBy_2(t *testing.T) {
	// Out of range slice indexes panic.

	s := stream{
		t: []token.Token{
			token.NewToken(token.ID, "abc", 0, 0, 3),
		},
	}

	assert.Panics(t, func() {
		s.SliceBy(func(tok []token.Token) (int, stat.Kind, token.Perror) {
			return 99, stat.ASSIGN_ID, nil
		})
	})

	assert.Panics(t, func() {
		s.SliceBy(func(tok []token.Token) (int, stat.Kind, token.Perror) {
			return 0, stat.ASSIGN_ID, nil
		})
	})
}