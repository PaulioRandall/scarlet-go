package lexeme

import (
	"testing"
)

func init() {

	lex := &Lexeme{}

	_ = Snippet(lex)
	_ = Node(lex)
	var _ string = lex.String()
}

// @Deprecated
func Test_Lexeme_ShiftUp(t *testing.T) {

	a := tok("true", BOOL)
	b := tok("1", NUMBER)
	c := tok(`"abc"`, STRING)

	_ = feign(a, b, c)

	a.ShiftUp()
	fullEqual(t, a, nil, b, a)
	fullEqual(t, b, a, c, b)
	fullEqual(t, c, b, nil, c)

	b.ShiftUp()
	fullEqual(t, b, nil, a, b)
	fullEqual(t, a, b, c, a)
	fullEqual(t, c, a, nil, c)

	c.ShiftUp()
	c.ShiftUp()
	fullEqual(t, c, nil, b, c)
	fullEqual(t, b, c, a, b)
	fullEqual(t, a, b, nil, a)
}

// @Deprecated
func Test_Lexeme_ShiftDown(t *testing.T) {

	a := tok("true", BOOL)
	b := tok("1", NUMBER)
	c := tok(`"abc"`, STRING)

	_ = feign(a, b, c)

	a.ShiftDown()
	fullEqual(t, b, nil, a, b)
	fullEqual(t, a, b, c, a)
	fullEqual(t, c, a, nil, c)

	a.ShiftDown()
	fullEqual(t, b, nil, c, b)
	fullEqual(t, c, b, a, c)
	fullEqual(t, a, c, nil, a)

	a.ShiftDown()
	fullEqual(t, b, nil, c, b)
	fullEqual(t, c, b, a, c)
	fullEqual(t, a, c, nil, a)

	b.ShiftDown()
	b.ShiftDown()
	fullEqual(t, c, nil, a, c)
	fullEqual(t, a, c, b, a)
	fullEqual(t, b, a, nil, b)
}

// @Deprecated
func Test_Lexeme_Prepend(t *testing.T) {

	a := tok("true", BOOL)
	b := tok("1", NUMBER)
	c := tok(`"abc"`, STRING)

	b.Prepend(a)
	fullEqual(t, a, nil, b, a)
	fullEqual(t, b, a, nil, b)

	c.Prepend(b)
	fullEqual(t, a, nil, b, a)
	fullEqual(t, b, a, c, b)
	fullEqual(t, c, b, nil, c)
}

// @Deprecated
func Test_Lexeme_Append(t *testing.T) {

	a := tok("true", BOOL)
	b := tok("1", NUMBER)
	c := tok(`"abc"`, STRING)

	b.Append(c)
	fullEqual(t, b, nil, c, b)
	fullEqual(t, c, b, nil, c)

	a.Append(b)
	fullEqual(t, a, nil, b, a)
	fullEqual(t, b, a, c, b)
	fullEqual(t, c, b, nil, c)
}

// @Deprecated
func Test_Lexeme_Remove(t *testing.T) {

	a, b, c, _ := setupList()
	a.Remove()
	fullEqual(t, b, nil, c, b)
	fullEqual(t, c, b, nil, c)

	a, b, c, _ = setupList()
	b.Remove()
	fullEqual(t, a, nil, c, a)
	fullEqual(t, c, a, nil, c)

	a, b, c, _ = setupList()
	c.Remove()
	fullEqual(t, a, nil, b, a)
	fullEqual(t, b, a, nil, b)
}

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
