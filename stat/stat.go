package stat

import (
	"github.com/PaulioRandall/scarlet-go/token"
)

// Statement represents program statement.
type Statement interface {

	// Kind returns the type of the statement.
	Kind() Kind

	// Tokens returns the tokens that make up the statement including those of
	// sub-statements.
	Tokens() []token.Token

	// Statements returns any sub-statements but not those of any sub-statements.
	Statements() []Statement
}

// statImpl is a simple implementation of the Statement interface.
type statImpl struct {
	k Kind
	t []token.Token
	s []Statement
}

// New creates a new statment.
func New(k Kind, t []token.Token, s []Statement) Statement {
	return statImpl{
		k: k,
		t: t,
		s: s,
	}
}

// Kind satisfies the Statement interface.
func (s statImpl) Kind() Kind {
	return s.k
}

// Tokens satisfies the Statement interface.
func (s statImpl) Tokens() []token.Token {
	return s.t
}

// Statements satisfies the Statement interface.
func (s statImpl) Statements() []Statement {
	return s.s
}
