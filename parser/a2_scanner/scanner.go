package scanner2

import (
	"github.com/PaulioRandall/scarlet-go/token/container"
	"github.com/PaulioRandall/scarlet-go/token/lexeme"
)

type token struct {
	line int
	col  int
	size int
	typ  lexeme.TokenType
}

func ScanString(s string) (*container.Container, error) {

	con := container.New()
	r := &reader{}
	r.data = []rune(s)
	r.size = len(r.data)

	for r.more() {
		tk := &token{
			line: r.line,
			col:  r.col,
		}

		if e := identifyToken(r, tk); e != nil {
			return nil, e
		}

		if e := scanToken(con, r, tk); e != nil {
			return nil, e
		}
	}

	return con, nil
}

func identifyToken(r *reader, tk *token) error {

	switch {
	case r.starts("\n"):
		tk.size, tk.typ = 1, lexeme.NEWLINE
	case r.starts("\r\n"):
		tk.size, tk.typ = 2, lexeme.NEWLINE

	default:
		return newErr(r.line, r.idx, "Unexpected symbol %q", r.at(0))
	}

	return nil
}

func isNewline(offset int, r *reader) bool {
	return false
}

func scanToken(con *container.Container, r *reader, tk *token) error {
	return nil
}
