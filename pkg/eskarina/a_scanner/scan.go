package scanner

import (
	"github.com/PaulioRandall/scarlet-go/pkg/eskarina/lexeme"
	"github.com/PaulioRandall/scarlet-go/pkg/eskarina/perror"
	"github.com/PaulioRandall/scarlet-go/pkg/eskarina/prop"
)

func ScanAll(s string) (*lexeme.Lexeme, error) {

	rr := &runeReader{}
	rr.runes = []rune(s)
	rr.size = len(rr.runes)

	var lex *lexeme.Lexeme

	for rr.more() {

		l, e := scanLexeme(rr)
		if e != nil {
			return nil, e
		}

		if lex != nil {
			l.Prev = lex
			lex.Next = l
		}

		lex = l
	}

	return lex, nil
}

func scanLexeme(rr *runeReader) (*lexeme.Lexeme, error) {

	switch rr.peek() {
	case '\r', '\n':
		return scanNewline(rr)
	}

	return nil, perror.New(
		"Unexpected terminal symbol %d:%d, have %q",
		rr.line, rr.idx, rr.peek(),
	)
}

func scanNewline(rr *runeReader) (*lexeme.Lexeme, error) {

	rr.accept('\r')

	e := rr.expect('\n')
	if e != nil {
		return nil, e
	}

	return rr.slice(prop.PR_REDUNDANT, prop.PR_NEWLINE), nil
}
