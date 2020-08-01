package inst

import (
	"fmt"
	"io"
	"strconv"
	"strings"
)

type inFmt struct {
	begin int
	code  int
	gap   int // gap between printed fields
}

func PrintAll(w io.StringWriter, ins []Instruction) error {

	inf := findInFmt(ins)

	for _, in := range ins {
		e := writeLine(w, &in, inf)
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

		line, col, _ := in.Snippet.At()
		size := len(strconv.Itoa(line) + ":" + strconv.Itoa(col))

		if size > inf.begin {
			inf.begin = size
		}

		size = len(in.Code.String())
		if size > inf.code {
			inf.code = size
		}
	}

	return inf
}

func writeLine(w io.StringWriter, in *Instruction, inf inFmt) error {

	e := writeInst(w, in, inf)
	if e != nil {
		return e
	}

	_, e = w.WriteString("\n")
	return e
}

func writeInst(w io.StringWriter, in *Instruction, inf inFmt) error {

	line, col, _ := in.Snippet.At()
	e := writePos(w, line, col, inf.begin)
	if e != nil {
		return e
	}

	e = writeGap(w, inf.gap)
	if e != nil {
		return e
	}

	e = writePadStr(w, in.Code.String(), inf.code)
	if e != nil {
		return e
	}

	e = writeGap(w, inf.gap)
	if e != nil {
		return e
	}

	s := fmt.Sprintf("%v", in.Data)
	return writeStr(w, s)
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
