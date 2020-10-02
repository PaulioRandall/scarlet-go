package series

import (
	"strings"

	"github.com/PaulioRandall/scarlet-go/token2/lexeme"
)

type List interface {
	Empty() bool
	More() bool
	Size() int
	Prepend(lexeme.Lexeme)
	Append(lexeme.Lexeme)
	String() string
}

type Iterator interface {
	JumpToStart()
	JumpToEnd()
	JumpToPrev(matcher func(ReadOnly) bool) bool
	JumpToNext(matcher func(ReadOnly) bool) bool
	Next() lexeme.Lexeme
	Get() lexeme.Lexeme
	Prev() lexeme.Lexeme
	LookAhead() lexeme.Lexeme
	LookBack() lexeme.Lexeme
	InsertAfter()
	InsertBefore()
	Remove() lexeme.Lexeme
	String() string
}

type ReadOnly interface {
	Empty() bool
	More() bool
	Size() int
	Get() lexeme.Lexeme
	LookAhead() lexeme.Lexeme
	LookBack() lexeme.Lexeme
	String() string
}

type Series struct {
	size int
	head *node
	tail *node
	prev *node
	curr *node
	next *node
}

func New() *Series {
	return &Series{}
}

func new(nodes ...*node) *Series {
	head, tail, size := chain(nodes...)
	return &Series{
		size: size,
		head: head,
		tail: tail,
		next: head,
	}
}

func (s *Series) Empty() bool {
	return s.size == 0
}

func (s *Series) More() bool {
	return s.size > 0
}

func (s *Series) Size() int {
	return s.size
}

func (s *Series) JumpToStart() {
	s.prev = nil
	s.curr = nil
	s.next = s.head
}

func (s *Series) JumpToEnd() {
	s.prev = s.tail
	s.curr = nil
	s.next = nil
}

func (s *Series) JumpToPrev(matcher func(ReadOnly) bool) bool {

	for n := s.prev; n != nil; n = n.prev {
		s.jumpTo(n)
		if matcher(s) {
			return true
		}
	}

	s.JumpToStart()
	return false
}

func (s *Series) JumpToNext(matcher func(ReadOnly) bool) bool {

	for n := s.next; n != nil; n = n.next {
		s.jumpTo(n)
		if matcher(s) {
			return true
		}
	}

	s.JumpToEnd()
	return false
}

func (s *Series) Next() lexeme.Lexeme {
	if s.next == nil {
		panic("Can't move beyond the end of the series")
	}
	s.jumpTo(s.next)
	return s.curr.data
}

func (s *Series) Get() lexeme.Lexeme {
	if s.curr == nil {
		return lexeme.Lexeme{}
	}
	return s.curr.data
}

func (s *Series) Prev() lexeme.Lexeme {
	if s.prev == nil {
		panic("Can't move beyond the start of the series")
	}
	s.jumpTo(s.prev)
	return s.curr.data
}

func (s *Series) LookAhead() lexeme.Lexeme {
	if s.next == nil {
		return lexeme.Lexeme{}
	}
	return s.next.data
}

func (s *Series) LookBack() lexeme.Lexeme {
	if s.prev == nil {
		return lexeme.Lexeme{}
	}
	return s.prev.data
}

func (s *Series) Prepend(l lexeme.Lexeme) {

	n := &node{
		data: l,
	}

	link(n, s.head)
	s.inserted(n)

	if s.curr != nil {
		s.jumpTo(s.curr)
	} else {
		s.JumpToStart()
	}
}

func (s *Series) Append(l lexeme.Lexeme) {

	n := &node{
		data: l,
	}

	link(s.tail, n)
	s.inserted(n)

	if s.curr != nil {
		s.jumpTo(s.curr)
	} else {
		s.JumpToEnd()
	}
}

func (s *Series) InsertAfter(l lexeme.Lexeme) {

	n := &node{
		data: l,
	}

	if s.curr == nil {
		panic("Current node missing, can't insert after it")
	}

	var next *node
	if s.curr != nil {
		next = s.curr.next
		unlink(s.curr, next)
	}

	chain(s.curr, n, next)
	s.inserted(n)
	s.jumpTo(s.curr)
}

func (s *Series) InsertBefore(l lexeme.Lexeme) {

	n := &node{
		data: l,
	}

	if s.curr == nil {
		panic("Current node missing, can't insert before it")
	}

	var prev *node
	if s.curr != nil {
		prev = s.curr.prev
		unlink(prev, s.curr)
	}

	chain(prev, n, s.curr)
	s.inserted(n)
	s.jumpTo(s.curr)
}

func (s *Series) Remove() lexeme.Lexeme {

	if s.curr == nil {
		return lexeme.Lexeme{}
	}

	n := s.curr
	s.curr = nil
	s.size--

	if n == s.head {
		s.head = n.next
	}

	if n == s.tail {
		s.tail = n.prev
	}

	n.remove()
	return n.data
}

func (s *Series) String() string {

	var sb strings.Builder
	for n := s.head; n != nil; n = n.next {
		if n != s.head {
			sb.WriteRune('\n')
		}

		sb.WriteString(n.data.String())
	}

	return sb.String()
}

func (s *Series) jumpTo(n *node) {
	s.prev = n.prev
	s.curr = n
	s.next = n.next
}

func (s *Series) inserted(n *node) {

	if s.size == 0 {
		s.head = n
		s.tail = n
		s.size = 1
		return
	}

	s.size++
	if n.next == s.head {
		s.head = n
	}

	if n.prev == s.tail {
		s.tail = n
	}
}
