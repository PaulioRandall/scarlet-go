package lexeme

import (
	"fmt"
	"strconv"
)

// Token represents a grammer token within a source file.
type Token struct {
	Lexeme Lexeme
	Value  string
	Line   int
	Col    int
}

// KindsToStrings converts the kind slice to a string slice.
func KindsToStrings(lexs []Lexeme) (strs []string) {

	for _, lex := range lexs {
		strs = append(strs, string(lex))
	}

	return
}

// New creates a new token.
func New(lex Lexeme, v string, l, c int) Token {
	return Token{
		Lexeme: lex,
		Value:  v,
		Line:   l,
		Col:    c,
	}
}

// OfKind creates a new token with the specified kind.
func OfKind(lex Lexeme) Token {
	return Token{
		Lexeme: lex,
	}
}

// OfValue creates a new token with the specified kind and value.
func OfValue(lex Lexeme, v string) Token {
	return Token{
		Lexeme: lex,
		Value:  v,
	}
}

// String returns a string representation of the token.
func (tk Token) String() string {

	var v interface{}

	if tk.Lexeme == LEXEME_TEMPLATE {
		v = strconv.QuoteToGraphic(tk.Value)
	} else if tk.Lexeme == LEXEME_STRING {
		v = "`" + tk.Value + "`"
	} else {
		v = tk.Value
	}

	// +1 for line index to number
	return fmt.Sprintf(`%d:%d %s %v`, tk.Line+1, tk.Col, tk.Lexeme, v)
}
