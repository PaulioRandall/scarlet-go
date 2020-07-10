package sanitise

import (
	"fmt"
	"testing"

	"github.com/PaulioRandall/scarlet-go/pkg/rincewind/shared/token"
	. "github.com/PaulioRandall/scarlet-go/pkg/rincewind/shared/token/types"

	tkt "github.com/PaulioRandall/scarlet-go/pkg/rincewind/shared/token/tokentest"
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

func tok(gen GenType, sub SubType, raw string) token.Tok {
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
		tok(GE_SPELL, SU_UNDEFINED, "@Print"),
		tok(GE_WHITESPACE, SU_UNDEFINED, " "),
		tok(GE_PARENTHESIS, SU_PAREN_OPEN, "("),
		tok(GE_WHITESPACE, SU_UNDEFINED, "  "),
		tok(GE_PARENTHESIS, SU_PAREN_CLOSE, ")"),
	}

	// THEN the whitespace is removed
	exp := []token.Token{
		tok(GE_SPELL, SU_UNDEFINED, "@Print"),
		tok(GE_PARENTHESIS, SU_PAREN_OPEN, "("),
		tok(GE_PARENTHESIS, SU_PAREN_CLOSE, ")"),
	}

	doTest(t, in, exp)
}

func Test2_1(t *testing.T) {

	// WHEN sanitising a spell call containing a newline after opening parenthesis
	// @Println(
	// )
	in := []token.Token{
		tok(GE_SPELL, SU_UNDEFINED, "@Print"),
		tok(GE_PARENTHESIS, SU_PAREN_OPEN, "("),
		tok(GE_TERMINATOR, SU_NEWLINE, "\n"),
		tok(GE_PARENTHESIS, SU_PAREN_CLOSE, ")"),
	}

	// THEN the newline is removed
	exp := []token.Token{
		tok(GE_SPELL, SU_UNDEFINED, "@Print"),
		tok(GE_PARENTHESIS, SU_PAREN_OPEN, "("),
		tok(GE_PARENTHESIS, SU_PAREN_CLOSE, ")"),
	}

	doTest(t, in, exp)
}

func Test3_1(t *testing.T) {

	// WHEN sanitising a spell call containing a newline after a value delimiter
	// @Println(1,
	// 1)
	in := []token.Token{
		tok(GE_SPELL, SU_UNDEFINED, "@Print"),
		tok(GE_PARENTHESIS, SU_PAREN_OPEN, "("),
		tok(GE_LITERAL, SU_NUMBER, "1"),
		tok(GE_DELIMITER, SU_VALUE_DELIM, ","),
		tok(GE_TERMINATOR, SU_NEWLINE, "\n"),
		tok(GE_LITERAL, SU_NUMBER, "1"),
		tok(GE_PARENTHESIS, SU_PAREN_CLOSE, ")"),
	}

	// THEN the newline is removed
	exp := []token.Token{
		tok(GE_SPELL, SU_UNDEFINED, "@Print"),
		tok(GE_PARENTHESIS, SU_PAREN_OPEN, "("),
		tok(GE_LITERAL, SU_NUMBER, "1"),
		tok(GE_DELIMITER, SU_VALUE_DELIM, ","),
		tok(GE_LITERAL, SU_NUMBER, "1"),
		tok(GE_PARENTHESIS, SU_PAREN_CLOSE, ")"),
	}

	doTest(t, in, exp)
}

func Test3_2(t *testing.T) {

	// WHEN sanitising a spell call containing a newline after a value delimiter
	// AND the next line only contains the closing parenthesis
	// @Println(1,
	// )
	in := []token.Token{
		tok(GE_SPELL, SU_UNDEFINED, "@Print"),
		tok(GE_PARENTHESIS, SU_PAREN_OPEN, "("),
		tok(GE_LITERAL, SU_NUMBER, "1"),
		tok(GE_DELIMITER, SU_VALUE_DELIM, ","),
		tok(GE_TERMINATOR, SU_NEWLINE, "\n"),
		tok(GE_PARENTHESIS, SU_PAREN_CLOSE, ")"),
	}

	// THEN the newline is removed along with the value delimiter
	exp := []token.Token{
		tok(GE_SPELL, SU_UNDEFINED, "@Print"),
		tok(GE_PARENTHESIS, SU_PAREN_OPEN, "("),
		tok(GE_LITERAL, SU_NUMBER, "1"),
		tok(GE_PARENTHESIS, SU_PAREN_CLOSE, ")"),
	}

	doTest(t, in, exp)
}
