package matching

import (
	"github.com/PaulioRandall/scarlet-go/pkg/token"
)

var patternCache []pattern = patterns()

func ScanAll(s string) []token.Token {
	tks := ReadAll(s)
	return SanitiseAll(tks)
}

func ReadAll(src string) []token.Token {

	var tk token.Token
	var tks []token.Token

	s := &symbols{[]rune(src), 0, 0}

	for tk.Type != token.EOF {
		tk = readNext(s)
		tks = append(tks, tk)
	}

	return tks
}

// SanitiseAll removes redundant tokens, such as comment and whitespace, as well
// as applying formatting to values, e.g trimming off the quotes from string
// literals and templates.
func SanitiseAll(in []token.Token) (out []token.Token) {

	itr := token.NewIterator(in)
	var prev token.Token

	for prev.Type != token.EOF && !itr.Empty() {
		p := sanitiseNext(itr, prev)

		if p != (token.Token{}) {
			out = append(out, p)
			prev = p
		}
	}

	return out
}

type scanErr struct {
	msg  string
	line int
	col  int
	len  int
}

func err(s *symbols, colOffset int, msg string) error {
	return scanErr{
		line: s.line,
		col:  s.col + colOffset,
		msg:  msg,
	}
}

func (se scanErr) Error() string {
	return se.msg
}

func (se scanErr) Line() int {
	return se.line
}

func (se scanErr) Col() int {
	return se.col
}

func (se scanErr) Len() int {
	return se.len
}
