package container

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/token2/lexeme"
	"github.com/PaulioRandall/scarlet-go/token2/token"

	"github.com/stretchr/testify/require"
)

func dummyNodes() (a, b, c, d *node) {

	a = &node{
		data: lexeme.New("true", token.TRUE, 0, 0),
	}

	b = &node{
		data: lexeme.New("1", token.NUMBER, 0, 4),
	}

	c = &node{
		data: lexeme.New("abc", token.STRING, 0, 5),
	}

	d = &node{
		data: lexeme.New("i", token.IDENT, 0, 8),
	}

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
