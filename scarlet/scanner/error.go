package scanner

import (
	"fmt"

	"github.com/PaulioRandall/scarlet-go/scarlet/position"
)

type scanErr struct {
	msg string
	position.Position
}

// scanErr implements the error interface.

func (e scanErr) Error() string {
	return e.msg
}

func err(p position.Position, msg string, args ...interface{}) scanErr {
	return scanErr{
		// TODO: Filename
		Position: p,
		msg:      fmt.Sprintf(msg, args...),
	}
}
