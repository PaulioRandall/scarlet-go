package scanner

import (
	"unicode"

	"github.com/PaulioRandall/scarlet-go/token/container"
	"github.com/PaulioRandall/scarlet-go/token2/lexeme"
	"github.com/PaulioRandall/scarlet-go/token2/token"
)

type lex struct {
	size int
	tk   token.Token
}

func ScanString(s string) (*container.Container, error) {

	con := container.New()
	r := &reader{}
	r.data = []rune(s)
	r.size = len(r.data)

	for r.more() {
		l := &lex{}
		if e := identifyLexeme(r, l); e != nil {
			return nil, e
		}

		line, col, val := r.read(l.size, l.tk == token.NEWLINE)
		lexTk := lexeme.New(val, l.tk, line, col)

		_ = lexTk
		//con.Put(lexTk)
	}

	return con, nil
}

func identifyLexeme(r *reader, l *lex) error {

	switch {
	case r.starts("\n"):
		l.size, l.tk = 1, token.NEWLINE
	case r.starts("\r\n"):
		l.size, l.tk = 2, token.NEWLINE
	case r.starts("\r"):
		return newErr(r.line, r.col, "Missing %q after %q", "\n", "\r")

	case unicode.IsSpace(r.at(0)):
		l.size, l.tk = 1, token.SPACE
		for r.inRange(l.size) && unicode.IsSpace(r.at(l.size)) {
			l.size++
		}

	case r.starts("#"):
		l.size, l.tk = 1, token.COMMENT
		for r.inRange(l.size) && !r.starts("\n") && !r.starts("\r") {
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
		return newErr(r.line, r.col, "Unexpected symbol %q", r.at(0))
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

	switch r.slice(l.size) {
	case "true":
		l.tk = token.TRUE
	case "false":
		l.tk = token.FALSE
	case "loop":
		l.tk = token.LOOP
	default:
		l.tk = token.IDENT
	}
}

func spell(r *reader, l *lex) error {

	// Valid:   @abc
	// Valid:   @abc.efg.xyz
	// Invalid: @
	// Invalid: @abc.

	part := func() error {
		if !r.inRange(l.size) {
			return newErr(r.line, r.col+l.size,
				"Bad spell name, have EOF, want letter")
		}

		if ru := r.at(l.size); !unicode.IsLetter(ru) {
			return newErr(r.line, r.col+l.size,
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
			return newErr(r.line, r.col+l.size, "Unterminated string")
		}
		l.size++
	}

	return nil

ERROR:
	return newErr(r.line, r.col+l.size, "Unterminated string")
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
		return newErr(r.line, r.col+l.size,
			"Unexpected symbol, have EOF, want [0-9]")
	}

	if ru := r.at(l.size); !unicode.IsDigit(ru) {
		return newErr(r.line, r.col+l.size,
			"Unexpected symbol, have %q want [0-9]", ru)
	}

	for r.inRange(l.size) && unicode.IsDigit(r.at(l.size)) {
		l.size++
	}

	return nil
}
