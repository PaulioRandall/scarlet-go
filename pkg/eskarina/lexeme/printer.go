package lexeme

import (
	"io"
	"strconv"
	"strings"

	"github.com/PaulioRandall/scarlet-go/pkg/eskarina/prop"
)

type lexFmt struct {
	begin int
	end   int
	props int
	gap   int // gap between printed fields
}

func PrintAll(w io.StringWriter, first *Lexeme) error {

	lxf := findLexFmt(first)

	for lex := first; lex != nil; lex = lex.Next {
		e := writeLine(w, lex, lxf)
		if e != nil {
			return e
		}
	}

	return nil
}

func findLexFmt(first *Lexeme) lexFmt {

	lxf := lexFmt{
		gap: 4,
	}

	for lex := first; lex != nil; lex = lex.Next {

		line, col, colEnd := lex.At()

		lineStr := strconv.Itoa(line)
		colStr := strconv.Itoa(col)
		colEndStr := strconv.Itoa(colEnd)

		size := len(lineStr + ":" + colStr)
		if size > lxf.begin {
			lxf.begin = size
		}

		size = len(lineStr + ":" + colEndStr)
		if size > lxf.end {
			lxf.end = size
		}

		size = len([]rune(prop.Join(", ", lex.Props...)))
		if size > lxf.props {
			lxf.props = size
		}
	}

	return lxf
}

func writeLine(w io.StringWriter, lex *Lexeme, lxf lexFmt) error {

	e := writeToken(w, lex, lxf)
	if e != nil {
		return e
	}

	_, e = w.WriteString("\n")
	return e
}

func writeToken(w io.StringWriter, lex *Lexeme, lxf lexFmt) error {

	e := writeSnippet(w, lex, lxf)
	if e != nil {
		return e
	}

	e = writeGap(w, lxf.gap)
	if e != nil {
		return e
	}

	e = writePadStr(w, prop.Join(", ", lex.Props...), lxf.props)
	if e != nil {
		return e
	}

	e = writeGap(w, lxf.gap)
	if e != nil {
		return e
	}

	if lex.Is(prop.PR_NEWLINE) {
		s := strconv.QuoteToGraphic(lex.Raw)
		return writeStr(w, s[1:len(s)-1])
	}

	return writeStr(w, lex.Raw)
}

func writeSnippet(w io.StringWriter, snip Snippet, lxf lexFmt) error {

	line, col, colEnd := snip.At()

	e := writePos(w, line, col, lxf.begin)
	if e != nil {
		return e
	}

	e = writeGap(w, lxf.gap)
	if e != nil {
		return e
	}

	return writePos(w, line, colEnd, lxf.end)
}

func writePos(w io.StringWriter, line, col, minLen int) error {
	s := strconv.Itoa(line) + ":" + strconv.Itoa(col)
	s = pad(minLen+1, s)
	return writeStr(w, s)
}

func pad(min int, s string) string {
	pad := strings.Repeat(" ", min-len(s))
	return s + pad
}

func writeGap(w io.StringWriter, gap int) error {
	return writePadStr(w, "", gap)
}

func writePadStr(w io.StringWriter, s string, maxLen int) error {
	s = pad(maxLen, s)
	return writeStr(w, s)
}

func writeStr(w io.StringWriter, s string) error {
	_, e := w.WriteString(s)
	return e
}
