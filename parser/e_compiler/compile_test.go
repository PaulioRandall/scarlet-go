package compiler

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/shared/inst"
	"github.com/PaulioRandall/scarlet-go/shared/inst/insttest"
	"github.com/PaulioRandall/scarlet-go/shared/lexeme"
	"github.com/PaulioRandall/scarlet-go/shared/lexeme/lextest"
	"github.com/PaulioRandall/scarlet-go/shared/number"
)

func doTest(t *testing.T, in *lexeme.Container, exps []inst.Instruction) {
	acts := CompileAll(in)
	insttest.Equal(t, exps, acts)
}

func Test1_1(t *testing.T) {

	// WHEN compiling a spell with no arguments
	// @Println()
	in := lextest.Feign(
		lextest.Tok("", lexeme.SPELL),
		lextest.Tok("@Println", lexeme.SPELL),
		lextest.Tok("\n", lexeme.NEWLINE),
	)

	// THEN these are the expected instructions
	exp := []inst.Instruction{
		insttest.NewIn(inst.CO_DELIM_PUSH, nil),
		insttest.NewIn(inst.CO_SPELL, "Println"),
	}

	doTest(t, in, exp)
}

func Test1_2(t *testing.T) {

	// WHEN compiling a spell with an identifier argument
	// @Println(x)
	in := lextest.Feign(
		lextest.Tok("", lexeme.SPELL),
		lextest.Tok("x", lexeme.IDENTIFIER),
		lextest.Tok("@Println", lexeme.SPELL),
		lextest.Tok("\n", lexeme.NEWLINE),
	)

	// THEN these are the expected instructions
	exp := []inst.Instruction{
		insttest.NewIn(inst.CO_DELIM_PUSH, nil),
		insttest.NewIn(inst.CO_CTX_GET, "x"),
		insttest.NewIn(inst.CO_SPELL, "Println"),
	}

	doTest(t, in, exp)
}

func Test1_3(t *testing.T) {

	// WHEN compiling a spell with a multiple arguments of different types
	// @Println(x, 1, "abc")
	in := lextest.Feign(
		lextest.Tok("", lexeme.SPELL),
		lextest.Tok("x", lexeme.IDENTIFIER),
		lextest.Tok(",", lexeme.SEPARATOR),
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok(",", lexeme.SEPARATOR),
		lextest.Tok(`"abc"`, lexeme.STRING),
		lextest.Tok("@Println", lexeme.SPELL),
		lextest.Tok("\n", lexeme.NEWLINE),
	)

	// THEN these are the expected instructions
	exp := []inst.Instruction{
		insttest.NewIn(inst.CO_DELIM_PUSH, nil),
		insttest.NewIn(inst.CO_CTX_GET, "x"),
		insttest.NewIn(inst.CO_VAL_PUSH, number.New("1")),
		insttest.NewIn(inst.CO_VAL_PUSH, "abc"),
		insttest.NewIn(inst.CO_SPELL, "Println"),
	}

	doTest(t, in, exp)
}

func Test2_1(t *testing.T) {

	// WHEN compiling an assignment
	// 1 := a
	in := lextest.Feign(
		lextest.Tok("", lexeme.ASSIGNMENT),
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok(":=", lexeme.ASSIGNMENT),
		lextest.Tok("x", lexeme.IDENTIFIER),
		lextest.Tok("\n", lexeme.NEWLINE),
	)

	// THEN these are the expected instructions
	exp := []inst.Instruction{
		insttest.NewIn(inst.CO_VAL_PUSH, number.New("1")),
		insttest.NewIn(inst.CO_CTX_SET, "x"),
	}

	doTest(t, in, exp)
}

func Test2_2(t *testing.T) {

	// WHEN compiling a multi assignment
	// 1 2 3 := c b a
	in := lextest.Feign(
		lextest.Tok("", lexeme.ASSIGNMENT),
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok(",", lexeme.SEPARATOR),
		lextest.Tok("2", lexeme.NUMBER),
		lextest.Tok(",", lexeme.SEPARATOR),
		lextest.Tok("3", lexeme.NUMBER),
		lextest.Tok(":=", lexeme.ASSIGNMENT),
		lextest.Tok("c", lexeme.IDENTIFIER),
		lextest.Tok(",", lexeme.SEPARATOR),
		lextest.Tok("b", lexeme.IDENTIFIER),
		lextest.Tok(",", lexeme.SEPARATOR),
		lextest.Tok("a", lexeme.IDENTIFIER),
		lextest.Tok("\n", lexeme.NEWLINE),
	)

	// THEN these are the expected instructions
	exp := []inst.Instruction{
		insttest.NewIn(inst.CO_VAL_PUSH, number.New("1")),
		insttest.NewIn(inst.CO_VAL_PUSH, number.New("2")),
		insttest.NewIn(inst.CO_VAL_PUSH, number.New("3")),
		insttest.NewIn(inst.CO_CTX_SET, "c"),
		insttest.NewIn(inst.CO_CTX_SET, "b"),
		insttest.NewIn(inst.CO_CTX_SET, "a"),
	}

	doTest(t, in, exp)
}

func Test3_1(t *testing.T) {

	// WHEN compiling a simple expression
	// 1 2 +
	in := lextest.Feign(
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok("2", lexeme.NUMBER),
		lextest.Tok("+", lexeme.ADD),
		lextest.Tok("\n", lexeme.NEWLINE),
	)

	// THEN these are the expected instructions
	exp := []inst.Instruction{
		insttest.NewIn(inst.CO_VAL_PUSH, number.New("1")),
		insttest.NewIn(inst.CO_VAL_PUSH, number.New("2")),
		insttest.NewIn(inst.CO_ADD, nil),
	}

	doTest(t, in, exp)
}

func Test3_2(t *testing.T) {

	// WHEN compiling a complex arithmetic expression
	// 1 2 3 * 4 / + 5 6 % -
	in := lextest.Feign(
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok("2", lexeme.NUMBER),
		lextest.Tok("3", lexeme.NUMBER),
		lextest.Tok("*", lexeme.MUL),
		lextest.Tok("4", lexeme.NUMBER),
		lextest.Tok("/", lexeme.DIV),
		lextest.Tok("+", lexeme.ADD),
		lextest.Tok("5", lexeme.NUMBER),
		lextest.Tok("6", lexeme.NUMBER),
		lextest.Tok("%", lexeme.REM),
		lextest.Tok("-", lexeme.SUB),
		lextest.Tok("\n", lexeme.NEWLINE),
	)

	// THEN these are the expected instructions
	exp := []inst.Instruction{
		insttest.NewIn(inst.CO_VAL_PUSH, number.New("1")),
		insttest.NewIn(inst.CO_VAL_PUSH, number.New("2")),
		insttest.NewIn(inst.CO_VAL_PUSH, number.New("3")),
		insttest.NewIn(inst.CO_MUL, nil),
		insttest.NewIn(inst.CO_VAL_PUSH, number.New("4")),
		insttest.NewIn(inst.CO_DIV, nil),
		insttest.NewIn(inst.CO_ADD, nil),
		insttest.NewIn(inst.CO_VAL_PUSH, number.New("5")),
		insttest.NewIn(inst.CO_VAL_PUSH, number.New("6")),
		insttest.NewIn(inst.CO_REM, nil),
		insttest.NewIn(inst.CO_SUB, nil),
	}

	doTest(t, in, exp)
}

func Test3_3(t *testing.T) {

	// WHEN compiling a spell with a simple expression as an argument
	// 1 2 + @Println
	in := lextest.Feign(
		lextest.Tok("", lexeme.SPELL),
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok("2", lexeme.NUMBER),
		lextest.Tok("+", lexeme.ADD),
		lextest.Tok("@Println", lexeme.SPELL),
		lextest.Tok("\n", lexeme.NEWLINE),
	)

	// THEN these are the expected instructions
	exp := []inst.Instruction{
		insttest.NewIn(inst.CO_DELIM_PUSH, nil),
		insttest.NewIn(inst.CO_VAL_PUSH, number.New("1")),
		insttest.NewIn(inst.CO_VAL_PUSH, number.New("2")),
		insttest.NewIn(inst.CO_ADD, nil),
		insttest.NewIn(inst.CO_SPELL, "Println"),
	}

	doTest(t, in, exp)
}

func Test3_4(t *testing.T) {

	// WHEN compiling an assignment with a simple expression as an argument
	// 1 2 + x
	in := lextest.Feign(
		lextest.Tok("", lexeme.ASSIGNMENT),
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok("2", lexeme.NUMBER),
		lextest.Tok("+", lexeme.ADD),
		lextest.Tok(":=", lexeme.ASSIGNMENT),
		lextest.Tok("x", lexeme.IDENTIFIER),
		lextest.Tok("\n", lexeme.NEWLINE),
	)

	// THEN these are the expected instructions
	exp := []inst.Instruction{
		insttest.NewIn(inst.CO_VAL_PUSH, number.New("1")),
		insttest.NewIn(inst.CO_VAL_PUSH, number.New("2")),
		insttest.NewIn(inst.CO_ADD, nil),
		insttest.NewIn(inst.CO_CTX_SET, "x"),
	}

	doTest(t, in, exp)
}

func Test3_5(t *testing.T) {

	// WHEN compiling a complex logical expression
	// false false true && || true true && ||
	in := lextest.Feign(
		lextest.Tok("false", lexeme.BOOL),
		lextest.Tok("false", lexeme.BOOL),
		lextest.Tok("true", lexeme.BOOL),
		lextest.Tok("&&", lexeme.AND),
		lextest.Tok("||", lexeme.OR),
		lextest.Tok("true", lexeme.BOOL),
		lextest.Tok("true", lexeme.BOOL),
		lextest.Tok("&&", lexeme.AND),
		lextest.Tok("||", lexeme.OR),
		lextest.Tok("\n", lexeme.NEWLINE),
	)

	// THEN these are the expected instructions
	exp := []inst.Instruction{
		insttest.NewIn(inst.CO_VAL_PUSH, false),
		insttest.NewIn(inst.CO_VAL_PUSH, false),
		insttest.NewIn(inst.CO_VAL_PUSH, true),
		insttest.NewIn(inst.CO_AND, nil),
		insttest.NewIn(inst.CO_OR, nil),
		insttest.NewIn(inst.CO_VAL_PUSH, true),
		insttest.NewIn(inst.CO_VAL_PUSH, true),
		insttest.NewIn(inst.CO_AND, nil),
		insttest.NewIn(inst.CO_OR, nil),
	}

	doTest(t, in, exp)
}
