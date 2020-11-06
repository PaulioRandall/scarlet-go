package sanitiser

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/scarlet/token"

	"github.com/stretchr/testify/require"
)

func TestRedundant_1(t *testing.T) {
	in := []token.Lexeme{
		token.MakeTok(" ", token.SPACE),
	}
	exp := []token.Lexeme{}
	require.Equal(t, exp, Sanitise(in))
}

func TestRedundant_2(t *testing.T) {
	in := []token.Lexeme{
		token.MakeTok("# Scarlet", token.COMMENT),
	}
	exp := []token.Lexeme{}
	require.Equal(t, exp, Sanitise(in))
}

func TestLeadingTerminators_1(t *testing.T) {
	in := []token.Lexeme{
		token.MakeTok("\n", token.TERMINATOR),
		token.MakeTok(";", token.TERMINATOR),
	}
	exp := []token.Lexeme{}
	require.Equal(t, exp, Sanitise(in))
}

func TestSuccessiveTerminators_1(t *testing.T) {
	in := []token.Lexeme{
		token.MakeTok("x", token.IDENT),
		token.MakeTok("\n", token.TERMINATOR),
		token.MakeTok(";", token.TERMINATOR),
	}
	exp := []token.Lexeme{
		token.MakeTok("x", token.IDENT),
		token.MakeTok("\n", token.TERMINATOR),
	}
	require.Equal(t, exp, Sanitise(in))
}

func TestNewlineAfterOpener_1(t *testing.T) {
	in := []token.Lexeme{
		token.MakeTok("(", token.L_PAREN),
		token.MakeTok("\n", token.NEWLINE),
	}
	exp := []token.Lexeme{
		token.MakeTok("(", token.L_PAREN),
	}
	require.Equal(t, exp, Sanitise(in))
}

func TestNewlineAfterDelim_1(t *testing.T) {
	in := []token.Lexeme{
		token.MakeTok(",", token.DELIM),
		token.MakeTok("\n", token.NEWLINE),
	}
	exp := []token.Lexeme{
		token.MakeTok(",", token.DELIM),
	}
	require.Equal(t, exp, Sanitise(in))
}

func TestDelimBeforeRParen_1(t *testing.T) {
	in := []token.Lexeme{
		token.MakeTok(",", token.DELIM),
		token.MakeTok(")", token.R_PAREN),
	}
	exp := []token.Lexeme{
		token.MakeTok(")", token.R_PAREN),
	}
	require.Equal(t, exp, Sanitise(in))
}

func TestTerminatorBeforeRCurly_1(t *testing.T) {
	in := []token.Lexeme{
		token.MakeTok("\n", token.NEWLINE),
		token.MakeTok("}", token.R_CURLY),
	}
	exp := []token.Lexeme{
		token.MakeTok("}", token.R_CURLY),
	}
	require.Equal(t, exp, Sanitise(in))
}

func TestFull_1(t *testing.T) {

	in := []token.Lexeme{
		token.MakeTok(" ", token.SPACE),
		token.MakeTok("\n", token.NEWLINE),
		token.MakeTok(" ", token.SPACE),
		token.MakeTok("\n", token.NEWLINE),
		token.MakeTok(" ", token.SPACE),
		token.MakeTok("@Println", token.SPELL),
		token.MakeTok(" ", token.SPACE),
		token.MakeTok("(", token.L_PAREN),
		token.MakeTok(" ", token.SPACE),
		token.MakeTok("\n", token.NEWLINE),
		token.MakeTok(" ", token.SPACE),
		token.MakeTok("1", token.NUMBER),
		token.MakeTok(" ", token.SPACE),
		token.MakeTok(",", token.DELIM),
		token.MakeTok("\n", token.NEWLINE),
		token.MakeTok("\n", token.NEWLINE),
		token.MakeTok("\n", token.NEWLINE),
		token.MakeTok(" ", token.SPACE),
		token.MakeTok("1", token.NUMBER),
		token.MakeTok(" ", token.SPACE),
		token.MakeTok(",", token.DELIM),
		token.MakeTok(" ", token.SPACE),
		token.MakeTok("\n", token.NEWLINE),
		token.MakeTok("\n", token.NEWLINE),
		token.MakeTok(" ", token.SPACE),
		token.MakeTok(")", token.R_PAREN),
		token.MakeTok(" ", token.SPACE),
		token.MakeTok("\n", token.NEWLINE),
	}

	// @Println(1,1)
	exp := []token.Lexeme{
		token.MakeTok("@Println", token.SPELL),
		token.MakeTok("(", token.L_PAREN),
		token.MakeTok("1", token.NUMBER),
		token.MakeTok(",", token.DELIM),
		token.MakeTok("1", token.NUMBER),
		token.MakeTok(")", token.R_PAREN),
		token.MakeTok("\n", token.NEWLINE),
	}

	require.Equal(t, exp, Sanitise(in))
}

func TestFull_2(t *testing.T) {

	// [true] {
	//   "abc"
	//   "xyz"
	// }
	in := []token.Lexeme{
		token.MakeTok("[", token.L_SQUARE),
		token.MakeTok("true", token.TRUE),
		token.MakeTok("]", token.R_SQUARE),
		token.MakeTok(" ", token.SPACE),
		token.MakeTok("{", token.L_CURLY),
		token.MakeTok("\n", token.NEWLINE),
		token.MakeTok("\t", token.SPACE),
		token.MakeTok(`"abc"`, token.STRING),
		token.MakeTok("\n", token.NEWLINE),
		token.MakeTok("\t", token.SPACE),
		token.MakeTok(`"xyz"`, token.STRING),
		token.MakeTok("\n", token.NEWLINE),
		token.MakeTok("}", token.R_CURLY),
	}

	// [true] {"abc"
	// "xyz"}
	exp := []token.Lexeme{
		token.MakeTok("[", token.L_SQUARE),
		token.MakeTok("true", token.TRUE),
		token.MakeTok("]", token.R_SQUARE),
		token.MakeTok("{", token.L_CURLY),
		token.MakeTok(`"abc"`, token.STRING),
		token.MakeTok("\n", token.NEWLINE),
		token.MakeTok(`"xyz"`, token.STRING),
		token.MakeTok("}", token.R_CURLY),
	}

	require.Equal(t, exp, Sanitise(in))
}