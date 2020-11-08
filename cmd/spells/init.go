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

	b.Inscribe("list.New", 1, List_New)
	b.Inscribe("list.Set", 0, List_Set)
	b.Inscribe("list.Prepend", 0, List_Prepend) // @Replace
	b.Inscribe("list.Append", 0, List_Append)   // @Replace
	b.Inscribe("list.Push", 0, List_Prepend)
	b.Inscribe("list.Pop", 1, List_Pop)
	b.Inscribe("list.Take", 1, List_Take)
	b.Inscribe("list.InRange", 1, List_InRange) // @Replace

	return b
}
