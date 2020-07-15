package scan

import (
	"strings"

	. "github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/prop"
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/token"
	. "github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/token/types"
)

func next(scn *scanner, tk *token.Tok) error {

	switch {
	case scn.match('#'):
		return comment(scn, tk)

	case scn.match('\r'), scn.match('\n'):
		return newline(scn, tk)

	case scn.matchSpace():
		return whitespace(scn, tk)

	case scn.matchLetter():
		return word(scn, tk)

	case scn.match('@'):
		return spell(scn, tk)

	case scn.match('_'):
		tk.Gen, tk.Sub = GEN_IDENTIFIER, SUB_VOID
		tk.RawProps = []Prop{PR_ASSIGNEE, PR_VOID}
		tk.RawStr = string(scn.next())
		return nil

	case scn.match(';'):
		tk.Gen, tk.Sub = GEN_TERMINATOR, SUB_TERMINATOR
		tk.RawProps = []Prop{PR_TERMINATOR}
		tk.RawStr = string(scn.next())
		return nil

	case scn.match(','):
		tk.Gen, tk.Sub = GEN_DELIMITER, SUB_VALUE_DELIM
		tk.RawProps = []Prop{PR_DELIMITER, PR_SEPARATOR}
		tk.RawStr = string(scn.next())
		return nil

	case scn.match('('):
		tk.Gen, tk.Sub = GEN_PARENTHESIS, SUB_PAREN_OPEN
		tk.RawProps = []Prop{PR_DELIMITER, PR_PARENTHESIS, PR_OPENER}
		tk.RawStr = string(scn.next())
		return nil

	case scn.match(')'):
		tk.Gen, tk.Sub = GEN_PARENTHESIS, SUB_PAREN_CLOSE
		tk.RawProps = []Prop{PR_DELIMITER, PR_PARENTHESIS, PR_CLOSER}
		tk.RawStr = string(scn.next())
		return nil

	case scn.match('"'):
		return stringLiteral(scn, tk)

	case scn.matchDigit():
		return numberLiteral(scn, tk)
	}

	return errorUnknownSymbol(scn)
}

func comment(scn *scanner, tk *token.Tok) error {

	sb := strings.Builder{}

	for !scn.empty() && !scn.matchNewline() {
		sb.WriteRune(scn.next())
	}

	tk.RawProps = []Prop{PR_REDUNDANT, PR_COMMENT}
	tk.Gen, tk.Sub, tk.RawStr = GEN_REDUNDANT, SUB_COMMENT, sb.String()
	return nil
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

	tk.RawProps = []Prop{PR_TERMINATOR, PR_NEWLINE}
	tk.Gen, tk.Sub, tk.RawStr = GEN_TERMINATOR, SUB_NEWLINE, sb.String()
	return nil
}

func whitespace(scn *scanner, tk *token.Tok) error {

	sb := strings.Builder{}

	for !scn.matchNewline() && scn.matchSpace() {
		sb.WriteRune(scn.next())
	}

	tk.RawProps = []Prop{PR_REDUNDANT, PR_WHITESPACE}
	tk.Gen, tk.Sub, tk.RawStr = GEN_REDUNDANT, SUB_WHITESPACE, sb.String()
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
		tk.RawProps = []Prop{PR_TERM, PR_LITERAL, PR_BOOL}
		tk.Gen, tk.Sub = GEN_LITERAL, SUB_BOOL

	default:
		tk.RawProps = []Prop{PR_TERM, PR_ASSIGNEE, PR_IDENTIFIER}
		tk.Gen, tk.Sub = GEN_IDENTIFIER, SUB_IDENTIFIER
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

	tk.RawProps = []Prop{PR_CALLABLE, PR_SPELL}
	tk.Gen, tk.Sub, tk.RawStr = GEN_SPELL, SUB_UNDEFINED, sb.String()
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

	tk.RawProps = []Prop{PR_TERM, PR_LITERAL, PR_STRING}
	tk.Gen, tk.Sub, tk.RawStr = GEN_LITERAL, SUB_STRING, sb.String()
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
	tk.RawProps = []Prop{PR_TERM, PR_LITERAL, PR_NUMBER}
	tk.Gen, tk.Sub, tk.RawStr = GEN_LITERAL, SUB_NUMBER, sb.String()
	return nil
}
