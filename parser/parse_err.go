package parser

import (
	"fmt"

	"github.com/PaulioRandall/scarlet-go/lexeme"
)

// parseErr represents an error with syntax.
type parseErr struct {
	msg       string
	cause     error
	lineIndex int
	colIndex  int
	length    int
}

// newErr returns a new parse error.
func newErr(msg string, args ...interface{}) *parseErr {
	return &parseErr{
		msg: fmt.Sprintf(msg, args...),
	}
}

// newTkErr returns a new parse error.
func newTkErr(tk lexeme.Token, msg string, args ...interface{}) *parseErr {
	return &parseErr{
		msg:       fmt.Sprintf(msg, args...),
		lineIndex: tk.Line,
		colIndex:  tk.Col,
	}
}

// newTkError returns a new parse error.
func newTkError(cause error, tk lexeme.Token, msg string, args ...interface{}) *parseErr {
	return &parseErr{
		msg:       fmt.Sprintf(msg, args...),
		cause:     cause,
		lineIndex: tk.Line,
		colIndex:  tk.Col,
	}
}

// Error satisfies the error interface.
func (pe parseErr) Error() string {
	return pe.msg
}

// Cause satisfies the Err interface.
func (pe parseErr) Cause() error {
	return pe.cause
}

// LineIndex satisfies the Err interface.
func (pe parseErr) LineIndex() int {
	return pe.lineIndex
}

// ColIndex satisfies the Err interface.
func (pe parseErr) ColIndex() int {
	return pe.colIndex
}

// Length satisfies the Err interface.
func (pe parseErr) Length() int {
	return pe.length
}
