package matching

import (
	"github.com/PaulioRandall/scarlet-go/pkg/lexeme"
)

// ReadAll parses all tokens from s into an array.
func ReadAll(s string) []lexeme.Token {

	var tk lexeme.Token
	var tokens []lexeme.Token
	ss := &symbolStream{
		runes: []rune(s),
	}

	for tk.Lexeme != lexeme.LEXEME_EOF {
		tk = read(ss)
		tokens = append(tokens, tk)
	}

	return tokens
}

func read(ss *symbolStream) lexeme.Token {

	if ss.empty() {
		// TokenStream.Read requires an EOF token be returned upon an empty stream.
		return lexeme.Token{
			Lexeme: lexeme.LEXEME_EOF,
			Line:   ss.lineIndex(),
			Col:    ss.colIndex(),
		}
	}

	tk := readToken(ss)

	if tk == (lexeme.Token{}) {
		panic(newErr(ss, 0, "Could not identify next token"))
	}

	return tk
}

// readToken attempts to match one of the non-terminal patterns to the next
// set of terminals in the script. If found, the terminals are removed and used
// to create a token.
func readToken(ss *symbolStream) (_ lexeme.Token) {

	ps := patterns()
	size := len(ps)

	for i := 0; i < size; i++ {

		p := ps[i]
		n := p.matcher(ss)

		if n > 0 {
			return tokenize(ss, n, p.lexeme)
		}
	}

	return
}

// tokenize creates a new token from the next non-terminal. It reads off n
// symbols from ss ready for scanning the next token.
func tokenize(ss *symbolStream, n int, l lexeme.Lexeme) lexeme.Token {

	tk := lexeme.Token{
		Lexeme: l,
		Line:   ss.lineIndex(),
		Col:    ss.colIndex(),
	}

	tk.Value = ss.readNonTerminal(n)
	return tk
}
