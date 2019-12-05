package tokeniser

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/token"
	"github.com/PaulioRandall/scarlet-go/tokeniser/source"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewlineThunk_1(t *testing.T) {

	s := source.New("\r\n")
	f := newlineThunk(s, nil)
	require.NotNil(t, f)

	exp := token.New("\r\n", token.NEWLINE, 0, 0, 2)
	act, f, e := f()

	assert.Nil(t, e)
	assert.Nil(t, f)
	assert.Equal(t, exp, act)
}

func TestNewlineThunk_2(t *testing.T) {

	s := source.New("abc")
	f := newlineThunk(s, nil)
	assert.NotNil(t, f)

	act, f, e := f()

	assert.NotNil(t, e)
	assert.Nil(t, f)
	assert.Empty(t, act)
}
