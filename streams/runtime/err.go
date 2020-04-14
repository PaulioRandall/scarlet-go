package runtime

import (
	"fmt"

	e "github.com/PaulioRandall/scarlet-go/err"
	"github.com/PaulioRandall/scarlet-go/lexeme"
)

// runtimeErr represents an error while executing statements or evaluating
// expressions.
type runtimeErr struct {
	msg       string
	cause     error
	lineIndex int
	colIndex  int
	length    int
}

// err returns a new runtime error.
func err(f string, tk lexeme.Token, msg string, args ...interface{}) e.Err {

	s := "[runtime." + f + "] " + fmt.Sprintf(msg, args...)

	return &runtimeErr{
		msg:       s,
		lineIndex: tk.Line,
		colIndex:  tk.Col,
	}
}

// Error satisfies the error interface.
func (re runtimeErr) Error() string {
	return re.msg
}

// Cause satisfies the Err interface.
func (re runtimeErr) Cause() error {
	return re.cause
}

// LineIndex satisfies the Err interface.
func (re runtimeErr) LineIndex() int {
	return re.lineIndex
}

// ColIndex satisfies the Err interface.
func (re runtimeErr) ColIndex() int {
	return re.colIndex
}

// Length satisfies the Err interface.
func (re runtimeErr) Length() int {
	return re.length
}
