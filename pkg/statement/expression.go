package statement

import (
	"github.com/PaulioRandall/scarlet-go/pkg/token"
)

type Expression interface {
	Token() token.Token

	String(indent int) string
}

type Identifier token.Token

func (id Identifier) Token() token.Token {
	return token.Token(id)
}

func (id Identifier) String(i int) string {

	var s str

	return s.indent(i).
		append("[Identifier] ").
		append(id.Token().String()).
		String()
}

type Value token.Token

func (v Value) Token() token.Token {
	return token.Token(v)
}

func (v Value) String(i int) string {

	var s str

	return s.indent(i).
		append("[Value] ").
		append(v.Token().String()).
		String()
}
