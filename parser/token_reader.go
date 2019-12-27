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
	err     lexor.ScanErr
}

// NewTokenReader makes a new TokenReader using the specified scanner.
func NewTokenReader(st lexor.ScanToken) *TokenReader {
	return &TokenReader{
		scanner: st,
	}
}

// Err returns the scanning error if one has occurred.
func (tb *TokenReader) Err() lexor.ScanErr {
	return tb.err
}

// HasMore returns true if there are tokens remaining to be read.
func (tb *TokenReader) HasMore() bool {
	return tb.buffer != nil || tb.scanner != nil
}

// Read returns the next token in the stream.
func (tb *TokenReader) Read() (t token.Token) {

	if tb.buffer == nil {
		t = tb.Peek()
	}

	tb.buffer = nil
	return t
}

// Peek returns the next token without iterating to the one after.
func (tb *TokenReader) Peek() token.Token {

	EMPTY_TOKEN := token.Token{}

	switch {
	case tb.err != nil:
		return EMPTY_TOKEN
	case tb.buffer != nil:
		return *tb.buffer
	case tb.scanner == nil:
		return EMPTY_TOKEN
	case tb.buff():
		return *tb.buffer
	}

	return EMPTY_TOKEN
}

// buff scans in another token and points the buffer to it. Assumes that the
// current buffer content is no longer needed and the scanner contains at least
// one more token.
func (tb *TokenReader) buff() bool {

	var t token.Token
	tb.buffer = nil

	t, tb.scanner, tb.err = tb.scanner()

	if tb.err != nil {
		return false
	}

	tb.buffer = &t
	return true
}
