package bard

import (
	"fmt"
)

// troubadour is a simple implementation of the Bard interface that outputs
// information simplistically without showing its reference to the source code.
type troubadour struct {
	Bard
}

// NewTroubadour create a simplistic outputter that won't link errors to their
// location within source code.
func NewTroubadour() Bard {
	return troubadour{}
}

// CatchNightmare satisfies the Bard interface.
func (t troubadour) CatchNightmare(f func()) {

	isPanic := true

	defer func() {

		if !isPanic {
			return
		}

		e := recover()
		t.printError(e)
	}()

	f()
	isPanic = false
}

// printError priunts an error.
func (_ troubadour) printError(e interface{}) {

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
