package scanner

import (
	"github.com/PaulioRandall/scarlet-go/lexeme"
	"github.com/PaulioRandall/scarlet-go/perror"
)

type Queue interface {
	AsContainer() *lexeme.Container
	Put(*lexeme.Lexeme)
}

func ScanStr(s string) (*lexeme.Container, error) {

	rr := &runeReader{}
	rr.runes = []rune(s)
	rr.size = len(rr.runes)

	que := Queue(&lexeme.Container{})

	for rr.more() {

		lex, e := scanLexeme(rr)
		if e != nil {
			return nil, e
		}

		que.Put(lex)
	}

	return que.AsContainer(), nil
}

func scanLexeme(rr *runeReader) (*lexeme.Lexeme, error) {

	switch {
	case rr.isNewline():
		return newline(rr)

	case rr.is('#'):
		return comment(rr)

	case rr.isSpace():
		return whitespace(rr)

	case rr.isLetter():
		return word(rr)

	case rr.is('@'):
		return spell(rr)

	case rr.is('"'):
		return stringLiteral(rr)

	case rr.isDigit():
		return numberLiteral(rr)

	case rr.accept('('):
		return rr.slice(lexeme.L_PAREN), nil

	case rr.accept(')'):
		return rr.slice(lexeme.R_PAREN), nil

	case rr.accept('['):
		return rr.slice(lexeme.L_SQUARE), nil

	case rr.accept(']'):
		return rr.slice(lexeme.R_SQUARE), nil

	case rr.accept('{'):
		return rr.slice(lexeme.L_CURLY), nil

	case rr.accept('}'):
		return rr.slice(lexeme.R_CURLY), nil

	case rr.accept(','):
		return rr.slice(lexeme.DELIM), nil

	case rr.accept(':'):
		if e := rr.expect('='); e != nil {
			return nil, e
		}
		return rr.slice(lexeme.ASSIGN), nil

	case rr.accept('_'):
		return rr.slice(lexeme.VOID), nil

	case rr.accept('+'):
		return rr.slice(lexeme.ADD), nil

	case rr.accept('-'):
		return rr.slice(lexeme.SUB), nil

	case rr.accept('*'):
		return rr.slice(lexeme.MUL), nil

	case rr.accept('/'):
		return rr.slice(lexeme.DIV), nil

	case rr.accept('%'):
		return rr.slice(lexeme.REM), nil

	case rr.accept('&'):
		if e := rr.expect('&'); e != nil {
			return nil, e
		}
		return rr.slice(lexeme.AND), nil

	case rr.accept('|'):
		if e := rr.expect('|'); e != nil {
			return nil, e
		}
		return rr.slice(lexeme.OR), nil

	case rr.accept('<'):
		if rr.accept('=') {
			return rr.slice(lexeme.LESS_EQUAL), nil
		}
		return rr.slice(lexeme.LESS), nil

	case rr.accept('>'):
		if rr.accept('=') {
			return rr.slice(lexeme.MORE_EQUAL), nil
		}
		return rr.slice(lexeme.MORE), nil

	case rr.accept('='):
		if e := rr.expect('='); e != nil {
			return nil, e
		}
		return rr.slice(lexeme.EQUAL), nil

	case rr.accept('!'):
		if e := rr.expect('='); e != nil {
			return nil, e
		}
		return rr.slice(lexeme.NOT_EQUAL), nil
	}

	return nil, perror.New(
		"Unexpected terminal symbol %d:%d, have %q",
		rr.line, rr.idx, rr.peek(),
	)
}

func newline(rr *runeReader) (*lexeme.Lexeme, error) {

	rr.accept('\r')

	e := rr.expect('\n')
	if e != nil {
		return nil, e
	}

	return rr.slice(lexeme.NEWLINE), nil
}

func comment(rr *runeReader) (*lexeme.Lexeme, error) {

	for rr.more() && !rr.isNewline() {
		rr.inc()
	}

	return rr.slice(lexeme.COMMENT), nil
}

func whitespace(rr *runeReader) (*lexeme.Lexeme, error) {

	for rr.isSpace() && !rr.isNewline() {
		rr.inc()
	}

	return rr.slice(lexeme.SPACE), nil
}

func word(rr *runeReader) (*lexeme.Lexeme, error) {

	rr.inc()

	for rr.isLetter() || rr.is('_') {
		rr.inc()
	}

	lex := rr.slice(lexeme.UNDEFINED)

	switch lex.Raw {
	case "false", "true":
		lex.Tok = lexeme.BOOL

	case "loop":
		lex.Tok = lexeme.LOOP

	default:
		lex.Tok = lexeme.IDENT
	}

	return lex, nil
}

func spell(rr *runeReader) (*lexeme.Lexeme, error) {

	rr.inc()

	for {
		if rr.empty() {
			return nil, perror.New(
				"Bad spell name %d:%d, want: letter, have: EOF",
				rr.line, rr.idx,
			)
		}

		if !rr.isLetter() {
			return nil, perror.New(
				"Bad spell name %d:%d, want: letter, have: %q",
				rr.line, rr.idx, rr.peek(),
			)
		}

		for rr.isLetter() {
			rr.inc()
		}

		if !rr.is('.') {
			break
		}

		rr.inc()
	}

	return rr.slice(lexeme.SPELL), nil
}

func stringLiteral(rr *runeReader) (*lexeme.Lexeme, error) {

	rr.inc()

	for !rr.accept('"') {

		rr.accept('\\')
		if rr.empty() || rr.isNewline() {
			return nil, perror.New("Unterminated string %d:%d", rr.line, rr.idx)
		}

		rr.inc()
	}

	return rr.slice(lexeme.STRING), nil
}

func numberLiteral(rr *runeReader) (*lexeme.Lexeme, error) {

	for rr.isDigit() {
		rr.inc()
	}

	if !rr.accept('.') {
		goto FINALISE
	}

	if rr.empty() {
		return nil, perror.New(
			"Unexpected symbol %d:%d, want: [0-9], have: EOF",
			rr.line, rr.idx,
		)
	}

	if !rr.isDigit() {
		return nil, perror.New(
			"Unexpected symbol %d:%d, want: [0-9], have: %q",
			rr.line, rr.idx, rr.peek(),
		)
	}

	for rr.isDigit() {
		rr.inc()
	}

FINALISE:
	return rr.slice(lexeme.NUMBER), nil
}
