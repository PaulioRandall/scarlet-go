package lexeme

import (
	"testing"
)

func Test_prepend(t *testing.T) {

	a := lex(0, 0, "1st", BOOL)
	b := lex(0, 4, "2nd", NUMBER)
	c := lex(0, 5, "3rd", STRING)

	prepend(b, a)
	fullEqual(t, a, nil, b, a)
	fullEqual(t, b, a, nil, b)

	prepend(c, b)
	fullEqual(t, a, nil, b, a)
	fullEqual(t, b, a, c, b)
	fullEqual(t, c, b, nil, c)
}

func Test_append(t *testing.T) {

	a := lex(0, 0, "1st", BOOL)
	b := lex(0, 4, "2nd", NUMBER)
	c := lex(0, 5, "3rd", STRING)

	append(b, c)
	fullEqual(t, b, nil, c, b)
	fullEqual(t, c, b, nil, c)

	append(a, b)
	fullEqual(t, a, nil, b, a)
	fullEqual(t, b, a, c, b)
	fullEqual(t, c, b, nil, c)
}

func Test_remove(t *testing.T) {

	a := lex(0, 0, "1st", BOOL)
	b := lex(0, 4, "2nd", NUMBER)
	c := lex(0, 5, "3rd", STRING)

	feign(a, b, c)
	remove(a)
	fullEqual(t, b, nil, c, b)
	fullEqual(t, c, b, nil, c)

	feign(a, b, c)
	remove(b)
	fullEqual(t, a, nil, c, a)
	fullEqual(t, c, a, nil, c)

	feign(a, b, c)
	remove(c)
	fullEqual(t, a, nil, b, a)
	fullEqual(t, b, a, nil, b)
}
