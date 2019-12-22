package token

import (
	"fmt"
)

// Kind represents a token type.
type Kind string

const (
	UNDEFINED Kind = ``
	// ------------------
	COMMENT      Kind = `COMMENT`
	WHITESPACE   Kind = `WHITESPACE`
	NEWLINE      Kind = `NEWLINE`
	FUNC         Kind = `FUNC`
	DO           Kind = `DO`
	END          Kind = `END`
	ID           Kind = `ID`
	ID_DELIM     Kind = `ID_DELIM`
	ASSIGN       Kind = `ASSIGN`
	OPEN_PAREN   Kind = `OPEN_PAREN`
	CLOSE_PAREN  Kind = `CLOSE_PAREN`
	SPELL        Kind = `SPELL`
	STR_LITERAL  Kind = `STR_LITERAL`
	STR_TEMPLATE Kind = `STR_TEMPLATE`
	NUM_LITERAL  Kind = `NUM_LITERAL`
	TRUE         Kind = `TRUE`
	FALSE        Kind = `FALSE`
	OPEN_LIST    Kind = `OPEN_LIST`
	CLOSE_LIST   Kind = `CLOSE_LIST`
	ADD          Kind = `ADD`
	SUBTRACT     Kind = `SUBTRACT`
	DIVIDE       Kind = `DIVIDE`
	MULTIPLY     Kind = `MULTIPLY`
	MODULO       Kind = `MODULO`
	OR           Kind = `OR`
	AND          Kind = `AND`
	EQUAL        Kind = `EQUAL`
	NOT_EQUAL    Kind = `NOT_EQUAL`
	LT           Kind = `LT`
	GT           Kind = `GT`
	LT_OR_EQUAL  Kind = `LT_OR_EQUAL`
	GT_OR_EQUAL  Kind = `GT_OR_EQUAL`
)

// Token represents a grammer token within a source file.
type Token struct {
	Kind  Kind
	Value string
	Line  int
	Col   int
}

// NewToken creates a new token.
func NewToken(k Kind, v string, l, c int) Token {
	return Token{
		Kind:  k,
		Value: v,
		Line:  l,
		Col:   c,
	}
}

// String returns a string representation of the token.
func (t Token) String() string {
	return fmt.Sprintf(`%d: '%s' (%s)`, t.Line, t.Value, t.Kind)
}
