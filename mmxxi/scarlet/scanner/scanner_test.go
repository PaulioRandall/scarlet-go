package scanner

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/mmxxi/scarlet/scroll"
	"github.com/PaulioRandall/scarlet-go/mmxxi/scarlet/token"

	"github.com/stretchr/testify/require"
)

func TestSoloTokens(t *testing.T) {
	tests := map[string]token.Lexeme{
		// Redundant
		" ":         token.MakeLex2(token.SPACE, " "),
		"   ":       token.MakeLex2(token.SPACE, "   "),
		"\t\f\v":    token.MakeLex2(token.SPACE, "\t\f\v"),
		"# Comment": token.MakeLex2(token.COMMENT, "# Comment"),

		// Identifiers
		"a":    token.MakeLex2(token.IDENT, "a"),
		"a2":   token.MakeLex2(token.IDENT, "a2"),
		"a2_b": token.MakeLex2(token.IDENT, "a2_b"),
		"_":    token.MakeLex2(token.IDENT, "_"),

		// Literals
		"true":      token.MakeLex2(token.BOOL, "true"),
		"false":     token.MakeLex2(token.BOOL, "false"),
		"123":       token.MakeLex2(token.NUM, "123"),
		"123.456":   token.MakeLex2(token.NUM, "123.456"),
		`""`:        token.MakeLex2(token.STR, `""`),
		`"Scarlet"`: token.MakeLex2(token.STR, `"Scarlet"`),

		// Keywords
		"B":     token.MakeLex2(token.T_BOOL, "B"),
		"N":     token.MakeLex2(token.T_NUM, "N"),
		"S":     token.MakeLex2(token.T_STR, "S"),
		"E":     token.MakeLex2(token.E_FUNC, "E"),
		"F":     token.MakeLex2(token.FUNC, "F"),
		"loop":  token.MakeLex2(token.LOOP, "loop"),
		"match": token.MakeLex2(token.MATCH, "match"),
		"type":  token.MakeLex2(token.TYPE, "type"),

		// Operators
		":=": token.MakeLex2(token.DEFINE, ":="),
		"<-": token.MakeLex2(token.ASSIGN, "<-"),
		"->": token.MakeLex2(token.OUTPUT, "->"),

		"+": token.MakeLex2(token.ADD, "+"),
		"-": token.MakeLex2(token.SUB, "-"),
		"*": token.MakeLex2(token.MUL, "*"),
		"/": token.MakeLex2(token.DIV, "/"),
		"%": token.MakeLex2(token.REM, "%"),

		"&&": token.MakeLex2(token.AND, "&&"),
		"||": token.MakeLex2(token.OR, "||"),

		"==": token.MakeLex2(token.EQU, "=="),
		"!=": token.MakeLex2(token.NEQ, "!="),
		"<=": token.MakeLex2(token.LTE, "<="),
		"<":  token.MakeLex2(token.LT, "<"),
		">=": token.MakeLex2(token.MTE, ">="),
		">":  token.MakeLex2(token.MT, ">"),

		"!": token.MakeLex2(token.NOT, "!"),
		"?": token.MakeLex2(token.QUE, "?"),

		// Delimiters
		";": token.MakeLex2(token.TERMINATOR, ";"),
		"@": token.MakeLex2(token.SPELL, "@"),
		",": token.MakeLex2(token.DELIM, ","),
		":": token.MakeLex2(token.REF, ":"),

		"(": token.MakeLex2(token.L_PAREN, "("),
		")": token.MakeLex2(token.R_PAREN, ")"),
		"[": token.MakeLex2(token.L_BRACK, "["),
		"]": token.MakeLex2(token.R_BRACK, "]"),
		"{": token.MakeLex2(token.L_BRACE, "{"),
		"}": token.MakeLex2(token.R_BRACE, "}"),
	}

	for in, exp := range tests {
		sr := scroll.NewReader(in)
		tks, e := ScanAll(sr)
		require.Nil(t, e, "%q: Unexpected error: %+v", in, e)
		require.Equal(t, 1, len(tks), "%q", in)
		require.Equal(t, exp, tks[0], "%q", in)
	}
}

func TestSoloNewlinesTokens(t *testing.T) {

	tests := []string{"\n", "\r\n"}

	for _, in := range tests {
		sr := scroll.NewReader(in)
		tks, e := ScanAll(sr)

		exp := token.MakeLex2(token.TERMINATOR, in)
		exp.Snippet.End.Line = 1
		exp.Snippet.End.ColByte = 0
		exp.Snippet.End.ColRune = 0

		require.Nil(t, e, "Unexpected errors: %+v", e)
		require.Equal(t, 1, len(tks))
		require.Equal(t, exp, tks[0])
	}
}

func TestBadSoloTokens(t *testing.T) {
	tests := []string{
		"£",
		"\r",
		".",
		".123",
		"123.",
		`"`,
		`"\"`,
	}

	for _, in := range tests {
		sr := scroll.NewReader(in)
		_, e := ScanAll(sr)
		require.NotNil(t, e, "%q: Expected error", in)
	}
}

func TestMultipleTokens(t *testing.T) {

	tm := &scroll.TextMarker{}
	genExp := func(tk token.Token, v string) token.Lexeme {
		return token.MakeLex(tk, tm.Advance(v))
	}

	in := "f := F(a N, b N -> c N) {\n" +
		"\tc <- a + b # Comment\r\n" +
		"}\n"

	exps := []token.Lexeme{
		genExp(token.IDENT, "f"),
		genExp(token.SPACE, " "),
		genExp(token.DEFINE, ":="),
		genExp(token.SPACE, " "),
		genExp(token.FUNC, "F"),
		genExp(token.L_PAREN, "("),
		genExp(token.IDENT, "a"),
		genExp(token.SPACE, " "),
		genExp(token.T_NUM, "N"),
		genExp(token.DELIM, ","),
		genExp(token.SPACE, " "),
		genExp(token.IDENT, "b"),
		genExp(token.SPACE, " "),
		genExp(token.T_NUM, "N"),
		genExp(token.SPACE, " "),
		genExp(token.OUTPUT, "->"),
		genExp(token.SPACE, " "),
		genExp(token.IDENT, "c"),
		genExp(token.SPACE, " "),
		genExp(token.T_NUM, "N"),
		genExp(token.R_PAREN, ")"),
		genExp(token.SPACE, " "),
		genExp(token.L_BRACE, "{"),
		genExp(token.TERMINATOR, "\n"),
		genExp(token.SPACE, "\t"),
		genExp(token.IDENT, "c"),
		genExp(token.SPACE, " "),
		genExp(token.ASSIGN, "<-"),
		genExp(token.SPACE, " "),
		genExp(token.IDENT, "a"),
		genExp(token.SPACE, " "),
		genExp(token.ADD, "+"),
		genExp(token.SPACE, " "),
		genExp(token.IDENT, "b"),
		genExp(token.SPACE, " "),
		genExp(token.COMMENT, "# Comment"),
		genExp(token.TERMINATOR, "\r\n"),
		genExp(token.R_BRACE, "}"),
		genExp(token.TERMINATOR, "\n"),
	}

	sr := scroll.NewReader(in)
	tks, e := ScanAll(sr)

	for i := range exps {
		require.Nil(t, e, "Unexpected errors: %+v", e)
		require.Equal(t, exps[i], tks[i])
	}
	require.Equal(t, len(exps), len(tks))
}
