package lexeme

import (
	"fmt"
	"strings"
)

type Stack interface {
	AsContainer() *Container
	ToIterator() *Iterator
	Empty() bool
	More() bool
	Size() int
	Top() *Lexeme
	Push(*Lexeme)
	Pop() *Lexeme
	String() string
}

type Queue interface {
	AsContainer() *Container
	ToIterator() *Iterator
	Empty() bool
	More() bool
	Size() int
	Head() *Lexeme
	Put(*Lexeme)
	Take() *Lexeme
	String() string
}

type Container struct {
	size int
	head *Lexeme
	tail *Lexeme
}

func checkIsSingle(lex *Lexeme) {
	if lex.next != nil || lex.prev != nil {
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

// @Retired
func (c *Container) ToIterator() *Iterator {
	head := c.head
	c.head, c.tail, c.size = nil, nil, 0
	return NewIterator(head)
}

func (c *Container) Iterator() *Iterator2 {
	return &Iterator2{
		after: c.head,
	}
}

func (c *Container) AsContainer() *Container {
	return c
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

	remove(r)
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
		append(c.tail, lex)
		c.tail = lex
	} else {
		prepend(c.head, lex)
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

func (c *Container) remove(lex *Lexeme) {

	if c.head == lex {
		c.head = lex.next
	}

	if c.tail == lex {
		c.tail = lex.prev
	}

	remove(lex)
	c.size--
}

func (c *Container) insertBefore(base *Lexeme, add *Lexeme) {

	prepend(base, add)
	c.size++

	if base == c.head {
		c.head = add
	}

	if base == c.tail {
		c.tail = add
	}
}

func (c *Container) insertAfter(base *Lexeme, add *Lexeme) {

	append(base, add)
	c.size++

	if base == c.head {
		c.head = add
	}

	if base == c.tail {
		c.tail = add
	}
}
