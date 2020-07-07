package check

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/pkg/rincewind/perror"
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

func doTest(t *testing.T, in []Token) {

	stream := &dummyStream{
		tks:  in,
		size: len(in),
	}
	acts := []Token{}

	var (
		tk Token
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

func doErrorTest(t *testing.T, expCode, in []Token) {

	itr := &dummyStream{
		tks:  in,
		size: len(in),
	}

	var e error
	for f := New(itr); f != nil; {
		if _, f, e = f(); e != nil {

			err := perror.Unwrap(e)
			require.NotNil(t, err,
				"All errors must be an perror.Error or have an perror.Error cause")
			require.Equal(t, expCode, err.Code())
			return
		}
	}

	require.Fail(t, "Expected error")
}

func Test1_1(t *testing.T) {

	// WHEN checking a spell with no arguments
	// THEN no errors should be returned
	// @Println()
	in := []Token{
		tkt.HalfTok(GE_SPELL, SU_UNDEFINED, "@Print"),
		tkt.HalfTok(GE_PARENTHESIS, SU_PAREN_OPEN, "("),
		tkt.HalfTok(GE_PARENTHESIS, SU_PAREN_CLOSE, ")"),
		tkt.HalfTok(GE_TERMINATOR, SU_NEWLINE, "\n"),
	}

	doTest(t, in)
}

func Test1_2(t *testing.T) {

	// WHEN checking a spell with one argument
	// THEN no errors should be returned
	// @Println(x)
	in := []Token{
		tkt.HalfTok(GE_SPELL, SU_UNDEFINED, "@Print"),
		tkt.HalfTok(GE_PARENTHESIS, SU_PAREN_OPEN, "("),
		tkt.HalfTok(GE_IDENTIFIER, SU_IDENTIFIER, "x"),
		tkt.HalfTok(GE_PARENTHESIS, SU_PAREN_CLOSE, ")"),
		tkt.HalfTok(GE_TERMINATOR, SU_NEWLINE, "\n"),
	}

	doTest(t, in)
}

func Test1_3(t *testing.T) {

	// WHEN checking a spell with multiple arguments
	// THEN no errors should be returned
	// @Println(x, y, z)
	in := []Token{
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

	doTest(t, in)
}
