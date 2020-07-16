package inst

import (
	"fmt"
	"io"
	"strconv"
	"strings"
)

type inFmt struct {
	begin int
	end   int
	code  int
	gap   int // gap between printed fields
}

func PrintAll(w io.StringWriter, ins []Instruction) error {

	inf := findInFmt(ins)

	for _, in := range ins {
		e := writeLine(w, in, inf)
		if e != nil {
			return e
		}
	}

	return nil
}

func findInFmt(ins []Instruction) inFmt {

	inf := inFmt{
		gap: 4,
	}

	for _, in := range ins {

		line, col := in.Begin()
		size := len(strconv.Itoa(line) + strconv.Itoa(col))
		if size > inf.begin {
			inf.begin = size
		}

		line, col = in.End()
		size = len(strconv.Itoa(line) + strconv.Itoa(col))
		if size > inf.end {
			inf.end = size
		}

		size = len(in.Code().String())
		if size > inf.code {
			inf.code = size
		}
	}

	return inf
}

func writeLine(w io.StringWriter, in Instruction, inf inFmt) error {

	e := writeToken(w, in, inf)
	if e != nil {
		return e
	}

	_, e = w.WriteString("\n")
	return e
}

func writeToken(w io.StringWriter, in Instruction, inf inFmt) error {

	e := writeSnippet(w, in, inf)
	if e != nil {
		return e
	}

	e = writeGap(w, inf.gap)
	if e != nil {
		return e
	}

	e = writePadStr(w, in.Code().String(), inf.code)
	if e != nil {
		return e
	}

	e = writeGap(w, inf.gap)
	if e != nil {
		return e
	}

	s := fmt.Sprintf("%v", in.Data())
	return writeStr(w, s)
}

func writeSnippet(w io.StringWriter, snip Snippet, inf inFmt) error {

	line, col := snip.Begin()
	e := writePos(w, line, col, inf.begin)
	if e != nil {
		return e
	}

	e = writeGap(w, inf.gap)
	if e != nil {
		return e
	}

	line, col = snip.End()
	return writePos(w, line, col, inf.end)
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
