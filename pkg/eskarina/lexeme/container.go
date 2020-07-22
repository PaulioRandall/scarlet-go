package lexeme

import (
	"fmt"

	"github.com/PaulioRandall/scarlet-go/pkg/eskarina/prop"
)

type Collection interface {
	Empty() bool
	More() bool
}

type List interface {
	Collection
	Size() int
	Get(int) *Lexeme
}

type MutableList interface {
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
	Front() *Lexeme
	Put(*Lexeme)
	Take() *Lexeme
}

// TODO
type TokenStream interface {
	Collection
	Token
	Accept(...prop.Prop) *Lexeme
	Expect(func([]prop.Prop) error, ...prop.Prop) (*Lexeme, error)
}

type Container struct {
	Collection
	List
	MutableList
	Stack
	Queue
	size  int
	first *Lexeme
	last  *Lexeme
}

func NewContainer(first *Lexeme) *Container {

	if first == nil {
		return &Container{}
	}

	if first.Prev != nil {
		panic("Can't use Lexeme as container as its not the head of a list")
	}

	c := &Container{
		size:  1,
		first: first,
		last:  first,
	}

	for c.last.Next != nil {
		c.last = c.last.Next
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

	node = c.first
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

	if c.first == nil {
		c.first = lex
		c.last = lex
		c.size = 1
		return
	}

	c.first.Prepend(lex)
	c.first = lex
	c.size++
}

func (c *Container) Append(lex *Lexeme) {

	lex.Remove()

	if c.last == nil {
		c.last = lex
		c.first = lex
		c.size = 1
		return
	}

	c.last.Append(lex)
	c.last = lex
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
		c.first = c.first.Next
	} else if idx == c.size-1 {
		c.last = c.last.Prev
	}

	node.Remove()
	c.size--

	if c.size == 0 {
		c.first, c.last = nil, nil
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

func (c *Container) Front() *Lexeme {
	return c.Get(0)
}

func (c *Container) Put(lex *Lexeme) {
	c.Append(lex)
}

func (c *Container) Take() *Lexeme {
	return c.Remove(0)
}

func (c *Container) Descend(f func(*Lexeme)) {
	for lex := c.first; lex != nil; lex = lex.Next {
		f(lex)
	}
}
