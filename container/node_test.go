package container

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/token"

	"github.com/stretchr/testify/require"
)

func dummyNodes() (a, b, c, d *node) {

	a = &node{
		data: token.New("true", token.BOOL, 0, 0),
	}

	b = &node{
		data: token.New("1", token.NUMBER, 0, 4),
	}

	c = &node{
		data: token.New("abc", token.STRING, 0, 5),
	}

	d = &node{
		data: token.New("i", token.IDENT, 0, 8),
	}

	return
}

func TestLinkAll(t *testing.T) {
	a, b, c, d := dummyNodes()
	head, tail := chain(a, b, c, d)
	require.Equal(t, a, head)
	require.Equal(t, d, tail)
}
