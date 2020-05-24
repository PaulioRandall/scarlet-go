package z_matching

import (
	. "github.com/PaulioRandall/scarlet-go/pkg/z_token"
)

var cache []pattern = patterns()

func ReadAll(src string) []tok {

	var tk tok
	var tks []tok

	s := &symbols{[]rune(src), 0, 0}

	for tk.Kind() != K_EOF {
		tk = readNext(s)
		tks = append(tks, tk)
	}

	return tks
}

func readNext(s *symbols) tok {

	if s.empty() {
		return tok{
			k: K_EOF,
			m: M_EOF,
			l: s.line,
			c: s.col,
		}
	}

	tk := readToken(s)

	if tk == (tok{}) {
		panic(err(s, 0, "Unknown token"))
	}

	if tk.k == K_EOF {
		s.drain()
	}

	return tk
}

func readToken(s *symbols) (_ tok) {

	for _, p := range cache {

		n := p.matcher(s)

		if n > 0 {
			return tokenize(s, n, p)
		}
	}

	return
}

func tokenize(s *symbols, terminalCount int, p pattern) tok {

	tk := tok{
		k: p.kind,
		m: p.morpheme,
		l: s.line,
		c: s.col,
	}

	tk.v = s.readNonTerminal(terminalCount)
	return tk
}
