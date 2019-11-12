package token

import (
	"github.com/PaulioRandall/scarlet-go/cookies/perror"
	"github.com/PaulioRandall/scarlet-go/cookies/where"
)

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

// EmptyTok returns an empty Token.
func Empty() Token {
	return Token{}
}
