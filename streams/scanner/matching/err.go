package matching

import (
	"fmt"

	"github.com/PaulioRandall/scarlet-go/err"
	"github.com/PaulioRandall/scarlet-go/streams/symbol"
)

// scannerErr represents an error with syntax.
type scannerErr struct {
	msg       string
	lineIndex int
	colIndex  int
	length    int
}

// newErr returns a new scanner error.
func newErr(ts *symbol.TerminalStream, colOffset int, msg string, args ...interface{}) err.Err {
	return &scannerErr{
		lineIndex: ts.LineIndex(),
		colIndex:  ts.ColIndex() + colOffset,
		msg:       fmt.Sprintf(msg, args...),
	}
}

// Error satisfies the error interface.
func (se scannerErr) Error() string {
	return se.msg
}

// Cause satisfies the Err interface.
func (se scannerErr) Cause() error {
	return nil
}

// LineIndex satisfies the Err interface.
func (se scannerErr) LineIndex() int {
	return se.lineIndex
}

// ColIndex satisfies the Err interface.
func (se scannerErr) ColIndex() int {
	return se.colIndex
}

// Length satisfies the Err interface.
func (se scannerErr) Length() int {
	return se.length
}
