package compiler

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/pkg/eskarina/code"
	"github.com/PaulioRandall/scarlet-go/pkg/eskarina/inst"
	"github.com/PaulioRandall/scarlet-go/pkg/eskarina/inst/insttest"
	"github.com/PaulioRandall/scarlet-go/pkg/eskarina/lexeme"
	"github.com/PaulioRandall/scarlet-go/pkg/eskarina/lexeme/lextest"
	"github.com/PaulioRandall/scarlet-go/pkg/eskarina/number"
	"github.com/PaulioRandall/scarlet-go/pkg/eskarina/prop"
)

func doTest(t *testing.T, in *lexeme.Lexeme, exps *inst.Instruction) {
	acts := CompileAll(in)
	insttest.Equal(t, exps, acts)
}

func Test1_1(t *testing.T) {

	// WHEN compiling a spell with no arguments
	// @Println()
	in := lextest.Feign(
		lextest.Tok("", prop.PR_CALLABLE),
		lextest.Tok("@Println", prop.PR_SPELL),
		lextest.Tok("\n"),
	)

	// THEN these are the expected instructions
	exp := insttest.Feign(
		insttest.NewIn(code.CO_VAL_PUSH, 0),
		insttest.NewIn(code.CO_SPELL, "Println"),
	)

	doTest(t, in, exp)
}

func Test1_2(t *testing.T) {

	// WHEN compiling a spell with an identifier argument
	// @Println(x)
	in := lextest.Feign(
		lextest.Tok("", prop.PR_CALLABLE),
		lextest.Tok("x", prop.PR_IDENTIFIER),
		lextest.Tok("@Println", prop.PR_SPELL),
		lextest.Tok("\n"),
	)

	// THEN these are the expected instructions
	exp := insttest.Feign(
		insttest.NewIn(code.CO_CTX_GET, "x"),
		insttest.NewIn(code.CO_VAL_PUSH, 1),
		insttest.NewIn(code.CO_SPELL, "Println"),
	)

	doTest(t, in, exp)
}

func Test1_3(t *testing.T) {

	// WHEN compiling a spell with a multiple arguments of different types
	// @Println(x, 1, "abc")
	in := lextest.Feign(
		lextest.Tok("", prop.PR_CALLABLE),
		lextest.Tok("x", prop.PR_IDENTIFIER),
		lextest.Tok("1", prop.PR_LITERAL, prop.PR_NUMBER),
		lextest.Tok(`"abc"`, prop.PR_TERM, prop.PR_LITERAL, prop.PR_STRING),
		lextest.Tok("@Println", prop.PR_SPELL),
		lextest.Tok("\n"),
	)

	// THEN these are the expected instructions
	exp := insttest.Feign(
		insttest.NewIn(code.CO_CTX_GET, "x"),
		insttest.NewIn(code.CO_VAL_PUSH, number.New("1")),
		insttest.NewIn(code.CO_VAL_PUSH, "abc"),
		insttest.NewIn(code.CO_VAL_PUSH, 3),
		insttest.NewIn(code.CO_SPELL, "Println"),
	)

	doTest(t, in, exp)
}
