package matching

import (
	"github.com/PaulioRandall/scarlet-go/lexeme"

	"github.com/PaulioRandall/scarlet-go/streams/symbol"
)

// ReadAll parses all tokens from ss into an array.
func ReadAll(ss symbol.SymbolStream) []lexeme.Token {

	var tk lexeme.Token
	var tokens []lexeme.Token

	for tk.Lexeme != lexeme.LEXEME_EOF {
		tk = Read(ss)
		tokens = append(tokens, tk)
	}

	return tokens
}

func Read(ss symbol.SymbolStream) lexeme.Token {

	if ss.Empty() {
		// TokenStream.Read requires an EOF token be returned upon an empty stream.
		return lexeme.Token{
			Lexeme: lexeme.LEXEME_EOF,
			Line:   ss.LineIndex(),
			Col:    ss.ColIndex(),
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
func readToken(ss symbol.SymbolStream) (_ lexeme.Token) {

	matchers := nonTerminals()
	size := len(matchers)

	for i := 0; i < size; i++ {

		nonTerminal := matchers[i]
		n := nonTerminal.matcher(ss)

		if n > 0 {
			return tokenize(ss, n, nonTerminal.lexeme)
		}
	}

	return
}

// tokenize creates a new token from the next non-terminal. It reads off n
// symbols from ss ready for scanning the next token.
func tokenize(ss symbol.SymbolStream, n int, l lexeme.Lexeme) lexeme.Token {

	tk := lexeme.Token{
		Lexeme: l,
		Line:   ss.LineIndex(),
		Col:    ss.ColIndex(),
	}

	tk.Value = ss.ReadNonTerminal(n)
	return tk
}
