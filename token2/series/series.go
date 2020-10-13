// Series package contains the Series struct and useful receiving functions
// for representing both a double linked list and iterator of Lexemes.
package series

import (
	"strings"

	"github.com/PaulioRandall/scarlet-go/token2/lexeme"
)

// Matcher function signiture is used by Series search functions such as
// Series.JumpToNext and Series.JumpToPrev. True is returned if some state in
// the SnapShot matches the consumers criteria.
type Matcher func(Snapshot) bool

// Builder documents the functions that may be applicable to any interface
// that appends or prepends to a Series. Subsetting the interface is desirable.
type Builder interface {
	Size() int
	Empty() bool
	Prepend(lexeme.Lexeme)
	Append(lexeme.Lexeme)
	String() string
}

// Snapshot is designed for use when iterating a Series to mask state changing
// functionality.
type Snapshot interface {
	Empty() bool
	More() bool
	Size() int
	Get() lexeme.Lexeme
	LookAhead() lexeme.Lexeme
	LookBack() lexeme.Lexeme
	String() string
}

// Iterator documents the functions that may be applicable to any interface
// that iterates a Series. Subsetting the interface is desirable.
type Iterator interface {
	More() bool
	Empty() bool
	JumpToStart()
	JumpToEnd()
	JumpToPrev(matcher Matcher) bool
	JumpToNext(matcher Matcher) bool
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

// Series represents a double linked list and iterator of Lexemes. The decision
// to combine the list and iteration functionality was not taken lightly but
// based on previous use, the separation of concerns appeared to cost more in
// usage complexity than an amalgamated structure did for implementation
// complexity. Could be wrong, only time will really tell. Splitting the structs
// shouldn't be much trouble if it comes to it.
//
// A Series always starts with its iterator mark before the first item such
// that Series.Get will return an empty Lexeme and Series.Next will move to the
// first item, if it exists.
//
// Series is best used through custom interfaces that limit functionality to
// the situation for readability. Series is not concurrent friendly.
type Series struct {
	list
	mark
}

// Make returns a new empty Series.
func Make() *Series {
	return &Series{
		list: list{},
		mark: mark{},
	}
}

func makeWith(nodes ...*node) *Series {
	head, tail, size := chain(nodes...)
	return &Series{
		list: list{
			size: size,
			head: head,
			tail: tail,
		},
		mark: mark{
			next: head,
		},
	}
}

// Size returns the length of the Series.
func (s *Series) Size() int {
	return s.list.size
}

// Empty returns true if the size of the Series is 0.
func (s *Series) Empty() bool {
	return s.list.size == 0
}

// JumpToStart resets the iterator mark before the first item in the Series.
func (s *Series) JumpToStart() {
	s.mark.jumpToStart(s.list.head)
}

// JumpToEnd puts the iterator mark after the last item in the Series.
func (s *Series) JumpToEnd() {
	s.mark.jumpToEnd(s.list.tail)
}

// JumpToPrev iterates backwards calling 'matcher' on each item on the way. If
// 'matcher' returns true then iteration stops and true is returned, if no match
// is found the iterator mark will end up before the first item in the Series
// and false is returned.
func (s *Series) JumpToPrev(matcher Matcher) bool {

	for n := s.mark.prev; n != nil; n = n.prev {
		s.mark.jumpTo(n)
		if matcher(s) {
			return true
		}
	}

	s.mark.jumpToStart(s.list.head)
	return false
}

// JumpToNext iterates forwards calling 'matcher' on each item on the way. If
// 'matcher' returns true then iteration stops and true is returned, if no match
// is found the iterator mark will end up after the last item in the Series
// and false is returned.
func (s *Series) JumpToNext(matcher Matcher) bool {

	for n := s.mark.next; n != nil; n = n.next {
		s.mark.jumpTo(n)
		if matcher(s) {
			return true
		}
	}

	s.mark.jumpToEnd(s.list.tail)
	return false
}

// Next moves the iterator mark onto the next item and returns it. A panic will
// ensue if the end of the iterator has already been reached so Series.More
// should be called before hand.
func (s *Series) Next() lexeme.Lexeme {
	return s.mark.nextLex()
}

// Prev moves the iterator mark onto the previous item and returns it. A panic
// will ensue if the start of the iterator has already been reached.
func (s *Series) Prev() lexeme.Lexeme {
	return s.mark.prevLex()
}

// Get returns the item at the current iterator mark or the Lexeme zero value if
// there is no item at the mark, Iie. before the first item, after the last
// item, and immediately after an item has been removed.
func (s *Series) Get() lexeme.Lexeme {
	if s.mark.curr == nil {
		return lexeme.Lexeme{}
	}
	return s.mark.curr.data
}

// LookAhead returns the Lexeme next in the iteration without incrementing the
// iterator mark. An empty Lexeme is returned if there is no item ahead.
func (s *Series) LookAhead() lexeme.Lexeme {
	if s.mark.next == nil {
		return lexeme.Lexeme{}
	}
	return s.mark.next.data
}

// Lookback returns the Lexeme previous in the iteration without decrementing
// the iterator mark. An empty Lexeme is returned if there is no item behind.
func (s *Series) LookBack() lexeme.Lexeme {
	if s.mark.prev == nil {
		return lexeme.Lexeme{}
	}
	return s.mark.prev.data
}

// Prepend inserts a Lexeme at the front of the Series.
func (s *Series) Prepend(l lexeme.Lexeme) {
	atStart := s.mark.prev == nil && s.mark.curr == nil
	if s.list.prepend(l); atStart {
		s.mark.jumpToStart(s.list.head)
	} else if s.curr != nil {
		s.mark.jumpTo(s.mark.curr)
	} else {
		s.mark.jumpToEnd(s.list.tail)
	}
}

// Append inserts a Lexeme at the back of the Series.
func (s *Series) Append(l lexeme.Lexeme) {
	atStart := s.mark.prev == nil && s.mark.curr == nil
	if s.list.append(l); atStart {
		s.mark.jumpToStart(s.list.head)
	} else if s.mark.curr != nil {
		s.mark.jumpTo(s.mark.curr)
	} else {
		s.mark.jumpToEnd(s.list.tail)
	}
}

// InsertAfter inserts a Lexeme after the item indicated by the iterator mark.
// A panic will ensue if the mark isn't pointing to an item.
func (s *Series) InsertAfter(l lexeme.Lexeme) {

	if s.mark.curr == nil {
		panic("Current node missing, can't insert after it")
	}

	n := &node{data: l}
	s.mark.insertAfter(n)
	s.list.inserted(n)
}

// InsertBefore inserts a Lexeme before the item indicated by the iterator mark.
// A panic will ensue if the mark isn't pointing to an item.
func (s *Series) InsertBefore(l lexeme.Lexeme) {

	if s.mark.curr == nil {
		panic("Current node missing, can't insert before it")
	}

	n := &node{data: l}
	s.mark.insertBefore(n)
	s.list.inserted(n)
}

// Remove removes a the Lexeme indicated by the iterator mark from the Series.
func (s *Series) Remove() lexeme.Lexeme {

	if s.mark.curr == nil {
		return lexeme.Lexeme{}
	}

	n := s.mark.curr
	s.mark.curr = nil
	s.list.removing(n)
	n.remove()
	return n.data
}

// String returns a human readable string representation of the Series.
func (s *Series) String() string {

	var sb strings.Builder
	for n := s.list.head; n != nil; n = n.next {
		if n != s.head {
			sb.WriteRune('\n')
		}

		sb.WriteString(n.data.String())
	}

	return sb.String()
}
