package container

import (
	"github.com/PaulioRandall/scarlet-go/token"
)

type node struct {
	data token.Lexeme
	next *node
	prev *node
}

func (n *node) unlink() {

	if n.prev != nil {
		n.prev.next = nil
	}

	if n.next != nil {
		n.next.prev = nil
	}
}

func (n *node) remove() {

	if n.prev != nil {
		n.prev.next = n.next
	}

	if n.next != nil {
		n.next.prev = n.prev
	}
}

func link(a, b *node) {
	a.next = b
	b.prev = a
}

func unlink(a, b *node) {
	a.next = nil
	b.prev = nil
}

func linkAll(nodes ...*node) (head, tail *node) {

	for _, n := range nodes {

		if head == nil {
			head, tail = n, n
			continue
		}

		link(tail, n)
	}

	return head, tail
}

func unlinkAll(nodes ...*node) {
	for _, n := range nodes {
		n.unlink()
	}
}
