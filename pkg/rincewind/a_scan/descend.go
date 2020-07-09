package scan

import (
	"strings"

	"github.com/PaulioRandall/scarlet-go/pkg/rincewind/token"
	. "github.com/PaulioRandall/scarlet-go/pkg/rincewind/token/types"
)

func scanNext(scn *scanner, tk *token.Tok) error {

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
		tk.Gen, tk.Sub, tk.RawStr = GE_IDENTIFIER, SU_VOID, string(scn.next())
		return nil

	case scn.match(';'):
		tk.Gen, tk.Sub, tk.RawStr = GE_TERMINATOR, SU_TERMINATOR, string(scn.next())
		return nil

	case scn.match(','):
		tk.Gen, tk.Sub, tk.RawStr = GE_DELIMITER, SU_VALUE_DELIM, string(scn.next())
		return nil

	case scn.match('('):
		tk.Gen, tk.Sub, tk.RawStr = GE_PARENTHESIS, SU_PAREN_OPEN, string(scn.next())
		return nil

	case scn.match(')'):
		tk.Gen, tk.Sub, tk.RawStr = GE_PARENTHESIS, SU_PAREN_CLOSE, string(scn.next())
		return nil

	case scn.match('"'):
		return stringLiteral(scn, tk)

	case scn.matchDigit():
		return numberLiteral(scn, tk)
	}

	return errorUnknownSymbol(scn)
}

func newline(scn *scanner, tk *token.Tok) error {

	sb := strings.Builder{}

	if scn.match('\r') {
		sb.WriteRune(scn.next())
	}

	if scn.notMatch('\n') {
		return errorBadNewline(scn)
	}
	sb.WriteRune(scn.next())

	tk.Gen, tk.Sub, tk.RawStr = GE_TERMINATOR, SU_NEWLINE, sb.String()
	return nil
}

func whitespace(scn *scanner, tk *token.Tok) error {

	sb := strings.Builder{}

	for !scn.matchNewline() && scn.matchSpace() {
		sb.WriteRune(scn.next())
	}

	tk.Gen, tk.Sub, tk.RawStr = GE_WHITESPACE, SU_UNDEFINED, sb.String()
	return nil
}

func word(scn *scanner, tk *token.Tok) error {

	sb := strings.Builder{}
	sb.WriteRune(scn.next())

	for scn.matchLetter() || scn.match('_') {
		sb.WriteRune(scn.next())
	}

	tk.RawStr = sb.String()

	switch tk.RawStr {
	case "false", "true":
		tk.Gen, tk.Sub = GE_LITERAL, SU_BOOL

	default:
		tk.Gen, tk.Sub = GE_IDENTIFIER, SU_IDENTIFIER
	}

	return nil
}

func spell(scn *scanner, tk *token.Tok) error {

	sb := strings.Builder{}
	sb.WriteRune(scn.next())

	for {
		if !scn.matchLetter() {
			return errorBadSpellName(scn)
		}

		for scn.matchLetter() {
			sb.WriteRune(scn.next())
		}

		if scn.notMatch('.') {
			break
		}

		sb.WriteRune(scn.next())
	}

	tk.Gen, tk.Sub, tk.RawStr = GE_SPELL, SU_UNDEFINED, sb.String()
	return nil
}

func stringLiteral(scn *scanner, tk *token.Tok) error {

	sb := strings.Builder{}
	sb.WriteRune(scn.next())

	for scn.notMatch('"') {

		if scn.match('\\') {
			sb.WriteRune(scn.next())
		}

		if scn.empty() || scn.matchNewline() {
			return errorBadString(scn)
		}

		sb.WriteRune(scn.next())
	}

	sb.WriteRune(scn.next())

	tk.Gen, tk.Sub, tk.RawStr = GE_LITERAL, SU_STRING, sb.String()
	return nil
}

func numberLiteral(scn *scanner, tk *token.Tok) error {

	sb := strings.Builder{}

	for scn.matchDigit() {
		sb.WriteRune(scn.next())
	}

	if scn.notMatch('.') {
		goto FINALISE
	}
	sb.WriteRune(scn.next())

	if !scn.matchDigit() {
		return errorBadNumber(scn)
	}

	for scn.matchDigit() {
		sb.WriteRune(scn.next())
	}

FINALISE:
	tk.Gen, tk.Sub, tk.RawStr = GE_LITERAL, SU_NUMBER, sb.String()
	return nil
}
