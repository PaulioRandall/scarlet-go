package errors

import (
	//"errors"

	"github.com/PaulioRandall/scarlet-go/scarlet/position"
)

type serror interface {
	error
	When() string
	From() position.Position
	To() position.Position
	Trace() string
}
