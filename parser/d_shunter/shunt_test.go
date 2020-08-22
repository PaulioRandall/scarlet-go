package shunter

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/lexeme"
	"github.com/PaulioRandall/scarlet-go/lexeme/lextest"
)

func doTest(t *testing.T, in, exp *lexeme.Container) {
	act := ShuntAll(in)
	lextest.Equal(t, exp.Head(), act.Head())
}

func Test1_1(t *testing.T) {

	// WHEN refixing a spell with no arguments
	// @Println()
	in := lextest.Feign(
		lextest.Tok("@Println", lexeme.SPELL),
		lextest.Tok("(", lexeme.L_PAREN),
		lextest.Tok(")", lexeme.R_PAREN),
		lextest.Tok("\n", lexeme.NEWLINE),
	)

	exp := lextest.Feign(
		lextest.Tok("", lexeme.SPELL),
		lextest.Tok("@Println", lexeme.SPELL),
	)

	doTest(t, in, exp)
}

func Test1_2(t *testing.T) {

	// WHEN refixing a spell with one argument
	// @Println(x)
	in := lextest.Feign(
		lextest.Tok("@Println", lexeme.SPELL),
		lextest.Tok("(", lexeme.L_PAREN),
		lextest.Tok("x", lexeme.IDENT),
		lextest.Tok(")", lexeme.R_PAREN),
		lextest.Tok("\n", lexeme.NEWLINE),
	)

	exp := lextest.Feign(
		lextest.Tok("", lexeme.SPELL),
		lextest.Tok("x", lexeme.IDENT),
		lextest.Tok("@Println", lexeme.SPELL),
	)

	doTest(t, in, exp)
}

func Test1_3(t *testing.T) {

	// WHEN refixing a spell with multiple arguments
	// @Println(x, 1, true)
	in := lextest.Feign(
		lextest.Tok("@Println", lexeme.SPELL),
		lextest.Tok("(", lexeme.L_PAREN),
		lextest.Tok("x", lexeme.IDENT),
		lextest.Tok(",", lexeme.DELIM),
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok(",", lexeme.DELIM),
		lextest.Tok("true", lexeme.BOOL),
		lextest.Tok(")", lexeme.R_PAREN),
		lextest.Tok("\n", lexeme.NEWLINE),
	)

	exp := lextest.Feign(
		lextest.Tok("", lexeme.SPELL),
		lextest.Tok("x", lexeme.IDENT),
		lextest.Tok(",", lexeme.DELIM),
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok(",", lexeme.DELIM),
		lextest.Tok("true", lexeme.BOOL),
		lextest.Tok("@Println", lexeme.SPELL),
	)

	doTest(t, in, exp)
}

func Test2_1(t *testing.T) {

	// WHEN refixing an assignment
	// x := 1
	in := lextest.Feign(
		lextest.Tok("x", lexeme.IDENT),
		lextest.Tok(":=", lexeme.ASSIGN),
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok("\n", lexeme.NEWLINE),
	)

	exp := lextest.Feign(
		lextest.Tok("", lexeme.ASSIGN),
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok(":=", lexeme.ASSIGN),
		lextest.Tok("x", lexeme.IDENT),
	)

	doTest(t, in, exp)
}

func Test2_2(t *testing.T) {

	// WHEN refixing a multi assignment
	// a, b, c := 1, 2, 3
	in := lextest.Feign(
		lextest.Tok("a", lexeme.IDENT),
		lextest.Tok(",", lexeme.DELIM),
		lextest.Tok("b", lexeme.IDENT),
		lextest.Tok(",", lexeme.DELIM),
		lextest.Tok("c", lexeme.IDENT),
		lextest.Tok(":=", lexeme.ASSIGN),
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok(",", lexeme.DELIM),
		lextest.Tok("2", lexeme.NUMBER),
		lextest.Tok(",", lexeme.DELIM),
		lextest.Tok("3", lexeme.NUMBER),
		lextest.Tok("\n", lexeme.NEWLINE),
	)

	exp := lextest.Feign(
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

	doTest(t, in, exp)
}

func Test3_1(t *testing.T) {

	// WHEN refixing a simple expression
	// 1 + 2
	in := lextest.Feign(
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok("+", lexeme.ADD),
		lextest.Tok("2", lexeme.NUMBER),
		lextest.Tok("\n", lexeme.NEWLINE),
	)

	exp := lextest.Feign(
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok("2", lexeme.NUMBER),
		lextest.Tok("+", lexeme.ADD),
	)

	doTest(t, in, exp)
}

func Test3_2(t *testing.T) {

	// WHEN refixing a simple expression containing identifiers
	// a + b
	in := lextest.Feign(
		lextest.Tok("a", lexeme.IDENT),
		lextest.Tok("+", lexeme.ADD),
		lextest.Tok("b", lexeme.IDENT),
		lextest.Tok("\n", lexeme.NEWLINE),
	)

	exp := lextest.Feign(
		lextest.Tok("a", lexeme.IDENT),
		lextest.Tok("b", lexeme.IDENT),
		lextest.Tok("+", lexeme.ADD),
	)

	doTest(t, in, exp)
}

func Test3_3(t *testing.T) {

	// WHEN refixing a complex expression with equal precedence operators
	// 1 + 2 - 3
	in := lextest.Feign(
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok("+", lexeme.ADD),
		lextest.Tok("2", lexeme.NUMBER),
		lextest.Tok("-", lexeme.SUB),
		lextest.Tok("3", lexeme.NUMBER),
		lextest.Tok("\n", lexeme.NEWLINE),
	)

	exp := lextest.Feign(
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok("2", lexeme.NUMBER),
		lextest.Tok("+", lexeme.ADD),
		lextest.Tok("3", lexeme.NUMBER),
		lextest.Tok("-", lexeme.SUB),
	)

	doTest(t, in, exp)
}

func Test3_4(t *testing.T) {

	// WHEN refixing a complex expression with unequal precedence operators
	// 1 + 2 * 3
	in := lextest.Feign(
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok("+", lexeme.ADD),
		lextest.Tok("2", lexeme.NUMBER),
		lextest.Tok("*", lexeme.MUL),
		lextest.Tok("3", lexeme.NUMBER),
		lextest.Tok("\n", lexeme.NEWLINE),
	)

	// 1 2 3 * +
	exp := lextest.Feign(
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok("2", lexeme.NUMBER),
		lextest.Tok("3", lexeme.NUMBER),
		lextest.Tok("*", lexeme.MUL),
		lextest.Tok("+", lexeme.ADD),
	)

	doTest(t, in, exp)
}

func Test3_5(t *testing.T) {

	// WHEN refixing a very complex expression
	// 1 + 2 * 3 / 4 - 5 % 6
	in := lextest.Feign(
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok("+", lexeme.ADD),
		lextest.Tok("2", lexeme.NUMBER),
		lextest.Tok("*", lexeme.MUL),
		lextest.Tok("3", lexeme.NUMBER),
		lextest.Tok("/", lexeme.DIV),
		lextest.Tok("4", lexeme.NUMBER),
		lextest.Tok("-", lexeme.SUB),
		lextest.Tok("5", lexeme.NUMBER),
		lextest.Tok("%", lexeme.REM),
		lextest.Tok("6", lexeme.NUMBER),
		lextest.Tok("\n", lexeme.NEWLINE),
	)

	// 1 2 3 * 4 / + 5 6 % -
	exp := lextest.Feign(
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

	doTest(t, in, exp)
}

func Test3_6(t *testing.T) {

	// WHEN refixing a spell with a single simple expression argument
	// @Println(1 + 2)
	in := lextest.Feign(
		lextest.Tok("@Println", lexeme.SPELL),
		lextest.Tok("(", lexeme.L_PAREN),
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok("+", lexeme.ADD),
		lextest.Tok("2", lexeme.NUMBER),
		lextest.Tok(")", lexeme.R_PAREN),
		lextest.Tok("\n", lexeme.NEWLINE),
	)

	exp := lextest.Feign(
		lextest.Tok("", lexeme.SPELL),
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok("2", lexeme.NUMBER),
		lextest.Tok("+", lexeme.ADD),
		lextest.Tok("@Println", lexeme.SPELL),
	)

	doTest(t, in, exp)
}

func Test3_7(t *testing.T) {

	// WHEN refixing an assignment with a single simple expression
	// x := 1 + 2
	in := lextest.Feign(
		lextest.Tok("x", lexeme.IDENT),
		lextest.Tok(":=", lexeme.ASSIGN),
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok("+", lexeme.ADD),
		lextest.Tok("2", lexeme.NUMBER),
		lextest.Tok("\n", lexeme.NEWLINE),
	)

	exp := lextest.Feign(
		lextest.Tok("", lexeme.ASSIGN),
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok("2", lexeme.NUMBER),
		lextest.Tok("+", lexeme.ADD),
		lextest.Tok(":=", lexeme.ASSIGN),
		lextest.Tok("x", lexeme.IDENT),
	)

	doTest(t, in, exp)
}

func Test3_8(t *testing.T) {

	// WHEN checking a complex logical and relational expression
	// THEN no errors should be returned
	// false || false && true || 1 < 2 && 3 >= 3 && 4 != 5
	in := lextest.Feign(
		lextest.Tok("false", lexeme.BOOL),
		lextest.Tok("||", lexeme.OR),
		lextest.Tok("false", lexeme.BOOL),
		lextest.Tok("&&", lexeme.AND),
		lextest.Tok("true", lexeme.BOOL),
		lextest.Tok("||", lexeme.OR),
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok("<", lexeme.LESS),
		lextest.Tok("2", lexeme.NUMBER),
		lextest.Tok("&&", lexeme.AND),
		lextest.Tok("3", lexeme.NUMBER),
		lextest.Tok(">=", lexeme.MORE_EQUAL),
		lextest.Tok("3", lexeme.NUMBER),
		lextest.Tok("&&", lexeme.AND),
		lextest.Tok("4", lexeme.NUMBER),
		lextest.Tok("!=", lexeme.NOT_EQUAL),
		lextest.Tok("5", lexeme.NUMBER),
		lextest.Tok("\n", lexeme.NEWLINE),
	)

	// false false true && || 1 2 < 3 3 >= && 4 5 != && ||
	exp := lextest.Feign(
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

	doTest(t, in, exp)
}

func Test4_1(t *testing.T) {

	// WHEN refixing an expression containing groups
	// 1 * (2 + 3)
	in := lextest.Feign(
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok("*", lexeme.MUL),
		lextest.Tok("(", lexeme.L_PAREN),
		lextest.Tok("2", lexeme.NUMBER),
		lextest.Tok("+", lexeme.ADD),
		lextest.Tok("3", lexeme.NUMBER),
		lextest.Tok(")", lexeme.R_PAREN),
		lextest.Tok("\n", lexeme.NEWLINE),
	)

	// 1 2 3 + *
	exp := lextest.Feign(
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok("2", lexeme.NUMBER),
		lextest.Tok("3", lexeme.NUMBER),
		lextest.Tok("+", lexeme.ADD),
		lextest.Tok("*", lexeme.MUL),
	)

	doTest(t, in, exp)
}

func Test4_2(t *testing.T) {

	// WHEN refixing an expression containing groups
	// 1 * ((2 + 3))
	in := lextest.Feign(
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok("*", lexeme.MUL),
		lextest.Tok("(", lexeme.L_PAREN),
		lextest.Tok("(", lexeme.L_PAREN),
		lextest.Tok("2", lexeme.NUMBER),
		lextest.Tok("+", lexeme.ADD),
		lextest.Tok("3", lexeme.NUMBER),
		lextest.Tok(")", lexeme.R_PAREN),
		lextest.Tok(")", lexeme.R_PAREN),
		lextest.Tok("\n", lexeme.NEWLINE),
	)

	// 1 2 3 + *
	exp := lextest.Feign(
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok("2", lexeme.NUMBER),
		lextest.Tok("3", lexeme.NUMBER),
		lextest.Tok("+", lexeme.ADD),
		lextest.Tok("*", lexeme.MUL),
	)

	doTest(t, in, exp)
}

func Test4_3(t *testing.T) {

	// WHEN refixing an expression containing groups
	// x := 1 * (2 + 3)
	in := lextest.Feign(
		lextest.Tok("x", lexeme.IDENT),
		lextest.Tok(":=", lexeme.ASSIGN),
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok("*", lexeme.MUL),
		lextest.Tok("(", lexeme.L_PAREN),
		lextest.Tok("2", lexeme.NUMBER),
		lextest.Tok("+", lexeme.ADD),
		lextest.Tok("3", lexeme.NUMBER),
		lextest.Tok(")", lexeme.R_PAREN),
		lextest.Tok("\n", lexeme.NEWLINE),
	)

	// 1 2 3 + *
	exp := lextest.Feign(
		lextest.Tok("", lexeme.ASSIGN),
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok("2", lexeme.NUMBER),
		lextest.Tok("3", lexeme.NUMBER),
		lextest.Tok("+", lexeme.ADD),
		lextest.Tok("*", lexeme.MUL),
		lextest.Tok(":=", lexeme.ASSIGN),
		lextest.Tok("x", lexeme.IDENT),
	)

	doTest(t, in, exp)
}

func Test4_4(t *testing.T) {

	// WHEN refixing an expression containing groups
	// @Println(1 * (2 + 3))
	in := lextest.Feign(
		lextest.Tok("@Println", lexeme.SPELL),
		lextest.Tok("(", lexeme.L_PAREN),
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok("*", lexeme.MUL),
		lextest.Tok("(", lexeme.L_PAREN),
		lextest.Tok("2", lexeme.NUMBER),
		lextest.Tok("+", lexeme.ADD),
		lextest.Tok("3", lexeme.NUMBER),
		lextest.Tok(")", lexeme.R_PAREN),
		lextest.Tok(")", lexeme.R_PAREN),
		lextest.Tok("\n", lexeme.NEWLINE),
	)

	// 1 2 3 + *
	exp := lextest.Feign(
		lextest.Tok("", lexeme.SPELL),
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok("2", lexeme.NUMBER),
		lextest.Tok("3", lexeme.NUMBER),
		lextest.Tok("+", lexeme.ADD),
		lextest.Tok("*", lexeme.MUL),
		lextest.Tok("@Println", lexeme.SPELL),
	)

	doTest(t, in, exp)
}

func Test5_1(t *testing.T) {

	// WHEN refixing a guard statement
	// [true] { @Println(1) }
	in := lextest.Feign(
		lextest.Tok("[", lexeme.L_SQUARE),
		lextest.Tok("true", lexeme.BOOL),
		lextest.Tok("]", lexeme.R_SQUARE),
		lextest.Tok("{", lexeme.L_CURLY),
		lextest.Tok("@Println", lexeme.SPELL),
		lextest.Tok("(", lexeme.L_PAREN),
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok(")", lexeme.R_PAREN),
		lextest.Tok("}", lexeme.R_CURLY),
		lextest.Tok("\n", lexeme.NEWLINE),
	)

	// true GUARD { SPELL 1 @Println }
	exp := lextest.Feign(
		lextest.Tok("true", lexeme.BOOL),
		lextest.Tok("", lexeme.GUARD),
		lextest.Tok("{", lexeme.L_CURLY),
		lextest.Tok("", lexeme.SPELL),
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok("@Println", lexeme.SPELL),
		lextest.Tok("}", lexeme.R_CURLY),
	)

	doTest(t, in, exp)
}

func Test5_2(t *testing.T) {

	// WHEN refixing a guard statement
	// [1 < 2] { @Println(1)
	// @Println(2)
	// @Println(3) }
	in := lextest.Feign(
		lextest.Tok("[", lexeme.L_SQUARE),
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok("<", lexeme.LESS),
		lextest.Tok("2", lexeme.NUMBER),
		lextest.Tok("]", lexeme.R_SQUARE),
		lextest.Tok("{", lexeme.L_CURLY),
		lextest.Tok("@Println", lexeme.SPELL),
		lextest.Tok("(", lexeme.L_PAREN),
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok(")", lexeme.R_PAREN),
		lextest.Tok("\n", lexeme.NEWLINE),
		lextest.Tok("@Println", lexeme.SPELL),
		lextest.Tok("(", lexeme.L_PAREN),
		lextest.Tok("2", lexeme.NUMBER),
		lextest.Tok(")", lexeme.R_PAREN),
		lextest.Tok("\n", lexeme.NEWLINE),
		lextest.Tok("@Println", lexeme.SPELL),
		lextest.Tok("(", lexeme.L_PAREN),
		lextest.Tok("3", lexeme.NUMBER),
		lextest.Tok(")", lexeme.R_PAREN),
		lextest.Tok("}", lexeme.R_CURLY),
		lextest.Tok("\n", lexeme.NEWLINE),
	)

	// 1 2 < GUARD { SPELL 1 @Println SPELL 2 @Println SPELL 3 @Println }
	exp := lextest.Feign(
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok("2", lexeme.NUMBER),
		lextest.Tok("<", lexeme.LESS),
		lextest.Tok("", lexeme.GUARD),
		lextest.Tok("{", lexeme.L_CURLY),
		lextest.Tok("", lexeme.SPELL),
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok("@Println", lexeme.SPELL),
		lextest.Tok("", lexeme.SPELL),
		lextest.Tok("2", lexeme.NUMBER),
		lextest.Tok("@Println", lexeme.SPELL),
		lextest.Tok("", lexeme.SPELL),
		lextest.Tok("3", lexeme.NUMBER),
		lextest.Tok("@Println", lexeme.SPELL),
		lextest.Tok("}", lexeme.R_CURLY),
	)

	doTest(t, in, exp)
}

func Test6_1(t *testing.T) {

	// WHEN refixing a loop statement
	// loop [true] { @Println(1) }
	in := lextest.Feign(
		lextest.Tok("loop", lexeme.LOOP),
		lextest.Tok("[", lexeme.L_SQUARE),
		lextest.Tok("true", lexeme.BOOL),
		lextest.Tok("]", lexeme.R_SQUARE),
		lextest.Tok("{", lexeme.L_CURLY),
		lextest.Tok("@Println", lexeme.SPELL),
		lextest.Tok("(", lexeme.L_PAREN),
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok(")", lexeme.R_PAREN),
		lextest.Tok("}", lexeme.R_CURLY),
		lextest.Tok("\n", lexeme.NEWLINE),
	)

	// loop true { SPELL 1 @Println }
	exp := lextest.Feign(
		lextest.Tok("", lexeme.LOOP),
		lextest.Tok("true", lexeme.BOOL),
		lextest.Tok("{", lexeme.L_CURLY),
		lextest.Tok("", lexeme.SPELL),
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok("@Println", lexeme.SPELL),
		lextest.Tok("}", lexeme.R_CURLY),
	)

	doTest(t, in, exp)
}
