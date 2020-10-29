// Scanner package scans Lexemes (Tokens) from a text source. The scanner will
// not sanitise any text in the process so the resultant lexemes will be an
// exact representation of the input including whitespace and other redundant
// Tokens. Pre-parsing should be performed via the sanitiser package if the
// Tokens are heading for compilation.
package scanner

import (
	"unicode"

	"github.com/PaulioRandall/scarlet-go/scarlet/token"
)

type (
	// ParseToken is designed to be used in a recursice fashion. It returns a
	// lexeme and another ParseToken function to obtain the subsequent lexeme.
	// On the last lexeme the ParseToken will be nil. Parsing may also end if an
	// error is returned.
	ParseToken func() (token.Lexeme, ParseToken, error)

	lex struct {
		size int // In runes
		tk   token.Token
	}
)

// New returns a ParseToken function. Calling the function will return a lexeme
// and another ParseToken function to obtain the subsequent token. On the last
// lexeme the ParseToken will be nil. Parsing may also end if an error is
// returned.
func New(in []rune) ParseToken {

	r := &reader{
		data:   in,
		remain: len(in),
	}

	if r.more() {
		return scan(r)
	}
	return nil
}

// ScanAll scans all lexemes within 'in' and returns them as an ordered slice.
func ScanAll(in []rune) ([]token.Lexeme, error) {

	var (
		r  []token.Lexeme
		l  token.Lexeme
		pt = New(in)
		e  error
	)

	for pt != nil {
		if l, pt, e = pt(); e != nil {
			return nil, e
		}
		r = append(r, l)
	}

	return r, nil
}

func scan(r *reader) ParseToken {
	return func() (token.Lexeme, ParseToken, error) {

		l := &lex{}
		if e := identifyLexeme(r, l); e != nil {
			return token.Lexeme{}, nil, e
		}

		snip, val := r.read(l.size)
		tk := token.Make(val, l.tk, snip)

		if r.more() {
			return tk, scan(r), nil
		}
		return tk, nil, nil
	}
}

func identifyLexeme(r *reader, l *lex) error {

	switch {
	case r.at(0) == LF:
		l.size, l.tk = 1, token.NEWLINE
	case r.starts(CRLF):
		l.size, l.tk = 2, token.NEWLINE
	case r.at(0) == CR:
		return errPos(r.Position(string(CR)), "Missing LF after CR")

	case unicode.IsSpace(r.at(0)):
		l.size, l.tk = 1, token.SPACE
		for r.inRange(l.size) && unicode.IsSpace(r.at(l.size)) {
			l.size++
		}

	case r.starts(COMMENT_PREFIX):
		l.size, l.tk = 1, token.COMMENT
		for r.inRange(l.size) && r.at(0) != LF && !r.starts(CRLF) {
			l.size++
		}

	case unicode.IsLetter(r.at(0)):
		identifyWord(r, l)

	case r.starts(TERMINATOR):
		l.size, l.tk = 1, token.TERMINATOR

	case r.starts(ASSIGN):
		l.size, l.tk = 2, token.ASSIGN

	case r.starts(DELIM):
		l.size, l.tk = 1, token.DELIM

	case r.starts(L_PAREN):
		l.size, l.tk = 1, token.L_PAREN

	case r.starts(R_PAREN):
		l.size, l.tk = 1, token.R_PAREN

	case r.starts(L_SQUARE):
		l.size, l.tk = 1, token.L_SQUARE

	case r.starts(R_SQUARE):
		l.size, l.tk = 1, token.R_SQUARE

	case r.starts(L_CURLY):
		l.size, l.tk = 1, token.L_CURLY

	case r.starts(R_CURLY):
		l.size, l.tk = 1, token.R_CURLY

	case r.at(0) == VOID:
		l.size, l.tk = 1, token.VOID

	case r.starts(ADD):
		l.size, l.tk = 1, token.ADD

	case r.starts(SUB):
		l.size, l.tk = 1, token.SUB

	case r.starts(MUL):
		l.size, l.tk = 1, token.MUL

	case r.starts(DIV):
		l.size, l.tk = 1, token.DIV

	case r.starts(REM):
		l.size, l.tk = 1, token.REM

	case r.starts(AND):
		l.size, l.tk = 2, token.AND

	case r.starts(OR):
		l.size, l.tk = 2, token.OR

	case r.starts(LESS_EQUAL):
		l.size, l.tk = 2, token.LESS_EQUAL

	case r.starts(LESS):
		l.size, l.tk = 1, token.LESS

	case r.starts(MORE_EQUAL):
		l.size, l.tk = 2, token.MORE_EQUAL

	case r.starts(MORE):
		l.size, l.tk = 1, token.MORE

	case r.starts(EQUAL):
		l.size, l.tk = 2, token.EQUAL

	case r.starts(NOT_EQUAL):
		l.size, l.tk = 2, token.NOT_EQUAL

	case r.starts(SPELL_PREFIX):
		if e := spell(r, l); e != nil {
			return e
		}

	case r.at(0) == STRING_PREFIX:
		if e := stringLiteral(r, l); e != nil {
			return e
		}

	case unicode.IsDigit(r.at(0)):
		if e := numberLiteral(r, l); e != nil {
			return e
		}

	default:
		return errPos(r.Snapshot(), "Unexpected symbol %q", r.at(0))
	}

	return nil
}

func identifyWord(r *reader, l *lex) {

	l.size = 1
	for r.inRange(l.size) {
		if ru := r.at(l.size); !unicode.IsLetter(ru) && ru != VOID {
			break
		}

		l.size++
	}

	l.tk = token.IdentifyWord(r.slice(l.size))
}

func spell(r *reader, l *lex) error {

	// Valid:   @abc
	// Valid:   @abc.efg.xyz
	// Invalid: @
	// Invalid: @abc.

	part := func() error {
		if !r.inRange(l.size) {
			// TODO: Snippet here, `colRune + l.size`
			return errPos(r.Snapshot(), "Bad spell name, have EOF, want letter")
		}

		if ru := r.at(l.size); !unicode.IsLetter(ru) {
			// TODO: Snippet here, `colRune + l.size`
			return errPos(r.Snapshot(), "Bad spell name, have %q, want letter", ru)
		}

		l.size++
		for r.inRange(l.size) && unicode.IsLetter(r.at(l.size)) {
			l.size++
		}

		return nil
	}

	l.size, l.tk = 1, token.SPELL
	for {
		if e := part(); e != nil {
			return e
		}

		if !r.inRange(l.size) || r.at(l.size) != SPELL_NAME_DELIM {
			break
		}
		l.size++
	}

	return nil
}

func stringLiteral(r *reader, l *lex) error {

	l.size, l.tk = 1, token.STRING
	for {
		if !r.inRange(l.size) {
			goto ERROR
		}

		if r.at(l.size) == STRING_SUFFIX {
			l.size++
			return nil
		}

		if r.at(l.size) == STRING_ESCAPE {
			l.size++
			if !r.inRange(l.size) {
				goto ERROR
			}
		}

		if ru := r.at(l.size); ru == CR || ru == LF {
			// TODO: Snippet here, `colRune + l.size`
			return errPos(r.Snapshot(), "Unterminated string")
		}
		l.size++
	}

ERROR:
	// TODO: Snippet here, `colRune + l.size`
	return errPos(r.Snapshot(), "Unterminated string")
}

func numberLiteral(r *reader, l *lex) error {

	l.size, l.tk = 1, token.NUMBER
	for r.inRange(l.size) && unicode.IsDigit(r.at(l.size)) {
		l.size++
	}

	if !r.inRange(l.size) || r.at(l.size) != NUMBER_FRAC_DELIM {
		return nil
	}
	l.size++

	if !r.inRange(l.size) {
		// TODO: Snippet here, `colRune + l.size`
		return errPos(r.Snapshot(), "Unexpected symbol, have EOF, want [0-9]")
	}

	if ru := r.at(l.size); !unicode.IsDigit(ru) {
		// TODO: Snippet here, `colRune + l.size`
		return errPos(r.Snapshot(), "Unexpected symbol, have %q want [0-9]", ru)
	}

	for r.inRange(l.size) && unicode.IsDigit(r.at(l.size)) {
		l.size++
	}

	return nil
}
