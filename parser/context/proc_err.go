package context

// ProcErr represents a error returned during the execution of a Scarlet
// function or spell.
type ProcErr interface {
	error

	// Cause returns the cause of the error if present.
	Why() error
}

// stdProcErr is the standard implementatin of a ProcErr.
type stdProcErr struct {
	why  error
	what string
}

// NewProcErr creates a new ProcErr. The caue may be null.
func NewProcErr(why error, what string) ProcErr {
	return stdProcErr{
		why:  why,
		what: what,
	}
}

// Error satisfies the error interface.
func (e stdProcErr) Error() string {
	return e.what
}

// Why satisfies the ProcErr interface.
func (e stdProcErr) Why() error {
	return e.why
}
