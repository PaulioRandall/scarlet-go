package recursive

import (
	"testing"

	. "github.com/PaulioRandall/scarlet-go/pkg/token"

	"github.com/stretchr/testify/require"
)

func Test_P1(t *testing.T) {

	// p._peek(), p._next(), p._prev()

	tok := func(m Morpheme) Token {
		return NewToken(m, "", 0, 0)
	}

	p := newPipe([]Token{
		tok(NUMBER),
		tok(ADD),
		tok(NUMBER),
	})

	require.Equal(t, Token(nil), p._prev())

	require.Equal(t, tok(NUMBER), p._peek())
	require.Equal(t, tok(NUMBER), p._next())
	require.Equal(t, tok(NUMBER), p._prev())

	require.Equal(t, tok(ADD), p._peek())
	require.Equal(t, tok(ADD), p._next())
	require.Equal(t, tok(ADD), p._prev())

	require.Equal(t, tok(NUMBER), p._peek())
	require.Equal(t, tok(NUMBER), p._next())
	require.Equal(t, tok(NUMBER), p._prev())

	require.Equal(t, Token(nil), p._peek())
	require.Equal(t, Token(nil), p._next())
}

func Test_P2(t *testing.T) {

	// p.hasMore(), p.match(), p.accept()

	tok := func(m Morpheme) Token {
		return NewToken(m, "", 0, 0)
	}

	p := newPipe([]Token{
		tok(NUMBER),
		tok(ADD),
		tok(NUMBER),
	})

	require.Equal(t, true, p.hasMore())
	require.Equal(t, true, p.match(NUMBER))
	require.Equal(t, false, p.accept(ADD))
	require.Equal(t, true, p.accept(NUMBER))

	require.Equal(t, true, p.hasMore())
	require.Equal(t, true, p.match(ADD))
	require.Equal(t, false, p.accept(NUMBER))
	require.Equal(t, true, p.accept(ADD))

	require.Equal(t, true, p.hasMore())
	require.Equal(t, true, p.match(NUMBER))
	require.Equal(t, false, p.accept(ADD))
	require.Equal(t, true, p.accept(NUMBER))

	require.Equal(t, false, p.hasMore())
	require.Equal(t, false, p.match(NUMBER))
	require.Equal(t, false, p.accept(NUMBER))
}

func Test_P3(t *testing.T) {

	// p.expect()

	tok := func(m Morpheme) Token {
		return NewToken(m, "", 0, 0)
	}

	checkOk := func(t *testing.T, p *pipe, exp Token, m Morpheme) {
		act, e := p.expect(m)
		require.Nil(t, nil, e)
		require.Equal(t, exp, act)
	}

	checkErr := func(t *testing.T, p *pipe, m Morpheme) {
		act, e := p.expect(m)
		require.NotNil(t, e)
		require.Nil(t, nil, act)
	}

	p := newPipe([]Token{
		tok(NUMBER),
		tok(ADD),
		tok(NUMBER),
	})

	checkErr(t, p, ADD)
	checkOk(t, p, tok(NUMBER), NUMBER)

	checkErr(t, p, NUMBER)
	checkOk(t, p, tok(ADD), ADD)

	checkErr(t, p, ADD)
	checkOk(t, p, tok(NUMBER), NUMBER)

	checkErr(t, p, NUMBER)
}
