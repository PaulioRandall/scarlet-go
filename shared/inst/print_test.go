package inst

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func newInst(code Code, data interface{}) Instruction {
	return Instruction{
		Code: code,
		Data: data,
	}
}

func Test_Print_1(t *testing.T) {

	ins := []Instruction{
		newInst(CO_CTX_GET, "x"),
		newInst(CO_VAL_PUSH, "Answer = "),
		newInst(CO_VAL_PUSH, 2),
		newInst(CO_SPELL, "println"),
	}

	exp := `CO_CTX_GET , "x"
CO_VAL_PUSH, "Answer = "
CO_VAL_PUSH, 2
CO_SPELL   , "println"
`

	sb := &strings.Builder{}
	Print(sb, ins)
	require.Equal(t, exp, sb.String())
}
