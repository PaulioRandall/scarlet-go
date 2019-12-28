package parser

import (
	"github.com/PaulioRandall/scarlet-go/lexor"
	"github.com/PaulioRandall/scarlet-go/token"
)

// TokenReader allows peeking of the next token returned by a ScanToken
// function.
type TokenReader struct {
	buffer  *token.Token
	scanner lexor.ScanToken
}

// NewTokenReader makes a new TokenReader using the specified scanner.
func NewTokenReader(st lexor.ScanToken) *TokenReader {
	return &TokenReader{
		scanner: st,
	}
}

// HasMore returns true if there are tokens remaining to be read.
func (tr *TokenReader) HasMore() bool {
	return tr.buffer != nil || tr.scanner != nil
}

// Read returns the next token in the stream.
func (tr *TokenReader) Read() (t token.Token) {

	if tr.buffer == nil {
		t = tr.Peek()
	}

	tr.buffer = nil
	return t
}

// Peek returns the next token without iterating to the one after.
func (tr *TokenReader) Peek() token.Token {

	if tr.buffer != nil {
		return *tr.buffer
	}

	if tr.scanner != nil {
		tr.buff()
		return *tr.buffer
	}

	return token.ZERO()
}

// buff scans in another token and points the buffer to it. Assumes that the
// current buffer content is no longer needed and the scanner contains at least
// one more token.
func (tr *TokenReader) buff() {

	var t token.Token
	var e lexor.ScanErr
	tr.buffer = nil

	t, tr.scanner, e = tr.scanner()

	if e != nil {
		panic(e)
	}

	tr.buffer = &t
}
