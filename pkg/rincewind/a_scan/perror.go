package scan

import (
	"fmt"

	"github.com/PaulioRandall/scarlet-go/pkg/rincewind/perror"
)

const (
	ERR_UNKNOWN_SYMBOL string = "SCAN_ERR_UNKNOWN_SYMBOL"
	ERR_BAD_NEWLINE    string = "SCAN_ERR_BAD_NEWLINE"
	ERR_BAD_SPELL_NAME string = "SCAN_ERR_BAD_SPELL_NAME"
	ERR_BAD_STRING     string = "SCAN_ERR_BAD_STRING"
	ERR_BAD_NUMBER     string = "SCAN_ERR_BAD_NUMBER"
)

func errorUnknownSymbol(scn *scanner) error {
	return fail(scn.line, scn.col, ERR_UNKNOWN_SYMBOL,
		"Unknown symbol %q", scn.buff)
}

func errorBadNewline(scn *scanner) error {
	return fail(scn.line, scn.col, ERR_BAD_NEWLINE,
		"Got %q, expected %q", '\r', string("\r\n"))
}

func errorBadSpellName(scn *scanner) error {
	return fail(scn.line, scn.col, ERR_BAD_SPELL_NAME,
		"Expected letter")
}

func errorBadString(scn *scanner) error {
	return fail(scn.line, scn.col, ERR_BAD_STRING,
		"Newline or EOF encountered before string was terminated")
}

func errorBadNumber(scn *scanner) error {
	return fail(scn.line, scn.col, ERR_BAD_NUMBER,
		"Expected digit after decimal point")
}

func fail(line, col int, code, msg string, args ...interface{}) error {
	msg = fmt.Sprintf(msg, args...)
	return perror.NewByPos(code, msg, line, col)
}
