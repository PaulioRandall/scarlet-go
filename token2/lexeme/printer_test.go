package lexeme

import (
	"strings"
	"testing"

	"github.com/PaulioRandall/scarlet-go/token2/position"
	"github.com/PaulioRandall/scarlet-go/token2/token"

	"github.com/stretchr/testify/require"
)

func makeTestData() ([]Lexeme, string) {

	tm := &position.TextMarker{}
	genLex := func(v string, tk token.Token) Lexeme {
		snip := tm.Snippet(v)
		tm.Advance(v, v == "\n")
		return New(v, tk, snip)
	}

	in := []Lexeme{
		genLex("x", token.IDENT),
		genLex(" ", token.SPACE),
		genLex(":=", token.ASSIGN),
		genLex(" ", token.SPACE),
		genLex(`"abc"`, token.STRING),
		genLex("\n", token.NEWLINE),
		genLex("@Println", token.SPELL),
		genLex("(", token.L_PAREN),
		genLex("1", token.NUMBER),
		genLex(")", token.R_PAREN),
		genLex("\n", token.NEWLINE),
	}

	tm.Offset, tm.Line, tm.ColByte, tm.ColRune = 10, 10, 10, 10
	doubleDigit := genLex("\n", token.NEWLINE)
	in = append(in, doubleDigit)

	tm.Offset, tm.Line, tm.ColByte, tm.ColRune = 100, 100, 100, 100
	tripleDigit := genLex("\n", token.NEWLINE)
	in = append(in, tripleDigit)

	exp := `  0:0   ->   0:1   IDENT   "x"
  0:1   ->   0:2   SPACE   " "
  0:2   ->   0:4   ASSIGN  ":="
  0:4   ->   0:5   SPACE   " "
  0:5   ->   0:10  STRING  "\"abc\""
  0:10  ->   0:11  NEWLINE "\n"
  1:0   ->   1:8   SPELL   "@Println"
  1:8   ->   1:9   L_PAREN "("
  1:9   ->   1:10  NUMBER  "1"
  1:10  ->   1:11  R_PAREN ")"
  1:11  ->   1:12  NEWLINE "\n"
 10:10  ->  10:11  NEWLINE "\n"
100:100 -> 100:101 NEWLINE "\n"
`

	return in, exp
}

type sbWriter struct {
	sb *strings.Builder
}

// io.StringWriter, because sb.WriteString has pointer receiver
func (sbw sbWriter) WriteString(s string) (int, error) {
	return sbw.sb.WriteString(s)
}

func TestPrinter(t *testing.T) {

	in, exp := makeTestData()
	sbw := sbWriter{
		sb: &strings.Builder{},
	}

	Print(sbw, in)
	act := sbw.sb.String()

	require.Equal(t, exp, act)
}
