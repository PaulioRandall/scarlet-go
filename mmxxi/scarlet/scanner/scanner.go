package scanner

import (
	"errors"
	"fmt"
	"unicode"

	"github.com/PaulioRandall/scarlet-go/mmxxi/scarlet/scroll"
	"github.com/PaulioRandall/scarlet-go/mmxxi/scarlet/token"
)

type (
	// ParseToken is a recursion based tokeniser. It returns a lexeme and another
	// ParseToken function to obtain the following lexeme. On error or while
	// obtaining the last lexeme, ParseToken will be nil.
	ParseToken func() (token.Lexeme, ParseToken, error)

	lex struct {
		size int // In runes
		tk   token.Token
	}
)

// New returns a ParseToken function.
func New(sr *scroll.ScrollReader) ParseToken {
	if sr.More() {
		return scan(sr)
	}
	return nil
}

// ScanAll scans all remaining lexemes as an ordered slice.
func ScanAll(sr *scroll.ScrollReader) ([]token.Lexeme, error) {

	var (
		r  []token.Lexeme
		lx token.Lexeme
		pt = New(sr)
		e  error
	)

	for pt != nil {
		if lx, pt, e = pt(); e != nil {
			return nil, e
		}
		r = append(r, lx)
	}

	return r, nil
}

func scan(sr *scroll.ScrollReader) ParseToken {
	return func() (token.Lexeme, ParseToken, error) {

		l := &lex{}
		if e := identifyLexeme(sr, l); e != nil {
			return token.Lexeme{}, nil, e
		}

		sn := sr.Read(l.size)
		lx := token.MakeLex(l.tk, sn)

		if sr.More() {
			return lx, scan(sr), nil
		}

		return lx, nil, nil
	}
}

func identifyLexeme(sr *scroll.ScrollReader, l *lex) error {

	switch {
	case sr.Starts(";") || sr.Starts("\n"):
		l.size, l.tk = 1, token.TERMINATOR
	case sr.Starts("\r\n"):
		l.size, l.tk = 2, token.TERMINATOR
	case sr.At(0) == '\r':
		return err(sr, "Missing LF after CR")

	// Redundant
	case unicode.IsSpace(sr.At(0)):
		l.size, l.tk = 1, token.SPACE
		for sr.InRange(l.size) &&
			unicode.IsSpace(sr.At(l.size)) &&
			sr.At(l.size) != '\r' &&
			sr.At(l.size) != '\n' {
			l.size++
		}

	case sr.At(0) == '#':
		l.size, l.tk = 1, token.COMMENT
		for sr.InRange(l.size) && sr.At(l.size) != '\n' && !sr.Starts("\r\n") {
			l.size++
		}

		// Keywords, identifiers, & bool literals
	case unicode.IsLetter(sr.At(0)) || sr.At(0) == '_':
		identifyWord(sr, l)

		// Operators
	case sr.Starts(":="):
		l.size, l.tk = 2, token.DEFINE

	case sr.Starts("<-"):
		l.size, l.tk = 2, token.ASSIGN

	case sr.Starts("->"):
		l.size, l.tk = 2, token.INTO

	case sr.Starts("+"):
		l.size, l.tk = 1, token.ADD

	case sr.Starts("-"):
		l.size, l.tk = 1, token.SUB

	case sr.Starts("*"):
		l.size, l.tk = 1, token.MUL

	case sr.Starts("/"):
		l.size, l.tk = 1, token.DIV

	case sr.Starts("%"):
		l.size, l.tk = 1, token.REM

	case sr.Starts("&&"):
		l.size, l.tk = 2, token.AND

	case sr.Starts("||"):
		l.size, l.tk = 2, token.OR

	case sr.Starts("=="):
		l.size, l.tk = 2, token.EQU

	case sr.Starts("!="):
		l.size, l.tk = 2, token.NEQ

	case sr.Starts("<="):
		l.size, l.tk = 2, token.LTE

	case sr.Starts("<"):
		l.size, l.tk = 1, token.LT

	case sr.Starts(">="):
		l.size, l.tk = 2, token.MTE

	case sr.Starts(">"):
		l.size, l.tk = 1, token.MT

	case sr.Starts("!"):
		l.size, l.tk = 1, token.NOT

	case sr.Starts("?"):
		l.size, l.tk = 1, token.QUE

	// Delimiters
	case sr.Starts("@"):
		l.size, l.tk = 1, token.SPELL

	case sr.Starts(","):
		l.size, l.tk = 1, token.DELIM

	case sr.Starts(":"):
		l.size, l.tk = 1, token.REF

	case sr.Starts("("):
		l.size, l.tk = 1, token.L_PAREN

	case sr.Starts(")"):
		l.size, l.tk = 1, token.R_PAREN

	case sr.Starts("["):
		l.size, l.tk = 1, token.L_BRACK

	case sr.Starts("]"):
		l.size, l.tk = 1, token.R_BRACK

	case sr.Starts("{"):
		l.size, l.tk = 1, token.L_BRACE

	case sr.Starts("}"):
		l.size, l.tk = 1, token.R_BRACE

	// Remaining literals
	case sr.Starts(`"`):
		if e := stringLiteral(sr, l); e != nil {
			return e
		}

	case unicode.IsDigit(sr.At(0)):
		if e := numberLiteral(sr, l); e != nil {
			return e
		}

	default:
		return err(sr, "Unexpected symbol %q", sr.At(0))
	}

	return nil
}

func identifyWord(sr *scroll.ScrollReader, l *lex) {

	l.size = 1
	for sr.InRange(l.size) {
		ru := sr.At(l.size)

		if !unicode.IsLetter(ru) &&
			!unicode.IsDigit(ru) &&
			ru != '_' {
			break
		}

		l.size++
	}

	l.tk = token.IdentifyWord(sr.Slice(l.size))
}

func stringLiteral(sr *scroll.ScrollReader, l *lex) error {

	l.size, l.tk = 1, token.STR
	for {
		if !sr.InRange(l.size) {
			goto ERROR
		}

		if sr.At(l.size) == '"' {
			l.size++
			return nil
		}

		if sr.At(l.size) == '\\' {
			l.size++
			if !sr.InRange(l.size) {
				goto ERROR
			}
		}

		if ru := sr.At(l.size); ru == '\r' || ru == '\n' {
			return err(sr, "Unterminated string")
		}
		l.size++
	}

ERROR:
	return err(sr, "Unterminated string")
}

func numberLiteral(sr *scroll.ScrollReader, l *lex) error {

	l.size, l.tk = 1, token.NUM
	for sr.InRange(l.size) && unicode.IsDigit(sr.At(l.size)) {
		l.size++
	}

	if !sr.InRange(l.size) || sr.At(l.size) != '.' {
		return nil
	}
	l.size++

	if !sr.InRange(l.size) {
		return err(sr, "Unexpected symbol, have EOF, want [0-9]")
	}

	if ru := sr.At(l.size); !unicode.IsDigit(ru) {
		return err(sr, "Unexpected symbol, have %q want [0-9]", ru)
	}

	for sr.InRange(l.size) && unicode.IsDigit(sr.At(l.size)) {
		l.size++
	}

	return nil
}

func err(sr *scroll.ScrollReader, m string, args ...interface{}) error {
	m = fmt.Sprintf(m, args...)
	m = fmt.Sprintf("Line %d: %s", sr.Line(), m)
	return errors.New(m)
}
