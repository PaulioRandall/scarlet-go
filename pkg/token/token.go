package token

import (
	"fmt"
	"strconv"
)

type Token struct {
	Type  TokenType // Meaning
	Value string    // Representation
	Line  int
	Col   int
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

func (tk Token) String() string {

	var v interface{}

	switch tk.Type {
	case TEMPLATE, TERMINATOR, NEWLINE, WHITESPACE:
		v = strconv.QuoteToGraphic(tk.Value)

	case STRING:
		v = "`" + tk.Value + "`"

	default:
		v = tk.Value
	}

	// +1 for line index to number
	return fmt.Sprintf(`%d:%d %s %v`, tk.Line+1, tk.Col, tk.Type, v)
}

func PrettyPrint(tks []Token) {
	for _, tk := range tks {
		switch k := tk.Type; k {
		case EOF:
			fmt.Println(k)
			fmt.Println()
		default:
			fmt.Print(k + " ")
		}
	}
}
