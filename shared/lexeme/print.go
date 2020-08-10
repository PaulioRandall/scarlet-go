package lexeme

import (
	"io"
	"strconv"
	"strings"
)

func Print(w io.StringWriter, head *Lexeme) error {
	return printLexemes(w, head)
}

func printLexemes(w io.StringWriter, head *Lexeme) error {

	if head == nil {
		return nil
	}

	lineMax, colMax, tokMax := findPrintPaddings(head)

	for lex := head; lex != nil; lex = lex.next {

		line := padFront(lineMax, strconv.Itoa(lex.Line))
		col := padBack(colMax+1, strconv.Itoa(lex.Col)+",")
		tok := padBack(tokMax+1, lex.Tok.String()+",")
		raw := strconv.QuoteToGraphic(lex.Raw)

		e := writeLine(w, line, ":", col, " ", tok, " ", raw)
		if e != nil {
			return e
		}
	}

	return nil
}

func findPrintPaddings(head *Lexeme) (line, col, tok int) {

	for lex := head; lex != nil; lex = lex.next {

		if lex.Line > line {
			line = lex.Line
		}

		if lex.Col > col {
			col = lex.Col
		}

		if tokLen := len(lex.Tok.String()); tokLen > tok {
			tok = tokLen
		}
	}

	line = len(strconv.Itoa(line))
	col = len(strconv.Itoa(col))
	return
}

func padFront(min int, s string) string {
	pad := strings.Repeat(" ", min-len(s))
	return pad + s
}

func padBack(min int, s string) string {
	pad := strings.Repeat(" ", min-len(s))
	return s + pad
}

func writeLine(w io.StringWriter, strs ...string) error {

	for _, s := range strs {
		_, e := w.WriteString(s)
		if e != nil {
			return e
		}
	}

	_, e := w.WriteString("\n")
	return e
}
