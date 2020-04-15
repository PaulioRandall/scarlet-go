package recursive

import (
	"fmt"

	e "github.com/PaulioRandall/scarlet-go/pkg/err"
	"github.com/PaulioRandall/scarlet-go/pkg/token"
)

// parseErr represents an error while parsing.
type parseErr struct {
	msg       string
	cause     error
	lineIndex int
	colIndex  int
	length    int
}

// err returns a new parse error.
func err(f string, tk token.Token, msg string, args ...interface{}) e.Err {

	s := "[parser." + f + "] " + fmt.Sprintf(msg, args...)

	return &parseErr{
		msg:       s,
		lineIndex: tk.Line,
		colIndex:  tk.Col,
	}
}

func unexpected(f string, tk token.Token, expected token.TokenType) e.Err {
	return err(f, tk, "Expected %v, got %s", token.ANOTHER, tk.String())
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
