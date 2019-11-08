package scanner

import (
	"github.com/PaulioRandall/scarlet-go/token"
)

// ScanToken is a recursive descent function that returns the next token
// followed by the callable (tail) function to get the token after next. If the
// function is null then the end of the token stream has been reached.
type ScanToken func() (token.Token, ScanToken)

// source represents the source code yet to be scanned.
type source struct {
	rs []rune // The source code
	i int    // The rune index
}

// scan removes `n` runes from the unscanned source code and returns it. The
// index is updated accordingly.
func (s* source) scan(n int) (out string, index int) {
	out = s.rs[:n]
	s.rs = s.rs[n:]
	s.i += n
	return out, s.i
}


// no_tok returns an empty Token.
func no_tok() token.Token {
	return token.Token{}
}

// New returns a ScanToken function that will return the first token in the
// input source.
func New(src string) ScanToken {
	s := source{
		rs: []rune(src),
	}
	return s.fileScope
}

// fileScope identifies and returns the next token in the source. The token must
// be one that appears at the start of a statement within the top level of a
// source file.
func (s* source) fileScope() (t token.Token, f ScanToken) {
	return no_tok(), nil
}
