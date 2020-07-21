package scanner

import (
	"fmt"

	"github.com/PaulioRandall/scarlet-go/pkg/eskarina/lexeme"
	"github.com/PaulioRandall/scarlet-go/pkg/eskarina/perror"
	"github.com/PaulioRandall/scarlet-go/pkg/eskarina/prop"
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
	start int
	end   int
	read  bool
}

func (lr *lexReader) accept(ru rune) bool {

	if lr.empty() {
		return false
	}

	if lr.peek() == ru {
		lr.idx++
		lr.read = true
		return true
	}

	return false
}

func (lr *lexReader) expect(ru rune) error {

	if lr.accept(ru) {
		return nil
	}

	if lr.empty() {
		return perror.New("Unexpected EOF %d:%d, wanted %q", lr.line, lr.end, ru)
	}

	return perror.New(
		"Unexpected terminal symbol %d:%d, want %q, have %q",
		lr.line, lr.end, ru, lr.peek(),
	)
}

func (lr *lexReader) slice(props ...prop.Prop) *lexeme.Lexeme {

	if !lr.read {
		failNow("You haven't accepted any terminal symbols yet")
	}

	begin := lr.idx - (lr.end - lr.start) - 1

	lex := &lexeme.Lexeme{
		Props: props,
		Raw:   string(lr.runes[begin:lr.idx]),
		Line:  lr.line,
		Col:   lr.start,
	}

	if lex.Has(prop.PR_NEWLINE) {
		lr.line++
		lr.start = 0
		lr.end = 0
	} else {
		lr.start = lr.end
	}

	lr.read = false
	return lex
}

func failNow(msg string, args ...interface{}) {
	msg = fmt.Sprintf(msg, args...)
	panic(fmt.Errorf("SANITY CHECK! %s", msg))
}
