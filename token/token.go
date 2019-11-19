package token

import (
	"github.com/PaulioRandall/scarlet-go/cookies/perror"
	"github.com/PaulioRandall/scarlet-go/cookies/where"
)

// Token represents a grammer token within a source file.
type Token interface {

	// Value returns the string representing the token in source.
	Value() string

	// Kind returns the type of the token.
	Kind() Kind

	// Where returns where the token is located within the source.
	Where() where.Where

	// IsSignificant returns true if the token is required for parsing the
	// program. Better put, false is returned if the token is whitespace or a
	// comment etc.
	IsSignificant() bool
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

// New creates a new token.
func New(v string, k Kind, w where.Where) Token {
	return tokenSimple{
		v: v,
		k: k,
		w: w,
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

// IsSignificant satisfies the Token interface.
func (t tokenSimple) IsSignificant() bool {
	switch t.k {
	case UNDEFINED:
	case WHITESPACE:
	default:
		return true
	}

	return false
}
