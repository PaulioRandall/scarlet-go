package tokeniser

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/token"
	"github.com/PaulioRandall/scarlet-go/tokeniser/source"

	"github.com/stretchr/testify/assert"
)

func TestNewlineEmitter_1(t *testing.T) {

	s := source.New("\r\n")
	f := newlineEmitter(s, nil)
	assert.NotNil(t, f)

	exp := token.New("", token.NEWLINE, 0, 0, 2)
	act, f, e := f()

	assert.Nil(t, e)
	assert.Nil(t, f)
	assert.Equal(t, exp.Where(), act.Where())
}

func TestNewlineEmitter_2(t *testing.T) {

	s := source.New("abc")
	f := newlineEmitter(s, nil)
	assert.NotNil(t, f)

	act, f, e := f()

	assert.NotNil(t, e)
	assert.Nil(t, f)
	assert.Empty(t, act)
}
