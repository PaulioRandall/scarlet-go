package scanner2

import (
	"github.com/PaulioRandall/scarlet-go/token/container"
)

type token struct {
	line int
	col  int
	size int
}

func ScanString(s string) (*container.Container, error) {

	var (
		con = container.New()
		tk  = &token{}
		r   = &reader{}
	)
	r.data = []rune(s)
	r.size = len(r.data)

	for r.more() {
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
	return nil
}

func scanToken(con *container.Container, r *reader, tk *token) error {
	return nil
}
