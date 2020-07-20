package token

import (
	"fmt"

	"github.com/PaulioRandall/scarlet-go/pkg/eskarina/props"
)

type Token struct {
	Props []props.Prop
	Raw   string
	Line  int
	Col   int
	Next  *Token
	Prev  *Token
}

func (tk Token) Has(o props.Prop) bool {

	for _, p := range tk.Props {
		if p == o {
			return true
		}
	}

	return false
}

func (tk Token) Is(others ...props.Prop) bool {

	for _, o := range others {
		if !tk.Has(o) {
			return false
		}
	}

	return true
}

func (tk Token) Any(others ...props.Prop) bool {

	for _, o := range others {
		if tk.Has(o) {
			return true
		}
	}

	return false
}

func (tk Token) String() string {
	return fmt.Sprintf("%d:%d %s %s",
		tk.Line,
		tk.Col,
		"["+props.Join(",", tk.Props...)+"]",
		tk.Raw,
	)
}
