package matching

import (
	"github.com/PaulioRandall/scarlet-go/pkg/token"
)

// ReadAll parses all tokens from s into an array.
func ReadAll(s string) []token.Token {

	var tk token.Token
	var tokens []token.Token
	ss := &symbolStream{
		runes: []rune(s),
	}

	for tk.Type != token.EOF {
		tk = read(ss)
		tokens = append(tokens, tk)
	}

	return tokens
}

func read(ss *symbolStream) token.Token {

	if ss.empty() {
		// TokenStream.Read requires an EOF token be returned upon an empty stream.
		return token.Token{
			Type: token.EOF,
			Line: ss.lineIndex(),
			Col:  ss.colIndex(),
		}
	}

	tk := readToken(ss)

	if tk == (token.Token{}) {
		panic(newErr(ss, 0, "Could not identify next token"))
	}

	if tk.Type == token.EOF {
		ss.drain()
	}

	return tk
}

// readToken attempts to match one of the non-terminal patterns to the next
// set of terminals in the script. If found, the terminals are removed and used
// to create a token.
func readToken(ss *symbolStream) (_ token.Token) {

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

// tokenize creates a new token from the next non-terminal. It reads off n
// symbols from ss ready for scanning the next token.
func tokenize(ss *symbolStream, n int, l token.TokenType) token.Token {

	tk := token.Token{
		Type: l,
		Line: ss.lineIndex(),
		Col:  ss.colIndex(),
	}

	tk.Value = ss.readNonTerminal(n)
	return tk
}
