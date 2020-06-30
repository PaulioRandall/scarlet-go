package scan

import (
	"fmt"
	"strings"

	"github.com/PaulioRandall/scarlet-go/pkg/rincewind/number"
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
		tk.ge, tk.su = GE_IDENTIFIER, SU_VOID
		tk.raw, tk.val = string(scn.next()), nil
		return nil

	case scn.match(';'):
		tk.ge, tk.su = GE_TERMINATOR, SU_TERMINATOR
		tk.raw, tk.val = string(scn.next()), nil
		return nil

	case scn.match(','):
		tk.ge, tk.su = GE_DELIMITER, SU_VALUE_DELIM
		tk.raw, tk.val = string(scn.next()), nil
		return nil

	case scn.match('('):
		tk.ge, tk.su = GE_BRACKET, SU_PAREN_OPEN
		tk.raw, tk.val = string(scn.next()), nil
		return nil

	case scn.match(')'):
		tk.ge, tk.su = GE_BRACKET, SU_PAREN_CLOSE
		tk.raw, tk.val = string(scn.next()), nil
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

	tk.ge, tk.su = GE_TERMINATOR, SU_NEWLINE
	tk.raw, tk.val = sb.String(), nil
	return nil
}

func whitespace(scn *scanner, tk *tok) error {

	sb := strings.Builder{}

	for !scn.matchNewline() && scn.matchSpace() {
		sb.WriteRune(scn.next())
	}

	tk.ge, tk.su = GE_WHITESPACE, SU_UNDEFINED
	tk.raw, tk.val = sb.String(), nil
	return nil
}

func word(scn *scanner, tk *tok) error {

	sb := strings.Builder{}
	sb.WriteRune(scn.next())

	for scn.matchLetter() || scn.match('_') {
		sb.WriteRune(scn.next())
	}

	switch sb.String() {
	case "false", "true":
		tk.ge, tk.su = GE_LITERAL, SU_BOOL

	default:
		tk.ge, tk.su = GE_IDENTIFIER, SU_IDENTIFIER
	}

	tk.raw, tk.val = sb.String(), nil
	return nil
}

func spell(scn *scanner, tk *tok) error {

	sb := strings.Builder{}

	for {
		if !scn.matchLetter() {
			return fail(scn, "Expected letter")
		}

		sb.WriteRune(scn.next())

		for scn.matchLetter() {
			sb.WriteRune(scn.next())
		}

		if scn.notMatch('.') {
			break
		}
	}

	tk.ge, tk.su = GE_IDENTIFIER, SU_IDENTIFIER
	tk.raw, tk.val = sb.String(), sb.String()[1:]
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

	size := len(sb.String())
	tk.ge, tk.su = GE_LITERAL, SU_STRING

	if size == 2 {
		tk.raw, tk.val = sb.String(), ""
	} else {
		tk.raw, tk.val = sb.String(), sb.String()[1:size-1]
	}

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
	tk.ge, tk.su = GE_LITERAL, SU_NUMBER
	tk.raw, tk.val = sb.String(), number.New(sb.String())
	return nil
}

func fail(scn *scanner, msg string, args ...interface{}) error {
	msg = fmt.Sprintf(msg, args...)
	return perror.NewByPos(msg, scn.line, scn.col)
}
