package z_tests

import (
	"testing"

	. "github.com/PaulioRandall/scarlet-go/pkg/z_token"
)

func T1_Newlines(t *testing.T, f ScanFunc) {
	checkOne(t, tok{K_NEWLINE, M_NEWLINE, "\n", 0, 0}, f("\n"))
	checkOne(t, tok{K_NEWLINE, M_NEWLINE, "\r\n", 0, 0}, f("\r\n"))
}

func T2_Whitespace(t *testing.T, f ScanFunc) {
	checkOne(t, tok{K_REDUNDANT, M_WHITESPACE, " \t\r\v\f", 0, 0}, f(" \t\r\v\f"))
}

func T3_Comments(t *testing.T, f ScanFunc) {
	in := "// This is a comment"
	exp := tok{K_REDUNDANT, M_COMMENT, "// This is a comment", 0, 0}
	checkOne(t, exp, f(in))
}

func T4_Match(t *testing.T, f ScanFunc) {
	checkOne(t, tok{K_KEYWORD, M_MATCH, "MATCH", 0, 0}, f("MATCH"))
	checkFirstNot(t, tok{K_KEYWORD, M_MATCH, "MATCH", 0, 0}, f("MATCHH"))
	checkFirstNot(t, tok{K_KEYWORD, M_MATCH, "MATCHH", 0, 0}, f("MATCHH"))
}

func T5_False(t *testing.T, f ScanFunc) {
	checkOne(t, tok{K_LITERAL, M_BOOL, "FALSE", 0, 0}, f("FALSE"))
	checkFirstNot(t, tok{K_LITERAL, M_BOOL, "FALSE", 0, 0}, f("FALSEE"))
	checkFirstNot(t, tok{K_LITERAL, M_BOOL, "FALSEE", 0, 0}, f("FALSEE"))
}

func T6_True(t *testing.T, f ScanFunc) {
	checkOne(t, tok{K_LITERAL, M_BOOL, "TRUE", 0, 0}, f("TRUE"))
	checkFirstNot(t, tok{K_LITERAL, M_BOOL, "TRUE", 0, 0}, f("TRUEE"))
	checkFirstNot(t, tok{K_LITERAL, M_BOOL, "TRUEE", 0, 0}, f("TRUEE"))
}

func T7_List(t *testing.T, f ScanFunc) {
	checkOne(t, tok{K_KEYWORD, M_LIST, "LIST", 0, 0}, f("LIST"))
	checkFirstNot(t, tok{K_KEYWORD, M_LIST, "LIST", 0, 0}, f("LISTT"))
	checkFirstNot(t, tok{K_KEYWORD, M_LIST, "LISTT", 0, 0}, f("LISTT"))
}

func T8_Fix(t *testing.T, f ScanFunc) {
	checkOne(t, tok{K_KEYWORD, M_FIX, "FIX", 0, 0}, f("FIX"))
	checkFirstNot(t, tok{K_KEYWORD, M_FIX, "FIX", 0, 0}, f("FIXX"))
	checkFirstNot(t, tok{K_KEYWORD, M_FIX, "FIXX", 0, 0}, f("FIXX"))
}

func T9_Eof(t *testing.T, f ScanFunc) {
	checkMany(t, []Token{}, f("EOF"))
}

func T10_F(t *testing.T, f ScanFunc) {
	checkOne(t, tok{K_KEYWORD, M_FUNC, "F", 0, 0}, f("F"))
	checkFirstNot(t, tok{K_KEYWORD, M_FUNC, "F", 0, 0}, f("FF"))
	checkFirstNot(t, tok{K_KEYWORD, M_FUNC, "FF", 0, 0}, f("FF"))
}

func T11_Identifiers(t *testing.T, f ScanFunc) {
	checkOne(t, tok{K_IDENTIFIER, M_IDENTIFIER, "a", 0, 0}, f("a"))
	checkOne(t, tok{K_IDENTIFIER, M_IDENTIFIER, "abc", 0, 0}, f("abc"))
	checkOne(t, tok{K_IDENTIFIER, M_IDENTIFIER, "a_c", 0, 0}, f("a_c"))
	checkOne(t, tok{K_IDENTIFIER, M_IDENTIFIER, "ab_", 0, 0}, f("ab_"))
	checkFirstNot(t, tok{K_IDENTIFIER, M_IDENTIFIER, "_", 0, 0}, f("_"))
}

func T12_Assign(t *testing.T, f ScanFunc) {
	checkOne(t, tok{K_DELIMITER, M_ASSIGN, ":=", 0, 0}, f(":="))
}

func T13_Output(t *testing.T, f ScanFunc) {
	checkOne(t, tok{K_KEYWORD, M_OUTPUT, "^", 0, 0}, f("^"))
}

func T14_LessThanOrEqual(t *testing.T, f ScanFunc) {
	checkOne(t, tok{K_COMPARISON, M_LESS_THAN_OR_EQUAL, "<=", 0, 0}, f("<="))
}

func T15_MoreThanOrEqual(t *testing.T, f ScanFunc) {
	checkOne(t, tok{K_COMPARISON, M_MORE_THAN_OR_EQUAL, ">=", 0, 0}, f(">="))
}

func T16_BlockOpen(t *testing.T, f ScanFunc) {
	checkOne(t, tok{K_DELIMITER, M_BLOCK_OPEN, "{", 0, 0}, f("{"))
}

func T17_BlockClose(t *testing.T, f ScanFunc) {
	checkOne(t, tok{K_DELIMITER, M_BLOCK_CLOSE, "}", 0, 0}, f("}"))
}

func T18_ParenOpen(t *testing.T, f ScanFunc) {
	checkOne(t, tok{K_DELIMITER, M_PAREN_OPEN, "(", 0, 0}, f("("))
}

func T19_ParenClose(t *testing.T, f ScanFunc) {
	checkOne(t, tok{K_DELIMITER, M_PAREN_CLOSE, ")", 0, 0}, f(")"))
}

func T20_GuardOpen(t *testing.T, f ScanFunc) {
	checkOne(t, tok{K_DELIMITER, M_GUARD_OPEN, "[", 0, 0}, f("["))
}

func T21_GuardClose(t *testing.T, f ScanFunc) {
	checkOne(t, tok{K_DELIMITER, M_GUARD_CLOSE, "]", 0, 0}, f("]"))
}

func T22_Delim(t *testing.T, f ScanFunc) {
	checkOne(t, tok{K_DELIMITER, M_DELIMITER, ",", 0, 0}, f(","))
}

func T23_Void(t *testing.T, f ScanFunc) {
	checkOne(t, tok{K_IDENTIFIER, M_VOID, "_", 0, 0}, f("_"))
}

func T24_Terminator(t *testing.T, f ScanFunc) {
	checkOne(t, tok{K_DELIMITER, M_TERMINATOR, ";", 0, 0}, f(";"))
}

func T25_Spell(t *testing.T, f ScanFunc) {
	checkOne(t, tok{K_KEYWORD, M_SPELL, "@", 0, 0}, f("@"))
}

func T26_Add(t *testing.T, f ScanFunc) {
	checkOne(t, tok{K_ARITHMETIC, M_ADD, "+", 0, 0}, f("+"))
}

func T27_Subtract(t *testing.T, f ScanFunc) {
	checkOne(t, tok{K_ARITHMETIC, M_SUBTRACT, "-", 0, 0}, f("-"))
}

func T28_Multiply(t *testing.T, f ScanFunc) {
	checkOne(t, tok{K_ARITHMETIC, M_MULTIPLY, "*", 0, 0}, f("*"))
}

func T29_Divide(t *testing.T, f ScanFunc) {
	checkOne(t, tok{K_ARITHMETIC, M_DIVIDE, "/", 0, 0}, f("/"))
}

func T30_Remainder(t *testing.T, f ScanFunc) {
	checkOne(t, tok{K_ARITHMETIC, M_REMAINDER, "%", 0, 0}, f("%"))
}

func T31_And(t *testing.T, f ScanFunc) {
	checkOne(t, tok{K_LOGIC, M_AND, "&", 0, 0}, f("&"))
}

func T32_Or(t *testing.T, f ScanFunc) {
	checkOne(t, tok{K_LOGIC, M_OR, "|", 0, 0}, f("|"))
}

func T33_Equal(t *testing.T, f ScanFunc) {
	checkOne(t, tok{K_COMPARISON, M_EQUAL, "==", 0, 0}, f("=="))
}

func T34_NotEqual(t *testing.T, f ScanFunc) {
	checkOne(t, tok{K_COMPARISON, M_NOT_EQUAL, "!=", 0, 0}, f("!="))
}

func T35_LessThan(t *testing.T, f ScanFunc) {
	checkOne(t, tok{K_COMPARISON, M_LESS_THAN, "<", 0, 0}, f("<"))
}

func T36_MoreThan(t *testing.T, f ScanFunc) {
	checkOne(t, tok{K_COMPARISON, M_MORE_THAN, ">", 0, 0}, f(">"))
}

func T37_String(t *testing.T, f ScanFunc) {
	checkOne(t, tok{K_LITERAL, M_STRING, "``", 0, 0}, f("``"))
	checkOne(t, tok{K_LITERAL, M_STRING, "`abc`", 0, 0}, f("`abc`"))
	checkPanic(t, func() { f("`") })
	checkPanic(t, func() { f("`abc") })
}

func T38_Template(t *testing.T, f ScanFunc) {
	checkOne(t, tok{K_LITERAL, M_TEMPLATE, `""`, 0, 0}, f(`""`))
	checkOne(t, tok{K_LITERAL, M_TEMPLATE, `"abc"`, 0, 0}, f(`"abc"`))
	checkOne(t, tok{K_LITERAL, M_TEMPLATE, `"\""`, 0, 0}, f(`"\""`))
	checkOne(t, tok{K_LITERAL, M_TEMPLATE, `"\\"`, 0, 0}, f(`"\\"`))
	checkOne(t, tok{K_LITERAL, M_TEMPLATE, `"\\\\\\"`, 0, 0}, f(`"\\\\\\"`))
	checkOne(t, tok{K_LITERAL, M_TEMPLATE, `"abc\"abc"`, 0, 0}, f(`"abc\"abc"`))
	checkPanic(t, func() { f(`"`) })
	checkPanic(t, func() { f(`"abc`) })
	checkPanic(t, func() { f(`"\"`) })
	checkPanic(t, func() { f(`"\\`) })
	checkPanic(t, func() { f(`"\\\\\"`) })
}

func T39_Number(t *testing.T, f ScanFunc) {
	checkOne(t, tok{K_LITERAL, M_NUMBER, "1", 0, 0}, f("1"))
	checkOne(t, tok{K_LITERAL, M_NUMBER, "123", 0, 0}, f("123"))
	checkOne(t, tok{K_LITERAL, M_NUMBER, "1.0", 0, 0}, f("1.0"))
	checkOne(t, tok{K_LITERAL, M_NUMBER, "123.456", 0, 0}, f("123.456"))
	checkPanic(t, func() { f("1.") })
}

func T40_Loop(t *testing.T, f ScanFunc) {
	checkOne(t, tok{K_KEYWORD, M_LOOP, "LOOP", 0, 0}, f("LOOP"))
	checkFirstNot(t, tok{K_KEYWORD, M_LOOP, "LOOP", 0, 0}, f("LOOPP"))
	checkFirstNot(t, tok{K_KEYWORD, M_LOOP, "LOOPP", 0, 0}, f("LOOPP"))
}

func T41_Append(t *testing.T, f ScanFunc) {
	checkOne(t, tok{K_REFERENCE, M_LIST_END, ">>", 0, 0}, f(">>"))
	checkFirstNot(t, tok{K_REFERENCE, M_LIST_END, ">>>", 0, 0}, f(">>>"))
}

func T42_Prepend(t *testing.T, f ScanFunc) {
	checkOne(t, tok{K_REFERENCE, M_LIST_START, "<<", 0, 0}, f("<<"))
	checkFirstNot(t, tok{K_REFERENCE, M_LIST_START, "<<<", 0, 0}, f("<<<"))
}
