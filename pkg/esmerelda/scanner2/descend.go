package scanner2

import (
	"fmt"

	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/err"
	. "github.com/PaulioRandall/scarlet-go/pkg/esmerelda/token"
)

func fail(scn *scanner, msg string) (TokenType, []rune, error) {
	return 0, nil, err.New(msg, err.Pos(scn.line, scn.col))
}

func scan(scn *scanner) (TokenType, []rune, error) {

	switch {
	case scn.match('\r'), scn.match('\n'):
		return newline(scn)

	case scn.matchSpace():
		return whitespace(scn)

	case scn.match('#'):
		return comment(scn)

	case scn.matchLetter():
		return word(scn)

	case scn.match('@'):
		return spell(scn)

	case scn.match(':'):
		return maybeTwoSymbols(scn, TK_THEN, TK_ASSIGNMENT, '=')

	case scn.match('{'):
		return oneSymbol(scn, TK_BLOCK_OPEN)

	case scn.match('}'):
		return oneSymbol(scn, TK_BLOCK_CLOSE)

	case scn.match('('):
		return oneSymbol(scn, TK_PAREN_OPEN)

	case scn.match(')'):
		return oneSymbol(scn, TK_PAREN_CLOSE)

	case scn.match('['):
		return oneSymbol(scn, TK_GUARD_OPEN)

	case scn.match(']'):
		return oneSymbol(scn, TK_GUARD_CLOSE)

	case scn.match(','):
		return oneSymbol(scn, TK_DELIMITER)

	case scn.match('_'):
		return oneSymbol(scn, TK_VOID)

	case scn.match(';'):
		return oneSymbol(scn, TK_TERMINATOR)

	case scn.match('+'):
		return oneSymbol(scn, TK_PLUS)

	case scn.match('-'):
		return maybeTwoSymbols(scn, TK_MINUS, TK_OUTPUTS, '>')

	case scn.match('*'):
		return oneSymbol(scn, TK_MULTIPLY)

	case scn.match('/'):
		return oneSymbol(scn, TK_DIVIDE)

	case scn.match('%'):
		return oneSymbol(scn, TK_REMAINDER)

	case scn.match('&'):
		return twoSymbols(scn, TK_AND, '&')

	case scn.match('|'):
		return twoSymbols(scn, TK_OR, '|')

	case scn.match('<'):
		return maybeTwoSymbols(scn, TK_LESS_THAN, TK_LESS_THAN_OR_EQUAL, '=')

	case scn.match('>'):
		return maybeTwoSymbols(scn, TK_MORE_THAN, TK_MORE_THAN_OR_EQUAL, '=')

	case scn.match('='):
		return twoSymbols(scn, TK_EQUAL, '=')

	case scn.match('!'):
		return twoSymbols(scn, TK_NOT_EQUAL, '=')

	case scn.match('?'):
		return oneSymbol(scn, TK_EXISTS)

	case scn.match('"'):
		return stringLiteral(scn)

	case scn.matchDigit():
		return numberLiteral(scn)
	}

	msg := fmt.Sprintf("Unknown symbol %q", scn.next())
	return fail(scn, msg)
}

func newline(scn *scanner) (TokenType, []rune, error) {

	var r []rune

	if scn.match('\r') {
		r = append(r, scn.next())
	}

	if scn.notMatch('\n') {
		msg := fmt.Sprintf("Got %q, expected %q", '\r', string("\r\n"))
		return fail(scn, msg)
	}

	r = append(r, scn.next())
	return TK_NEWLINE, r, nil
}

func whitespace(scn *scanner) (TokenType, []rune, error) {

	var r []rune

	for !scn.matchNewline() && scn.matchSpace() {
		r = append(r, scn.next())
	}

	return TK_WHITESPACE, r, nil
}

func comment(scn *scanner) (TokenType, []rune, error) {

	var r []rune

	for scn.hasNext() && !scn.matchNewline() {
		r = append(r, scn.next())
	}

	return TK_COMMENT, r, nil
}

func word(scn *scanner) (TokenType, []rune, error) {

	r := []rune{scn.next()}

	for scn.matchLetter() || scn.match('_') {
		r = append(r, scn.next())
	}

	switch string(r) {
	case "false", "true":
		return TK_BOOL, r, nil
	case "def":
		return TK_DEFINITION, r, nil
	case "F":
		return TK_FUNCTION, r, nil
	case "E":
		return TK_EXPR_FUNC, r, nil
	case "when":
		return TK_WHEN, r, nil
	case "loop":
		return TK_LOOP, r, nil
	case "exit":
		return TK_EXIT, r, nil
	}

	return TK_IDENTIFIER, r, nil
}

func spell(scn *scanner) (TokenType, []rune, error) {

	var r []rune

	for {
		r = append(r, scn.next())

		if !scn.matchLetter() {
			return fail(scn, "Expected letter")
		}

		r = append(r, scn.next())
		for scn.matchLetter() {
			r = append(r, scn.next())
		}

		if scn.notMatch('.') {
			break
		}
	}

	return TK_SPELL, r, nil
}

func oneSymbol(scn *scanner, ty TokenType) (TokenType, []rune, error) {
	return ty, []rune{scn.next()}, nil
}

func twoSymbols(scn *scanner, ty TokenType, second rune) (TokenType, []rune, error) {

	first := scn.next()

	if scn.notMatch(second) {
		exp := string([]rune{first, second})
		msg := fmt.Sprintf("Got %q, expected %q", first, exp)
		return fail(scn, msg)
	}

	return ty, []rune{first, scn.next()}, nil
}

func maybeTwoSymbols(scn *scanner, ifOne, ifTwo TokenType, second rune) (TokenType, []rune, error) {

	first := scn.next()
	if scn.notMatch(second) {
		return ifOne, []rune{first}, nil
	}

	return ifTwo, []rune{first, scn.next()}, nil
}

func stringLiteral(scn *scanner) (TokenType, []rune, error) {

	r := []rune{scn.next()}

	for scn.notMatch('"') {

		if scn.match('\\') {
			r = append(r, scn.next())
		}

		if scn.empty() {
			return fail(scn, "EOF encountered before string was terminated")
		}

		if scn.matchNewline() {
			return fail(scn, "Newline encountered before string was terminated")
		}

		r = append(r, scn.next())
	}

	r = append(r, scn.next())
	return TK_STRING, r, nil
}

func numberLiteral(scn *scanner) (TokenType, []rune, error) {

	var r []rune

	for scn.matchDigit() {
		r = append(r, scn.next())
	}

	if scn.empty() || scn.notMatch('.') {
		return TK_NUMBER, r, nil
	}
	r = append(r, scn.next())

	if !scn.matchDigit() {
		return fail(scn, "Expected digit after decimal point")
	}

	for scn.matchDigit() {
		r = append(r, scn.next())
	}

	return TK_NUMBER, r, nil
}
