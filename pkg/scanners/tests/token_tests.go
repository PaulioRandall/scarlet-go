package tests

import (
	"testing"

	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

func T1_Newlines(t *testing.T, f ScanFunc) {
	checkOne(t, Token{NEWLINE, "\n", 0, 0}, f("\n"))
	checkOne(t, Token{NEWLINE, "\r\n", 0, 0}, f("\r\n"))
}

func T2_Whitespace(t *testing.T, f ScanFunc) {
	checkOne(t, Token{WHITESPACE, " \t\r\v\f", 0, 0}, f(" \t\r\v\f"))
}

func T3_Comments(t *testing.T, f ScanFunc) {
	in := "// This is a comment"
	exp := Token{COMMENT, "// This is a comment", 0, 0}
	checkOne(t, exp, f(in))
}

func T4_Match(t *testing.T, f ScanFunc) {
	checkOne(t, Token{MATCH, "MATCH", 0, 0}, f("MATCH"))
	checkOneNot(t, Token{MATCH, "MATCH", 0, 0}, f("MATCHH"))
	checkOneNot(t, Token{MATCH, "MATCHH", 0, 0}, f("MATCHH"))
}

func T5_False(t *testing.T, f ScanFunc) {
	checkOne(t, Token{BOOL, "FALSE", 0, 0}, f("FALSE"))
	checkOneNot(t, Token{BOOL, "FALSE", 0, 0}, f("FALSEE"))
	checkOneNot(t, Token{BOOL, "FALSEE", 0, 0}, f("FALSEE"))
}

func T6_True(t *testing.T, f ScanFunc) {
	checkOne(t, Token{BOOL, "TRUE", 0, 0}, f("TRUE"))
	checkOneNot(t, Token{BOOL, "TRUE", 0, 0}, f("TRUEE"))
	checkOneNot(t, Token{BOOL, "TRUEE", 0, 0}, f("TRUEE"))
}

func T7_List(t *testing.T, f ScanFunc) {
	checkOne(t, Token{LIST, "LIST", 0, 0}, f("LIST"))
	checkOneNot(t, Token{LIST, "LIST", 0, 0}, f("LISTT"))
	checkOneNot(t, Token{LIST, "LISTT", 0, 0}, f("LISTT"))
}

func T8_Fix(t *testing.T, f ScanFunc) {
	checkOne(t, Token{FIX, "FIX", 0, 0}, f("FIX"))
	checkOneNot(t, Token{FIX, "FIX", 0, 0}, f("FIXX"))
	checkOneNot(t, Token{FIX, "FIXX", 0, 0}, f("FIXX"))
}

func T9_Eof(t *testing.T, f ScanFunc) {
	checkMany(t, []Token{}, f("EOF"))
}

func T10_F(t *testing.T, f ScanFunc) {
	checkOne(t, Token{FUNC, "F", 0, 0}, f("F"))
	checkOneNot(t, Token{FUNC, "F", 0, 0}, f("FF"))
	checkOneNot(t, Token{FUNC, "FF", 0, 0}, f("FF"))
}

func T11_Identifiers(t *testing.T, f ScanFunc) {
	checkOne(t, Token{ID, "a", 0, 0}, f("a"))
	checkOne(t, Token{ID, "abc", 0, 0}, f("abc"))
	checkOne(t, Token{ID, "a_c", 0, 0}, f("a_c"))
	checkOne(t, Token{ID, "ab_", 0, 0}, f("ab_"))
	checkOneNot(t, Token{ID, "_", 0, 0}, f("_"))
}

func T12_Assign(t *testing.T, f ScanFunc) {
	checkOne(t, Token{ASSIGN, ":=", 0, 0}, f(":="))
}

func T13_Output(t *testing.T, f ScanFunc) {
	checkOne(t, Token{OUTPUT, "^", 0, 0}, f("^"))
}

func T14_LessThanOrEqual(t *testing.T, f ScanFunc) {
	checkOne(t, Token{LESS_THAN_OR_EQUAL, "<=", 0, 0}, f("<="))
}

func T15_MoreThanOrEqual(t *testing.T, f ScanFunc) {
	checkOne(t, Token{MORE_THAN_OR_EQUAL, ">=", 0, 0}, f(">="))
}

func T16_BlockOpen(t *testing.T, f ScanFunc) {
	checkOne(t, Token{BLOCK_OPEN, "{", 0, 0}, f("{"))
}

func T17_BlockClose(t *testing.T, f ScanFunc) {
	checkOne(t, Token{BLOCK_CLOSE, "}", 0, 0}, f("}"))
}

func T18_ParenOpen(t *testing.T, f ScanFunc) {
	checkOne(t, Token{PAREN_OPEN, "(", 0, 0}, f("("))
}

func T19_ParenClose(t *testing.T, f ScanFunc) {
	checkOne(t, Token{PAREN_CLOSE, ")", 0, 0}, f(")"))
}

func T20_GuardOpen(t *testing.T, f ScanFunc) {
	checkOne(t, Token{GUARD_OPEN, "[", 0, 0}, f("["))
}

func T21_GuardClose(t *testing.T, f ScanFunc) {
	checkOne(t, Token{GUARD_CLOSE, "]", 0, 0}, f("]"))
}

func T22_Delim(t *testing.T, f ScanFunc) {
	checkOne(t, Token{DELIM, ",", 0, 0}, f(","))
}

func T23_Void(t *testing.T, f ScanFunc) {
	checkOne(t, Token{VOID, "_", 0, 0}, f("_"))
}

func T24_Terminator(t *testing.T, f ScanFunc) {
	checkOne(t, Token{TERMINATOR, ";", 0, 0}, f(";"))
}

func T25_Spell(t *testing.T, f ScanFunc) {
	checkOne(t, Token{SPELL, "@", 0, 0}, f("@"))
}

func T26_Add(t *testing.T, f ScanFunc) {
	checkOne(t, Token{ADD, "+", 0, 0}, f("+"))
}

func T27_Subtract(t *testing.T, f ScanFunc) {
	checkOne(t, Token{SUBTRACT, "-", 0, 0}, f("-"))
}

func T28_Multiply(t *testing.T, f ScanFunc) {
	checkOne(t, Token{MULTIPLY, "*", 0, 0}, f("*"))
}

func T29_Divide(t *testing.T, f ScanFunc) {
	checkOne(t, Token{DIVIDE, "/", 0, 0}, f("/"))
}

func T30_Remainder(t *testing.T, f ScanFunc) {
	checkOne(t, Token{REMAINDER, "%", 0, 0}, f("%"))
}

func T31_And(t *testing.T, f ScanFunc) {
	checkOne(t, Token{AND, "&", 0, 0}, f("&"))
}

func T32_Or(t *testing.T, f ScanFunc) {
	checkOne(t, Token{OR, "|", 0, 0}, f("|"))
}

func T33_Equal(t *testing.T, f ScanFunc) {
	checkOne(t, Token{EQUAL, "==", 0, 0}, f("=="))
}

func T34_NotEqual(t *testing.T, f ScanFunc) {
	checkOne(t, Token{NOT_EQUAL, "!=", 0, 0}, f("!="))
}

func T35_LessThan(t *testing.T, f ScanFunc) {
	checkOne(t, Token{LESS_THAN, "<", 0, 0}, f("<"))
}

func T36_MoreThan(t *testing.T, f ScanFunc) {
	checkOne(t, Token{MORE_THAN, ">", 0, 0}, f(">"))
}

func T37_String(t *testing.T, f ScanFunc) {
	checkOne(t, Token{STRING, "``", 0, 0}, f("``"))
	checkOne(t, Token{STRING, "`abc`", 0, 0}, f("`abc`"))
	checkPanic(t, func() { f("`") })
	checkPanic(t, func() { f("`abc") })
}

func T38_Template(t *testing.T, f ScanFunc) {
	checkOne(t, Token{TEMPLATE, `""`, 0, 0}, f(`""`))
	checkOne(t, Token{TEMPLATE, `"abc"`, 0, 0}, f(`"abc"`))
	checkOne(t, Token{TEMPLATE, `"\""`, 0, 0}, f(`"\""`))
	checkOne(t, Token{TEMPLATE, `"\\"`, 0, 0}, f(`"\\"`))
	checkOne(t, Token{TEMPLATE, `"\\\\\\"`, 0, 0}, f(`"\\\\\\"`))
	checkOne(t, Token{TEMPLATE, `"abc\"abc"`, 0, 0}, f(`"abc\"abc"`))
	checkPanic(t, func() { f(`"`) })
	checkPanic(t, func() { f(`"abc`) })
	checkPanic(t, func() { f(`"\"`) })
	checkPanic(t, func() { f(`"\\`) })
	checkPanic(t, func() { f(`"\\\\\"`) })
}

func T39_Number(t *testing.T, f ScanFunc) {
	checkOne(t, Token{NUMBER, "1", 0, 0}, f("1"))
	checkOne(t, Token{NUMBER, "123", 0, 0}, f("123"))
	checkOne(t, Token{NUMBER, "1.0", 0, 0}, f("1.0"))
	checkOne(t, Token{NUMBER, "123.456", 0, 0}, f("123.456"))
	checkPanic(t, func() { f("1.") })
}

func T40_Loop(t *testing.T, f ScanFunc) {
	checkOne(t, Token{LOOP, "LOOP", 0, 0}, f("LOOP"))
	checkOneNot(t, Token{LOOP, "LOOP", 0, 0}, f("LOOPP"))
	checkOneNot(t, Token{LOOP, "LOOPP", 0, 0}, f("LOOPP"))
}
