package tests

import (
	"testing"

	. "github.com/PaulioRandall/scarlet-go/pkg/gytha/token"
)

func T1_Newlines(t *testing.T, f ScanFunc) {
	checkOne(t, tok(TK_NEWLINE, "\n"), f("\n"))
	checkOne(t, tok(TK_NEWLINE, "\r\n"), f("\r\n"))
}

func T2_Whitespace(t *testing.T, f ScanFunc) {
	checkOne(t, tok(TK_WHITESPACE, " \t\r\v\f"), f(" \t\r\v\f"))
}

func T3_Comments(t *testing.T, f ScanFunc) {
	in := "// This is a comment"
	exp := tok(TK_COMMENT, "// This is a comment")
	checkOne(t, exp, f(in))
}

func T4_When(t *testing.T, f ScanFunc) {
	checkOne(t, tok(TK_WHEN, "WHEN"), f("WHEN"))
	checkFirstNot(t, tok(TK_WHEN, "WHEN"), f("WHENH"))
	checkFirstNot(t, tok(TK_WHEN, "WHENH"), f("WHENH"))
}

func T5_False(t *testing.T, f ScanFunc) {
	checkOne(t, tok(TK_BOOL, "FALSE"), f("FALSE"))
	checkFirstNot(t, tok(TK_BOOL, "FALSE"), f("FALSEE"))
	checkFirstNot(t, tok(TK_BOOL, "FALSEE"), f("FALSEE"))
}

func T6_True(t *testing.T, f ScanFunc) {
	checkOne(t, tok(TK_BOOL, "TRUE"), f("TRUE"))
	checkFirstNot(t, tok(TK_BOOL, "TRUE"), f("TRUEE"))
	checkFirstNot(t, tok(TK_BOOL, "TRUEE"), f("TRUEE"))
}

func T7_List(t *testing.T, f ScanFunc) {
	checkOne(t, tok(TK_LIST, "LIST"), f("LIST"))
	checkFirstNot(t, tok(TK_LIST, "LIST"), f("LISTT"))
	checkFirstNot(t, tok(TK_LIST, "LISTT"), f("LISTT"))
}

func T8_Def(t *testing.T, f ScanFunc) {
	checkOne(t, tok(TK_DEFINITION, "DEF"), f("DEF"))
	checkFirstNot(t, tok(TK_DEFINITION, "DEF"), f("DEFX"))
	checkFirstNot(t, tok(TK_DEFINITION, "DEFX"), f("DEFX"))
}

func T9_Eof(t *testing.T, f ScanFunc) {
	checkMany(t, []Token{}, f("EOF"))
}

func T10_F(t *testing.T, f ScanFunc) {
	checkOne(t, tok(TK_FUNCTION, "F"), f("F"))
	checkFirstNot(t, tok(TK_FUNCTION, "F"), f("FF"))
	checkFirstNot(t, tok(TK_FUNCTION, "FF"), f("FF"))
}

func T11_Identifiers(t *testing.T, f ScanFunc) {
	checkOne(t, tok(TK_IDENTIFIER, "a"), f("a"))
	checkOne(t, tok(TK_IDENTIFIER, "abc"), f("abc"))
	checkOne(t, tok(TK_IDENTIFIER, "a_c"), f("a_c"))
	checkOne(t, tok(TK_IDENTIFIER, "ab_"), f("ab_"))
	checkFirstNot(t, tok(TK_IDENTIFIER, "_"), f("_"))
}

func T12_Assign(t *testing.T, f ScanFunc) {
	checkOne(t, tok(TK_ASSIGNMENT, ":"), f(":"))
}

func T13_Output(t *testing.T, f ScanFunc) {
	checkOne(t, tok(TK_OUTPUT, "^"), f("^"))
}

func T14_LessThanOrEqual(t *testing.T, f ScanFunc) {
	checkOne(t, tok(TK_LESS_THAN_OR_EQUAL, "<="), f("<="))
}

func T15_MoreThanOrEqual(t *testing.T, f ScanFunc) {
	checkOne(t, tok(TK_MORE_THAN_OR_EQUAL, ">="), f(">="))
}

func T16_BlockOpen(t *testing.T, f ScanFunc) {
	checkOne(t, tok(TK_BLOCK_OPEN, "{"), f("{"))
}

func T17_BlockClose(t *testing.T, f ScanFunc) {
	checkOne(t, tok(TK_BLOCK_CLOSE, "}"), f("}"))
}

func T18_ParenOpen(t *testing.T, f ScanFunc) {
	checkOne(t, tok(TK_PAREN_OPEN, "("), f("("))
}

func T19_ParenClose(t *testing.T, f ScanFunc) {
	checkOne(t, tok(TK_PAREN_CLOSE, ")"), f(")"))
}

func T20_GuardOpen(t *testing.T, f ScanFunc) {
	checkOne(t, tok(TK_GUARD_OPEN, "["), f("["))
}

func T21_GuardClose(t *testing.T, f ScanFunc) {
	checkOne(t, tok(TK_GUARD_CLOSE, "]"), f("]"))
}

func T22_Delim(t *testing.T, f ScanFunc) {
	checkOne(t, tok(TK_DELIMITER, ","), f(","))
}

func T23_Void(t *testing.T, f ScanFunc) {
	checkOne(t, tok(TK_VOID, "_"), f("_"))
}

func T24_Terminator(t *testing.T, f ScanFunc) {
	checkOne(t, tok(TK_TERMINATOR, ";"), f(";"))
}

func T25_Spell(t *testing.T, f ScanFunc) {
	checkOne(t, tok(TK_SPELL, "@"), f("@"))
}

func T26_Add(t *testing.T, f ScanFunc) {
	checkOne(t, tok(TK_PLUS, "+"), f("+"))
}

func T27_Subtract(t *testing.T, f ScanFunc) {
	checkOne(t, tok(TK_MINUS, "-"), f("-"))
}

func T28_Multiply(t *testing.T, f ScanFunc) {
	checkOne(t, tok(TK_MULTIPLY, "*"), f("*"))
}

func T29_Divide(t *testing.T, f ScanFunc) {
	checkOne(t, tok(TK_DIVIDE, "/"), f("/"))
}

func T30_Remainder(t *testing.T, f ScanFunc) {
	checkOne(t, tok(TK_REMAINDER, "%"), f("%"))
}

func T31_And(t *testing.T, f ScanFunc) {
	checkOne(t, tok(TK_AND, "&"), f("&"))
}

func T32_Or(t *testing.T, f ScanFunc) {
	checkOne(t, tok(TK_OR, "|"), f("|"))
}

func T33_Equal(t *testing.T, f ScanFunc) {
	checkOne(t, tok(TK_EQUAL, "=="), f("=="))
}

func T34_NotEqual(t *testing.T, f ScanFunc) {
	checkOne(t, tok(TK_NOT_EQUAL, "!="), f("!="))
}

func T35_LessThan(t *testing.T, f ScanFunc) {
	checkOne(t, tok(TK_LESS_THAN, "<"), f("<"))
}

func T36_MoreThan(t *testing.T, f ScanFunc) {
	checkOne(t, tok(TK_MORE_THAN, ">"), f(">"))
}

func T37_String(t *testing.T, f ScanFunc) {
	checkOne(t, tok(TK_STRING, `""`), f(`""`))
	checkOne(t, tok(TK_STRING, `"abc"`), f(`"abc"`))
	checkPanic(t, func() { f(`"`) })
	checkPanic(t, func() { f(`"abc`) })
}

func T39_Number(t *testing.T, f ScanFunc) {
	checkOne(t, tok(TK_NUMBER, "1"), f("1"))
	checkOne(t, tok(TK_NUMBER, "123"), f("123"))
	checkOne(t, tok(TK_NUMBER, "1.0"), f("1.0"))
	checkOne(t, tok(TK_NUMBER, "123.456"), f("123.456"))
	checkPanic(t, func() { f("1.") })
}

func T40_Loop(t *testing.T, f ScanFunc) {
	checkOne(t, tok(TK_LOOP, "LOOP"), f("LOOP"))
	checkFirstNot(t, tok(TK_LOOP, "LOOP"), f("LOOPP"))
	checkFirstNot(t, tok(TK_LOOP, "LOOPP"), f("LOOPP"))
}

func T41_Append(t *testing.T, f ScanFunc) {
	checkOne(t, tok(TK_LIST_END, ">>"), f(">>"))
	checkFirstNot(t, tok(TK_LIST_END, ">>>"), f(">>>"))
}

func T42_Prepend(t *testing.T, f ScanFunc) {
	checkOne(t, tok(TK_LIST_START, "<<"), f("<<"))
	checkFirstNot(t, tok(TK_LIST_START, "<<<"), f("<<<"))
}

func T43_Updates(t *testing.T, f ScanFunc) {
	checkOne(t, tok(TK_UPDATES, "<-"), f("<-"))
}

func T44_ExprFunc(t *testing.T, f ScanFunc) {
	checkOne(t, tok(TK_EXPR_FUNC, "E"), f("E"))
	checkFirstNot(t, tok(TK_EXPR_FUNC, "E"), f("EE"))
	checkFirstNot(t, tok(TK_EXPR_FUNC, "EE"), f("EE"))
}
