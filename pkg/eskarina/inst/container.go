package inst

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
	Get(int) *Instruction
}

type Queue interface {
	Collection
	Head() *Instruction
	Put(*Instruction)
	Take() *Instruction
}

type Container struct {
	size int
	head *Instruction
	tail *Instruction
}

func NewContainer(head *Instruction) *Container {

	if head == nil {
		return &Container{}
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

func (c *Container) Get(idx int) *Instruction {

	var node *Instruction

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

func (c *Container) Head() *Instruction {
	return c.head
}

func (c *Container) Put(in *Instruction) {

	in.Next = nil

	if c.tail == nil {
		c.tail = in
		c.head = in
		c.size = 1
		return
	}

	c.tail.Next = in
	c.tail = in
	c.size++
}

func (c *Container) Take() *Instruction {

	in := c.Get(0)
	c.head = c.head.Next
	c.size--

	if c.size == 0 {
		c.head, c.tail = nil, nil
	}

	in.Next = nil
	return in
}

func (c *Container) Descend(f func(*Instruction)) {
	for in := c.head; in != nil; in = in.Next {
		f(in)
	}
}
