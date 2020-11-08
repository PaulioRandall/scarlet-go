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

	b.Inscribe("list.New", 1, List_New)
	b.Inscribe("list.Set", 0, List_Set)
	b.Inscribe("list.Get", 1, List_Get)
	b.Inscribe("list.Prepend", 0, List_Prepend)
	b.Inscribe("list.Append", 0, List_Append)
	b.Inscribe("list.Push", 0, List_Prepend)
	b.Inscribe("list.Pop", 1, List_Pop)

	return b
}
