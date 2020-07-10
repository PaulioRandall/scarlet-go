package scan

import (
	"fmt"

	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/perror"
)

func errorUnknownSymbol(scn *scanner) error {
	return fail(scn.line, scn.col, "Unknown symbol %q", scn.buff)
}

func errorBadNewline(scn *scanner) error {
	return fail(scn.line, scn.col,
		"Got %q, expected %q", '\r', string("\r\n"))
}

func errorBadSpellName(scn *scanner) error {
	return fail(scn.line, scn.col, "Expected letter")
}

func errorBadString(scn *scanner) error {
	return fail(scn.line, scn.col,
		"Newline or EOF encountered before string was terminated")
}

func errorBadNumber(scn *scanner) error {
	return fail(scn.line, scn.col, "Expected digit after decimal point")
}

func fail(line, col int, msg string, args ...interface{}) error {
	msg = fmt.Sprintf(msg, args...)
	return perror.NewByPos(msg, line, col)
}

func failNow(msg string, args ...interface{}) {
	perror.Panic(msg, args...)
}
