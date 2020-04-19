package statement

import (
	"github.com/PaulioRandall/scarlet-go/pkg/token"
)

type Block struct {
	Open  token.Token
	Stats []Statement
	Close token.Token
}

func (b Block) Token() token.Token {
	return b.Open
}

func (b Block) String(i int) string {

	s := indent(i) + "[Block]" + newline()
	s += indent(i+1) + "Open: " + b.Open.String() + newline()

	s += indent(i+1) + "Statements:" + newline()
	for _, st := range b.Stats {
		s += st.String(i+2) + newline()
	}

	s += indent(i+1) + "Close: " + b.Close.String()
	return s
}
