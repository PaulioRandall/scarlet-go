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

// Precedence returns the priorty of the token within an expression.
func (tk Token) Precedence() int {
	switch tk.Type {
	case MULTIPLY, DIVIDE, REMAINDER:
		return 6 // Multiplicative

	case ADD, SUBTRACT:
		return 5 // Additive

	case LESS_THAN, LESS_THAN_OR_EQUAL, MORE_THAN, MORE_THAN_OR_EQUAL:
		return 4 // Relational

	case EQUAL, NOT_EQUAL:
		return 3 // Equalitive

	case AND:
		return 2

	case OR:
		return 1
	}

	return 0
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
