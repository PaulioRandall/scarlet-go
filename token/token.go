package token

import (
	"github.com/PaulioRandall/scarlet-go/perror"
	"github.com/PaulioRandall/scarlet-go/where"
)

// Token represents a grammer token within a source file.
type Token interface {

	// Value returns the string representing the token in source.
	Value() string

	// Kind returns the type of the token.
	Kind() Kind

	// Where returns where the token is located within the source.
	Where() where.Where
}

// tokenSimple is a simple implementation of the Token interface.
type tokenSimple struct {
	v string
	k Kind
	w where.Where
}

// ScanToken is a recursive descent function that returns the next token
// followed by the callable (tail) function to get the token after next. If the
// function is null then the end of the token stream has been reached.
type ScanToken func() (Token, ScanToken, perror.Perror)

// Newish creates a new token.
func Newish(v string, k Kind, w where.Where) Token {
	return tokenSimple{
		v: v,
		k: k,
		w: w,
	}
}

// New creates a new token.
func New(v string, k Kind, line, start, end int) Token {
	return tokenSimple{
		v: v,
		k: k,
		w: where.New(line, start, end),
	}
}

// EmptyTok returns an empty Token.
func Empty() Token {
	return tokenSimple{}
}

// Value satisfies the Token interface.
func (t tokenSimple) Value() string {
	return t.v
}

// Kind satisfies the Token interface.
func (t tokenSimple) Kind() Kind {
	return t.k
}

// Where satisfies the Token interface.
func (t tokenSimple) Where() where.Where {
	return t.w
}
