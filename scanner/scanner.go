package scanner

import (
	"github.com/PaulioRandall/scarlet-go/token"
)

type source struct {
	s []rune // The source code
	i int    // The rune index
}

// ScanToken is a recursive descent function that returns the next token
// followed by the callable tail function to get the token after next. If the
// function is null then the end of the token stream has been reached.
type ScanToken func() (token.Token, ScanToken)

// no_tok returns an empty Token.
func no_tok() token.Token {
	return token.Token{}
}

// New returns a ScanToken function that will return the first token in the
// input source.
func New(src string) ScanToken {
	s := source{
		s: []rune(src),
	}
	return s.fileScope
}

// fileScope identifies and returns the next token in the source. The token must
// be one that appears at the start of a statement within the top level of a
// source file.
func (src *source) fileScope() (t token.Token, f ScanToken) {
	return no_tok(), nil
}
