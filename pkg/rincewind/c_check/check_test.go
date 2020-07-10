package check

import (
	"fmt"
	"testing"

	"github.com/PaulioRandall/scarlet-go/pkg/rincewind/shared/token"
	. "github.com/PaulioRandall/scarlet-go/pkg/rincewind/shared/token/types"

	"github.com/PaulioRandall/scarlet-go/pkg/rincewind/shared/testutils"
	"github.com/stretchr/testify/require"
)

func doTest(t *testing.T, in []token.Token) {

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

	testutils.RequireTokenSlice(t, in, acts)
}

func doErrorTest(t *testing.T, in []token.Token) {

	var (
		e   error
		itr = token.NewStream(in)
	)

	for f := New(itr); f != nil; {
		if _, f, e = f(); e != nil {
			return
		}
	}

	require.Fail(t, "Expected error")
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

	// WHEN checking a spell with no arguments
	// THEN no errors should be returned
	// @Println()
	in := []token.Token{
		tok(GE_SPELL, SU_UNDEFINED, "@Print"),
		tok(GE_PARENTHESIS, SU_PAREN_OPEN, "("),
		tok(GE_PARENTHESIS, SU_PAREN_CLOSE, ")"),
		tok(GE_TERMINATOR, SU_NEWLINE, "\n"),
	}

	doTest(t, in)
}

func Test1_2(t *testing.T) {

	// WHEN checking a spell with one argument
	// THEN no errors should be returned
	// @Println(x)
	in := []token.Token{
		tok(GE_SPELL, SU_UNDEFINED, "@Print"),
		tok(GE_PARENTHESIS, SU_PAREN_OPEN, "("),
		tok(GE_IDENTIFIER, SU_IDENTIFIER, "x"),
		tok(GE_PARENTHESIS, SU_PAREN_CLOSE, ")"),
		tok(GE_TERMINATOR, SU_NEWLINE, "\n"),
	}

	doTest(t, in)
}

func Test1_3(t *testing.T) {

	// WHEN checking a spell with multiple arguments
	// THEN no errors should be returned
	// @Println(x, 1, true)
	in := []token.Token{
		tok(GE_SPELL, SU_UNDEFINED, "@Println"),
		tok(GE_PARENTHESIS, SU_PAREN_OPEN, "("),
		tok(GE_IDENTIFIER, SU_IDENTIFIER, "x"),
		tok(GE_DELIMITER, SU_VALUE_DELIM, ","),
		tok(GE_LITERAL, SU_NUMBER, "1"),
		tok(GE_DELIMITER, SU_VALUE_DELIM, ","),
		tok(GE_LITERAL, SU_BOOL, "true"),
		tok(GE_PARENTHESIS, SU_PAREN_CLOSE, ")"),
		tok(GE_TERMINATOR, SU_NEWLINE, "\n"),
	}

	doTest(t, in)
}

func Test2_1(t *testing.T) {

	// WHEN checking a spell with missing opening parenthesis
	// THEN an error should be returned
	// @Println)
	in := []token.Token{
		tok(GE_SPELL, SU_UNDEFINED, "@Println"),
		tok(GE_PARENTHESIS, SU_PAREN_CLOSE, ")"),
		tok(GE_TERMINATOR, SU_NEWLINE, "\n"),
	}

	doErrorTest(t, in)
}

func Test2_2(t *testing.T) {

	// WHEN checking a spell with missing closing parenthesis
	// THEN an error should be returned
	// @Println(
	in := []token.Token{
		tok(GE_SPELL, SU_UNDEFINED, "@Println"),
		tok(GE_PARENTHESIS, SU_PAREN_OPEN, "("),
		tok(GE_TERMINATOR, SU_NEWLINE, "\n"),
	}

	doErrorTest(t, in)
}

func Test2_3(t *testing.T) {

	// WHEN checking a spell with a stray value delimiter
	// THEN an error should be returned
	// @Println(
	in := []token.Token{
		tok(GE_SPELL, SU_UNDEFINED, "@Println"),
		tok(GE_PARENTHESIS, SU_PAREN_OPEN, "("),
		tok(GE_DELIMITER, SU_VALUE_DELIM, ","),
		tok(GE_PARENTHESIS, SU_PAREN_CLOSE, ")"),
		tok(GE_TERMINATOR, SU_NEWLINE, "\n"),
	}

	doErrorTest(t, in)
}

func Test2_4(t *testing.T) {

	// WHEN checking a spell with a stray value delimiter
	// THEN an error should be returned
	// @Println(
	in := []token.Token{
		tok(GE_SPELL, SU_UNDEFINED, "@Println"),
		tok(GE_PARENTHESIS, SU_PAREN_OPEN, "("),
		tok(GE_IDENTIFIER, SU_IDENTIFIER, "x"),
		tok(GE_DELIMITER, SU_VALUE_DELIM, ","),
		tok(GE_PARENTHESIS, SU_PAREN_CLOSE, ")"),
		tok(GE_TERMINATOR, SU_NEWLINE, "\n"),
	}

	doErrorTest(t, in)
}

func Test2_5(t *testing.T) {

	// WHEN checking a spell with a missing value delimiter
	// THEN an error should be returned
	// @Println(
	in := []token.Token{
		tok(GE_SPELL, SU_UNDEFINED, "@Println"),
		tok(GE_PARENTHESIS, SU_PAREN_OPEN, "("),
		tok(GE_IDENTIFIER, SU_IDENTIFIER, "x"),
		tok(GE_IDENTIFIER, SU_IDENTIFIER, "y"),
		tok(GE_PARENTHESIS, SU_PAREN_CLOSE, ")"),
		tok(GE_TERMINATOR, SU_NEWLINE, "\n"),
	}

	doErrorTest(t, in)
}

func Test2_6(t *testing.T) {

	// WHEN checking a spell with a missing final terminator
	// THEN an error should be returned
	// @Println(
	in := []token.Token{
		tok(GE_SPELL, SU_UNDEFINED, "@Println"),
		tok(GE_PARENTHESIS, SU_PAREN_OPEN, "("),
		tok(GE_PARENTHESIS, SU_PAREN_CLOSE, ")"),
	}

	doErrorTest(t, in)
}
