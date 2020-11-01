package spells

import (
	"github.com/PaulioRandall/scarlet-go/scarlet/spellbook"
)

// NewBook returns a new spell book inscribed with all spells.
func NewBook() spellbook.Book {

	const (
		NO_ARGS  = spellbook.NO_ARGS
		VAR_ARGS = spellbook.VAR_ARGS
	)

	b := spellbook.Book{}

	b.Inscribe("Exit", 1, NO_ARGS, Exit)
	b.Inscribe("Print", VAR_ARGS, NO_ARGS, Print)
	b.Inscribe("Println", VAR_ARGS, NO_ARGS, Println)

	return b
}
