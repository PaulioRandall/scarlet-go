package matching

import (
	"github.com/PaulioRandall/scarlet-go/pkg/token"
)

func ScanAll(src string) []token.Token {

	var tk token.Token
	var tks []token.Token

	s := &symbols{[]rune(src), 0, 0}

	for tk.Type != token.EOF {
		tk = scanNext(s)
		tks = append(tks, tk)
	}

	return tks
}

func scanNext(s *symbols) token.Token {

	if s.empty() {
		// TokenStream.Read requires an EOF token be returned upon an empty stream.
		return token.Token{
			Type: token.EOF,
			Line: s.lineIndex(),
			Col:  s.colIndex(),
		}
	}

	tk := scanToken(s)

	if tk == (token.Token{}) {
		panic(newErr(s, 0, "Could not identify next token"))
	}

	if tk.Type == token.EOF {
		s.drain()
	}

	return tk
}

func scanToken(s *symbols) (_ token.Token) {

	ps := patterns()
	size := len(ps)

	for i := 0; i < size; i++ {

		p := ps[i]
		n := p.matcher(s)

		if n > 0 {
			return tokenize(s, n, p.tokenType)
		}
	}

	return
}

func tokenize(s *symbols, n int, l token.TokenType) token.Token {

	tk := token.Token{
		Type: l,
		Line: s.lineIndex(),
		Col:  s.colIndex(),
	}

	tk.Value = s.readNonTerminal(n)
	return tk
}
