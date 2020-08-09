package lexeme

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_prettyPrint_1(t *testing.T) {

	a := lex(0, 0, "true", BOOL)
	b := lex(1, 2, "2", NUMBER)
	c := lex(99, 66, `"three"`, STRING)
	d := lex(666, 999, "x", IDENTIFIER)

	feign(a, b, c, d)

	exp := `  0:0  , BOOL      , "true"
  1:2  , NUMBER    , "2"
 99:66 , STRING    , "\"three\""
666:999, IDENTIFIER, "x"
`

	sb := &strings.Builder{}
	Print(sb, a)
	require.Equal(t, exp, sb.String())
}
