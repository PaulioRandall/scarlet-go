package parser

import (
	"fmt"

	"github.com/PaulioRandall/scarlet-go/token"
)

// ParseErr represents an error while syntax parsing.
type ParseErr interface {
	error

	// Token returns the token that caused the error.
	Token() token.Token
}

// perr is the standard ParseErr implementation.
type perr struct {
	what string
	why  error
	tok  token.Token
}

// NewParseErr returns a new instance of ParseErr.
func NewParseErr(what string, why error, tok token.Token) ParseErr {
	return perr{
		what: what,
		why:  why,
		tok:  tok,
	}
}

// Error satisfies the error interface.
func (e perr) Error() string {

	s := fmt.Sprintf("%v: %s", e.tok.String(), e.what)

	if e.why != nil {
		s += fmt.Sprintf("\n\t...caused by: %s", e.why.Error())
	}

	return s
}

// Token satisfies the ParseErr interface.
func (e perr) Token() token.Token {
	return e.tok
}
