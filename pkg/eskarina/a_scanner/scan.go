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

	lr := &lexReader{
		runeReader: rr,
	}

	var lex *lexeme.Lexeme

	for lr.more() {

		l, e := scanLexeme(lr)
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

func scanLexeme(lr *lexReader) (*lexeme.Lexeme, error) {

	switch lr.peek() {
	case '\r', '\n':
		return scanNewline(lr)
	}

	return nil, perror.New(
		"Unexpected terminal symbol %d:%d, have %q",
		lr.line, lr.idx, lr.peek(),
	)
}

func scanNewline(lr *lexReader) (*lexeme.Lexeme, error) {

	lr.accept('\r')

	e := lr.expect('\n')
	if e != nil {
		return nil, e
	}

	return lr.slice(prop.PR_REDUNDANT, prop.PR_NEWLINE), nil
}
