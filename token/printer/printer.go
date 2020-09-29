package printer

import (
	"io"
	"strconv"
	"strings"

	"github.com/PaulioRandall/scarlet-go/token/lexeme"
)

type colWidths struct {
	line int
	col  int
	typ  int
}

func Print(w io.StringWriter, lexs []lexeme.Lexeme) error {
	cw := findColWidths(lexs)
	return printLexemes(w, cw, lexs)
}

func findColWidths(lexs []lexeme.Lexeme) colWidths {

	var r colWidths

	for _, l := range lexs {
		if l.Line() > r.line {
			r.line = l.Line()
		}

		if l.Col() > r.col {
			r.col = l.Col()
		}

		if len(l.Type().String()) > r.typ {
			r.typ = len(l.Type().String())
		}
	}

	r.line = len(strconv.Itoa(r.line))
	r.col = len(strconv.Itoa(r.col))
	return r
}

func printLexemes(w io.StringWriter, cw colWidths, lexs []lexeme.Lexeme) error {

	for _, l := range lexs {

		line := padFront(cw.line, strconv.Itoa(l.Line()))
		col := padBack(cw.col+1, strconv.Itoa(l.Col())+",")
		tok := padBack(cw.typ+1, l.Type().String()+",")
		val := strconv.QuoteToGraphic(l.Val())

		e := writeLine(w, line, ":", col, " ", tok, " ", val)
		if e != nil {
			return e
		}
	}

	return nil
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
