package spells

import (
	"github.com/PaulioRandall/scarlet-go/scarlet/spell"
)

// NewBook returns a new spell book inscribed with all spells.
func NewBook() spell.Book {

	b := spell.Book{}

	b.Inscribe("Len", 1, Len)
	b.Inscribe("Exit", 0, Exit)
	b.Inscribe("Print", 0, Print)
	b.Inscribe("Println", 0, Println)
	b.Inscribe("ParseNum", 2, ParseNum)
	b.Inscribe("PrintScope", 0, PrintScope)

	b.Inscribe("List.New", 1, List_New)
	b.Inscribe("List.Set", 0, List_Set)
	b.Inscribe("List.Get", 1, List_Get)
	b.Inscribe("List.Prepend", 0, List_Prepend)
	b.Inscribe("List.Append", 0, List_Append)

	return b
}
