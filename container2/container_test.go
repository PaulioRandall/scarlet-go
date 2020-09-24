package container2

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/token"

	"github.com/stretchr/testify/require"
)

func randomLexemes() (a, b, c, d token.Lexeme) {
	a = token.New("true", token.BOOL, 0, 0)
	b = token.New("1", token.NUMBER, 0, 4)
	c = token.New("abc", token.STRING, 0, 5)
	d = token.New("i", token.IDENT, 0, 8)
	return
}

func setup() (con *Container, a, b, c, d token.Lexeme) {

	a, b, c, d = randomLexemes()
	con = &Container{}

	con.append(a)
	con.append(b)
	con.append(c)
	con.append(d)

	return
}

func requireNodes(t *testing.T, con *Container, data ...token.Lexeme) {

	var prev *node
	var next *node = con.head

	for _, l := range data {
		require.NotNil(t, next)
		require.Equal(t, l, next.data)
		require.Equal(t, prev, next.prev)
		prev, next = next, next.next
	}

	require.Equal(t, prev, con.tail)
	require.Nil(t, next)
	require.Equal(t, len(data), con.size)
}

func Test_Container_prepend(t *testing.T) {

	a, b, c, d := randomLexemes()
	con := &Container{}

	con.prepend(d)
	requireNodes(t, con, d)

	con.prepend(c)
	requireNodes(t, con, c, d)

	con.prepend(b)
	requireNodes(t, con, b, c, d)

	con.prepend(a)
	requireNodes(t, con, a, b, c, d)
}

func Test_Container_append(t *testing.T) {

	a, b, c, d := randomLexemes()
	con := &Container{}

	con.append(a)
	requireNodes(t, con, a)

	con.append(b)
	requireNodes(t, con, a, b)

	con.append(c)
	requireNodes(t, con, a, b, c)

	con.append(d)
	requireNodes(t, con, a, b, c, d)
}

func Test_Container_pop(t *testing.T) {

	con, a, b, c, d := setup()
	var l token.Lexeme

	l = con.pop()
	require.Equal(t, a, l)
	requireNodes(t, con, b, c, d)

	l = con.pop()
	require.Equal(t, b, l)
	requireNodes(t, con, c, d)

	l = con.pop()
	require.Equal(t, c, l)
	requireNodes(t, con, d)

	l = con.pop()
	require.Equal(t, d, l)
	requireNodes(t, con)
}
