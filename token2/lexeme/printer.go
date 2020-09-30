package lexeme

import (
	"io"
	"strconv"
	"strings"

	"github.com/PaulioRandall/scarlet-go/token2/token"
)

func Print(w io.StringWriter, lexs []Lexeme) error {
	cw := findCellWidths(lexs)
	return printLexemes(w, cw, lexs)
}

type cellWidths struct {
	line int
	col  int
	tk   int
}

func findCellWidths(lexs []Lexeme) cellWidths {

	var r cellWidths
	for _, l := range lexs {
		r.line = max(r.line, l.Position.LineIdx, l.End.LineIdx)
		r.col = max(r.col, l.Position.ColRune, l.End.ColRune)
		r.tk = max(r.tk, len(l.Token.String()))
	}

	r.line = len(strconv.Itoa(r.line))
	r.col = len(strconv.Itoa(r.col))
	return r
}

func printLexemes(w io.StringWriter, cw cellWidths, lexs []Lexeme) error {
	for _, l := range lexs {

		// Examples:
		// `  1:2   ->   1:4   ASSIGN   ":="`
		// `100:100 -> 100:101 NEWLINE  "\n"`
		start := positionString(cw, l.Position)
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

func positionString(cw cellWidths, p token.Position) string {
	// Examples:
	// `  1:2  `
	// `100:100`
	line := padFront(cw.line, strconv.Itoa(p.LineIdx))
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
