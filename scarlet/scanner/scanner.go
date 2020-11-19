// Scanner package scans Lexemes (Tokens) from a text source. The scanner will
// not sanitise any text in the process so the resultant lexemes will be an
// exact representation of the input including whitespace and other redundant
// Tokens.
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

		rng, val := r.read(l.size)
		tk := token.Make(val, l.tk, token.Snippet{})
		tk.Range = rng

		if r.more() {
			return tk, scan(r), nil
		}

		return tk, nil, nil
	}
}

func identifyLexeme(r *reader, l *lex) error {

	switch {
	case r.at(0) == '\n':
		l.size, l.tk = 1, token.NEWLINE
	case r.starts("\r\n"):
		l.size, l.tk = 2, token.NEWLINE
	case r.at(0) == '\r':
		return err(r.tm.PosAfter("\r"), "Missing LF after CR")

	case unicode.IsSpace(r.at(0)):
		l.size, l.tk = 1, token.SPACE
		for r.inRange(l.size) && unicode.IsSpace(r.at(l.size)) &&
			r.at(l.size) != '\r' && r.at(l.size) != '\n' {
			l.size++
		}

	case r.at(0) == '#':
		l.size, l.tk = 1, token.COMMENT
		for r.inRange(l.size) && r.at(l.size) != '\n' && !r.starts("\r\n") {
			l.size++
		}

	case unicode.IsLetter(r.at(0)):
		identifyWord(r, l)

	case r.at(0) == ';':
		l.size, l.tk = 1, token.TERMINATOR

	case r.starts("<-"):
		l.size, l.tk = 2, token.ASSIGN

	case r.at(0) == ',':
		l.size, l.tk = 1, token.DELIM

	case r.at(0) == '?':
		l.size, l.tk = 1, token.EXIST

	case r.at(0) == '(':
		l.size, l.tk = 1, token.L_PAREN

	case r.at(0) == ')':
		l.size, l.tk = 1, token.R_PAREN

	case r.at(0) == '[':
		l.size, l.tk = 1, token.L_SQUARE

	case r.at(0) == ']':
		l.size, l.tk = 1, token.R_SQUARE

	case r.at(0) == '{':
		l.size, l.tk = 1, token.L_CURLY

	case r.at(0) == '}':
		l.size, l.tk = 1, token.R_CURLY

	case r.at(0) == '_':
		l.size, l.tk = 1, token.VOID

	case r.at(0) == '+':
		l.size, l.tk = 1, token.ADD

	case r.at(0) == '-':
		l.size, l.tk = 1, token.SUB

	case r.at(0) == '*':
		l.size, l.tk = 1, token.MUL

	case r.at(0) == '/':
		l.size, l.tk = 1, token.DIV

	case r.at(0) == '%':
		l.size, l.tk = 1, token.REM

	case r.starts("&&"):
		l.size, l.tk = 2, token.AND

	case r.starts("||"):
		l.size, l.tk = 2, token.OR

	case r.starts("<="):
		l.size, l.tk = 2, token.LTE

	case r.at(0) == '<':
		l.size, l.tk = 1, token.LT

	case r.starts(">="):
		l.size, l.tk = 2, token.MTE

	case r.at(0) == '>':
		l.size, l.tk = 1, token.MT

	case r.starts("=="):
		l.size, l.tk = 2, token.EQU

	case r.starts("!="):
		l.size, l.tk = 2, token.NEQ

	case r.at(0) == '@':
		if e := spell(r, l); e != nil {
			return e
		}

	case r.at(0) == '"':
		if e := stringLiteral(r, l); e != nil {
			return e
		}

	case unicode.IsDigit(r.at(0)):
		if e := numberLiteral(r, l); e != nil {
			return e
		}

	default:
		return err(r.tm.PosAfter(""), "Unexpected symbol %q", r.at(0))
	}

	return nil
}

func identifyWord(r *reader, l *lex) {

	l.size = 1
	for r.inRange(l.size) {
		if ru := r.at(l.size); !unicode.IsLetter(ru) && ru != '_' {
			break
		}

		l.size++
	}

	l.tk = token.IdentifyWord(r.slice(l.size))
}

func spell(r *reader, l *lex) error {

	// Valid:   @abc
	// Valid:   @abc.xyz
	// Invalid: @
	// Invalid: @abc.
	// Invalid: @abc..xyz
	// Invalid: @abc.efg.hij

	parsePart := func() error {
		if !r.inRange(l.size) {
			// TODO: Snippet here, `colRune + l.size`
			return err(r.tm.PosAfter(""), "Bad spell name, have EOF, want letter")
		}

		if ru := r.at(l.size); !unicode.IsLetter(ru) {
			// TODO: Snippet here, `colRune + l.size`
			return err(r.tm.PosAfter(""), "Bad spell name, have %q, want letter", ru)
		}

		l.size++
		for r.inRange(l.size) && unicode.IsLetter(r.at(l.size)) {
			l.size++
		}

		return nil
	}

	l.size, l.tk = 1, token.SPELL

	if e := parsePart(); e != nil {
		return e
	}

	if !r.inRange(l.size) || r.at(l.size) != '.' {
		return nil
	}
	l.size++

	if e := parsePart(); e != nil {
		return e
	}

	return nil
}

func stringLiteral(r *reader, l *lex) error {

	l.size, l.tk = 1, token.STRING
	for {
		if !r.inRange(l.size) {
			goto ERROR
		}

		if r.at(l.size) == '"' {
			l.size++
			return nil
		}

		if r.at(l.size) == '\\' {
			l.size++
			if !r.inRange(l.size) {
				goto ERROR
			}
		}

		if ru := r.at(l.size); ru == '\r' || ru == '\n' {
			// TODO: Snippet here, `colRune + l.size`
			return err(r.tm.PosAfter(""), "Unterminated string")
		}
		l.size++
	}

ERROR:
	// TODO: Snippet here, `colRune + l.size`
	return err(r.tm.PosAfter(""), "Unterminated string")
}

func numberLiteral(r *reader, l *lex) error {

	l.size, l.tk = 1, token.NUMBER
	for r.inRange(l.size) && unicode.IsDigit(r.at(l.size)) {
		l.size++
	}

	if !r.inRange(l.size) || r.at(l.size) != '.' {
		return nil
	}
	l.size++

	if !r.inRange(l.size) {
		// TODO: Snippet here, `colRune + l.size`
		return err(r.tm.PosAfter(""), "Unexpected symbol, have EOF, want [0-9]")
	}

	if ru := r.at(l.size); !unicode.IsDigit(ru) {
		// TODO: Snippet here, `colRune + l.size`
		return err(r.tm.PosAfter(""), "Unexpected symbol, have %q want [0-9]", ru)
	}

	for r.inRange(l.size) && unicode.IsDigit(r.at(l.size)) {
		l.size++
	}

	return nil
}
