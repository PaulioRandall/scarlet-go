package spells

import (
	"github.com/PaulioRandall/scarlet-go/scarlet/spell"
)

// NewBook returns a new spell book inscribed with all spells.
func NewBook() spell.Book {

	const (
		NO_ARGS  = spell.NO_ARGS
		VAR_ARGS = spell.VAR_ARGS
	)

	b := spell.Book{}

	b.Inscribe("Exit", 1, NO_ARGS, Exit)
	b.Inscribe("Print", VAR_ARGS, NO_ARGS, Print)
	b.Inscribe("Println", VAR_ARGS, NO_ARGS, Println)

	return b
}
