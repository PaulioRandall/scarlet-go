package parser

import (
	"github.com/PaulioRandall/scarlet-go/lexor"
	"github.com/PaulioRandall/scarlet-go/token"
)

// TokenBuffer enables token buffering for tokens returned by a ScanToken
// function. Further functionality has been added to enable parser code to be
// more concise.
type TokenBuffer struct {
	buffer  []token.Token
	scanner lexor.ScanToken
	err     lexor.ScanErr
}

// NewTokenBuffer makes a new TokenBuffer using the specified scanner.
func NewTokenBuffer(st lexor.ScanToken) *TokenBuffer {
	return &TokenBuffer{
		scanner: st,
	}
}

// Err returns the scanning error if one has occurred.
func (tb *TokenBuffer) Err() lexor.ScanErr {
	return tb.err
}

// Read satisfies the ScanToken function signature returning the next token in
// the stream.
func (tb *TokenBuffer) Read() (token.Token, lexor.ScanErr) {

	t, e := tb.Peek()

	if e == nil && len(tb.buffer) > 0 {
		tb.buffer = tb.buffer[1:]
	}

	return t, e
}

// Peek returns the next token but does NOT remove it from the buffer.
func (tb *TokenBuffer) Peek() (token.Token, lexor.ScanErr) {

	EMPTY_TOKEN := token.Token{}

	if tb.err != nil {
		return EMPTY_TOKEN, tb.err
	}

	if !tb.tryBuff() || tb.err != nil {
		return EMPTY_TOKEN, tb.err
	}

	return tb.buffer[0], nil
}

// HasMore returns true if there are tokens remaining to be read.
func (tb *TokenBuffer) HasMore() bool {
	return tb.bufferHasMore() || tb.scannerHasMore()
}

// Push allows a token to be pushed into the front of the buffer so it becomes
// the next token returned by a read.
func (tb *TokenBuffer) Push(t token.Token) lexor.ScanErr {

	if tb.err == nil {
		tb.buffer = append([]token.Token{t}, tb.buffer...)
	}

	return tb.err
}

// tryBuff attempts to scan in the next token returning true if the buffer
// contains at least one token.
func (tb *TokenBuffer) tryBuff() bool {

	if tb.bufferHasMore() {
		return true
	}

	if tb.scannerHasMore() {
		return tb.buff()
	}

	return false
}

// buff scans in another token and adds it to the buffer. Assumes that the
// scanner contains at least one more token.
func (tb *TokenBuffer) buff() bool {

	var t token.Token
	t, tb.scanner, tb.err = tb.scanner()

	if tb.err != nil {
		return false
	}

	tb.buffer = append(tb.buffer, t)
	return true
}

// bufferHasMore returns true if there are tokens sitting in the buffer.
func (tb *TokenBuffer) bufferHasMore() bool {
	return len(tb.buffer) > 0
}

// scannerHasMore returns true if there are more tokens still to scan via the
// scanner.
func (tb *TokenBuffer) scannerHasMore() bool {
	return tb.scanner != nil
}
