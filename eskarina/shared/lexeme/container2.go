package lexeme

import (
	"fmt"
	"strings"
)

type Collection2 interface {
	Empty() bool
	More() bool
	Size() int
	Container() *Container2
}

type Stack2 interface {
	Collection2
	Top() *Lexeme
	Push(*Lexeme)
	Pop() *Lexeme
}

type Queue2 interface {
	Collection2
	Head() *Lexeme
	Put(*Lexeme)
	Take() *Lexeme
}

type Container2 struct {
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

func NewContainer2() *Container2 {
	return &Container2{}
}

func (c *Container2) AsStack() Stack2 {
	return c
}

func (c *Container2) AsQueue() Queue2 {
	return c
}

func (c *Container2) Empty() bool {
	return c.size == 0
}

func (c *Container2) More() bool {
	return c.size > 0
}

func (c *Container2) Size() int {
	return c.size
}

func (c *Container2) Container() *Container2 {
	return c
}

func (c *Container2) Top() *Lexeme {
	return c.head
}

func (c *Container2) Push(lex *Lexeme) {
	c.push(lex, false)
}

func (c *Container2) Pop() *Lexeme {
	return c.pop(false)
}

func (c *Container2) Head() *Lexeme {
	return c.head
}

func (c *Container2) Put(lex *Lexeme) {
	c.push(lex, true)
}

func (c *Container2) Take() *Lexeme {
	return c.pop(false)
}

func (c *Container2) pop(fromBack bool) *Lexeme {

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

func (c *Container2) push(lex *Lexeme, toBack bool) {

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

func (c *Container2) String() string {

	sb := strings.Builder{}

	for lex := c.head; lex != nil; lex = lex.Next {

		if lex != c.head {
			sb.WriteRune('\n')
		}

		sb.WriteString(lex.String())
	}

	return sb.String()
}
