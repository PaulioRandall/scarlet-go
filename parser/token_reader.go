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
func (tr *TokenReader) Err() lexor.ScanErr {
	return tr.err
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

	EMPTY_TOKEN := token.Token{}

	switch {
	case tr.err != nil:
		return EMPTY_TOKEN
	case tr.buffer != nil:
		return *tr.buffer
	case tr.scanner == nil:
		return EMPTY_TOKEN
	case tr.buff():
		return *tr.buffer
	}

	return EMPTY_TOKEN
}

// buff scans in another token and points the buffer to it. Assumes that the
// current buffer content is no longer needed and the scanner contains at least
// one more token.
func (tr *TokenReader) buff() bool {

	var t token.Token
	tr.buffer = nil

	t, tr.scanner, tr.err = tr.scanner()

	if tr.err != nil {
		return false
	}

	tr.buffer = &t
	return true
}
