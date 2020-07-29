package lexeme

import (
	"testing"
)

func Test_prepend(t *testing.T) {

	a, b, c, _ := setup()

	prepend(b, a)
	fullEqual(t, a, nil, b, a)
	fullEqual(t, b, a, nil, b)

	prepend(c, b)
	fullEqual(t, a, nil, b, a)
	fullEqual(t, b, a, c, b)
	fullEqual(t, c, b, nil, c)
}

func Test_append(t *testing.T) {

	a, b, c, _ := setup()

	append(b, c)
	fullEqual(t, b, nil, c, b)
	fullEqual(t, c, b, nil, c)

	append(a, b)
	fullEqual(t, a, nil, b, a)
	fullEqual(t, b, a, c, b)
	fullEqual(t, c, b, nil, c)
}

func Test_remove(t *testing.T) {

	a, b, c, _ := setup()
	feign(a, b, c)
	remove(a)
	fullEqual(t, b, nil, c, b)
	fullEqual(t, c, b, nil, c)

	a, b, c, _ = setup()
	feign(a, b, c)
	remove(b)
	fullEqual(t, a, nil, c, a)
	fullEqual(t, c, a, nil, c)

	a, b, c, _ = setup()
	feign(a, b, c)
	remove(c)
	fullEqual(t, a, nil, b, a)
	fullEqual(t, b, a, nil, b)
}
