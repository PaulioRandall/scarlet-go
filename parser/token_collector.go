package parser

import (
	"github.com/PaulioRandall/scarlet-go/token"
)

// TokenCollector mimics bulk token put back functionality for a TokenReader.
type TokenCollector struct {
	buffer []token.Token
	index  int
	reader *TokenReader
}

// TokenCollector makes a new TokenCollector using the specified reader.
func NewTokenCollector(r *TokenReader) *TokenCollector {
	return &TokenCollector{
		reader: r,
	}
}

// HasMore returns true if there are tokens remaining to be read. This excludes
// tokens read but buffered.
func (tc *TokenCollector) HasMore() bool {
	return tc.index < len(tc.buffer) || tc.reader.HasMore()
}

// Read returns the next token in the stream. If an empty token is returned
// then either an error has occurred or the end of the token stream has been
// reached.
func (tc *TokenCollector) Read() token.Token {

	t := tc.Peek()
	if t.IsNotZero() {
		tc.index++
	}

	return t
}

// Peek returns the next token without iterating to the one after. If an empty
// token is returned then either an error has occurred or the end of the token
// stream has been reached.
func (tc *TokenCollector) Peek() token.Token {

	if tc.index < len(tc.buffer) {
		return tc.buffer[tc.index]
	}

	if tc.HasMore() && tc.tryBuff() {
		return tc.buffer[tc.index]
	}

	return token.ZERO()
}

// PutBack puts back upto `n` tokens so they may be re-read.
func (tc *TokenCollector) PutBack(n int) {
	for n > 0 && tc.index > 0 {
		n--
		tc.index--
	}
}

// PutBackAll puts back all tokens so they may be re-read.
func (tc *TokenCollector) PutBackAll() {
	tc.index = 0
}

// Clear removes any collected tokens so they can no longer be put back.
func (tc *TokenCollector) Clear(n int) {
	tc.buffer = []token.Token{}
	tc.index = 0
}

// tryBuff reads in another token and adds it to the buffer. Returns true if a
// valid token was added to the buffer.
func (tc *TokenCollector) tryBuff() bool {

	t := tc.reader.Read()

	if t.IsZero() {
		return false
	}

	tc.buffer = append(tc.buffer, t)
	return true
}
