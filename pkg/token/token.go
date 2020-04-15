package token

import (
	"fmt"
	"strconv"
)

// Token represents a grammar token within a script.
type Token struct {
	Type  TokenType // Meaning
	Value string    // Representation
	Line  int       // Location within script
	Col   int       // Location within line
}

// String returns a string representation of the token.
func (tk Token) String() string {

	var v interface{}

	if tk.Type == TEMPLATE {
		v = strconv.QuoteToGraphic(tk.Value)

	} else if tk.Type == STRING {
		v = "`" + tk.Value + "`"

	} else {
		v = tk.Value
	}

	// +1 for line index to number
	return fmt.Sprintf(`%d:%d %s %v`, tk.Line+1, tk.Col, tk.Type, v)
}

// PrintTokens pretty prints all tokens in tks.
func PrintTokens(tks []Token) {
	for _, tk := range tks {
		switch k := tk.Type; k {
		case EOF:
			println(k)
			println()
		default:
			print(k + " ")
		}
	}
}
