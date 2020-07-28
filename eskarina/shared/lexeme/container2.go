package lexeme

import (
	"fmt"
)

type Collection2 interface {
	Empty() bool
	More() bool
	Size() int
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

func NewContainer2(head *Lexeme) *Container2 {

	if head == nil {
		return &Container2{}
	}

	checkIsSingle(head)

	return &Container2{
		size: 1,
		head: head,
		tail: head,
	}
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
		c.tail = c.tail.Next
	} else {
		r = c.head
		c.head = c.head.Next
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

func (c *Container2) Descend(f func(*Lexeme)) {
	for lex := c.head; lex != nil; lex = lex.Next {
		f(lex)
	}
}
