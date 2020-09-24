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
