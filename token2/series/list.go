package series

import (
	"github.com/PaulioRandall/scarlet-go/token2/lexeme"
)

type list struct {
	size int
	head *node
	tail *node
}

// Size returns the length of the Series.
func (li *list) Size() int {
	return li.size
}

// Empty returns true if the size of the Series is 0.
func (li *list) Empty() bool {
	return li.size == 0
}

func (li *list) prepend(l lexeme.Lexeme) {
	n := &node{
		data: l,
	}
	link(n, li.head)
	li.inserted(n)
}

func (li *list) append(l lexeme.Lexeme) {
	n := &node{
		data: l,
	}
	link(li.tail, n)
	li.inserted(n)
}

func (li *list) inserted(n *node) {

	if li.size == 0 {
		li.head = n
		li.tail = n
		li.size = 1
		return
	}

	li.size++
	if n.next == li.head {
		li.head = n
	}

	if n.prev == li.tail {
		li.tail = n
	}
}

func (li *list) removing(n *node) {
	li.size--

	if n == li.head {
		li.head = n.next
	}

	if n == li.tail {
		li.tail = n.prev
	}
}
