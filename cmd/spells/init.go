package spells

import (
	"github.com/PaulioRandall/scarlet-go/scarlet/spell"
)

// NewBook returns a new spell book inscribed with all spells.
func NewBook() spell.Book {

	b := spell.Book{}

	b.Inscribe("Exit", 0, Exit)
	b.Inscribe("Print", 0, Print)
	b.Inscribe("Println", 0, Println)
	b.Inscribe("ParseNum", 2, ParseNum)
	b.Inscribe("PrintScope", 0, PrintScope)
	b.Inscribe("Len", 1, Len)
	b.Inscribe("Slice", 1, Slice)
	b.Inscribe("At", 1, At)
	b.Inscribe("InRange", 1, InRange)
	b.Inscribe("Prepend", 1, Prepend)
	b.Inscribe("Append", 1, Append)

	b.Inscribe("NewList", 1, NewList)
	b.Inscribe("Set", 1, Set)
	b.Inscribe("Push", 1, Prepend)
	b.Inscribe("Pop", 1, Pop)
	b.Inscribe("Take", 1, Take)

	return b
}
