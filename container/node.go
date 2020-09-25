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
	unlinkAll(n.prev, n, n.next)
}

func (n *node) remove() {
	link(n.prev, n.next)
	n.prev, n.next = nil, nil
}

func link(a, b *node) {

	if a != nil {
		a.next = b
	}

	if b != nil {
		b.prev = a
	}
}

func unlink(a, b *node) {

	if a != nil {
		a.next = nil
	}

	if b != nil {
		b.prev = nil
	}
}

func chain(nodes ...*node) (head, tail *node, size int) {

	for _, n := range nodes {

		if n == nil {
			continue
		}

		size++

		if head == nil {
			head, tail = n, n
			continue
		}

		link(tail, n)
		tail = n
	}

	return
}

func unlinkAll(nodes ...*node) {

	var prev *node

	for _, n := range nodes {

		if n == nil {
			continue
		}

		if prev == nil {
			prev = n
			continue
		}

		unlink(prev, n)
		prev = n
	}
}
