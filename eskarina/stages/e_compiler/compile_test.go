package compiler

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/eskarina/shared/inst"
	"github.com/PaulioRandall/scarlet-go/eskarina/shared/inst/insttest"
	"github.com/PaulioRandall/scarlet-go/eskarina/shared/lexeme"
	"github.com/PaulioRandall/scarlet-go/eskarina/shared/lexeme/lextest"
	"github.com/PaulioRandall/scarlet-go/eskarina/shared/number"
)

func doTest(t *testing.T, in *lexeme.Lexeme, exps *inst.Instruction) {
	acts := CompileAll(in)
	insttest.Equal(t, exps, acts)
}

func Test1_1(t *testing.T) {

	// WHEN compiling a spell with no arguments
	// @Println()
	in := lextest.Feign(
		lextest.Tok("", lexeme.PR_CALLABLE),
		lextest.Tok("@Println", lexeme.PR_SPELL),
		lextest.Tok("\n"),
	)

	// THEN these are the expected instructions
	exp := insttest.Feign(
		insttest.NewIn(inst.CO_VAL_PUSH, 0),
		insttest.NewIn(inst.CO_SPELL, "Println"),
	)

	doTest(t, in, exp)
}

func Test1_2(t *testing.T) {

	// WHEN compiling a spell with an identifier argument
	// @Println(x)
	in := lextest.Feign(
		lextest.Tok("", lexeme.PR_CALLABLE),
		lextest.Tok("x", lexeme.PR_IDENTIFIER),
		lextest.Tok("@Println", lexeme.PR_SPELL),
		lextest.Tok("\n"),
	)

	// THEN these are the expected instructions
	exp := insttest.Feign(
		insttest.NewIn(inst.CO_CTX_GET, "x"),
		insttest.NewIn(inst.CO_VAL_PUSH, 1),
		insttest.NewIn(inst.CO_SPELL, "Println"),
	)

	doTest(t, in, exp)
}

func Test1_3(t *testing.T) {

	// WHEN compiling a spell with a multiple arguments of different types
	// @Println(x, 1, "abc")
	in := lextest.Feign(
		lextest.Tok("", lexeme.PR_CALLABLE),
		lextest.Tok("x", lexeme.PR_IDENTIFIER),
		lextest.Tok("1", lexeme.PR_LITERAL, lexeme.PR_NUMBER),
		lextest.Tok(`"abc"`, lexeme.PR_TERM, lexeme.PR_LITERAL, lexeme.PR_STRING),
		lextest.Tok("@Println", lexeme.PR_SPELL),
		lextest.Tok("\n"),
	)

	// THEN these are the expected instructions
	exp := insttest.Feign(
		insttest.NewIn(inst.CO_CTX_GET, "x"),
		insttest.NewIn(inst.CO_VAL_PUSH, number.New("1")),
		insttest.NewIn(inst.CO_VAL_PUSH, "abc"),
		insttest.NewIn(inst.CO_VAL_PUSH, 3),
		insttest.NewIn(inst.CO_SPELL, "Println"),
	)

	doTest(t, in, exp)
}
