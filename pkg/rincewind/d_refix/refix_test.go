package refix

import (
	"fmt"
	"testing"

	"github.com/PaulioRandall/scarlet-go/pkg/rincewind/token"
	. "github.com/PaulioRandall/scarlet-go/pkg/rincewind/token/types"

	tkt "github.com/PaulioRandall/scarlet-go/pkg/rincewind/token/tokentest"

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

	tkt.RequireSlice(t, exps, acts)
}

func halfTok(gen GenType, sub SubType, raw string) token.Tok {
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
		halfTok(GE_SPELL, SU_UNDEFINED, "@Print"),
		halfTok(GE_PARENTHESIS, SU_PAREN_OPEN, "("),
		halfTok(GE_PARENTHESIS, SU_PAREN_CLOSE, ")"),
		halfTok(GE_TERMINATOR, SU_NEWLINE, "\n"),
	}

	// THEN parenthesis are removed
	// AND magic token indicating parameter start inserted before spell
	exp := []token.Token{
		halfTok(GE_PARAMS, SU_UNDEFINED, "("),
		halfTok(GE_SPELL, SU_UNDEFINED, "@Print"),
		halfTok(GE_TERMINATOR, SU_NEWLINE, "\n"),
	}

	doTest(t, in, exp)
}

func Test1_2(t *testing.T) {

	// WHEN refixing a spell with one argument
	// @Println(x)
	in := []token.Token{
		halfTok(GE_SPELL, SU_UNDEFINED, "@Print"),
		halfTok(GE_PARENTHESIS, SU_PAREN_OPEN, "("),
		halfTok(GE_IDENTIFIER, SU_IDENTIFIER, "x"),
		halfTok(GE_PARENTHESIS, SU_PAREN_CLOSE, ")"),
		halfTok(GE_TERMINATOR, SU_NEWLINE, "\n"),
	}

	// THEN parenthesis are removed
	// AND the argument is placed before the spell
	// AND magic token indicating parameter start inserted before the argument
	exp := []token.Token{
		halfTok(GE_PARAMS, SU_UNDEFINED, "("),
		halfTok(GE_IDENTIFIER, SU_IDENTIFIER, "x"),
		halfTok(GE_SPELL, SU_UNDEFINED, "@Print"),
		halfTok(GE_TERMINATOR, SU_NEWLINE, "\n"),
	}

	doTest(t, in, exp)
}

func Test1_3(t *testing.T) {

	// WHEN refixing a spell with multiple arguments
	// @Println(x, y, z)
	in := []token.Token{
		halfTok(GE_SPELL, SU_UNDEFINED, "@Println"),
		halfTok(GE_PARENTHESIS, SU_PAREN_OPEN, "("),
		halfTok(GE_IDENTIFIER, SU_IDENTIFIER, "x"),
		halfTok(GE_DELIMITER, SU_VALUE_DELIM, ","),
		halfTok(GE_IDENTIFIER, SU_IDENTIFIER, "y"),
		halfTok(GE_DELIMITER, SU_VALUE_DELIM, ","),
		halfTok(GE_IDENTIFIER, SU_IDENTIFIER, "z"),
		halfTok(GE_PARENTHESIS, SU_PAREN_CLOSE, ")"),
		halfTok(GE_TERMINATOR, SU_NEWLINE, "\n"),
	}

	// THEN parenthesis are removed
	// AND the delimiters are removed
	// AND the arguments are placed before the spell
	// AND magic token indicating parameter start inserted before the arguments
	exp := []token.Token{
		halfTok(GE_PARAMS, SU_UNDEFINED, "("),
		halfTok(GE_IDENTIFIER, SU_IDENTIFIER, "x"),
		halfTok(GE_IDENTIFIER, SU_IDENTIFIER, "y"),
		halfTok(GE_IDENTIFIER, SU_IDENTIFIER, "z"),
		halfTok(GE_SPELL, SU_UNDEFINED, "@Println"),
		halfTok(GE_TERMINATOR, SU_NEWLINE, "\n"),
	}

	doTest(t, in, exp)
}
