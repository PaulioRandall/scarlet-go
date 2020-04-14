package matching

import (
	"github.com/PaulioRandall/scarlet-go/lexeme"

	"github.com/PaulioRandall/scarlet-go/streams/terminal"
)

// ReadAll parses all tokens from ts into an array.
func ReadAll(ts *terminal.TerminalStream) []lexeme.Token {

	var tk lexeme.Token
	var tokens []lexeme.Token

	for tk.Lexeme != lexeme.LEXEME_EOF {
		tk = Read(ts)
		tokens = append(tokens, tk)
	}

	return tokens
}

func Read(ts *terminal.TerminalStream) lexeme.Token {

	if ts.Empty() {
		// TokenStream.Read requires an EOF token be returned upon an empty stream.
		return lexeme.Token{
			Lexeme: lexeme.LEXEME_EOF,
			Line:   ts.LineIndex(),
			Col:    ts.ColIndex(),
		}
	}

	tk := readToken(ts)

	if tk == (lexeme.Token{}) {
		panic(newErr(ts, 0, "Could not identify next token"))
	}

	return tk
}

// readToken attempts to match one of the non-terminal patterns to the next
// set of terminals in the script. If found, the terminals are removed and used
// to create a token.
func readToken(ts *terminal.TerminalStream) (_ lexeme.Token) {

	ps := patterns()
	size := len(ps)

	for i := 0; i < size; i++ {

		p := ps[i]
		n := p.matcher(ts)

		if n > 0 {
			return tokenize(ts, n, p.lexeme)
		}
	}

	return
}

// tokenize creates a new token from the next non-terminal. It reads off n
// symbols from ts ready for scanning the next token.
func tokenize(ts *terminal.TerminalStream, n int, l lexeme.Lexeme) lexeme.Token {

	tk := lexeme.Token{
		Lexeme: l,
		Line:   ts.LineIndex(),
		Col:    ts.ColIndex(),
	}

	tk.Value = ts.ReadNonTerminal(n)
	return tk
}
