package lexeme

import (
	"testing"
)

func Test_Lexeme_prepend(t *testing.T) {

	a, b, c, _ := setup()

	b.prepend(a)
	fullEqual(t, a, nil, b, a)
	fullEqual(t, b, a, nil, b)

	c.prepend(b)
	fullEqual(t, a, nil, b, a)
	fullEqual(t, b, a, c, b)
	fullEqual(t, c, b, nil, c)
}

func Test_Lexeme_append(t *testing.T) {

	a, b, c, _ := setup()

	b.append(c)
	fullEqual(t, b, nil, c, b)
	fullEqual(t, c, b, nil, c)

	a.append(b)
	fullEqual(t, a, nil, b, a)
	fullEqual(t, b, a, c, b)
	fullEqual(t, c, b, nil, c)
}

func Test_Lexeme_remove(t *testing.T) {

	a, b, c, _ := setup()
	feign(a, b, c)
	a.remove()
	fullEqual(t, b, nil, c, b)
	fullEqual(t, c, b, nil, c)

	a, b, c, _ = setup()
	feign(a, b, c)
	b.remove()
	fullEqual(t, a, nil, c, a)
	fullEqual(t, c, a, nil, c)

	a, b, c, _ = setup()
	feign(a, b, c)
	c.remove()
	fullEqual(t, a, nil, b, a)
	fullEqual(t, b, a, nil, b)
}
