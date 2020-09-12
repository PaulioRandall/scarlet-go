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
		newInst(CO_VAL_GET, "x"),
		newInst(CO_VAL_PVAL, "Answer = "),
		newInst(CO_VAL_PVAL, 2),
		newInst(CO_SPL_CALL, "println"),
	}

	exp := `CO_VAL_GET,  "x"
CO_VAL_PVAL, "Answer = "
CO_VAL_PVAL, 2
CO_SPL_CALL, "println"
`

	sb := &strings.Builder{}
	Print(sb, ins)
	require.Equal(t, exp, sb.String())
}
