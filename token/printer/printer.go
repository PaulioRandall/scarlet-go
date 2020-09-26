package printer

import (
	"io"
	"strconv"
	"strings"

	"github.com/PaulioRandall/scarlet-go/token/lexeme"
)

type colWidths struct {
	line int
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
	}

	r.line = len(strconv.Itoa(r.line))
	return r
}

func printLexemes(w io.StringWriter, cw colWidths, lexs []lexeme.Lexeme) error {

	for _, l := range lexs {

		line := padFront(cw.line, strconv.Itoa(l.Line()))
		//col := padBack(colMax+1, strconv.Itoa(lex.Col)+",")
		//tok := padBack(tokMax+1, lex.Tok.String()+",")
		//raw := strconv.QuoteToGraphic(lex.Raw)

		e := writeLine(w, line) //, ":", col, " ", tok, " ", raw)
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
