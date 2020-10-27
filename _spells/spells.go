package spells

import (
	"github.com/PaulioRandall/scarlet-go/spells/fmtr"
	"github.com/PaulioRandall/scarlet-go/spells/spellbook"
	"github.com/PaulioRandall/scarlet-go/spells/std"
)

func NewSpellbook() spellbook.Spellbook {

	spbk := spellbook.Spellbook{}
	checkError := func(e error) {
		if e != nil {
			panic(e)
		}
	}

	checkError(std.RegisterAll(spbk))
	checkError(fmtr.RegisterAll(spbk))

	return spbk
}
