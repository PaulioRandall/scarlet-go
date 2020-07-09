package compile

import (
	"fmt"
	"testing"

	"github.com/PaulioRandall/scarlet-go/pkg/rincewind/shared/inst"
	"github.com/PaulioRandall/scarlet-go/pkg/rincewind/shared/number"

	"github.com/PaulioRandall/scarlet-go/pkg/rincewind/shared/token"
	. "github.com/PaulioRandall/scarlet-go/pkg/rincewind/shared/token/types"

	ist "github.com/PaulioRandall/scarlet-go/pkg/rincewind/shared/inst/insttest"

	"github.com/stretchr/testify/require"
)

func doTest(t *testing.T, rpn []token.Token, exps []inst.Instruction) {

	var (
		in     inst.Instruction
		e      error
		stream = token.NewStream(rpn)
		acts   = []inst.Instruction{}
	)

	for f := New(stream); f != nil; {
		if in, f, e = f(); e != nil {
			require.NotNil(t, fmt.Sprintf("%+v", e))
		}

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
