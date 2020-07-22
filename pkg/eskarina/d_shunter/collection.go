package shunter

import (
	"github.com/PaulioRandall/scarlet-go/pkg/eskarina/lexeme"
	"github.com/PaulioRandall/scarlet-go/pkg/eskarina/perror"
)

type Collection interface {
	Empty() bool
	More() bool
}

type List interface {
	Collection
	Size() int
	At(int) *lexeme.Lexeme
}

type MutableList interface {
	List
	Prepend(*lexeme.Lexeme)
	Append(*lexeme.Lexeme)
	InsertBefore(int, *lexeme.Lexeme)
	InsertAfter(int, *lexeme.Lexeme)
	Remove(int) *lexeme.Lexeme
}

type Stack interface {
	Collection
	Top() *lexeme.Lexeme
	Push(*lexeme.Lexeme)
	Pop() *lexeme.Lexeme
}

type Queue interface {
	Collection
	Front() *lexeme.Lexeme
	Put(*lexeme.Lexeme)
	Take() *lexeme.Lexeme
}

type OrderedCollection struct {
	Collection
	List
	MutableList
	Stack
	Queue
	size  int
	first *lexeme.Lexeme
	last  *lexeme.Lexeme
}

func NewOrderedCollection(first *lexeme.Lexeme) *OrderedCollection {

	if first == nil {
		return &OrderedCollection{}
	}

	if first.Prev != nil {
		perror.Panic("Can't use as collection front as it has a previous link")
	}

	c := &OrderedCollection{
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

func (c *OrderedCollection) Empty() bool {
	return c.size == 0
}

func (c *OrderedCollection) More() bool {
	return c.size > 0
}

func (c *OrderedCollection) Size() int {
	return c.size
}

func (c *OrderedCollection) At(idx int) *lexeme.Lexeme {

	if idx < 0 || c.size == 0 {
		perror.Panic("Index out of range: len %d, idx %d", c.size, idx)
	}

	node := c.first
	for i := 0; node != nil && i < idx; i++ {
		node = node.Next
	}

	if node != nil {
		perror.Panic("Index out of range: len %d, idx %d", c.size, idx)
	}

	return node
}

func (c *OrderedCollection) Prepend(lex *lexeme.Lexeme) {

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

func (c *OrderedCollection) Append(lex *lexeme.Lexeme) {

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

func (c *OrderedCollection) InsertBefore(idx int, lex *lexeme.Lexeme) {

	if idx == 0 {
		c.Prepend(lex)
		return
	}

	node := c.At(idx)
	lex.Remove()

	node.Prepend(lex)
	c.size++
}

func (c *OrderedCollection) InsertAfter(idx int, lex *lexeme.Lexeme) {

	if idx == c.size-1 {
		c.Append(lex)
		return
	}

	node := c.At(idx)
	lex.Remove()

	node.Append(lex)
	c.size++
}

func (c *OrderedCollection) Remove(idx int) *lexeme.Lexeme {

	node := c.At(idx)

	if idx == 0 {
		c.first = c.first.Next
	} else if idx == c.size-1 {
		c.last = c.last.Prev
	}

	node.Remove()
	c.size--

	return node
}

func (c *OrderedCollection) Top() *lexeme.Lexeme {
	return c.At(0)
}

func (c *OrderedCollection) Push(lex *lexeme.Lexeme) {
	c.Prepend(lex)
}

func (c *OrderedCollection) Pop() *lexeme.Lexeme {
	return c.Remove(0)
}

func (c *OrderedCollection) Front() *lexeme.Lexeme {
	return c.At(0)
}

func (c *OrderedCollection) Put(lex *lexeme.Lexeme) {
	c.Append(lex)
}

func (c *OrderedCollection) Take() *lexeme.Lexeme {
	return c.Remove(0)
}

func (c *OrderedCollection) Descend(f func(*lexeme.Lexeme)) {
	for lex := c.first; lex != nil; lex = lex.Next {
		f(lex)
	}
}
