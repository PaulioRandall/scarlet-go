package token

import (
	"github.com/PaulioRandall/scarlet-go/where"
)

// Token represents a grammer token within a source file.
type Token interface {

	// Kind returns the type of the token.
	Kind() Kind

	// Value returns the string representing the token in source.
	Value() string

	// Where returns where the token is located within the source.
	Where() where.Where
}

// tokenImpl is a simple implementation of the Token interface.
type tokenImpl struct {
	k Kind
	v string
	w where.Where
}

// Newish creates a new token.
func Newish(k Kind, v string, w where.Where) Token {
	return tokenImpl{
		k: k,
		v: v,
		w: w,
	}
}

// New creates a new token.
func New(k Kind, v string, line, start, end int) Token {
	return tokenImpl{
		k: k,
		v: v,
		w: where.New(line, start, end),
	}
}

// Kind satisfies the Token interface.
func (t tokenImpl) Kind() Kind {
	return t.k
}

// Value satisfies the Token interface.
func (t tokenImpl) Value() string {
	return t.v
}

// Where satisfies the Token interface.
func (t tokenImpl) Where() where.Where {
	return t.w
}
