package tests

import (
	"testing"

	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

func T1_Newlines(t *testing.T, f ScanFunc) {
	checkOne(t, tok(NEWLINE, "\n"), f("\n"))
	checkOne(t, tok(NEWLINE, "\r\n"), f("\r\n"))
}

func T2_Whitespace(t *testing.T, f ScanFunc) {
	checkOne(t, tok(WHITESPACE, " \t\r\v\f"), f(" \t\r\v\f"))
}

func T3_Comments(t *testing.T, f ScanFunc) {
	in := "// This is a comment"
	exp := tok(COMMENT, "// This is a comment")
	checkOne(t, exp, f(in))
}

func T4_When(t *testing.T, f ScanFunc) {
	checkOne(t, tok(WHEN, "WHEN"), f("WHEN"))
	checkFirstNot(t, tok(WHEN, "WHEN"), f("WHENH"))
	checkFirstNot(t, tok(WHEN, "WHENH"), f("WHENH"))
}

func T5_False(t *testing.T, f ScanFunc) {
	checkOne(t, tok(BOOL, "FALSE"), f("FALSE"))
	checkFirstNot(t, tok(BOOL, "FALSE"), f("FALSEE"))
	checkFirstNot(t, tok(BOOL, "FALSEE"), f("FALSEE"))
}

func T6_True(t *testing.T, f ScanFunc) {
	checkOne(t, tok(BOOL, "TRUE"), f("TRUE"))
	checkFirstNot(t, tok(BOOL, "TRUE"), f("TRUEE"))
	checkFirstNot(t, tok(BOOL, "TRUEE"), f("TRUEE"))
}

func T7_List(t *testing.T, f ScanFunc) {
	checkOne(t, tok(LIST, "LIST"), f("LIST"))
	checkFirstNot(t, tok(LIST, "LIST"), f("LISTT"))
	checkFirstNot(t, tok(LIST, "LISTT"), f("LISTT"))
}

func T8_Def(t *testing.T, f ScanFunc) {
	checkOne(t, tok(DEF, "DEF"), f("DEF"))
	checkFirstNot(t, tok(DEF, "DEF"), f("DEFX"))
	checkFirstNot(t, tok(DEF, "DEFX"), f("DEFX"))
}

func T9_Eof(t *testing.T, f ScanFunc) {
	checkMany(t, []Token{}, f("EOF"))
}

func T10_F(t *testing.T, f ScanFunc) {
	checkOne(t, tok(FUNC, "F"), f("F"))
	checkFirstNot(t, tok(FUNC, "F"), f("FF"))
	checkFirstNot(t, tok(FUNC, "FF"), f("FF"))
}

func T11_Identifiers(t *testing.T, f ScanFunc) {
	checkOne(t, tok(IDENTIFIER, "a"), f("a"))
	checkOne(t, tok(IDENTIFIER, "abc"), f("abc"))
	checkOne(t, tok(IDENTIFIER, "a_c"), f("a_c"))
	checkOne(t, tok(IDENTIFIER, "ab_"), f("ab_"))
	checkFirstNot(t, tok(IDENTIFIER, "_"), f("_"))
}

func T12_Assign(t *testing.T, f ScanFunc) {
	checkOne(t, tok(ASSIGN, ":"), f(":"))
}

func T13_Output(t *testing.T, f ScanFunc) {
	checkOne(t, tok(OUTPUT, "^"), f("^"))
}

func T14_LessThanOrEqual(t *testing.T, f ScanFunc) {
	checkOne(t, tok(LESS_THAN_OR_EQUAL, "<="), f("<="))
}

func T15_MoreThanOrEqual(t *testing.T, f ScanFunc) {
	checkOne(t, tok(MORE_THAN_OR_EQUAL, ">="), f(">="))
}

func T16_BlockOpen(t *testing.T, f ScanFunc) {
	checkOne(t, tok(BLOCK_OPEN, "{"), f("{"))
}

func T17_BlockClose(t *testing.T, f ScanFunc) {
	checkOne(t, tok(BLOCK_CLOSE, "}"), f("}"))
}

func T18_ParenOpen(t *testing.T, f ScanFunc) {
	checkOne(t, tok(PAREN_OPEN, "("), f("("))
}

func T19_ParenClose(t *testing.T, f ScanFunc) {
	checkOne(t, tok(PAREN_CLOSE, ")"), f(")"))
}

func T20_GuardOpen(t *testing.T, f ScanFunc) {
	checkOne(t, tok(GUARD_OPEN, "["), f("["))
}

func T21_GuardClose(t *testing.T, f ScanFunc) {
	checkOne(t, tok(GUARD_CLOSE, "]"), f("]"))
}

func T22_Delim(t *testing.T, f ScanFunc) {
	checkOne(t, tok(DELIMITER, ","), f(","))
}

func T23_Void(t *testing.T, f ScanFunc) {
	checkOne(t, tok(VOID, "_"), f("_"))
}

func T24_Terminator(t *testing.T, f ScanFunc) {
	checkOne(t, tok(TERMINATOR, ";"), f(";"))
}

func T25_Spell(t *testing.T, f ScanFunc) {
	checkOne(t, tok(SPELL, "@"), f("@"))
}

func T26_Add(t *testing.T, f ScanFunc) {
	checkOne(t, tok(ADD, "+"), f("+"))
}

func T27_Subtract(t *testing.T, f ScanFunc) {
	checkOne(t, tok(SUBTRACT, "-"), f("-"))
}

func T28_Multiply(t *testing.T, f ScanFunc) {
	checkOne(t, tok(MULTIPLY, "*"), f("*"))
}

func T29_Divide(t *testing.T, f ScanFunc) {
	checkOne(t, tok(DIVIDE, "/"), f("/"))
}

func T30_Remainder(t *testing.T, f ScanFunc) {
	checkOne(t, tok(REMAINDER, "%"), f("%"))
}

func T31_And(t *testing.T, f ScanFunc) {
	checkOne(t, tok(AND, "&"), f("&"))
}

func T32_Or(t *testing.T, f ScanFunc) {
	checkOne(t, tok(OR, "|"), f("|"))
}

func T33_Equal(t *testing.T, f ScanFunc) {
	checkOne(t, tok(EQUAL, "=="), f("=="))
}

func T34_NotEqual(t *testing.T, f ScanFunc) {
	checkOne(t, tok(NOT_EQUAL, "!="), f("!="))
}

func T35_LessThan(t *testing.T, f ScanFunc) {
	checkOne(t, tok(LESS_THAN, "<"), f("<"))
}

func T36_MoreThan(t *testing.T, f ScanFunc) {
	checkOne(t, tok(MORE_THAN, ">"), f(">"))
}

func T37_String(t *testing.T, f ScanFunc) {
	checkOne(t, tok(STRING, `""`), f(`""`))
	checkOne(t, tok(STRING, `"abc"`), f(`"abc"`))
	checkPanic(t, func() { f(`"`) })
	checkPanic(t, func() { f(`"abc`) })
}

func T39_Number(t *testing.T, f ScanFunc) {
	checkOne(t, tok(NUMBER, "1"), f("1"))
	checkOne(t, tok(NUMBER, "123"), f("123"))
	checkOne(t, tok(NUMBER, "1.0"), f("1.0"))
	checkOne(t, tok(NUMBER, "123.456"), f("123.456"))
	checkPanic(t, func() { f("1.") })
}

func T40_Loop(t *testing.T, f ScanFunc) {
	checkOne(t, tok(LOOP, "LOOP"), f("LOOP"))
	checkFirstNot(t, tok(LOOP, "LOOP"), f("LOOPP"))
	checkFirstNot(t, tok(LOOP, "LOOPP"), f("LOOPP"))
}

func T41_Append(t *testing.T, f ScanFunc) {
	checkOne(t, tok(LIST_END, ">>"), f(">>"))
	checkFirstNot(t, tok(LIST_END, ">>>"), f(">>>"))
}

func T42_Prepend(t *testing.T, f ScanFunc) {
	checkOne(t, tok(LIST_START, "<<"), f("<<"))
	checkFirstNot(t, tok(LIST_START, "<<<"), f("<<<"))
}

func T43_Updates(t *testing.T, f ScanFunc) {
	checkOne(t, tok(UPDATES, "<-"), f("<-"))
}

func T44_ExprFunc(t *testing.T, f ScanFunc) {
	checkOne(t, tok(EXPR_FUNC, "E"), f("E"))
	checkFirstNot(t, tok(EXPR_FUNC, "E"), f("EE"))
	checkFirstNot(t, tok(EXPR_FUNC, "EE"), f("EE"))
}
