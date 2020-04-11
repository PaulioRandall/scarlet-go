package lexeme

import (
	"fmt"
	"strconv"
)

// Token represents a grammar token within a script.
type Token struct {
	Lexeme Lexeme // Meaning
	Value  string // Representation
	Line   int    // Location within script
	Col    int    // Location within line
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
