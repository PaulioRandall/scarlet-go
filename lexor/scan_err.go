package lexor

import (
	"fmt"
)

// ScanErr represents an error while scanning.
type ScanErr interface {
	error

	// Line returns the line where the error occurred
	Line() int

	// Col returns the line where the error occurred
	Col() int
}

// serr is the standard ScanErr implementation.
type serr struct {
	what string
	why  error
	line int
	col  int
}

// NewScanErr returns a new instance of ScanErr.
func NewScanErr(what string, why error, line, col int) ScanErr {
	return serr{
		what: what,
		why:  why,
		line: line,
		col:  col,
	}
}

// Error satisfies the error interface.
func (e serr) Error() string {

	s := fmt.Sprintf("%d:%d: %s", e.line, e.col, e.what)

	if e.why != nil {
		s += fmt.Sprintf("\n\t...caused by: %s", e.why.Error())
	}

	return s
}

// Line satisfies the ScanErr interface.
func (e serr) Line() int {
	return e.line
}

// Col satisfies the ScanErr interface.
func (e serr) Col() int {
	return e.col
}
