package container

import (
	"strings"

	"github.com/PaulioRandall/scarlet-go/token"
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

func (c *Container) Top() token.Lexeme {
	return c.headNode()
}

func (c *Container) Push(l token.Lexeme) {
	c.prepend(&node{
		data: l,
	})
}

func (c *Container) Pop() token.Lexeme {
	return c.popLex()
}

func (c *Container) Head() token.Lexeme {
	return c.headNode()
}

func (c *Container) Put(l token.Lexeme) {
	c.append(&node{
		data: l,
	})
}

func (c *Container) Take() token.Lexeme {
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

func (c *Container) headNode() token.Lexeme {
	if c.head == nil {
		return token.Lexeme{}
	}
	return c.head.data
}

func (c *Container) popLex() token.Lexeme {
	if n := c.pop(); n != nil {
		return n.data
	}
	return token.Lexeme{}
}

func (c *Container) pop() *node {

	if c.size == 0 {
		return nil
	}

	n := c.head
	c.head = c.head.next
	n.remove()
	c.size--

	if c.size == 0 {
		c.tail = nil
	}

	return n
}

func (c *Container) prepend(n *node) {

	if c.size == 0 {
		c.head = n
		c.tail = n
		c.size = 1
		return
	}

	link(n, c.head)
	c.head = n
	c.size++
}

func (c *Container) append(n *node) {

	if c.size == 0 {
		c.head = n
		c.tail = n
		c.size = 1
		return
	}

	link(c.tail, n)
	c.tail = n
	c.size++
}
