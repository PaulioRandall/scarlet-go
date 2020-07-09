package compile

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/pkg/rincewind/inst"
	"github.com/PaulioRandall/scarlet-go/pkg/rincewind/number"

	"github.com/PaulioRandall/scarlet-go/pkg/rincewind/token"
	. "github.com/PaulioRandall/scarlet-go/pkg/rincewind/token/types"

	ist "github.com/PaulioRandall/scarlet-go/pkg/rincewind/inst/insttest"
	pet "github.com/PaulioRandall/scarlet-go/pkg/rincewind/perror/perrortest"

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

func doTest(t *testing.T, rpn []token.Token, exps []inst.Instruction) {

	require.NotNil(t, exps, "SANITY CHECK! Expected tokens missing")

	stream := &dummyStream{
		tks:  rpn,
		size: len(rpn),
	}

	acts := []inst.Instruction{}

	var (
		in inst.Instruction
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

func halfTok(gen GenType, sub SubType, raw string) token.Tok {
	return token.Tok{
		Gen:    gen,
		Sub:    sub,
		RawStr: raw,
	}
}

func Test1_1(t *testing.T) {

	// WHEN compiling a spell with no arguments
	// @Println()
	in := []token.Token{
		halfTok(GE_PARAMS, SU_UNDEFINED, "("),
		halfTok(GE_SPELL, SU_UNDEFINED, "@Println"),
		halfTok(GE_TERMINATOR, SU_NEWLINE, "\n"),
	}

	// THEN these are the expected instructions
	exp := []inst.Instruction{
		ist.HalfIns(inst.IN_SPELL, []interface{}{0, "Println"}),
	}

	doTest(t, in, exp)
}

func Test1_2(t *testing.T) {

	// WHEN compiling a spell with an identifier argument
	// @Println(x)
	in := []token.Token{
		halfTok(GE_PARAMS, SU_UNDEFINED, "("),
		halfTok(GE_IDENTIFIER, SU_IDENTIFIER, "x"),
		halfTok(GE_SPELL, SU_UNDEFINED, "@Println"),
		halfTok(GE_TERMINATOR, SU_NEWLINE, "\n"),
	}

	// THEN these are the expected instructions
	exp := []inst.Instruction{
		ist.HalfIns(inst.IN_CTX_GET, "x"),
		ist.HalfIns(inst.IN_SPELL, []interface{}{1, "Println"}),
	}

	doTest(t, in, exp)
}

func Test1_3(t *testing.T) {

	// WHEN compiling a spell with a multiple arguments of different types
	// @Println(x, 1, "abc")
	in := []token.Token{
		halfTok(GE_PARAMS, SU_UNDEFINED, "("),
		halfTok(GE_IDENTIFIER, SU_IDENTIFIER, "x"),
		halfTok(GE_LITERAL, SU_NUMBER, "1"),
		halfTok(GE_LITERAL, SU_STRING, `"abc"`),
		halfTok(GE_SPELL, SU_UNDEFINED, "@Println"),
		halfTok(GE_TERMINATOR, SU_NEWLINE, "\n"),
	}

	// THEN these are the expected instructions
	exp := []inst.Instruction{
		ist.HalfIns(inst.IN_CTX_GET, "x"),
		ist.HalfIns(inst.IN_VAL_PUSH, number.New("1")),
		ist.HalfIns(inst.IN_VAL_PUSH, "abc"),
		ist.HalfIns(inst.IN_SPELL, []interface{}{3, "Println"}),
	}

	doTest(t, in, exp)
}
