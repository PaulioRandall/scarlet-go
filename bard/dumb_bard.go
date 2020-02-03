package bard

import (
	"fmt"
)

// dumbBard is a simple implementation of the Bard interface that outputs
// information simplistically without showing its reference to the source code.
type dumbBard struct {
	Bard
}

// NewDumbBard create a simplistic outputter that won't link errors to their
// location within source code.
func NewDumbBard() Bard {
	return dumbBard{}
}

// CatchNightmare satisfies the Bard interface.
func (b dumbBard) CatchNightmare(f func()) {

	isPanic := true

	defer func() {

		if !isPanic {
			return
		}

		e := recover()
		b.printError(e)
	}()

	f()
	isPanic = false
}

// printError priunts an error.
func (b dumbBard) printError(e interface{}) {

	printNightmare := func(e error, at string, args ...interface{}) {

		println("[ERROR] " + e.Error())
		at = fmt.Sprintf(at, args...)
		println("[AT]    " + at)

		if n, ok := e.(Nightmare); ok && n.cause != nil {
			println("[CAUSE] " + n.cause.Error())
		}
	}

	switch err := e.(type) {
	case Horror:
		printNightmare(err, err.tk.String())
	case Terror:
		printNightmare(err, "%v:%v", err.line, err.col)
	case Nightmare:
		printNightmare(err, `¯\_(ツ)_/¯`)
	case error:
		printNightmare(err, `¯\_(ツ)_/¯`)
	default:
		panic("SANITY CHECK!" +
			" A panic was generated but the content was not an error of any sort",
		)
	}
}
