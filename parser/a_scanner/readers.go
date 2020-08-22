package scanner

import (
	"fmt"
	"unicode"

	"github.com/PaulioRandall/scarlet-go/lexeme"
	"github.com/PaulioRandall/scarlet-go/perror"
)

type runeReader struct {
	runes []rune
	size  int
	idx   int
	line  int
	col   int
	count int
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

func (rr *runeReader) inc() {
	rr.idx++
	rr.count++
}

func (rr *runeReader) is(ru rune) bool {
	return rr.more() && rr.peek() == ru
}

func (rr *runeReader) isNewline() bool {
	return rr.more() && (rr.peek() == '\r' || rr.peek() == '\n')
}

func (rr *runeReader) isSpace() bool {
	return rr.more() && unicode.IsSpace(rr.peek())
}

func (rr *runeReader) isLetter() bool {
	return rr.more() && unicode.IsLetter(rr.peek())
}

func (rr *runeReader) isDigit() bool {
	return rr.more() && unicode.IsDigit(rr.peek())
}

func (rr *runeReader) accept(ru rune) bool {

	if rr.more() && rr.peek() == ru {
		rr.inc()
		return true
	}

	return false
}

func (rr *runeReader) expect(ru rune) error {

	if rr.accept(ru) {
		return nil
	}

	if rr.empty() {
		return perror.New(
			"Unexpected EOF %d:%d, wanted %q",
			rr.line, rr.idx-rr.count, ru,
		)
	}

	return perror.New(
		"Unexpected terminal symbol %d:%d, want %q, have %q",
		rr.line, rr.idx-rr.count, ru, rr.peek(),
	)
}

func (rr *runeReader) slice(tk lexeme.Token) *lexeme.Lexeme {

	if rr.count == 0 {
		failNow("You haven't accepted any terminal symbols yet")
	}

	lex := &lexeme.Lexeme{
		Tok: tk,
		Raw: string(rr.runes[rr.idx-rr.count : rr.idx]),
	}

	rr.update(lex)
	return lex
}

func (rr *runeReader) update(lex *lexeme.Lexeme) {

	lex.Line = rr.line
	lex.Col = rr.col

	if lex.Tok == lexeme.NEWLINE {
		rr.line++
		rr.col = 0
	} else {
		rr.col += len(lex.Raw)
	}

	rr.count = 0
}

func (rr *runeReader) syntaxError(msg string, args interface{}) error {
	msg = fmt.Sprintf(msg, args)
	return perror.New("At %d:%d, %s", rr.line+1, rr.col, msg)
}

func failNow(msg string, args ...interface{}) {
	msg = fmt.Sprintf(msg, args...)
	panic(fmt.Errorf("SANITY CHECK! %s", msg))
}
