package program

import (
	"io"
	"os"
	"strconv"
	"strings"

	. "github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/prop"
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/token"
)

func writeTokenPhaseFile(filename string, tks []token.Token) error {

	f, e := os.Create(filename)
	if e != nil {
		return e
	}

	defer f.Close()

	tkf := findTkFmt(tks)

	for _, tk := range tks {
		e = writeLine(f, tk, tkf)
		if e != nil {
			return e
		}
	}

	return nil
}

type tkFmt struct {
	begin int
	end   int
	props int
	gap   int // gap between printed fields
}

func findTkFmt(tks []token.Token) tkFmt {

	tkf := tkFmt{
		gap: 4,
	}

	for _, tk := range tks {

		line, col := tk.Begin()
		size := len(strconv.Itoa(line) + strconv.Itoa(col))
		if size > tkf.begin {
			tkf.begin = size
		}

		line, col = tk.End()
		size = len(strconv.Itoa(line) + strconv.Itoa(col))
		if size > tkf.end {
			tkf.end = size
		}

		size = len([]rune(JoinProps(", ", tk.Props()...)))
		if size > tkf.props {
			tkf.props = size
		}
	}

	return tkf
}

func writeLine(w io.StringWriter, tk token.Token, tkf tkFmt) error {

	e := writeToken(w, tk, tkf)
	if e != nil {
		return e
	}

	_, e = w.WriteString("\n")
	return e
}

func writeToken(w io.StringWriter, tk token.Token, tkf tkFmt) error {

	e := writeSnippet(w, tk, tkf)
	if e != nil {
		return e
	}

	e = writeGap(w, tkf.gap)
	if e != nil {
		return e
	}

	e = writePadStr(w, JoinProps(", ", tk.Props()...), tkf.props)
	if e != nil {
		return e
	}

	e = writeGap(w, tkf.gap)
	if e != nil {
		return e
	}

	if tk.Is(PR_NEWLINE) {
		s := strconv.QuoteToGraphic(tk.Raw())
		return writeStr(w, s[1:len(s)-1])
	}

	return writeStr(w, tk.Raw())
}

func writeSnippet(w io.StringWriter, snip token.Snippet, tkf tkFmt) error {

	line, col := snip.Begin()
	e := writePos(w, line, col, tkf.begin)
	if e != nil {
		return e
	}

	e = writeGap(w, tkf.gap)
	if e != nil {
		return e
	}

	line, col = snip.End()
	return writePos(w, line, col, tkf.end)
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
