package processor2

import (
	"testing"

	//"github.com/PaulioRandall/scarlet-go/scarlet/token"
	"github.com/PaulioRandall/scarlet-go/scarlet/tree"
	"github.com/PaulioRandall/scarlet-go/scarlet/value"
	"github.com/PaulioRandall/scarlet-go/scarlet/value/number"

	"github.com/stretchr/testify/require"
)

func numValue(n string) value.Num {
	return value.Num{Number: number.New(n)}
}

func TestLiteral_Bool_1(t *testing.T) {

	in := tree.BoolLit{Val: true}
	exp := value.Bool(true)

	env := newTestEnv()
	act := Literal(env, in)

	require.Nil(t, env.err, "ERROR: %+v", env.err)
	require.False(t, env.exitFlag)
	require.Equal(t, 0, env.exitCode)
	require.Equal(t, exp, act)
}

func TestLiteral_Bool_2(t *testing.T) {

	in := tree.BoolLit{Val: false}
	exp := value.Bool(false)

	env := newTestEnv()
	act := Literal(env, in)

	require.Nil(t, env.err, "ERROR: %+v", env.err)
	require.False(t, env.exitFlag)
	require.Equal(t, 0, env.exitCode)
	require.Equal(t, exp, act)
}
