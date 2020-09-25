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
	c.prepend(l)
}

func (c *Container) Pop() token.Lexeme {
	return c.pop()
}

func (c *Container) Head() token.Lexeme {
	return c.headNode()
}

func (c *Container) Put(l token.Lexeme) {
	c.append(l)
}

func (c *Container) Take() token.Lexeme {
	return c.pop()
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

func (c *Container) pop() token.Lexeme {

	if c.size == 0 {
		return token.Lexeme{}
	}

	n := c.head
	c.head = c.head.next
	n.remove()
	c.size--

	if c.size == 0 {
		c.tail = nil
	}

	return n.data
}

func (c *Container) prepend(l token.Lexeme) {

	n := &node{
		data: l,
	}

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

func (c *Container) append(l token.Lexeme) {

	n := &node{
		data: l,
	}

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
