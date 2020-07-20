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
	line  int
	col   int
	idx   int
	flag  bool
	prev  *lexeme.Lexeme
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

func (rr *runeReader) accept(ru rune) bool {

	if rr.empty() {
		return false
	}

	if rr.runes[rr.idx] == ru {
		rr.idx++
		rr.flag = true
		return true
	}

	return false
}

func (rr *runeReader) expect(ru rune) error {

	if rr.accept(ru) {
		return nil
	}

	return perror.New("Unexpected rune %d:%d %q", rr.line, rr.idx, ru)
}

func (rr *runeReader) slice(props ...prop.Prop) *lexeme.Lexeme {

	if !rr.flag {
		failNow("You haven't accepted any runes yet")
	}

	lex := &lexeme.Lexeme{
		Props: props,
		Raw:   string(rr.runes[rr.col:rr.idx]),
		Line:  rr.line,
		Col:   rr.col,
		Prev:  rr.prev,
	}

	rr.prev = lex
	if lex.Prev != nil {
		lex.Prev.Next = lex
	}

	if lex.Has(prop.PR_NEWLINE) {
		rr.line++
		rr.col = 0
		rr.idx = 0
	} else {
		rr.col = rr.idx
	}

	return lex
}

func failNow(msg string, args ...interface{}) {
	msg = fmt.Sprintf(msg, args...)
	panic(fmt.Errorf("SANITY CHECK! %s", msg))
}
