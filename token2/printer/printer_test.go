package printer

import (
	"strings"
	"testing"

	"github.com/PaulioRandall/scarlet-go/token2/lexeme"
	"github.com/PaulioRandall/scarlet-go/token2/token"

	"github.com/stretchr/testify/require"
)

func makeTestData() ([]lexeme.Lexeme, string) {

	in := []lexeme.Lexeme{
		lexeme.New("x", token.IDENT, 0, 0),
		lexeme.New(" ", token.SPACE, 0, 1),
		lexeme.New(":=", token.ASSIGN, 0, 2),
		lexeme.New(" ", token.SPACE, 0, 4),
		lexeme.New(`"abc"`, token.STRING, 0, 5),
		lexeme.New("\n", token.NEWLINE, 0, 10),
		lexeme.New("@Println", token.SPELL, 1, 0),
		lexeme.New("(", token.L_PAREN, 1, 8),
		lexeme.New("1", token.NUMBER, 1, 9),
		lexeme.New(")", token.R_PAREN, 1, 10),
		lexeme.New("\n", token.NEWLINE, 1, 11),
		lexeme.New("\n", token.NEWLINE, 10, 10),
		lexeme.New("\n", token.NEWLINE, 100, 100),
	}

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
