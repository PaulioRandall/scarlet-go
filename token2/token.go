package token

import (
	"fmt"
)

// Kind represents a token type.
type Kind string

const (
	UNDEFINED Kind = ``
	// ------------------
	COMMENT     Kind = `COMMENT`
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
