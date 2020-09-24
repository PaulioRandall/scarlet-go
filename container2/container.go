package container2

import (
	//"fmt"
	//"strings"

	"github.com/PaulioRandall/scarlet-go/token"
)

type node struct {
	data token.Lexeme
	next *node
	prev *node
}

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
	if c.head == nil {
		return token.Lexeme{}
	}
	return c.head.data
}

func (c *Container) Push(l token.Lexeme) {
	c.prepend(l)
}

/*
func (c *Container) Pop() *Lexeme {
	return c.pop(false)
}

func (c *Container) Head() *Lexeme {
	return c.head
}

func (c *Container) Put(lex *Lexeme) {
	c.append(l)
}

func (c *Container) Take() *Lexeme {
	return c.pop(false)
}

func (c *Container) String() string {

	var sb strings.Builder

	for lex := c.head; lex != nil; lex = lex.next {

		if lex != c.head {
			sb.WriteRune('\n')
		}

		sb.WriteString(lex.String())
	}

	return sb.String()
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
*/
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

/*
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

	if c.head == base {
		c.head = add
	}
}

func (c *Container) insertAfter(base *Lexeme, add *Lexeme) {

	append(base, add)
	c.size++

	if base == c.tail {
		c.tail = add
	}
}

func (lex *Lexeme) checkIsSingle() {
	if lex.next != nil || lex.prev != nil {
		m := fmt.Sprintf(
			"Lexeme `%s` is already part of another collection, remove first",
			lex.String(),
		)
		panic(m)
	}
}
*/
