package scanner

import (
	"fmt"
	"unicode"

	"github.com/PaulioRandall/scarlet-go/eskarina/shared/lexeme"
	"github.com/PaulioRandall/scarlet-go/eskarina/shared/perror"
)

type runeReader struct {
	runes []rune
	size  int
	idx   int
}

func (rr *runeReader) empty() bool {
	return rr.idx >= rr.size
}

func (rr *runeReader) more() bool {
	return rr.idx < rr.size
}

func (rr *runeReader) peek() rune {
	return rr.runes[rr.idx]
}

type lexReader struct {
	*runeReader
	line  int
	col   int
	start int
	read  bool
}

func (lr *lexReader) inc() {
	lr.idx++
	lr.read = true
}

func (lr *lexReader) is(ru rune) bool {
	return lr.more() && lr.peek() == ru
}

func (lr *lexReader) isNewline() bool {
	return lr.more() && (lr.peek() == '\r' || lr.peek() == '\n')
}

func (lr *lexReader) isSpace() bool {
	return lr.more() && unicode.IsSpace(lr.peek())
}

func (lr *lexReader) isLetter() bool {
	return lr.more() && unicode.IsLetter(lr.peek())
}

func (lr *lexReader) isDigit() bool {
	return lr.more() && unicode.IsDigit(lr.peek())
}

func (lr *lexReader) accept(ru rune) bool {

	if lr.more() && lr.peek() == ru {
		lr.inc()
		return true
	}

	return false
}

func (lr *lexReader) expect(ru rune) error {

	if lr.accept(ru) {
		return nil
	}

	if lr.empty() {
		return perror.New(
			"Unexpected EOF %d:%d, wanted %q",
			lr.line, lr.idx-lr.start, ru,
		)
	}

	return perror.New(
		"Unexpected terminal symbol %d:%d, want %q, have %q",
		lr.line, lr.idx-lr.start, ru, lr.peek(),
	)
}

func (lr *lexReader) slice(tk lexeme.Token) *lexeme.Lexeme {

	if !lr.read {
		failNow("You haven't accepted any terminal symbols yet")
	}

	lex := &lexeme.Lexeme{
		Tok:  tk,
		Raw:  string(lr.runes[lr.start:lr.idx]),
		Line: lr.line,
		Col:  lr.start,
	}

	if lex.Tok == lexeme.NEWLINE {
		lr.line++
		lr.col = 0
	} else {
		lr.col = lr.idx - lr.start
	}

	lr.start, lr.read = lr.idx, false
	return lex
}

func failNow(msg string, args ...interface{}) {
	msg = fmt.Sprintf(msg, args...)
	panic(fmt.Errorf("SANITY CHECK! %s", msg))
}
