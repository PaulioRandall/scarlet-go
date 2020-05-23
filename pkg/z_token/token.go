package z_token

import (
	"fmt"
	"strconv"
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
	Kind() Kind
	Lexeme
	TextPos
}

func ToString(tk Token) string {

	if v, ok := tk.(fmt.Stringer); ok {
		return v.String()
	}

	var s interface{}
	v := tk.Value()
	m := tk.Morpheme()

	switch m {
	case M_TEMPLATE, M_TERMINATOR, M_NEWLINE, M_WHITESPACE:
		s = strconv.QuoteToGraphic(v)

	case M_STRING:
		s = "`" + v + "`"

	default:
		s = v
	}

	// +1 for line index to number
	return fmt.Sprintf(`%d:%d %s %v`,
		tk.Line()+1,
		tk.Col(),
		m.String(),
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
