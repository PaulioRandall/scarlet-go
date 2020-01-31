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
	SOF          Kind = `SOF`
	EOF          Kind = `EOF`
	COMMENT      Kind = `COMMENT`
	WHITESPACE   Kind = `WHITESPACE`
	NEWLINE      Kind = `NEWLINE`
	STICKY       Kind = `STICKY`
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
	TERMINATOR   Kind = `TERMINATOR`
)

// Token represents a grammer token within a source file.
type Token struct {
	Kind  Kind
	Value string
	Line  int
	Col   int
}

// KindsToStrings converts the kind slice to a string slice.
func KindsToStrings(ks []Kind) (strs []string) {

	for _, k := range ks {
		strs = append(strs, string(k))
	}

	return
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
func (tk Token) String() string {

	var v interface{}

	if tk.Kind == STR_TEMPLATE {
		v = strconv.QuoteToGraphic(tk.Value)
	} else if tk.Kind == STR_LITERAL {
		v = "`" + tk.Value + "`"
	} else {
		v = tk.Value
	}

	// +1 for line index to number
	return fmt.Sprintf(`%d:%d %s %v`, tk.Line+1, tk.Col, tk.Kind, v)
}
