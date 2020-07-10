package compile

import (
	"fmt"
	"testing"

	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/number"

	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/inst"
	. "github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/inst/codes"

	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/token"
	. "github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/token/types"

	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/testutils"
	"github.com/stretchr/testify/require"
)

func doTest(t *testing.T, in []token.Token, exps []inst.Instruction) {

	acts, e := CompileAll(in)
	if e != nil {
		require.Nil(t, fmt.Sprintf("%+v", e))
	}

	testutils.RequireInstructionSlice(t, exps, acts)
}

func doErrorTest(t *testing.T, in []token.Token) {
	_, e := CompileAll(in)
	require.NotNil(t, e, "Expected an error")
}

func tok(gen GenType, sub SubType, raw string) token.Tok {
	return token.Tok{
		Gen:    gen,
		Sub:    sub,
		RawStr: raw,
	}
}

func instruction(code Code, data interface{}) inst.Instruction {
	return inst.Inst{
		InstCode: code,
		InstData: data,
		Opener:   token.Tok{},
		Closer:   token.Tok{},
	}
}

func Test1_1(t *testing.T) {

	// WHEN compiling a spell with no arguments
	// @Println()
	in := []token.Token{
		tok(GEN_PARAMS, SUB_UNDEFINED, "("),
		tok(GEN_SPELL, SUB_UNDEFINED, "@Println"),
		tok(GEN_TERMINATOR, SUB_NEWLINE, "\n"),
	}

	// THEN these are the expected instructions
	exp := []inst.Instruction{
		instruction(IN_SPELL, []interface{}{0, "Println"}),
	}

	doTest(t, in, exp)
}

func Test1_2(t *testing.T) {

	// WHEN compiling a spell with an identifier argument
	// @Println(x)
	in := []token.Token{
		tok(GEN_PARAMS, SUB_UNDEFINED, "("),
		tok(GEN_IDENTIFIER, SUB_IDENTIFIER, "x"),
		tok(GEN_SPELL, SUB_UNDEFINED, "@Println"),
		tok(GEN_TERMINATOR, SUB_NEWLINE, "\n"),
	}

	// THEN these are the expected instructions
	exp := []inst.Instruction{
		instruction(IN_CTX_GET, "x"),
		instruction(IN_SPELL, []interface{}{1, "Println"}),
	}

	doTest(t, in, exp)
}

func Test1_3(t *testing.T) {

	// WHEN compiling a spell with a multiple arguments of different types
	// @Println(x, 1, "abc")
	in := []token.Token{
		tok(GEN_PARAMS, SUB_UNDEFINED, "("),
		tok(GEN_IDENTIFIER, SUB_IDENTIFIER, "x"),
		tok(GEN_LITERAL, SUB_NUMBER, "1"),
		tok(GEN_LITERAL, SUB_STRING, `"abc"`),
		tok(GEN_SPELL, SUB_UNDEFINED, "@Println"),
		tok(GEN_TERMINATOR, SUB_NEWLINE, "\n"),
	}

	// THEN these are the expected instructions
	exp := []inst.Instruction{
		instruction(IN_CTX_GET, "x"),
		instruction(IN_VAL_PUSH, number.New("1")),
		instruction(IN_VAL_PUSH, "abc"),
		instruction(IN_SPELL, []interface{}{3, "Println"}),
	}

	doTest(t, in, exp)
}
