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

		line, col, raw := r.slice(tk.size)
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

	case unicode.IsSpace(r.at(0)):
		tk.size, tk.typ = 1, lexeme.SPACE
		for r.inRange(tk.size) && unicode.IsSpace(r.at(tk.size)) {
			tk.size++
		}

	default:
		return newErr(r.line, r.col, "Unexpected symbol %q", r.at(0))
	}

	return nil
}
