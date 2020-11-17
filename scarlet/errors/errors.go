package errors

import (
	//"errors"

	"github.com/PaulioRandall/scarlet-go/scarlet/position"
)

type serror interface {
	error
	When() string
	From() position.Pos
	To() position.Pos
	Trace() string
}
