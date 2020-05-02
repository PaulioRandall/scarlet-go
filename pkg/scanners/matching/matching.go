package matching

import (
	"github.com/PaulioRandall/scarlet-go/pkg/token"
)

func ScanAll(src string) []token.Token {

	var tk token.Token
	var tks []token.Token

	s := &symbolStream{[]rune(src), 0, 0}

	for tk.Type != token.EOF {
		tk = scanNext(s)
		tks = append(tks, tk)
	}

	return tks
}

func scanNext(ss *symbolStream) token.Token {

	if ss.empty() {
		// TokenStream.Read requires an EOF token be returned upon an empty stream.
		return token.Token{
			Type: token.EOF,
			Line: ss.lineIndex(),
			Col:  ss.colIndex(),
		}
	}

	tk := scanToken(ss)

	if tk == (token.Token{}) {
		panic(newErr(ss, 0, "Could not identify next token"))
	}

	if tk.Type == token.EOF {
		ss.drain()
	}

	return tk
}

func scanToken(ss *symbolStream) (_ token.Token) {

	ps := patterns()
	size := len(ps)

	for i := 0; i < size; i++ {

		p := ps[i]
		n := p.matcher(ss)

		if n > 0 {
			return tokenize(ss, n, p.tokenType)
		}
	}

	return
}

func tokenize(ss *symbolStream, n int, l token.TokenType) token.Token {

	tk := token.Token{
		Type: l,
		Line: ss.lineIndex(),
		Col:  ss.colIndex(),
	}

	tk.Value = ss.readNonTerminal(n)
	return tk
}
