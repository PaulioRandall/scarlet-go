package z_token

import (
	"fmt"
	"strconv"
)

type Token interface {
	Morpheme() Morpheme
	Value() string
	Line() int
	Col() int
}

func ToString(tk Token) string {

	if tk == nil {
		return `NIL-TOKEN`
	}

	if v, ok := tk.(fmt.Stringer); ok {
		return v.String()
	}

	var s interface{}
	v := tk.Value()
	m := tk.Morpheme()

	switch m {
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
