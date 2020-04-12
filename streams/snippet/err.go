package snippet

import (
	"fmt"

	"github.com/PaulioRandall/scarlet-go/err"
	"github.com/PaulioRandall/scarlet-go/lexeme"
)

// snippetErr represents an error with syntax.
type snippetErr struct {
	msg       string
	cause     error
	lineIndex int
	colIndex  int
	length    int
}

// newErr returns a new snippet error.
func newErr(tk lexeme.Token, msg string, args ...interface{}) err.Err {
	return &snippetErr{
		msg:       fmt.Sprintf(msg, args...),
		lineIndex: tk.Line,
		colIndex:  tk.Col,
	}
}

// Error satisfies the error interface.
func (se snippetErr) Error() string {
	return se.msg
}

// Cause satisfies the Err interface.
func (se snippetErr) Cause() error {
	return se.cause
}

// LineIndex satisfies the Err interface.
func (se snippetErr) LineIndex() int {
	return se.lineIndex
}

// ColIndex satisfies the Err interface.
func (se snippetErr) ColIndex() int {
	return se.colIndex
}

// Length satisfies the Err interface.
func (se snippetErr) Length() int {
	return se.length
}
