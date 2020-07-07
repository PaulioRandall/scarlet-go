package sanitise

import (
	"testing"

	. "github.com/PaulioRandall/scarlet-go/pkg/rincewind/token"

	pet "github.com/PaulioRandall/scarlet-go/pkg/rincewind/perror/perrortest"
	tkt "github.com/PaulioRandall/scarlet-go/pkg/rincewind/token/tokentest"
	"github.com/stretchr/testify/require"
)

type dummyStream struct {
	tks  []Token
	size int
	idx  int
}

func (d *dummyStream) Next() Token {

	if d.idx >= d.size {
		return nil
	}

	tk := d.tks[d.idx]
	d.idx++
	return tk
}

func doTest(t *testing.T, in []Token, exps []Token) {

	require.NotNil(t, exps, "SANITY CHECK! Expected Tokens missing")

	stream := &dummyStream{
		tks:  in,
		size: len(in),
	}
	acts := []Token{}

	var (
		tk Token
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

func Test1_1(t *testing.T) {

	// WHEN sanitising a statement containing redudant whitespace
	// @Println (  )
	in := []Token{
		tkt.HalfTok(GE_SPELL, SU_UNDEFINED, "@Print"),
		tkt.HalfTok(GE_WHITESPACE, SU_UNDEFINED, " "),
		tkt.HalfTok(GE_PARENTHESIS, SU_PAREN_OPEN, "("),
		tkt.HalfTok(GE_WHITESPACE, SU_UNDEFINED, "  "),
		tkt.HalfTok(GE_PARENTHESIS, SU_PAREN_CLOSE, ")"),
	}

	// THEN the whitespace is removed
	exp := []Token{
		tkt.HalfTok(GE_SPELL, SU_UNDEFINED, "@Print"),
		tkt.HalfTok(GE_PARENTHESIS, SU_PAREN_OPEN, "("),
		tkt.HalfTok(GE_PARENTHESIS, SU_PAREN_CLOSE, ")"),
	}

	doTest(t, in, exp)
}

func Test2_1(t *testing.T) {

	// WHEN sanitising a spell call containing a newline after opening parenthesis
	// @Println(
	// )
	in := []Token{
		tkt.HalfTok(GE_SPELL, SU_UNDEFINED, "@Print"),
		tkt.HalfTok(GE_PARENTHESIS, SU_PAREN_OPEN, "("),
		tkt.HalfTok(GE_TERMINATOR, SU_NEWLINE, "\n"),
		tkt.HalfTok(GE_PARENTHESIS, SU_PAREN_CLOSE, ")"),
	}

	// THEN the newline is removed
	exp := []Token{
		tkt.HalfTok(GE_SPELL, SU_UNDEFINED, "@Print"),
		tkt.HalfTok(GE_PARENTHESIS, SU_PAREN_OPEN, "("),
		tkt.HalfTok(GE_PARENTHESIS, SU_PAREN_CLOSE, ")"),
	}

	doTest(t, in, exp)
}

func Test3_1(t *testing.T) {

	// WHEN sanitising a spell call containing a newline after a value delimiter
	// @Println(1,
	// 1)
	in := []Token{
		tkt.HalfTok(GE_SPELL, SU_UNDEFINED, "@Print"),
		tkt.HalfTok(GE_PARENTHESIS, SU_PAREN_OPEN, "("),
		tkt.HalfTok(GE_LITERAL, SU_NUMBER, "1"),
		tkt.HalfTok(GE_DELIMITER, SU_VALUE_DELIM, ","),
		tkt.HalfTok(GE_TERMINATOR, SU_NEWLINE, "\n"),
		tkt.HalfTok(GE_LITERAL, SU_NUMBER, "1"),
		tkt.HalfTok(GE_PARENTHESIS, SU_PAREN_CLOSE, ")"),
	}

	// THEN the newline is removed
	exp := []Token{
		tkt.HalfTok(GE_SPELL, SU_UNDEFINED, "@Print"),
		tkt.HalfTok(GE_PARENTHESIS, SU_PAREN_OPEN, "("),
		tkt.HalfTok(GE_LITERAL, SU_NUMBER, "1"),
		tkt.HalfTok(GE_DELIMITER, SU_VALUE_DELIM, ","),
		tkt.HalfTok(GE_LITERAL, SU_NUMBER, "1"),
		tkt.HalfTok(GE_PARENTHESIS, SU_PAREN_CLOSE, ")"),
	}

	doTest(t, in, exp)
}

func Test3_2(t *testing.T) {

	// WHEN sanitising a spell call containing a newline after a value delimiter
	// AND the next line only contains the closing parenthesis
	// @Println(1,
	// )
	in := []Token{
		tkt.HalfTok(GE_SPELL, SU_UNDEFINED, "@Print"),
		tkt.HalfTok(GE_PARENTHESIS, SU_PAREN_OPEN, "("),
		tkt.HalfTok(GE_LITERAL, SU_NUMBER, "1"),
		tkt.HalfTok(GE_DELIMITER, SU_VALUE_DELIM, ","),
		tkt.HalfTok(GE_TERMINATOR, SU_NEWLINE, "\n"),
		tkt.HalfTok(GE_PARENTHESIS, SU_PAREN_CLOSE, ")"),
	}

	// THEN the newline is removed along with the value delimiter
	exp := []Token{
		tkt.HalfTok(GE_SPELL, SU_UNDEFINED, "@Print"),
		tkt.HalfTok(GE_PARENTHESIS, SU_PAREN_OPEN, "("),
		tkt.HalfTok(GE_LITERAL, SU_NUMBER, "1"),
		tkt.HalfTok(GE_PARENTHESIS, SU_PAREN_CLOSE, ")"),
	}

	doTest(t, in, exp)
}
