package token

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

type iterator struct {
	i    int
	lexs []Lexeme
}

func (itr *iterator) More() bool {
	return itr.i+1 < len(itr.lexs)
}

func (itr *iterator) Next() Lexeme {
	itr.i++
	return itr.lexs[itr.i]
}

func (itr *iterator) JumpToStart() {
	itr.i = -1
}

func makeTestData() (LexemeIterator, string) {

	tm := &TextMarker{}
	genLex := func(v string, tk Token) Lexeme {
		snip := tm.Snippet(v)
		tm.Advance(v)
		return Make(v, tk, snip)
	}

	itr := &iterator{
		i: -1,
		lexs: []Lexeme{
			genLex("x", IDENT),
			genLex(" ", SPACE),
			genLex(":=", ASSIGN),
			genLex(" ", SPACE),
			genLex(`"abc"`, STRING),
			genLex("\n", NEWLINE),
			genLex("@Println", SPELL),
			genLex("(", L_PAREN),
			genLex("1", NUMBER),
			genLex(")", R_PAREN),
			genLex("\n", NEWLINE),
		},
	}

	// Add lexeme with line and rune column with two digits
	tm.Offset, tm.Line, tm.ColByte, tm.ColRune = 10, 10, 10, 10
	doubleDigit := genLex("\n", NEWLINE)
	itr.lexs = append(itr.lexs, doubleDigit)

	// Add lexeme with line and rune column with three digits
	tm.Offset, tm.Line, tm.ColByte, tm.ColRune = 100, 100, 100, 100
	tripleDigit := genLex("\n", NEWLINE)
	itr.lexs = append(itr.lexs, tripleDigit)

	exp := `  0:0   ->   0:1   IDENT   "x"
  0:1   ->   0:2   SPACE   " "
  0:2   ->   0:4   ASSIGN  ":="
  0:4   ->   0:5   SPACE   " "
  0:5   ->   0:10  STRING  "\"abc\""
  0:10  ->   1:0   NEWLINE "\n"
  1:0   ->   1:8   SPELL   "@Println"
  1:8   ->   1:9   L_PAREN "("
  1:9   ->   1:10  NUMBER  "1"
  1:10  ->   1:11  R_PAREN ")"
  1:11  ->   2:0   NEWLINE "\n"
 10:10  ->  11:0   NEWLINE "\n"
100:100 -> 101:0   NEWLINE "\n"
`

	return itr, exp
}

type sbWriter struct {
	sb *strings.Builder
}

// io.StringWriter, because sb.WriteString has pointer receiver
func (sbw sbWriter) WriteString(s string) (int, error) {
	return sbw.sb.WriteString(s)
}

func TestPrinter(t *testing.T) {

	itr, exp := makeTestData()
	sbw := sbWriter{
		sb: &strings.Builder{},
	}

	Print(sbw, itr)
	act := sbw.sb.String()

	require.Equal(t, exp, act)
}
