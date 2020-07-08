package compile

import (
	"testing"

	. "github.com/PaulioRandall/scarlet-go/pkg/rincewind/inst"
	"github.com/PaulioRandall/scarlet-go/pkg/rincewind/number"
	. "github.com/PaulioRandall/scarlet-go/pkg/rincewind/token"

	ist "github.com/PaulioRandall/scarlet-go/pkg/rincewind/inst/insttest"
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

func doTest(t *testing.T, rpn []Token, exps []Instruction) {

	require.NotNil(t, exps, "SANITY CHECK! Expected tokens missing")

	stream := &dummyStream{
		tks:  rpn,
		size: len(rpn),
	}

	acts := []Instruction{}

	var (
		in Instruction
		f  CompileFunc
		e  error
	)

	for f = New(stream); f != nil; {
		in, f, e = f()
		pet.RequireNil(t, e)
		acts = append(acts, in)
	}

	ist.RequireSlice(t, exps, acts)
}

func Test1_1(t *testing.T) {

	// WHEN compiling a spell with no arguments
	// @Println()
	in := []Token{
		tkt.MinTok(GE_PARAMS, SU_UNDEFINED, "("),
		tkt.MinTok(GE_SPELL, SU_UNDEFINED, "@Println"),
		tkt.MinTok(GE_TERMINATOR, SU_NEWLINE, "\n"),
	}

	// THEN these are the expected instructions
	exp := []Instruction{
		ist.HalfIns(IN_SPELL, []interface{}{0, "Println"}),
	}

	doTest(t, in, exp)
}

func Test1_2(t *testing.T) {

	// WHEN compiling a spell with an identifier argument
	// @Println(x)
	in := []Token{
		tkt.MinTok(GE_PARAMS, SU_UNDEFINED, "("),
		tkt.MinTok(GE_IDENTIFIER, SU_IDENTIFIER, "x"),
		tkt.MinTok(GE_SPELL, SU_UNDEFINED, "@Println"),
		tkt.MinTok(GE_TERMINATOR, SU_NEWLINE, "\n"),
	}

	// THEN these are the expected instructions
	exp := []Instruction{
		ist.HalfIns(IN_CTX_GET, "x"),
		ist.HalfIns(IN_SPELL, []interface{}{1, "Println"}),
	}

	doTest(t, in, exp)
}

func Test1_3(t *testing.T) {

	// WHEN compiling a spell with a multiple arguments of diffeerent types
	// @Println(x, 1, "abc")
	in := []Token{
		tkt.MinTok(GE_PARAMS, SU_UNDEFINED, "("),
		tkt.MinTok(GE_IDENTIFIER, SU_IDENTIFIER, "x"),
		tkt.MinTok(GE_LITERAL, SU_NUMBER, "1"),
		tkt.MinTok(GE_LITERAL, SU_STRING, `"abc"`),
		tkt.MinTok(GE_SPELL, SU_UNDEFINED, "@Println"),
		tkt.MinTok(GE_TERMINATOR, SU_NEWLINE, "\n"),
	}

	// THEN these are the expected instructions
	exp := []Instruction{
		ist.HalfIns(IN_CTX_GET, "x"),
		ist.HalfIns(IN_VAL_PUSH, number.New("1")),
		ist.HalfIns(IN_VAL_PUSH, "abc"),
		ist.HalfIns(IN_SPELL, []interface{}{3, "Println"}),
	}

	doTest(t, in, exp)
}
