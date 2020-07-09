package check

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/pkg/rincewind/token"
	. "github.com/PaulioRandall/scarlet-go/pkg/rincewind/token/types"

	pet "github.com/PaulioRandall/scarlet-go/pkg/rincewind/perror/perrortest"
	tkt "github.com/PaulioRandall/scarlet-go/pkg/rincewind/token/tokentest"

	"github.com/stretchr/testify/require"
)

func doTest(t *testing.T, in []token.Token) {

	stream := token.NewStream(in)
	acts := []token.Token{}

	var (
		tk token.Token
		f  CheckFunc
		e  error
	)

	for f = New(stream); f != nil; {
		tk, f, e = f()
		pet.RequireNil(t, e)
		acts = append(acts, tk)
	}

	tkt.RequireSlice(t, in, acts)
}

func doErrorTest(t *testing.T, in []token.Token) {

	itr := token.NewStream(in)

	var e error
	for f := New(itr); f != nil; {
		if _, f, e = f(); e != nil {
			return
		}
	}

	require.Fail(t, "Expected error")
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

	// WHEN checking a spell with no arguments
	// THEN no errors should be returned
	// @Println()
	in := []token.Token{
		halfTok(GE_SPELL, SU_UNDEFINED, "@Print"),
		halfTok(GE_PARENTHESIS, SU_PAREN_OPEN, "("),
		halfTok(GE_PARENTHESIS, SU_PAREN_CLOSE, ")"),
		halfTok(GE_TERMINATOR, SU_NEWLINE, "\n"),
	}

	doTest(t, in)
}

func Test1_2(t *testing.T) {

	// WHEN checking a spell with one argument
	// THEN no errors should be returned
	// @Println(x)
	in := []token.Token{
		halfTok(GE_SPELL, SU_UNDEFINED, "@Print"),
		halfTok(GE_PARENTHESIS, SU_PAREN_OPEN, "("),
		halfTok(GE_IDENTIFIER, SU_IDENTIFIER, "x"),
		halfTok(GE_PARENTHESIS, SU_PAREN_CLOSE, ")"),
		halfTok(GE_TERMINATOR, SU_NEWLINE, "\n"),
	}

	doTest(t, in)
}

func Test1_3(t *testing.T) {

	// WHEN checking a spell with multiple arguments
	// THEN no errors should be returned
	// @Println(x, 1, true)
	in := []token.Token{
		halfTok(GE_SPELL, SU_UNDEFINED, "@Println"),
		halfTok(GE_PARENTHESIS, SU_PAREN_OPEN, "("),
		halfTok(GE_IDENTIFIER, SU_IDENTIFIER, "x"),
		halfTok(GE_DELIMITER, SU_VALUE_DELIM, ","),
		halfTok(GE_LITERAL, SU_NUMBER, "1"),
		halfTok(GE_DELIMITER, SU_VALUE_DELIM, ","),
		halfTok(GE_LITERAL, SU_BOOL, "true"),
		halfTok(GE_PARENTHESIS, SU_PAREN_CLOSE, ")"),
		halfTok(GE_TERMINATOR, SU_NEWLINE, "\n"),
	}

	doTest(t, in)
}

func Test2_1(t *testing.T) {

	// WHEN checking a spell with missing opening parenthesis
	// THEN an error should be returned
	// @Println)
	in := []token.Token{
		halfTok(GE_SPELL, SU_UNDEFINED, "@Println"),
		halfTok(GE_PARENTHESIS, SU_PAREN_CLOSE, ")"),
		halfTok(GE_TERMINATOR, SU_NEWLINE, "\n"),
	}

	doErrorTest(t, in)
}

func Test2_2(t *testing.T) {

	// WHEN checking a spell with missing closing parenthesis
	// THEN an error should be returned
	// @Println(
	in := []token.Token{
		halfTok(GE_SPELL, SU_UNDEFINED, "@Println"),
		halfTok(GE_PARENTHESIS, SU_PAREN_OPEN, "("),
		halfTok(GE_TERMINATOR, SU_NEWLINE, "\n"),
	}

	doErrorTest(t, in)
}

func Test2_3(t *testing.T) {

	// WHEN checking a spell with a stray value delimiter
	// THEN an error should be returned
	// @Println(
	in := []token.Token{
		halfTok(GE_SPELL, SU_UNDEFINED, "@Println"),
		halfTok(GE_PARENTHESIS, SU_PAREN_OPEN, "("),
		halfTok(GE_DELIMITER, SU_VALUE_DELIM, ","),
		halfTok(GE_PARENTHESIS, SU_PAREN_CLOSE, ")"),
		halfTok(GE_TERMINATOR, SU_NEWLINE, "\n"),
	}

	doErrorTest(t, in)
}

func Test2_4(t *testing.T) {

	// WHEN checking a spell with a stray value delimiter
	// THEN an error should be returned
	// @Println(
	in := []token.Token{
		halfTok(GE_SPELL, SU_UNDEFINED, "@Println"),
		halfTok(GE_PARENTHESIS, SU_PAREN_OPEN, "("),
		halfTok(GE_IDENTIFIER, SU_IDENTIFIER, "x"),
		halfTok(GE_DELIMITER, SU_VALUE_DELIM, ","),
		halfTok(GE_PARENTHESIS, SU_PAREN_CLOSE, ")"),
		halfTok(GE_TERMINATOR, SU_NEWLINE, "\n"),
	}

	doErrorTest(t, in)
}

func Test2_5(t *testing.T) {

	// WHEN checking a spell with a missing value delimiter
	// THEN an error should be returned
	// @Println(
	in := []token.Token{
		halfTok(GE_SPELL, SU_UNDEFINED, "@Println"),
		halfTok(GE_PARENTHESIS, SU_PAREN_OPEN, "("),
		halfTok(GE_IDENTIFIER, SU_IDENTIFIER, "x"),
		halfTok(GE_IDENTIFIER, SU_IDENTIFIER, "y"),
		halfTok(GE_PARENTHESIS, SU_PAREN_CLOSE, ")"),
		halfTok(GE_TERMINATOR, SU_NEWLINE, "\n"),
	}

	doErrorTest(t, in)
}

func Test2_6(t *testing.T) {

	// WHEN checking a spell with a missing final terminator
	// THEN an error should be returned
	// @Println(
	in := []token.Token{
		halfTok(GE_SPELL, SU_UNDEFINED, "@Println"),
		halfTok(GE_PARENTHESIS, SU_PAREN_OPEN, "("),
		halfTok(GE_PARENTHESIS, SU_PAREN_CLOSE, ")"),
	}

	doErrorTest(t, in)
}
