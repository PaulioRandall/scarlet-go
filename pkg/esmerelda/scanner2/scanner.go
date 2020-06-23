package scanner2

import (
	. "github.com/PaulioRandall/scarlet-go/pkg/esmerelda/token"
)

type ScanFunc func() (Token, ScanFunc, error)

type SymItr interface {
	Next() (rune, bool)
}

func New(s SymItr) ScanFunc {

	if s == nil {
		panic("PROGRAMMERS ERROR! SymItr input is nil")
	}

	b := &buffer{
		SymItr: s,
	}
	b.bufferNext()

	if b.empty() {
		return nil
	}

	scn := &scanner{
		buffer: b,
	}

	return scn.scan
}

type scanner struct {
	*buffer
	line int
	col  int
}

func (scn *scanner) scan() (Token, ScanFunc, error) {

	if scn.empty() {
		return nil, nil, nil
	}

	ty, runes, e := scan(scn)
	if e != nil {
		return nil, nil, e
	}

	if ty == TK_UNDEFINED {
		progError("Token type not set for " + string(runes))
	}

	if runes == nil || len(runes) == 0 {
		progError("Missing runes for " + ty.String())
	}

	tk := tokenise(scn, ty, runes)
	if scn.empty() {
		return tk, nil, nil
	}

	return tk, scn.scan, nil
}

func tokenise(scn *scanner, ty TokenType, runes []rune) Token {

	val := string(runes)
	tk := NewToken(
		ty,
		val,
		scn.line,
		scn.col,
	)

	update(scn, len(runes), ty == TK_NEWLINE)
	return tk
}

func update(scn *scanner, runeCount int, newline bool) {

	if newline {
		scn.line++
		scn.col = 0
		return
	}

	scn.col += runeCount
}

func progError(msg string) {
	panic("PROGRAMMERS ERROR! " + msg)
}
