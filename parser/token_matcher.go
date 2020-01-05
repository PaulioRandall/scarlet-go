package parser

import (
	"github.com/PaulioRandall/scarlet-go/lexor"
	"github.com/PaulioRandall/scarlet-go/token"
)

// TokenMatcher provides functionality to buffer and match tokens from a token
// stream.
type TokenMatcher struct {
	buffer  []token.Token
	scanner lexor.ScanToken
}

// Matcher is the function signature for all match functions.
type Matcher func() int

// NewTokenMatcher creates a new token matcher.
func NewTokenMatcher(sc lexor.ScanToken) *TokenMatcher {
	return &TokenMatcher{
		scanner: sc,
	}
}

// Skip performs a single Read without returning the token.
func (tm *TokenMatcher) Skip() bool {
	_, ok := tm.Read()
	return ok
}

// SkipMany performs a ReadMany with `n` as input but does NOT return the token
// slice.
func (tm *TokenMatcher) SkipMany(n int) int {
	_, n = tm.ReadMany(n)
	return n
}

// Peek returns the next token in the stream without removing it from the
// stream. True is returned if the token is valid and the end of the stream
// has NOT been reached.
func (tm *TokenMatcher) Peek() (_ token.Token, _ bool) {

	if 0 == len(tm.buffer) {
		if 0 == tm.readIntoBuffer(1) {
			return
		}
	}

	return tm.buffer[0], true
}

// Read returns the next token in the stream. True is returned if the token is
// valid and the end of the stream has NOT been reached.
func (tm *TokenMatcher) Read() (_ token.Token, _ bool) {
	r, ok := tm.Peek()

	if ok {
		tm.buffer = tm.buffer[1:]
	}

	return r, ok
}

// Read returns the next `n` tokens in the stream. The length of the slice is
// returned; if it is less than the input `n` then the end of the stream has
// been reached and all remaining tokens were returned.
func (tm *TokenMatcher) ReadMany(n int) (_ []token.Token, _ int) {

	bufLen := len(tm.buffer)
	need := n - bufLen

	if need > 0 {
		tm.readIntoBuffer(need)
		n = len(tm.buffer)
	}

	r := tm.buffer[0:n]
	tm.buffer = tm.buffer[n:]

	return r, n
}

// Match compares the input kind `k` with the next tokens kind returning 1 if
// there was a match; else returns 0.
func (tm *TokenMatcher) Match(k token.Kind) (_ int) {

	if 0 == len(tm.buffer) {
		if 0 == tm.readIntoBuffer(1) {
			return
		}
	}

	if k == tm.buffer[0].Kind {
		return 1
	}

	return
}

// Match compares the input kinds `ks` with the next tokens kind returning 1 if
// any of them matched; else returns 0.
func (tm *TokenMatcher) MatchAny(ks ...token.Kind) (_ int) {

	if 0 == len(tm.buffer) {
		if 0 == tm.readIntoBuffer(1) {
			return
		}
	}

	for _, k := range ks {
		if k == tm.buffer[0].Kind {
			return 1
		}
	}

	return
}

// Match compares the sequence of input kinds `seq` with the kinds in the next
// sequence of tokens returning the length of `seq` if the sequences match in
// both kind and order; else returns 0.
func (tm *TokenMatcher) MatchSeq(seq ...token.Kind) (_ int) {

	seqLen := len(seq)
	need := seqLen - len(tm.buffer)

	if need > 0 {
		if tm.readIntoBuffer(need) < need {
			return
		}
	}

	for i, k := range seq {
		if k != tm.buffer[i].Kind {
			return
		}
	}

	return seqLen
}

// MatchRepeat compares the sequence of input kinds `seq` with the kinds in the
// next sequence allowing for the input sequence to repeat any number of times
// (longest match is used). The length of the repeating sequence is returned if
// it matches at least once else 0 is returned.
func (tm *TokenMatcher) MatchRepeat(seq ...token.Kind) (n int) {

	seqLen := len(seq)

	for {
		need := n + seqLen - len(tm.buffer)

		if need > 0 {
			if tm.readIntoBuffer(need) < need {
				return
			}
		}

		for i, k := range seq {
			if k != tm.buffer[n+i].Kind {
				return
			}
		}

		n += seqLen
	}

	return
}

// HasMore returns true if there are more tokens in the stream, however, the
// last token may be a zero value.
func (tm *TokenMatcher) HasMore() bool {
	return len(tm.buffer) > 0 || tm.scanner != nil
}

// readIntoBuffer reads upto `n` tokens into the buffer. The number of tokens
// added is returned; if it is less than the input `n` then the end of the
// stream has been reached and all remaining tokens were buffered.
func (tm *TokenMatcher) readIntoBuffer(n int) (totalRead int) {

	var t token.Token
	var e lexor.ScanErr

	for totalRead < n && tm.scanner != nil {

		t, tm.scanner, e = tm.scanner()
		if e != nil {
			panic(e)
		}

		if t.IsZero() {
			// Last token in the stream may have been removed by the evaluator.
			tm.scanner = nil
			break
		}

		tm.buffer = append(tm.buffer, t)
		totalRead++
	}

	return
}
