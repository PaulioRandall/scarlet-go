package testutils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type Snippet interface {
	Begin() (int, int)
	End() (int, int)
}

/*
func RequireSnippet(t *testing.T, exp, act Snippet) {

	if v, ok := exp.(fmt.Stringer); ok {
		requireSnippet(t, exp, act, v.String())
		return
	}

		line, col := exp.Begin()
		endLine, endCol := exp.End()
		msg := fmt.Sprintf("Begin %d:%d, End %d:%d", line, col, endLine, endCol)
	requireSnippet(t, exp, act, msg)
}
*/
func requireSnippet(t *testing.T, exp, act Snippet, msg string) {
	requirePos(t, exp.Begin, act.Begin, msg)
	requirePos(t, exp.End, act.End, msg)
}

func requirePos(t *testing.T, exp, act func() (int, int), msg string) {
	expLine, expCol := exp()
	actLine, actCol := act()
	require.Equal(t, expLine, actLine, msg)
	require.Equal(t, expCol, actCol, msg)
}
