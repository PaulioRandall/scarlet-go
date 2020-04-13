package err

import (
	"errors"
)

// Err represents an error .
type Err interface {
	Error() string

	Cause() error

	LineIndex() int

	ColIndex() int

	Length() int
}

// Try executes f. If f panics, a recovery is made and an Err representing the
// error is returned.
func Try(f func()) (err Err) {

	func() {
		defer func() {

			switch v := recover().(type) {
			case nil:
				err = nil
			case Err:
				err = v
			case string:
				err = goErr{errors.New(v)}
			case error:
				err = goErr{v}
			default:
				s := `¯\_(ツ)_/¯ Something panicked, but I don't understand the error`
				err = goErr{errors.New(s)}
			}

		}()

		f()
	}()

	return
}
