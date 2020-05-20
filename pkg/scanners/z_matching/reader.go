package z_matching

import (
	"github.com/PaulioRandall/scarlet-go/pkg/token"
)

func readNext(s *symbols) token.Token {

	if s.empty() {
		// TokenStream.Read requires an EOF token be returned upon an empty stream.
		return newToken(s, token.EOF)
	}

	tk := readToken(s)

	if tk == (token.Token{}) {
		panic(err(s, 0, "Unknown token"))
	}

	if tk.Type == token.EOF {
		s.drain()
	}

	return tk
}

func readToken(s *symbols) (_ token.Token) {

	for _, p := range patternCache {

		n := p.matcher(s)

		if n > 0 {
			return tokenize(s, n, p.tokenType)
		}
	}

	return
}

func tokenize(s *symbols, numOfTerminals int, t token.TokenType) token.Token {
	tk := newToken(s, t)
	tk.Value = s.readNonTerminal(numOfTerminals)
	return tk
}

func newToken(s *symbols, t token.TokenType) token.Token {
	return token.Token{
		Type: t,
		Line: s.line,
		Col:  s.col,
	}
}
