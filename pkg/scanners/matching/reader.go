package matching

import (
	"github.com/PaulioRandall/scarlet-go/pkg/err"
	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

var cache []pattern = patterns()

func ScanAll(s string) []Token {

	var tks []Token
	sym := &symbols{[]rune(s), 0, 0}

	for tk := readNext(sym); tk != nil; tk = readNext(sym) {
		tks = append(tks, tk)
	}

	return tks
}

func readNext(s *symbols) Token {

	if s.empty() {
		return nil
	}

	tk := readToken(s)

	if tk == nil {
		err.Panic("Unknown token", err.Pos(s.line, s.col))
	}

	return tk
}

func readToken(s *symbols) Token {

	for _, p := range cache {

		n := p.matcher(s)

		if n > 0 {
			return tokenize(s, n, p)
		}
	}

	return nil
}

func tokenize(s *symbols, terminalCount int, p pattern) Token {

	ty := p.tokenType
	l := s.line
	c := s.col
	v := s.readNonTerminal(terminalCount)

	return NewToken(ty, v, l, c)
}
