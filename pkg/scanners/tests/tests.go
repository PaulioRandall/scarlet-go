package tests

import (
	"testing"

	. "github.com/PaulioRandall/scarlet-go/pkg/token"

	"github.com/stretchr/testify/require"
)

type ScanFunc func(in string) []Token

func DoTests(t *testing.T, f ScanFunc) {
	a1(t, f)
}

func a1(t *testing.T, f ScanFunc) {

	in := "\r\n" +
		" \t\r\v\f// comment\n" +
		"123 123.456\r\n" +
		"`abc`" + `"abc"` + "\n" +
		"abc_xyz\r\n" +
		"F TRUE\n" +
		"@:=*->)\r\n"

	exp := []Token{
		// Line 0
		Token{NEWLINE, "\r\n", 0, 0},
		// Line 1
		Token{WHITESPACE, " \t\r\v\f", 1, 0},
		Token{COMMENT, "// comment", 1, 5},
		Token{NEWLINE, "\n", 1, 15},
		// Line 2
		Token{NUMBER, "123", 2, 0},
		Token{WHITESPACE, " ", 2, 3},
		Token{NUMBER, "123.456", 2, 4},
		Token{NEWLINE, "\r\n", 2, 11},
		// Line 3
		Token{STRING, "`abc`", 3, 0},
		Token{TEMPLATE, `"abc"`, 3, 5},
		Token{NEWLINE, "\n", 3, 10},
		// Line 4
		Token{ID, "abc_xyz", 4, 0},
		Token{NEWLINE, "\r\n", 4, 7},
		// Line 5
		Token{FUNC, "F", 5, 0},
		Token{WHITESPACE, " ", 5, 1},
		Token{BOOL, "TRUE", 5, 2},
		Token{NEWLINE, "\n", 5, 6},
		// Line 6
		Token{SPELL, "@", 6, 0},
		Token{ASSIGN, ":=", 6, 1},
		Token{MULTIPLY, "*", 6, 3},
		Token{RETURNS, "->", 6, 4},
		Token{PAREN_CLOSE, ")", 6, 6},
		Token{NEWLINE, "\r\n", 6, 7},
		// Line 7
		Token{EOF, "", 7, 0},
	}

	act := f(in)

	for _, tk := range exp {

		require.True(t, len(act) > 0, "[Scanner Test A1]"+
			" Expected ("+tk.String()+") but no actual tokens remain")

		require.Equal(t, tk, act[0], "[Scanner Test A1]"+
			" Expected ("+tk.String()+") but got ("+act[0].String()+")")

		if tk.Type == EOF {
			break
		}

		act = act[1:]
	}

	require.True(t, len(act) > 0, "[Scanner Test A1]"+
		" Did not expected anymore tokens but got ("+act[0].String()+")"+
		" and "+string(len(act)-1)+" others")
}
