package container

import (
	"strings"

	"github.com/PaulioRandall/scarlet-go/token2/lexeme"
)

type Container struct {
	size int
	head *node
	tail *node
}

func New() *Container {
	return &Container{}
}

func (c *Container) Iterator() *Iterator {
	return &Iterator{
		con:  c,
		next: c.head,
	}
}

func (c *Container) Empty() bool {
	return c.size == 0
}

func (c *Container) More() bool {
	return c.size > 0
}

func (c *Container) Size() int {
	return c.size
}

func (c *Container) Top() lexeme.Lexeme {
	return c.headNode()
}

func (c *Container) Push(l lexeme.Lexeme) {
	c.prepend(&node{
		data: l,
	})
}

func (c *Container) Pop() lexeme.Lexeme {
	return c.popLex()
}

func (c *Container) Head() lexeme.Lexeme {
	return c.headNode()
}

func (c *Container) Put(l lexeme.Lexeme) {
	c.append(&node{
		data: l,
	})
}

func (c *Container) Take() lexeme.Lexeme {
	return c.popLex()
}

func (c *Container) String() string {

	var sb strings.Builder

	for n := c.head; n != nil; n = n.next {
		if n != c.head {
			sb.WriteRune('\n')
		}

		sb.WriteString(n.data.String())
	}

	return sb.String()
}

func (c *Container) headNode() lexeme.Lexeme {
	if c.head == nil {
		return lexeme.Lexeme{}
	}
	return c.head.data
}

func (c *Container) popLex() lexeme.Lexeme {
	if n := c.pop(); n != nil {
		return n.data
	}
	return lexeme.Lexeme{}
}

func (c *Container) pop() *node {

	if c.size == 0 {
		return nil
	}

	n := c.head
	c.remove(n)
	return n
}

func (c *Container) prepend(n *node) {
	link(n, c.head)
	c.inserted(n)
}

func (c *Container) append(n *node) {
	link(c.tail, n)
	c.inserted(n)
}

func (c *Container) insertBefore(ref, n *node) {

	var prev *node

	if ref != nil {
		prev = ref.prev
		unlink(prev, ref)
	}

	chain(prev, n, ref)
	c.inserted(n)
}

func (c *Container) insertAfter(ref, n *node) {

	var next *node

	if ref != nil {
		next = ref.next
		unlink(ref, next)
	}

	chain(ref, n, next)
	c.inserted(n)
}

func (c *Container) inserted(n *node) {

	if c.size == 0 {
		c.head = n
		c.tail = n
		c.size = 1
		return
	}

	if n.next == c.head {
		c.head = n
	}

	if n.prev == c.tail {
		c.tail = n
	}

	c.size++
}

func (c *Container) remove(n *node) {

	if n == c.head {
		c.head = n.next
	}

	if n == c.tail {
		c.tail = n.prev
	}

	c.size--
	n.remove()
}
