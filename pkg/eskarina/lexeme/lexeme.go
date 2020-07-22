package lexeme

import (
	"fmt"
	"strings"

	"github.com/PaulioRandall/scarlet-go/pkg/eskarina/prop"
)

type Token interface {
	Has(prop.Prop) bool
	Is(...prop.Prop) bool
	Any(...prop.Prop) bool
}

type Snippet interface {
	At() (line, start, end int)
}

type Node interface {
	NextNode() *Lexeme
	PrevNode() *Lexeme
	ShiftUp()
	ShiftDown()
	Prepend(*Lexeme)
	Append(*Lexeme)
	Remove()
}

type Lexeme struct {
	Props []prop.Prop
	Raw   string
	Line  int
	Col   int
	Next  *Lexeme
	Prev  *Lexeme
}

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

func (lex Lexeme) NextNode() *Lexeme {
	return lex.Next
}

func (lex Lexeme) PrevNode() *Lexeme {
	return lex.Prev
}

func (lex *Lexeme) ShiftUp() {

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

func (lex *Lexeme) ShiftDown() {

	if lex.Next == nil {
		return
	}

	lex.Next.ShiftUp()
}

func (lex *Lexeme) Prepend(other *Lexeme) {

	if lex.Prev != nil {
		lex.Prev.Next = other
		other.Prev = lex.Prev
	}

	other.Next = lex
	lex.Prev = other
}

func (lex *Lexeme) Append(other *Lexeme) {

	if lex.Next != nil {
		lex.Next.Prepend(other)
		return
	}

	lex.Next = other
	other.Prev = lex
}

func (lex *Lexeme) Remove() {

	if lex.Next != nil {
		lex.Next.Prev = lex.Prev
	}

	if lex.Prev != nil {
		lex.Prev.Next = lex.Next
	}

	lex.Next, lex.Prev = nil, nil
}

func (lex Lexeme) String() string {
	return fmt.Sprintf("%d:%d %s %q",
		lex.Line,
		lex.Col,
		"["+prop.Join(",", lex.Props...)+"]",
		lex.Raw,
	)
}

func DiffPrint(left, right *Lexeme) {

	pad := func(s string, n int) string {

		if len(s) >= n {
			return s
		}

		p := n - len(s)
		return s + strings.Repeat(" ", p)
	}

	const padding = 38

	fmt.Print("\n  ")
	fmt.Print(pad("Left", padding))
	fmt.Println("right")

	for left != nil || right != nil {

		var lStr, rStr string

		if left != nil {
			lStr = pad(left.String(), padding)
			left = left.Next
		}

		if right != nil {
			rStr = pad(right.String(), padding)
			right = right.Next
		}

		if lStr != rStr {
			fmt.Print("- ")
		} else {
			fmt.Print("+ ")
		}

		fmt.Print(lStr)
		fmt.Println(rStr)
	}

	fmt.Println()
}
