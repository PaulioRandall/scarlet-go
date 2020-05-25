package matching

import (
	"github.com/PaulioRandall/scarlet-go/pkg/err"
	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

var cache []pattern = patterns()

func ScanAll(s string) []Token {

	in := ReadAll(s)
	out := make([]Token, len(in))

	for i := range in {
		out[i] = in[i]
	}

	return out
}

func ReadAll(s string) []tok {

	var tks []tok
	sym := &symbols{[]rune(s), 0, 0}

	for tk, ok := readNext(sym); ok; tk, ok = readNext(sym) {
		tks = append(tks, tk)
	}

	return tks
}

func readNext(s *symbols) (tok, bool) {

	if s.empty() {
		return tok{}, false
	}

	tk := readToken(s)

	if tk == (tok{}) {
		err.Panic(
			"Unknown token",
			err.Pos(s.line, s.col),
		)
	}

	return tk, true
}

func readToken(s *symbols) tok {

	for _, p := range cache {

		n := p.matcher(s)

		if n > 0 {
			return tokenize(s, n, p)
		}
	}

	return tok{}
}

func tokenize(s *symbols, terminalCount int, p pattern) tok {

	tk := tok{
		m: p.morpheme,
		l: s.line,
		c: s.col,
	}

	tk.v = s.readNonTerminal(terminalCount)
	return tk
}
