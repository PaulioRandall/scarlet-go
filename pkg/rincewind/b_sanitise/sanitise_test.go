package sanitise

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/pkg/rincewind/token"
	. "github.com/PaulioRandall/scarlet-go/pkg/rincewind/token/types"

	pet "github.com/PaulioRandall/scarlet-go/pkg/rincewind/perror/perrortest"
	tkt "github.com/PaulioRandall/scarlet-go/pkg/rincewind/token/tokentest"
	"github.com/stretchr/testify/require"
)

func doTest(t *testing.T, in []token.Token, exps []token.Token) {

	require.NotNil(t, exps, "SANITY CHECK! Expected Tokens missing")

	stream := token.NewStream(in)
	acts := []token.Token{}

	var (
		tk token.Token
		f  SanitiseFunc
		e  error
	)

	for f = New(stream); f != nil; {
		tk, f, e = f()
		pet.RequireNil(t, e)
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

	// WHEN sanitising a statement containing redudant whitespace
	// @Println (  )
	in := []token.Token{
		halfTok(GE_SPELL, SU_UNDEFINED, "@Print"),
		halfTok(GE_WHITESPACE, SU_UNDEFINED, " "),
		halfTok(GE_PARENTHESIS, SU_PAREN_OPEN, "("),
		halfTok(GE_WHITESPACE, SU_UNDEFINED, "  "),
		halfTok(GE_PARENTHESIS, SU_PAREN_CLOSE, ")"),
	}

	// THEN the whitespace is removed
	exp := []token.Token{
		halfTok(GE_SPELL, SU_UNDEFINED, "@Print"),
		halfTok(GE_PARENTHESIS, SU_PAREN_OPEN, "("),
		halfTok(GE_PARENTHESIS, SU_PAREN_CLOSE, ")"),
	}

	doTest(t, in, exp)
}

func Test2_1(t *testing.T) {

	// WHEN sanitising a spell call containing a newline after opening parenthesis
	// @Println(
	// )
	in := []token.Token{
		halfTok(GE_SPELL, SU_UNDEFINED, "@Print"),
		halfTok(GE_PARENTHESIS, SU_PAREN_OPEN, "("),
		halfTok(GE_TERMINATOR, SU_NEWLINE, "\n"),
		halfTok(GE_PARENTHESIS, SU_PAREN_CLOSE, ")"),
	}

	// THEN the newline is removed
	exp := []token.Token{
		halfTok(GE_SPELL, SU_UNDEFINED, "@Print"),
		halfTok(GE_PARENTHESIS, SU_PAREN_OPEN, "("),
		halfTok(GE_PARENTHESIS, SU_PAREN_CLOSE, ")"),
	}

	doTest(t, in, exp)
}

func Test3_1(t *testing.T) {

	// WHEN sanitising a spell call containing a newline after a value delimiter
	// @Println(1,
	// 1)
	in := []token.Token{
		halfTok(GE_SPELL, SU_UNDEFINED, "@Print"),
		halfTok(GE_PARENTHESIS, SU_PAREN_OPEN, "("),
		halfTok(GE_LITERAL, SU_NUMBER, "1"),
		halfTok(GE_DELIMITER, SU_VALUE_DELIM, ","),
		halfTok(GE_TERMINATOR, SU_NEWLINE, "\n"),
		halfTok(GE_LITERAL, SU_NUMBER, "1"),
		halfTok(GE_PARENTHESIS, SU_PAREN_CLOSE, ")"),
	}

	// THEN the newline is removed
	exp := []token.Token{
		halfTok(GE_SPELL, SU_UNDEFINED, "@Print"),
		halfTok(GE_PARENTHESIS, SU_PAREN_OPEN, "("),
		halfTok(GE_LITERAL, SU_NUMBER, "1"),
		halfTok(GE_DELIMITER, SU_VALUE_DELIM, ","),
		halfTok(GE_LITERAL, SU_NUMBER, "1"),
		halfTok(GE_PARENTHESIS, SU_PAREN_CLOSE, ")"),
	}

	doTest(t, in, exp)
}

func Test3_2(t *testing.T) {

	// WHEN sanitising a spell call containing a newline after a value delimiter
	// AND the next line only contains the closing parenthesis
	// @Println(1,
	// )
	in := []token.Token{
		halfTok(GE_SPELL, SU_UNDEFINED, "@Print"),
		halfTok(GE_PARENTHESIS, SU_PAREN_OPEN, "("),
		halfTok(GE_LITERAL, SU_NUMBER, "1"),
		halfTok(GE_DELIMITER, SU_VALUE_DELIM, ","),
		halfTok(GE_TERMINATOR, SU_NEWLINE, "\n"),
		halfTok(GE_PARENTHESIS, SU_PAREN_CLOSE, ")"),
	}

	// THEN the newline is removed along with the value delimiter
	exp := []token.Token{
		halfTok(GE_SPELL, SU_UNDEFINED, "@Print"),
		halfTok(GE_PARENTHESIS, SU_PAREN_OPEN, "("),
		halfTok(GE_LITERAL, SU_NUMBER, "1"),
		halfTok(GE_PARENTHESIS, SU_PAREN_CLOSE, ")"),
	}

	doTest(t, in, exp)
}
