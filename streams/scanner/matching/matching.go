package matching

import (
	"github.com/PaulioRandall/scarlet-go/lexeme"

	"github.com/PaulioRandall/scarlet-go/streams/symbol"
)

// ReadAll parses all tokens from ts into an array.
func ReadAll(ts *symbol.TerminalStream) []lexeme.Token {

	var tk lexeme.Token
	var tokens []lexeme.Token

	for tk.Lexeme != lexeme.LEXEME_EOF {
		tk = Read(ts)
		tokens = append(tokens, tk)
	}

	return tokens
}

func Read(ts *symbol.TerminalStream) lexeme.Token {

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
func readToken(ts *symbol.TerminalStream) (_ lexeme.Token) {

	matchers := nonTerminals()
	size := len(matchers)

	for i := 0; i < size; i++ {

		nonTerminal := matchers[i]
		n := nonTerminal.matcher(ts)

		if n > 0 {
			return tokenize(ts, n, nonTerminal.lexeme)
		}
	}

	return
}

// tokenize creates a new token from the next non-terminal. It reads off n
// symbols from ts ready for scanning the next token.
func tokenize(ts *symbol.TerminalStream, n int, l lexeme.Lexeme) lexeme.Token {

	tk := lexeme.Token{
		Lexeme: l,
		Line:   ts.LineIndex(),
		Col:    ts.ColIndex(),
	}

	tk.Value = ts.ReadNonTerminal(n)
	return tk
}
