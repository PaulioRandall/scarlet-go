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
	)

	// THEN these are the expected instructions
	exp := []inst.Instruction{
		insttest.NewIn(inst.CO_DLM_PUSH, nil),
		insttest.NewIn(inst.CO_SPL_CALL, "Println"),
	}

	doTest(t, in, exp)
}

func Test1_2(t *testing.T) {

	// WHEN compiling a spell with an identifier argument
	// @Println(x)
	in := lextest.Feign(
		lextest.Tok("", lexeme.SPELL),
		lextest.Tok("x", lexeme.IDENT),
		lextest.Tok("@Println", lexeme.SPELL),
	)

	// THEN these are the expected instructions
	exp := []inst.Instruction{
		insttest.NewIn(inst.CO_DLM_PUSH, nil),
		insttest.NewIn(inst.CO_VAL_GET, "x"),
		insttest.NewIn(inst.CO_SPL_CALL, "Println"),
	}

	doTest(t, in, exp)
}

func Test1_3(t *testing.T) {

	// WHEN compiling a spell with a multiple arguments of different types
	// @Println(x, 1, "abc")
	in := lextest.Feign(
		lextest.Tok("", lexeme.SPELL),
		lextest.Tok("x", lexeme.IDENT),
		lextest.Tok(",", lexeme.DELIM),
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok(",", lexeme.DELIM),
		lextest.Tok(`"abc"`, lexeme.STRING),
		lextest.Tok("@Println", lexeme.SPELL),
	)

	// THEN these are the expected instructions
	exp := []inst.Instruction{
		insttest.NewIn(inst.CO_DLM_PUSH, nil),
		insttest.NewIn(inst.CO_VAL_GET, "x"),
		insttest.NewIn(inst.CO_VAL_PUSH, number.New("1")),
		insttest.NewIn(inst.CO_VAL_PUSH, "abc"),
		insttest.NewIn(inst.CO_SPL_CALL, "Println"),
	}

	doTest(t, in, exp)
}

func Test2_1(t *testing.T) {

	// WHEN compiling an assignment
	// 1 := a
	in := lextest.Feign(
		lextest.Tok("", lexeme.ASSIGN),
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok(":=", lexeme.ASSIGN),
		lextest.Tok("x", lexeme.IDENT),
	)

	// THEN these are the expected instructions
	exp := []inst.Instruction{
		insttest.NewIn(inst.CO_VAL_PUSH, number.New("1")),
		insttest.NewIn(inst.CO_VAL_BIND, "x"),
	}

	doTest(t, in, exp)
}

func Test2_2(t *testing.T) {

	// WHEN compiling a multi assignment
	// 1 2 3 := c b a
	in := lextest.Feign(
		lextest.Tok("", lexeme.ASSIGN),
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok(",", lexeme.DELIM),
		lextest.Tok("2", lexeme.NUMBER),
		lextest.Tok(",", lexeme.DELIM),
		lextest.Tok("3", lexeme.NUMBER),
		lextest.Tok(":=", lexeme.ASSIGN),
		lextest.Tok("c", lexeme.IDENT),
		lextest.Tok(",", lexeme.DELIM),
		lextest.Tok("b", lexeme.IDENT),
		lextest.Tok(",", lexeme.DELIM),
		lextest.Tok("a", lexeme.IDENT),
	)

	// THEN these are the expected instructions
	exp := []inst.Instruction{
		insttest.NewIn(inst.CO_VAL_PUSH, number.New("1")),
		insttest.NewIn(inst.CO_VAL_PUSH, number.New("2")),
		insttest.NewIn(inst.CO_VAL_PUSH, number.New("3")),
		insttest.NewIn(inst.CO_VAL_BIND, "c"),
		insttest.NewIn(inst.CO_VAL_BIND, "b"),
		insttest.NewIn(inst.CO_VAL_BIND, "a"),
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
	)

	// THEN these are the expected instructions
	exp := []inst.Instruction{
		insttest.NewIn(inst.CO_DLM_PUSH, nil),
		insttest.NewIn(inst.CO_VAL_PUSH, number.New("1")),
		insttest.NewIn(inst.CO_VAL_PUSH, number.New("2")),
		insttest.NewIn(inst.CO_ADD, nil),
		insttest.NewIn(inst.CO_SPL_CALL, "Println"),
	}

	doTest(t, in, exp)
}

func Test3_4(t *testing.T) {

	// WHEN compiling an assignment with a simple expression as an argument
	// 1 2 + x
	in := lextest.Feign(
		lextest.Tok("", lexeme.ASSIGN),
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok("2", lexeme.NUMBER),
		lextest.Tok("+", lexeme.ADD),
		lextest.Tok(":=", lexeme.ASSIGN),
		lextest.Tok("x", lexeme.IDENT),
	)

	// THEN these are the expected instructions
	exp := []inst.Instruction{
		insttest.NewIn(inst.CO_VAL_PUSH, number.New("1")),
		insttest.NewIn(inst.CO_VAL_PUSH, number.New("2")),
		insttest.NewIn(inst.CO_ADD, nil),
		insttest.NewIn(inst.CO_VAL_BIND, "x"),
	}

	doTest(t, in, exp)
}

func Test3_5(t *testing.T) {

	// WHEN compiling a complex logical expression
	// false false true && || 1 2 < 3 3 >= && 4 5 != && ||
	in := lextest.Feign(
		lextest.Tok("false", lexeme.BOOL),
		lextest.Tok("false", lexeme.BOOL),
		lextest.Tok("true", lexeme.BOOL),
		lextest.Tok("&&", lexeme.AND),
		lextest.Tok("||", lexeme.OR),
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok("2", lexeme.NUMBER),
		lextest.Tok("<", lexeme.LESS),
		lextest.Tok("3", lexeme.NUMBER),
		lextest.Tok("3", lexeme.NUMBER),
		lextest.Tok(">=", lexeme.MORE_EQUAL),
		lextest.Tok("&&", lexeme.AND),
		lextest.Tok("4", lexeme.NUMBER),
		lextest.Tok("5", lexeme.NUMBER),
		lextest.Tok("!=", lexeme.NOT_EQUAL),
		lextest.Tok("&&", lexeme.AND),
		lextest.Tok("||", lexeme.OR),
	)

	// THEN these are the expected instructions
	exp := []inst.Instruction{
		insttest.NewIn(inst.CO_VAL_PUSH, false),
		insttest.NewIn(inst.CO_VAL_PUSH, false),
		insttest.NewIn(inst.CO_VAL_PUSH, true),
		insttest.NewIn(inst.CO_AND, nil),
		insttest.NewIn(inst.CO_OR, nil),
		insttest.NewIn(inst.CO_VAL_PUSH, number.New("1")),
		insttest.NewIn(inst.CO_VAL_PUSH, number.New("2")),
		insttest.NewIn(inst.CO_LESS, nil),
		insttest.NewIn(inst.CO_VAL_PUSH, number.New("3")),
		insttest.NewIn(inst.CO_VAL_PUSH, number.New("3")),
		insttest.NewIn(inst.CO_MORE_EQU, nil),
		insttest.NewIn(inst.CO_AND, nil),
		insttest.NewIn(inst.CO_VAL_PUSH, number.New("4")),
		insttest.NewIn(inst.CO_VAL_PUSH, number.New("5")),
		insttest.NewIn(inst.CO_NOT_EQU, nil),
		insttest.NewIn(inst.CO_AND, nil),
		insttest.NewIn(inst.CO_OR, nil),
	}

	doTest(t, in, exp)
}

func Test4_1(t *testing.T) {

	// WHEN compiling a simple guard with an empty body
	// true GUARD { }
	in := lextest.Feign(
		lextest.Tok("true", lexeme.BOOL),
		lextest.Tok("", lexeme.GUARD),
		lextest.Tok("{", lexeme.L_CURLY),
		lextest.Tok("}", lexeme.R_CURLY),
	)

	// THEN these are the expected instructions
	exp := []inst.Instruction{
		insttest.NewIn(inst.CO_VAL_PUSH, true),
		insttest.NewIn(inst.CO_JMP_FALSE, 2),
		insttest.NewIn(inst.CO_SUB_CTX_PUSH, nil),
		insttest.NewIn(inst.CO_SUB_CTX_POP, nil),
	}

	doTest(t, in, exp)
}

func Test4_2(t *testing.T) {

	// WHEN compiling a simple guard with an single statement body
	// true GUARD { SPELL 1 @Println }
	in := lextest.Feign(
		lextest.Tok("true", lexeme.BOOL),
		lextest.Tok("", lexeme.GUARD),
		lextest.Tok("{", lexeme.L_CURLY),
		/**/ lextest.Tok("", lexeme.SPELL),
		/**/ lextest.Tok("1", lexeme.NUMBER),
		/**/ lextest.Tok("@Println", lexeme.SPELL),
		lextest.Tok("}", lexeme.R_CURLY),
	)

	// THEN these are the expected instructions
	exp := []inst.Instruction{
		insttest.NewIn(inst.CO_VAL_PUSH, true),
		insttest.NewIn(inst.CO_JMP_FALSE, 5),
		insttest.NewIn(inst.CO_SUB_CTX_PUSH, nil),
		insttest.NewIn(inst.CO_DLM_PUSH, nil),
		insttest.NewIn(inst.CO_VAL_PUSH, number.New("1")),
		insttest.NewIn(inst.CO_SPL_CALL, "Println"),
		insttest.NewIn(inst.CO_SUB_CTX_POP, nil),
	}

	doTest(t, in, exp)
}

func Test4_3(t *testing.T) {

	// WHEN compiling nested guards
	// true GUARD { true GUARD { SPELL 1 @Println } SPELL 2 @Println }
	in := lextest.Feign(
		lextest.Tok("true", lexeme.BOOL),
		lextest.Tok("", lexeme.GUARD),
		lextest.Tok("{", lexeme.L_CURLY),
		/**/ lextest.Tok("true", lexeme.BOOL),
		/**/ lextest.Tok("", lexeme.GUARD),
		/**/ lextest.Tok("{", lexeme.L_CURLY),
		/**/ /**/ lextest.Tok("", lexeme.SPELL),
		/**/ /**/ lextest.Tok("1", lexeme.NUMBER),
		/**/ /**/ lextest.Tok("@Println", lexeme.SPELL),
		/**/ lextest.Tok("}", lexeme.R_CURLY),
		/**/ lextest.Tok("", lexeme.SPELL),
		/**/ lextest.Tok("2", lexeme.NUMBER),
		/**/ lextest.Tok("@Println", lexeme.SPELL),
		lextest.Tok("}", lexeme.R_CURLY),
	)

	// THEN these are the expected instructions
	exp := []inst.Instruction{
		insttest.NewIn(inst.CO_VAL_PUSH, true),
		insttest.NewIn(inst.CO_JMP_FALSE, 12),
		insttest.NewIn(inst.CO_SUB_CTX_PUSH, nil),
		insttest.NewIn(inst.CO_VAL_PUSH, true),
		insttest.NewIn(inst.CO_JMP_FALSE, 5),
		insttest.NewIn(inst.CO_SUB_CTX_PUSH, nil),
		insttest.NewIn(inst.CO_DLM_PUSH, nil),
		insttest.NewIn(inst.CO_VAL_PUSH, number.New("1")),
		insttest.NewIn(inst.CO_SPL_CALL, "Println"),
		insttest.NewIn(inst.CO_SUB_CTX_POP, nil),
		insttest.NewIn(inst.CO_DLM_PUSH, nil),
		insttest.NewIn(inst.CO_VAL_PUSH, number.New("2")),
		insttest.NewIn(inst.CO_SPL_CALL, "Println"),
		insttest.NewIn(inst.CO_SUB_CTX_POP, nil),
	}

	doTest(t, in, exp)
}

func Test5_1(t *testing.T) {

	// WHEN compiling a simple loop with an empty body
	// LOOP true { }
	in := lextest.Feign(
		lextest.Tok("", lexeme.LOOP),
		lextest.Tok("true", lexeme.BOOL),
		lextest.Tok("{", lexeme.L_CURLY),
		lextest.Tok("}", lexeme.R_CURLY),
	)

	// THEN these are the expected instructions
	exp := []inst.Instruction{
		insttest.NewIn(inst.CO_VAL_PUSH, true),
		insttest.NewIn(inst.CO_JMP_FALSE, 3),
		insttest.NewIn(inst.CO_SUB_CTX_PUSH, nil),
		insttest.NewIn(inst.CO_SUB_CTX_POP, nil),
		insttest.NewIn(inst.CO_JMP_BACK, 5),
	}

	doTest(t, in, exp)
}

func Test5_2(t *testing.T) {

	// WHEN compiling a loop with a complex condition
	// LOOP true { }
	in := lextest.Feign(
		lextest.Tok("", lexeme.LOOP),
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok("2", lexeme.NUMBER),
		lextest.Tok("<", lexeme.LESS),
		lextest.Tok("{", lexeme.L_CURLY),
		lextest.Tok("}", lexeme.R_CURLY),
	)

	// THEN these are the expected instructions
	exp := []inst.Instruction{
		insttest.NewIn(inst.CO_VAL_PUSH, number.New("1")),
		insttest.NewIn(inst.CO_VAL_PUSH, number.New("2")),
		insttest.NewIn(inst.CO_LESS, nil),
		insttest.NewIn(inst.CO_JMP_FALSE, 3),
		insttest.NewIn(inst.CO_SUB_CTX_PUSH, nil),
		insttest.NewIn(inst.CO_SUB_CTX_POP, nil),
		insttest.NewIn(inst.CO_JMP_BACK, 7),
	}

	doTest(t, in, exp)
}

func Test5_3(t *testing.T) {

	// WHEN compiling nested loops
	// LOOP true { LOOP true { SPELL 1 @Println } }
	in := lextest.Feign(
		lextest.Tok("", lexeme.LOOP),
		lextest.Tok("true", lexeme.BOOL),
		lextest.Tok("{", lexeme.L_CURLY),
		/**/ lextest.Tok("", lexeme.LOOP),
		/**/ lextest.Tok("true", lexeme.BOOL),
		/**/ lextest.Tok("{", lexeme.L_CURLY),
		/**/ /**/ lextest.Tok("", lexeme.SPELL),
		/**/ /**/ lextest.Tok("1", lexeme.NUMBER),
		/**/ /**/ lextest.Tok("@Println", lexeme.SPELL),
		/**/ lextest.Tok("}", lexeme.R_CURLY),
		lextest.Tok("}", lexeme.R_CURLY),
	)

	// THEN these are the expected instructions
	exp := []inst.Instruction{
		insttest.NewIn(inst.CO_VAL_PUSH, true),
		insttest.NewIn(inst.CO_JMP_FALSE, 11),
		insttest.NewIn(inst.CO_SUB_CTX_PUSH, nil),
		/**/ insttest.NewIn(inst.CO_VAL_PUSH, true),
		/**/ insttest.NewIn(inst.CO_JMP_FALSE, 6),
		/**/ insttest.NewIn(inst.CO_SUB_CTX_PUSH, nil),
		/**/ /**/ insttest.NewIn(inst.CO_DLM_PUSH, nil),
		/**/ /**/ insttest.NewIn(inst.CO_VAL_PUSH, number.New("1")),
		/**/ /**/ insttest.NewIn(inst.CO_SPL_CALL, "Println"),
		/**/ insttest.NewIn(inst.CO_SUB_CTX_POP, nil),
		/**/ insttest.NewIn(inst.CO_JMP_BACK, 8),
		insttest.NewIn(inst.CO_SUB_CTX_POP, nil),
		insttest.NewIn(inst.CO_JMP_BACK, 13),
	}

	doTest(t, in, exp)
}
