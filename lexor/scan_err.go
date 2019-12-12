package lexor

import (
	"fmt"

	"github.com/PaulioRandall/scarlet-go/token"
)

// ScanErr represents an error while scanning.
type ScanErr interface {
	Error() string
	Where() token.Snippet
	Unwrap() error
	String() string
}

// stdScanErr is the standard ScanErr implementation.
type stdScanErr struct {
	what  string
	where token.Snippet
	why   error
}

// Error satisfies the Perror interface.
func (e stdScanErr) Error() string {
	return e.what
}

// Where satisfies the Perror interface.
func (e stdScanErr) Where() token.Snippet {
	return e.where
}

// Unwrap satisfies the Perror interface.
func (e stdScanErr) Unwrap() error {
	return e.why
}

// String returns a simple string representation of the error.
func (e stdScanErr) String() string {
	return fmt.Sprintf("%s at %s", e.what, e.where.String())
}

// NewScanErr returns a new instance of ScanErr.
func NewScanErr(what string, line, start, end int) ScanErr {
	return NewScanErr_2(what, token.NewSnippet(line, start, end))
}

// NewScanErr_2 returns a new instance of ScanErr.
func NewScanErr_2(what string, where token.Snippet) stdScanErr {
	return stdScanErr{
		what:  what,
		where: where,
	}
}

// WrapScanErr wraps an error in a ScanErr.
func WrapScanErr(what string, where token.Snippet, why error) ScanErr {
	return stdScanErr{
		what:  what,
		where: where,
		why:   why,
	}
}
