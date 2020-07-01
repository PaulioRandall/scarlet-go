package scan

import (
	"fmt"
	"strings"

	"github.com/PaulioRandall/scarlet-go/pkg/rincewind/perror"

	. "github.com/PaulioRandall/scarlet-go/pkg/rincewind/token"
)

func scanNext(scn *scanner, tk *tok) error {

	switch {
	case scn.match('\r'), scn.match('\n'):
		return newline(scn, tk)

	case scn.matchSpace():
		return whitespace(scn, tk)

	case scn.matchLetter():
		return word(scn, tk)

	case scn.match('@'):
		return spell(scn, tk)

	case scn.match('_'):
		tk.ge, tk.su, tk.raw = GE_IDENTIFIER, SU_VOID, string(scn.next())
		return nil

	case scn.match(';'):
		tk.ge, tk.su, tk.raw = GE_TERMINATOR, SU_TERMINATOR, string(scn.next())
		return nil

	case scn.match(','):
		tk.ge, tk.su, tk.raw = GE_DELIMITER, SU_VALUE_DELIM, string(scn.next())
		return nil

	case scn.match('('):
		tk.ge, tk.su, tk.raw = GE_PARENTHESIS, SU_PAREN_OPEN, string(scn.next())
		return nil

	case scn.match(')'):
		tk.ge, tk.su, tk.raw = GE_PARENTHESIS, SU_PAREN_CLOSE, string(scn.next())
		return nil

	case scn.match('"'):
		return stringLiteral(scn, tk)

	case scn.matchDigit():
		return numberLiteral(scn, tk)
	}

	return fail(scn, "Unknown symbol %q", scn.next())
}

func newline(scn *scanner, tk *tok) error {

	sb := strings.Builder{}

	if scn.match('\r') {
		sb.WriteRune(scn.next())
	}

	if scn.notMatch('\n') {
		return fail(scn, "Got %q, expected %q", '\r', string("\r\n"))
	}
	sb.WriteRune(scn.next())

	tk.ge, tk.su, tk.raw = GE_TERMINATOR, SU_NEWLINE, sb.String()
	return nil
}

func whitespace(scn *scanner, tk *tok) error {

	sb := strings.Builder{}

	for !scn.matchNewline() && scn.matchSpace() {
		sb.WriteRune(scn.next())
	}

	tk.ge, tk.su, tk.raw = GE_WHITESPACE, SU_UNDEFINED, sb.String()
	return nil
}

func word(scn *scanner, tk *tok) error {

	sb := strings.Builder{}
	sb.WriteRune(scn.next())

	for scn.matchLetter() || scn.match('_') {
		sb.WriteRune(scn.next())
	}

	tk.raw = sb.String()

	switch tk.raw {
	case "false", "true":
		tk.ge, tk.su = GE_LITERAL, SU_BOOL

	default:
		tk.ge, tk.su = GE_IDENTIFIER, SU_IDENTIFIER
	}

	return nil
}

func spell(scn *scanner, tk *tok) error {

	sb := strings.Builder{}
	sb.WriteRune(scn.next())

	for {
		if !scn.matchLetter() {
			return fail(scn, "Expected letter")
		}

		for scn.matchLetter() {
			sb.WriteRune(scn.next())
		}

		if scn.notMatch('.') {
			break
		}

		sb.WriteRune(scn.next())
	}

	tk.ge, tk.su, tk.raw = GE_SPELL, SU_UNDEFINED, sb.String()
	return nil
}

func stringLiteral(scn *scanner, tk *tok) error {

	sb := strings.Builder{}
	sb.WriteRune(scn.next())

	for scn.notMatch('"') {

		if scn.match('\\') {
			sb.WriteRune(scn.next())
		}

		if scn.empty() {
			return fail(scn, "EOF encountered before string was terminated")
		}

		if scn.matchNewline() {
			return fail(scn, "Newline encountered before string was terminated")
		}

		sb.WriteRune(scn.next())
	}

	sb.WriteRune(scn.next())

	tk.ge, tk.su, tk.raw = GE_LITERAL, SU_STRING, sb.String()
	return nil
}

func numberLiteral(scn *scanner, tk *tok) error {

	sb := strings.Builder{}

	for scn.matchDigit() {
		sb.WriteRune(scn.next())
	}

	if scn.notMatch('.') {
		goto FINALISE
	}
	sb.WriteRune(scn.next())

	if !scn.matchDigit() {
		return fail(scn, "Expected digit after decimal point")
	}

	for scn.matchDigit() {
		sb.WriteRune(scn.next())
	}

FINALISE:
	tk.ge, tk.su, tk.raw = GE_LITERAL, SU_NUMBER, sb.String()
	return nil
}

func fail(scn *scanner, msg string, args ...interface{}) error {
	msg = fmt.Sprintf(msg, args...)
	return perror.NewByPos(msg, scn.line, scn.col)
}
