package token

import (
	"fmt"
)

// Perror represents an error within a file including its location.
type Perror interface {
	Error() string
	Where() Snippet
	Unwrap() error
	String() string
}

// perr is simple implementation of Perror.
type perr struct {
	what  string
	where Snippet
	why   error
}

// Error satisfies the Perror interface.
func (e perr) Error() string {
	return e.what
}

// Where satisfies the Perror interface.
func (e perr) Where() Snippet {
	return e.where
}

// Unwrap satisfies the Perror interface.
func (e perr) Unwrap() error {
	return e.why
}

// String returns a simple string representation of the error.
func (e perr) String() string {
	return fmt.Sprintf("%s at %s", e.what, e.where.String())
}

// NewPerror returns a new instance of Perror.
func NewPerror(what string, line, start, end int) Perror {
	return PerrorBySnippet(what, NewSnippet(line, start, end))
}

// PerrorBySnippet returns a new instance of Perror.
func PerrorBySnippet(what string, where Snippet) Perror {
	return perr{
		what:  what,
		where: where,
	}
}

// WrapPerror wraps an error in a Perror.
func WrapPerror(what string, where Snippet, why error) Perror {
	return perr{
		what:  what,
		where: where,
		why:   why,
	}
}
