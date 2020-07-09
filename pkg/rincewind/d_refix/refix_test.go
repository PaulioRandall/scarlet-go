package refix

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/pkg/rincewind/token"
	. "github.com/PaulioRandall/scarlet-go/pkg/rincewind/token/types"

	pet "github.com/PaulioRandall/scarlet-go/pkg/rincewind/perror/perrortest"
	tkt "github.com/PaulioRandall/scarlet-go/pkg/rincewind/token/tokentest"

	"github.com/stretchr/testify/require"
)

type dummyStream struct {
	tks  []token.Token
	size int
	idx  int
}

func (d *dummyStream) Next() token.Token {

	if d.idx >= d.size {
		return nil
	}

	tk := d.tks[d.idx]
	d.idx++
	return tk
}

func doTest(t *testing.T, in, exps []token.Token) {

	require.NotNil(t, exps, "SANITY CHECK! Expected tokens missing")

	stream := &dummyStream{
		tks:  in,
		size: len(in),
	}
	acts := []token.Token{}

	var (
		tk token.Token
		f  RefixFunc
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

	// WHEN refixing a spell with no arguments
	// @Println()
	in := []token.Token{
		tkt.HalfTok(GE_SPELL, SU_UNDEFINED, "@Print"),
		tkt.HalfTok(GE_PARENTHESIS, SU_PAREN_OPEN, "("),
		tkt.HalfTok(GE_PARENTHESIS, SU_PAREN_CLOSE, ")"),
		tkt.HalfTok(GE_TERMINATOR, SU_NEWLINE, "\n"),
	}

	// THEN parenthesis are removed
	// AND magic token indicating parameter start inserted before spell
	exp := []token.Token{
		tkt.HalfTok(GE_PARAMS, SU_UNDEFINED, "("),
		tkt.HalfTok(GE_SPELL, SU_UNDEFINED, "@Print"),
		tkt.HalfTok(GE_TERMINATOR, SU_NEWLINE, "\n"),
	}

	doTest(t, in, exp)
}

func Test1_2(t *testing.T) {

	// WHEN refixing a spell with one argument
	// @Println(x)
	in := []token.Token{
		tkt.HalfTok(GE_SPELL, SU_UNDEFINED, "@Print"),
		tkt.HalfTok(GE_PARENTHESIS, SU_PAREN_OPEN, "("),
		tkt.HalfTok(GE_IDENTIFIER, SU_IDENTIFIER, "x"),
		tkt.HalfTok(GE_PARENTHESIS, SU_PAREN_CLOSE, ")"),
		tkt.HalfTok(GE_TERMINATOR, SU_NEWLINE, "\n"),
	}

	// THEN parenthesis are removed
	// AND the argument is placed before the spell
	// AND magic token indicating parameter start inserted before the argument
	exp := []token.Token{
		tkt.HalfTok(GE_PARAMS, SU_UNDEFINED, "("),
		tkt.HalfTok(GE_IDENTIFIER, SU_IDENTIFIER, "x"),
		tkt.HalfTok(GE_SPELL, SU_UNDEFINED, "@Print"),
		tkt.HalfTok(GE_TERMINATOR, SU_NEWLINE, "\n"),
	}

	doTest(t, in, exp)
}

func Test1_3(t *testing.T) {

	// WHEN refixing a spell with multiple arguments
	// @Println(x, y, z)
	in := []token.Token{
		tkt.HalfTok(GE_SPELL, SU_UNDEFINED, "@Println"),
		tkt.HalfTok(GE_PARENTHESIS, SU_PAREN_OPEN, "("),
		tkt.HalfTok(GE_IDENTIFIER, SU_IDENTIFIER, "x"),
		tkt.HalfTok(GE_DELIMITER, SU_VALUE_DELIM, ","),
		tkt.HalfTok(GE_IDENTIFIER, SU_IDENTIFIER, "y"),
		tkt.HalfTok(GE_DELIMITER, SU_VALUE_DELIM, ","),
		tkt.HalfTok(GE_IDENTIFIER, SU_IDENTIFIER, "z"),
		tkt.HalfTok(GE_PARENTHESIS, SU_PAREN_CLOSE, ")"),
		tkt.HalfTok(GE_TERMINATOR, SU_NEWLINE, "\n"),
	}

	// THEN parenthesis are removed
	// AND the delimiters are removed
	// AND the arguments are placed before the spell
	// AND magic token indicating parameter start inserted before the arguments
	exp := []token.Token{
		tkt.HalfTok(GE_PARAMS, SU_UNDEFINED, "("),
		tkt.HalfTok(GE_IDENTIFIER, SU_IDENTIFIER, "x"),
		tkt.HalfTok(GE_IDENTIFIER, SU_IDENTIFIER, "y"),
		tkt.HalfTok(GE_IDENTIFIER, SU_IDENTIFIER, "z"),
		tkt.HalfTok(GE_SPELL, SU_UNDEFINED, "@Println"),
		tkt.HalfTok(GE_TERMINATOR, SU_NEWLINE, "\n"),
	}

	doTest(t, in, exp)
}
