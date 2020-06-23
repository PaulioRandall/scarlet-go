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

	case scn.match('/'):
		return comment(scn)

	case scn.matchLetter():
		return word(scn)

	case scn.match('_'):
		return TK_VOID, []rune{scn.next()}, nil
	}

	msg := fmt.Sprintf("Unknown symbol %q", scn.peek())
	return fail(scn, msg)
}

func newline(scn *scanner) (TokenType, []rune, error) {

	var r []rune

	if scn.match('\r') {
		r = append(r, scn.next())
	}

	if scn.notMatch('\n') {
		msg := fmt.Sprintf("Expected %q after %q", '\n', '\r')
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
