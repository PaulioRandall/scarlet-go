package alpha

import (
	"fmt"

	"github.com/PaulioRandall/scarlet-go/err"
	"github.com/PaulioRandall/scarlet-go/lexeme"
)

// statErr represents an error while partitioning statements.
type statErr struct {
	msg       string
	cause     error
	lineIndex int
	colIndex  int
	length    int
}

// newErr returns a new statement error.
func newErr(tk lexeme.Token, msg string, args ...interface{}) err.Err {
	return &statErr{
		msg:       fmt.Sprintf(msg, args...),
		lineIndex: tk.Line,
		colIndex:  tk.Col,
	}
}

// Error satisfies the error interface.
func (se statErr) Error() string {
	return se.msg
}

// Cause satisfies the Err interface.
func (se statErr) Cause() error {
	return se.cause
}

// LineIndex satisfies the Err interface.
func (se statErr) LineIndex() int {
	return se.lineIndex
}

// ColIndex satisfies the Err interface.
func (se statErr) ColIndex() int {
	return se.colIndex
}

// Length satisfies the Err interface.
func (se statErr) Length() int {
	return se.length
}
