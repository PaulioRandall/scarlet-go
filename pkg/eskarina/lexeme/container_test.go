package lexeme

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/pkg/eskarina/prop"

	"github.com/stretchr/testify/require"
)

func setup() (a *Lexeme, b *Lexeme, c *Lexeme) {
	a = tok("true", prop.PR_BOOL)
	b = tok("1", prop.PR_NUMBER)
	c = tok(`"abc"`, prop.PR_STRING)
	return
}

func setupList() (a *Lexeme, b *Lexeme, c *Lexeme) {
	a, b, c = setup()
	_ = feign(a, b, c)
	return
}

func equal(t *testing.T, exp, act *Lexeme) {

	if exp == nil {
		require.Nil(t, act)
		return
	}

	require.NotNil(t, act)
	require.Equal(t, exp.Props, act.Props)
	require.Equal(t, exp.Raw, act.Raw)
}

func requireEqual(t *testing.T, exp, prev, next, act *Lexeme) {

	require.NotNil(t, act)
	require.Equal(t, exp.Props, act.Props)
	require.Equal(t, exp.Raw, act.Raw)

	equal(t, prev, act.Prev)
	equal(t, next, act.Next)
}

func Test_NewContainer(t *testing.T) {

	a, b, c := setupList()
	con := NewContainer(a)

	require.Equal(t, 3, con.size)
	requireEqual(t, a, nil, b, con.first)
	requireEqual(t, b, a, c, con.first.Next)
	requireEqual(t, b, a, c, con.last.Prev)
	requireEqual(t, c, b, nil, con.last)
}

func Test_Container_Get(t *testing.T) {

	a, b, c := setupList()
	con := NewContainer(a)

	equal(t, a, con.Get(0))
	equal(t, b, con.Get(1))
	equal(t, c, con.Get(2))

	require.Panics(t, func() {
		con.Get(-1)
	})

	require.Panics(t, func() {
		con.Get(3)
	})
}

func Test_Container_Prepend(t *testing.T) {

	a, b, c := setup()
	con := Container{}

	con.Prepend(a)
	requireEqual(t, a, nil, nil, con.first)
	requireEqual(t, a, nil, nil, con.last)

	con.Prepend(b)
	requireEqual(t, b, nil, a, con.first)
	requireEqual(t, a, b, nil, con.last)

	con.Prepend(c)
	requireEqual(t, c, nil, b, con.first)
	requireEqual(t, b, c, a, con.first.Next)
	requireEqual(t, b, c, a, con.last.Prev)
	requireEqual(t, a, b, nil, con.last)
}

func Test_Container_Append(t *testing.T) {

	a, b, c := setup()
	con := Container{}

	con.Append(a)
	requireEqual(t, a, nil, nil, con.first)
	requireEqual(t, a, nil, nil, con.last)

	con.Append(b)
	requireEqual(t, a, nil, b, con.first)
	requireEqual(t, b, a, nil, con.last)

	con.Append(c)
	requireEqual(t, a, nil, b, con.first)
	requireEqual(t, b, a, c, con.first.Next)
	requireEqual(t, b, a, c, con.last.Prev)
	requireEqual(t, c, b, nil, con.last)
}
