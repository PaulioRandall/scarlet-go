package shunt

import (
	"fmt"
	"testing"

	"github.com/PaulioRandall/scarlet-go/pkg/rincewind/shared/token"
	. "github.com/PaulioRandall/scarlet-go/pkg/rincewind/shared/token/types"

	"github.com/PaulioRandall/scarlet-go/pkg/rincewind/shared/testutils"
	"github.com/stretchr/testify/require"
)

func doTest(t *testing.T, in, exps []token.Token) {

	var (
		tk     token.Token
		e      error
		stream = token.NewStream(in)
		acts   = []token.Token{}
	)

	for f := New(stream); f != nil; {
		if tk, f, e = f(); e != nil {
			require.NotNil(t, fmt.Sprintf("%+v", e))
		}

		acts = append(acts, tk)
	}

	testutils.RequireTokenSlice(t, exps, acts)
}

func tok(gen GenType, sub SubType, raw string) token.Tok {
	return token.Tok{
		Gen:    gen,
		Sub:    sub,
		RawStr: raw,
		ColEnd: len(raw),
	}
}

func Test1_1(t *testing.T) {

	// WHEN refixing a spell with no arguments
	// @Println()
	in := []token.Token{
		tok(GEN_SPELL, SU_UNDEFINED, "@Print"),
		tok(GEN_PARENTHESIS, SU_PAREN_OPEN, "("),
		tok(GEN_PARENTHESIS, SU_PAREN_CLOSE, ")"),
		tok(GEN_TERMINATOR, SU_NEWLINE, "\n"),
	}

	// THEN parenthesis are removed
	// AND magic token indicating parameter start inserted before spell
	exp := []token.Token{
		tok(GEN_PARAMS, SU_UNDEFINED, "("),
		tok(GEN_SPELL, SU_UNDEFINED, "@Print"),
		tok(GEN_TERMINATOR, SU_NEWLINE, "\n"),
	}

	doTest(t, in, exp)
}

func Test1_2(t *testing.T) {

	// WHEN refixing a spell with one argument
	// @Println(x)
	in := []token.Token{
		tok(GEN_SPELL, SU_UNDEFINED, "@Print"),
		tok(GEN_PARENTHESIS, SU_PAREN_OPEN, "("),
		tok(GEN_IDENTIFIER, SU_IDENTIFIER, "x"),
		tok(GEN_PARENTHESIS, SU_PAREN_CLOSE, ")"),
		tok(GEN_TERMINATOR, SU_NEWLINE, "\n"),
	}

	// THEN parenthesis are removed
	// AND the argument is placed before the spell
	// AND magic token indicating parameter start inserted before the argument
	exp := []token.Token{
		tok(GEN_PARAMS, SU_UNDEFINED, "("),
		tok(GEN_IDENTIFIER, SU_IDENTIFIER, "x"),
		tok(GEN_SPELL, SU_UNDEFINED, "@Print"),
		tok(GEN_TERMINATOR, SU_NEWLINE, "\n"),
	}

	doTest(t, in, exp)
}

func Test1_3(t *testing.T) {

	// WHEN refixing a spell with multiple arguments
	// @Println(x, y, z)
	in := []token.Token{
		tok(GEN_SPELL, SU_UNDEFINED, "@Println"),
		tok(GEN_PARENTHESIS, SU_PAREN_OPEN, "("),
		tok(GEN_IDENTIFIER, SU_IDENTIFIER, "x"),
		tok(GEN_DELIMITER, SU_VALUE_DELIM, ","),
		tok(GEN_IDENTIFIER, SU_IDENTIFIER, "y"),
		tok(GEN_DELIMITER, SU_VALUE_DELIM, ","),
		tok(GEN_IDENTIFIER, SU_IDENTIFIER, "z"),
		tok(GEN_PARENTHESIS, SU_PAREN_CLOSE, ")"),
		tok(GEN_TERMINATOR, SU_NEWLINE, "\n"),
	}

	// THEN parenthesis are removed
	// AND the delimiters are removed
	// AND the arguments are placed before the spell
	// AND magic token indicating parameter start inserted before the arguments
	exp := []token.Token{
		tok(GEN_PARAMS, SU_UNDEFINED, "("),
		tok(GEN_IDENTIFIER, SU_IDENTIFIER, "x"),
		tok(GEN_IDENTIFIER, SU_IDENTIFIER, "y"),
		tok(GEN_IDENTIFIER, SU_IDENTIFIER, "z"),
		tok(GEN_SPELL, SU_UNDEFINED, "@Println"),
		tok(GEN_TERMINATOR, SU_NEWLINE, "\n"),
	}

	doTest(t, in, exp)
}
