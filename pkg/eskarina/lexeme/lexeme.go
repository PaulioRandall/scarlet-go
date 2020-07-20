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

func (lex Lexeme) String() string {
	return fmt.Sprintf("%d:%d %s %s",
		lex.Line,
		lex.Col,
		"["+prop.Join(",", lex.Props...)+"]",
		lex.Raw,
	)
}
