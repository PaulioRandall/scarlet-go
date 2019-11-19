package token

import (
	"github.com/PaulioRandall/scarlet-go/cookies/perror"
	"github.com/PaulioRandall/scarlet-go/cookies/where"
)

type TokenInterface interface {

	// Value returns the string representing the token in source.
	O_Value() string

	// Kind returns the type of the token.
	O_Kind() Kind

	// Where returns where the token is located within the source.
	O_Where() where.Where
}

// Token represents a grammer token within a source file.
type Token struct {
	Value string      // The value of the token within the source code
	Kind  Kind        // The type of the token
	Where where.Where // The location of the token within the source file
}

// ScanToken is a recursive descent function that returns the next token
// followed by the callable (tail) function to get the token after next. If the
// function is null then the end of the token stream has been reached.
type ScanToken func() (Token, ScanToken, perror.Perror)

// New creates a new token.
func New(v string, k Kind, w where.Where) Token {
	return Token{
		Value: v,
		Kind:  k,
		Where: w,
	}
}

// EmptyTok returns an empty Token.
func Empty() Token {
	return Token{}
}

// Value satisfies the TokenInterface interface.
func (t Token) O_Value() string {
	return t.Value
}

// Kind satisfies the TokenInterface interface.
func (t Token) O_Kind() Kind {
	return t.Kind
}

// Where satisfies the TokenInterface interface.
func (t Token) O_Where() where.Where {
	return t.Where
}

// IsSignificant returns true if the token is required for parsing the program.
// Better put, false is returned if the token is whitespace or a comment etc.
func (t Token) IsSignificant() bool {
	switch t.Kind {
	case UNDEFINED:
	case WHITESPACE:
	default:
		return true
	}

	return false
}
