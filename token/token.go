package token

import (
	"fmt"
	"strconv"
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

	if tk.Kind == TEMPLATE {
		v = strconv.QuoteToGraphic(tk.Value)
	} else if tk.Kind == STR {
		v = "`" + tk.Value + "`"
	} else {
		v = tk.Value
	}

	// +1 for line index to number
	return fmt.Sprintf(`%d:%d %s %v`, tk.Line+1, tk.Col, tk.Kind, v)
}
