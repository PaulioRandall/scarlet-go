package token

import (
	"fmt"
	"strconv"
)

// Kind represents a token type.
type Kind string

const (
	UNDEFINED Kind = ``
	// ------------------
	COMMENT      Kind = `COMMENT`
	WHITESPACE   Kind = `WHITESPACE`
	NEWLINE      Kind = `NEWLINE`
	GLOBAL       Kind = `GLOBAL`
	FUNC         Kind = `FUNC`
	DO           Kind = `DO`
	MATCH        Kind = `MATCH`
	WATCH        Kind = `WATCH`
	END          Kind = `END`
	ID           Kind = `ID`
	DELIM        Kind = `DELIM`
	ASSIGN       Kind = `ASSIGN`
	RETURNS      Kind = `RETURNS`
	OPEN_PAREN   Kind = `OPEN_PAREN`
	CLOSE_PAREN  Kind = `CLOSE_PAREN`
	OPEN_GUARD   Kind = `OPEN_GUARD`
	CLOSE_GUARD  Kind = `CLOSE_GUARD`
	OPEN_LIST    Kind = `OPEN_LIST`
	CLOSE_LIST   Kind = `CLOSE_LIST`
	SPELL        Kind = `SPELL`
	STR_LITERAL  Kind = `STR_LITERAL`
	STR_TEMPLATE Kind = `STR_TEMPLATE`
	INT_LITERAL  Kind = `INT_LITERAL`
	REAL_LITERAL Kind = `REAL_LITERAL`
	BOOL_LITERAL Kind = `BOOL`
	NOT          Kind = `NOT`
	OPERATOR     Kind = `OPERATOR`
	VOID         Kind = `VOID`
)

// Token represents a grammer token within a source file.
type Token struct {
	Kind  Kind
	Value string
	Line  int
	Col   int
}

// New creates a new token.
func New(k Kind, v string, l, c int) Token {
	return Token{
		Kind:  k,
		Value: v,
		Line:  l,
		Col:   c,
	}
}

// OfKind creates a new token with the specified kind.
func OfKind(k Kind) Token {
	return Token{
		Kind: k,
	}
}

// OfValue creates a new token with the specified kind and value.
func OfValue(k Kind, v string) Token {
	return Token{
		Kind:  k,
		Value: v,
	}
}

// ZERO returns a zero token value.
func ZERO() Token {
	return Token{}
}

// IsZero returns true if the token is a zero value.
func (t Token) IsZero() bool {
	return t == Token{}
}

// IsNotZero returns true if the token is NOT a zero value.
func (t Token) IsNotZero() bool {
	return t != Token{}
}

// String returns a string representation of the token.
func (t Token) String() string {
	s := strconv.QuoteToGraphic(t.Value)
	return fmt.Sprintf(`%d: '%s' (%s)`, t.Line, s, t.Kind)
}
