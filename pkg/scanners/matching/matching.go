package matching

import (
	"github.com/PaulioRandall/scarlet-go/pkg/token"
)

var patternCache []pattern = patterns()

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
		return newToken(s, token.EOF)
	}

	tk := scanToken(s)

	if tk == (token.Token{}) {
		panic(err(s, 0, "Unknown token"))
	}

	if tk.Type == token.EOF {
		s.drain()
	}

	return tk
}

func scanToken(s *symbols) (_ token.Token) {

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
