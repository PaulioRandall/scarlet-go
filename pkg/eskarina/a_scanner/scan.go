package scanner

import (
	"github.com/PaulioRandall/scarlet-go/pkg/eskarina/lexeme"
	"github.com/PaulioRandall/scarlet-go/pkg/eskarina/perror"
	"github.com/PaulioRandall/scarlet-go/pkg/eskarina/prop"
)

func ScanStr(s string) (*lexeme.Lexeme, error) {

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

	switch {
	case lr.isNewline():
		return newline(lr)
	case lr.is('#'):
		return comment(lr)
	case lr.isSpace():
		return whitespace(lr)
	}

	return nil, perror.New(
		"Unexpected terminal symbol %d:%d, have %q",
		lr.line, lr.idx, lr.peek(),
	)
}

func newline(lr *lexReader) (*lexeme.Lexeme, error) {

	lr.accept('\r')

	e := lr.expect('\n')
	if e != nil {
		return nil, e
	}

	return lr.slice(prop.PR_TERMINATOR, prop.PR_NEWLINE), nil
}

func comment(lr *lexReader) (*lexeme.Lexeme, error) {

	for lr.more() && !lr.isNewline() {
		lr.inc()
	}

	return lr.slice(prop.PR_REDUNDANT, prop.PR_COMMENT), nil
}

func whitespace(lr *lexReader) (*lexeme.Lexeme, error) {

	for lr.more() && lr.isSpace() && !lr.isNewline() {
		lr.inc()
	}

	return lr.slice(prop.PR_REDUNDANT, prop.PR_WHITESPACE), nil
}
