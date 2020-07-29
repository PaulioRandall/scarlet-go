package lexeme

import (
	"fmt"
	"strings"
)

type Collection interface {
	To() *To
	Empty() bool
	More() bool
	Size() int
}

type Stack interface {
	Collection
	Top() *Lexeme
	Push(*Lexeme)
	Pop() *Lexeme
}

type Queue interface {
	Collection
	Head() *Lexeme
	Put(*Lexeme)
	Take() *Lexeme
}

type Container struct {
	size int
	head *Lexeme
	tail *Lexeme
}

func checkIsSingle(lex *Lexeme) {
	if !lex.IsSingle() {
		m := fmt.Sprintf(
			"Lexeme `%s` is already part of another collection, remove first",
			lex.String(),
		)
		panic(m)
	}
}

func NewContainer(head *Lexeme) *Container {

	if head == nil {
		return &Container{}
	}

	c := &Container{
		head: head,
		tail: head,
		size: 1,
	}

	for lex := c.head.next; lex != nil; lex = lex.next {
		c.tail = lex
		c.size++
	}

	return c
}

func (c *Container) vacate() *Lexeme {
	head := c.head
	c.head, c.tail, c.size = nil, nil, 0
	return head
}

func (c *Container) To() *To {
	return &To{
		b: c,
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

func (c *Container) Top() *Lexeme {
	return c.head
}

func (c *Container) Push(lex *Lexeme) {
	c.push(lex, false)
}

func (c *Container) Pop() *Lexeme {
	return c.pop(false)
}

func (c *Container) Head() *Lexeme {
	return c.head
}

func (c *Container) Put(lex *Lexeme) {
	c.push(lex, true)
}

func (c *Container) Take() *Lexeme {
	return c.pop(false)
}

func (c *Container) pop(fromBack bool) *Lexeme {

	if c.size == 0 {
		return nil
	}

	var r *Lexeme
	if fromBack {
		r = c.tail
		c.tail = c.tail.prev
	} else {
		r = c.head
		c.head = c.head.next
	}

	r.remove()
	c.size--

	if c.size == 0 {
		c.head, c.tail = nil, nil
	}

	return r
}

func (c *Container) push(lex *Lexeme, toBack bool) {

	checkIsSingle(lex)

	if c.size == 0 {
		c.head = lex
		c.tail = lex
		c.size = 1
		return
	}

	if toBack {
		c.tail.append(lex)
		c.tail = lex
	} else {
		c.head.prepend(lex)
		c.head = lex
	}

	c.size++
}

func (c *Container) String() string {

	sb := strings.Builder{}

	for lex := c.head; lex != nil; lex = lex.next {

		if lex != c.head {
			sb.WriteRune('\n')
		}

		sb.WriteString(lex.String())
	}

	return sb.String()
}
