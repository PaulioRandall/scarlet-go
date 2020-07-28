package lexeme

import (
	"io"
	"strconv"
	"strings"
)

func CopyAll(head *Lexeme) *Lexeme {

	var que Queue = &Container{}

	for lex := head; lex != nil; lex = lex.Next {
		que.Put(Copy(lex))
	}

	return que.Head()
}

func Copy(lex *Lexeme) *Lexeme {
	return &Lexeme{
		Raw: lex.Raw,
	}
}

func Split(head *Lexeme, delim Token) []*Lexeme {

	r := []*Lexeme{}

	if head == nil {
		return r
	}

	findNextDelim := func(head *Lexeme) *Lexeme {
		for next := head; next != nil; next = next.Next {
			if next.Tok == delim {
				return next
			}
		}
		return nil
	}

	for head != nil {
		r = append(r, head)
		delim := findNextDelim(head)
		head = delim.Next
		delim.SplitBelow()
	}

	return r
}

type lexFmt struct {
	begin int
	end   int
	tok   int
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

		size = len(lex.Tok.String())
		if size > lxf.tok {
			lxf.tok = size
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

	e = writePadStr(w, lex.Tok.String(), lxf.tok)
	if e != nil {
		return e
	}

	e = writeGap(w, lxf.gap)
	if e != nil {
		return e
	}

	if lex.Tok == NEWLINE {
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
