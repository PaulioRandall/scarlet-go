package eval

// EvalErr represents a error returned during the evaluation of a Scarlet
// expression.
type EvalErr interface {
	error

	// Who returns the ID key of the range of tokens that represent the error
	// location in source.
	Who() int64

	// Why returns the cause of the error if presesnt.
	Why() error
}

type stdEvalErr struct {
	what string
	who  int64
	why  error
}

// NewEvalErr creates a new EvalErr. The `why` may be null.
func NewEvalErr(why error, who int64, what string) EvalErr {
	return stdEvalErr{
		what: what,
		who:  who,
		why:  why,
	}
}

// Error satisfies the error interface.
func (e stdEvalErr) Error() string {
	return e.what
}

// Who satisfies the EvalErr interface.
func (e stdEvalErr) Who() int64 {
	return e.who
}

// Why satisfies the EvalErr interface.
func (e stdEvalErr) Why() error {
	return e.why
}
