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

/*
func (c *Container) Iterator() *Iterator {
	return &Iterator{
		con:   c,
		after: c.head,
	}
}

func (c *Container) AsContainer() *Container {
	return c
}
*/

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

	var r token.Lexeme

	if c.size == 0 {
		return r
	}

	r = c.head.data
	c.head = c.head.next
	c.size--

	if c.size == 0 {
		c.tail = nil
	} else {
		c.head.prev = nil
	}

	return r
}

func (c *Container) prepend(l token.Lexeme) {

	n := &node{
		data: l,
		next: c.head,
	}

	if c.size == 0 {
		c.head = n
		c.tail = n
		c.size = 1
		return
	}

	c.head.prev = n
	c.head = n
	c.size++
}

func (c *Container) append(l token.Lexeme) {

	n := &node{
		data: l,
		prev: c.tail,
	}

	if c.size == 0 {
		c.head = n
		c.tail = n
		c.size = 1
		return
	}

	c.tail.next = n
	c.tail = n
	c.size++
}
