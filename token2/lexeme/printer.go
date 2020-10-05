package lexeme

import (
	"io"
	"strconv"
	"strings"

	"github.com/PaulioRandall/scarlet-go/token2/position"
)

// LexemeIterator provides sequential access to the ordered list of Lexemes
// that are to be printed when using this packages Print function.
type LexemeIterator interface {
	More() bool
	Next() Lexeme
	JumpToStart()
}

// Print writes the string representation of all Lexemes in 'itr', in order, to
// the writer 'w'.
func Print(w io.StringWriter, itr LexemeIterator) error {
	itr.JumpToStart()
	cw := findCellWidths(itr)
	itr.JumpToStart()
	return printLexemes(w, cw, itr)
}

type cellWidths struct {
	line int
	col  int
	tk   int
}

func findCellWidths(itr LexemeIterator) cellWidths {

	var r cellWidths
	for itr.More() {
		l := itr.Next()
		r.line = max(r.line, l.UTF8Pos.Line, l.End.Line)
		r.col = max(r.col, l.UTF8Pos.ColRune, l.End.ColRune)
		r.tk = max(r.tk, len(l.Token.String()))
	}

	r.line = len(strconv.Itoa(r.line))
	r.col = len(strconv.Itoa(r.col))
	return r
}

func printLexemes(w io.StringWriter, cw cellWidths, itr LexemeIterator) error {
	for itr.More() {
		l := itr.Next()

		// Examples:
		// `  1:2   ->   1:4   ASSIGN   ":="`
		// `100:100 -> 100:101 NEWLINE  "\n"`
		start := positionString(cw, l.UTF8Pos)
		end := positionString(cw, l.End)
		tok := padBack(cw.tk, l.Token.String())
		val := strconv.QuoteToGraphic(l.Val)

		e := writeLine(w, start, " -> ", end, " ", tok, " ", val)
		if e != nil {
			return e
		}
	}

	return nil
}

func positionString(cw cellWidths, p position.UTF8Pos) string {
	// Examples:
	// `  1:2  `
	// `100:100`
	line := padFront(cw.line, strconv.Itoa(p.Line))
	col := padBack(cw.col, strconv.Itoa(p.ColRune))
	return line + ":" + col
}

func max(first int, others ...int) int {
	r := first
	for _, o := range others {
		if o > r {
			r = o
		}
	}
	return r
}

func padFront(min int, s string) string {
	pad := strings.Repeat(" ", min-len(s))
	return pad + s
}

func padBack(min int, s string) string {
	pad := strings.Repeat(" ", min-len(s))
	return s + pad
}

func writeLine(w io.StringWriter, items ...string) error {

	const LINE_SEPARATOR = "\n"

	for _, s := range items {
		_, e := w.WriteString(s)
		if e != nil {
			return e
		}
	}

	_, e := w.WriteString(LINE_SEPARATOR)
	return e
}
