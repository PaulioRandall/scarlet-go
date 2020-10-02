package series

import (
	"fmt"
	"testing"

	"github.com/PaulioRandall/scarlet-go/token2/lexeme"
	"github.com/PaulioRandall/scarlet-go/token2/token"

	"github.com/stretchr/testify/require"
)

func dummyLexemes() (l1, l2, l3, l4 lexeme.Lexeme) {
	l1 = lexeme.Tok("true", token.TRUE)
	l2 = lexeme.Tok("1", token.NUMBER)
	l3 = lexeme.Tok("abc", token.STRING)
	l4 = lexeme.Tok("i", token.IDENT)
	return
}

func dummyNodes() (n1, n2, n3, n4 *node) {
	l1, l2, l3, l4 := dummyLexemes()
	n1 = &node{data: l1}
	n2 = &node{data: l2}
	n3 = &node{data: l3}
	n4 = &node{data: l4}
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

	var expNode, actNode = exp, act
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
	n1, n2, n3, n4 := dummyNodes()
	head, tail, size := chain(n1, n2, n3, n4)

	require.Equal(t, n1, head)
	require.Equal(t, n2, head.next)
	require.Equal(t, n3, head.next.next)
	require.Equal(t, n4, head.next.next.next)
	require.Nil(t, head.next.next.next.next)

	require.Equal(t, n4, tail)
	require.Equal(t, n3, tail.prev)
	require.Equal(t, n2, tail.prev.prev)
	require.Equal(t, n1, tail.prev.prev.prev)
	require.Nil(t, tail.prev.prev.prev.prev)

	require.Equal(t, 4, size)
}

func TestUnlinkAll(t *testing.T) {

	n1, n2, n3, n4 := dummyNodes()
	n1.next = n2
	n2.next = n3
	n3.next = n4
	n2.prev = n1
	n3.prev = n2
	n4.prev = n3

	unlinkAll(n1, n2, n3)
	require.Nil(t, n1.prev)
	require.Nil(t, n1.next)
	require.Nil(t, n2.prev)
	require.Nil(t, n2.next)
	require.Nil(t, n3.prev)
	require.Equal(t, n4, n3.next)
}
