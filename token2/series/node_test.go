package series

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/token2/lexeme"
	"github.com/PaulioRandall/scarlet-go/token2/token"

	"github.com/stretchr/testify/require"
)

func dummyLexemes() (la, lb, lc, ld lexeme.Lexeme) {
	la = lexeme.Tok("true", token.TRUE)
	lb = lexeme.Tok("1", token.NUMBER)
	lc = lexeme.Tok("abc", token.STRING)
	ld = lexeme.Tok("i", token.IDENT)
	return
}

func dummyNodes() (a, b, c, d *node) {
	la, lb, lc, ld := dummyLexemes()
	a = &node{data: la}
	b = &node{data: lb}
	c = &node{data: lc}
	d = &node{data: ld}
	return
}

func TestChain(t *testing.T) {
	a, b, c, d := dummyNodes()
	head, tail, size := chain(a, b, c, d)
	require.Equal(t, a, head)
	require.Equal(t, d, tail)
	require.Equal(t, 4, size)
}

func TestUnlinkAll(t *testing.T) {

	a, b, c, d := dummyNodes()
	a.next = b
	b.next = c
	c.next = d
	b.prev = a
	c.prev = b
	d.prev = c

	unlinkAll(a, b, c)
	require.Nil(t, a.prev)
	require.Nil(t, a.next)
	require.Nil(t, b.prev)
	require.Nil(t, b.next)
	require.Nil(t, c.prev)
	require.Equal(t, d, c.next)
}
