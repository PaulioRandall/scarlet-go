package tokeniser

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/token"
	"github.com/PaulioRandall/scarlet-go/tokeniser/source"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIdThunk_1(t *testing.T) {

	s := source.New("abc")
	f := idThunk(s, nil)
	require.NotNil(t, f)

	exp := token.New("abc", token.ID, 0, 0, 3)
	act, f, e := f()

	assert.Nil(t, e)
	assert.Nil(t, f)
	assert.Equal(t, exp, act)
}
