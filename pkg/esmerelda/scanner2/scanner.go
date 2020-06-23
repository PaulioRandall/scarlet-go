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

	lex := lexeme{
		scn: scn,
	}
	if e := lex.scan(); e != nil {
		return nil, nil, e
	}

	tk := scn.tokenise(lex)
	if scn.empty() {
		return tk, nil, nil
	}

	return tk, scn.scan, nil
}

func (scn *scanner) tokenise(lex lexeme) Token {

	if lex.ty == TK_UNDEFINED {
		panic("PROGRAMMERS ERROR! Token type not set")
	}

	val := string(lex.tok)
	tk := NewToken(
		lex.ty,
		val,
		scn.line,
		scn.col,
	)

	scn.update(len(val), lex.ty == TK_NEWLINE)
	return tk
}

func (scn *scanner) update(runeCount int, newline bool) {

	if newline {
		scn.line++
		scn.col = 0
		return
	}

	scn.col += runeCount
}
