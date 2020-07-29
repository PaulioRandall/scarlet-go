package lexeme

import (
	"testing"
)

func Test_Lexeme_prepend(t *testing.T) {

	a, b, c, _ := setup2()

	b.prepend(a)
	fullEqual2(t, a, nil, b, a)
	fullEqual2(t, b, a, nil, b)

	c.prepend(b)
	fullEqual2(t, a, nil, b, a)
	fullEqual2(t, b, a, c, b)
	fullEqual2(t, c, b, nil, c)
}

func Test_Lexeme_append(t *testing.T) {

	a, b, c, _ := setup2()

	b.append(c)
	fullEqual2(t, b, nil, c, b)
	fullEqual2(t, c, b, nil, c)

	a.append(b)
	fullEqual2(t, a, nil, b, a)
	fullEqual2(t, b, a, c, b)
	fullEqual2(t, c, b, nil, c)
}

func Test_Lexeme_remove(t *testing.T) {

	a, b, c, _ := setup2()
	feign2(a, b, c)
	a.remove()
	fullEqual2(t, b, nil, c, b)
	fullEqual2(t, c, b, nil, c)

	a, b, c, _ = setup2()
	feign2(a, b, c)
	b.remove()
	fullEqual2(t, a, nil, c, a)
	fullEqual2(t, c, a, nil, c)

	a, b, c, _ = setup2()
	feign2(a, b, c)
	c.remove()
	fullEqual2(t, a, nil, b, a)
	fullEqual2(t, b, a, nil, b)
}
