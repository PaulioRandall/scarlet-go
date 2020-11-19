package scanner

import (
	"fmt"

	"github.com/PaulioRandall/scarlet-go/scarlet/position"
)

type scanErr struct {
	msg string
	position.Range
}

// scanErr implements the error interface.

func (e scanErr) Error() string {
	return e.msg
}

func err(rng position.Range, msg string, args ...interface{}) scanErr {
	return scanErr{
		Range: rng,
		msg:   fmt.Sprintf(msg, args...),
	}
}
