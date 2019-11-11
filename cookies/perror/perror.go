package perror

import (
	"fmt"

	w "github.com/PaulioRandall/scarlet-go/cookies/where"
)

// Perror represents an error within a file including its location.
type Perror interface {
	Error() string
	Where() w.Where
	Unwrap() error
	String() string
}

// perr is simple implementation of Perror.
type perr struct {
	what  string
	where w.Where
	why   error
}

// Error satisfies the Perror interface.
func (e perr) Error() string {
	return e.what
}

// Where satisfies the Perror interface.
func (e perr) Where() w.Where {
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

// New returns an instantiated implementation of Perror.
func New(what string, line, start, end int) Perror {
	return Newish(what, w.New(line, start, end))
}

// Newish returns an instantiated implementation of Perror.
func Newish(what string, where w.Where) Perror {
	return perr{
		what:  what,
		where: where,
	}
}

// Wrap wraps an error in an implementation of Perror.
func Wrap(what string, where w.Where, why error) Perror {
	return perr{
		what:  what,
		where: where,
		why:   why,
	}
}
