package scanner2

import (
	"unicode"

	"github.com/PaulioRandall/scarlet-go/token/container"
	"github.com/PaulioRandall/scarlet-go/token/lexeme"
)

type token struct {
	size int
	typ  lexeme.TokenType
}

func ScanString(s string) (*container.Container, error) {

	con := container.New()
	r := &reader{}
	r.data = []rune(s)
	r.size = len(r.data)

	for r.more() {
		tk := &token{}
		if e := identifyToken(r, tk); e != nil {
			return nil, e
		}

		line, col, raw := r.read(tk.size)
		l := lexeme.New(raw, tk.typ, line, col)
		con.Put(l)
	}

	return con, nil
}

func identifyToken(r *reader, tk *token) error {

	switch {
	case r.starts("\n"):
		tk.size, tk.typ = 1, lexeme.NEWLINE
	case r.starts("\r\n"):
		tk.size, tk.typ = 2, lexeme.NEWLINE
	case r.starts("\r"):
		return newErr(r.line, r.col, "Missing %q after %q", "\n", "\r")

	case unicode.IsSpace(r.at(0)):
		tk.size, tk.typ = 1, lexeme.SPACE
		for r.inRange(tk.size) && unicode.IsSpace(r.at(tk.size)) {
			tk.size++
		}

	case r.starts("#"):
		tk.size, tk.typ = 1, lexeme.COMMENT
		for r.inRange(tk.size) && !r.starts("\n") && !r.starts("\r") {
			tk.size++
		}

	case unicode.IsLetter(r.at(0)):
		identifyWord(r, tk)

	case r.starts("@"):
		if e := spell(r, tk); e != nil {
			return e
		}

	case r.starts(`"`):
		if e := stringLiteral(r, tk); e != nil {
			return e
		}

	default:
		return newErr(r.line, r.col, "Unexpected symbol %q", r.at(0))
	}

	return nil
}

func identifyWord(r *reader, tk *token) {

	tk.size = 1
	for r.inRange(tk.size) {
		if ru := r.at(tk.size); !unicode.IsLetter(ru) && ru != '_' {
			break
		}

		tk.size++
	}

	switch r.slice(tk.size) {
	case "false", "true":
		tk.typ = lexeme.BOOL
	case "loop":
		tk.typ = lexeme.LOOP
	default:
		tk.typ = lexeme.IDENT
	}
}

func spell(r *reader, tk *token) error {

	// Valid:   @abc
	// Valid:   @abc.efg.xyz
	// Invalid: @
	// Invalid: @abc.

	part := func() error {
		if !r.inRange(tk.size) {
			return newErr(r.line, r.col+tk.size,
				"Bad spell name, have EOF, want letter")
		}

		if ru := r.at(tk.size); !unicode.IsLetter(ru) {
			return newErr(r.line, r.col+tk.size,
				"Bad spell name, have %q, want letter", ru)
		}

		tk.size++
		for r.inRange(tk.size) && unicode.IsLetter(r.at(tk.size)) {
			tk.size++
		}

		return nil
	}

	tk.size, tk.typ = 1, lexeme.SPELL
	for {
		if e := part(); e != nil {
			return e
		}

		if !r.inRange(tk.size) || r.at(tk.size) != '.' {
			break
		}
		tk.size++
	}

	return nil
}

func stringLiteral(r *reader, tk *token) error {

	tk.size, tk.typ = 1, lexeme.STRING
	for {
		if !r.inRange(tk.size) {
			goto ERROR
		}

		if r.at(tk.size) == '"' {
			tk.size++
			return nil
		}

		if r.at(tk.size) == '\\' {
			tk.size++
			if !r.inRange(tk.size) {
				goto ERROR
			}
		}

		if ru := r.at(tk.size); ru == '\r' || ru == '\n' {
			return newErr(r.line, r.col+tk.size, "Unterminated string")
		}
		tk.size++
	}

	return nil

ERROR:
	return newErr(r.line, r.col+tk.size, "Unterminated string")
}
