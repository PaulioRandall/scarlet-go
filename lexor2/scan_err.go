package lexor

// ScanErr represents an error while scanning.
type ScanErr interface {
	error

	// Unwrap returns the underlying error or nil if there isn't one.
	Unwrap() error

	// Line returns the line where the error occurred
	Line() int

	// Col returns the line where the error occurred
	Col() int
}

// stdScanErr is the standard ScanErr implementation.
type stdScanErr struct {
	what string
	why  error
	line int
	col  int
}

// NewScanErr returns a new instance of ScanErr.
func NewScanErr(what string, why error, line, col int) ScanErr {
	return stdScanErr{
		what: what,
		why:  why,
		line: line,
		col:  col,
	}
}

// Error satisfies the error interface.
func (e stdScanErr) Error() string {
	return e.what
}

// Unwrap satisfies the ScanErr interface.
func (e stdScanErr) Unwrap() error {
	return e.why
}

// Line satisfies the ScanErr interface.
func (e stdScanErr) Line() int {
	return e.line
}

// Col satisfies the ScanErr interface.
func (e stdScanErr) Col() int {
	return e.col
}
