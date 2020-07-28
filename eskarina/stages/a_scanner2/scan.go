package scanner

import (
	"github.com/PaulioRandall/scarlet-go/eskarina/shared/lexeme"
	"github.com/PaulioRandall/scarlet-go/eskarina/shared/perror"
)

func ScanStr(s string) (*lexeme.Container2, error) {

	rr := &runeReader{}
	rr.runes = []rune(s)
	rr.size = len(rr.runes)

	lr := &lexReader{
		runeReader: rr,
	}

	que := &lexeme.Container2{}

	for lr.more() {

		lex, e := scanLexeme(lr)
		if e != nil {
			return nil, e
		}

		que.Put(lex)
	}

	return que, nil
}

func scanLexeme(lr *lexReader) (*lexeme.Lexeme, error) {

	switch {
	case lr.isNewline():
		return newline(lr)

	case lr.is('#'):
		return comment(lr)

	case lr.isSpace():
		return whitespace(lr)

	case lr.isLetter():
		return word(lr)

	case lr.is('@'):
		return spell(lr)

	case lr.is('"'):
		return stringLiteral(lr)

	case lr.isDigit():
		return numberLiteral(lr)

	case lr.accept('('):
		return lr.slice(lexeme.LEFT_PAREN), nil

	case lr.accept(')'):
		return lr.slice(lexeme.RIGHT_PAREN), nil

	case lr.accept(','):
		return lr.slice(lexeme.SEPARATOR), nil
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

	return lr.slice(lexeme.NEWLINE), nil
}

func comment(lr *lexReader) (*lexeme.Lexeme, error) {

	for lr.more() && !lr.isNewline() {
		lr.inc()
	}

	return lr.slice(lexeme.COMMENT), nil
}

func whitespace(lr *lexReader) (*lexeme.Lexeme, error) {

	for lr.isSpace() && !lr.isNewline() {
		lr.inc()
	}

	return lr.slice(lexeme.WHITESPACE), nil
}

func word(lr *lexReader) (*lexeme.Lexeme, error) {

	lr.inc()

	for lr.isLetter() || lr.is('_') {
		lr.inc()
	}

	lex := lr.slice(lexeme.UNDEFINED)

	if lex.Raw == "false" || lex.Raw == "true" {
		lex.Tok = lexeme.BOOL
	} else {
		lex.Tok = lexeme.IDENTIFIER
	}

	return lex, nil
}

func spell(lr *lexReader) (*lexeme.Lexeme, error) {

	lr.inc()

	for {
		if lr.empty() {
			return nil, perror.New(
				"Bad spell name %d:%d, want: letter, have: EOF",
				lr.line, lr.idx,
			)
		}

		if !lr.isLetter() {
			return nil, perror.New(
				"Bad spell name %d:%d, want: letter, have: %q",
				lr.line, lr.idx, lr.peek(),
			)
		}

		for lr.isLetter() {
			lr.inc()
		}

		if !lr.is('.') {
			break
		}

		lr.inc()
	}

	return lr.slice(lexeme.SPELL), nil
}

func stringLiteral(lr *lexReader) (*lexeme.Lexeme, error) {

	lr.inc()

	for !lr.accept('"') {

		lr.accept('\\')
		if lr.empty() || lr.isNewline() {
			return nil, perror.New("Unterminated string %d:%d", lr.line, lr.idx)
		}

		lr.inc()
	}

	return lr.slice(lexeme.STRING), nil
}

func numberLiteral(lr *lexReader) (*lexeme.Lexeme, error) {

	for lr.isDigit() {
		lr.inc()
	}

	if !lr.accept('.') {
		goto FINALISE
	}

	if lr.empty() {
		return nil, perror.New(
			"Unexpected symbol %d:%d, want: [0-9], have: EOF",
			lr.line, lr.idx,
		)
	}

	if !lr.isDigit() {
		return nil, perror.New(
			"Unexpected symbol %d:%d, want: [0-9], have: %q",
			lr.line, lr.idx, lr.peek(),
		)
	}

	for lr.isDigit() {
		lr.inc()
	}

FINALISE:
	return lr.slice(lexeme.NUMBER), nil
}