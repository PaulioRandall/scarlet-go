// Scanner package is scans in Lexemes (Tokens) from a text source into a
// Series. The scanner is will not sanitise any text in the process so the
// resultant Series of Lexemes will be an exact representation of the input
// source code including whitespace and other redundant Tokens. Pre-parsing
// should be performed via the sanitiser module if the Tokens are heading for
// compilation.
package scanner

import (
	"unicode"

	"github.com/PaulioRandall/scarlet-go/token/lexeme"
	"github.com/PaulioRandall/scarlet-go/token/series"
	"github.com/PaulioRandall/scarlet-go/token/symbol"
	"github.com/PaulioRandall/scarlet-go/token/token"
)

type lex struct {
	size int // In runes
	tk   token.Token
}

// ScanAll converts the input 'in' into a Series of Lexemes (Tokens).
// Redundant Tokens are not removed in the process so the result will be a
// lossless representation of the original input 'in'.
func ScanAll(in []rune) (*series.Series, error) {

	se := series.Make()
	r := &reader{
		data:   in,
		remain: len(in),
	}

	for r.more() {
		l := &lex{}
		if e := identifyLexeme(r, l); e != nil {
			return nil, e
		}

		snip, val := r.read(l.size)
		tk := lexeme.Make(val, l.tk, snip)
		se.Append(tk)
	}

	return se, nil
}

func identifyLexeme(r *reader, l *lex) error {

	switch {
	case r.at(0) == symbol.LF:
		l.size, l.tk = 1, token.NEWLINE
	case r.starts(symbol.CRLF):
		l.size, l.tk = 2, token.NEWLINE
	case r.at(0) == symbol.CR:
		return errPos(r.Position(string(symbol.CR)), "Missing LF after CR")

	case unicode.IsSpace(r.at(0)):
		l.size, l.tk = 1, token.SPACE
		for r.inRange(l.size) && unicode.IsSpace(r.at(l.size)) {
			l.size++
		}

	case r.starts(symbol.COMMENT_PREFIX):
		l.size, l.tk = 1, token.COMMENT
		for r.inRange(l.size) && r.at(0) != symbol.LF && !r.starts(symbol.CRLF) {
			l.size++
		}

	case unicode.IsLetter(r.at(0)):
		identifyWord(r, l)

	case r.starts(symbol.TERMINATOR):
		l.size, l.tk = 1, token.TERMINATOR

	case r.starts(symbol.ASSIGN):
		l.size, l.tk = 2, token.ASSIGN

	case r.starts(symbol.DELIM):
		l.size, l.tk = 1, token.DELIM

	case r.starts(symbol.L_PAREN):
		l.size, l.tk = 1, token.L_PAREN

	case r.starts(symbol.R_PAREN):
		l.size, l.tk = 1, token.R_PAREN

	case r.starts(symbol.L_SQUARE):
		l.size, l.tk = 1, token.L_SQUARE

	case r.starts(symbol.R_SQUARE):
		l.size, l.tk = 1, token.R_SQUARE

	case r.starts(symbol.L_CURLY):
		l.size, l.tk = 1, token.L_CURLY

	case r.starts(symbol.R_CURLY):
		l.size, l.tk = 1, token.R_CURLY

	case r.at(0) == symbol.VOID:
		l.size, l.tk = 1, token.VOID

	case r.starts(symbol.ADD):
		l.size, l.tk = 1, token.ADD

	case r.starts(symbol.SUB):
		l.size, l.tk = 1, token.SUB

	case r.starts(symbol.MUL):
		l.size, l.tk = 1, token.MUL

	case r.starts(symbol.DIV):
		l.size, l.tk = 1, token.DIV

	case r.starts(symbol.REM):
		l.size, l.tk = 1, token.REM

	case r.starts(symbol.AND):
		l.size, l.tk = 2, token.AND

	case r.starts(symbol.OR):
		l.size, l.tk = 2, token.OR

	case r.starts(symbol.LESS_EQUAL):
		l.size, l.tk = 2, token.LESS_EQUAL

	case r.starts(symbol.LESS):
		l.size, l.tk = 1, token.LESS

	case r.starts(symbol.MORE_EQUAL):
		l.size, l.tk = 2, token.MORE_EQUAL

	case r.starts(symbol.MORE):
		l.size, l.tk = 1, token.MORE

	case r.starts(symbol.EQUAL):
		l.size, l.tk = 2, token.EQUAL

	case r.starts(symbol.NOT_EQUAL):
		l.size, l.tk = 2, token.NOT_EQUAL

	case r.starts(symbol.SPELL_PREFIX):
		if e := spell(r, l); e != nil {
			return e
		}

	case r.at(0) == symbol.STRING_PREFIX:
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
		if ru := r.at(l.size); !unicode.IsLetter(ru) && ru != symbol.VOID {
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

		if !r.inRange(l.size) || r.at(l.size) != symbol.SPELL_NAME_DELIM {
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

		if r.at(l.size) == symbol.STRING_SUFFIX {
			l.size++
			return nil
		}

		if r.at(l.size) == symbol.STRING_ESCAPE {
			l.size++
			if !r.inRange(l.size) {
				goto ERROR
			}
		}

		if ru := r.at(l.size); ru == symbol.CR || ru == symbol.LF {
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

	if !r.inRange(l.size) || r.at(l.size) != symbol.NUMBER_FRAC_DELIM {
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
