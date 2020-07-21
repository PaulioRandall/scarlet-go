package lexeme

import (
	"fmt"

	"github.com/PaulioRandall/scarlet-go/pkg/eskarina/prop"
)

type Lexeme struct {
	Props []prop.Prop
	Raw   string
	Line  int
	Col   int
	Next  *Lexeme
	Prev  *Lexeme
}

//type Token interface {
//	Has(Prop) bool
//	Is(...Prop) bool
//	Any(...Prop) bool
//}

//type Snippet interface {
//	At() (line, start, end, int)
//}

//type LinkNode interface {
//	Promote()
//	Demote()
//	Prepend(*lexeme.Lexeme)
//	Append(*lexeme.Lexeme)
//	Remove()
//}

func (lex Lexeme) Has(o prop.Prop) bool {

	for _, p := range lex.Props {
		if p == o {
			return true
		}
	}

	return false
}

func (lex Lexeme) Is(others ...prop.Prop) bool {

	for _, o := range others {
		if !lex.Has(o) {
			return false
		}
	}

	return true
}

func (lex Lexeme) Any(others ...prop.Prop) bool {

	for _, o := range others {
		if lex.Has(o) {
			return true
		}
	}

	return false
}

func (lex Lexeme) At() (line, start, end int) {
	return lex.Line, lex.Col, len(lex.Raw)
}

func (lex *Lexeme) Promote() {

	if lex.Prev == nil {
		return
	}

	prev := lex.Prev

	if prev.Prev != nil {
		prev.Prev.Next = lex
	}

	if lex.Next != nil {
		lex.Next.Prev = prev
	}

	lex.Prev, prev.Next = prev.Prev, lex.Next
	lex.Next, prev.Prev = prev, lex
}

func (lex *Lexeme) Demote() {

	if lex.Next == nil {
		return
	}

	lex.Next.Promote()
}

func (lex *Lexeme) Prepend(new *Lexeme) {

	if lex.Prev != nil {
		lex.Prev.Next = new
	}

	new.Next = lex
	new.Prev = lex.Prev
	lex.Prev = new
}

func (lex *Lexeme) Append(new *Lexeme) {

	if lex.Next != nil {
		lex.Next.Prepend(new)
		return
	}

	lex.Next = new
	new.Prev = lex
}

func (lex *Lexeme) Remove() {

	if lex.Next != nil {
		lex.Next.Prev = lex.Prev
	}

	if lex.Prev != nil {
		lex.Prev.Next = lex.Next
	}
}

func (lex Lexeme) String() string {
	return fmt.Sprintf("%d:%d %s %q",
		lex.Line,
		lex.Col,
		"["+prop.Join(",", lex.Props...)+"]",
		lex.Raw,
	)
}
