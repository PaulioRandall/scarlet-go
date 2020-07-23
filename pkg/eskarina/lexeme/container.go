package lexeme

import (
	"fmt"
)

type Collection interface {
	Empty() bool
	More() bool
	Size() int
}

type List interface {
	Collection
	Get(int) *Lexeme
}

type MutList interface {
	List
	Prepend(*Lexeme)
	Append(*Lexeme)
	InsertBefore(int, *Lexeme)
	InsertAfter(int, *Lexeme)
	Remove(int) *Lexeme
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

func NewContainer(head *Lexeme) *Container {

	if head == nil {
		return &Container{}
	}

	if head.Prev != nil {
		panic("Can't use Lexeme as head since it's not the head of its linked list")
	}

	c := &Container{
		size: 1,
		head: head,
		tail: head,
	}

	for c.tail.Next != nil {
		c.tail = c.tail.Next
		c.size++
	}

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

func (c *Container) Get(idx int) *Lexeme {

	var node *Lexeme

	if idx < 0 || c.size == 0 {
		goto ERROR
	}

	node = c.head
	for i := 0; node != nil && i < idx; i++ {
		node = node.Next
	}

	if node == nil {
		goto ERROR
	}

	return node

ERROR:
	e := fmt.Errorf("Index out of range: len %d, idx %d", c.size, idx)
	panic(e)
}

func (c *Container) Prepend(lex *Lexeme) {

	lex.Remove()

	if c.head == nil {
		c.head = lex
		c.tail = lex
		c.size = 1
		return
	}

	c.head.Prepend(lex)
	c.head = lex
	c.size++
}

func (c *Container) Append(lex *Lexeme) {

	lex.Remove()

	if c.tail == nil {
		c.tail = lex
		c.head = lex
		c.size = 1
		return
	}

	c.tail.Append(lex)
	c.tail = lex
	c.size++
}

func (c *Container) InsertBefore(idx int, lex *Lexeme) {

	if idx == 0 {
		c.Prepend(lex)
		return
	}

	node := c.Get(idx)
	lex.Remove()

	node.Prepend(lex)
	c.size++
}

func (c *Container) InsertAfter(idx int, lex *Lexeme) {

	if idx == c.size-1 {
		c.Append(lex)
		return
	}

	node := c.Get(idx)
	lex.Remove()

	node.Append(lex)
	c.size++
}

func (c *Container) Remove(idx int) *Lexeme {

	node := c.Get(idx)

	if idx == 0 {
		c.head = c.head.Next
	} else if idx == c.size-1 {
		c.tail = c.tail.Prev
	}

	node.Remove()
	c.size--

	if c.size == 0 {
		c.head, c.tail = nil, nil
	}

	return node
}

func (c *Container) Top() *Lexeme {
	return c.Get(0)
}

func (c *Container) Push(lex *Lexeme) {
	c.Prepend(lex)
}

func (c *Container) Pop() *Lexeme {
	return c.Remove(0)
}

func (c *Container) Head() *Lexeme {
	return c.Get(0)
}

func (c *Container) Put(lex *Lexeme) {
	c.Append(lex)
}

func (c *Container) Take() *Lexeme {
	return c.Remove(0)
}

func (c *Container) Descend(f func(*Lexeme)) {
	for lex := c.head; lex != nil; lex = lex.Next {
		f(lex)
	}
}
