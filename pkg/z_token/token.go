package z_token

import (
	"fmt"
	"strconv"

	. "github.com/PaulioRandall/scarlet-go/pkg/morpheme"
)

type Lexeme interface {
	Morpheme() Morpheme
	Value() string
}

type TextPos interface {
	Line() int
	Col() int
}

type Token interface {
	Lexeme
	TextPos
}

func ToString(tk Token) string {

	if v, ok := tk.(fmt.Stringer); ok {
		return v.String()
	}

	var s interface{}
	v := tk.Value()

	switch tk.Morpheme() {
	case TEMPLATE, TERMINATOR, NEWLINE, WHITESPACE:
		s = strconv.QuoteToGraphic(v)

	case STRING:
		s = "`" + v + "`"

	default:
		s = v
	}

	// +1 for line index to number
	return fmt.Sprintf(`%d:%d %s %v`,
		tk.Line()+1,
		tk.Col(),
		tk.Morpheme().String(),
		s,
	)
}

func PrettyPrint(tks []Token) {

	for _, tk := range tks {
		s := tk.Morpheme().String()
		fmt.Print(s + " ")
	}

	fmt.Println()
}
