package token

import (
	"github.com/PaulioRandall/scarlet-go/where"
)

// Kind represents a token type.
type Kind string

const (
	UNDEFINED Kind = ``
	// ------------------
	WHITESPACE  Kind = `WHITESPACE`
	NEWLINE     Kind = `NEWLINE`
	FUNC        Kind = `FUNC`
	DO          Kind = `DO`
	END         Kind = `END`
	ID          Kind = `ID`
	ID_DELIM    Kind = `ID_DELIM`
	ASSIGN      Kind = `ASSIGN`
	OPEN_PAREN  Kind = `OPEN_PAREN`
	CLOSE_PAREN Kind = `CLOSE_PAREN`
	SPELL       Kind = `SPELL`
	STR_LITERAL Kind = `STR_LITERAL`
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

// New creates a new token.
func New(k Kind, v string, line, start, end int) Token {
	return tokenImpl{
		k: k,
		v: v,
		w: where.New(line, start, end),
	}
}

// NewByWhere creates a new token using Where as an input parameter.
func NewByWhere(k Kind, v string, w where.Where) Token {
	return tokenImpl{
		k: k,
		v: v,
		w: w,
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
