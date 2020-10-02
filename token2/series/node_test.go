package series

import (
	"fmt"
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

func chainLexemes(lexs ...lexeme.Lexeme) (head, tail *node, size int) {

	size = len(lexs)
	if size == 0 {
		return
	}

	nodes := make([]*node, size)
	for i, l := range lexs {
		nodes[i] = &node{data: l}
	}

	return chain(nodes...)
}

func requireChain(t *testing.T, exp *node, act *node) {
	// TODO: Too complex?

	errMsg := func(i int, expNode, actNode *node) string {

		var exp, act string
		if expNode == nil {
			exp = "nil"
		} else {
			exp = expNode.String()
		}

		if actNode == nil {
			act = "nil"
		} else {
			act = actNode.String()
		}

		return fmt.Sprintf("Unexpected node at %d; have %s, want %s", i, act, exp)
	}

	var expNode, actNode *node
	var expTail, actTail *node
	i := 0

	// Test node.next by working forward through the chain
	for expNode != nil || actNode != nil {

		require.NotNil(t, expNode, errMsg(i, expNode, actNode))
		require.NotNil(t, actNode, errMsg(i, expNode, actNode))
		require.Equal(t, expNode.data, actNode.data, errMsg(i, expNode, actNode))

		expTail, expNode = expNode, expNode.next
		actTail, actNode = actNode, actNode.next
		i++
	}

	// Test node.prev by working back through the chain
	expNode = expTail
	actNode = actTail
	for expNode != nil || actNode != nil {

		require.NotNil(t, expNode, errMsg(i, expNode, actNode))
		require.NotNil(t, actNode, errMsg(i, expNode, actNode))
		require.Equal(t, expNode.data, actNode.data, errMsg(i, expNode, actNode))

		expNode = expNode.prev
		actNode = actNode.prev
		i--
	}
}

func TestChain(t *testing.T) {
	a, b, c, d := dummyNodes()
	head, tail, size := chain(a, b, c, d)

	require.Equal(t, a, head)
	require.Equal(t, b, head.next)
	require.Equal(t, c, head.next.next)
	require.Equal(t, d, head.next.next.next)
	require.Nil(t, head.next.next.next.next)

	require.Equal(t, d, tail)
	require.Equal(t, c, tail.prev)
	require.Equal(t, b, tail.prev.prev)
	require.Equal(t, a, tail.prev.prev.prev)
	require.Nil(t, tail.prev.prev.prev.prev)

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
