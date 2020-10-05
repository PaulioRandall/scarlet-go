// Scanner package is scans in Lexemes (Tokens) from a text source into a
// Series. The scanner is will not sanitise any text in the process so the
// resultant Series of Lexemes will be an exact representation of the input
// source code including whitespace and other redundant Tokens. Pre-parsing
// should be performed via the sanitiser module if the Tokens are heading for
// compilation.
package scanner

import (
	"unicode"

	"github.com/PaulioRandall/scarlet-go/token2/lexeme"
	"github.com/PaulioRandall/scarlet-go/token2/series"
	"github.com/PaulioRandall/scarlet-go/token2/token"
)

type lex struct {
	size int // In runes
	tk   token.Token
}

// ScanString converts the text 's' into a Series of Lexemes (Tokens).
// Redundant Tokens are not removed in the process so the result will be a
// lossless representation of the original input text 's'.
func ScanString(s string) (series.Series, error) {

	se := series.Make()
	r := newReader(s)

	for r.more() {
		l := &lex{}
		if e := identifyLexeme(r, l); e != nil {
			return series.Series{}, e
		}

		snip, val := r.read(l.size)
		lexTk := lexeme.Make(val, l.tk, snip)
		se.Append(lexTk)
	}

	return se, nil
}

func identifyLexeme(r *reader, l *lex) error {

	switch {
	case r.starts("\n"):
		l.size, l.tk = 1, token.NEWLINE
	case r.starts("\r\n"):
		l.size, l.tk = 2, token.NEWLINE
	case r.starts("\r"):
		return newErr(r.Line, r.ColRune, "Missing %q after %q", "\n", "\r")

	case unicode.IsSpace(r.at(0)):
		l.size, l.tk = 1, token.SPACE
		for r.inRange(l.size) && unicode.IsSpace(r.at(l.size)) {
			l.size++
		}

	case r.starts("#"):
		l.size, l.tk = 1, token.COMMENT
		for r.inRange(l.size) && !r.starts("\n") && !r.starts("\r\n") {
			l.size++
		}

	case unicode.IsLetter(r.at(0)):
		identifyWord(r, l)

	case r.starts(";"):
		l.size, l.tk = 1, token.TERMINATOR

	case r.starts(":="):
		l.size, l.tk = 2, token.ASSIGN

	case r.starts(","):
		l.size, l.tk = 1, token.DELIM

	case r.starts("("):
		l.size, l.tk = 1, token.L_PAREN

	case r.starts(")"):
		l.size, l.tk = 1, token.R_PAREN

	case r.starts("["):
		l.size, l.tk = 1, token.L_SQUARE

	case r.starts("]"):
		l.size, l.tk = 1, token.R_SQUARE

	case r.starts("{"):
		l.size, l.tk = 1, token.L_CURLY

	case r.starts("}"):
		l.size, l.tk = 1, token.R_CURLY

	case r.starts("_"):
		l.size, l.tk = 1, token.VOID

	case r.starts("+"):
		l.size, l.tk = 1, token.ADD

	case r.starts("-"):
		l.size, l.tk = 1, token.SUB

	case r.starts("*"):
		l.size, l.tk = 1, token.MUL

	case r.starts("/"):
		l.size, l.tk = 1, token.DIV

	case r.starts("%"):
		l.size, l.tk = 1, token.REM

	case r.starts("&&"):
		l.size, l.tk = 2, token.AND

	case r.starts("||"):
		l.size, l.tk = 2, token.OR

	case r.starts("<="):
		l.size, l.tk = 2, token.LESS_EQUAL

	case r.starts("<"):
		l.size, l.tk = 1, token.LESS

	case r.starts(">="):
		l.size, l.tk = 2, token.MORE_EQUAL

	case r.starts(">"):
		l.size, l.tk = 1, token.MORE

	case r.starts("=="):
		l.size, l.tk = 2, token.EQUAL

	case r.starts("!="):
		l.size, l.tk = 2, token.NOT_EQUAL

	case r.starts("@"):
		if e := spell(r, l); e != nil {
			return e
		}

	case r.starts(`"`):
		if e := stringLiteral(r, l); e != nil {
			return e
		}

	case unicode.IsDigit(r.at(0)):
		if e := numberLiteral(r, l); e != nil {
			return e
		}

	default:
		return newErr(r.Line, r.ColRune, "Unexpected symbol %q", r.at(0))
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
	// Valid:   @abc.efg.xyz
	// Invalid: @
	// Invalid: @abc.

	part := func() error {
		if !r.inRange(l.size) {
			return newErr(r.Line, r.ColRune+l.size,
				"Bad spell name, have EOF, want letter")
		}

		if ru := r.at(l.size); !unicode.IsLetter(ru) {
			return newErr(r.Line, r.ColRune+l.size,
				"Bad spell name, have %q, want letter", ru)
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

		if !r.inRange(l.size) || r.at(l.size) != '.' {
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
			return newErr(r.Line, r.ColRune+l.size, "Unterminated string")
		}
		l.size++
	}

	return nil

ERROR:
	return newErr(r.Line, r.ColRune+l.size, "Unterminated string")
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
		return newErr(r.Line, r.ColRune+l.size,
			"Unexpected symbol, have EOF, want [0-9]")
	}

	if ru := r.at(l.size); !unicode.IsDigit(ru) {
		return newErr(r.Line, r.ColRune+l.size,
			"Unexpected symbol, have %q want [0-9]", ru)
	}

	for r.inRange(l.size) && unicode.IsDigit(r.at(l.size)) {
		l.size++
	}

	return nil
}
